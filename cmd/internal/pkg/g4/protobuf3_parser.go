// Code generated from Protobuf3.g4 by ANTLR 4.13.1. DO NOT EDIT.

package g4 // Protobuf3
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

type Protobuf3Parser struct {
	*antlr.BaseParser
}

var Protobuf3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protobuf3ParserInit() {
	staticData := &Protobuf3ParserStaticData
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
		"proto", "syntax", "importStatement", "packageStatement", "optionStatement",
		"optionName", "fieldLabel", "field", "fieldOptions", "fieldOption",
		"fieldNumber", "oneof", "oneofField", "mapField", "keyType", "type_",
		"reserved", "ranges", "range_", "reservedFieldNames", "topLevelDef",
		"enumDef", "enumBody", "enumElement", "enumField", "enumValueOptions",
		"enumValueOption", "messageDef", "messageBody", "messageElement", "extendDef",
		"serviceDef", "serviceElement", "rpc", "constant", "blockLit", "emptyStatement_",
		"ident", "fullIdent", "messageName", "enumName", "fieldName", "oneofName",
		"mapName", "serviceName", "rpcName", "messageType", "enumType", "intLit",
		"strLit", "boolLit", "floatLit", "keywords",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 496, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 113, 8, 0, 10, 0, 12, 0,
		116, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2,
		127, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 148, 8, 5,
		3, 5, 150, 8, 5, 1, 6, 1, 6, 1, 7, 3, 7, 155, 8, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 165, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 5, 8, 172, 8, 8, 10, 8, 12, 8, 175, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 189, 8, 11,
		10, 11, 12, 11, 192, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 204, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 221, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 244, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16,
		249, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 256, 8, 17, 10, 17,
		12, 17, 259, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 265, 8, 18, 3, 18,
		267, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 272, 8, 19, 10, 19, 12, 19, 275,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 281, 8, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 5, 22, 289, 8, 22, 10, 22, 12, 22, 292, 9, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 299, 8, 23, 1, 24, 1, 24, 1,
		24, 3, 24, 304, 8, 24, 1, 24, 1, 24, 3, 24, 308, 8, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 5, 25, 316, 8, 25, 10, 25, 12, 25, 319, 9, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 5, 28, 333, 8, 28, 10, 28, 12, 28, 336, 9, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 349,
		8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 356, 8, 30, 10, 30, 12,
		30, 359, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 367, 8,
		31, 10, 31, 12, 31, 370, 9, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 3, 32,
		377, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 383, 8, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 3, 33, 390, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 5, 33, 397, 8, 33, 10, 33, 12, 33, 400, 9, 33, 1, 33, 1, 33, 3, 33,
		404, 8, 33, 1, 34, 1, 34, 3, 34, 408, 8, 34, 1, 34, 1, 34, 3, 34, 412,
		8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 418, 8, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 5, 35, 425, 8, 35, 10, 35, 12, 35, 428, 9, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37, 436, 8, 37, 1, 38, 1, 38, 1,
		38, 5, 38, 441, 8, 38, 10, 38, 12, 38, 444, 9, 38, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 46, 3, 46, 461, 8, 46, 1, 46, 1, 46, 1, 46, 5, 46, 466, 8, 46, 10,
		46, 12, 46, 469, 9, 46, 1, 46, 1, 46, 1, 47, 3, 47, 474, 8, 47, 1, 47,
		1, 47, 1, 47, 5, 47, 479, 8, 47, 10, 47, 12, 47, 482, 9, 47, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 52, 0, 0, 53, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100,
		102, 104, 0, 7, 1, 0, 36, 37, 1, 0, 3, 4, 1, 0, 7, 8, 1, 0, 11, 22, 1,
		0, 51, 52, 2, 0, 36, 37, 53, 53, 2, 0, 1, 35, 54, 54, 521, 0, 106, 1, 0,
		0, 0, 2, 119, 1, 0, 0, 0, 4, 124, 1, 0, 0, 0, 6, 131, 1, 0, 0, 0, 8, 135,
		1, 0, 0, 0, 10, 149, 1, 0, 0, 0, 12, 151, 1, 0, 0, 0, 14, 154, 1, 0, 0,
		0, 16, 168, 1, 0, 0, 0, 18, 176, 1, 0, 0, 0, 20, 180, 1, 0, 0, 0, 22, 182,
		1, 0, 0, 0, 24, 195, 1, 0, 0, 0, 26, 207, 1, 0, 0, 0, 28, 224, 1, 0, 0,
		0, 30, 243, 1, 0, 0, 0, 32, 245, 1, 0, 0, 0, 34, 252, 1, 0, 0, 0, 36, 260,
		1, 0, 0, 0, 38, 268, 1, 0, 0, 0, 40, 280, 1, 0, 0, 0, 42, 282, 1, 0, 0,
		0, 44, 286, 1, 0, 0, 0, 46, 298, 1, 0, 0, 0, 48, 300, 1, 0, 0, 0, 50, 311,
		1, 0, 0, 0, 52, 322, 1, 0, 0, 0, 54, 326, 1, 0, 0, 0, 56, 330, 1, 0, 0,
		0, 58, 348, 1, 0, 0, 0, 60, 350, 1, 0, 0, 0, 62, 362, 1, 0, 0, 0, 64, 376,
		1, 0, 0, 0, 66, 378, 1, 0, 0, 0, 68, 417, 1, 0, 0, 0, 70, 419, 1, 0, 0,
		0, 72, 431, 1, 0, 0, 0, 74, 435, 1, 0, 0, 0, 76, 437, 1, 0, 0, 0, 78, 445,
		1, 0, 0, 0, 80, 447, 1, 0, 0, 0, 82, 449, 1, 0, 0, 0, 84, 451, 1, 0, 0,
		0, 86, 453, 1, 0, 0, 0, 88, 455, 1, 0, 0, 0, 90, 457, 1, 0, 0, 0, 92, 460,
		1, 0, 0, 0, 94, 473, 1, 0, 0, 0, 96, 485, 1, 0, 0, 0, 98, 487, 1, 0, 0,
		0, 100, 489, 1, 0, 0, 0, 102, 491, 1, 0, 0, 0, 104, 493, 1, 0, 0, 0, 106,
		114, 3, 2, 1, 0, 107, 113, 3, 4, 2, 0, 108, 113, 3, 6, 3, 0, 109, 113,
		3, 8, 4, 0, 110, 113, 3, 40, 20, 0, 111, 113, 3, 72, 36, 0, 112, 107, 1,
		0, 0, 0, 112, 108, 1, 0, 0, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0,
		0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118,
		5, 0, 0, 1, 118, 1, 1, 0, 0, 0, 119, 120, 5, 1, 0, 0, 120, 121, 5, 39,
		0, 0, 121, 122, 7, 0, 0, 0, 122, 123, 5, 38, 0, 0, 123, 3, 1, 0, 0, 0,
		124, 126, 5, 2, 0, 0, 125, 127, 7, 1, 0, 0, 126, 125, 1, 0, 0, 0, 126,
		127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 3, 98, 49, 0, 129, 130,
		5, 38, 0, 0, 130, 5, 1, 0, 0, 0, 131, 132, 5, 5, 0, 0, 132, 133, 3, 76,
		38, 0, 133, 134, 5, 38, 0, 0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 6, 0, 0,
		136, 137, 3, 10, 5, 0, 137, 138, 5, 39, 0, 0, 138, 139, 3, 68, 34, 0, 139,
		140, 5, 38, 0, 0, 140, 9, 1, 0, 0, 0, 141, 150, 3, 76, 38, 0, 142, 143,
		5, 40, 0, 0, 143, 144, 3, 76, 38, 0, 144, 147, 5, 41, 0, 0, 145, 146, 5,
		48, 0, 0, 146, 148, 3, 76, 38, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 150, 1, 0, 0, 0, 149, 141, 1, 0, 0, 0, 149, 142, 1, 0, 0, 0,
		150, 11, 1, 0, 0, 0, 151, 152, 7, 2, 0, 0, 152, 13, 1, 0, 0, 0, 153, 155,
		3, 12, 6, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 157, 3, 30, 15, 0, 157, 158, 3, 82, 41, 0, 158, 159, 5, 39,
		0, 0, 159, 164, 3, 20, 10, 0, 160, 161, 5, 42, 0, 0, 161, 162, 3, 16, 8,
		0, 162, 163, 5, 43, 0, 0, 163, 165, 1, 0, 0, 0, 164, 160, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 38, 0, 0, 167, 15,
		1, 0, 0, 0, 168, 173, 3, 18, 9, 0, 169, 170, 5, 49, 0, 0, 170, 172, 3,
		18, 9, 0, 171, 169, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0,
		0, 173, 174, 1, 0, 0, 0, 174, 17, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176,
		177, 3, 10, 5, 0, 177, 178, 5, 39, 0, 0, 178, 179, 3, 68, 34, 0, 179, 19,
		1, 0, 0, 0, 180, 181, 3, 96, 48, 0, 181, 21, 1, 0, 0, 0, 182, 183, 5, 9,
		0, 0, 183, 184, 3, 84, 42, 0, 184, 190, 5, 44, 0, 0, 185, 189, 3, 8, 4,
		0, 186, 189, 3, 24, 12, 0, 187, 189, 3, 72, 36, 0, 188, 185, 1, 0, 0, 0,
		188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 193, 194, 5, 45, 0, 0, 194, 23, 1, 0, 0, 0, 195, 196, 3, 30,
		15, 0, 196, 197, 3, 82, 41, 0, 197, 198, 5, 39, 0, 0, 198, 203, 3, 20,
		10, 0, 199, 200, 5, 42, 0, 0, 200, 201, 3, 16, 8, 0, 201, 202, 5, 43, 0,
		0, 202, 204, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		205, 1, 0, 0, 0, 205, 206, 5, 38, 0, 0, 206, 25, 1, 0, 0, 0, 207, 208,
		5, 10, 0, 0, 208, 209, 5, 46, 0, 0, 209, 210, 3, 28, 14, 0, 210, 211, 5,
		49, 0, 0, 211, 212, 3, 30, 15, 0, 212, 213, 5, 47, 0, 0, 213, 214, 3, 86,
		43, 0, 214, 215, 5, 39, 0, 0, 215, 220, 3, 20, 10, 0, 216, 217, 5, 42,
		0, 0, 217, 218, 3, 16, 8, 0, 218, 219, 5, 43, 0, 0, 219, 221, 1, 0, 0,
		0, 220, 216, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		223, 5, 38, 0, 0, 223, 27, 1, 0, 0, 0, 224, 225, 7, 3, 0, 0, 225, 29, 1,
		0, 0, 0, 226, 244, 5, 23, 0, 0, 227, 244, 5, 24, 0, 0, 228, 244, 5, 11,
		0, 0, 229, 244, 5, 12, 0, 0, 230, 244, 5, 13, 0, 0, 231, 244, 5, 14, 0,
		0, 232, 244, 5, 15, 0, 0, 233, 244, 5, 16, 0, 0, 234, 244, 5, 17, 0, 0,
		235, 244, 5, 18, 0, 0, 236, 244, 5, 19, 0, 0, 237, 244, 5, 20, 0, 0, 238,
		244, 5, 21, 0, 0, 239, 244, 5, 22, 0, 0, 240, 244, 5, 25, 0, 0, 241, 244,
		3, 92, 46, 0, 242, 244, 3, 94, 47, 0, 243, 226, 1, 0, 0, 0, 243, 227, 1,
		0, 0, 0, 243, 228, 1, 0, 0, 0, 243, 229, 1, 0, 0, 0, 243, 230, 1, 0, 0,
		0, 243, 231, 1, 0, 0, 0, 243, 232, 1, 0, 0, 0, 243, 233, 1, 0, 0, 0, 243,
		234, 1, 0, 0, 0, 243, 235, 1, 0, 0, 0, 243, 236, 1, 0, 0, 0, 243, 237,
		1, 0, 0, 0, 243, 238, 1, 0, 0, 0, 243, 239, 1, 0, 0, 0, 243, 240, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 31, 1, 0, 0, 0,
		245, 248, 5, 26, 0, 0, 246, 249, 3, 34, 17, 0, 247, 249, 3, 38, 19, 0,
		248, 246, 1, 0, 0, 0, 248, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		251, 5, 38, 0, 0, 251, 33, 1, 0, 0, 0, 252, 257, 3, 36, 18, 0, 253, 254,
		5, 49, 0, 0, 254, 256, 3, 36, 18, 0, 255, 253, 1, 0, 0, 0, 256, 259, 1,
		0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 35, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 260, 266, 3, 96, 48, 0, 261, 264, 5, 27, 0, 0,
		262, 265, 3, 96, 48, 0, 263, 265, 5, 28, 0, 0, 264, 262, 1, 0, 0, 0, 264,
		263, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 261, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 37, 1, 0, 0, 0, 268, 273, 3, 98, 49, 0, 269, 270, 5, 49,
		0, 0, 270, 272, 3, 98, 49, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0,
		0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 39, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 276, 281, 3, 54, 27, 0, 277, 281, 3, 42, 21, 0, 278, 281,
		3, 60, 30, 0, 279, 281, 3, 62, 31, 0, 280, 276, 1, 0, 0, 0, 280, 277, 1,
		0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 41, 1, 0, 0,
		0, 282, 283, 5, 29, 0, 0, 283, 284, 3, 80, 40, 0, 284, 285, 3, 44, 22,
		0, 285, 43, 1, 0, 0, 0, 286, 290, 5, 44, 0, 0, 287, 289, 3, 46, 23, 0,
		288, 287, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290,
		291, 1, 0, 0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294,
		5, 45, 0, 0, 294, 45, 1, 0, 0, 0, 295, 299, 3, 8, 4, 0, 296, 299, 3, 48,
		24, 0, 297, 299, 3, 72, 36, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0,
		0, 298, 297, 1, 0, 0, 0, 299, 47, 1, 0, 0, 0, 300, 301, 3, 74, 37, 0, 301,
		303, 5, 39, 0, 0, 302, 304, 5, 52, 0, 0, 303, 302, 1, 0, 0, 0, 303, 304,
		1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 3, 96, 48, 0, 306, 308, 3,
		50, 25, 0, 307, 306, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0,
		0, 0, 309, 310, 5, 38, 0, 0, 310, 49, 1, 0, 0, 0, 311, 312, 5, 42, 0, 0,
		312, 317, 3, 52, 26, 0, 313, 314, 5, 49, 0, 0, 314, 316, 3, 52, 26, 0,
		315, 313, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321,
		5, 43, 0, 0, 321, 51, 1, 0, 0, 0, 322, 323, 3, 10, 5, 0, 323, 324, 5, 39,
		0, 0, 324, 325, 3, 68, 34, 0, 325, 53, 1, 0, 0, 0, 326, 327, 5, 30, 0,
		0, 327, 328, 3, 78, 39, 0, 328, 329, 3, 56, 28, 0, 329, 55, 1, 0, 0, 0,
		330, 334, 5, 44, 0, 0, 331, 333, 3, 58, 29, 0, 332, 331, 1, 0, 0, 0, 333,
		336, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337,
		1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 338, 5, 45, 0, 0, 338, 57, 1, 0,
		0, 0, 339, 349, 3, 14, 7, 0, 340, 349, 3, 42, 21, 0, 341, 349, 3, 54, 27,
		0, 342, 349, 3, 60, 30, 0, 343, 349, 3, 8, 4, 0, 344, 349, 3, 22, 11, 0,
		345, 349, 3, 26, 13, 0, 346, 349, 3, 32, 16, 0, 347, 349, 3, 72, 36, 0,
		348, 339, 1, 0, 0, 0, 348, 340, 1, 0, 0, 0, 348, 341, 1, 0, 0, 0, 348,
		342, 1, 0, 0, 0, 348, 343, 1, 0, 0, 0, 348, 344, 1, 0, 0, 0, 348, 345,
		1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 347, 1, 0, 0, 0, 349, 59, 1, 0,
		0, 0, 350, 351, 5, 32, 0, 0, 351, 352, 3, 92, 46, 0, 352, 357, 5, 44, 0,
		0, 353, 356, 3, 14, 7, 0, 354, 356, 3, 72, 36, 0, 355, 353, 1, 0, 0, 0,
		355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357,
		358, 1, 0, 0, 0, 358, 360, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 361,
		5, 45, 0, 0, 361, 61, 1, 0, 0, 0, 362, 363, 5, 31, 0, 0, 363, 364, 3, 88,
		44, 0, 364, 368, 5, 44, 0, 0, 365, 367, 3, 64, 32, 0, 366, 365, 1, 0, 0,
		0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369,
		371, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 372, 5, 45, 0, 0, 372, 63,
		1, 0, 0, 0, 373, 377, 3, 8, 4, 0, 374, 377, 3, 66, 33, 0, 375, 377, 3,
		72, 36, 0, 376, 373, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0,
		0, 0, 377, 65, 1, 0, 0, 0, 378, 379, 5, 33, 0, 0, 379, 380, 3, 90, 45,
		0, 380, 382, 5, 40, 0, 0, 381, 383, 5, 34, 0, 0, 382, 381, 1, 0, 0, 0,
		382, 383, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 385, 3, 92, 46, 0, 385,
		386, 5, 41, 0, 0, 386, 387, 5, 35, 0, 0, 387, 389, 5, 40, 0, 0, 388, 390,
		5, 34, 0, 0, 389, 388, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 1, 0,
		0, 0, 391, 392, 3, 92, 46, 0, 392, 403, 5, 41, 0, 0, 393, 398, 5, 44, 0,
		0, 394, 397, 3, 8, 4, 0, 395, 397, 3, 72, 36, 0, 396, 394, 1, 0, 0, 0,
		396, 395, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398,
		399, 1, 0, 0, 0, 399, 401, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 404,
		5, 45, 0, 0, 402, 404, 5, 38, 0, 0, 403, 393, 1, 0, 0, 0, 403, 402, 1,
		0, 0, 0, 404, 67, 1, 0, 0, 0, 405, 418, 3, 76, 38, 0, 406, 408, 7, 4, 0,
		0, 407, 406, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409,
		418, 3, 96, 48, 0, 410, 412, 7, 4, 0, 0, 411, 410, 1, 0, 0, 0, 411, 412,
		1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 418, 3, 102, 51, 0, 414, 418, 3,
		98, 49, 0, 415, 418, 3, 100, 50, 0, 416, 418, 3, 70, 35, 0, 417, 405, 1,
		0, 0, 0, 417, 407, 1, 0, 0, 0, 417, 411, 1, 0, 0, 0, 417, 414, 1, 0, 0,
		0, 417, 415, 1, 0, 0, 0, 417, 416, 1, 0, 0, 0, 418, 69, 1, 0, 0, 0, 419,
		426, 5, 44, 0, 0, 420, 421, 3, 74, 37, 0, 421, 422, 5, 50, 0, 0, 422, 423,
		3, 68, 34, 0, 423, 425, 1, 0, 0, 0, 424, 420, 1, 0, 0, 0, 425, 428, 1,
		0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 429, 1, 0, 0,
		0, 428, 426, 1, 0, 0, 0, 429, 430, 5, 45, 0, 0, 430, 71, 1, 0, 0, 0, 431,
		432, 5, 38, 0, 0, 432, 73, 1, 0, 0, 0, 433, 436, 5, 57, 0, 0, 434, 436,
		3, 104, 52, 0, 435, 433, 1, 0, 0, 0, 435, 434, 1, 0, 0, 0, 436, 75, 1,
		0, 0, 0, 437, 442, 3, 74, 37, 0, 438, 439, 5, 48, 0, 0, 439, 441, 3, 74,
		37, 0, 440, 438, 1, 0, 0, 0, 441, 444, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0,
		442, 443, 1, 0, 0, 0, 443, 77, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 446,
		3, 74, 37, 0, 446, 79, 1, 0, 0, 0, 447, 448, 3, 74, 37, 0, 448, 81, 1,
		0, 0, 0, 449, 450, 3, 74, 37, 0, 450, 83, 1, 0, 0, 0, 451, 452, 3, 74,
		37, 0, 452, 85, 1, 0, 0, 0, 453, 454, 3, 74, 37, 0, 454, 87, 1, 0, 0, 0,
		455, 456, 3, 74, 37, 0, 456, 89, 1, 0, 0, 0, 457, 458, 3, 74, 37, 0, 458,
		91, 1, 0, 0, 0, 459, 461, 5, 48, 0, 0, 460, 459, 1, 0, 0, 0, 460, 461,
		1, 0, 0, 0, 461, 467, 1, 0, 0, 0, 462, 463, 3, 74, 37, 0, 463, 464, 5,
		48, 0, 0, 464, 466, 1, 0, 0, 0, 465, 462, 1, 0, 0, 0, 466, 469, 1, 0, 0,
		0, 467, 465, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 470, 1, 0, 0, 0, 469,
		467, 1, 0, 0, 0, 470, 471, 3, 78, 39, 0, 471, 93, 1, 0, 0, 0, 472, 474,
		5, 48, 0, 0, 473, 472, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 480, 1, 0,
		0, 0, 475, 476, 3, 74, 37, 0, 476, 477, 5, 48, 0, 0, 477, 479, 1, 0, 0,
		0, 478, 475, 1, 0, 0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480,
		481, 1, 0, 0, 0, 481, 483, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 483, 484,
		3, 80, 40, 0, 484, 95, 1, 0, 0, 0, 485, 486, 5, 56, 0, 0, 486, 97, 1, 0,
		0, 0, 487, 488, 7, 5, 0, 0, 488, 99, 1, 0, 0, 0, 489, 490, 5, 54, 0, 0,
		490, 101, 1, 0, 0, 0, 491, 492, 5, 55, 0, 0, 492, 103, 1, 0, 0, 0, 493,
		494, 7, 6, 0, 0, 494, 105, 1, 0, 0, 0, 45, 112, 114, 126, 147, 149, 154,
		164, 173, 188, 190, 203, 220, 243, 248, 257, 264, 266, 273, 280, 290, 298,
		303, 307, 317, 334, 348, 355, 357, 368, 376, 382, 389, 396, 398, 403, 407,
		411, 417, 426, 435, 442, 460, 467, 473, 480,
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

// Protobuf3ParserInit initializes any static state used to implement Protobuf3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtobuf3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Protobuf3ParserInit() {
	staticData := &Protobuf3ParserStaticData
	staticData.once.Do(protobuf3ParserInit)
}

