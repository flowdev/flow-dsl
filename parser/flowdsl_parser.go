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
		"", "'import'", "'flow'", "'{'", "", "", "", "", "'->'", "';'", "':'",
		"", "", "", "'('", "", "'}'", "", "", "", "", "", "", "')'", "", "",
		"", "", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "FLOW", "LBRACE", "ID", "BLOCK_COMMENT", "LINE_COMMENT",
		"WS", "ARROW", "SEMI", "COLON", "STRING", "NAME", "INT", "LPAREN", "LBRACK",
		"RBRACE", "BLOCK_COMMENTF", "LINE_COMMENTF", "WSF", "COMMA", "PIPE",
		"DOTI", "RPAREN", "LBRACKI", "RBRACK", "IDI", "BLOCK_COMMENTI", "LINE_COMMENTI",
		"WSI", "ASSIGN", "COMMAP", "DOTP", "PIPEP", "RBRACKP", "IDP", "BLOCK_COMMENTP",
		"LINE_COMMENTP", "WSP",
	}
	staticData.RuleNames = []string{
		"flows", "imp", "flow", "statement", "statementMiddle", "statementStart",
		"statementEnd", "arrowComponent", "component", "componentTypeName",
		"plugin", "pluginPart", "allData", "data", "packageIDI", "packageIDP",
		"port",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 178, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 3, 0, 36, 8, 0, 1, 0, 4, 0, 39, 8, 0, 11, 0, 12, 0,
		40, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 49, 8, 1, 11, 1, 12, 1, 50,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 59, 8, 2, 11, 2, 12, 2, 60, 1,
		2, 1, 2, 1, 3, 3, 3, 66, 8, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 5, 4, 76, 8, 4, 10, 4, 12, 4, 79, 9, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 85, 8, 5, 1, 6, 3, 6, 88, 8, 6, 1, 6, 3, 6, 91, 8, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 3, 7, 97, 8, 7, 1, 7, 3, 7, 100, 8, 7, 1, 7, 1, 7, 3, 7,
		104, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 111, 8, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 3, 9, 117, 8, 9, 1, 9, 1, 9, 3, 9, 121, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 5, 10, 127, 8, 10, 10, 10, 12, 10, 130, 9, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 140, 8, 11, 10, 11, 12,
		11, 143, 9, 11, 3, 11, 145, 8, 11, 1, 12, 1, 12, 3, 12, 149, 8, 12, 1,
		12, 1, 12, 5, 12, 153, 8, 12, 10, 12, 12, 12, 156, 9, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 164, 8, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 3, 15, 171, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 176, 8, 16, 1,
		16, 1, 77, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 0, 0, 184, 0, 35, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 54, 1, 0, 0,
		0, 6, 65, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 87, 1,
		0, 0, 0, 14, 96, 1, 0, 0, 0, 16, 107, 1, 0, 0, 0, 18, 114, 1, 0, 0, 0,
		20, 122, 1, 0, 0, 0, 22, 144, 1, 0, 0, 0, 24, 146, 1, 0, 0, 0, 26, 159,
		1, 0, 0, 0, 28, 165, 1, 0, 0, 0, 30, 167, 1, 0, 0, 0, 32, 172, 1, 0, 0,
		0, 34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38,
		1, 0, 0, 0, 37, 39, 3, 4, 2, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0,
		40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 5,
		0, 0, 1, 43, 1, 1, 0, 0, 0, 44, 45, 5, 1, 0, 0, 45, 48, 5, 3, 0, 0, 46,
		47, 5, 11, 0, 0, 47, 49, 5, 9, 0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 1, 0,
		0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53,
		5, 16, 0, 0, 53, 3, 1, 0, 0, 0, 54, 55, 5, 2, 0, 0, 55, 56, 5, 4, 0, 0,
		56, 58, 5, 3, 0, 0, 57, 59, 3, 6, 3, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1,
		0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62,
		63, 5, 16, 0, 0, 63, 5, 1, 0, 0, 0, 64, 66, 3, 10, 5, 0, 65, 64, 1, 0,
		0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 3, 8, 4, 0, 68, 70,
		3, 12, 6, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 72, 5, 9, 0, 0, 72, 7, 1, 0, 0, 0, 73, 77, 3, 16, 8, 0, 74, 76, 3,
		14, 7, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 78, 9, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 3, 32, 16,
		0, 81, 82, 3, 24, 12, 0, 82, 84, 5, 8, 0, 0, 83, 85, 3, 32, 16, 0, 84,
		83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 11, 1, 0, 0, 0, 86, 88, 3, 32,
		16, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89,
		91, 3, 24, 12, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0,
		0, 0, 92, 93, 5, 8, 0, 0, 93, 94, 3, 32, 16, 0, 94, 13, 1, 0, 0, 0, 95,
		97, 3, 32, 16, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1, 0,
		0, 0, 98, 100, 3, 24, 12, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100,
		101, 1, 0, 0, 0, 101, 103, 5, 8, 0, 0, 102, 104, 3, 32, 16, 0, 103, 102,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 3, 16,
		8, 0, 106, 15, 1, 0, 0, 0, 107, 108, 5, 15, 0, 0, 108, 110, 3, 18, 9, 0,
		109, 111, 3, 20, 10, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 113, 5, 25, 0, 0, 113, 17, 1, 0, 0, 0, 114, 116,
		5, 26, 0, 0, 115, 117, 5, 26, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1,
		0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 119, 5, 22, 0, 0, 119, 121, 5, 26,
		0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 19, 1, 0, 0, 0,
		122, 123, 5, 24, 0, 0, 123, 128, 3, 22, 11, 0, 124, 125, 5, 33, 0, 0, 125,
		127, 3, 22, 11, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0,
		0, 0, 131, 132, 5, 34, 0, 0, 132, 21, 1, 0, 0, 0, 133, 145, 3, 30, 15,
		0, 134, 135, 3, 30, 15, 0, 135, 136, 5, 30, 0, 0, 136, 141, 3, 30, 15,
		0, 137, 138, 5, 31, 0, 0, 138, 140, 3, 30, 15, 0, 139, 137, 1, 0, 0, 0,
		140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 133, 1, 0, 0, 0, 144, 134,
		1, 0, 0, 0, 145, 23, 1, 0, 0, 0, 146, 148, 5, 14, 0, 0, 147, 149, 3, 26,
		13, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 154, 1, 0, 0, 0,
		150, 151, 5, 20, 0, 0, 151, 153, 3, 26, 13, 0, 152, 150, 1, 0, 0, 0, 153,
		156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157,
		1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 23, 0, 0, 158, 25, 1, 0,
		0, 0, 159, 160, 5, 26, 0, 0, 160, 163, 5, 26, 0, 0, 161, 162, 5, 22, 0,
		0, 162, 164, 5, 26, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		27, 1, 0, 0, 0, 165, 166, 5, 26, 0, 0, 166, 29, 1, 0, 0, 0, 167, 170, 5,
		35, 0, 0, 168, 169, 5, 32, 0, 0, 169, 171, 5, 35, 0, 0, 170, 168, 1, 0,
		0, 0, 170, 171, 1, 0, 0, 0, 171, 31, 1, 0, 0, 0, 172, 175, 5, 12, 0, 0,
		173, 174, 5, 10, 0, 0, 174, 176, 5, 13, 0, 0, 175, 173, 1, 0, 0, 0, 175,
		176, 1, 0, 0, 0, 176, 33, 1, 0, 0, 0, 24, 35, 40, 50, 60, 65, 69, 77, 84,
		87, 90, 96, 99, 103, 110, 116, 120, 128, 141, 144, 148, 154, 163, 170,
		175,
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
	flowDslParserIMPORT         = 1
	flowDslParserFLOW           = 2
	flowDslParserLBRACE         = 3
	flowDslParserID             = 4
	flowDslParserBLOCK_COMMENT  = 5
	flowDslParserLINE_COMMENT   = 6
	flowDslParserWS             = 7
	flowDslParserARROW          = 8
	flowDslParserSEMI           = 9
	flowDslParserCOLON          = 10
	flowDslParserSTRING         = 11
	flowDslParserNAME           = 12
	flowDslParserINT            = 13
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
	flowDslParserRULE_flows             = 0
	flowDslParserRULE_imp               = 1
	flowDslParserRULE_flow              = 2
	flowDslParserRULE_statement         = 3
	flowDslParserRULE_statementMiddle   = 4
	flowDslParserRULE_statementStart    = 5
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
	flowDslParserRULE_port              = 16
)

