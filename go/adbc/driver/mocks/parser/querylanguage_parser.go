// Code generated from QueryLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // QueryLanguage

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

type QueryLanguageParser struct {
	*antlr.BaseParser
}

var QueryLanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylanguageParserInit() {
	staticData := &QueryLanguageParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'<'", "'%'", "'>'", "'dict'", "'dictionary'", "'list'",
		"'struct'", "'fixed_size_binary'", "'decimal128'", "'ree'", "'run_end_encoded'",
		"'_'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "'timestamp_s'", "'timestamp_ms'", "'timestamp_us'",
		"'timestamp_ns'", "'duration_s'", "'duration_ms'", "'duration_us'",
		"'duration_ns'", "'interval_month'", "'interval_daytime'", "'interval_monthdaynano'",
		"'sparse_union'", "'dense_union'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NULL", "BOOL",
		"UINT8", "INT8", "UINT16", "INT16", "UINT32", "INT32", "UINT64", "INT64",
		"FLOAT16", "FLOAT32", "FLOAT64", "BINARY", "STRING", "DATE32", "DATE64",
		"TIME32S", "TIME32MS", "TIME64US", "TIME64NS", "TIMESTAMP_S", "TIMESTAMP_MS",
		"TIMESTAMP_US", "TIMESTAMP_NS", "DURATION_S", "DURATION_MS", "DURATION_US",
		"DURATION_NS", "INTERVAL_MONTH", "INTERVAL_DAYTIME", "INTERVAL_MONTHDAYNANO",
		"SPARSE_UNION", "DENSE_UNION", "DECIMAL_PS", "COUNT", "FIELD_NAME",
		"UNION_VALUE_NAME", "BYTE_WIDTH", "DICT_ENTRY", "WS",
	}
	staticData.RuleNames = []string{
		"query", "topLevelField", "structField", "unionField", "type", "union",
		"dictionary", "list", "struct", "fixedSizeBinary", "decimal128", "decimal256",
		"runEndEncoded", "unionValue", "dictEntry", "simpleTypes",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 153, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 1, 0, 5, 0, 39, 8, 0, 10, 0, 12, 0, 42,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 2, 1, 2, 3, 2, 52, 8,
		2, 1, 3, 1, 3, 3, 3, 56, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 67, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 74, 8,
		5, 10, 5, 12, 5, 77, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 83, 8, 5, 10,
		5, 12, 5, 86, 9, 5, 3, 5, 88, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 100, 8, 6, 10, 6, 12, 6, 103, 9, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 110, 8, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 120, 8, 8, 10, 8, 12, 8, 123, 9, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 6, 1, 0, 46, 47, 1, 0, 5, 6, 1,
		0, 11, 12, 2, 0, 13, 13, 51, 51, 2, 0, 13, 13, 53, 53, 1, 0, 14, 45, 155,
		0, 33, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 53, 1, 0, 0,
		0, 8, 66, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 106,
		1, 0, 0, 0, 16, 114, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 131, 1, 0, 0,
		0, 22, 136, 1, 0, 0, 0, 24, 141, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 148,
		1, 0, 0, 0, 30, 150, 1, 0, 0, 0, 32, 34, 5, 49, 0, 0, 33, 32, 1, 0, 0,
		0, 33, 34, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 40, 3, 2, 1, 0, 36, 37,
		5, 1, 0, 0, 37, 39, 3, 2, 1, 0, 38, 36, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0,
		40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 40, 1,
		0, 0, 0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 47, 3, 8, 4, 0, 46,
		48, 5, 50, 0, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 3, 1, 0, 0,
		0, 49, 51, 3, 8, 4, 0, 50, 52, 5, 50, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52,
		1, 0, 0, 0, 52, 5, 1, 0, 0, 0, 53, 55, 3, 8, 4, 0, 54, 56, 5, 50, 0, 0,
		55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 67, 3, 30,
		15, 0, 58, 67, 3, 14, 7, 0, 59, 67, 3, 16, 8, 0, 60, 67, 3, 12, 6, 0, 61,
		67, 3, 18, 9, 0, 62, 67, 3, 20, 10, 0, 63, 67, 3, 22, 11, 0, 64, 67, 3,
		10, 5, 0, 65, 67, 3, 24, 12, 0, 66, 57, 1, 0, 0, 0, 66, 58, 1, 0, 0, 0,
		66, 59, 1, 0, 0, 0, 66, 60, 1, 0, 0, 0, 66, 61, 1, 0, 0, 0, 66, 62, 1,
		0, 0, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67,
		9, 1, 0, 0, 0, 68, 69, 7, 0, 0, 0, 69, 70, 5, 2, 0, 0, 70, 75, 3, 6, 3,
		0, 71, 72, 5, 1, 0, 0, 72, 74, 3, 6, 3, 0, 73, 71, 1, 0, 0, 0, 74, 77,
		1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 87, 1, 0, 0, 0,
		77, 75, 1, 0, 0, 0, 78, 79, 5, 3, 0, 0, 79, 84, 3, 26, 13, 0, 80, 81, 5,
		1, 0, 0, 81, 83, 3, 26, 13, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1,
		0, 0, 0, 87, 78, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		90, 5, 4, 0, 0, 90, 11, 1, 0, 0, 0, 91, 92, 7, 1, 0, 0, 92, 93, 5, 2, 0,
		0, 93, 94, 5, 49, 0, 0, 94, 95, 3, 8, 4, 0, 95, 96, 5, 1, 0, 0, 96, 101,
		3, 28, 14, 0, 97, 98, 5, 1, 0, 0, 98, 100, 3, 28, 14, 0, 99, 97, 1, 0,
		0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0,
		102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 4, 0, 0, 105,
		13, 1, 0, 0, 0, 106, 107, 5, 7, 0, 0, 107, 109, 5, 2, 0, 0, 108, 110, 5,
		49, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0,
		0, 111, 112, 3, 8, 4, 0, 112, 113, 5, 4, 0, 0, 113, 15, 1, 0, 0, 0, 114,
		115, 5, 8, 0, 0, 115, 116, 5, 2, 0, 0, 116, 121, 3, 4, 2, 0, 117, 118,
		5, 1, 0, 0, 118, 120, 3, 4, 2, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0,
		123, 121, 1, 0, 0, 0, 124, 125, 5, 4, 0, 0, 125, 17, 1, 0, 0, 0, 126, 127,
		5, 9, 0, 0, 127, 128, 5, 2, 0, 0, 128, 129, 5, 52, 0, 0, 129, 130, 5, 4,
		0, 0, 130, 19, 1, 0, 0, 0, 131, 132, 5, 10, 0, 0, 132, 133, 5, 2, 0, 0,
		133, 134, 5, 48, 0, 0, 134, 135, 5, 4, 0, 0, 135, 21, 1, 0, 0, 0, 136,
		137, 5, 10, 0, 0, 137, 138, 5, 2, 0, 0, 138, 139, 5, 48, 0, 0, 139, 140,
		5, 4, 0, 0, 140, 23, 1, 0, 0, 0, 141, 142, 7, 2, 0, 0, 142, 143, 5, 2,
		0, 0, 143, 144, 3, 8, 4, 0, 144, 145, 5, 4, 0, 0, 145, 25, 1, 0, 0, 0,
		146, 147, 7, 3, 0, 0, 147, 27, 1, 0, 0, 0, 148, 149, 7, 4, 0, 0, 149, 29,
		1, 0, 0, 0, 150, 151, 7, 5, 0, 0, 151, 31, 1, 0, 0, 0, 12, 33, 40, 47,
		51, 55, 66, 75, 84, 87, 101, 109, 121,
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

// QueryLanguageParserInit initializes any static state used to implement QueryLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQueryLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLanguageParserInit() {
	staticData := &QueryLanguageParserStaticData
	staticData.once.Do(querylanguageParserInit)
}

// NewQueryLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQueryLanguageParser(input antlr.TokenStream) *QueryLanguageParser {
	QueryLanguageParserInit()
	this := new(QueryLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &QueryLanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "QueryLanguage.g4"

	return this
}

// QueryLanguageParser tokens.
const (
	QueryLanguageParserEOF                   = antlr.TokenEOF
	QueryLanguageParserT__0                  = 1
	QueryLanguageParserT__1                  = 2
	QueryLanguageParserT__2                  = 3
	QueryLanguageParserT__3                  = 4
	QueryLanguageParserT__4                  = 5
	QueryLanguageParserT__5                  = 6
	QueryLanguageParserT__6                  = 7
	QueryLanguageParserT__7                  = 8
	QueryLanguageParserT__8                  = 9
	QueryLanguageParserT__9                  = 10
	QueryLanguageParserT__10                 = 11
	QueryLanguageParserT__11                 = 12
	QueryLanguageParserT__12                 = 13
	QueryLanguageParserNULL                  = 14
	QueryLanguageParserBOOL                  = 15
	QueryLanguageParserUINT8                 = 16
	QueryLanguageParserINT8                  = 17
	QueryLanguageParserUINT16                = 18
	QueryLanguageParserINT16                 = 19
	QueryLanguageParserUINT32                = 20
	QueryLanguageParserINT32                 = 21
	QueryLanguageParserUINT64                = 22
	QueryLanguageParserINT64                 = 23
	QueryLanguageParserFLOAT16               = 24
	QueryLanguageParserFLOAT32               = 25
	QueryLanguageParserFLOAT64               = 26
	QueryLanguageParserBINARY                = 27
	QueryLanguageParserSTRING                = 28
	QueryLanguageParserDATE32                = 29
	QueryLanguageParserDATE64                = 30
	QueryLanguageParserTIME32S               = 31
	QueryLanguageParserTIME32MS              = 32
	QueryLanguageParserTIME64US              = 33
	QueryLanguageParserTIME64NS              = 34
	QueryLanguageParserTIMESTAMP_S           = 35
	QueryLanguageParserTIMESTAMP_MS          = 36
	QueryLanguageParserTIMESTAMP_US          = 37
	QueryLanguageParserTIMESTAMP_NS          = 38
	QueryLanguageParserDURATION_S            = 39
	QueryLanguageParserDURATION_MS           = 40
	QueryLanguageParserDURATION_US           = 41
	QueryLanguageParserDURATION_NS           = 42
	QueryLanguageParserINTERVAL_MONTH        = 43
	QueryLanguageParserINTERVAL_DAYTIME      = 44
	QueryLanguageParserINTERVAL_MONTHDAYNANO = 45
	QueryLanguageParserSPARSE_UNION          = 46
	QueryLanguageParserDENSE_UNION           = 47
	QueryLanguageParserDECIMAL_PS            = 48
	QueryLanguageParserCOUNT                 = 49
	QueryLanguageParserFIELD_NAME            = 50
	QueryLanguageParserUNION_VALUE_NAME      = 51
	QueryLanguageParserBYTE_WIDTH            = 52
	QueryLanguageParserDICT_ENTRY            = 53
	QueryLanguageParserWS                    = 54
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_query           = 0
	QueryLanguageParserRULE_topLevelField   = 1
	QueryLanguageParserRULE_structField     = 2
	QueryLanguageParserRULE_unionField      = 3
	QueryLanguageParserRULE_type            = 4
	QueryLanguageParserRULE_union           = 5
	QueryLanguageParserRULE_dictionary      = 6
	QueryLanguageParserRULE_list            = 7
	QueryLanguageParserRULE_struct          = 8
	QueryLanguageParserRULE_fixedSizeBinary = 9
	QueryLanguageParserRULE_decimal128      = 10
	QueryLanguageParserRULE_decimal256      = 11
	QueryLanguageParserRULE_runEndEncoded   = 12
	QueryLanguageParserRULE_unionValue      = 13
	QueryLanguageParserRULE_dictEntry       = 14
	QueryLanguageParserRULE_simpleTypes     = 15
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTopLevelField() []ITopLevelFieldContext
	TopLevelField(i int) ITopLevelFieldContext
	EOF() antlr.TerminalNode
	COUNT() antlr.TerminalNode

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) AllTopLevelField() []ITopLevelFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelFieldContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelFieldContext); ok {
			tst[i] = t.(ITopLevelFieldContext)
			i++
		}
	}

	return tst
}