// NewProtobuf3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtobuf3Parser(input antlr.TokenStream) *Protobuf3Parser {
	Protobuf3ParserInit()
	this := new(Protobuf3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Protobuf3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Protobuf3.g4"

	return this
}

// Protobuf3Parser tokens.
const (
	Protobuf3ParserEOF               = antlr.TokenEOF
	Protobuf3ParserSYNTAX            = 1
	Protobuf3ParserIMPORT            = 2
	Protobuf3ParserWEAK              = 3
	Protobuf3ParserPUBLIC            = 4
	Protobuf3ParserPACKAGE           = 5
	Protobuf3ParserOPTION            = 6
	Protobuf3ParserOPTIONAL          = 7
	Protobuf3ParserREPEATED          = 8
	Protobuf3ParserONEOF             = 9
	Protobuf3ParserMAP               = 10
	Protobuf3ParserINT32             = 11
	Protobuf3ParserINT64             = 12
	Protobuf3ParserUINT32            = 13
	Protobuf3ParserUINT64            = 14
	Protobuf3ParserSINT32            = 15
	Protobuf3ParserSINT64            = 16
	Protobuf3ParserFIXED32           = 17
	Protobuf3ParserFIXED64           = 18
	Protobuf3ParserSFIXED32          = 19
	Protobuf3ParserSFIXED64          = 20
	Protobuf3ParserBOOL              = 21
	Protobuf3ParserSTRING            = 22
	Protobuf3ParserDOUBLE            = 23
	Protobuf3ParserFLOAT             = 24
	Protobuf3ParserBYTES             = 25
	Protobuf3ParserRESERVED          = 26
	Protobuf3ParserTO                = 27
	Protobuf3ParserMAX               = 28
	Protobuf3ParserENUM              = 29
	Protobuf3ParserMESSAGE           = 30
	Protobuf3ParserSERVICE           = 31
	Protobuf3ParserEXTEND            = 32
	Protobuf3ParserRPC               = 33
	Protobuf3ParserSTREAM            = 34
	Protobuf3ParserRETURNS           = 35
	Protobuf3ParserPROTO3_LIT_SINGLE = 36
	Protobuf3ParserPROTO3_LIT_DOBULE = 37
	Protobuf3ParserSEMI              = 38
	Protobuf3ParserEQ                = 39
	Protobuf3ParserLP                = 40
	Protobuf3ParserRP                = 41
	Protobuf3ParserLB                = 42
	Protobuf3ParserRB                = 43
	Protobuf3ParserLC                = 44
	Protobuf3ParserRC                = 45
	Protobuf3ParserLT                = 46
	Protobuf3ParserGT                = 47
	Protobuf3ParserDOT               = 48
	Protobuf3ParserCOMMA             = 49
	Protobuf3ParserCOLON             = 50
	Protobuf3ParserPLUS              = 51
	Protobuf3ParserMINUS             = 52
	Protobuf3ParserSTR_LIT           = 53
	Protobuf3ParserBOOL_LIT          = 54
	Protobuf3ParserFLOAT_LIT         = 55
	Protobuf3ParserINT_LIT           = 56
	Protobuf3ParserIDENTIFIER        = 57
	Protobuf3ParserWS                = 58
	Protobuf3ParserLINE_COMMENT      = 59
	Protobuf3ParserCOMMENT           = 60
)

