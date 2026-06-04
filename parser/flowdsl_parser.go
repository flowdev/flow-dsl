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
		"flows", "imports", "flow", "statement", "statementMiddle", "statementStart",
		"statementEnd", "component", "componentTypeName", "plugin", "pluginPart",
		"data", "packageIDI", "packageIDP", "port", "dataSep",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 171, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 4, 0, 34, 8, 0, 11, 0, 12, 0, 35, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 4, 1, 44, 8, 1, 11, 1, 12, 1, 45, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		4, 2, 54, 8, 2, 11, 2, 12, 2, 55, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 4, 1, 4, 3, 4,
		75, 8, 4, 1, 4, 3, 4, 78, 8, 4, 1, 4, 1, 4, 3, 4, 82, 8, 4, 1, 4, 5, 4,
		85, 8, 4, 10, 4, 12, 4, 88, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 94, 8,
		5, 1, 6, 3, 6, 97, 8, 6, 1, 6, 3, 6, 100, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 3, 7, 109, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 114, 8, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 122, 8, 9, 10, 9, 12, 9, 125, 9,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 135, 8,
		10, 10, 10, 12, 10, 138, 9, 10, 3, 10, 140, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 147, 8, 11, 10, 11, 12, 11, 150, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 3, 12, 157, 8, 12, 1, 13, 1, 13, 1, 13, 3, 13,
		162, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14, 167, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 86, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		0, 1, 1, 0, 20, 21, 175, 0, 33, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 49, 1,
		0, 0, 0, 6, 70, 1, 0, 0, 0, 8, 72, 1, 0, 0, 0, 10, 89, 1, 0, 0, 0, 12,
		96, 1, 0, 0, 0, 14, 105, 1, 0, 0, 0, 16, 113, 1, 0, 0, 0, 18, 117, 1, 0,
		0, 0, 20, 139, 1, 0, 0, 0, 22, 141, 1, 0, 0, 0, 24, 153, 1, 0, 0, 0, 26,
		158, 1, 0, 0, 0, 28, 163, 1, 0, 0, 0, 30, 168, 1, 0, 0, 0, 32, 34, 3, 4,
		2, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5, 0, 0, 1, 38, 1, 1, 0, 0, 0,
		39, 40, 5, 1, 0, 0, 40, 43, 5, 3, 0, 0, 41, 42, 5, 11, 0, 0, 42, 44, 5,
		9, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45,
		46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 5, 16, 0, 0, 48, 3, 1, 0, 0,
		0, 49, 50, 5, 2, 0, 0, 50, 51, 5, 4, 0, 0, 51, 53, 5, 3, 0, 0, 52, 54,
		3, 6, 3, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 5, 16, 0, 0, 58, 5, 1,
		0, 0, 0, 59, 60, 3, 10, 5, 0, 60, 61, 3, 8, 4, 0, 61, 62, 3, 12, 6, 0,
		62, 71, 1, 0, 0, 0, 63, 64, 3, 10, 5, 0, 64, 65, 3, 8, 4, 0, 65, 66, 5,
		9, 0, 0, 66, 71, 1, 0, 0, 0, 67, 68, 3, 8, 4, 0, 68, 69, 3, 12, 6, 0, 69,
		71, 1, 0, 0, 0, 70, 59, 1, 0, 0, 0, 70, 63, 1, 0, 0, 0, 70, 67, 1, 0, 0,
		0, 71, 7, 1, 0, 0, 0, 72, 86, 3, 14, 7, 0, 73, 75, 3, 28, 14, 0, 74, 73,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 78, 3, 22, 11,
		0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81,
		5, 8, 0, 0, 80, 82, 3, 28, 14, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0,
		0, 82, 83, 1, 0, 0, 0, 83, 85, 3, 14, 7, 0, 84, 74, 1, 0, 0, 0, 85, 88,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 9, 1, 0, 0, 0,
		88, 86, 1, 0, 0, 0, 89, 90, 3, 28, 14, 0, 90, 91, 3, 22, 11, 0, 91, 93,
		5, 8, 0, 0, 92, 94, 3, 28, 14, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0,
		0, 94, 11, 1, 0, 0, 0, 95, 97, 3, 28, 14, 0, 96, 95, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 100, 3, 22, 11, 0, 99, 98, 1, 0, 0,
		0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 5, 8, 0, 0, 102,
		103, 3, 28, 14, 0, 103, 104, 5, 9, 0, 0, 104, 13, 1, 0, 0, 0, 105, 106,
		5, 15, 0, 0, 106, 108, 3, 16, 8, 0, 107, 109, 3, 18, 9, 0, 108, 107, 1,
		0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 5, 25, 0,
		0, 111, 15, 1, 0, 0, 0, 112, 114, 5, 26, 0, 0, 113, 112, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 3, 24, 12, 0, 116, 17,
		1, 0, 0, 0, 117, 118, 5, 24, 0, 0, 118, 123, 3, 20, 10, 0, 119, 120, 5,
		33, 0, 0, 120, 122, 3, 20, 10, 0, 121, 119, 1, 0, 0, 0, 122, 125, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 126, 127, 5, 34, 0, 0, 127, 19, 1, 0, 0, 0, 128,
		140, 3, 26, 13, 0, 129, 130, 3, 26, 13, 0, 130, 131, 5, 30, 0, 0, 131,
		136, 3, 26, 13, 0, 132, 133, 5, 31, 0, 0, 133, 135, 3, 26, 13, 0, 134,
		132, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137,
		1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 128, 1, 0,
		0, 0, 139, 129, 1, 0, 0, 0, 140, 21, 1, 0, 0, 0, 141, 142, 5, 14, 0, 0,
		142, 148, 3, 24, 12, 0, 143, 144, 3, 30, 15, 0, 144, 145, 3, 24, 12, 0,
		145, 147, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 151, 152, 5, 23, 0, 0, 152, 23, 1, 0, 0, 0, 153, 156, 5, 26,
		0, 0, 154, 155, 5, 22, 0, 0, 155, 157, 5, 26, 0, 0, 156, 154, 1, 0, 0,
		0, 156, 157, 1, 0, 0, 0, 157, 25, 1, 0, 0, 0, 158, 161, 5, 35, 0, 0, 159,
		160, 5, 32, 0, 0, 160, 162, 5, 35, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162,
		1, 0, 0, 0, 162, 27, 1, 0, 0, 0, 163, 166, 5, 12, 0, 0, 164, 165, 5, 10,
		0, 0, 165, 167, 5, 13, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0,
		167, 29, 1, 0, 0, 0, 168, 169, 7, 0, 0, 0, 169, 31, 1, 0, 0, 0, 20, 35,
		45, 55, 70, 74, 77, 81, 86, 93, 96, 99, 108, 113, 123, 136, 139, 148, 156,
		161, 166,
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
	flowDslParserRULE_imports           = 1
	flowDslParserRULE_flow              = 2
	flowDslParserRULE_statement         = 3
	flowDslParserRULE_statementMiddle   = 4
	flowDslParserRULE_statementStart    = 5
	flowDslParserRULE_statementEnd      = 6
	flowDslParserRULE_component         = 7
	flowDslParserRULE_componentTypeName = 8
	flowDslParserRULE_plugin            = 9
	flowDslParserRULE_pluginPart        = 10
	flowDslParserRULE_data              = 11
	flowDslParserRULE_packageIDI        = 12
	flowDslParserRULE_packageIDP        = 13
	flowDslParserRULE_port              = 14
	flowDslParserRULE_dataSep           = 15
)

