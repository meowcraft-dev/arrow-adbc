// Code generated from QueryLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type QueryLanguageLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var QueryLanguageLexerLexerStaticData struct {
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

func querylanguagelexerLexerInit() {
	staticData := &QueryLanguageLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'list'", "'<'", "'>'", "'struct'", "'int8'", "'string'",
		"'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "IDENTIFIER",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 70, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 5, 8, 59, 8, 8, 10, 8, 12, 8, 62, 9, 8, 1, 9, 4, 9, 65, 8, 9, 11, 9,
		12, 9, 66, 1, 9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 71, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0,
		5, 28, 1, 0, 0, 0, 7, 30, 1, 0, 0, 0, 9, 32, 1, 0, 0, 0, 11, 39, 1, 0,
		0, 0, 13, 44, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 56, 1, 0, 0, 0, 19, 64,
		1, 0, 0, 0, 21, 22, 5, 44, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 108, 0,
		0, 24, 25, 5, 105, 0, 0, 25, 26, 5, 115, 0, 0, 26, 27, 5, 116, 0, 0, 27,
		4, 1, 0, 0, 0, 28, 29, 5, 60, 0, 0, 29, 6, 1, 0, 0, 0, 30, 31, 5, 62, 0,
		0, 31, 8, 1, 0, 0, 0, 32, 33, 5, 115, 0, 0, 33, 34, 5, 116, 0, 0, 34, 35,
		5, 114, 0, 0, 35, 36, 5, 117, 0, 0, 36, 37, 5, 99, 0, 0, 37, 38, 5, 116,
		0, 0, 38, 10, 1, 0, 0, 0, 39, 40, 5, 105, 0, 0, 40, 41, 5, 110, 0, 0, 41,
		42, 5, 116, 0, 0, 42, 43, 5, 56, 0, 0, 43, 12, 1, 0, 0, 0, 44, 45, 5, 115,
		0, 0, 45, 46, 5, 116, 0, 0, 46, 47, 5, 114, 0, 0, 47, 48, 5, 105, 0, 0,
		48, 49, 5, 110, 0, 0, 49, 50, 5, 103, 0, 0, 50, 14, 1, 0, 0, 0, 51, 52,
		5, 98, 0, 0, 52, 53, 5, 111, 0, 0, 53, 54, 5, 111, 0, 0, 54, 55, 5, 108,
		0, 0, 55, 16, 1, 0, 0, 0, 56, 60, 7, 0, 0, 0, 57, 59, 7, 1, 0, 0, 58, 57,
		1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0,
		61, 18, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 65, 7, 2, 0, 0, 64, 63, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67,
		68, 1, 0, 0, 0, 68, 69, 6, 9, 0, 0, 69, 20, 1, 0, 0, 0, 3, 0, 60, 66, 1,
		6, 0, 0,
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

// QueryLanguageLexerInit initializes any static state used to implement QueryLanguageLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLanguageLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLanguageLexerInit() {
	staticData := &QueryLanguageLexerLexerStaticData
	staticData.once.Do(querylanguagelexerLexerInit)
}

// NewQueryLanguageLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLanguageLexer(input antlr.CharStream) *QueryLanguageLexer {
	QueryLanguageLexerInit()
	l := new(QueryLanguageLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &QueryLanguageLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "QueryLanguage.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLanguageLexer tokens.
const (
	QueryLanguageLexerT__0       = 1
	QueryLanguageLexerT__1       = 2
	QueryLanguageLexerT__2       = 3
	QueryLanguageLexerT__3       = 4
	QueryLanguageLexerT__4       = 5
	QueryLanguageLexerT__5       = 6
	QueryLanguageLexerT__6       = 7
	QueryLanguageLexerT__7       = 8
	QueryLanguageLexerIDENTIFIER = 9
	QueryLanguageLexerWS         = 10
)
