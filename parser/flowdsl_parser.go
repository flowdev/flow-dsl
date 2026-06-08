// Code generated from ./flowDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // flowDsl

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type flowDslParser struct {
	*antlr.BaseParser
}

var FlowDslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func flowdslParserInit() {
	staticData := &FlowDslParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'import'", "'flow'", "'{'", "", "", "", "", "", "", "':'",
		"", "", "'('", "", "'}'", "", "", "", "", "", "", "')'", "", "", "",
		"", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "SEMI", "IMPORT", "FLOW", "LBRACE", "ID", "BLOCK_COMMENT", "LINE_COMMENT",
		"WS", "ARROW", "SEMIF", "COLON", "STRING", "PORT", "LPAREN", "LBRACK",
		"RBRACE", "BLOCK_COMMENTF", "LINE_COMMENTF", "WSF", "COMMA", "PIPE",
		"DOTI", "RPAREN", "LBRACKI", "RBRACK", "IDI", "BLOCK_COMMENTI", "LINE_COMMENTI",
		"WSI", "ASSIGN", "COMMAP", "DOTP", "PIPEP", "RBRACKP", "IDP", "BLOCK_COMMENTP",
		"LINE_COMMENTP", "WSP",
	}
	staticData.RuleNames = []string{
		"flowFile", "imports", "flow", "statement", "statementStart", "statementMiddle",
		"statementEnd", "arrowComponent", "component", "componentTypeName",
		"plugin", "pluginPart", "allData", "data", "packageIDI", "packageIDP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 171, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 3, 0, 34, 8, 0, 1, 0, 4, 0, 37, 8, 0, 11, 0, 12, 0, 38, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 47, 8, 1, 11, 1, 12, 1, 48, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 57, 8, 2, 11, 2, 12, 2, 58, 1, 2, 1, 2, 1,
		3, 3, 3, 64, 8, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 76, 8, 4, 1, 5, 1, 5, 5, 5, 80, 8, 5, 10, 5, 12, 5, 83,
		9, 5, 1, 6, 3, 6, 86, 8, 6, 1, 6, 3, 6, 89, 8, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 3, 7, 95, 8, 7, 1, 7, 3, 7, 98, 8, 7, 1, 7, 1, 7, 3, 7, 102, 8, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 109, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 117, 8, 9, 3, 9, 119, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 125, 8, 10, 10, 10, 12, 10, 128, 9, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 138, 8, 11, 10, 11, 12, 11, 141,
		9, 11, 3, 11, 143, 8, 11, 1, 12, 1, 12, 3, 12, 147, 8, 12, 1, 12, 1, 12,
		5, 12, 151, 8, 12, 10, 12, 12, 12, 154, 9, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 3, 13, 162, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		3, 15, 169, 8, 15, 1, 15, 1, 81, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 0, 0, 177, 0, 33, 1, 0, 0, 0, 2, 42, 1, 0,
		0, 0, 4, 52, 1, 0, 0, 0, 6, 63, 1, 0, 0, 0, 8, 71, 1, 0, 0, 0, 10, 77,
		1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 94, 1, 0, 0, 0, 16, 105, 1, 0, 0, 0,
		18, 112, 1, 0, 0, 0, 20, 120, 1, 0, 0, 0, 22, 142, 1, 0, 0, 0, 24, 144,
		1, 0, 0, 0, 26, 157, 1, 0, 0, 0, 28, 163, 1, 0, 0, 0, 30, 165, 1, 0, 0,
		0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36,
		1, 0, 0, 0, 35, 37, 3, 4, 2, 0, 36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 5,
		0, 0, 1, 41, 1, 1, 0, 0, 0, 42, 43, 5, 2, 0, 0, 43, 46, 5, 4, 0, 0, 44,
		45, 5, 12, 0, 0, 45, 47, 5, 10, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 1, 0,
		0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51,
		5, 16, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 5, 3, 0, 0, 53, 54, 5, 5, 0, 0,
		54, 56, 5, 4, 0, 0, 55, 57, 3, 6, 3, 0, 56, 55, 1, 0, 0, 0, 57, 58, 1,
		0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60,
		61, 5, 16, 0, 0, 61, 5, 1, 0, 0, 0, 62, 64, 3, 8, 4, 0, 63, 62, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 3, 10, 5, 0, 66, 68,
		3, 12, 6, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 70, 5, 10, 0, 0, 70, 7, 1, 0, 0, 0, 71, 72, 5, 13, 0, 0, 72, 73, 3,
		24, 12, 0, 73, 75, 5, 9, 0, 0, 74, 76, 5, 13, 0, 0, 75, 74, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 81, 3, 16, 8, 0, 78, 80, 3,
		14, 7, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 11, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 5, 13,
		0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 89,
		3, 24, 12, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 91, 5, 9, 0, 0, 91, 92, 5, 13, 0, 0, 92, 13, 1, 0, 0, 0, 93, 95,
		5, 13, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0,
		96, 98, 3, 24, 12, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 101, 5, 9, 0, 0, 100, 102, 5, 13, 0, 0, 101, 100, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 3, 16, 8, 0, 104,
		15, 1, 0, 0, 0, 105, 106, 5, 15, 0, 0, 106, 108, 3, 18, 9, 0, 107, 109,
		3, 20, 10, 0, 108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1,
		0, 0, 0, 110, 111, 5, 25, 0, 0, 111, 17, 1, 0, 0, 0, 112, 118, 5, 26, 0,
		0, 113, 116, 5, 26, 0, 0, 114, 115, 5, 22, 0, 0, 115, 117, 5, 26, 0, 0,
		116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0, 118,
		113, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 19, 1, 0, 0, 0, 120, 121, 5,
		24, 0, 0, 121, 126, 3, 22, 11, 0, 122, 123, 5, 33, 0, 0, 123, 125, 3, 22,
		11, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0,
		126, 127, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129,
		130, 5, 34, 0, 0, 130, 21, 1, 0, 0, 0, 131, 143, 3, 30, 15, 0, 132, 133,
		3, 30, 15, 0, 133, 134, 5, 30, 0, 0, 134, 139, 3, 30, 15, 0, 135, 136,
		5, 31, 0, 0, 136, 138, 3, 30, 15, 0, 137, 135, 1, 0, 0, 0, 138, 141, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 143, 1, 0, 0,
		0, 141, 139, 1, 0, 0, 0, 142, 131, 1, 0, 0, 0, 142, 132, 1, 0, 0, 0, 143,
		23, 1, 0, 0, 0, 144, 146, 5, 14, 0, 0, 145, 147, 3, 26, 13, 0, 146, 145,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 152, 1, 0, 0, 0, 148, 149, 5, 20,
		0, 0, 149, 151, 3, 26, 13, 0, 150, 148, 1, 0, 0, 0, 151, 154, 1, 0, 0,
		0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 155, 156, 5, 23, 0, 0, 156, 25, 1, 0, 0, 0, 157, 158,
		5, 26, 0, 0, 158, 161, 5, 26, 0, 0, 159, 160, 5, 22, 0, 0, 160, 162, 5,
		26, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 27, 1, 0, 0,
		0, 163, 164, 5, 26, 0, 0, 164, 29, 1, 0, 0, 0, 165, 168, 5, 35, 0, 0, 166,
		167, 5, 32, 0, 0, 167, 169, 5, 35, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169,
		1, 0, 0, 0, 169, 31, 1, 0, 0, 0, 23, 33, 38, 48, 58, 63, 67, 75, 81, 85,
		88, 94, 97, 101, 108, 116, 118, 126, 139, 142, 146, 152, 161, 168,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// flowDslParserInit initializes any static state used to implement flowDslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewflowDslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FlowDslParserInit() {
	staticData := &FlowDslParserStaticData
	staticData.once.Do(flowdslParserInit)
}

// NewflowDslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewflowDslParser(input antlr.TokenStream) *flowDslParser {
	FlowDslParserInit()
	this := new(flowDslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FlowDslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "flowDsl.g4"

	return this
}

// flowDslParser tokens.
const (
	flowDslParserEOF            = antlr.TokenEOF
	flowDslParserSEMI           = 1
	flowDslParserIMPORT         = 2
	flowDslParserFLOW           = 3
	flowDslParserLBRACE         = 4
	flowDslParserID             = 5
	flowDslParserBLOCK_COMMENT  = 6
	flowDslParserLINE_COMMENT   = 7
	flowDslParserWS             = 8
	flowDslParserARROW          = 9
	flowDslParserSEMIF          = 10
	flowDslParserCOLON          = 11
	flowDslParserSTRING         = 12
	flowDslParserPORT           = 13
	flowDslParserLPAREN         = 14
	flowDslParserLBRACK         = 15
	flowDslParserRBRACE         = 16
	flowDslParserBLOCK_COMMENTF = 17
	flowDslParserLINE_COMMENTF  = 18
	flowDslParserWSF            = 19
	flowDslParserCOMMA          = 20
	flowDslParserPIPE           = 21
	flowDslParserDOTI           = 22
	flowDslParserRPAREN         = 23
	flowDslParserLBRACKI        = 24
	flowDslParserRBRACK         = 25
	flowDslParserIDI            = 26
	flowDslParserBLOCK_COMMENTI = 27
	flowDslParserLINE_COMMENTI  = 28
	flowDslParserWSI            = 29
	flowDslParserASSIGN         = 30
	flowDslParserCOMMAP         = 31
	flowDslParserDOTP           = 32
	flowDslParserPIPEP          = 33
	flowDslParserRBRACKP        = 34
	flowDslParserIDP            = 35
	flowDslParserBLOCK_COMMENTP = 36
	flowDslParserLINE_COMMENTP  = 37
	flowDslParserWSP            = 38
)

// flowDslParser rules.
const (
	flowDslParserRULE_flowFile          = 0
	flowDslParserRULE_imports           = 1
	flowDslParserRULE_flow              = 2
	flowDslParserRULE_statement         = 3
	flowDslParserRULE_statementStart    = 4
	flowDslParserRULE_statementMiddle   = 5
	flowDslParserRULE_statementEnd      = 6
	flowDslParserRULE_arrowComponent    = 7
	flowDslParserRULE_component         = 8
	flowDslParserRULE_componentTypeName = 9
	flowDslParserRULE_plugin            = 10
	flowDslParserRULE_pluginPart        = 11
	flowDslParserRULE_allData           = 12
	flowDslParserRULE_data              = 13
	flowDslParserRULE_packageIDI        = 14
	flowDslParserRULE_packageIDP        = 15
)

// IFlowFileContext is an interface to support dynamic dispatch.
type IFlowFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAllImports returns the AllImports rule contexts.
	GetAllImports() IImportsContext

	// Get_flow returns the _flow rule contexts.
	Get_flow() IFlowContext

	// SetAllImports sets the AllImports rule contexts.
	SetAllImports(IImportsContext)

	// Set_flow sets the _flow rule contexts.
	Set_flow(IFlowContext)

	// GetFlows returns the Flows rule context list.
	GetFlows() []IFlowContext

	// SetFlows sets the Flows rule context list.
	SetFlows([]IFlowContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	Imports() IImportsContext
	AllFlow() []IFlowContext
	Flow(i int) IFlowContext

	// IsFlowFileContext differentiates from other interfaces.
	IsFlowFileContext()
}

type FlowFileContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	AllImports IImportsContext
	_flow      IFlowContext
	Flows      []IFlowContext
}

func NewEmptyFlowFileContext() *FlowFileContext {
	var p = new(FlowFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flowFile
	return p
}

func InitEmptyFlowFileContext(p *FlowFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flowFile
}

func (*FlowFileContext) IsFlowFileContext() {}

func NewFlowFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowFileContext {
	var p = new(FlowFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_flowFile

	return p
}

func (s *FlowFileContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowFileContext) GetAllImports() IImportsContext { return s.AllImports }

func (s *FlowFileContext) Get_flow() IFlowContext { return s._flow }

func (s *FlowFileContext) SetAllImports(v IImportsContext) { s.AllImports = v }

func (s *FlowFileContext) Set_flow(v IFlowContext) { s._flow = v }

func (s *FlowFileContext) GetFlows() []IFlowContext { return s.Flows }

func (s *FlowFileContext) SetFlows(v []IFlowContext) { s.Flows = v }

func (s *FlowFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(flowDslParserEOF, 0)
}

func (s *FlowFileContext) Imports() IImportsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *FlowFileContext) AllFlow() []IFlowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFlowContext); ok {
			len++
		}
	}

	tst := make([]IFlowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFlowContext); ok {
			tst[i] = t.(IFlowContext)
			i++
		}
	}

	return tst
}

