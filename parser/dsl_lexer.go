// Code generated from dsl.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type dslLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DslLexerLexerStaticData struct {
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

func dsllexerLexerInit() {
	staticData := &DslLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'types:'", "','", "'('", "')'", "'action'", "'{'", "'}'", "'monitor'",
		"':'", "'var'", "'rand'", "'tag'", "'store'", "'type'", "'=>'", "'<='",
		"'.'", "'\"'", "'int'", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "Wildcard", "Assign", "StringLiteral", "DigitSeq", "Comment",
		"Whitespace", "Newline",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "Wildcard", "Assign", "StringLiteral", "Char", "DigitSeq",
		"Digit", "Comment", "Whitespace", "Newline",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 175, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 4, 21, 136, 8, 21, 11, 21, 12, 21, 137, 1, 22,
		1, 22, 1, 23, 4, 23, 143, 8, 23, 11, 23, 12, 23, 144, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 5, 25, 153, 8, 25, 10, 25, 12, 25, 156, 9, 25,
		1, 25, 1, 25, 1, 26, 4, 26, 161, 8, 26, 11, 26, 12, 26, 162, 1, 26, 1,
		26, 1, 27, 1, 27, 3, 27, 169, 8, 27, 1, 27, 3, 27, 172, 8, 27, 1, 27, 1,
		27, 0, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 0, 47, 23, 49, 0, 51, 24, 53, 25, 55,
		26, 1, 0, 4, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0,
		10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 178, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0,
		0, 0, 55, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 64, 1, 0, 0, 0, 5, 66, 1, 0,
		0, 0, 7, 68, 1, 0, 0, 0, 9, 70, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13, 79,
		1, 0, 0, 0, 15, 81, 1, 0, 0, 0, 17, 89, 1, 0, 0, 0, 19, 91, 1, 0, 0, 0,
		21, 95, 1, 0, 0, 0, 23, 100, 1, 0, 0, 0, 25, 104, 1, 0, 0, 0, 27, 110,
		1, 0, 0, 0, 29, 115, 1, 0, 0, 0, 31, 118, 1, 0, 0, 0, 33, 121, 1, 0, 0,
		0, 35, 123, 1, 0, 0, 0, 37, 125, 1, 0, 0, 0, 39, 129, 1, 0, 0, 0, 41, 132,
		1, 0, 0, 0, 43, 135, 1, 0, 0, 0, 45, 139, 1, 0, 0, 0, 47, 142, 1, 0, 0,
		0, 49, 146, 1, 0, 0, 0, 51, 148, 1, 0, 0, 0, 53, 160, 1, 0, 0, 0, 55, 171,
		1, 0, 0, 0, 57, 58, 5, 116, 0, 0, 58, 59, 5, 121, 0, 0, 59, 60, 5, 112,
		0, 0, 60, 61, 5, 101, 0, 0, 61, 62, 5, 115, 0, 0, 62, 63, 5, 58, 0, 0,
		63, 2, 1, 0, 0, 0, 64, 65, 5, 44, 0, 0, 65, 4, 1, 0, 0, 0, 66, 67, 5, 40,
		0, 0, 67, 6, 1, 0, 0, 0, 68, 69, 5, 41, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71,
		5, 97, 0, 0, 71, 72, 5, 99, 0, 0, 72, 73, 5, 116, 0, 0, 73, 74, 5, 105,
		0, 0, 74, 75, 5, 111, 0, 0, 75, 76, 5, 110, 0, 0, 76, 10, 1, 0, 0, 0, 77,
		78, 5, 123, 0, 0, 78, 12, 1, 0, 0, 0, 79, 80, 5, 125, 0, 0, 80, 14, 1,
		0, 0, 0, 81, 82, 5, 109, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 110, 0,
		0, 84, 85, 5, 105, 0, 0, 85, 86, 5, 116, 0, 0, 86, 87, 5, 111, 0, 0, 87,
		88, 5, 114, 0, 0, 88, 16, 1, 0, 0, 0, 89, 90, 5, 58, 0, 0, 90, 18, 1, 0,
		0, 0, 91, 92, 5, 118, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5, 114, 0, 0,
		94, 20, 1, 0, 0, 0, 95, 96, 5, 114, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98,
		5, 110, 0, 0, 98, 99, 5, 100, 0, 0, 99, 22, 1, 0, 0, 0, 100, 101, 5, 116,
		0, 0, 101, 102, 5, 97, 0, 0, 102, 103, 5, 103, 0, 0, 103, 24, 1, 0, 0,
		0, 104, 105, 5, 115, 0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5, 111, 0,
		0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 101, 0, 0, 109, 26, 1, 0, 0, 0,
		110, 111, 5, 116, 0, 0, 111, 112, 5, 121, 0, 0, 112, 113, 5, 112, 0, 0,
		113, 114, 5, 101, 0, 0, 114, 28, 1, 0, 0, 0, 115, 116, 5, 61, 0, 0, 116,
		117, 5, 62, 0, 0, 117, 30, 1, 0, 0, 0, 118, 119, 5, 60, 0, 0, 119, 120,
		5, 61, 0, 0, 120, 32, 1, 0, 0, 0, 121, 122, 5, 46, 0, 0, 122, 34, 1, 0,
		0, 0, 123, 124, 5, 34, 0, 0, 124, 36, 1, 0, 0, 0, 125, 126, 5, 105, 0,
		0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 116, 0, 0, 128, 38, 1, 0, 0, 0,
		129, 130, 5, 95, 0, 0, 130, 131, 3, 47, 23, 0, 131, 40, 1, 0, 0, 0, 132,
		133, 5, 61, 0, 0, 133, 42, 1, 0, 0, 0, 134, 136, 3, 45, 22, 0, 135, 134,
		1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0,
		0, 0, 138, 44, 1, 0, 0, 0, 139, 140, 7, 0, 0, 0, 140, 46, 1, 0, 0, 0, 141,
		143, 3, 49, 24, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 48, 1, 0, 0, 0, 146, 147, 7, 1,
		0, 0, 147, 50, 1, 0, 0, 0, 148, 149, 5, 47, 0, 0, 149, 150, 5, 47, 0, 0,
		150, 154, 1, 0, 0, 0, 151, 153, 8, 2, 0, 0, 152, 151, 1, 0, 0, 0, 153,
		156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157,
		1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 6, 25, 0, 0, 158, 52, 1, 0,
		0, 0, 159, 161, 7, 3, 0, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0,
		162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		165, 6, 26, 1, 0, 165, 54, 1, 0, 0, 0, 166, 168, 5, 13, 0, 0, 167, 169,
		5, 10, 0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 172, 1, 0,
		0, 0, 170, 172, 5, 10, 0, 0, 171, 166, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 6, 27, 1, 0, 174, 56, 1, 0, 0, 0, 7, 0,
		137, 144, 154, 162, 168, 171, 2, 6, 0, 0, 0, 1, 0,
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

// dslLexerInit initializes any static state used to implement dslLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewdslLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DslLexerInit() {
	staticData := &DslLexerLexerStaticData
	staticData.once.Do(dsllexerLexerInit)
}

// NewdslLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewdslLexer(input antlr.CharStream) *dslLexer {
	DslLexerInit()
	l := new(dslLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DslLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "dsl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// dslLexer tokens.
const (
	dslLexerT__0          = 1
	dslLexerT__1          = 2
	dslLexerT__2          = 3
	dslLexerT__3          = 4
	dslLexerT__4          = 5
	dslLexerT__5          = 6
	dslLexerT__6          = 7
	dslLexerT__7          = 8
	dslLexerT__8          = 9
	dslLexerT__9          = 10
	dslLexerT__10         = 11
	dslLexerT__11         = 12
	dslLexerT__12         = 13
	dslLexerT__13         = 14
	dslLexerT__14         = 15
	dslLexerT__15         = 16
	dslLexerT__16         = 17
	dslLexerT__17         = 18
	dslLexerT__18         = 19
	dslLexerWildcard      = 20
	dslLexerAssign        = 21
	dslLexerStringLiteral = 22
	dslLexerDigitSeq      = 23
	dslLexerComment       = 24
	dslLexerWhitespace    = 25
	dslLexerNewline       = 26
)
