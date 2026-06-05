package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/flowdev/flow-dsl/data"
)

func ConvertFlowFileToData(flowFile IFlowFileContext) (*data.FlowFile, error) {
	antlrFlows := flowFile.GetFlows()
	dc := &dataConverter{compMap: make(map[string]*data.Comp), compLinkTable: make([][]*compLink, len(antlrFlows))}
	dataFlowFile := &data.FlowFile{}
	dataFlowFile.Imports = dc.convertImportsToData(flowFile.GetAllImports())
	dataFlowFile.Flows = dc.convertFlowsToData(antlrFlows)

	return nil, nil
}

type dataConverter struct {
	compMap       map[string]*data.Comp
	compLinkTable [][]*compLink
}

type compLink struct {
	statIdx, compIdx int
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
	dc.convertStartPortToData(antlrStatement.StatementStart(), dataStartComp)
	dc.convertMidCompToData(antlrStatement.StatementMiddle(), dataStartComp)
	return dataStartComp
}

func (dc *dataConverter) convertStartPortToData(antlrStatementStart IStatementStartContext, dataStartComp *data.StartComp) {
	if antlrStatementStart != nil {
		dataStartComp.PortName = antlrStatementStart.GetStartPort().GetText()
		dataArrow := &data.Arrow{}
		dataArrow.DstPort = antlrStatementStart.GetDstPort().GetText()
		dataArrow.DataTypes = dc.convertArrowDataToData(antlrStatementStart.GetAllArrData())
		dataStartComp.Output = dataArrow
	}
}

func (dc *dataConverter) convertMidCompToData(antlrMidComp IStatementMiddleContext, dataStartComp *data.StartComp) {
	panic("unimplemented")
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
		Typ:  antlrData.GetTypPack().GetText() + "." + antlrData.GetTypName().GetText(),
	}
}