func (s *FlowFileContext) Flow(i int) IFlowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlowContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlowContext)
}

func (s *FlowFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) FlowFile() (localctx IFlowFileContext) {
	localctx = NewFlowFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, flowDslParserRULE_flowFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserIMPORT {
		{
			p.SetState(32)

			var _x = p.Imports()

			localctx.(*FlowFileContext).AllImports = _x
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserFLOW {
		{
			p.SetState(35)

			var _x = p.Flow()

			localctx.(*FlowFileContext)._flow = _x
		}
		localctx.(*FlowFileContext).Flows = append(localctx.(*FlowFileContext).Flows, localctx.(*FlowFileContext)._flow)

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(flowDslParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportsContext is an interface to support dynamic dispatch.
type IImportsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetImports returns the Imports token list.
	GetImports() []antlr.Token

	// SetImports sets the Imports token list.
	SetImports([]antlr.Token)

	// Getter signatures
	IMPORT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSEMIF() []antlr.TerminalNode
	SEMIF(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	_STRING antlr.Token
	Imports []antlr.Token
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_imports
	return p
}

func InitEmptyImportsContext(p *ImportsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_imports
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportsContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ImportsContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ImportsContext) GetImports() []antlr.Token { return s.Imports }

func (s *ImportsContext) SetImports(v []antlr.Token) { s.Imports = v }

func (s *ImportsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(flowDslParserIMPORT, 0)
}

func (s *ImportsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACE, 0)
}

func (s *ImportsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACE, 0)
}

func (s *ImportsContext) AllSEMIF() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSEMIF)
}

func (s *ImportsContext) SEMIF(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMIF, i)
}

func (s *ImportsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSTRING)
}