// Protobuf3Parser rules.
const (
	Protobuf3ParserRULE_proto              = 0
	Protobuf3ParserRULE_syntax             = 1
	Protobuf3ParserRULE_importStatement    = 2
	Protobuf3ParserRULE_packageStatement   = 3
	Protobuf3ParserRULE_optionStatement    = 4
	Protobuf3ParserRULE_optionName         = 5
	Protobuf3ParserRULE_fieldLabel         = 6
	Protobuf3ParserRULE_field              = 7
	Protobuf3ParserRULE_fieldOptions       = 8
	Protobuf3ParserRULE_fieldOption        = 9
	Protobuf3ParserRULE_fieldNumber        = 10
	Protobuf3ParserRULE_oneof              = 11
	Protobuf3ParserRULE_oneofField         = 12
	Protobuf3ParserRULE_mapField           = 13
	Protobuf3ParserRULE_keyType            = 14
	Protobuf3ParserRULE_type_              = 15
	Protobuf3ParserRULE_reserved           = 16
	Protobuf3ParserRULE_ranges             = 17
	Protobuf3ParserRULE_range_             = 18
	Protobuf3ParserRULE_reservedFieldNames = 19
	Protobuf3ParserRULE_topLevelDef        = 20
	Protobuf3ParserRULE_enumDef            = 21
	Protobuf3ParserRULE_enumBody           = 22
	Protobuf3ParserRULE_enumElement        = 23
	Protobuf3ParserRULE_enumField          = 24
	Protobuf3ParserRULE_enumValueOptions   = 25
	Protobuf3ParserRULE_enumValueOption    = 26
	Protobuf3ParserRULE_messageDef         = 27
	Protobuf3ParserRULE_messageBody        = 28
	Protobuf3ParserRULE_messageElement     = 29
	Protobuf3ParserRULE_extendDef          = 30
	Protobuf3ParserRULE_serviceDef         = 31
	Protobuf3ParserRULE_serviceElement     = 32
	Protobuf3ParserRULE_rpc                = 33
	Protobuf3ParserRULE_constant           = 34
	Protobuf3ParserRULE_blockLit           = 35
	Protobuf3ParserRULE_emptyStatement_    = 36
	Protobuf3ParserRULE_ident              = 37
	Protobuf3ParserRULE_fullIdent          = 38
	Protobuf3ParserRULE_messageName        = 39
	Protobuf3ParserRULE_enumName           = 40
	Protobuf3ParserRULE_fieldName          = 41
	Protobuf3ParserRULE_oneofName          = 42
	Protobuf3ParserRULE_mapName            = 43
	Protobuf3ParserRULE_serviceName        = 44
	Protobuf3ParserRULE_rpcName            = 45
	Protobuf3ParserRULE_messageType        = 46
	Protobuf3ParserRULE_enumType           = 47
	Protobuf3ParserRULE_intLit             = 48
	Protobuf3ParserRULE_strLit             = 49
	Protobuf3ParserRULE_boolLit            = 50
	Protobuf3ParserRULE_floatLit           = 51
	Protobuf3ParserRULE_keywords           = 52
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Syntax() ISyntaxContext
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllPackageStatement() []IPackageStatementContext
	PackageStatement(i int) IPackageStatementContext
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllTopLevelDef() []ITopLevelDefContext
	TopLevelDef(i int) ITopLevelDefContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
	return p
}