// IFlowsContext is an interface to support dynamic dispatch.
type IFlowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImport returns the Import rule contexts.
	GetImport() IImpContext

	// Get_flow returns the _flow rule contexts.
	Get_flow() IFlowContext

	// SetImport sets the Import rule contexts.
	SetImport(IImpContext)

	// Set_flow sets the _flow rule contexts.
	Set_flow(IFlowContext)

	// GetFlows returns the Flows rule context list.
	GetFlows() []IFlowContext

	// SetFlows sets the Flows rule context list.
	SetFlows([]IFlowContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	Imp() IImpContext
	AllFlow() []IFlowContext
	Flow(i int) IFlowContext

	// IsFlowsContext differentiates from other interfaces.
	IsFlowsContext()
}

type FlowsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	Import IImpContext
	_flow  IFlowContext
	Flows  []IFlowContext
}

func NewEmptyFlowsContext() *FlowsContext {
	var p = new(FlowsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flows
	return p
}

func InitEmptyFlowsContext(p *FlowsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_flows
}

func (*FlowsContext) IsFlowsContext() {}

func NewFlowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowsContext {
	var p = new(FlowsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_flows

	return p
}

func (s *FlowsContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowsContext) GetImport() IImpContext { return s.Import }

func (s *FlowsContext) Get_flow() IFlowContext { return s._flow }

func (s *FlowsContext) SetImport(v IImpContext) { s.Import = v }

func (s *FlowsContext) Set_flow(v IFlowContext) { s._flow = v }

func (s *FlowsContext) GetFlows() []IFlowContext { return s.Flows }

func (s *FlowsContext) SetFlows(v []IFlowContext) { s.Flows = v }

func (s *FlowsContext) EOF() antlr.TerminalNode {
	return s.GetToken(flowDslParserEOF, 0)
}

func (s *FlowsContext) Imp() IImpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImpContext)
}