func (s *QueryContext) TopLevelField(i int) ITopLevelFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelFieldContext); ok {
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

	return t.(ITopLevelFieldContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserEOF, 0)
}

func (s *QueryContext) COUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserCOUNT, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *QueryLanguageParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryLanguageParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserCOUNT {
		{
			p.SetState(32)
			p.Match(QueryLanguageParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(35)
		p.TopLevelField()
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(36)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.TopLevelField()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(QueryLanguageParserEOF)
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

// ITopLevelFieldContext is an interface to support dynamic dispatch.
type ITopLevelFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	FIELD_NAME() antlr.TerminalNode

	// IsTopLevelFieldContext differentiates from other interfaces.
	IsTopLevelFieldContext()
}

type TopLevelFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelFieldContext() *TopLevelFieldContext {
	var p = new(TopLevelFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_topLevelField
	return p
}

func InitEmptyTopLevelFieldContext(p *TopLevelFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_topLevelField
}

func (*TopLevelFieldContext) IsTopLevelFieldContext() {}

func NewTopLevelFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelFieldContext {
	var p = new(TopLevelFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_topLevelField

	return p
}

func (s *TopLevelFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelFieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TopLevelFieldContext) FIELD_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFIELD_NAME, 0)
}

func (s *TopLevelFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterTopLevelField(s)
	}
}

func (s *TopLevelFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitTopLevelField(s)
	}
}

