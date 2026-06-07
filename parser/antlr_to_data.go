package parser

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/flowdev/flow-dsl/data"
)

func ConvertFlowFileToData(flowFile IFlowFileContext) (*data.FlowFile, error) {
	antlrFlows := flowFile.GetFlows()
	dc := &dataConverter{compMap: make(map[string]compLink), compLinkTable: make([][]compLink, len(antlrFlows))}
	dataFlowFile := &data.FlowFile{}
	dataFlowFile.Imports = dc.convertImportsToData(flowFile.GetAllImports())
	dataFlowFile.Flows = dc.convertFlowsToData(antlrFlows)
	return dataFlowFile, errors.Join(dc.errs...)
}

type dataConverter struct {
	compMap       map[string]compLink
	compLinkTable [][]compLink
	errs          []error
}

type compLink struct {
	stmtIdx,
	compIdx int
	isLoop bool
	comp   *data.Comp
}

func (dc *dataConverter) convertImportsToData(imports IImportsContext) []string {
	antlrImports := imports.GetImports()
	dataImports := make([]string, len(antlrImports))
	for i := 0; i < len(antlrImports); i++ {
		dataImports[i] = dc.convertImportToData(antlrImports[i])
	}
	return dataImports
}

func (dc *dataConverter) convertImportToData(imp antlr.Token) string {
	fullImp := imp.GetText()
	l := len(fullImp)
	return fullImp[1 : l-1] // cut of leading and trailing quote (")
}

func (dc *dataConverter) convertFlowsToData(antlrFlows []IFlowContext) []*data.Flow {
	dataFlows := make([]*data.Flow, len(antlrFlows))
	for i := 0; i < len(antlrFlows); i++ {
		dataFlows[i] = dc.convertFlowToData(antlrFlows[i])
	}
	return dataFlows
}

func (dc *dataConverter) convertFlowToData(antlrFlow IFlowContext) *data.Flow {
	dataFlow := &data.Flow{}
	dataFlow.Name = dc.convertFlowNameToData(antlrFlow.GetName())
	antlrStatements := antlrFlow.GetStatements()
	dataFlow.StatementStarts = make([]*data.StartComp, len(antlrStatements))
	for i := 0; i < len(antlrStatements); i++ {
		dataFlow.StatementStarts[i] = dc.convertStatementToData(antlrStatements[i], i)
	}
	return dataFlow
}

func (dc *dataConverter) convertFlowNameToData(antlrFlowName antlr.Token) string {
	return antlrFlowName.GetText()
}

func (dc *dataConverter) convertStatementToData(antlrStatement IStatementContext, stmtIdx int) *data.StartComp {
	dc.compLinkTable[stmtIdx] = make([]*compLink, 0, 32)
	dataStartComp := &data.StartComp{}
	dc.convertStatementStartToData(antlrStatement.StatementStart(), dataStartComp)
	dataEndComp := dc.convertStatementMiddleToData(antlrStatement.StatementMiddle(), dataStartComp, stmtIdx)
	dc.convertStatementEndToData(antlrStatement.StatementEnd(), dataEndComp)
	return dataStartComp
}

func (dc *dataConverter) convertStatementStartToData(antlrStatementStart IStatementStartContext, dataStartComp *data.StartComp) {
	if antlrStatementStart != nil {
		dataStartComp.PortName = antlrStatementStart.GetStartPort().GetText()
		dataArrow := &data.Arrow{}
		dataArrow.DstPort = antlrStatementStart.GetDstPort().GetText()
		dataArrow.DataTypes = dc.convertArrowDataToData(antlrStatementStart.GetAllArrData())
		dataStartComp.Arrow = dataArrow
	}
}

func (dc *dataConverter) convertStatementMiddleToData(antlrMidComp IStatementMiddleContext, dataStartComp *data.StartComp, stmtIdx int) (dataEndComp *data.EndComp) {
	dataComp := dc.convertComponentToData(antlrMidComp.GetFirstComponent(), stmtIdx, 0)
	dataStartComp.Comp = dataComp
	dataEndComp = &data.EndComp{Comp: dataComp}
	dataStartComp.Arrow.DstComp = dataEndComp
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
	dataArrow.SrcPort = antlrArrComp.GetSrcPort().GetText()
	dataArrow.DataTypes = dc.convertArrowDataToData(antlrArrComp.GetAllArrData())
	dataArrow.DstPort = antlrArrComp.GetDstPort().GetText()
	dataDstComp := dc.convertComponentToData(antlrArrComp.Component(), stmtIdx, compIdx)
	dataArrow.DstComp = &data.EndComp{Comp: dataDstComp}
	dataDstComp.Inputs = append(dataDstComp.Inputs, dataArrow)
	return dataArrow.DstComp
}

func (dc *dataConverter) convertStatementEndToData(antlrStatementEnd IStatementEndContext, dataEndComp *data.EndComp) {
	if antlrStatementEnd != nil {
		dataLastComp := dataEndComp.Comp
		dataEndComp = &data.EndComp{} // we create a new end
		dataEndComp.PortName = antlrStatementEnd.GetEndPort().GetText()
		dataArrow := &data.Arrow{}
		dataArrow.SrcPort = antlrStatementEnd.GetSrcPort().GetText()
		dataArrow.DataTypes = dc.convertArrowDataToData(antlrStatementEnd.GetAllArrData())
		dataArrow.SrcComp = &data.StartComp{Comp: dataLastComp}
		dataArrow.DstComp = dataEndComp
		dataEndComp.Arrow = dataArrow
		dataLastComp.Outputs = append(dataLastComp.Outputs, dataArrow)
	}
}

func (dc *dataConverter) convertComponentToData(antlrComp IComponentContext, stmtIdx, compIdx int) *data.Comp {
	dataComp := &data.Comp{}
	antlrCompCore := antlrComp.GetCore()
	dataComp.Name = antlrCompCore.GetName().GetText()
	dataComp.Typ = convertPackageTypeToString(antlrCompCore.GetTypPack(), antlrCompCore.GetTypName())
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
		}
	} else {
		if existingComp.comp != nil {
			if !dc.isJumpComponent(newCompLink, existingComp) {
				dc.compLinkTable[stmtIdx][compIdx] = existingComp
			} else {
				// TODO: Change dataComp to Jump
			}
		} else {
			dc.compMap[dataComp.Name] = newCompLink
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
	dataArrowData := make([]data.DataType, len(antlrArrowData.GetDatas()))
	for i := 0; i < len(antlrArrowData.GetDatas()); i++ {
		dataArrowData[i] = dc.convertDataToData(antlrArrowData.GetDatas()[i])
	}
	return dataArrowData
}

func (dc *dataConverter) convertDataToData(antlrData IDataContext) data.DataType {
	return data.DataType{
		Name: antlrData.GetName().GetText(),
		Typ:  convertPackageTypeToString(antlrData.GetTypPack(), antlrData.GetTypName()),
	}
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