func InitEmptyProtoContext(p *ProtoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEOF, 0)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
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

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageStatementContext); ok {
			len++
		}
	}

	tst := make([]IPackageStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageStatementContext); ok {
			tst[i] = t.(IPackageStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
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

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDefContext); ok {
			tst[i] = t.(ITopLevelDefContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
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

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitProto(s)
	}
}

func (s *ProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf3ParserRULE_proto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Syntax()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&282930970724) != 0 {
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserIMPORT:
			{
				p.SetState(107)
				p.ImportStatement()
			}

		case Protobuf3ParserPACKAGE:
			{
				p.SetState(108)
				p.PackageStatement()
			}

		case Protobuf3ParserOPTION:
			{
				p.SetState(109)
				p.OptionStatement()
			}

		case Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserEXTEND:
			{
				p.SetState(110)
				p.TopLevelDef()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(111)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(Protobuf3ParserEOF)
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

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNTAX() antlr.TerminalNode
	EQ() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	PROTO3_LIT_SINGLE() antlr.TerminalNode
	PROTO3_LIT_DOBULE() antlr.TerminalNode

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
	return p
}

func InitEmptySyntaxContext(p *SyntaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *SyntaxContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *SyntaxContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (s *SyntaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitSyntax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf3ParserRULE_syntax)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(Protobuf3ParserSYNTAX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserPROTO3_LIT_SINGLE || _la == Protobuf3ParserPROTO3_LIT_DOBULE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(122)
		p.Match(Protobuf3ParserSEMI)
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

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	StrLit() IStrLitContext
	SEMI() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf3ParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(Protobuf3ParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC {
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(128)
		p.StrLit()
	}
	{
		p.SetState(129)
		p.Match(Protobuf3ParserSEMI)
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

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	FullIdent() IFullIdentContext
	SEMI() antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf3ParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(Protobuf3ParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.FullIdent()
	}
	{
		p.SetState(133)
		p.Match(Protobuf3ParserSEMI)
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

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext
	SEMI() antlr.TerminalNode

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *OptionStatementContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionStatement(s)
	}
}

func (s *OptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionStatement(s)
	}
}

func (s *OptionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf3ParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(Protobuf3ParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.OptionName()
	}
	{
		p.SetState(137)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Constant()
	}
	{
		p.SetState(139)
		p.Match(Protobuf3ParserSEMI)
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

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFullIdent() []IFullIdentContext
	FullIdent(i int) IFullIdentContext
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllFullIdent() []IFullIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFullIdentContext); ok {
			len++
		}
	}

	tst := make([]IFullIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFullIdentContext); ok {
			tst[i] = t.(IFullIdentContext)
			i++
		}
	}

	return tst
}

func (s *OptionNameContext) FullIdent(i int) IFullIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
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

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf3ParserRULE_optionName)
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserOPTIONAL, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserEXTEND, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT, Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.FullIdent()
		}

	case Protobuf3ParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(Protobuf3ParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.FullIdent()
		}
		{
			p.SetState(144)
			p.Match(Protobuf3ParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserDOT {
			{
				p.SetState(145)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(146)
				p.FullIdent()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IFieldLabelContext is an interface to support dynamic dispatch.
type IFieldLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTIONAL() antlr.TerminalNode
	REPEATED() antlr.TerminalNode

	// IsFieldLabelContext differentiates from other interfaces.
	IsFieldLabelContext()
}

type FieldLabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldLabelContext() *FieldLabelContext {
	var p = new(FieldLabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldLabel
	return p
}

func InitEmptyFieldLabelContext(p *FieldLabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldLabel
}

func (*FieldLabelContext) IsFieldLabelContext() {}

func NewFieldLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldLabelContext {
	var p = new(FieldLabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldLabel

	return p
}

func (s *FieldLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldLabelContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTIONAL, 0)
}

func (s *FieldLabelContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *FieldLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldLabel(s)
	}
}

func (s *FieldLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldLabel(s)
	}
}

func (s *FieldLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldLabel() (localctx IFieldLabelContext) {
	localctx = NewFieldLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf3ParserRULE_fieldLabel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserOPTIONAL || _la == Protobuf3ParserREPEATED) {
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	FieldLabel() IFieldLabelContext
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *FieldContext) FieldLabel() IFieldLabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLabelContext)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf3ParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.FieldLabel()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(156)
		p.Type_()
	}
	{
		p.SetState(157)
		p.FieldName()
	}
	{
		p.SetState(158)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.FieldNumber()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(160)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.FieldOptions()
		}
		{
			p.SetState(162)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(166)
		p.Match(Protobuf3ParserSEMI)
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

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldOption() []IFieldOptionContext
	FieldOption(i int) IFieldOptionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
	return p
}

func InitEmptyFieldOptionsContext(p *FieldOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldOptionContext); ok {
			len++
		}
	}

	tst := make([]IFieldOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldOptionContext); ok {
			tst[i] = t.(IFieldOptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionContext); ok {
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

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf3ParserRULE_fieldOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.FieldOption()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(169)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.FieldOption()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
	return p
}

func InitEmptyFieldOptionContext(p *FieldOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (s *FieldOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf3ParserRULE_fieldOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.OptionName()
	}
	{
		p.SetState(177)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Constant()
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

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
	return p
}

func InitEmptyFieldNumberContext(p *FieldNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (s *FieldNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf3ParserRULE_fieldNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.IntLit()
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

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ONEOF() antlr.TerminalNode
	OneofName() IOneofNameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllOneofField() []IOneofFieldContext
	OneofField(i int) IOneofFieldContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
	return p
}

func InitEmptyOneofContext(p *OneofContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *OneofContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *OneofContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOneofFieldContext); ok {
			len++
		}
	}

	tst := make([]IOneofFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOneofFieldContext); ok {
			tst[i] = t.(IOneofFieldContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofFieldContext); ok {
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

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneof(s)
	}
}

func (s *OneofContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneof(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf3ParserRULE_oneof)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(Protobuf3ParserONEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.OneofName()
	}
	{
		p.SetState(184)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162411405159432190) != 0 {
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(185)
				p.OptionStatement()
			}

		case 2:
			{
				p.SetState(186)
				p.OneofField()
			}

		case 3:
			{
				p.SetState(187)
				p.EmptyStatement_()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(Protobuf3ParserRC)
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

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
	return p
}

func InitEmptyOneofFieldContext(p *OneofFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OneofFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (s *OneofFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneofField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf3ParserRULE_oneofField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Type_()
	}
	{
		p.SetState(196)
		p.FieldName()
	}
	{
		p.SetState(197)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.FieldNumber()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(199)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.FieldOptions()
		}
		{
			p.SetState(201)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(205)
		p.Match(Protobuf3ParserSEMI)
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

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LT() antlr.TerminalNode
	KeyType() IKeyTypeContext
	COMMA() antlr.TerminalNode
	Type_() IType_Context
	GT() antlr.TerminalNode
	MapName() IMapNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
	return p
}

func InitEmptyMapFieldContext(p *MapFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLT, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, 0)
}

func (s *MapFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapFieldContext) GT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserGT, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapField(s)
	}
}

func (s *MapFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf3ParserRULE_mapField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(Protobuf3ParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(Protobuf3ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.KeyType()
	}
	{
		p.SetState(210)
		p.Match(Protobuf3ParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Type_()
	}
	{
		p.SetState(212)
		p.Match(Protobuf3ParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.MapName()
	}
	{
		p.SetState(214)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.FieldNumber()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(216)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.FieldOptions()
		}
		{
			p.SetState(218)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(222)
		p.Match(Protobuf3ParserSEMI)
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

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (s *KeyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitKeyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf3ParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8386560) != 0) {
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

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MessageType() IMessageTypeContext
	EnumType() IEnumTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf3ParserRULE_type_)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Match(Protobuf3ParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(Protobuf3ParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			p.Match(Protobuf3ParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			p.Match(Protobuf3ParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.Match(Protobuf3ParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			p.Match(Protobuf3ParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(232)
			p.Match(Protobuf3ParserSINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(233)
			p.Match(Protobuf3ParserSINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(234)
			p.Match(Protobuf3ParserFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(235)
			p.Match(Protobuf3ParserFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(236)
			p.Match(Protobuf3ParserSFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(237)
			p.Match(Protobuf3ParserSFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(238)
			p.Match(Protobuf3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(239)
			p.Match(Protobuf3ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(240)
			p.Match(Protobuf3ParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(241)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(242)
			p.EnumType()
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

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESERVED() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	Ranges() IRangesContext
	ReservedFieldNames() IReservedFieldNamesContext

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reserved
	return p
}

func InitEmptyReservedContext(p *ReservedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reserved
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) ReservedFieldNames() IReservedFieldNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedFieldNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReserved(s)
	}
}

func (s *ReservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf3ParserRULE_reserved)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(Protobuf3ParserRESERVED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserINT_LIT:
		{
			p.SetState(246)
			p.Ranges()
		}

	case Protobuf3ParserPROTO3_LIT_SINGLE, Protobuf3ParserPROTO3_LIT_DOBULE, Protobuf3ParserSTR_LIT:
		{
			p.SetState(247)
			p.ReservedFieldNames()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(250)
		p.Match(Protobuf3ParserSEMI)
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

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRange_() []IRange_Context
	Range_(i int) IRange_Context
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ranges
	return p
}

func InitEmptyRangesContext(p *RangesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ranges
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRange_() []IRange_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRange_Context); ok {
			len++
		}
	}

	tst := make([]IRange_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRange_Context); ok {
			tst[i] = t.(IRange_Context)
			i++
		}
	}

	return tst
}

func (s *RangesContext) Range_(i int) IRange_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_Context); ok {
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

	return t.(IRange_Context)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRanges(s)
	}
}

func (s *RangesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRanges(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf3ParserRULE_ranges)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Range_()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(253)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Range_()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IRange_Context is an interface to support dynamic dispatch.
type IRange_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntLit() []IIntLitContext
	IntLit(i int) IIntLitContext
	TO() antlr.TerminalNode
	MAX() antlr.TerminalNode

	// IsRange_Context differentiates from other interfaces.
	IsRange_Context()
}

type Range_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_Context() *Range_Context {
	var p = new(Range_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_range_
	return p
}

func InitEmptyRange_Context(p *Range_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_range_
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_range_

	return p
}

func (s *Range_Context) GetParser() antlr.Parser { return s.parser }

func (s *Range_Context) AllIntLit() []IIntLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntLitContext); ok {
			len++
		}
	}

	tst := make([]IIntLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntLitContext); ok {
			tst[i] = t.(IIntLitContext)
			i++
		}
	}

	return tst
}

func (s *Range_Context) IntLit(i int) IIntLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
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

	return t.(IIntLitContext)
}

func (s *Range_Context) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *Range_Context) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRange_(s)
	}
}

func (s *Range_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRange_(s)
	}
}

func (s *Range_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRange_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Range_() (localctx IRange_Context) {
	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf3ParserRULE_range_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.IntLit()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserTO {
		{
			p.SetState(261)
			p.Match(Protobuf3ParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserINT_LIT:
			{
				p.SetState(262)
				p.IntLit()
			}

		case Protobuf3ParserMAX:
			{
				p.SetState(263)
				p.Match(Protobuf3ParserMAX)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IReservedFieldNamesContext is an interface to support dynamic dispatch.
type IReservedFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStrLit() []IStrLitContext
	StrLit(i int) IStrLitContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReservedFieldNamesContext differentiates from other interfaces.
	IsReservedFieldNamesContext()
}

type ReservedFieldNamesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedFieldNamesContext() *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames
	return p
}

func InitEmptyReservedFieldNamesContext(p *ReservedFieldNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames
}

func (*ReservedFieldNamesContext) IsReservedFieldNamesContext() {}

func NewReservedFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames

	return p
}

func (s *ReservedFieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedFieldNamesContext) AllStrLit() []IStrLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStrLitContext); ok {
			len++
		}
	}

	tst := make([]IStrLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStrLitContext); ok {
			tst[i] = t.(IStrLitContext)
			i++
		}
	}

	return tst
}

func (s *ReservedFieldNamesContext) StrLit(i int) IStrLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
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

	return t.(IStrLitContext)
}

func (s *ReservedFieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *ReservedFieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *ReservedFieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReservedFieldNames(s)
	}
}

func (s *ReservedFieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReservedFieldNames(s)
	}
}

func (s *ReservedFieldNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitReservedFieldNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ReservedFieldNames() (localctx IReservedFieldNamesContext) {
	localctx = NewReservedFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf3ParserRULE_reservedFieldNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.StrLit()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(269)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.StrLit()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageDef() IMessageDefContext
	EnumDef() IEnumDefContext
	ExtendDef() IExtendDefContext
	ServiceDef() IServiceDefContext

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
	return p
}

func InitEmptyTopLevelDefContext(p *TopLevelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) MessageDef() IMessageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) ExtendDef() IExtendDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendDefContext)
}