func (s *FlowsContext) AllFlow() []IFlowContext {
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

func (s *FlowsContext) Flow(i int) IFlowContext {
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

func (s *FlowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Flows() (localctx IFlowsContext) {
	localctx = NewFlowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, flowDslParserRULE_flows)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserIMPORT {
		{
			p.SetState(34)

			var _x = p.Imp()

			localctx.(*FlowsContext).Import = _x
		}

	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserFLOW {
		{
			p.SetState(37)

			var _x = p.Flow()

			localctx.(*FlowsContext)._flow = _x
		}
		localctx.(*FlowsContext).Flows = append(localctx.(*FlowsContext).Flows, localctx.(*FlowsContext)._flow)

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
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

// IImpContext is an interface to support dynamic dispatch.
type IImpContext interface {
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
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsImpContext differentiates from other interfaces.
	IsImpContext()
}

type ImpContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	_STRING antlr.Token
	Imports []antlr.Token
}

func NewEmptyImpContext() *ImpContext {
	var p = new(ImpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_imp
	return p
}

func InitEmptyImpContext(p *ImpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_imp
}

func (*ImpContext) IsImpContext() {}

func NewImpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpContext {
	var p = new(ImpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_imp

	return p
}

func (s *ImpContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ImpContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ImpContext) GetImports() []antlr.Token { return s.Imports }

func (s *ImpContext) SetImports(v []antlr.Token) { s.Imports = v }

func (s *ImpContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(flowDslParserIMPORT, 0)
}

func (s *ImpContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACE, 0)
}

func (s *ImpContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACE, 0)
}

func (s *ImpContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSEMI)
}

func (s *ImpContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMI, i)
}

func (s *ImpContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSTRING)
}

func (s *ImpContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSTRING, i)
}