// IFlowsContext is an interface to support dynamic dispatch.
type IFlowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFlow() []IFlowContext
	Flow(i int) IFlowContext

	// IsFlowsContext differentiates from other interfaces.
	IsFlowsContext()
}

type FlowsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *FlowsContext) EOF() antlr.TerminalNode {
	return s.GetToken(flowDslParserEOF, 0)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserFLOW {
		{
			p.SetState(32)
			p.Flow()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
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

	// Getter signatures
	IMPORT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ImportsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(flowDslParserIMPORT, 0)
}

func (s *ImportsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACE, 0)
}

func (s *ImportsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACE, 0)
}

func (s *ImportsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSTRING)
}

func (s *ImportsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSTRING, i)
}

func (s *ImportsContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserSEMI)
}

func (s *ImportsContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMI, i)
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
		p.SetState(39)
		p.Match(flowDslParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(flowDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserSTRING {
		{
			p.SetState(41)
			p.Match(flowDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(flowDslParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
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

	// Getter signatures
	FLOW() antlr.TerminalNode
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFlowContext differentiates from other interfaces.
	IsFlowContext()
}

type FlowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *FlowContext) FLOW() antlr.TerminalNode {
	return s.GetToken(flowDslParserFLOW, 0)
}

func (s *FlowContext) ID() antlr.TerminalNode {
	return s.GetToken(flowDslParserID, 0)
}

func (s *FlowContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACE, 0)
}

func (s *FlowContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACE, 0)
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
		p.SetState(49)
		p.Match(flowDslParserFLOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(flowDslParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(flowDslParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowDslParserNAME || _la == flowDslParserLBRACK {
		{
			p.SetState(52)
			p.Statement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
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
	StatementStart() IStatementStartContext
	StatementMiddle() IStatementMiddleContext
	StatementEnd() IStatementEndContext
	SEMI() antlr.TerminalNode

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

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMI, 0)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.StatementStart()
		}
		{
			p.SetState(60)
			p.StatementMiddle()
		}
		{
			p.SetState(61)
			p.StatementEnd()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.StatementStart()
		}
		{
			p.SetState(64)
			p.StatementMiddle()
		}
		{
			p.SetState(65)
			p.Match(flowDslParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.StatementMiddle()
		}
		{
			p.SetState(68)
			p.StatementEnd()
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

// IStatementMiddleContext is an interface to support dynamic dispatch.
type IStatementMiddleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComponent() []IComponentContext
	Component(i int) IComponentContext
	AllARROW() []antlr.TerminalNode
	ARROW(i int) antlr.TerminalNode
	AllPort() []IPortContext
	Port(i int) IPortContext
	AllData() []IDataContext
	Data(i int) IDataContext

	// IsStatementMiddleContext differentiates from other interfaces.
	IsStatementMiddleContext()
}

type StatementMiddleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *StatementMiddleContext) AllComponent() []IComponentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComponentContext); ok {
			len++
		}
	}

	tst := make([]IComponentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComponentContext); ok {
			tst[i] = t.(IComponentContext)
			i++
		}
	}

	return tst
}

func (s *StatementMiddleContext) Component(i int) IComponentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentContext); ok {
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

	return t.(IComponentContext)
}

func (s *StatementMiddleContext) AllARROW() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserARROW)
}

