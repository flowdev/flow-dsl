package parser

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/flowdev/flow-dsl/data"

	"github.com/antlr4-go/antlr/v4"
)

type LinkGenerator interface {
	GenerateLink(imprt, typ string, isData bool) (link string, isLinkToFlow bool)
}

func ConvertFlowFileToData(flowFile IFlowFileContext, linkGenerator LinkGenerator) (*data.FlowFile, error) {
	antlrFlows := flowFile.GetFlows()
	dataFlowFile := &data.FlowFile{}
	dataFlowFile.Imports = convertImportsToData(flowFile.GetAllImports())
	dataFlows, err := convertFlowsToData(antlrFlows, dataFlowFile.Imports, linkGenerator)
	dataFlowFile.Flows = dataFlows
	return dataFlowFile, err
}

type dataConverter struct {
	linkGenerator LinkGenerator
	imports       map[string]string
	compMap       map[string]compLink
	compLinkTable [][]compLink
	errs          []error
}

type compLink struct {
	stmtIdx,
	compIdx int
	comp *data.Comp
}

func convertImportsToData(imports IImportsContext) []string {
	if imports == nil {
		return nil
	}
	antlrImports := imports.GetImports()
	dataImports := make([]string, len(antlrImports))
	for i := 0; i < len(antlrImports); i++ {
		dataImports[i] = convertImportToData(antlrImports[i])
	}
	return dataImports
}

func convertImportToData(imp antlr.Token) string {
	fullImp := imp.GetText()
	l := len(fullImp)
	return fullImp[1 : l-1] // cut off leading and trailing quote (")
}

func convertFlowsToData(antlrFlows []IFlowContext, dataImports []string, linkGenerator LinkGenerator) ([]*data.Flow, error) {
	dc := &dataConverter{linkGenerator: linkGenerator}
	dc.imports = dc.convertDataImportsToMap(dataImports)
	dataFlows := make([]*data.Flow, len(antlrFlows))
	for i, antlrFlow := range antlrFlows {
		dataFlows[i] = dc.convertFlowToData(antlrFlow)
	}
	return dataFlows, errors.Join(dc.errs...)
}

func (dc *dataConverter) convertDataImportsToMap(dataImports []string) map[string]string {
	importMap := make(map[string]string)
	for _, dataImport := range dataImports {
		i := strings.LastIndexFunc(dataImport, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' && r != '-'
		})
		key := dataImport
		if i >= len(dataImport)-1 {
			key = "missingImportPart"
			dc.errs = append(dc.errs, fmt.Errorf("import is missing a final part: %q", dataImport))
		} else if i >= 0 {
			key = dataImport[i+1:]
		}
		importMap[key] = dataImport
	}
	return importMap
}

func (dc *dataConverter) convertFlowToData(antlrFlow IFlowContext) *data.Flow {
	dataFlow := &data.Flow{}
	dataFlow.Name = dc.convertFlowNameToData(antlrFlow.GetName())
	antlrStatements := antlrFlow.GetStatements()
	dc.compMap = make(map[string]compLink)
	dc.compLinkTable = make([][]compLink, len(antlrStatements))
	dataFlow.StatementStarts = make([]*data.StartComp, len(antlrStatements))
	for i := 0; i < len(antlrStatements); i++ {
		dataFlow.StatementStarts[i] = dc.convertStatementToData(antlrStatements[i], i)
	}
	dataFlow.StatementStarts = cleanFlowStatements(dataFlow.StatementStarts)
	return dataFlow
}

func cleanFlowStatements(stmts []*data.StartComp) []*data.StartComp {
	cleanStmts := make([]*data.StartComp, 0, len(stmts))
	for _, stmt := range stmts {
		if stmt.Arrow != nil || stmt.Comp != nil {
			cleanStmts = append(cleanStmts, stmt)
		}
	}
	return cleanStmts
}

func (dc *dataConverter) convertFlowNameToData(antlrFlowName antlr.Token) string {
	return antlrFlowName.GetText()
}

func (dc *dataConverter) convertStatementToData(antlrStatement IStatementContext, stmtIdx int) *data.StartComp {
	dc.compLinkTable[stmtIdx] = make([]compLink, 0, 32)
	dataStartComp := &data.StartComp{}
	dc.convertStatementStartToData(antlrStatement.StatementStart(), dataStartComp)
	dataEndComp := dc.convertStatementMiddleToData(antlrStatement.StatementMiddle(), dataStartComp, stmtIdx)
	dc.convertStatementEndToData(antlrStatement.StatementEnd(), dataEndComp)
	return dataStartComp
}