func (s *ImportsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSTRING, i)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, flowDslParserRULE_imports)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(flowDslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(flowDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserSTRING {
		{
			p.SetState(44)

			var _m = p.Match(flowDslParserSTRING)

			localctx.(*ImportsContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ImportsContext).Imports = append(localctx.(*ImportsContext).Imports, localctx.(*ImportsContext)._STRING)
		{
			p.SetState(45)
			p.Match(flowDslParserSEMIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(flowDslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFlowContext is an interface to support dynamic dispatch.
type IFlowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name token.
	GetName() antlr.Token

	// SetName sets the Name token.
	SetName(antlr.Token)

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStatements returns the Statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the Statements rule context list.
	SetStatements([]IStatementContext)

	// Getter signatures
	FLOW() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFlowContext differentiates from other interfaces.
	IsFlowContext()
}

type FlowContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	Name       antlr.Token
	_statement IStatementContext
	Statements []IStatementContext
}

func NewEmptyFlowContext() *FlowContext {
	var p = new(FlowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flow
	return p
}

func InitEmptyFlowContext(p *FlowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flow
}

func (*FlowContext) IsFlowContext() {}

func NewFlowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowContext {
	var p = new(FlowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_flow

	return p
}

func (s *FlowContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowContext) GetName() antlr.Token { return s.Name }

func (s *FlowContext) SetName(v antlr.Token) { s.Name = v }

func (s *FlowContext) Get_statement() IStatementContext { return s._statement }

func (s *FlowContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *FlowContext) GetStatements() []IStatementContext { return s.Statements }

func (s *FlowContext) SetStatements(v []IStatementContext) { s.Statements = v }

func (s *FlowContext) FLOW() antlr.TerminalNode {
	return s.GetToken(flowDslParserFLOW, 0)
}

func (s *FlowContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACE, 0)
}

func (s *FlowContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACE, 0)
}

func (s *FlowContext) ID() antlr.TerminalNode {
	return s.GetToken(flowDslParserID, 0)
}

func (s *FlowContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FlowContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FlowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Flow() (localctx IFlowContext) {
	localctx = NewFlowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, flowDslParserRULE_flow)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(flowDslParserFLOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)

		var _m = p.Match(flowDslParserID)

		localctx.(*FlowContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(flowDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserPORT || _la == flowDslParserLBRACK {
		{
			p.SetState(55)

			var _x = p.Statement()

			localctx.(*FlowContext)._statement = _x
		}
		localctx.(*FlowContext).Statements = append(localctx.(*FlowContext).Statements, localctx.(*FlowContext)._statement)

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(flowDslParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementMiddle() IStatementMiddleContext
	SEMIF() antlr.TerminalNode
	StatementStart() IStatementStartContext
	StatementEnd() IStatementEndContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StatementMiddle() IStatementMiddleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementMiddleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementMiddleContext)
}

func (s *StatementContext) SEMIF() antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMIF, 0)
}

func (s *StatementContext) StatementStart() IStatementStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementStartContext)
}

func (s *StatementContext) StatementEnd() IStatementEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementEndContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, flowDslParserRULE_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserPORT {
		{
			p.SetState(62)
			p.StatementStart()
		}

	}
	{
		p.SetState(65)
		p.StatementMiddle()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25088) != 0 {
		{
			p.SetState(66)
			p.StatementEnd()
		}

	}
	{
		p.SetState(69)
		p.Match(flowDslParserSEMIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementStartContext is an interface to support dynamic dispatch.
type IStatementStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStartPort returns the StartPort token.
	GetStartPort() antlr.Token

	// GetDstPort returns the DstPort token.
	GetDstPort() antlr.Token

	// SetStartPort sets the StartPort token.
	SetStartPort(antlr.Token)

	// SetDstPort sets the DstPort token.
	SetDstPort(antlr.Token)

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	AllPORT() []antlr.TerminalNode
	PORT(i int) antlr.TerminalNode
	AllData() IAllDataContext

	// IsStatementStartContext differentiates from other interfaces.
	IsStatementStartContext()
}

type StatementStartContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	StartPort  antlr.Token
	AllArrData IAllDataContext
	DstPort    antlr.Token
}

func NewEmptyStatementStartContext() *StatementStartContext {
	var p = new(StatementStartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementStart
	return p
}

func InitEmptyStatementStartContext(p *StatementStartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementStart
}

func (*StatementStartContext) IsStatementStartContext() {}

func NewStatementStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementStartContext {
	var p = new(StatementStartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_statementStart

	return p
}

func (s *StatementStartContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementStartContext) GetStartPort() antlr.Token { return s.StartPort }

func (s *StatementStartContext) GetDstPort() antlr.Token { return s.DstPort }

func (s *StatementStartContext) SetStartPort(v antlr.Token) { s.StartPort = v }

func (s *StatementStartContext) SetDstPort(v antlr.Token) { s.DstPort = v }

func (s *StatementStartContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *StatementStartContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *StatementStartContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
}

func (s *StatementStartContext) AllPORT() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserPORT)
}

func (s *StatementStartContext) PORT(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserPORT, i)
}

func (s *StatementStartContext) AllData() IAllDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllDataContext)
}

func (s *StatementStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) StatementStart() (localctx IStatementStartContext) {
	localctx = NewStatementStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, flowDslParserRULE_statementStart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)

		var _m = p.Match(flowDslParserPORT)

		localctx.(*StatementStartContext).StartPort = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)

		var _x = p.AllData()

		localctx.(*StatementStartContext).AllArrData = _x
	}
	{
		p.SetState(73)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserPORT {
		{
			p.SetState(74)

			var _m = p.Match(flowDslParserPORT)

			localctx.(*StatementStartContext).DstPort = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementMiddleContext is an interface to support dynamic dispatch.
type IStatementMiddleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirstComponent returns the FirstComponent rule contexts.
	GetFirstComponent() IComponentContext

	// Get_arrowComponent returns the _arrowComponent rule contexts.
	Get_arrowComponent() IArrowComponentContext

	// SetFirstComponent sets the FirstComponent rule contexts.
	SetFirstComponent(IComponentContext)

	// Set_arrowComponent sets the _arrowComponent rule contexts.
	Set_arrowComponent(IArrowComponentContext)

	// GetArrowComponents returns the ArrowComponents rule context list.
	GetArrowComponents() []IArrowComponentContext

	// SetArrowComponents sets the ArrowComponents rule context list.
	SetArrowComponents([]IArrowComponentContext)

	// Getter signatures
	Component() IComponentContext
	AllArrowComponent() []IArrowComponentContext
	ArrowComponent(i int) IArrowComponentContext

	// IsStatementMiddleContext differentiates from other interfaces.
	IsStatementMiddleContext()
}

type StatementMiddleContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	FirstComponent  IComponentContext
	_arrowComponent IArrowComponentContext
	ArrowComponents []IArrowComponentContext
}

func NewEmptyStatementMiddleContext() *StatementMiddleContext {
	var p = new(StatementMiddleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementMiddle
	return p
}

func InitEmptyStatementMiddleContext(p *StatementMiddleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementMiddle
}

func (*StatementMiddleContext) IsStatementMiddleContext() {}

func NewStatementMiddleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementMiddleContext {
	var p = new(StatementMiddleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_statementMiddle

	return p
}

func (s *StatementMiddleContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementMiddleContext) GetFirstComponent() IComponentContext { return s.FirstComponent }

func (s *StatementMiddleContext) Get_arrowComponent() IArrowComponentContext {
	return s._arrowComponent
}

func (s *StatementMiddleContext) SetFirstComponent(v IComponentContext) { s.FirstComponent = v }

func (s *StatementMiddleContext) Set_arrowComponent(v IArrowComponentContext) { s._arrowComponent = v }

func (s *StatementMiddleContext) GetArrowComponents() []IArrowComponentContext {
	return s.ArrowComponents
}

func (s *StatementMiddleContext) SetArrowComponents(v []IArrowComponentContext) {
	s.ArrowComponents = v
}

func (s *StatementMiddleContext) Component() IComponentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *StatementMiddleContext) AllArrowComponent() []IArrowComponentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrowComponentContext); ok {
			len++
		}
	}

	tst := make([]IArrowComponentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrowComponentContext); ok {
			tst[i] = t.(IArrowComponentContext)
			i++
		}
	}

	return tst
}

func (s *StatementMiddleContext) ArrowComponent(i int) IArrowComponentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrowComponentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrowComponentContext)
}

func (s *StatementMiddleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementMiddleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) StatementMiddle() (localctx IStatementMiddleContext) {
	localctx = NewStatementMiddleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, flowDslParserRULE_statementMiddle)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _x = p.Component()

		localctx.(*StatementMiddleContext).FirstComponent = _x
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(78)

				var _x = p.ArrowComponent()

				localctx.(*StatementMiddleContext)._arrowComponent = _x
			}
			localctx.(*StatementMiddleContext).ArrowComponents = append(localctx.(*StatementMiddleContext).ArrowComponents, localctx.(*StatementMiddleContext)._arrowComponent)

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementEndContext is an interface to support dynamic dispatch.
type IStatementEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSrcPort returns the SrcPort token.
	GetSrcPort() antlr.Token

	// GetEndPort returns the EndPort token.
	GetEndPort() antlr.Token

	// SetSrcPort sets the SrcPort token.
	SetSrcPort(antlr.Token)

	// SetEndPort sets the EndPort token.
	SetEndPort(antlr.Token)

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	AllPORT() []antlr.TerminalNode
	PORT(i int) antlr.TerminalNode
	AllData() IAllDataContext

	// IsStatementEndContext differentiates from other interfaces.
	IsStatementEndContext()
}

type StatementEndContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	SrcPort    antlr.Token
	AllArrData IAllDataContext
	EndPort    antlr.Token
}

