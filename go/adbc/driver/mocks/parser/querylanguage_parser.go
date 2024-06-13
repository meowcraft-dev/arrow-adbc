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
		"", "','", "'list'", "'<'", "'>'", "'struct'", "'fixed_size_binary'",
		"'decimal128'", "'uint8'", "'int8'", "'uint16'", "'int16'", "'uint32'",
		"'int32'", "'uint64'", "'int64'", "'float16'", "'float32'", "'float64'",
		"'binary'", "'string'", "'date32'", "'date64'", "'bool'", "'time32s'",
		"'time32ms'", "'time64us'", "'time64ns'", "'timestamp_s'", "'timestamp_ms'",
		"'timestamp_us'", "'timestamp_ns'", "'duration_s'", "'duration_ms'",
		"'duration_us'", "'duration_ns'", "'interval_month'", "'interval_daytime'",
		"'interval_monthdaynano'", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "DECIMALPS", "ROWCOUNT", "BYTEWIDTH", "FIELDNAME",
		"WS",
	}
	staticData.RuleNames = []string{
		"query", "fields", "structFields", "innerType", "list", "struct", "fixedSizeBinary",
		"decimal128", "decimal256", "simpleTypes",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 98, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 3,
		0, 22, 8, 0, 1, 0, 1, 0, 1, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0, 30, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 2,
		3, 2, 54, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 59, 8, 3, 1, 4, 1, 4, 1, 4, 3,
		4, 64, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 74,
		8, 5, 10, 5, 12, 5, 77, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 1, 1, 0, 8, 39, 105,
		0, 21, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 58, 1, 0, 0,
		0, 8, 60, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14, 85, 1,
		0, 0, 0, 16, 90, 1, 0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 22, 5, 41, 0, 0, 21,
		20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 28, 3, 2, 1,
		0, 24, 25, 5, 1, 0, 0, 25, 27, 3, 2, 1, 0, 26, 24, 1, 0, 0, 0, 27, 30,
		1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 31, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 31, 32, 5, 0, 0, 1, 32, 1, 1, 0, 0, 0, 33, 40, 3, 18,
		9, 0, 34, 40, 3, 8, 4, 0, 35, 40, 3, 10, 5, 0, 36, 40, 3, 12, 6, 0, 37,
		40, 3, 14, 7, 0, 38, 40, 3, 16, 8, 0, 39, 33, 1, 0, 0, 0, 39, 34, 1, 0,
		0, 0, 39, 35, 1, 0, 0, 0, 39, 36, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38,
		1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 43, 5, 43, 0, 0, 42, 41, 1, 0, 0, 0,
		42, 43, 1, 0, 0, 0, 43, 3, 1, 0, 0, 0, 44, 51, 3, 18, 9, 0, 45, 51, 3,
		8, 4, 0, 46, 51, 3, 10, 5, 0, 47, 51, 3, 12, 6, 0, 48, 51, 3, 14, 7, 0,
		49, 51, 3, 16, 8, 0, 50, 44, 1, 0, 0, 0, 50, 45, 1, 0, 0, 0, 50, 46, 1,
		0, 0, 0, 50, 47, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51,
		53, 1, 0, 0, 0, 52, 54, 5, 43, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0,
		0, 0, 54, 5, 1, 0, 0, 0, 55, 59, 3, 18, 9, 0, 56, 59, 3, 8, 4, 0, 57, 59,
		3, 10, 5, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0,
		59, 7, 1, 0, 0, 0, 60, 61, 5, 2, 0, 0, 61, 63, 5, 3, 0, 0, 62, 64, 5, 41,
		0, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66,
		3, 6, 3, 0, 66, 67, 5, 4, 0, 0, 67, 9, 1, 0, 0, 0, 68, 69, 5, 5, 0, 0,
		69, 70, 5, 3, 0, 0, 70, 75, 3, 4, 2, 0, 71, 72, 5, 1, 0, 0, 72, 74, 3,
		4, 2, 0, 73, 71, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 4, 0,
		0, 79, 11, 1, 0, 0, 0, 80, 81, 5, 6, 0, 0, 81, 82, 5, 3, 0, 0, 82, 83,
		5, 42, 0, 0, 83, 84, 5, 4, 0, 0, 84, 13, 1, 0, 0, 0, 85, 86, 5, 7, 0, 0,
		86, 87, 5, 3, 0, 0, 87, 88, 5, 40, 0, 0, 88, 89, 5, 4, 0, 0, 89, 15, 1,
		0, 0, 0, 90, 91, 5, 7, 0, 0, 91, 92, 5, 3, 0, 0, 92, 93, 5, 40, 0, 0, 93,
		94, 5, 4, 0, 0, 94, 17, 1, 0, 0, 0, 95, 96, 7, 0, 0, 0, 96, 19, 1, 0, 0,
		0, 9, 21, 28, 39, 42, 50, 53, 58, 63, 75,
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
	QueryLanguageParserEOF       = antlr.TokenEOF
	QueryLanguageParserT__0      = 1
	QueryLanguageParserT__1      = 2
	QueryLanguageParserT__2      = 3
	QueryLanguageParserT__3      = 4
	QueryLanguageParserT__4      = 5
	QueryLanguageParserT__5      = 6
	QueryLanguageParserT__6      = 7
	QueryLanguageParserT__7      = 8
	QueryLanguageParserT__8      = 9
	QueryLanguageParserT__9      = 10
	QueryLanguageParserT__10     = 11
	QueryLanguageParserT__11     = 12
	QueryLanguageParserT__12     = 13
	QueryLanguageParserT__13     = 14
	QueryLanguageParserT__14     = 15
	QueryLanguageParserT__15     = 16
	QueryLanguageParserT__16     = 17
	QueryLanguageParserT__17     = 18
	QueryLanguageParserT__18     = 19
	QueryLanguageParserT__19     = 20
	QueryLanguageParserT__20     = 21
	QueryLanguageParserT__21     = 22
	QueryLanguageParserT__22     = 23
	QueryLanguageParserT__23     = 24
	QueryLanguageParserT__24     = 25
	QueryLanguageParserT__25     = 26
	QueryLanguageParserT__26     = 27
	QueryLanguageParserT__27     = 28
	QueryLanguageParserT__28     = 29
	QueryLanguageParserT__29     = 30
	QueryLanguageParserT__30     = 31
	QueryLanguageParserT__31     = 32
	QueryLanguageParserT__32     = 33
	QueryLanguageParserT__33     = 34
	QueryLanguageParserT__34     = 35
	QueryLanguageParserT__35     = 36
	QueryLanguageParserT__36     = 37
	QueryLanguageParserT__37     = 38
	QueryLanguageParserT__38     = 39
	QueryLanguageParserDECIMALPS = 40
	QueryLanguageParserROWCOUNT  = 41
	QueryLanguageParserBYTEWIDTH = 42
	QueryLanguageParserFIELDNAME = 43
	QueryLanguageParserWS        = 44
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_query           = 0
	QueryLanguageParserRULE_fields          = 1
	QueryLanguageParserRULE_structFields    = 2
	QueryLanguageParserRULE_innerType       = 3
	QueryLanguageParserRULE_list            = 4
	QueryLanguageParserRULE_struct          = 5
	QueryLanguageParserRULE_fixedSizeBinary = 6
	QueryLanguageParserRULE_decimal128      = 7
	QueryLanguageParserRULE_decimal256      = 8
	QueryLanguageParserRULE_simpleTypes     = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFields() []IFieldsContext
	Fields(i int) IFieldsContext
	EOF() antlr.TerminalNode
	ROWCOUNT() antlr.TerminalNode

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

func (s *QueryContext) AllFields() []IFieldsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsContext); ok {
			len++
		}
	}

	tst := make([]IFieldsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsContext); ok {
			tst[i] = t.(IFieldsContext)
			i++
		}
	}

	return tst
}