func (s *ImpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Imp() (localctx IImpContext) {
	localctx = NewImpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, flowDslParserRULE_imp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(flowDslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(flowDslParserLBRACE)
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

	for ok := true; ok; ok = _la == flowDslParserSTRING {
		{
			p.SetState(46)

			var _m = p.Match(flowDslParserSTRING)

			localctx.(*ImpContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ImpContext).Imports = append(localctx.(*ImpContext).Imports, localctx.(*ImpContext)._STRING)
		{
			p.SetState(47)
			p.Match(flowDslParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
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
		p.SetState(54)
		p.Match(flowDslParserFLOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)

		var _m = p.Match(flowDslParserID)

		localctx.(*FlowContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(flowDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserNAME || _la == flowDslParserLBRACK {
		{
			p.SetState(57)

			var _x = p.Statement()

			localctx.(*FlowContext)._statement = _x
		}
		localctx.(*FlowContext).Statements = append(localctx.(*FlowContext).Statements, localctx.(*FlowContext)._statement)

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
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
	SEMI() antlr.TerminalNode
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

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMI, 0)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(64)
			p.StatementStart()
		}

	}
	{
		p.SetState(67)
		p.StatementMiddle()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20736) != 0 {
		{
			p.SetState(68)
			p.StatementEnd()
		}

	}
	{
		p.SetState(71)
		p.Match(flowDslParserSEMI)
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
	p.EnterRule(localctx, 8, flowDslParserRULE_statementMiddle)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)

		var _x = p.Component()

		localctx.(*StatementMiddleContext).FirstComponent = _x
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(74)

				var _x = p.ArrowComponent()

				localctx.(*StatementMiddleContext)._arrowComponent = _x
			}
			localctx.(*StatementMiddleContext).ArrowComponents = append(localctx.(*StatementMiddleContext).ArrowComponents, localctx.(*StatementMiddleContext)._arrowComponent)

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

// IStatementStartContext is an interface to support dynamic dispatch.
type IStatementStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStartPort returns the StartPort rule contexts.
	GetStartPort() IPortContext

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// GetDstPort returns the DstPort rule contexts.
	GetDstPort() IPortContext

	// SetStartPort sets the StartPort rule contexts.
	SetStartPort(IPortContext)

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// SetDstPort sets the DstPort rule contexts.
	SetDstPort(IPortContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	AllPort() []IPortContext
	Port(i int) IPortContext
	AllData() IAllDataContext

	// IsStatementStartContext differentiates from other interfaces.
	IsStatementStartContext()
}

type StatementStartContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	StartPort  IPortContext
	AllArrData IAllDataContext
	DstPort    IPortContext
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

func (s *StatementStartContext) GetStartPort() IPortContext { return s.StartPort }

func (s *StatementStartContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *StatementStartContext) GetDstPort() IPortContext { return s.DstPort }

func (s *StatementStartContext) SetStartPort(v IPortContext) { s.StartPort = v }

func (s *StatementStartContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *StatementStartContext) SetDstPort(v IPortContext) { s.DstPort = v }

func (s *StatementStartContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
}

func (s *StatementStartContext) AllPort() []IPortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPortContext); ok {
			len++
		}
	}

	tst := make([]IPortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPortContext); ok {
			tst[i] = t.(IPortContext)
			i++
		}
	}

	return tst
}

func (s *StatementStartContext) Port(i int) IPortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
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

	return t.(IPortContext)
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
	p.EnterRule(localctx, 10, flowDslParserRULE_statementStart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)

		var _x = p.Port()

		localctx.(*StatementStartContext).StartPort = _x
	}
	{
		p.SetState(81)

		var _x = p.AllData()

		localctx.(*StatementStartContext).AllArrData = _x
	}
	{
		p.SetState(82)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(83)

			var _x = p.Port()

			localctx.(*StatementStartContext).DstPort = _x
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

	// GetSrcPort returns the SrcPort rule contexts.
	GetSrcPort() IPortContext

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// GetEndPort returns the EndPort rule contexts.
	GetEndPort() IPortContext

	// SetSrcPort sets the SrcPort rule contexts.
	SetSrcPort(IPortContext)

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// SetEndPort sets the EndPort rule contexts.
	SetEndPort(IPortContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	AllPort() []IPortContext
	Port(i int) IPortContext
	AllData() IAllDataContext

	// IsStatementEndContext differentiates from other interfaces.
	IsStatementEndContext()
}

type StatementEndContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	SrcPort    IPortContext
	AllArrData IAllDataContext
	EndPort    IPortContext
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

func (s *StatementEndContext) GetSrcPort() IPortContext { return s.SrcPort }

func (s *StatementEndContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *StatementEndContext) GetEndPort() IPortContext { return s.EndPort }

func (s *StatementEndContext) SetSrcPort(v IPortContext) { s.SrcPort = v }

func (s *StatementEndContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *StatementEndContext) SetEndPort(v IPortContext) { s.EndPort = v }

func (s *StatementEndContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
}

func (s *StatementEndContext) AllPort() []IPortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPortContext); ok {
			len++
		}
	}

	tst := make([]IPortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPortContext); ok {
			tst[i] = t.(IPortContext)
			i++
		}
	}

	return tst
}