func NewEmptyStatementEndContext() *StatementEndContext {
	var p = new(StatementEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementEnd
	return p
}

func InitEmptyStatementEndContext(p *StatementEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_statementEnd
}

func (*StatementEndContext) IsStatementEndContext() {}

func NewStatementEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementEndContext {
	var p = new(StatementEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_statementEnd

	return p
}

func (s *StatementEndContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementEndContext) GetSrcPort() antlr.Token { return s.SrcPort }

func (s *StatementEndContext) GetEndPort() antlr.Token { return s.EndPort }

func (s *StatementEndContext) SetSrcPort(v antlr.Token) { s.SrcPort = v }

func (s *StatementEndContext) SetEndPort(v antlr.Token) { s.EndPort = v }

func (s *StatementEndContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *StatementEndContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *StatementEndContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
}

func (s *StatementEndContext) AllPORT() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserPORT)
}

func (s *StatementEndContext) PORT(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserPORT, i)
}

func (s *StatementEndContext) AllData() IAllDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllDataContext)
}

func (s *StatementEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) StatementEnd() (localctx IStatementEndContext) {
	localctx = NewStatementEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, flowDslParserRULE_statementEnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserPORT {
		{
			p.SetState(84)

			var _m = p.Match(flowDslParserPORT)

			localctx.(*StatementEndContext).SrcPort = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLPAREN {
		{
			p.SetState(87)

			var _x = p.AllData()

			localctx.(*StatementEndContext).AllArrData = _x
		}

	}
	{
		p.SetState(90)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)

		var _m = p.Match(flowDslParserPORT)

		localctx.(*StatementEndContext).EndPort = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrowComponentContext is an interface to support dynamic dispatch.
type IArrowComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSrcPort returns the SrcPort token.
	GetSrcPort() antlr.Token

	// GetDstPort returns the DstPort token.
	GetDstPort() antlr.Token

	// SetSrcPort sets the SrcPort token.
	SetSrcPort(antlr.Token)

	// SetDstPort sets the DstPort token.
	SetDstPort(antlr.Token)

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// GetDstComponent returns the DstComponent rule contexts.
	GetDstComponent() IComponentContext

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// SetDstComponent sets the DstComponent rule contexts.
	SetDstComponent(IComponentContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	Component() IComponentContext
	AllPORT() []antlr.TerminalNode
	PORT(i int) antlr.TerminalNode
	AllData() IAllDataContext

	// IsArrowComponentContext differentiates from other interfaces.
	IsArrowComponentContext()
}

type ArrowComponentContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	SrcPort      antlr.Token
	AllArrData   IAllDataContext
	DstPort      antlr.Token
	DstComponent IComponentContext
}

func NewEmptyArrowComponentContext() *ArrowComponentContext {
	var p = new(ArrowComponentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_arrowComponent
	return p
}

func InitEmptyArrowComponentContext(p *ArrowComponentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_arrowComponent
}

func (*ArrowComponentContext) IsArrowComponentContext() {}

func NewArrowComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowComponentContext {
	var p = new(ArrowComponentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_arrowComponent

	return p
}

func (s *ArrowComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowComponentContext) GetSrcPort() antlr.Token { return s.SrcPort }

func (s *ArrowComponentContext) GetDstPort() antlr.Token { return s.DstPort }

func (s *ArrowComponentContext) SetSrcPort(v antlr.Token) { s.SrcPort = v }

func (s *ArrowComponentContext) SetDstPort(v antlr.Token) { s.DstPort = v }

func (s *ArrowComponentContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *ArrowComponentContext) GetDstComponent() IComponentContext { return s.DstComponent }

func (s *ArrowComponentContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *ArrowComponentContext) SetDstComponent(v IComponentContext) { s.DstComponent = v }

func (s *ArrowComponentContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
}

func (s *ArrowComponentContext) Component() IComponentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *ArrowComponentContext) AllPORT() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserPORT)
}

func (s *ArrowComponentContext) PORT(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserPORT, i)
}

func (s *ArrowComponentContext) AllData() IAllDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllDataContext)
}

func (s *ArrowComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) ArrowComponent() (localctx IArrowComponentContext) {
	localctx = NewArrowComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, flowDslParserRULE_arrowComponent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserPORT {
		{
			p.SetState(93)

			var _m = p.Match(flowDslParserPORT)

			localctx.(*ArrowComponentContext).SrcPort = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLPAREN {
		{
			p.SetState(96)

			var _x = p.AllData()

			localctx.(*ArrowComponentContext).AllArrData = _x
		}

	}
	{
		p.SetState(99)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserPORT {
		{
			p.SetState(100)

			var _m = p.Match(flowDslParserPORT)

			localctx.(*ArrowComponentContext).DstPort = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(103)

		var _x = p.Component()

		localctx.(*ArrowComponentContext).DstComponent = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCore returns the Core rule contexts.
	GetCore() IComponentTypeNameContext

	// GetAllPlugins returns the AllPlugins rule contexts.
	GetAllPlugins() IPluginContext

	// SetCore sets the Core rule contexts.
	SetCore(IComponentTypeNameContext)

	// SetAllPlugins sets the AllPlugins rule contexts.
	SetAllPlugins(IPluginContext)

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ComponentTypeName() IComponentTypeNameContext
	Plugin() IPluginContext

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	Core       IComponentTypeNameContext
	AllPlugins IPluginContext
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_component
	return p
}

func InitEmptyComponentContext(p *ComponentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_component
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) GetCore() IComponentTypeNameContext { return s.Core }

func (s *ComponentContext) GetAllPlugins() IPluginContext { return s.AllPlugins }

func (s *ComponentContext) SetCore(v IComponentTypeNameContext) { s.Core = v }

func (s *ComponentContext) SetAllPlugins(v IPluginContext) { s.AllPlugins = v }

func (s *ComponentContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACK, 0)
}

func (s *ComponentContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACK, 0)
}

func (s *ComponentContext) ComponentTypeName() IComponentTypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentTypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentTypeNameContext)
}

func (s *ComponentContext) Plugin() IPluginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPluginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPluginContext)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Component() (localctx IComponentContext) {
	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, flowDslParserRULE_component)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(flowDslParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)

		var _x = p.ComponentTypeName()

		localctx.(*ComponentContext).Core = _x
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLBRACKI {
		{
			p.SetState(107)

			var _x = p.Plugin()

			localctx.(*ComponentContext).AllPlugins = _x
		}

	}
	{
		p.SetState(110)
		p.Match(flowDslParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentTypeNameContext is an interface to support dynamic dispatch.
type IComponentTypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name token.
	GetName() antlr.Token

	// GetTypPack returns the TypPack token.
	GetTypPack() antlr.Token

	// GetTypName returns the TypName token.
	GetTypName() antlr.Token

	// SetName sets the Name token.
	SetName(antlr.Token)

	// SetTypPack sets the TypPack token.
	SetTypPack(antlr.Token)

	// SetTypName sets the TypName token.
	SetTypName(antlr.Token)

	// Getter signatures
	AllIDI() []antlr.TerminalNode
	IDI(i int) antlr.TerminalNode
	DOTI() antlr.TerminalNode

	// IsComponentTypeNameContext differentiates from other interfaces.
	IsComponentTypeNameContext()
}

type ComponentTypeNameContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	Name    antlr.Token
	TypPack antlr.Token
	TypName antlr.Token
}