func (s *TopLevelDefContext) ServiceDef() IServiceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf3ParserRULE_topLevelDef)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.MessageDef()
		}

	case Protobuf3ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.EnumDef()
		}

	case Protobuf3ParserEXTEND:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.ExtendDef()
		}

	case Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(279)
			p.ServiceDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	EnumName() IEnumNameContext
	EnumBody() IEnumBodyContext

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf3ParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(Protobuf3ParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.EnumName()
	}
	{
		p.SetState(284)
		p.EnumBody()
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

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllEnumElement() []IEnumElementContext
	EnumElement(i int) IEnumElementContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumElementContext); ok {
			len++
		}
	}

	tst := make([]IEnumElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumElementContext); ok {
			tst[i] = t.(IEnumElementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumElementContext); ok {
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

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf3ParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162129930182721534) != 0 {
		{
			p.SetState(287)
			p.EnumElement()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(293)
		p.Match(Protobuf3ParserRC)
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

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionStatement() IOptionStatementContext
	EnumField() IEnumFieldContext
	EmptyStatement_() IEmptyStatement_Context

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumElement
	return p
}

func InitEmptyEnumElementContext(p *EnumElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumElement
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (s *EnumElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf3ParserRULE_enumElement)
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.OptionStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(297)
			p.EmptyStatement_()
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

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	EQ() antlr.TerminalNode
	IntLit() IIntLitContext
	SEMI() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EnumValueOptions() IEnumValueOptionsContext

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMINUS, 0)
}

func (s *EnumFieldContext) EnumValueOptions() IEnumValueOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionsContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf3ParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Ident()
	}
	{
		p.SetState(301)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserMINUS {
		{
			p.SetState(302)
			p.Match(Protobuf3ParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(305)
		p.IntLit()
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(306)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(309)
		p.Match(Protobuf3ParserSEMI)
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

// IEnumValueOptionsContext is an interface to support dynamic dispatch.
type IEnumValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LB() antlr.TerminalNode
	AllEnumValueOption() []IEnumValueOptionContext
	EnumValueOption(i int) IEnumValueOptionContext
	RB() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEnumValueOptionsContext differentiates from other interfaces.
	IsEnumValueOptionsContext()
}

type EnumValueOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionsContext() *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions
	return p
}

func InitEmptyEnumValueOptionsContext(p *EnumValueOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *EnumValueOptionsContext) AllEnumValueOption() []IEnumValueOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueOptionContext); ok {
			tst[i] = t.(IEnumValueOptionContext)
			i++
		}
	}

	return tst
}