func (dc *dataConverter) convertStatementStartToData(antlrStatementStart IStatementStartContext, dataStartComp *data.StartComp) {
	if antlrStatementStart == nil {
		return
	}
	dataStartComp.PortName = antlrStatementStart.GetStartPort().GetText()
	dataArrow := &data.Arrow{}
	if antlrStatementStart.GetDstPort() != nil {
		dataArrow.DstPort = antlrStatementStart.GetDstPort().GetText()
	}
	dataArrow.DataTypes = dc.convertArrowDataToData(antlrStatementStart.GetAllArrData())
	dataStartComp.Arrow = dataArrow
}

func (dc *dataConverter) convertStatementMiddleToData(antlrMidComp IStatementMiddleContext, dataStartComp *data.StartComp, stmtIdx int) (dataEndComp *data.EndComp) {
	dataComp := dc.convertComponentToData(antlrMidComp.GetFirstComponent(), stmtIdx, 0)
	dataEndComp = &data.EndComp{Comp: dataComp}
	if dataStartComp.Arrow != nil {
		dataStartComp.Arrow.DstComp = dataEndComp
		dataComp.Inputs = append(dataComp.Inputs, dataStartComp.Arrow)
	} else if dc.compLinkTable[stmtIdx][0].comp == nil {
		dataStartComp.Comp = dataComp
	}
	antlrArrComps := antlrMidComp.GetArrowComponents()
	compIdx := 1
	for _, antlrArrComp := range antlrArrComps {
		dataEndComp = dc.convertArrowComponent(dataEndComp, antlrArrComp, stmtIdx, compIdx)
		compIdx++
	}
	return dataEndComp
}

func (dc *dataConverter) convertArrowComponent(dataSrcComp *data.EndComp, antlrArrComp IArrowComponentContext, stmtIdx, compIdx int) *data.EndComp {
	dataArrow := &data.Arrow{}
	dataArrow.SrcComp = &data.StartComp{Comp: dataSrcComp.Comp}
	dataSrcComp.Comp.Outputs = append(dataSrcComp.Comp.Outputs, dataArrow)
	if antlrArrComp.GetSrcPort() != nil {
		dataArrow.SrcPort = antlrArrComp.GetSrcPort().GetText()
	}
	dataArrow.DataTypes = dc.convertArrowDataToData(antlrArrComp.GetAllArrData())
	if antlrArrComp.GetDstPort() != nil {
		dataArrow.DstPort = antlrArrComp.GetDstPort().GetText()
	}
	dataDstComp := dc.convertComponentToData(antlrArrComp.Component(), stmtIdx, compIdx)
	dataArrow.DstComp = &data.EndComp{Comp: dataDstComp}
	dataDstComp.Inputs = append(dataDstComp.Inputs, dataArrow)
	return dataArrow.DstComp
}

func (dc *dataConverter) convertStatementEndToData(antlrStatementEnd IStatementEndContext, dataEndComp *data.EndComp) {
	if antlrStatementEnd == nil {
		return
	}
	dataLastComp := dataEndComp.Comp
	dataEndComp = &data.EndComp{} // we create a new end
	dataEndComp.PortName = antlrStatementEnd.GetEndPort().GetText()
	dataArrow := &data.Arrow{}
	if antlrStatementEnd.GetSrcPort() != nil {
		dataArrow.SrcPort = antlrStatementEnd.GetSrcPort().GetText()
	}
	dataArrow.DataTypes = dc.convertArrowDataToData(antlrStatementEnd.GetAllArrData())
	dataArrow.SrcComp = &data.StartComp{Comp: dataLastComp}
	dataLastComp.Outputs = append(dataLastComp.Outputs, dataArrow)
	dataArrow.DstComp = dataEndComp
	dataEndComp.Arrow = dataArrow
}

func (dc *dataConverter) convertComponentToData(antlrComp IComponentContext, stmtIdx, compIdx int) *data.Comp {
	dataComp := &data.Comp{}
	antlrCompCore := antlrComp.GetCore()
	dataComp.Name = antlrCompCore.GetName().GetText()
	dataComp.Typ = convertPackageTypeToString(antlrCompCore.GetTypPack(), antlrCompCore.GetTypName())
	if dataComp.Typ == "" {
		dataComp.Typ = dataComp.Name
	}
	dataComp.PluginGroups = dc.convertAllPluginGroupsToData(antlrComp.GetAllPlugins())
	if dataComp.Typ == "" {
		dataComp.Typ = dataComp.Name
	}

	dc.compLinkTable[stmtIdx] = append(dc.compLinkTable[stmtIdx], compLink{})
	newCompLink := compLink{stmtIdx: stmtIdx, compIdx: compIdx, comp: dataComp}
	existingComp := dc.compMap[dataComp.Name]
	if dataComp.Name != dataComp.Typ || len(dataComp.PluginGroups) > 0 { // must be a new component
		if existingComp.comp != nil {
			dc.errs = append(dc.errs, fmt.Errorf(
				"component with name %q exists already at statement %d, component %d and now at statement %d, component %d",
				dataComp.Name, existingComp.stmtIdx+1, existingComp.compIdx+1, stmtIdx+1, compIdx+1))
		} else {
			dc.compMap[dataComp.Name] = newCompLink
			dataComp.Link, _ = dc.createLink(dataComp.Typ, false)
		}
	} else {
		if existingComp.comp != nil {
			if dc.isJumpComponent(newCompLink, existingComp) {
				dataComp.IsJump = true
				dataComp.JumpPort = "" // TODO: fill port from arrow
				dataComp.Link, _ = dc.createLink(dataComp.Typ, false)
			} else {
				dc.compLinkTable[stmtIdx][compIdx] = existingComp
				dataComp = existingComp.comp
			}
		} else {
			dc.compMap[dataComp.Name] = newCompLink
			dataComp.Link, _ = dc.createLink(dataComp.Typ, false)
		}
	}

	return dataComp
}

