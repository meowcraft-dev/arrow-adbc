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
		"", "','", "'list'", "'<'", "'>'", "'struct'", "'int8'", "'string'",
		"'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"query", "typeSpec", "list", "struct", "primitiveType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 45, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 3, 1, 24, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 36, 8, 3, 10, 3, 12, 3, 39, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 1, 1, 0, 6, 8, 43, 0, 10,
		1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 25, 1, 0, 0, 0, 6, 30, 1, 0, 0, 0, 8,
		42, 1, 0, 0, 0, 10, 15, 3, 2, 1, 0, 11, 12, 5, 1, 0, 0, 12, 14, 3, 2, 1,
		0, 13, 11, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 18, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 0, 0, 1,
		19, 1, 1, 0, 0, 0, 20, 24, 3, 8, 4, 0, 21, 24, 3, 4, 2, 0, 22, 24, 3, 6,
		3, 0, 23, 20, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 22, 1, 0, 0, 0, 24, 3,
		1, 0, 0, 0, 25, 26, 5, 2, 0, 0, 26, 27, 5, 3, 0, 0, 27, 28, 3, 2, 1, 0,
		28, 29, 5, 4, 0, 0, 29, 5, 1, 0, 0, 0, 30, 31, 5, 5, 0, 0, 31, 32, 5, 3,
		0, 0, 32, 37, 3, 2, 1, 0, 33, 34, 5, 1, 0, 0, 34, 36, 3, 2, 1, 0, 35, 33,
		1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 4, 0, 0, 41, 7, 1, 0,
		0, 0, 42, 43, 7, 0, 0, 0, 43, 9, 1, 0, 0, 0, 3, 15, 23, 37,
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
	QueryLanguageParserEOF        = antlr.TokenEOF
	QueryLanguageParserT__0       = 1
	QueryLanguageParserT__1       = 2
	QueryLanguageParserT__2       = 3
	QueryLanguageParserT__3       = 4
	QueryLanguageParserT__4       = 5
	QueryLanguageParserT__5       = 6
	QueryLanguageParserT__6       = 7
	QueryLanguageParserT__7       = 8
	QueryLanguageParserIDENTIFIER = 9
	QueryLanguageParserWS         = 10
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_query         = 0
	QueryLanguageParserRULE_typeSpec      = 1
	QueryLanguageParserRULE_list          = 2
	QueryLanguageParserRULE_struct        = 3
	QueryLanguageParserRULE_primitiveType = 4
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeSpec() []ITypeSpecContext
	TypeSpec(i int) ITypeSpecContext
	EOF() antlr.TerminalNode

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

func (s *QueryContext) AllTypeSpec() []ITypeSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSpecContext); ok {
			len++
		}
	}

	tst := make([]ITypeSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSpecContext); ok {
			tst[i] = t.(ITypeSpecContext)
			i++
		}
	}

	return tst
}

func (s *QueryContext) TypeSpec(i int) ITypeSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
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

	return t.(ITypeSpecContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserEOF, 0)
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
	{
		p.SetState(10)
		p.TypeSpec()
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(11)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)
			p.TypeSpec()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
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

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	List() IListContext
	Struct_() IStructContext

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_typeSpec
	return p
}

func InitEmptyTypeSpecContext(p *TypeSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_typeSpec
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeSpecContext) List() IListContext {
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

func (s *TypeSpecContext) Struct_() IStructContext {
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

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (p *QueryLanguageParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryLanguageParserRULE_typeSpec)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__5, QueryLanguageParserT__6, QueryLanguageParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.PrimitiveType()
		}

	case QueryLanguageParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.List()
		}

	case QueryLanguageParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
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
	TypeSpec() ITypeSpecContext

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

func (s *ListContext) TypeSpec() ITypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
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
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(QueryLanguageParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.TypeSpec()
	}
	{
		p.SetState(28)
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
	AllTypeSpec() []ITypeSpecContext
	TypeSpec(i int) ITypeSpecContext

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

func (s *StructContext) AllTypeSpec() []ITypeSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSpecContext); ok {
			len++
		}
	}

	tst := make([]ITypeSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSpecContext); ok {
			tst[i] = t.(ITypeSpecContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) TypeSpec(i int) ITypeSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
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

	return t.(ITypeSpecContext)
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
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(QueryLanguageParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(QueryLanguageParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.TypeSpec()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == QueryLanguageParserT__0 {
		{
			p.SetState(33)
			p.Match(QueryLanguageParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.TypeSpec()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
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

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_primitiveType
	return p
}

func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_primitiveType
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (p *QueryLanguageParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
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