func NewEmptyComponentTypeNameContext() *ComponentTypeNameContext {
	var p = new(ComponentTypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_componentTypeName
	return p
}

func InitEmptyComponentTypeNameContext(p *ComponentTypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_componentTypeName
}

func (*ComponentTypeNameContext) IsComponentTypeNameContext() {}

func NewComponentTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentTypeNameContext {
	var p = new(ComponentTypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_componentTypeName

	return p
}

func (s *ComponentTypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentTypeNameContext) GetName() antlr.Token { return s.Name }

func (s *ComponentTypeNameContext) GetTypPack() antlr.Token { return s.TypPack }

func (s *ComponentTypeNameContext) GetTypName() antlr.Token { return s.TypName }

func (s *ComponentTypeNameContext) SetName(v antlr.Token) { s.Name = v }

func (s *ComponentTypeNameContext) SetTypPack(v antlr.Token) { s.TypPack = v }

func (s *ComponentTypeNameContext) SetTypName(v antlr.Token) { s.TypName = v }

func (s *ComponentTypeNameContext) AllIDI() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserIDI)
}

func (s *ComponentTypeNameContext) IDI(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserIDI, i)
}

func (s *ComponentTypeNameContext) DOTI() antlr.TerminalNode {
	return s.GetToken(flowDslParserDOTI, 0)
}

func (s *ComponentTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) ComponentTypeName() (localctx IComponentTypeNameContext) {
	localctx = NewComponentTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, flowDslParserRULE_componentTypeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*ComponentTypeNameContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserIDI {
		{
			p.SetState(113)

			var _m = p.Match(flowDslParserIDI)

			localctx.(*ComponentTypeNameContext).TypPack = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == flowDslParserDOTI {
			{
				p.SetState(114)
				p.Match(flowDslParserDOTI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(115)

				var _m = p.Match(flowDslParserIDI)

				localctx.(*ComponentTypeNameContext).TypName = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPluginContext is an interface to support dynamic dispatch.
type IPluginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_pluginPart returns the _pluginPart rule contexts.
	Get_pluginPart() IPluginPartContext

	// Set_pluginPart sets the _pluginPart rule contexts.
	Set_pluginPart(IPluginPartContext)

	// GetPluginGroups returns the PluginGroups rule context list.
	GetPluginGroups() []IPluginPartContext

	// SetPluginGroups sets the PluginGroups rule context list.
	SetPluginGroups([]IPluginPartContext)

	// Getter signatures
	LBRACKI() antlr.TerminalNode
	RBRACKP() antlr.TerminalNode
	AllPluginPart() []IPluginPartContext
	PluginPart(i int) IPluginPartContext
	AllPIPEP() []antlr.TerminalNode
	PIPEP(i int) antlr.TerminalNode

	// IsPluginContext differentiates from other interfaces.
	IsPluginContext()
}

type PluginContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	_pluginPart  IPluginPartContext
	PluginGroups []IPluginPartContext
}

func NewEmptyPluginContext() *PluginContext {
	var p = new(PluginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_plugin
	return p
}

func InitEmptyPluginContext(p *PluginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_plugin
}

func (*PluginContext) IsPluginContext() {}

func NewPluginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PluginContext {
	var p = new(PluginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_plugin

	return p
}

func (s *PluginContext) GetParser() antlr.Parser { return s.parser }

func (s *PluginContext) Get_pluginPart() IPluginPartContext { return s._pluginPart }

func (s *PluginContext) Set_pluginPart(v IPluginPartContext) { s._pluginPart = v }

func (s *PluginContext) GetPluginGroups() []IPluginPartContext { return s.PluginGroups }

func (s *PluginContext) SetPluginGroups(v []IPluginPartContext) { s.PluginGroups = v }

func (s *PluginContext) LBRACKI() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACKI, 0)
}

func (s *PluginContext) RBRACKP() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACKP, 0)
}

func (s *PluginContext) AllPluginPart() []IPluginPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPluginPartContext); ok {
			len++
		}
	}

	tst := make([]IPluginPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPluginPartContext); ok {
			tst[i] = t.(IPluginPartContext)
			i++
		}
	}

	return tst
}

func (s *PluginContext) PluginPart(i int) IPluginPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPluginPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPluginPartContext)
}

func (s *PluginContext) AllPIPEP() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserPIPEP)
}

func (s *PluginContext) PIPEP(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserPIPEP, i)
}