func (p *QueryLanguageParser) TopLevelField() (localctx ITopLevelFieldContext) {
	localctx = NewTopLevelFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryLanguageParserRULE_topLevelField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Type_()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELD_NAME {
		{
			p.SetState(46)
			p.Match(QueryLanguageParserFIELD_NAME)
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

// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	FIELD_NAME() antlr.TerminalNode

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructFieldContext) FIELD_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFIELD_NAME, 0)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (p *QueryLanguageParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_structField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Type_()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELD_NAME {
		{
			p.SetState(50)
			p.Match(QueryLanguageParserFIELD_NAME)
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

// IUnionFieldContext is an interface to support dynamic dispatch.
type IUnionFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	FIELD_NAME() antlr.TerminalNode

	// IsUnionFieldContext differentiates from other interfaces.
	IsUnionFieldContext()
}

type UnionFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionFieldContext() *UnionFieldContext {
	var p = new(UnionFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_unionField
	return p
}

func InitEmptyUnionFieldContext(p *UnionFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_unionField
}

func (*UnionFieldContext) IsUnionFieldContext() {}

func NewUnionFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionFieldContext {
	var p = new(UnionFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_unionField

	return p
}

func (s *UnionFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionFieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *UnionFieldContext) FIELD_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFIELD_NAME, 0)
}

func (s *UnionFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterUnionField(s)
	}
}

func (s *UnionFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitUnionField(s)
	}
}

func (p *QueryLanguageParser) UnionField() (localctx IUnionFieldContext) {
	localctx = NewUnionFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_unionField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Type_()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELD_NAME {
		{
			p.SetState(54)
			p.Match(QueryLanguageParserFIELD_NAME)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleTypes() ISimpleTypesContext
	List() IListContext
	Struct_() IStructContext
	Dictionary() IDictionaryContext
	FixedSizeBinary() IFixedSizeBinaryContext
	Decimal128() IDecimal128Context
	Decimal256() IDecimal256Context
	Union() IUnionContext
	RunEndEncoded() IRunEndEncodedContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) SimpleTypes() ISimpleTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTypesContext)
}

func (s *TypeContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *TypeContext) Struct_() IStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructContext)
}

func (s *TypeContext) Dictionary() IDictionaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictionaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictionaryContext)
}

func (s *TypeContext) FixedSizeBinary() IFixedSizeBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixedSizeBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixedSizeBinaryContext)
}

func (s *TypeContext) Decimal128() IDecimal128Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimal128Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimal128Context)
}

func (s *TypeContext) Decimal256() IDecimal256Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimal256Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimal256Context)
}

func (s *TypeContext) Union() IUnionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *TypeContext) RunEndEncoded() IRunEndEncodedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunEndEncodedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunEndEncodedContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *QueryLanguageParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_type)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.SimpleTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.List()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Struct_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
			p.Dictionary()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)
			p.FixedSizeBinary()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(62)
			p.Decimal128()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(63)
			p.Decimal256()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(64)
			p.Union()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(65)
			p.RunEndEncoded()
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

// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnionField() []IUnionFieldContext
	UnionField(i int) IUnionFieldContext
	SPARSE_UNION() antlr.TerminalNode
	DENSE_UNION() antlr.TerminalNode
	AllUnionValue() []IUnionValueContext
	UnionValue(i int) IUnionValueContext

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	var p = new(UnionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_union
	return p
}

func InitEmptyUnionContext(p *UnionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_union
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	var p = new(UnionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) AllUnionField() []IUnionFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnionFieldContext); ok {
			len++
		}
	}

	tst := make([]IUnionFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnionFieldContext); ok {
			tst[i] = t.(IUnionFieldContext)
			i++
		}
	}

	return tst
}

func (s *UnionContext) UnionField(i int) IUnionFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionFieldContext); ok {
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

	return t.(IUnionFieldContext)
}

func (s *UnionContext) SPARSE_UNION() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserSPARSE_UNION, 0)
}

func (s *UnionContext) DENSE_UNION() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDENSE_UNION, 0)
}

func (s *UnionContext) AllUnionValue() []IUnionValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnionValueContext); ok {
			len++
		}
	}

	tst := make([]IUnionValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnionValueContext); ok {
			tst[i] = t.(IUnionValueContext)
			i++
		}
	}

	return tst
}

func (s *UnionContext) UnionValue(i int) IUnionValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionValueContext); ok {
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

	return t.(IUnionValueContext)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (p *QueryLanguageParser) Union() (localctx IUnionContext) {
	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryLanguageParserRULE_union)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserSPARSE_UNION || _la == QueryLanguageParserDENSE_UNION) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(69)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.UnionField()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(71)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.UnionField()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserT__2 {
		{
			p.SetState(78)
			p.Match(QueryLanguageParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.UnionValue()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == QueryLanguageParserT__0 {
			{
				p.SetState(80)
				p.Match(QueryLanguageParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.UnionValue()
			}

			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(89)
		p.Match(QueryLanguageParserT__3)
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

// IDictionaryContext is an interface to support dynamic dispatch.
type IDictionaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	Type_() ITypeContext
	AllDictEntry() []IDictEntryContext
	DictEntry(i int) IDictEntryContext

	// IsDictionaryContext differentiates from other interfaces.
	IsDictionaryContext()
}

type DictionaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictionaryContext() *DictionaryContext {
	var p = new(DictionaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_dictionary
	return p
}

func InitEmptyDictionaryContext(p *DictionaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_dictionary
}

func (*DictionaryContext) IsDictionaryContext() {}

func NewDictionaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryContext {
	var p = new(DictionaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_dictionary

	return p
}

func (s *DictionaryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryContext) COUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserCOUNT, 0)
}

func (s *DictionaryContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DictionaryContext) AllDictEntry() []IDictEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictEntryContext); ok {
			len++
		}
	}

	tst := make([]IDictEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictEntryContext); ok {
			tst[i] = t.(IDictEntryContext)
			i++
		}
	}

	return tst
}

func (s *DictionaryContext) DictEntry(i int) IDictEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictEntryContext); ok {
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

	return t.(IDictEntryContext)
}

func (s *DictionaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterDictionary(s)
	}
}

func (s *DictionaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitDictionary(s)
	}
}

