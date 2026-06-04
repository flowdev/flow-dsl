// Code generated from ./flowDslLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type flowDslLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FlowDslLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func flowdsllexerLexerInit() {
	staticData := &FlowDslLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "F", "I", "P",
	}
	staticData.LiteralNames = []string{
		"", "'flow'", "'{'", "", "", "", "", "'->'", "';'", "':'", "", "", "'('",
		"", "'}'", "", "", "", "", "", "", "')'", "", "", "", "", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "FLOW", "LBRACE", "ID", "BLOCK_COMMENT", "LINE_COMMENT", "WS", "ARROW",
		"SEMI", "COLON", "NAME", "INT", "LPAREN", "LBRACK", "RBRACE", "BLOCK_COMMENTF",
		"LINE_COMMENTF", "WSF", "COMMA", "PIPE", "DOTI", "RPAREN", "LBRACKI",
		"RBRACK", "IDI", "BLOCK_COMMENTI", "LINE_COMMENTI", "WSI", "ASSIGN",
		"COMMAP", "DOTP", "PIPEP", "RBRACKP", "IDP", "BLOCK_COMMENTP", "LINE_COMMENTP",
		"WSP",
	}
	staticData.RuleNames = []string{
		"FLOW", "LBRACE", "ID", "BLOCK_COMMENT", "LINE_COMMENT", "WS", "ARROW",
		"SEMI", "COLON", "NAME", "INT", "LPAREN", "LBRACK", "RBRACE", "BLOCK_COMMENTF",
		"LINE_COMMENTF", "WSF", "COMMA", "PIPE", "DOTI", "RPAREN", "LBRACKI",
		"RBRACK", "IDI", "BLOCK_COMMENTI", "LINE_COMMENTI", "WSI", "ASSIGN",
		"COMMAP", "DOTP", "PIPEP", "RBRACKP", "IDP", "BLOCK_COMMENTP", "LINE_COMMENTP",
		"WSP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 283, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2,
		7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8,
		7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13,
		2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2,
		19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24,
		7, 24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7,
		29, 2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34,
		2, 35, 7, 35, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 5, 2, 88, 8, 2, 10, 2, 12, 2, 91, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 97, 8, 3, 10, 3, 12, 3, 100, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 111, 8, 4, 10, 4, 12, 4, 114, 9, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 131, 8, 9, 10, 9, 12, 9, 134, 9, 9, 1, 10, 4, 10, 137,
		8, 10, 11, 10, 12, 10, 138, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14,
		157, 8, 14, 10, 14, 12, 14, 160, 9, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 171, 8, 15, 10, 15, 12, 15, 174,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 5, 23, 202, 8, 23, 10, 23, 12,
		23, 205, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 211, 8, 24, 10, 24,
		12, 24, 214, 9, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 5, 25, 225, 8, 25, 10, 25, 12, 25, 228, 9, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 250, 8, 32,
		10, 32, 12, 32, 253, 9, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 259, 8,
		33, 10, 33, 12, 33, 262, 9, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 34, 5, 34, 273, 8, 34, 10, 34, 12, 34, 276, 9, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 4, 98, 158, 212, 260, 0, 36, 4,
		1, 6, 2, 8, 3, 10, 4, 12, 5, 14, 6, 16, 7, 18, 8, 20, 9, 22, 10, 24, 11,
		26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 17, 38, 18, 40, 19, 42, 20,
		44, 21, 46, 22, 48, 23, 50, 24, 52, 25, 54, 26, 56, 27, 58, 28, 60, 29,
		62, 30, 64, 31, 66, 32, 68, 33, 70, 34, 72, 35, 74, 36, 4, 0, 1, 2, 3,
		7, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		2, 0, 10, 10, 13, 13, 10, 0, 9, 13, 32, 32, 133, 133, 160, 160, 5760, 5760,
		8192, 8202, 8232, 8233, 8239, 8239, 8287, 8287, 12288, 12288, 1, 0, 97,
		122, 3, 0, 48, 57, 65, 90, 97, 122, 1, 0, 48, 57, 292, 0, 4, 1, 0, 0, 0,
		0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0,
		0, 0, 14, 1, 0, 0, 0, 1, 16, 1, 0, 0, 0, 1, 18, 1, 0, 0, 0, 1, 20, 1, 0,
		0, 0, 1, 22, 1, 0, 0, 0, 1, 24, 1, 0, 0, 0, 1, 26, 1, 0, 0, 0, 1, 28, 1,
		0, 0, 0, 1, 30, 1, 0, 0, 0, 1, 32, 1, 0, 0, 0, 1, 34, 1, 0, 0, 0, 1, 36,
		1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 2, 42, 1, 0, 0, 0, 2,
		44, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0,
		2, 52, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 3, 58, 1, 0, 0,
		0, 3, 60, 1, 0, 0, 0, 3, 62, 1, 0, 0, 0, 3, 64, 1, 0, 0, 0, 3, 66, 1, 0,
		0, 0, 3, 68, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 3, 72, 1, 0, 0, 0, 3, 74, 1,
		0, 0, 0, 4, 76, 1, 0, 0, 0, 6, 81, 1, 0, 0, 0, 8, 85, 1, 0, 0, 0, 10, 92,
		1, 0, 0, 0, 12, 106, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 121, 1, 0, 0,
		0, 18, 124, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 128, 1, 0, 0, 0, 24, 136,
		1, 0, 0, 0, 26, 140, 1, 0, 0, 0, 28, 144, 1, 0, 0, 0, 30, 148, 1, 0, 0,
		0, 32, 152, 1, 0, 0, 0, 34, 166, 1, 0, 0, 0, 36, 177, 1, 0, 0, 0, 38, 181,
		1, 0, 0, 0, 40, 183, 1, 0, 0, 0, 42, 185, 1, 0, 0, 0, 44, 187, 1, 0, 0,
		0, 46, 191, 1, 0, 0, 0, 48, 195, 1, 0, 0, 0, 50, 199, 1, 0, 0, 0, 52, 206,
		1, 0, 0, 0, 54, 220, 1, 0, 0, 0, 56, 231, 1, 0, 0, 0, 58, 235, 1, 0, 0,
		0, 60, 237, 1, 0, 0, 0, 62, 239, 1, 0, 0, 0, 64, 241, 1, 0, 0, 0, 66, 243,
		1, 0, 0, 0, 68, 247, 1, 0, 0, 0, 70, 254, 1, 0, 0, 0, 72, 268, 1, 0, 0,
		0, 74, 279, 1, 0, 0, 0, 76, 77, 5, 102, 0, 0, 77, 78, 5, 108, 0, 0, 78,
		79, 5, 111, 0, 0, 79, 80, 5, 119, 0, 0, 80, 5, 1, 0, 0, 0, 81, 82, 5, 123,
		0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6, 1, 0, 0, 84, 7, 1, 0, 0, 0, 85, 89,
		7, 0, 0, 0, 86, 88, 7, 1, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0,
		89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 9, 1, 0, 0, 0, 91, 89, 1, 0,
		0, 0, 92, 93, 5, 47, 0, 0, 93, 94, 5, 42, 0, 0, 94, 98, 1, 0, 0, 0, 95,
		97, 9, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 98, 96, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		102, 5, 42, 0, 0, 102, 103, 5, 47, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105,
		6, 3, 1, 0, 105, 11, 1, 0, 0, 0, 106, 107, 5, 47, 0, 0, 107, 108, 5, 47,
		0, 0, 108, 112, 1, 0, 0, 0, 109, 111, 8, 2, 0, 0, 110, 109, 1, 0, 0, 0,
		111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 6, 4, 1, 0, 116, 13, 1,
		0, 0, 0, 117, 118, 7, 3, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 6, 5, 1,
		0, 120, 15, 1, 0, 0, 0, 121, 122, 5, 45, 0, 0, 122, 123, 5, 62, 0, 0, 123,
		17, 1, 0, 0, 0, 124, 125, 5, 59, 0, 0, 125, 19, 1, 0, 0, 0, 126, 127, 5,
		58, 0, 0, 127, 21, 1, 0, 0, 0, 128, 132, 7, 4, 0, 0, 129, 131, 7, 5, 0,
		0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 23, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 137, 7,
		6, 0, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 25, 1, 0, 0, 0, 140, 141, 5, 40, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 6, 11, 2, 0, 143, 27, 1, 0, 0, 0, 144, 145,
		5, 91, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 6, 12, 2, 0, 147, 29, 1, 0,
		0, 0, 148, 149, 5, 125, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 6, 13, 3,
		0, 151, 31, 1, 0, 0, 0, 152, 153, 5, 47, 0, 0, 153, 154, 5, 42, 0, 0, 154,
		158, 1, 0, 0, 0, 155, 157, 9, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 160,
		1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 161, 1, 0,
		0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 5, 42, 0, 0, 162, 163, 5, 47, 0,
		0, 163, 164, 1, 0, 0, 0, 164, 165, 6, 14, 1, 0, 165, 33, 1, 0, 0, 0, 166,
		167, 5, 47, 0, 0, 167, 168, 5, 47, 0, 0, 168, 172, 1, 0, 0, 0, 169, 171,
		8, 2, 0, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0,
		0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0,
		175, 176, 6, 15, 1, 0, 176, 35, 1, 0, 0, 0, 177, 178, 7, 3, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 180, 6, 16, 1, 0, 180, 37, 1, 0, 0, 0, 181, 182,
		5, 44, 0, 0, 182, 39, 1, 0, 0, 0, 183, 184, 5, 124, 0, 0, 184, 41, 1, 0,
		0, 0, 185, 186, 5, 46, 0, 0, 186, 43, 1, 0, 0, 0, 187, 188, 5, 41, 0, 0,
		188, 189, 1, 0, 0, 0, 189, 190, 6, 20, 3, 0, 190, 45, 1, 0, 0, 0, 191,
		192, 5, 91, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 6, 21, 4, 0, 194, 47,
		1, 0, 0, 0, 195, 196, 5, 93, 0, 0, 196, 197, 1, 0, 0, 0, 197, 198, 6, 22,
		3, 0, 198, 49, 1, 0, 0, 0, 199, 203, 7, 0, 0, 0, 200, 202, 7, 1, 0, 0,
		201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 51, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207, 5,
		47, 0, 0, 207, 208, 5, 42, 0, 0, 208, 212, 1, 0, 0, 0, 209, 211, 9, 0,
		0, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0,
		212, 210, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215,
		216, 5, 42, 0, 0, 216, 217, 5, 47, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219,
		6, 24, 1, 0, 219, 53, 1, 0, 0, 0, 220, 221, 5, 47, 0, 0, 221, 222, 5, 47,
		0, 0, 222, 226, 1, 0, 0, 0, 223, 225, 8, 2, 0, 0, 224, 223, 1, 0, 0, 0,
		225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		229, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 6, 25, 1, 0, 230, 55,
		1, 0, 0, 0, 231, 232, 7, 3, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 6, 26,
		1, 0, 234, 57, 1, 0, 0, 0, 235, 236, 5, 61, 0, 0, 236, 59, 1, 0, 0, 0,
		237, 238, 5, 44, 0, 0, 238, 61, 1, 0, 0, 0, 239, 240, 5, 46, 0, 0, 240,
		63, 1, 0, 0, 0, 241, 242, 5, 124, 0, 0, 242, 65, 1, 0, 0, 0, 243, 244,
		5, 93, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 6, 31, 3, 0, 246, 67, 1, 0,
		0, 0, 247, 251, 7, 0, 0, 0, 248, 250, 7, 1, 0, 0, 249, 248, 1, 0, 0, 0,
		250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		69, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 5, 47, 0, 0, 255, 256,
		5, 42, 0, 0, 256, 260, 1, 0, 0, 0, 257, 259, 9, 0, 0, 0, 258, 257, 1, 0,
		0, 0, 259, 262, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0,
		261, 263, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5, 42, 0, 0, 264,
		265, 5, 47, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 6, 33, 1, 0, 267, 71,
		1, 0, 0, 0, 268, 269, 5, 47, 0, 0, 269, 270, 5, 47, 0, 0, 270, 274, 1,
		0, 0, 0, 271, 273, 8, 2, 0, 0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0,
		0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276,
		274, 1, 0, 0, 0, 277, 278, 6, 34, 1, 0, 278, 73, 1, 0, 0, 0, 279, 280,
		7, 3, 0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 6, 35, 1, 0, 282, 75, 1, 0,
		0, 0, 17, 0, 1, 2, 3, 89, 98, 112, 132, 138, 158, 172, 203, 212, 226, 251,
		260, 274, 5, 5, 1, 0, 0, 1, 0, 5, 2, 0, 4, 0, 0, 5, 3, 0,
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

// flowDslLexerInit initializes any static state used to implement flowDslLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewflowDslLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FlowDslLexerInit() {
	staticData := &FlowDslLexerLexerStaticData
	staticData.once.Do(flowdsllexerLexerInit)
}

// NewflowDslLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewflowDslLexer(input antlr.CharStream) *flowDslLexer {
	FlowDslLexerInit()
	l := new(flowDslLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FlowDslLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "flowDslLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// flowDslLexer tokens.
const (
	flowDslLexerFLOW           = 1
	flowDslLexerLBRACE         = 2
	flowDslLexerID             = 3
	flowDslLexerBLOCK_COMMENT  = 4
	flowDslLexerLINE_COMMENT   = 5
	flowDslLexerWS             = 6
	flowDslLexerARROW          = 7
	flowDslLexerSEMI           = 8
	flowDslLexerCOLON          = 9
	flowDslLexerNAME           = 10
	flowDslLexerINT            = 11
	flowDslLexerLPAREN         = 12
	flowDslLexerLBRACK         = 13
	flowDslLexerRBRACE         = 14
	flowDslLexerBLOCK_COMMENTF = 15
	flowDslLexerLINE_COMMENTF  = 16
	flowDslLexerWSF            = 17
	flowDslLexerCOMMA          = 18
	flowDslLexerPIPE           = 19
	flowDslLexerDOTI           = 20
	flowDslLexerRPAREN         = 21
	flowDslLexerLBRACKI        = 22
	flowDslLexerRBRACK         = 23
	flowDslLexerIDI            = 24
	flowDslLexerBLOCK_COMMENTI = 25
	flowDslLexerLINE_COMMENTI  = 26
	flowDslLexerWSI            = 27
	flowDslLexerASSIGN         = 28
	flowDslLexerCOMMAP         = 29
	flowDslLexerDOTP           = 30
	flowDslLexerPIPEP          = 31
	flowDslLexerRBRACKP        = 32
	flowDslLexerIDP            = 33
	flowDslLexerBLOCK_COMMENTP = 34
	flowDslLexerLINE_COMMENTP  = 35
	flowDslLexerWSP            = 36
)

// flowDslLexer modes.
const (
	flowDslLexerF = iota + 1
	flowDslLexerI
	flowDslLexerP
)