func (dc *dataConverter) isJumpComponent(dataComp, existingComp compLink) bool {
	if dataComp.stmtIdx <= existingComp.stmtIdx {
		return true
	}
	// TODO: look for compExtensions in own and previous lines
	//       - the higher they go (lower stmtIdx in link), the more important
	//       - the later in the statement the compExtension, the more important
	//       - these extensions form barriers we can't cross
	return false
}

func (dc *dataConverter) findCompLinkMax(stmtIdx, maxCompIdx int) compLink {
	compLinkRow := dc.compLinkTable[stmtIdx]
	for i := maxCompIdx; i >= 0; i-- {
		if compLinkRow[i].comp != nil {
			return compLinkRow[i]
		}
	}
	return compLink{stmtIdx: stmtIdx, compIdx: -1}
}

func (dc *dataConverter) convertAllPluginGroupsToData(antlrAllPluginGroups IPluginContext) []data.PluginGroup {
	if antlrAllPluginGroups == nil {
		return nil
	}
	antlrPluginGroups := antlrAllPluginGroups.GetPluginGroups()
	dataPluginGroups := make([]data.PluginGroup, len(antlrPluginGroups))
	for i := 0; i < len(antlrAllPluginGroups.GetPluginGroups()); i++ {
		dataPluginGroups[i] = dc.convertPluginGroup(antlrPluginGroups[i])
	}
	return dataPluginGroups
}

func (dc *dataConverter) convertPluginGroup(antlrPluginGroup IPluginPartContext) data.PluginGroup {
	dataPluginGroup := data.PluginGroup{}
	dataPluginGroup.Interface = convertPackageTypeToString(antlrPluginGroup.GetInterface().GetID1(), antlrPluginGroup.GetInterface().GetID2())
	dataPluginGroup.Plugins = make([]data.Plugin, len(antlrPluginGroup.GetPlugins()))
	for i := 0; i < len(antlrPluginGroup.GetPlugins()); i++ {
		antlrPlugin := antlrPluginGroup.GetPlugins()[i]
		dataPluginGroup.Plugins[i] = data.Plugin{Typ: convertPackageTypeToString(antlrPlugin.GetID1(), antlrPlugin.GetID2())}
	}
	return dataPluginGroup
}

func (dc *dataConverter) convertArrowDataToData(antlrArrowData IAllDataContext) []data.DataType {
	if antlrArrowData == nil {
		return nil
	}
	dataArrowData := make([]data.DataType, len(antlrArrowData.GetDatas()))
	for i := 0; i < len(antlrArrowData.GetDatas()); i++ {
		dataArrowData[i] = dc.convertDataToData(antlrArrowData.GetDatas()[i])
	}
	return dataArrowData
}

func (dc *dataConverter) convertDataToData(antlrData IDataContext) data.DataType {
	dataData := data.DataType{
		Name: antlrData.GetName().GetText(),
		Typ:  convertPackageTypeToString(antlrData.GetTypPack(), antlrData.GetTypName()),
	}
	dataData.Link, _ = dc.createLink(dataData.Typ, true)
	return dataData
}

func (dc *dataConverter) createLink(typ string, isData bool) (link string, isFlowLink bool) {
	lastImportPart := ""
	dotIdx := strings.IndexByte(typ, '.')
	if dotIdx > 0 {
		lastImportPart = typ[:dotIdx]
	}
	imprt := ""
	if lastImportPart != "" {
		imprt = dc.imports[lastImportPart]
	}
	return dc.linkGenerator.GenerateLink(imprt, typ, isData)
}

func convertPackageTypeToString(antlrPackage, antlrType antlr.Token) string {
	dataPackage := ""
	dataType := ""
	if antlrPackage != nil {
		dataPackage = antlrPackage.GetText()
		if antlrType != nil {
			dataType = antlrType.GetText()
		}
	}

	if dataPackage != "" && dataType != "" {
		return dataPackage + "." + dataType
	}
	return dataPackage // this is really the type
}