func (s *EnumValueOptionsContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
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

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOptions(s)
	}
}

func (s *EnumValueOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOptions(s)
	}
}

func (s *EnumValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf3ParserRULE_enumValueOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(Protobuf3ParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.EnumValueOption()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(313)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.EnumValueOption()
		}

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(320)
		p.Match(Protobuf3ParserRB)
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

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
	return p
}

func InitEmptyEnumValueOptionContext(p *EnumValueOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf3ParserRULE_enumValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.OptionName()
	}
	{
		p.SetState(323)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Constant()
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

// IMessageDefContext is an interface to support dynamic dispatch.
type IMessageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MESSAGE() antlr.TerminalNode
	MessageName() IMessageNameContext
	MessageBody() IMessageBodyContext

	// IsMessageDefContext differentiates from other interfaces.
	IsMessageDefContext()
}

type MessageDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefContext() *MessageDefContext {
	var p = new(MessageDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageDef
	return p
}

func InitEmptyMessageDefContext(p *MessageDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageDef
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *MessageDefContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageDefContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageDef(s)
	}
}

func (s *MessageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageDef(s)
	}
}

func (s *MessageDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageDef() (localctx IMessageDefContext) {
	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf3ParserRULE_messageDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(Protobuf3ParserMESSAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.MessageName()
	}
	{
		p.SetState(328)
		p.MessageBody()
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

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllMessageElement() []IMessageElementContext
	MessageElement(i int) IMessageElementContext

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
	return p
}

func InitEmptyMessageBodyContext(p *MessageBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *MessageBodyContext) AllMessageElement() []IMessageElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageElementContext); ok {
			len++
		}
	}

	tst := make([]IMessageElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageElementContext); ok {
			tst[i] = t.(IMessageElementContext)
			i++
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageElement(i int) IMessageElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageElementContext); ok {
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

	return t.(IMessageElementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (s *MessageBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf3ParserRULE_messageBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162411405159432190) != 0 {
		{
			p.SetState(331)
			p.MessageElement()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(337)
		p.Match(Protobuf3ParserRC)
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

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	EnumDef() IEnumDefContext
	MessageDef() IMessageDefContext
	ExtendDef() IExtendDefContext
	OptionStatement() IOptionStatementContext
	Oneof() IOneofContext
	MapField() IMapFieldContext
	Reserved() IReservedContext
	EmptyStatement_() IEmptyStatement_Context

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageElement
	return p
}

func InitEmptyMessageElementContext(p *MessageElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageElement
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *MessageElementContext) MessageDef() IMessageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *MessageElementContext) ExtendDef() IExtendDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendDefContext)
}

