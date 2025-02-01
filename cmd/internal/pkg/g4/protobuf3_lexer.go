// Code generated from Protobuf3.g4 by ANTLR 4.13.1. DO NOT EDIT.

package g4

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

type Protobuf3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Protobuf3LexerLexerStaticData struct {
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

func protobuf3lexerLexerInit() {
	staticData := &Protobuf3LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
		"'optional'", "'repeated'", "'oneof'", "'map'", "'int32'", "'int64'",
		"'uint32'", "'uint64'", "'sint32'", "'sint64'", "'fixed32'", "'fixed64'",
		"'sfixed32'", "'sfixed64'", "'bool'", "'string'", "'double'", "'float'",
		"'bytes'", "'reserved'", "'to'", "'max'", "'enum'", "'message'", "'service'",
		"'extend'", "'rpc'", "'stream'", "'returns'", "'\"proto3\"'", "''proto3''",
		"';'", "'='", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'",
		"'.'", "','", "':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "OPTIONAL",
		"REPEATED", "ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32",
		"SINT64", "FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING",
		"DOUBLE", "FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE",
		"SERVICE", "EXTEND", "RPC", "STREAM", "RETURNS", "PROTO3_LIT_SINGLE",
		"PROTO3_LIT_DOBULE", "SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC",
		"LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT",
		"FLOAT_LIT", "INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "OPTIONAL",
		"REPEATED", "ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32",
		"SINT64", "FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING",
		"DOUBLE", "FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE",
		"SERVICE", "EXTEND", "RPC", "STREAM", "RETURNS", "PROTO3_LIT_SINGLE",
		"PROTO3_LIT_DOBULE", "SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC",
		"LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "CHAR_VALUE",
		"HEX_ESCAPE", "OCT_ESCAPE", "CHAR_ESCAPE", "BOOL_LIT", "FLOAT_LIT",
		"EXPONENT", "DECIMALS", "INT_LIT", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT",
		"IDENTIFIER", "LETTER", "DECIMAL_DIGIT", "OCTAL_DIGIT", "HEX_DIGIT",
		"WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 60, 592, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 5, 52, 435, 8, 52, 10, 52, 12, 52, 438,
		9, 52, 1, 52, 1, 52, 1, 52, 5, 52, 443, 8, 52, 10, 52, 12, 52, 446, 9,
		52, 1, 52, 3, 52, 449, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 455, 8,
		53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 3, 57, 479, 8, 57, 1, 58, 1, 58, 1, 58, 3, 58, 484, 8, 58, 1,
		58, 3, 58, 487, 8, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 3, 58,
		495, 8, 58, 3, 58, 497, 8, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58,
		3, 58, 505, 8, 58, 1, 59, 1, 59, 1, 59, 3, 59, 510, 8, 59, 1, 59, 1, 59,
		1, 60, 4, 60, 515, 8, 60, 11, 60, 12, 60, 516, 1, 61, 1, 61, 1, 61, 3,
		61, 522, 8, 61, 1, 62, 1, 62, 5, 62, 526, 8, 62, 10, 62, 12, 62, 529, 9,
		62, 1, 63, 1, 63, 5, 63, 533, 8, 63, 10, 63, 12, 63, 536, 9, 63, 1, 64,
		1, 64, 1, 64, 4, 64, 541, 8, 64, 11, 64, 12, 64, 542, 1, 65, 1, 65, 1,
		65, 5, 65, 548, 8, 65, 10, 65, 12, 65, 551, 9, 65, 1, 66, 1, 66, 1, 67,
		1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 4, 70, 562, 8, 70, 11, 70, 12,
		70, 563, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 572, 8, 71, 10,
		71, 12, 71, 575, 9, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 72, 5, 72,
		583, 8, 72, 10, 72, 12, 72, 586, 9, 72, 1, 72, 1, 72, 1, 72, 1, 72, 1,
		72, 3, 436, 444, 584, 0, 73, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87,
		44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 0, 109, 0, 111, 0, 113, 0, 115, 54, 117, 55, 119, 0, 121, 0, 123,
		56, 125, 0, 127, 0, 129, 0, 131, 57, 133, 0, 135, 0, 137, 0, 139, 0, 141,
		58, 143, 59, 145, 60, 1, 0, 11, 3, 0, 0, 0, 10, 10, 92, 92, 2, 0, 88, 88,
		120, 120, 9, 0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 118, 118, 2, 0, 69, 69, 101, 101, 1, 0, 49, 57, 3, 0, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70,
		97, 102, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 10, 10, 13, 13, 605, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1,
		0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0,
		0, 117, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 141, 1,
		0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 1, 147, 1, 0, 0, 0, 3,
		154, 1, 0, 0, 0, 5, 161, 1, 0, 0, 0, 7, 166, 1, 0, 0, 0, 9, 173, 1, 0,
		0, 0, 11, 181, 1, 0, 0, 0, 13, 188, 1, 0, 0, 0, 15, 197, 1, 0, 0, 0, 17,
		206, 1, 0, 0, 0, 19, 212, 1, 0, 0, 0, 21, 216, 1, 0, 0, 0, 23, 222, 1,
		0, 0, 0, 25, 228, 1, 0, 0, 0, 27, 235, 1, 0, 0, 0, 29, 242, 1, 0, 0, 0,
		31, 249, 1, 0, 0, 0, 33, 256, 1, 0, 0, 0, 35, 264, 1, 0, 0, 0, 37, 272,
		1, 0, 0, 0, 39, 281, 1, 0, 0, 0, 41, 290, 1, 0, 0, 0, 43, 295, 1, 0, 0,
		0, 45, 302, 1, 0, 0, 0, 47, 309, 1, 0, 0, 0, 49, 315, 1, 0, 0, 0, 51, 321,
		1, 0, 0, 0, 53, 330, 1, 0, 0, 0, 55, 333, 1, 0, 0, 0, 57, 337, 1, 0, 0,
		0, 59, 342, 1, 0, 0, 0, 61, 350, 1, 0, 0, 0, 63, 358, 1, 0, 0, 0, 65, 365,
		1, 0, 0, 0, 67, 369, 1, 0, 0, 0, 69, 376, 1, 0, 0, 0, 71, 384, 1, 0, 0,
		0, 73, 393, 1, 0, 0, 0, 75, 402, 1, 0, 0, 0, 77, 404, 1, 0, 0, 0, 79, 406,
		1, 0, 0, 0, 81, 408, 1, 0, 0, 0, 83, 410, 1, 0, 0, 0, 85, 412, 1, 0, 0,
		0, 87, 414, 1, 0, 0, 0, 89, 416, 1, 0, 0, 0, 91, 418, 1, 0, 0, 0, 93, 420,
		1, 0, 0, 0, 95, 422, 1, 0, 0, 0, 97, 424, 1, 0, 0, 0, 99, 426, 1, 0, 0,
		0, 101, 428, 1, 0, 0, 0, 103, 430, 1, 0, 0, 0, 105, 448, 1, 0, 0, 0, 107,
		454, 1, 0, 0, 0, 109, 456, 1, 0, 0, 0, 111, 461, 1, 0, 0, 0, 113, 466,
		1, 0, 0, 0, 115, 478, 1, 0, 0, 0, 117, 504, 1, 0, 0, 0, 119, 506, 1, 0,
		0, 0, 121, 514, 1, 0, 0, 0, 123, 521, 1, 0, 0, 0, 125, 523, 1, 0, 0, 0,
		127, 530, 1, 0, 0, 0, 129, 537, 1, 0, 0, 0, 131, 544, 1, 0, 0, 0, 133,
		552, 1, 0, 0, 0, 135, 554, 1, 0, 0, 0, 137, 556, 1, 0, 0, 0, 139, 558,
		1, 0, 0, 0, 141, 561, 1, 0, 0, 0, 143, 567, 1, 0, 0, 0, 145, 578, 1, 0,
		0, 0, 147, 148, 5, 115, 0, 0, 148, 149, 5, 121, 0, 0, 149, 150, 5, 110,
		0, 0, 150, 151, 5, 116, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153, 5, 120,
		0, 0, 153, 2, 1, 0, 0, 0, 154, 155, 5, 105, 0, 0, 155, 156, 5, 109, 0,
		0, 156, 157, 5, 112, 0, 0, 157, 158, 5, 111, 0, 0, 158, 159, 5, 114, 0,
		0, 159, 160, 5, 116, 0, 0, 160, 4, 1, 0, 0, 0, 161, 162, 5, 119, 0, 0,
		162, 163, 5, 101, 0, 0, 163, 164, 5, 97, 0, 0, 164, 165, 5, 107, 0, 0,
		165, 6, 1, 0, 0, 0, 166, 167, 5, 112, 0, 0, 167, 168, 5, 117, 0, 0, 168,
		169, 5, 98, 0, 0, 169, 170, 5, 108, 0, 0, 170, 171, 5, 105, 0, 0, 171,
		172, 5, 99, 0, 0, 172, 8, 1, 0, 0, 0, 173, 174, 5, 112, 0, 0, 174, 175,
		5, 97, 0, 0, 175, 176, 5, 99, 0, 0, 176, 177, 5, 107, 0, 0, 177, 178, 5,
		97, 0, 0, 178, 179, 5, 103, 0, 0, 179, 180, 5, 101, 0, 0, 180, 10, 1, 0,
		0, 0, 181, 182, 5, 111, 0, 0, 182, 183, 5, 112, 0, 0, 183, 184, 5, 116,
		0, 0, 184, 185, 5, 105, 0, 0, 185, 186, 5, 111, 0, 0, 186, 187, 5, 110,
		0, 0, 187, 12, 1, 0, 0, 0, 188, 189, 5, 111, 0, 0, 189, 190, 5, 112, 0,
		0, 190, 191, 5, 116, 0, 0, 191, 192, 5, 105, 0, 0, 192, 193, 5, 111, 0,
		0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 97, 0, 0, 195, 196, 5, 108, 0,
		0, 196, 14, 1, 0, 0, 0, 197, 198, 5, 114, 0, 0, 198, 199, 5, 101, 0, 0,
		199, 200, 5, 112, 0, 0, 200, 201, 5, 101, 0, 0, 201, 202, 5, 97, 0, 0,
		202, 203, 5, 116, 0, 0, 203, 204, 5, 101, 0, 0, 204, 205, 5, 100, 0, 0,
		205, 16, 1, 0, 0, 0, 206, 207, 5, 111, 0, 0, 207, 208, 5, 110, 0, 0, 208,
		209, 5, 101, 0, 0, 209, 210, 5, 111, 0, 0, 210, 211, 5, 102, 0, 0, 211,
		18, 1, 0, 0, 0, 212, 213, 5, 109, 0, 0, 213, 214, 5, 97, 0, 0, 214, 215,
		5, 112, 0, 0, 215, 20, 1, 0, 0, 0, 216, 217, 5, 105, 0, 0, 217, 218, 5,
		110, 0, 0, 218, 219, 5, 116, 0, 0, 219, 220, 5, 51, 0, 0, 220, 221, 5,
		50, 0, 0, 221, 22, 1, 0, 0, 0, 222, 223, 5, 105, 0, 0, 223, 224, 5, 110,
		0, 0, 224, 225, 5, 116, 0, 0, 225, 226, 5, 54, 0, 0, 226, 227, 5, 52, 0,
		0, 227, 24, 1, 0, 0, 0, 228, 229, 5, 117, 0, 0, 229, 230, 5, 105, 0, 0,
		230, 231, 5, 110, 0, 0, 231, 232, 5, 116, 0, 0, 232, 233, 5, 51, 0, 0,
		233, 234, 5, 50, 0, 0, 234, 26, 1, 0, 0, 0, 235, 236, 5, 117, 0, 0, 236,
		237, 5, 105, 0, 0, 237, 238, 5, 110, 0, 0, 238, 239, 5, 116, 0, 0, 239,
		240, 5, 54, 0, 0, 240, 241, 5, 52, 0, 0, 241, 28, 1, 0, 0, 0, 242, 243,
		5, 115, 0, 0, 243, 244, 5, 105, 0, 0, 244, 245, 5, 110, 0, 0, 245, 246,
		5, 116, 0, 0, 246, 247, 5, 51, 0, 0, 247, 248, 5, 50, 0, 0, 248, 30, 1,
		0, 0, 0, 249, 250, 5, 115, 0, 0, 250, 251, 5, 105, 0, 0, 251, 252, 5, 110,
		0, 0, 252, 253, 5, 116, 0, 0, 253, 254, 5, 54, 0, 0, 254, 255, 5, 52, 0,
		0, 255, 32, 1, 0, 0, 0, 256, 257, 5, 102, 0, 0, 257, 258, 5, 105, 0, 0,
		258, 259, 5, 120, 0, 0, 259, 260, 5, 101, 0, 0, 260, 261, 5, 100, 0, 0,
		261, 262, 5, 51, 0, 0, 262, 263, 5, 50, 0, 0, 263, 34, 1, 0, 0, 0, 264,
		265, 5, 102, 0, 0, 265, 266, 5, 105, 0, 0, 266, 267, 5, 120, 0, 0, 267,
		268, 5, 101, 0, 0, 268, 269, 5, 100, 0, 0, 269, 270, 5, 54, 0, 0, 270,
		271, 5, 52, 0, 0, 271, 36, 1, 0, 0, 0, 272, 273, 5, 115, 0, 0, 273, 274,
		5, 102, 0, 0, 274, 275, 5, 105, 0, 0, 275, 276, 5, 120, 0, 0, 276, 277,
		5, 101, 0, 0, 277, 278, 5, 100, 0, 0, 278, 279, 5, 51, 0, 0, 279, 280,
		5, 50, 0, 0, 280, 38, 1, 0, 0, 0, 281, 282, 5, 115, 0, 0, 282, 283, 5,
		102, 0, 0, 283, 284, 5, 105, 0, 0, 284, 285, 5, 120, 0, 0, 285, 286, 5,
		101, 0, 0, 286, 287, 5, 100, 0, 0, 287, 288, 5, 54, 0, 0, 288, 289, 5,
		52, 0, 0, 289, 40, 1, 0, 0, 0, 290, 291, 5, 98, 0, 0, 291, 292, 5, 111,
		0, 0, 292, 293, 5, 111, 0, 0, 293, 294, 5, 108, 0, 0, 294, 42, 1, 0, 0,
		0, 295, 296, 5, 115, 0, 0, 296, 297, 5, 116, 0, 0, 297, 298, 5, 114, 0,
		0, 298, 299, 5, 105, 0, 0, 299, 300, 5, 110, 0, 0, 300, 301, 5, 103, 0,
		0, 301, 44, 1, 0, 0, 0, 302, 303, 5, 100, 0, 0, 303, 304, 5, 111, 0, 0,
		304, 305, 5, 117, 0, 0, 305, 306, 5, 98, 0, 0, 306, 307, 5, 108, 0, 0,
		307, 308, 5, 101, 0, 0, 308, 46, 1, 0, 0, 0, 309, 310, 5, 102, 0, 0, 310,
		311, 5, 108, 0, 0, 311, 312, 5, 111, 0, 0, 312, 313, 5, 97, 0, 0, 313,
		314, 5, 116, 0, 0, 314, 48, 1, 0, 0, 0, 315, 316, 5, 98, 0, 0, 316, 317,
		5, 121, 0, 0, 317, 318, 5, 116, 0, 0, 318, 319, 5, 101, 0, 0, 319, 320,
		5, 115, 0, 0, 320, 50, 1, 0, 0, 0, 321, 322, 5, 114, 0, 0, 322, 323, 5,
		101, 0, 0, 323, 324, 5, 115, 0, 0, 324, 325, 5, 101, 0, 0, 325, 326, 5,
		114, 0, 0, 326, 327, 5, 118, 0, 0, 327, 328, 5, 101, 0, 0, 328, 329, 5,
		100, 0, 0, 329, 52, 1, 0, 0, 0, 330, 331, 5, 116, 0, 0, 331, 332, 5, 111,
		0, 0, 332, 54, 1, 0, 0, 0, 333, 334, 5, 109, 0, 0, 334, 335, 5, 97, 0,
		0, 335, 336, 5, 120, 0, 0, 336, 56, 1, 0, 0, 0, 337, 338, 5, 101, 0, 0,
		338, 339, 5, 110, 0, 0, 339, 340, 5, 117, 0, 0, 340, 341, 5, 109, 0, 0,
		341, 58, 1, 0, 0, 0, 342, 343, 5, 109, 0, 0, 343, 344, 5, 101, 0, 0, 344,
		345, 5, 115, 0, 0, 345, 346, 5, 115, 0, 0, 346, 347, 5, 97, 0, 0, 347,
		348, 5, 103, 0, 0, 348, 349, 5, 101, 0, 0, 349, 60, 1, 0, 0, 0, 350, 351,
		5, 115, 0, 0, 351, 352, 5, 101, 0, 0, 352, 353, 5, 114, 0, 0, 353, 354,
		5, 118, 0, 0, 354, 355, 5, 105, 0, 0, 355, 356, 5, 99, 0, 0, 356, 357,
		5, 101, 0, 0, 357, 62, 1, 0, 0, 0, 358, 359, 5, 101, 0, 0, 359, 360, 5,
		120, 0, 0, 360, 361, 5, 116, 0, 0, 361, 362, 5, 101, 0, 0, 362, 363, 5,
		110, 0, 0, 363, 364, 5, 100, 0, 0, 364, 64, 1, 0, 0, 0, 365, 366, 5, 114,
		0, 0, 366, 367, 5, 112, 0, 0, 367, 368, 5, 99, 0, 0, 368, 66, 1, 0, 0,
		0, 369, 370, 5, 115, 0, 0, 370, 371, 5, 116, 0, 0, 371, 372, 5, 114, 0,
		0, 372, 373, 5, 101, 0, 0, 373, 374, 5, 97, 0, 0, 374, 375, 5, 109, 0,
		0, 375, 68, 1, 0, 0, 0, 376, 377, 5, 114, 0, 0, 377, 378, 5, 101, 0, 0,
		378, 379, 5, 116, 0, 0, 379, 380, 5, 117, 0, 0, 380, 381, 5, 114, 0, 0,
		381, 382, 5, 110, 0, 0, 382, 383, 5, 115, 0, 0, 383, 70, 1, 0, 0, 0, 384,
		385, 5, 34, 0, 0, 385, 386, 5, 112, 0, 0, 386, 387, 5, 114, 0, 0, 387,
		388, 5, 111, 0, 0, 388, 389, 5, 116, 0, 0, 389, 390, 5, 111, 0, 0, 390,
		391, 5, 51, 0, 0, 391, 392, 5, 34, 0, 0, 392, 72, 1, 0, 0, 0, 393, 394,
		5, 39, 0, 0, 394, 395, 5, 112, 0, 0, 395, 396, 5, 114, 0, 0, 396, 397,
		5, 111, 0, 0, 397, 398, 5, 116, 0, 0, 398, 399, 5, 111, 0, 0, 399, 400,
		5, 51, 0, 0, 400, 401, 5, 39, 0, 0, 401, 74, 1, 0, 0, 0, 402, 403, 5, 59,
		0, 0, 403, 76, 1, 0, 0, 0, 404, 405, 5, 61, 0, 0, 405, 78, 1, 0, 0, 0,
		406, 407, 5, 40, 0, 0, 407, 80, 1, 0, 0, 0, 408, 409, 5, 41, 0, 0, 409,
		82, 1, 0, 0, 0, 410, 411, 5, 91, 0, 0, 411, 84, 1, 0, 0, 0, 412, 413, 5,
		93, 0, 0, 413, 86, 1, 0, 0, 0, 414, 415, 5, 123, 0, 0, 415, 88, 1, 0, 0,
		0, 416, 417, 5, 125, 0, 0, 417, 90, 1, 0, 0, 0, 418, 419, 5, 60, 0, 0,
		419, 92, 1, 0, 0, 0, 420, 421, 5, 62, 0, 0, 421, 94, 1, 0, 0, 0, 422, 423,
		5, 46, 0, 0, 423, 96, 1, 0, 0, 0, 424, 425, 5, 44, 0, 0, 425, 98, 1, 0,
		0, 0, 426, 427, 5, 58, 0, 0, 427, 100, 1, 0, 0, 0, 428, 429, 5, 43, 0,
		0, 429, 102, 1, 0, 0, 0, 430, 431, 5, 45, 0, 0, 431, 104, 1, 0, 0, 0, 432,
		436, 5, 39, 0, 0, 433, 435, 3, 107, 53, 0, 434, 433, 1, 0, 0, 0, 435, 438,
		1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 439, 1, 0,
		0, 0, 438, 436, 1, 0, 0, 0, 439, 449, 5, 39, 0, 0, 440, 444, 5, 34, 0,
		0, 441, 443, 3, 107, 53, 0, 442, 441, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0,
		444, 445, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 447, 1, 0, 0, 0, 446,
		444, 1, 0, 0, 0, 447, 449, 5, 34, 0, 0, 448, 432, 1, 0, 0, 0, 448, 440,
		1, 0, 0, 0, 449, 106, 1, 0, 0, 0, 450, 455, 3, 109, 54, 0, 451, 455, 3,
		111, 55, 0, 452, 455, 3, 113, 56, 0, 453, 455, 8, 0, 0, 0, 454, 450, 1,
		0, 0, 0, 454, 451, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 454, 453, 1, 0, 0,
		0, 455, 108, 1, 0, 0, 0, 456, 457, 5, 92, 0, 0, 457, 458, 7, 1, 0, 0, 458,
		459, 3, 139, 69, 0, 459, 460, 3, 139, 69, 0, 460, 110, 1, 0, 0, 0, 461,
		462, 5, 92, 0, 0, 462, 463, 3, 137, 68, 0, 463, 464, 3, 137, 68, 0, 464,
		465, 3, 137, 68, 0, 465, 112, 1, 0, 0, 0, 466, 467, 5, 92, 0, 0, 467, 468,
		7, 2, 0, 0, 468, 114, 1, 0, 0, 0, 469, 470, 5, 116, 0, 0, 470, 471, 5,
		114, 0, 0, 471, 472, 5, 117, 0, 0, 472, 479, 5, 101, 0, 0, 473, 474, 5,
		102, 0, 0, 474, 475, 5, 97, 0, 0, 475, 476, 5, 108, 0, 0, 476, 477, 5,
		115, 0, 0, 477, 479, 5, 101, 0, 0, 478, 469, 1, 0, 0, 0, 478, 473, 1, 0,
		0, 0, 479, 116, 1, 0, 0, 0, 480, 481, 3, 121, 60, 0, 481, 483, 3, 95, 47,
		0, 482, 484, 3, 121, 60, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0,
		484, 486, 1, 0, 0, 0, 485, 487, 3, 119, 59, 0, 486, 485, 1, 0, 0, 0, 486,
		487, 1, 0, 0, 0, 487, 497, 1, 0, 0, 0, 488, 489, 3, 121, 60, 0, 489, 490,
		3, 119, 59, 0, 490, 497, 1, 0, 0, 0, 491, 492, 3, 95, 47, 0, 492, 494,
		3, 121, 60, 0, 493, 495, 3, 119, 59, 0, 494, 493, 1, 0, 0, 0, 494, 495,
		1, 0, 0, 0, 495, 497, 1, 0, 0, 0, 496, 480, 1, 0, 0, 0, 496, 488, 1, 0,
		0, 0, 496, 491, 1, 0, 0, 0, 497, 505, 1, 0, 0, 0, 498, 499, 5, 105, 0,
		0, 499, 500, 5, 110, 0, 0, 500, 505, 5, 102, 0, 0, 501, 502, 5, 110, 0,
		0, 502, 503, 5, 97, 0, 0, 503, 505, 5, 110, 0, 0, 504, 496, 1, 0, 0, 0,
		504, 498, 1, 0, 0, 0, 504, 501, 1, 0, 0, 0, 505, 118, 1, 0, 0, 0, 506,
		509, 7, 3, 0, 0, 507, 510, 3, 101, 50, 0, 508, 510, 3, 103, 51, 0, 509,
		507, 1, 0, 0, 0, 509, 508, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 511,
		1, 0, 0, 0, 511, 512, 3, 121, 60, 0, 512, 120, 1, 0, 0, 0, 513, 515, 3,
		135, 67, 0, 514, 513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 514, 1, 0,
		0, 0, 516, 517, 1, 0, 0, 0, 517, 122, 1, 0, 0, 0, 518, 522, 3, 125, 62,
		0, 519, 522, 3, 127, 63, 0, 520, 522, 3, 129, 64, 0, 521, 518, 1, 0, 0,
		0, 521, 519, 1, 0, 0, 0, 521, 520, 1, 0, 0, 0, 522, 124, 1, 0, 0, 0, 523,
		527, 7, 4, 0, 0, 524, 526, 3, 135, 67, 0, 525, 524, 1, 0, 0, 0, 526, 529,
		1, 0, 0, 0, 527, 525, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 126, 1, 0,
		0, 0, 529, 527, 1, 0, 0, 0, 530, 534, 5, 48, 0, 0, 531, 533, 3, 137, 68,
		0, 532, 531, 1, 0, 0, 0, 533, 536, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534,
		535, 1, 0, 0, 0, 535, 128, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 537, 538,
		5, 48, 0, 0, 538, 540, 7, 1, 0, 0, 539, 541, 3, 139, 69, 0, 540, 539, 1,
		0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0, 542, 543, 1, 0, 0,
		0, 543, 130, 1, 0, 0, 0, 544, 549, 3, 133, 66, 0, 545, 548, 3, 133, 66,
		0, 546, 548, 3, 135, 67, 0, 547, 545, 1, 0, 0, 0, 547, 546, 1, 0, 0, 0,
		548, 551, 1, 0, 0, 0, 549, 547, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550,
		132, 1, 0, 0, 0, 551, 549, 1, 0, 0, 0, 552, 553, 7, 5, 0, 0, 553, 134,
		1, 0, 0, 0, 554, 555, 7, 6, 0, 0, 555, 136, 1, 0, 0, 0, 556, 557, 7, 7,
		0, 0, 557, 138, 1, 0, 0, 0, 558, 559, 7, 8, 0, 0, 559, 140, 1, 0, 0, 0,
		560, 562, 7, 9, 0, 0, 561, 560, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563,
		561, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565, 566,
		6, 70, 0, 0, 566, 142, 1, 0, 0, 0, 567, 568, 5, 47, 0, 0, 568, 569, 5,
		47, 0, 0, 569, 573, 1, 0, 0, 0, 570, 572, 8, 10, 0, 0, 571, 570, 1, 0,
		0, 0, 572, 575, 1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0,
		574, 576, 1, 0, 0, 0, 575, 573, 1, 0, 0, 0, 576, 577, 6, 71, 1, 0, 577,
		144, 1, 0, 0, 0, 578, 579, 5, 47, 0, 0, 579, 580, 5, 42, 0, 0, 580, 584,
		1, 0, 0, 0, 581, 583, 9, 0, 0, 0, 582, 581, 1, 0, 0, 0, 583, 586, 1, 0,
		0, 0, 584, 585, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 585, 587, 1, 0, 0, 0,
		586, 584, 1, 0, 0, 0, 587, 588, 5, 42, 0, 0, 588, 589, 5, 47, 0, 0, 589,
		590, 1, 0, 0, 0, 590, 591, 6, 72, 1, 0, 591, 146, 1, 0, 0, 0, 22, 0, 436,
		444, 448, 454, 478, 483, 486, 494, 496, 504, 509, 516, 521, 527, 534, 542,
		547, 549, 563, 573, 584, 2, 6, 0, 0, 0, 1, 0,
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

// Protobuf3LexerInit initializes any static state used to implement Protobuf3Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewProtobuf3Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Protobuf3LexerInit() {
	staticData := &Protobuf3LexerLexerStaticData
	staticData.once.Do(protobuf3lexerLexerInit)
}

// NewProtobuf3Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewProtobuf3Lexer(input antlr.CharStream) *Protobuf3Lexer {
	Protobuf3LexerInit()
	l := new(Protobuf3Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Protobuf3LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Protobuf3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Protobuf3Lexer tokens.
const (
	Protobuf3LexerSYNTAX            = 1
	Protobuf3LexerIMPORT            = 2
	Protobuf3LexerWEAK              = 3
	Protobuf3LexerPUBLIC            = 4
	Protobuf3LexerPACKAGE           = 5
	Protobuf3LexerOPTION            = 6
	Protobuf3LexerOPTIONAL          = 7
	Protobuf3LexerREPEATED          = 8
	Protobuf3LexerONEOF             = 9
	Protobuf3LexerMAP               = 10
	Protobuf3LexerINT32             = 11
	Protobuf3LexerINT64             = 12
	Protobuf3LexerUINT32            = 13
	Protobuf3LexerUINT64            = 14
	Protobuf3LexerSINT32            = 15
	Protobuf3LexerSINT64            = 16
	Protobuf3LexerFIXED32           = 17
	Protobuf3LexerFIXED64           = 18
	Protobuf3LexerSFIXED32          = 19
	Protobuf3LexerSFIXED64          = 20
	Protobuf3LexerBOOL              = 21
	Protobuf3LexerSTRING            = 22
	Protobuf3LexerDOUBLE            = 23
	Protobuf3LexerFLOAT             = 24
	Protobuf3LexerBYTES             = 25
	Protobuf3LexerRESERVED          = 26
	Protobuf3LexerTO                = 27
	Protobuf3LexerMAX               = 28
	Protobuf3LexerENUM              = 29
	Protobuf3LexerMESSAGE           = 30
	Protobuf3LexerSERVICE           = 31
	Protobuf3LexerEXTEND            = 32
	Protobuf3LexerRPC               = 33
	Protobuf3LexerSTREAM            = 34
	Protobuf3LexerRETURNS           = 35
	Protobuf3LexerPROTO3_LIT_SINGLE = 36
	Protobuf3LexerPROTO3_LIT_DOBULE = 37
	Protobuf3LexerSEMI              = 38
	Protobuf3LexerEQ                = 39
	Protobuf3LexerLP                = 40
	Protobuf3LexerRP                = 41
	Protobuf3LexerLB                = 42
	Protobuf3LexerRB                = 43
	Protobuf3LexerLC                = 44
	Protobuf3LexerRC                = 45
	Protobuf3LexerLT                = 46
	Protobuf3LexerGT                = 47
	Protobuf3LexerDOT               = 48
	Protobuf3LexerCOMMA             = 49
	Protobuf3LexerCOLON             = 50
	Protobuf3LexerPLUS              = 51
	Protobuf3LexerMINUS             = 52
	Protobuf3LexerSTR_LIT           = 53
	Protobuf3LexerBOOL_LIT          = 54
	Protobuf3LexerFLOAT_LIT         = 55
	Protobuf3LexerINT_LIT           = 56
	Protobuf3LexerIDENTIFIER        = 57
	Protobuf3LexerWS                = 58
	Protobuf3LexerLINE_COMMENT      = 59
	Protobuf3LexerCOMMENT           = 60
)