func (s *PluginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PluginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Plugin() (localctx IPluginContext) {
	localctx = NewPluginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, flowDslParserRULE_plugin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(flowDslParserLBRACKI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)

		var _x = p.PluginPart()

		localctx.(*PluginContext)._pluginPart = _x
	}
	localctx.(*PluginContext).PluginGroups = append(localctx.(*PluginContext).PluginGroups, localctx.(*PluginContext)._pluginPart)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == flowDslParserPIPEP {
		{
			p.SetState(122)
			p.Match(flowDslParserPIPEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

			var _x = p.PluginPart()

			localctx.(*PluginContext)._pluginPart = _x
		}
		localctx.(*PluginContext).PluginGroups = append(localctx.(*PluginContext).PluginGroups, localctx.(*PluginContext)._pluginPart)

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(flowDslParserRBRACKP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPluginPartContext is an interface to support dynamic dispatch.
type IPluginPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInterface returns the Interface rule contexts.
	GetInterface() IPackageIDPContext

	// Get_packageIDP returns the _packageIDP rule contexts.
	Get_packageIDP() IPackageIDPContext

	// SetInterface sets the Interface rule contexts.
	SetInterface(IPackageIDPContext)

	// Set_packageIDP sets the _packageIDP rule contexts.
	Set_packageIDP(IPackageIDPContext)

	// GetPlugins returns the Plugins rule context list.
	GetPlugins() []IPackageIDPContext

	// SetPlugins sets the Plugins rule context list.
	SetPlugins([]IPackageIDPContext)

	// Getter signatures
	AllPackageIDP() []IPackageIDPContext
	PackageIDP(i int) IPackageIDPContext
	ASSIGN() antlr.TerminalNode
	AllCOMMAP() []antlr.TerminalNode
	COMMAP(i int) antlr.TerminalNode

	// IsPluginPartContext differentiates from other interfaces.
	IsPluginPartContext()
}

type PluginPartContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	Interface   IPackageIDPContext
	_packageIDP IPackageIDPContext
	Plugins     []IPackageIDPContext
}

func NewEmptyPluginPartContext() *PluginPartContext {
	var p = new(PluginPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_pluginPart
	return p
}

func InitEmptyPluginPartContext(p *PluginPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_pluginPart
}

func (*PluginPartContext) IsPluginPartContext() {}

func NewPluginPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PluginPartContext {
	var p = new(PluginPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_pluginPart

	return p
}

func (s *PluginPartContext) GetParser() antlr.Parser { return s.parser }

func (s *PluginPartContext) GetInterface() IPackageIDPContext { return s.Interface }

func (s *PluginPartContext) Get_packageIDP() IPackageIDPContext { return s._packageIDP }

func (s *PluginPartContext) SetInterface(v IPackageIDPContext) { s.Interface = v }

func (s *PluginPartContext) Set_packageIDP(v IPackageIDPContext) { s._packageIDP = v }

func (s *PluginPartContext) GetPlugins() []IPackageIDPContext { return s.Plugins }

func (s *PluginPartContext) SetPlugins(v []IPackageIDPContext) { s.Plugins = v }

func (s *PluginPartContext) AllPackageIDP() []IPackageIDPContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageIDPContext); ok {
			len++
		}
	}

	tst := make([]IPackageIDPContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageIDPContext); ok {
			tst[i] = t.(IPackageIDPContext)
			i++
		}
	}

	return tst
}

func (s *PluginPartContext) PackageIDP(i int) IPackageIDPContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageIDPContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageIDPContext)
}

func (s *PluginPartContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(flowDslParserASSIGN, 0)
}

func (s *PluginPartContext) AllCOMMAP() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserCOMMAP)
}

func (s *PluginPartContext) COMMAP(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserCOMMAP, i)
}

func (s *PluginPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PluginPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) PluginPart() (localctx IPluginPartContext) {
	localctx = NewPluginPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, flowDslParserRULE_pluginPart)
	var _la int

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)

			var _x = p.PackageIDP()

			localctx.(*PluginPartContext).Interface = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)

			var _x = p.PackageIDP()

			localctx.(*PluginPartContext).Interface = _x
		}
		{
			p.SetState(133)
			p.Match(flowDslParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.PackageIDP()

			localctx.(*PluginPartContext)._packageIDP = _x
		}
		localctx.(*PluginPartContext).Plugins = append(localctx.(*PluginPartContext).Plugins, localctx.(*PluginPartContext)._packageIDP)
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == flowDslParserCOMMAP {
			{
				p.SetState(135)
				p.Match(flowDslParserCOMMAP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(136)

				var _x = p.PackageIDP()

				localctx.(*PluginPartContext)._packageIDP = _x
			}
			localctx.(*PluginPartContext).Plugins = append(localctx.(*PluginPartContext).Plugins, localctx.(*PluginPartContext)._packageIDP)

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllDataContext is an interface to support dynamic dispatch.
type IAllDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetDatas returns the Datas rule context list.
	GetDatas() []IDataContext

	// SetDatas sets the Datas rule context list.
	SetDatas([]IDataContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllData() []IDataContext
	Data(i int) IDataContext

	// IsAllDataContext differentiates from other interfaces.
	IsAllDataContext()
}

type AllDataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	_data  IDataContext
	Datas  []IDataContext
}

func NewEmptyAllDataContext() *AllDataContext {
	var p = new(AllDataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_allData
	return p
}

func InitEmptyAllDataContext(p *AllDataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_allData
}

func (*AllDataContext) IsAllDataContext() {}

func NewAllDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllDataContext {
	var p = new(AllDataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_allData

	return p
}

func (s *AllDataContext) GetParser() antlr.Parser { return s.parser }

func (s *AllDataContext) Get_data() IDataContext { return s._data }

func (s *AllDataContext) Set_data(v IDataContext) { s._data = v }

func (s *AllDataContext) GetDatas() []IDataContext { return s.Datas }

func (s *AllDataContext) SetDatas(v []IDataContext) { s.Datas = v }

func (s *AllDataContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(flowDslParserLPAREN, 0)
}

func (s *AllDataContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(flowDslParserRPAREN, 0)
}

func (s *AllDataContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserCOMMA)
}

func (s *AllDataContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserCOMMA, i)
}

func (s *AllDataContext) AllData() []IDataContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataContext); ok {
			len++
		}
	}

	tst := make([]IDataContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataContext); ok {
			tst[i] = t.(IDataContext)
			i++
		}
	}

	return tst
}