func (s *MessageElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *MessageElementContext) Oneof() IOneofContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageElementContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageElementContext) Reserved() IReservedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageElement(s)
	}
}

func (s *MessageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageElement(s)
	}
}

func (s *MessageElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageElement() (localctx IMessageElementContext) {
	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf3ParserRULE_messageElement)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.EnumDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
			p.MessageDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(342)
			p.ExtendDef()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(343)
			p.OptionStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(344)
			p.Oneof()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(345)
			p.MapField()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(346)
			p.Reserved()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(347)
			p.EmptyStatement_()
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

// IExtendDefContext is an interface to support dynamic dispatch.
type IExtendDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTEND() antlr.TerminalNode
	MessageType() IMessageTypeContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsExtendDefContext differentiates from other interfaces.
	IsExtendDefContext()
}

type ExtendDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendDefContext() *ExtendDefContext {
	var p = new(ExtendDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_extendDef
	return p
}

func InitEmptyExtendDefContext(p *ExtendDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_extendDef
}

func (*ExtendDefContext) IsExtendDefContext() {}

func NewExtendDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendDefContext {
	var p = new(ExtendDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_extendDef

	return p
}

func (s *ExtendDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendDefContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEXTEND, 0)
}

func (s *ExtendDefContext) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *ExtendDefContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *ExtendDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *ExtendDefContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *ExtendDefContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *ExtendDefContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *ExtendDefContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *ExtendDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterExtendDef(s)
	}
}

func (s *ExtendDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitExtendDef(s)
	}
}

func (s *ExtendDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitExtendDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ExtendDef() (localctx IExtendDefContext) {
	localctx = NewExtendDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf3ParserRULE_extendDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(Protobuf3ParserEXTEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.MessageType()
	}
	{
		p.SetState(352)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162411405159432190) != 0 {
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserOPTIONAL, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserEXTEND, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserDOT, Protobuf3ParserBOOL_LIT, Protobuf3ParserIDENTIFIER:
			{
				p.SetState(353)
				p.Field()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(354)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
		p.Match(Protobuf3ParserRC)
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

// IServiceDefContext is an interface to support dynamic dispatch.
type IServiceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	ServiceName() IServiceNameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllServiceElement() []IServiceElementContext
	ServiceElement(i int) IServiceElementContext

	// IsServiceDefContext differentiates from other interfaces.
	IsServiceDefContext()
}

type ServiceDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefContext() *ServiceDefContext {
	var p = new(ServiceDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceDef
	return p
}

func InitEmptyServiceDefContext(p *ServiceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceDef
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *ServiceDefContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceDefContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *ServiceDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *ServiceDefContext) AllServiceElement() []IServiceElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceElementContext); ok {
			len++
		}
	}

	tst := make([]IServiceElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceElementContext); ok {
			tst[i] = t.(IServiceElementContext)
			i++
		}
	}

	return tst
}

func (s *ServiceDefContext) ServiceElement(i int) IServiceElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceElementContext); ok {
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

	return t.(IServiceElementContext)
}

func (s *ServiceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceDef(s)
	}
}

func (s *ServiceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceDef(s)
	}
}

func (s *ServiceDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitServiceDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceDef() (localctx IServiceDefContext) {
	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf3ParserRULE_serviceDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(Protobuf3ParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.ServiceName()
	}
	{
		p.SetState(364)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283467841600) != 0 {
		{
			p.SetState(365)
			p.ServiceElement()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(371)
		p.Match(Protobuf3ParserRC)
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

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionStatement() IOptionStatementContext
	Rpc() IRpcContext
	EmptyStatement_() IEmptyStatement_Context

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceElement
	return p
}

func InitEmptyServiceElementContext(p *ServiceElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceElement
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ServiceElementContext) Rpc() IRpcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceElement(s)
	}
}

func (s *ServiceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceElement(s)
	}
}