func (p *QueryLanguageParser) Dictionary() (localctx IDictionaryContext) {
	localctx = NewDictionaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryLanguageParserRULE_dictionary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserT__4 || _la == QueryLanguageParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(92)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(QueryLanguageParserCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Type_()
	}
	{
		p.SetState(95)
		p.Match(QueryLanguageParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.DictEntry()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(97)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.DictEntry()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(QueryLanguageParserT__3)
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	COUNT() antlr.TerminalNode

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListContext) COUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserCOUNT, 0)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *QueryLanguageParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryLanguageParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(QueryLanguageParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserCOUNT {
		{
			p.SetState(108)
			p.Match(QueryLanguageParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(111)
		p.Type_()
	}
	{
		p.SetState(112)
		p.Match(QueryLanguageParserT__3)
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

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	var p = new(StructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_struct
	return p
}

func InitEmptyStructContext(p *StructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_struct
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	var p = new(StructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
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

	return t.(IStructFieldContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (p *QueryLanguageParser) Struct_() (localctx IStructContext) {
	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryLanguageParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(QueryLanguageParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.StructField()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(117)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.StructField()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(QueryLanguageParserT__3)
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

// IFixedSizeBinaryContext is an interface to support dynamic dispatch.
type IFixedSizeBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BYTE_WIDTH() antlr.TerminalNode

	// IsFixedSizeBinaryContext differentiates from other interfaces.
	IsFixedSizeBinaryContext()
}

type FixedSizeBinaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedSizeBinaryContext() *FixedSizeBinaryContext {
	var p = new(FixedSizeBinaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fixedSizeBinary
	return p
}

func InitEmptyFixedSizeBinaryContext(p *FixedSizeBinaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fixedSizeBinary
}

func (*FixedSizeBinaryContext) IsFixedSizeBinaryContext() {}

func NewFixedSizeBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedSizeBinaryContext {
	var p = new(FixedSizeBinaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_fixedSizeBinary

	return p
}

func (s *FixedSizeBinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedSizeBinaryContext) BYTE_WIDTH() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserBYTE_WIDTH, 0)
}

func (s *FixedSizeBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedSizeBinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedSizeBinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterFixedSizeBinary(s)
	}
}

func (s *FixedSizeBinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitFixedSizeBinary(s)
	}
}

func (p *QueryLanguageParser) FixedSizeBinary() (localctx IFixedSizeBinaryContext) {
	localctx = NewFixedSizeBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryLanguageParserRULE_fixedSizeBinary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(QueryLanguageParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(QueryLanguageParserBYTE_WIDTH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(QueryLanguageParserT__3)
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

// IDecimal128Context is an interface to support dynamic dispatch.
type IDecimal128Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_PS() antlr.TerminalNode

	// IsDecimal128Context differentiates from other interfaces.
	IsDecimal128Context()
}

type Decimal128Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimal128Context() *Decimal128Context {
	var p = new(Decimal128Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_decimal128
	return p
}

func InitEmptyDecimal128Context(p *Decimal128Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_decimal128
}

func (*Decimal128Context) IsDecimal128Context() {}

func NewDecimal128Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decimal128Context {
	var p = new(Decimal128Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_decimal128

	return p
}

func (s *Decimal128Context) GetParser() antlr.Parser { return s.parser }

func (s *Decimal128Context) DECIMAL_PS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDECIMAL_PS, 0)
}

func (s *Decimal128Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decimal128Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decimal128Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterDecimal128(s)
	}
}

func (s *Decimal128Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitDecimal128(s)
	}
}

func (p *QueryLanguageParser) Decimal128() (localctx IDecimal128Context) {
	localctx = NewDecimal128Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QueryLanguageParserRULE_decimal128)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(QueryLanguageParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(QueryLanguageParserDECIMAL_PS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(QueryLanguageParserT__3)
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

// IDecimal256Context is an interface to support dynamic dispatch.
type IDecimal256Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_PS() antlr.TerminalNode

	// IsDecimal256Context differentiates from other interfaces.
	IsDecimal256Context()
}

type Decimal256Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimal256Context() *Decimal256Context {
	var p = new(Decimal256Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_decimal256
	return p
}

func InitEmptyDecimal256Context(p *Decimal256Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_decimal256
}

func (*Decimal256Context) IsDecimal256Context() {}

func NewDecimal256Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decimal256Context {
	var p = new(Decimal256Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_decimal256

	return p
}

func (s *Decimal256Context) GetParser() antlr.Parser { return s.parser }

func (s *Decimal256Context) DECIMAL_PS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDECIMAL_PS, 0)
}

func (s *Decimal256Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decimal256Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decimal256Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterDecimal256(s)
	}
}

func (s *Decimal256Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitDecimal256(s)
	}
}

func (p *QueryLanguageParser) Decimal256() (localctx IDecimal256Context) {
	localctx = NewDecimal256Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QueryLanguageParserRULE_decimal256)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(QueryLanguageParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(QueryLanguageParserDECIMAL_PS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(QueryLanguageParserT__3)
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

// IRunEndEncodedContext is an interface to support dynamic dispatch.
type IRunEndEncodedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsRunEndEncodedContext differentiates from other interfaces.
	IsRunEndEncodedContext()
}

type RunEndEncodedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunEndEncodedContext() *RunEndEncodedContext {
	var p = new(RunEndEncodedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_runEndEncoded
	return p
}

func InitEmptyRunEndEncodedContext(p *RunEndEncodedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_runEndEncoded
}

func (*RunEndEncodedContext) IsRunEndEncodedContext() {}

func NewRunEndEncodedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunEndEncodedContext {
	var p = new(RunEndEncodedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_runEndEncoded

	return p
}

func (s *RunEndEncodedContext) GetParser() antlr.Parser { return s.parser }

func (s *RunEndEncodedContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *RunEndEncodedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunEndEncodedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunEndEncodedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterRunEndEncoded(s)
	}
}

func (s *RunEndEncodedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitRunEndEncoded(s)
	}
}

func (p *QueryLanguageParser) RunEndEncoded() (localctx IRunEndEncodedContext) {
	localctx = NewRunEndEncodedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QueryLanguageParserRULE_runEndEncoded)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserT__10 || _la == QueryLanguageParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(142)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Type_()
	}
	{
		p.SetState(144)
		p.Match(QueryLanguageParserT__3)
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

// IUnionValueContext is an interface to support dynamic dispatch.
type IUnionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNION_VALUE_NAME() antlr.TerminalNode

	// IsUnionValueContext differentiates from other interfaces.
	IsUnionValueContext()
}

type UnionValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionValueContext() *UnionValueContext {
	var p = new(UnionValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_unionValue
	return p
}

func InitEmptyUnionValueContext(p *UnionValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_unionValue
}

func (*UnionValueContext) IsUnionValueContext() {}

func NewUnionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionValueContext {
	var p = new(UnionValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_unionValue

	return p
}

func (s *UnionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionValueContext) UNION_VALUE_NAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserUNION_VALUE_NAME, 0)
}

func (s *UnionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterUnionValue(s)
	}
}

func (s *UnionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitUnionValue(s)
	}
}

func (p *QueryLanguageParser) UnionValue() (localctx IUnionValueContext) {
	localctx = NewUnionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QueryLanguageParserRULE_unionValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserT__12 || _la == QueryLanguageParserUNION_VALUE_NAME) {
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

// IDictEntryContext is an interface to support dynamic dispatch.
type IDictEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DICT_ENTRY() antlr.TerminalNode

	// IsDictEntryContext differentiates from other interfaces.
	IsDictEntryContext()
}

type DictEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryContext() *DictEntryContext {
	var p = new(DictEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_dictEntry
	return p
}

func InitEmptyDictEntryContext(p *DictEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_dictEntry
}

func (*DictEntryContext) IsDictEntryContext() {}

func NewDictEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryContext {
	var p = new(DictEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_dictEntry

	return p
}

func (s *DictEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryContext) DICT_ENTRY() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDICT_ENTRY, 0)
}

func (s *DictEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterDictEntry(s)
	}
}

func (s *DictEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitDictEntry(s)
	}
}

func (p *QueryLanguageParser) DictEntry() (localctx IDictEntryContext) {
	localctx = NewDictEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, QueryLanguageParserRULE_dictEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryLanguageParserT__12 || _la == QueryLanguageParserDICT_ENTRY) {
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

// ISimpleTypesContext is an interface to support dynamic dispatch.
type ISimpleTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UINT8() antlr.TerminalNode
	INT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	INT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	INT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	INT64() antlr.TerminalNode
	FLOAT16() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	BINARY() antlr.TerminalNode
	STRING() antlr.TerminalNode
	DATE32() antlr.TerminalNode
	DATE64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	TIME32S() antlr.TerminalNode
	TIME32MS() antlr.TerminalNode
	TIME64US() antlr.TerminalNode
	TIME64NS() antlr.TerminalNode
	TIMESTAMP_S() antlr.TerminalNode
	TIMESTAMP_MS() antlr.TerminalNode
	TIMESTAMP_US() antlr.TerminalNode
	TIMESTAMP_NS() antlr.TerminalNode
	DURATION_S() antlr.TerminalNode
	DURATION_MS() antlr.TerminalNode
	DURATION_US() antlr.TerminalNode
	DURATION_NS() antlr.TerminalNode
	INTERVAL_MONTH() antlr.TerminalNode
	INTERVAL_DAYTIME() antlr.TerminalNode
	INTERVAL_MONTHDAYNANO() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsSimpleTypesContext differentiates from other interfaces.
	IsSimpleTypesContext()
}

type SimpleTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTypesContext() *SimpleTypesContext {
	var p = new(SimpleTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_simpleTypes
	return p
}

func InitEmptySimpleTypesContext(p *SimpleTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_simpleTypes
}

func (*SimpleTypesContext) IsSimpleTypesContext() {}

func NewSimpleTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTypesContext {
	var p = new(SimpleTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_simpleTypes

	return p
}

func (s *SimpleTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTypesContext) UINT8() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserUINT8, 0)
}

func (s *SimpleTypesContext) INT8() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINT8, 0)
}

