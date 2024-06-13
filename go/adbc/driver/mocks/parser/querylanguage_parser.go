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
		"", "','", "'list'", "'<'", "'>'", "'struct'", "'uint8'", "'int8'",
		"'uint16'", "'int16'", "'uint32'", "'int32'", "'uint64'", "'int64'",
		"'float16'", "'float32'", "'float64'", "'binary'", "'string'", "'date32'",
		"'date64'", "'bool'", "'time32s'", "'time32ms'", "'time64us'", "'time64ns'",
		"'timestamp_s'", "'timestamp_ms'", "'timestamp_us'", "'timestamp_ns'",
		"'duration_s'", "'duration_ms'", "'duration_us'", "'duration_ns'", "'interval_month'",
		"'interval_daytime'", "'interval_monthdaynano'", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "WS", "ROWCOUNT", "FIELDNAME",
	}
	staticData.RuleNames = []string{
		"query", "fields", "innerType", "list", "struct", "simpleTypes",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 61, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 3, 0, 14, 8, 0, 1, 0, 1, 0, 1, 0, 5, 0, 19, 8, 0,
		10, 0, 12, 0, 22, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1,
		1, 3, 1, 32, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3,
		3, 3, 42, 8, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 52,
		8, 4, 10, 4, 12, 4, 55, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0,
		2, 4, 6, 8, 10, 0, 1, 1, 0, 6, 37, 63, 0, 13, 1, 0, 0, 0, 2, 28, 1, 0,
		0, 0, 4, 36, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 58,
		1, 0, 0, 0, 12, 14, 5, 39, 0, 0, 13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0,
		14, 15, 1, 0, 0, 0, 15, 20, 3, 2, 1, 0, 16, 17, 5, 1, 0, 0, 17, 19, 3,
		2, 1, 0, 18, 16, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20,
		21, 1, 0, 0, 0, 21, 23, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 5, 0, 0,
		1, 24, 1, 1, 0, 0, 0, 25, 29, 3, 10, 5, 0, 26, 29, 3, 6, 3, 0, 27, 29,
		3, 8, 4, 0, 28, 25, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0,
		29, 31, 1, 0, 0, 0, 30, 32, 5, 40, 0, 0, 31, 30, 1, 0, 0, 0, 31, 32, 1,
		0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 37, 3, 10, 5, 0, 34, 37, 3, 6, 3, 0, 35,
		37, 3, 8, 4, 0, 36, 33, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0,
		0, 37, 5, 1, 0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 41, 5, 3, 0, 0, 40, 42, 5,
		39, 0, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43,
		44, 3, 4, 2, 0, 44, 45, 5, 4, 0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 5, 0,
		0, 47, 48, 5, 3, 0, 0, 48, 53, 3, 4, 2, 0, 49, 50, 5, 1, 0, 0, 50, 52,
		3, 4, 2, 0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0,
		53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5,
		4, 0, 0, 57, 9, 1, 0, 0, 0, 58, 59, 7, 0, 0, 0, 59, 11, 1, 0, 0, 0, 7,
		13, 20, 28, 31, 36, 41, 53,
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
	QueryLanguageParserWS        = 38
	QueryLanguageParserROWCOUNT  = 39
	QueryLanguageParserFIELDNAME = 40
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_query       = 0
	QueryLanguageParserRULE_fields      = 1
	QueryLanguageParserRULE_innerType   = 2
	QueryLanguageParserRULE_list        = 3
	QueryLanguageParserRULE_struct      = 4
	QueryLanguageParserRULE_simpleTypes = 5
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserROWCOUNT {
		{
			p.SetState(12)
			p.Match(QueryLanguageParserROWCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(15)
		p.Fields()
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(16)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Fields()
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__5, QueryLanguageParserT__6, QueryLanguageParserT__7, QueryLanguageParserT__8, QueryLanguageParserT__9, QueryLanguageParserT__10, QueryLanguageParserT__11, QueryLanguageParserT__12, QueryLanguageParserT__13, QueryLanguageParserT__14, QueryLanguageParserT__15, QueryLanguageParserT__16, QueryLanguageParserT__17, QueryLanguageParserT__18, QueryLanguageParserT__19, QueryLanguageParserT__20, QueryLanguageParserT__21, QueryLanguageParserT__22, QueryLanguageParserT__23, QueryLanguageParserT__24, QueryLanguageParserT__25, QueryLanguageParserT__26, QueryLanguageParserT__27, QueryLanguageParserT__28, QueryLanguageParserT__29, QueryLanguageParserT__30, QueryLanguageParserT__31, QueryLanguageParserT__32, QueryLanguageParserT__33, QueryLanguageParserT__34, QueryLanguageParserT__35, QueryLanguageParserT__36:
		{
			p.SetState(25)
			p.SimpleTypes()
		}

	case QueryLanguageParserT__1:
		{
			p.SetState(26)
			p.List()
		}

	case QueryLanguageParserT__4:
		{
			p.SetState(27)
			p.Struct_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserFIELDNAME {
		{
			p.SetState(30)
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
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_innerType)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__5, QueryLanguageParserT__6, QueryLanguageParserT__7, QueryLanguageParserT__8, QueryLanguageParserT__9, QueryLanguageParserT__10, QueryLanguageParserT__11, QueryLanguageParserT__12, QueryLanguageParserT__13, QueryLanguageParserT__14, QueryLanguageParserT__15, QueryLanguageParserT__16, QueryLanguageParserT__17, QueryLanguageParserT__18, QueryLanguageParserT__19, QueryLanguageParserT__20, QueryLanguageParserT__21, QueryLanguageParserT__22, QueryLanguageParserT__23, QueryLanguageParserT__24, QueryLanguageParserT__25, QueryLanguageParserT__26, QueryLanguageParserT__27, QueryLanguageParserT__28, QueryLanguageParserT__29, QueryLanguageParserT__30, QueryLanguageParserT__31, QueryLanguageParserT__32, QueryLanguageParserT__33, QueryLanguageParserT__34, QueryLanguageParserT__35, QueryLanguageParserT__36:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.SimpleTypes()
		}

	case QueryLanguageParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.List()
		}

	case QueryLanguageParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
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
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == QueryLanguageParserROWCOUNT {
		{
			p.SetState(40)
			p.Match(QueryLanguageParserROWCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(43)
		p.InnerType()
	}
	{
		p.SetState(44)
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
	AllInnerType() []IInnerTypeContext
	InnerType(i int) IInnerTypeContext

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

func (s *StructContext) AllInnerType() []IInnerTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInnerTypeContext); ok {
			len++
		}
	}

	tst := make([]IInnerTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInnerTypeContext); ok {
			tst[i] = t.(IInnerTypeContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) InnerType(i int) IInnerTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeContext); ok {
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

	return t.(IInnerTypeContext)
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
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(QueryLanguageParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.InnerType()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(49)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.InnerType()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
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
	p.EnterRule(localctx, 10, QueryLanguageParserRULE_simpleTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877906880) != 0) {
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
