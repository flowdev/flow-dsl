package parser

import (
	"github.com/flowdev/flow-dsl/data"

	"github.com/antlr4-go/antlr/v4"
)

func ParseFile(fileName string, linkGenerator LinkGenerator) (*data.FlowFile, error) {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		return nil, err
	}
	lexer := NewflowDslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewflowDslParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	flowFile := p.FlowFile()

	return ConvertFlowFileToData(flowFile, linkGenerator)
}