func (s *QueryContext) Fields(i int) IFieldsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
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

	return t.(IFieldsContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserEOF, 0)
}

func (s *QueryContext) ROWCOUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserROWCOUNT, 0)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserROWCOUNT {
		{
			p.SetState(20)
			p.Match(QueryLanguageParserROWCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(23)
		p.Fields()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(24)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Fields()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
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

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleTypes() ISimpleTypesContext
	List() IListContext
	Struct_() IStructContext
	FixedSizeBinary() IFixedSizeBinaryContext
	Decimal128() IDecimal128Context
	Decimal256() IDecimal256Context
	FIELDNAME() antlr.TerminalNode

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fields
	return p
}

func InitEmptyFieldsContext(p *FieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fields
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) SimpleTypes() ISimpleTypesContext {
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

func (s *FieldsContext) List() IListContext {
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

func (s *FieldsContext) Struct_() IStructContext {
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

func (s *FieldsContext) FixedSizeBinary() IFixedSizeBinaryContext {
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

func (s *FieldsContext) Decimal128() IDecimal128Context {
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

func (s *FieldsContext) Decimal256() IDecimal256Context {
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

func (s *FieldsContext) FIELDNAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFIELDNAME, 0)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *QueryLanguageParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryLanguageParserRULE_fields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(33)
			p.SimpleTypes()
		}

	case 2:
		{
			p.SetState(34)
			p.List()
		}

	case 3:
		{
			p.SetState(35)
			p.Struct_()
		}

	case 4:
		{
			p.SetState(36)
			p.FixedSizeBinary()
		}

	case 5:
		{
			p.SetState(37)
			p.Decimal128()
		}

	case 6:
		{
			p.SetState(38)
			p.Decimal256()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELDNAME {
		{
			p.SetState(41)
			p.Match(QueryLanguageParserFIELDNAME)
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

// IStructFieldsContext is an interface to support dynamic dispatch.
type IStructFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleTypes() ISimpleTypesContext
	List() IListContext
	Struct_() IStructContext
	FixedSizeBinary() IFixedSizeBinaryContext
	Decimal128() IDecimal128Context
	Decimal256() IDecimal256Context
	FIELDNAME() antlr.TerminalNode

	// IsStructFieldsContext differentiates from other interfaces.
	IsStructFieldsContext()
}

type StructFieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldsContext() *StructFieldsContext {
	var p = new(StructFieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_structFields
	return p
}

func InitEmptyStructFieldsContext(p *StructFieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_structFields
}

func (*StructFieldsContext) IsStructFieldsContext() {}

func NewStructFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldsContext {
	var p = new(StructFieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_structFields

	return p
}

func (s *StructFieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldsContext) SimpleTypes() ISimpleTypesContext {
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

func (s *StructFieldsContext) List() IListContext {
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

func (s *StructFieldsContext) Struct_() IStructContext {
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

func (s *StructFieldsContext) FixedSizeBinary() IFixedSizeBinaryContext {
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

func (s *StructFieldsContext) Decimal128() IDecimal128Context {
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

func (s *StructFieldsContext) Decimal256() IDecimal256Context {
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

func (s *StructFieldsContext) FIELDNAME() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserFIELDNAME, 0)
}

func (s *StructFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterStructFields(s)
	}
}

func (s *StructFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitStructFields(s)
	}
}

func (p *QueryLanguageParser) StructFields() (localctx IStructFieldsContext) {
	localctx = NewStructFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_structFields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(44)
			p.SimpleTypes()
		}

	case 2:
		{
			p.SetState(45)
			p.List()
		}

	case 3:
		{
			p.SetState(46)
			p.Struct_()
		}

	case 4:
		{
			p.SetState(47)
			p.FixedSizeBinary()
		}

	case 5:
		{
			p.SetState(48)
			p.Decimal128()
		}

	case 6:
		{
			p.SetState(49)
			p.Decimal256()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELDNAME {
		{
			p.SetState(52)
			p.Match(QueryLanguageParserFIELDNAME)
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

// IInnerTypeContext is an interface to support dynamic dispatch.
type IInnerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleTypes() ISimpleTypesContext
	List() IListContext
	Struct_() IStructContext

	// IsInnerTypeContext differentiates from other interfaces.
	IsInnerTypeContext()
}

type InnerTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeContext() *InnerTypeContext {
	var p = new(InnerTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_innerType
	return p
}

func InitEmptyInnerTypeContext(p *InnerTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_innerType
}

func (*InnerTypeContext) IsInnerTypeContext() {}

func NewInnerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeContext {
	var p = new(InnerTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_innerType

	return p
}

func (s *InnerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeContext) SimpleTypes() ISimpleTypesContext {
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

func (s *InnerTypeContext) List() IListContext {
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

func (s *InnerTypeContext) Struct_() IStructContext {
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

func (s *InnerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterInnerType(s)
	}
}

func (s *InnerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitInnerType(s)
	}
}

func (p *QueryLanguageParser) InnerType() (localctx IInnerTypeContext) {
	localctx = NewInnerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_innerType)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__7, QueryLanguageParserT__8, QueryLanguageParserT__9, QueryLanguageParserT__10, QueryLanguageParserT__11, QueryLanguageParserT__12, QueryLanguageParserT__13, QueryLanguageParserT__14, QueryLanguageParserT__15, QueryLanguageParserT__16, QueryLanguageParserT__17, QueryLanguageParserT__18, QueryLanguageParserT__19, QueryLanguageParserT__20, QueryLanguageParserT__21, QueryLanguageParserT__22, QueryLanguageParserT__23, QueryLanguageParserT__24, QueryLanguageParserT__25, QueryLanguageParserT__26, QueryLanguageParserT__27, QueryLanguageParserT__28, QueryLanguageParserT__29, QueryLanguageParserT__30, QueryLanguageParserT__31, QueryLanguageParserT__32, QueryLanguageParserT__33, QueryLanguageParserT__34, QueryLanguageParserT__35, QueryLanguageParserT__36, QueryLanguageParserT__37, QueryLanguageParserT__38:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.SimpleTypes()
		}

	case QueryLanguageParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.List()
		}

	case QueryLanguageParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Struct_()
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InnerType() IInnerTypeContext
	ROWCOUNT() antlr.TerminalNode

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

func (s *ListContext) InnerType() IInnerTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeContext)
}

func (s *ListContext) ROWCOUNT() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserROWCOUNT, 0)
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
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserROWCOUNT {
		{
			p.SetState(62)
			p.Match(QueryLanguageParserROWCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(65)
		p.InnerType()
	}
	{
		p.SetState(66)
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
	AllStructFields() []IStructFieldsContext
	StructFields(i int) IStructFieldsContext

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

func (s *StructContext) AllStructFields() []IStructFieldsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldsContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldsContext); ok {
			tst[i] = t.(IStructFieldsContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) StructFields(i int) IStructFieldsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldsContext); ok {
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

	return t.(IStructFieldsContext)
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
	p.EnterRule(localctx, 10, QueryLanguageParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(QueryLanguageParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.StructFields()
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
			p.StructFields()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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
	BYTEWIDTH() antlr.TerminalNode

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

func (s *FixedSizeBinaryContext) BYTEWIDTH() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserBYTEWIDTH, 0)
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
	p.EnterRule(localctx, 12, QueryLanguageParserRULE_fixedSizeBinary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(QueryLanguageParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(QueryLanguageParserBYTEWIDTH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
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
	DECIMALPS() antlr.TerminalNode

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

func (s *Decimal128Context) DECIMALPS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDECIMALPS, 0)
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
	p.EnterRule(localctx, 14, QueryLanguageParserRULE_decimal128)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(QueryLanguageParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(QueryLanguageParserDECIMALPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
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
	DECIMALPS() antlr.TerminalNode

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

func (s *Decimal256Context) DECIMALPS() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserDECIMALPS, 0)
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
	p.EnterRule(localctx, 16, QueryLanguageParserRULE_decimal256)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(QueryLanguageParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(QueryLanguageParserDECIMALPS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
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

// ISimpleTypesContext is an interface to support dynamic dispatch.
type ISimpleTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.EnterRule(localctx, 18, QueryLanguageParserRULE_simpleTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1099511627520) != 0) {
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