func (s *StatementEndContext) Port(i int) IPortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
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

	return t.(IPortContext)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(86)

			var _x = p.Port()

			localctx.(*StatementEndContext).SrcPort = _x
		}

	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLPAREN {
		{
			p.SetState(89)

			var _x = p.AllData()

			localctx.(*StatementEndContext).AllArrData = _x
		}

	}
	{
		p.SetState(92)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)

		var _x = p.Port()

		localctx.(*StatementEndContext).EndPort = _x
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

	// GetSrcPort returns the SrcPort rule contexts.
	GetSrcPort() IPortContext

	// GetAllArrData returns the AllArrData rule contexts.
	GetAllArrData() IAllDataContext

	// GetDstPort returns the DstPort rule contexts.
	GetDstPort() IPortContext

	// GetDstComponent returns the DstComponent rule contexts.
	GetDstComponent() IComponentContext

	// SetSrcPort sets the SrcPort rule contexts.
	SetSrcPort(IPortContext)

	// SetAllArrData sets the AllArrData rule contexts.
	SetAllArrData(IAllDataContext)

	// SetDstPort sets the DstPort rule contexts.
	SetDstPort(IPortContext)

	// SetDstComponent sets the DstComponent rule contexts.
	SetDstComponent(IComponentContext)

	// Getter signatures
	ARROW() antlr.TerminalNode
	Component() IComponentContext
	AllPort() []IPortContext
	Port(i int) IPortContext
	AllData() IAllDataContext

	// IsArrowComponentContext differentiates from other interfaces.
	IsArrowComponentContext()
}

type ArrowComponentContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	SrcPort      IPortContext
	AllArrData   IAllDataContext
	DstPort      IPortContext
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

func (s *ArrowComponentContext) GetSrcPort() IPortContext { return s.SrcPort }

func (s *ArrowComponentContext) GetAllArrData() IAllDataContext { return s.AllArrData }

func (s *ArrowComponentContext) GetDstPort() IPortContext { return s.DstPort }

func (s *ArrowComponentContext) GetDstComponent() IComponentContext { return s.DstComponent }

func (s *ArrowComponentContext) SetSrcPort(v IPortContext) { s.SrcPort = v }

func (s *ArrowComponentContext) SetAllArrData(v IAllDataContext) { s.AllArrData = v }

func (s *ArrowComponentContext) SetDstPort(v IPortContext) { s.DstPort = v }

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

func (s *ArrowComponentContext) AllPort() []IPortContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPortContext); ok {
			len++
		}
	}

	tst := make([]IPortContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPortContext); ok {
			tst[i] = t.(IPortContext)
			i++
		}
	}

	return tst
}