func (s *StatementMiddleContext) ARROW(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, i)
}

func (s *StatementMiddleContext) AllPort() []IPortContext {
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

func (s *StatementMiddleContext) Port(i int) IPortContext {
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

func (s *StatementMiddleContext) AllData() []IDataContext {
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

func (s *StatementMiddleContext) Data(i int) IDataContext {
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

func (s *StatementMiddleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementMiddleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) StatementMiddle() (localctx IStatementMiddleContext) {
	localctx = NewStatementMiddleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, flowDslParserRULE_statementMiddle)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Component()
	}
	p.SetState(86)
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
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == flowDslParserNAME {
				{
					p.SetState(73)
					p.Port()
				}

			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == flowDslParserLPAREN {
				{
					p.SetState(76)
					p.Data()
				}

			}
			{
				p.SetState(79)
				p.Match(flowDslParserARROW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == flowDslParserNAME {
				{
					p.SetState(80)
					p.Port()
				}

			}
			{
				p.SetState(83)
				p.Component()
			}

		}
		p.SetState(88)
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

// IStatementStartContext is an interface to support dynamic dispatch.
type IStatementStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPort() []IPortContext
	Port(i int) IPortContext
	Data() IDataContext
	ARROW() antlr.TerminalNode

	// IsStatementStartContext differentiates from other interfaces.
	IsStatementStartContext()
}

type StatementStartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *StatementStartContext) Data() IDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *StatementStartContext) ARROW() antlr.TerminalNode {
	return s.GetToken(flowDslParserARROW, 0)
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
		p.SetState(89)
		p.Port()
	}
	{
		p.SetState(90)
		p.Data()
	}
	{
		p.SetState(91)
		p.Match(flowDslParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(92)
			p.Port()
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

	// Getter signatures
	ARROW() antlr.TerminalNode
	AllPort() []IPortContext
	Port(i int) IPortContext
	SEMI() antlr.TerminalNode
	Data() IDataContext

	// IsStatementEndContext differentiates from other interfaces.
	IsStatementEndContext()
}

type StatementEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *StatementEndContext) SEMI() antlr.TerminalNode {
	return s.GetToken(flowDslParserSEMI, 0)
}

func (s *StatementEndContext) Data() IDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataContext)
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
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserNAME {
		{
			p.SetState(95)
			p.Port()
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
			p.Data()
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
	{
		p.SetState(102)
		p.Port()
	}
	{
		p.SetState(103)
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

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	ComponentTypeName() IComponentTypeNameContext
	RBRACK() antlr.TerminalNode
	Plugin() IPluginContext

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ComponentContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACK, 0)
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

func (s *ComponentContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACK, 0)
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
	p.EnterRule(localctx, 14, flowDslParserRULE_component)
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
		p.ComponentTypeName()
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
			p.Plugin()
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

	// Getter signatures
	PackageIDI() IPackageIDIContext
	IDI() antlr.TerminalNode

	// IsComponentTypeNameContext differentiates from other interfaces.
	IsComponentTypeNameContext()
}

type ComponentTypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ComponentTypeNameContext) PackageIDI() IPackageIDIContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageIDIContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageIDIContext)
}