func (s *ServiceElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitServiceElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf3ParserRULE_serviceElement)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.OptionStatement()
		}

	case Protobuf3ParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.Rpc()
		}

	case Protobuf3ParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(375)
			p.EmptyStatement_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPC() antlr.TerminalNode
	RpcName() IRpcNameContext
	AllLP() []antlr.TerminalNode
	LP(i int) antlr.TerminalNode
	AllMessageType() []IMessageTypeContext
	MessageType(i int) IMessageTypeContext
	AllRP() []antlr.TerminalNode
	RP(i int) antlr.TerminalNode
	RETURNS() antlr.TerminalNode
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	AllSTREAM() []antlr.TerminalNode
	STREAM(i int) antlr.TerminalNode
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
	return p
}

func InitEmptyRpcContext(p *RpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllLP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserLP)
}

func (s *RpcContext) LP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, i)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageTypeContext); ok {
			len++
		}
	}

	tst := make([]IMessageTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageTypeContext); ok {
			tst[i] = t.(IMessageTypeContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
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

	return t.(IMessageTypeContext)
}

func (s *RpcContext) AllRP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserRP)
}

func (s *RpcContext) RP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *RpcContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *RpcContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *RpcContext) AllSTREAM() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserSTREAM)
}

func (s *RpcContext) STREAM(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, i)
}

func (s *RpcContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *RpcContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpc(s)
	}
}

func (s *RpcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf3ParserRULE_rpc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(Protobuf3ParserRPC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.RpcName()
	}
	{
		p.SetState(380)
		p.Match(Protobuf3ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(381)
			p.Match(Protobuf3ParserSTREAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(384)
		p.MessageType()
	}
	{
		p.SetState(385)
		p.Match(Protobuf3ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(Protobuf3ParserRETURNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Match(Protobuf3ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(388)
			p.Match(Protobuf3ParserSTREAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(391)
		p.MessageType()
	}
	{
		p.SetState(392)
		p.Match(Protobuf3ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserLC:
		{
			p.SetState(393)
			p.Match(Protobuf3ParserLC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserSEMI {
			p.SetState(396)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserOPTION:
				{
					p.SetState(394)
					p.OptionStatement()
				}

			case Protobuf3ParserSEMI:
				{
					p.SetState(395)
					p.EmptyStatement_()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(401)
			p.Match(Protobuf3ParserRC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf3ParserSEMI:
		{
			p.SetState(402)
			p.Match(Protobuf3ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FullIdent() IFullIdentContext
	IntLit() IIntLitContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	FloatLit() IFloatLitContext
	StrLit() IStrLitContext
	BoolLit() IBoolLitContext
	BlockLit() IBlockLitContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf3ParserRULE_constant)
	var _la int

	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(406)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(409)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(410)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(413)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(414)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(415)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(416)
			p.BlockLit()
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

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_blockLit
	return p
}

func InitEmptyBlockLitContext(p *BlockLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_blockLit
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
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

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterBlockLit(s)
	}
}

func (s *BlockLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitBlockLit(s)
	}
}

func (s *BlockLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitBlockLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) BlockLit() (localctx IBlockLitContext) {
	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf3ParserRULE_blockLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&162129655304814590) != 0 {
		{
			p.SetState(420)
			p.Ident()
		}
		{
			p.SetState(421)
			p.Match(Protobuf3ParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Constant()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(429)
		p.Match(Protobuf3ParserRC)
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

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_
	return p
}

func InitEmptyEmptyStatement_Context(p *EmptyStatement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf3ParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Match(Protobuf3ParserSEMI)
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

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Keywords() IKeywordsContext

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf3ParserRULE_ident)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(Protobuf3ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserOPTIONAL, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserEXTEND, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.Keywords()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf3ParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Ident()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserDOT {
		{
			p.SetState(438)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Ident()
		}

		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
	return p
}

func InitEmptyMessageNameContext(p *MessageNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (s *MessageNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf3ParserRULE_messageName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Ident()
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

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
	return p
}

func InitEmptyEnumNameContext(p *EnumNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (s *EnumNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf3ParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Ident()
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

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Protobuf3ParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Ident()
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

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
	return p
}

func InitEmptyOneofNameContext(p *OneofNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (s *OneofNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneofName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Protobuf3ParserRULE_oneofName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Ident()
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

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
	return p
}

func InitEmptyMapNameContext(p *MapNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapName(s)
	}
}

func (s *MapNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMapName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Protobuf3ParserRULE_mapName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Ident()
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

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
	return p
}

func InitEmptyServiceNameContext(p *ServiceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Protobuf3ParserRULE_serviceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Ident()
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

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
	return p
}

func InitEmptyRpcNameContext(p *RpcNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (s *RpcNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRpcName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Protobuf3ParserRULE_rpcName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.Ident()
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

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageName() IMessageNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
	return p
}

func InitEmptyMessageTypeContext(p *MessageTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (s *MessageTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Protobuf3ParserRULE_messageType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(459)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(462)
				p.Ident()
			}
			{
				p.SetState(463)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.MessageName()
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

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumName() IEnumNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumType
	return p
}

func InitEmptyEnumTypeContext(p *EnumTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumType
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (s *EnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Protobuf3ParserRULE_enumType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(472)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(475)
				p.Ident()
			}
			{
				p.SetState(476)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(483)
		p.EnumName()
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

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LIT() antlr.TerminalNode

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_intLit
	return p
}

func InitEmptyIntLitContext(p *IntLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_intLit
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, Protobuf3ParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Match(Protobuf3ParserINT_LIT)
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

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR_LIT() antlr.TerminalNode
	PROTO3_LIT_SINGLE() antlr.TerminalNode
	PROTO3_LIT_DOBULE() antlr.TerminalNode

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_strLit
	return p
}

func InitEmptyStrLitContext(p *StrLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_strLit
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitStrLit(s)
	}
}

func (s *StrLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitStrLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, Protobuf3ParserRULE_strLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9007405413171200) != 0) {
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

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL_LIT() antlr.TerminalNode

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_boolLit
	return p
}

func InitEmptyBoolLitContext(p *BoolLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_boolLit
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, Protobuf3ParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(Protobuf3ParserBOOL_LIT)
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

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LIT() antlr.TerminalNode

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_floatLit
	return p
}

func InitEmptyFloatLitContext(p *FloatLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_floatLit
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, Protobuf3ParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(Protobuf3ParserFLOAT_LIT)
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

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNTAX() antlr.TerminalNode
	IMPORT() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode
	PACKAGE() antlr.TerminalNode
	OPTION() antlr.TerminalNode
	OPTIONAL() antlr.TerminalNode
	REPEATED() antlr.TerminalNode
	ONEOF() antlr.TerminalNode
	MAP() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	RESERVED() antlr.TerminalNode
	TO() antlr.TerminalNode
	MAX() antlr.TerminalNode
	ENUM() antlr.TerminalNode
	MESSAGE() antlr.TerminalNode
	SERVICE() antlr.TerminalNode
	EXTEND() antlr.TerminalNode
	RPC() antlr.TerminalNode
	STREAM() antlr.TerminalNode
	RETURNS() antlr.TerminalNode
	BOOL_LIT() antlr.TerminalNode

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *KeywordsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *KeywordsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *KeywordsContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *KeywordsContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTIONAL, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *KeywordsContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *KeywordsContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *KeywordsContext) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *KeywordsContext) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *KeywordsContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEXTEND, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *KeywordsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *KeywordsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, Protobuf3ParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18014467228958718) != 0) {
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