func (s *SimpleTypesContext) UINT16() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserUINT16, 0)
}

func (s *SimpleTypesContext) INT16() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINT16, 0)
}

func (s *SimpleTypesContext) UINT32() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserUINT32, 0)
}

func (s *SimpleTypesContext) INT32() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINT32, 0)
}

func (s *SimpleTypesContext) UINT64() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserUINT64, 0)
}

func (s *SimpleTypesContext) INT64() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINT64, 0)
}

func (s *SimpleTypesContext) FLOAT16() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFLOAT16, 0)
}

func (s *SimpleTypesContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFLOAT32, 0)
}

func (s *SimpleTypesContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFLOAT64, 0)
}

func (s *SimpleTypesContext) BINARY() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserBINARY, 0)
}

func (s *SimpleTypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserSTRING, 0)
}

func (s *SimpleTypesContext) DATE32() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDATE32, 0)
}

func (s *SimpleTypesContext) DATE64() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDATE64, 0)
}

func (s *SimpleTypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserBOOL, 0)
}

func (s *SimpleTypesContext) TIME32S() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIME32S, 0)
}

func (s *SimpleTypesContext) TIME32MS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIME32MS, 0)
}

func (s *SimpleTypesContext) TIME64US() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIME64US, 0)
}

func (s *SimpleTypesContext) TIME64NS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIME64NS, 0)
}