func (s *AllDataContext) Data(i int) IDataContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *AllDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) AllData() (localctx IAllDataContext) {
	localctx = NewAllDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, flowDslParserRULE_allData)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(flowDslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserIDI {
		{
			p.SetState(145)

			var _x = p.Data()

			localctx.(*AllDataContext)._data = _x
		}
		localctx.(*AllDataContext).Datas = append(localctx.(*AllDataContext).Datas, localctx.(*AllDataContext)._data)

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == flowDslParserCOMMA {
		{
			p.SetState(148)
			p.Match(flowDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)

			var _x = p.Data()

			localctx.(*AllDataContext)._data = _x
		}
		localctx.(*AllDataContext).Datas = append(localctx.(*AllDataContext).Datas, localctx.(*AllDataContext)._data)

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(flowDslParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name token.
	GetName() antlr.Token

	// GetTypPack returns the TypPack token.
	GetTypPack() antlr.Token

	// GetTypName returns the TypName token.
	GetTypName() antlr.Token

	// SetName sets the Name token.
	SetName(antlr.Token)

	// SetTypPack sets the TypPack token.
	SetTypPack(antlr.Token)

	// SetTypName sets the TypName token.
	SetTypName(antlr.Token)

	// Getter signatures
	AllIDI() []antlr.TerminalNode
	IDI(i int) antlr.TerminalNode
	DOTI() antlr.TerminalNode

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	Name    antlr.Token
	TypPack antlr.Token
	TypName antlr.Token
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_data
	return p
}

func InitEmptyDataContext(p *DataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_data
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) GetName() antlr.Token { return s.Name }

func (s *DataContext) GetTypPack() antlr.Token { return s.TypPack }

func (s *DataContext) GetTypName() antlr.Token { return s.TypName }

func (s *DataContext) SetName(v antlr.Token) { s.Name = v }

func (s *DataContext) SetTypPack(v antlr.Token) { s.TypPack = v }

func (s *DataContext) SetTypName(v antlr.Token) { s.TypName = v }

func (s *DataContext) AllIDI() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserIDI)
}

func (s *DataContext) IDI(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserIDI, i)
}

func (s *DataContext) DOTI() antlr.TerminalNode {
	return s.GetToken(flowDslParserDOTI, 0)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, flowDslParserRULE_data)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*DataContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*DataContext).TypPack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTI {
		{
			p.SetState(159)
			p.Match(flowDslParserDOTI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _m = p.Match(flowDslParserIDI)

			localctx.(*DataContext).TypName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageIDIContext is an interface to support dynamic dispatch.
type IPackageIDIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetID1 returns the ID1 token.
	GetID1() antlr.Token

	// SetID1 sets the ID1 token.
	SetID1(antlr.Token)

	// Getter signatures
	IDI() antlr.TerminalNode

	// IsPackageIDIContext differentiates from other interfaces.
	IsPackageIDIContext()
}

type PackageIDIContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ID1    antlr.Token
}

func NewEmptyPackageIDIContext() *PackageIDIContext {
	var p = new(PackageIDIContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_packageIDI
	return p
}

func InitEmptyPackageIDIContext(p *PackageIDIContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_packageIDI
}

func (*PackageIDIContext) IsPackageIDIContext() {}

func NewPackageIDIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageIDIContext {
	var p = new(PackageIDIContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_packageIDI

	return p
}

func (s *PackageIDIContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageIDIContext) GetID1() antlr.Token { return s.ID1 }

func (s *PackageIDIContext) SetID1(v antlr.Token) { s.ID1 = v }

func (s *PackageIDIContext) IDI() antlr.TerminalNode {
	return s.GetToken(flowDslParserIDI, 0)
}

func (s *PackageIDIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageIDIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) PackageIDI() (localctx IPackageIDIContext) {
	localctx = NewPackageIDIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, flowDslParserRULE_packageIDI)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*PackageIDIContext).ID1 = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageIDPContext is an interface to support dynamic dispatch.
type IPackageIDPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetID1 returns the ID1 token.
	GetID1() antlr.Token

	// GetID2 returns the ID2 token.
	GetID2() antlr.Token

	// SetID1 sets the ID1 token.
	SetID1(antlr.Token)

	// SetID2 sets the ID2 token.
	SetID2(antlr.Token)

	// Getter signatures
	AllIDP() []antlr.TerminalNode
	IDP(i int) antlr.TerminalNode
	DOTP() antlr.TerminalNode

	// IsPackageIDPContext differentiates from other interfaces.
	IsPackageIDPContext()
}

type PackageIDPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ID1    antlr.Token
	ID2    antlr.Token
}

func NewEmptyPackageIDPContext() *PackageIDPContext {
	var p = new(PackageIDPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_packageIDP
	return p
}

func InitEmptyPackageIDPContext(p *PackageIDPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_packageIDP
}

func (*PackageIDPContext) IsPackageIDPContext() {}

func NewPackageIDPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageIDPContext {
	var p = new(PackageIDPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_packageIDP

	return p
}

func (s *PackageIDPContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageIDPContext) GetID1() antlr.Token { return s.ID1 }

func (s *PackageIDPContext) GetID2() antlr.Token { return s.ID2 }

func (s *PackageIDPContext) SetID1(v antlr.Token) { s.ID1 = v }

func (s *PackageIDPContext) SetID2(v antlr.Token) { s.ID2 = v }

func (s *PackageIDPContext) AllIDP() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserIDP)
}

func (s *PackageIDPContext) IDP(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserIDP, i)
}

func (s *PackageIDPContext) DOTP() antlr.TerminalNode {
	return s.GetToken(flowDslParserDOTP, 0)
}

func (s *PackageIDPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageIDPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) PackageIDP() (localctx IPackageIDPContext) {
	localctx = NewPackageIDPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, flowDslParserRULE_packageIDP)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)

		var _m = p.Match(flowDslParserIDP)

		localctx.(*PackageIDPContext).ID1 = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTP {
		{
			p.SetState(166)
			p.Match(flowDslParserDOTP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)

			var _m = p.Match(flowDslParserIDP)

			localctx.(*PackageIDPContext).ID2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