func (s *ComponentTypeNameContext) IDI() antlr.TerminalNode {
	return s.GetToken(flowDslParserIDI, 0)
}

func (s *ComponentTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) ComponentTypeName() (localctx IComponentTypeNameContext) {
	localctx = NewComponentTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, flowDslParserRULE_componentTypeName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(112)
			p.Match(flowDslParserIDI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(115)
		p.PackageIDI()
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

	// Getter signatures
	LBRACKI() antlr.TerminalNode
	AllPluginPart() []IPluginPartContext
	PluginPart(i int) IPluginPartContext
	RBRACKP() antlr.TerminalNode
	AllPIPEP() []antlr.TerminalNode
	PIPEP(i int) antlr.TerminalNode

	// IsPluginContext differentiates from other interfaces.
	IsPluginContext()
}

type PluginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *PluginContext) LBRACKI() antlr.TerminalNode {
	return s.GetToken(flowDslParserLBRACKI, 0)
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

func (s *PluginContext) RBRACKP() antlr.TerminalNode {
	return s.GetToken(flowDslParserRBRACKP, 0)
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
	p.EnterRule(localctx, 18, flowDslParserRULE_plugin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(flowDslParserLBRACKI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.PluginPart()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == flowDslParserPIPEP {
		{
			p.SetState(119)
			p.Match(flowDslParserPIPEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.PluginPart()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
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
	parser antlr.Parser
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
	p.EnterRule(localctx, 20, flowDslParserRULE_pluginPart)
	var _la int

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.PackageIDP()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.PackageIDP()
		}
		{
			p.SetState(130)
			p.Match(flowDslParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.PackageIDP()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == flowDslParserCOMMAP {
			{
				p.SetState(132)
				p.Match(flowDslParserCOMMAP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(133)
				p.PackageIDP()
			}

			p.SetState(138)
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

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllPackageIDI() []IPackageIDIContext
	PackageIDI(i int) IPackageIDIContext
	RPAREN() antlr.TerminalNode
	AllDataSep() []IDataSepContext
	DataSep(i int) IDataSepContext

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *DataContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(flowDslParserLPAREN, 0)
}

func (s *DataContext) AllPackageIDI() []IPackageIDIContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageIDIContext); ok {
			len++
		}
	}

	tst := make([]IPackageIDIContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageIDIContext); ok {
			tst[i] = t.(IPackageIDIContext)
			i++
		}
	}

	return tst
}

func (s *DataContext) PackageIDI(i int) IPackageIDIContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageIDIContext); ok {
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

	return t.(IPackageIDIContext)
}

func (s *DataContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(flowDslParserRPAREN, 0)
}

func (s *DataContext) AllDataSep() []IDataSepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataSepContext); ok {
			len++
		}
	}

	tst := make([]IDataSepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataSepContext); ok {
			tst[i] = t.(IDataSepContext)
			i++
		}
	}

	return tst
}