func (s *SimpleTypesContext) TIMESTAMP_S() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIMESTAMP_S, 0)
}

func (s *SimpleTypesContext) TIMESTAMP_MS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIMESTAMP_MS, 0)
}

func (s *SimpleTypesContext) TIMESTAMP_US() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIMESTAMP_US, 0)
}

func (s *SimpleTypesContext) TIMESTAMP_NS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserTIMESTAMP_NS, 0)
}

func (s *SimpleTypesContext) DURATION_S() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDURATION_S, 0)
}

func (s *SimpleTypesContext) DURATION_MS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDURATION_MS, 0)
}

func (s *SimpleTypesContext) DURATION_US() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDURATION_US, 0)
}

func (s *SimpleTypesContext) DURATION_NS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDURATION_NS, 0)
}

func (s *SimpleTypesContext) INTERVAL_MONTH() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINTERVAL_MONTH, 0)
}

func (s *SimpleTypesContext) INTERVAL_DAYTIME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINTERVAL_DAYTIME, 0)
}

func (s *SimpleTypesContext) INTERVAL_MONTHDAYNANO() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserINTERVAL_MONTHDAYNANO, 0)
}

func (s *SimpleTypesContext) NULL() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserNULL, 0)
}

func (s *SimpleTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterSimpleTypes(s)
	}
}

func (s *SimpleTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitSimpleTypes(s)
	}
}

func (p *QueryLanguageParser) SimpleTypes() (localctx ISimpleTypesContext) {
	localctx = NewSimpleTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, QueryLanguageParserRULE_simpleTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70368744161280) != 0) {
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