func (s *ArrowComponentContext) Port(i int) IPortContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
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

	return t.(IPortContext)
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
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(95)

			var _x = p.Port()

			localctx.(*ArrowComponentContext).SrcPort = _x
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLPAREN {
		{
			p.SetState(98)

			var _x = p.AllData()

			localctx.(*ArrowComponentContext).AllArrData = _x
		}

	}
	{
		p.SetState(101)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(102)

			var _x = p.Port()

			localctx.(*ArrowComponentContext).DstPort = _x
		}

	}
	{
		p.SetState(105)

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
		p.SetState(107)
		p.Match(flowDslParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)

		var _x = p.ComponentTypeName()

		localctx.(*ComponentContext).Core = _x
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserLBRACKI {
		{
			p.SetState(109)

			var _x = p.Plugin()

			localctx.(*ComponentContext).AllPlugins = _x
		}

	}
	{
		p.SetState(112)
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
		p.SetState(114)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*ComponentTypeNameContext).Name = _m
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

	if _la == flowDslParserIDI {
		{
			p.SetState(115)

			var _m = p.Match(flowDslParserIDI)

			localctx.(*ComponentTypeNameContext).TypPack = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTI {
		{
			p.SetState(118)
			p.Match(flowDslParserDOTI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)

			var _m = p.Match(flowDslParserIDI)

			localctx.(*ComponentTypeNameContext).TypName = _m
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
		p.SetState(122)
		p.Match(flowDslParserLBRACKI)
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

	for _la == flowDslParserPIPEP {
		{
			p.SetState(124)
			p.Match(flowDslParserPIPEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)

			var _x = p.PluginPart()

			localctx.(*PluginContext)._pluginPart = _x
		}
		localctx.(*PluginContext).PluginGroups = append(localctx.(*PluginContext).PluginGroups, localctx.(*PluginContext)._pluginPart)

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)

			var _x = p.PackageIDP()

			localctx.(*PluginPartContext).Interface = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)

			var _x = p.PackageIDP()

			localctx.(*PluginPartContext).Interface = _x
		}
		{
			p.SetState(135)
			p.Match(flowDslParserASSIGN)
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

		for _la == flowDslParserCOMMAP {
			{
				p.SetState(137)
				p.Match(flowDslParserCOMMAP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(138)

				var _x = p.PackageIDP()

				localctx.(*PluginPartContext)._packageIDP = _x
			}
			localctx.(*PluginPartContext).Plugins = append(localctx.(*PluginPartContext).Plugins, localctx.(*PluginPartContext)._packageIDP)

			p.SetState(143)
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
		p.SetState(146)
		p.Match(flowDslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserIDI {
		{
			p.SetState(147)

			var _x = p.Data()

			localctx.(*AllDataContext)._data = _x
		}
		localctx.(*AllDataContext).Datas = append(localctx.(*AllDataContext).Datas, localctx.(*AllDataContext)._data)

	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == flowDslParserCOMMA {
		{
			p.SetState(150)
			p.Match(flowDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)

			var _x = p.Data()

			localctx.(*AllDataContext)._data = _x
		}
		localctx.(*AllDataContext).Datas = append(localctx.(*AllDataContext).Datas, localctx.(*AllDataContext)._data)

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
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
		p.SetState(159)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*DataContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)

		var _m = p.Match(flowDslParserIDI)

		localctx.(*DataContext).TypPack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTI {
		{
			p.SetState(161)
			p.Match(flowDslParserDOTI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

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
		p.SetState(165)

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
		p.SetState(167)

		var _m = p.Match(flowDslParserIDP)

		localctx.(*PackageIDPContext).ID1 = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTP {
		{
			p.SetState(168)
			p.Match(flowDslParserDOTP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)

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

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the Name token.
	GetName() antlr.Token

	// GetNum returns the Num token.
	GetNum() antlr.Token

	// SetName sets the Name token.
	SetName(antlr.Token)

	// SetNum sets the Num token.
	SetNum(antlr.Token)

	// Getter signatures
	NAME() antlr.TerminalNode
	COLON() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	Name   antlr.Token
	Num    antlr.Token
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_port
	return p
}

func InitEmptyPortContext(p *PortContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_port
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) GetName() antlr.Token { return s.Name }

func (s *PortContext) GetNum() antlr.Token { return s.Num }

func (s *PortContext) SetName(v antlr.Token) { s.Name = v }

func (s *PortContext) SetNum(v antlr.Token) { s.Num = v }

func (s *PortContext) NAME() antlr.TerminalNode {
	return s.GetToken(flowDslParserNAME, 0)
}

func (s *PortContext) COLON() antlr.TerminalNode {
	return s.GetToken(flowDslParserCOLON, 0)
}

func (s *PortContext) INT() antlr.TerminalNode {
	return s.GetToken(flowDslParserINT, 0)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, flowDslParserRULE_port)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)

		var _m = p.Match(flowDslParserNAME)

		localctx.(*PortContext).Name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserCOLON {
		{
			p.SetState(173)
			p.Match(flowDslParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _m = p.Match(flowDslParserINT)

			localctx.(*PortContext).Num = _m
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