func (s *DataContext) DataSep(i int) IDataSepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSepContext); ok {
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

	return t.(IDataSepContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, flowDslParserRULE_data)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(flowDslParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.PackageIDI()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == flowDslParserCOMMA || _la == flowDslParserPIPE {
		{
			p.SetState(143)
			p.DataSep()
		}
		{
			p.SetState(144)
			p.PackageIDI()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
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

// IPackageIDIContext is an interface to support dynamic dispatch.
type IPackageIDIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDI() []antlr.TerminalNode
	IDI(i int) antlr.TerminalNode
	DOTI() antlr.TerminalNode

	// IsPackageIDIContext differentiates from other interfaces.
	IsPackageIDIContext()
}

type PackageIDIContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *PackageIDIContext) AllIDI() []antlr.TerminalNode {
	return s.GetTokens(flowDslParserIDI)
}

func (s *PackageIDIContext) IDI(i int) antlr.TerminalNode {
	return s.GetToken(flowDslParserIDI, i)
}

func (s *PackageIDIContext) DOTI() antlr.TerminalNode {
	return s.GetToken(flowDslParserDOTI, 0)
}

func (s *PackageIDIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageIDIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) PackageIDI() (localctx IPackageIDIContext) {
	localctx = NewPackageIDIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, flowDslParserRULE_packageIDI)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(flowDslParserIDI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserDOTI {
		{
			p.SetState(154)
			p.Match(flowDslParserDOTI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(flowDslParserIDI)
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

// IPackageIDPContext is an interface to support dynamic dispatch.
type IPackageIDPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.EnterRule(localctx, 26, flowDslParserRULE_packageIDP)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(flowDslParserIDP)
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

	if _la == flowDslParserDOTP {
		{
			p.SetState(159)
			p.Match(flowDslParserDOTP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(flowDslParserIDP)
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
	p.EnterRule(localctx, 28, flowDslParserRULE_port)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(flowDslParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == flowDslParserCOLON {
		{
			p.SetState(164)
			p.Match(flowDslParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(flowDslParserINT)
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

// IDataSepContext is an interface to support dynamic dispatch.
type IDataSepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPE() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsDataSepContext differentiates from other interfaces.
	IsDataSepContext()
}

type DataSepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSepContext() *DataSepContext {
	var p = new(DataSepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_dataSep
	return p
}

func InitEmptyDataSepContext(p *DataSepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = flowDslParserRULE_dataSep
}

func (*DataSepContext) IsDataSepContext() {}

func NewDataSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSepContext {
	var p = new(DataSepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowDslParserRULE_dataSep

	return p
}

func (s *DataSepContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSepContext) PIPE() antlr.TerminalNode {
	return s.GetToken(flowDslParserPIPE, 0)
}

func (s *DataSepContext) COMMA() antlr.TerminalNode {
	return s.GetToken(flowDslParserCOMMA, 0)
}

func (s *DataSepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *flowDslParser) DataSep() (localctx IDataSepContext) {
	localctx = NewDataSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, flowDslParserRULE_dataSep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !(_la == flowDslParserCOMMA || _la == flowDslParserPIPE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
