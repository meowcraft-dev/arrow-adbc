// Code generated from QueryLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // QueryLanguage

import "github.com/antlr4-go/antlr/v4"

// BaseQueryLanguageListener is a complete listener for a parse tree produced by QueryLanguageParser.
type BaseQueryLanguageListener struct{}

var _ QueryLanguageListener = &BaseQueryLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQueryLanguageListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQueryLanguageListener) ExitQuery(ctx *QueryContext) {}

// EnterTopLevelField is called when production topLevelField is entered.
func (s *BaseQueryLanguageListener) EnterTopLevelField(ctx *TopLevelFieldContext) {}

// ExitTopLevelField is called when production topLevelField is exited.
func (s *BaseQueryLanguageListener) ExitTopLevelField(ctx *TopLevelFieldContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseQueryLanguageListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseQueryLanguageListener) ExitStructField(ctx *StructFieldContext) {}

// EnterUnionField is called when production unionField is entered.
func (s *BaseQueryLanguageListener) EnterUnionField(ctx *UnionFieldContext) {}

// ExitUnionField is called when production unionField is exited.
func (s *BaseQueryLanguageListener) ExitUnionField(ctx *UnionFieldContext) {}

// EnterType is called when production type is entered.
func (s *BaseQueryLanguageListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseQueryLanguageListener) ExitType(ctx *TypeContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseQueryLanguageListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseQueryLanguageListener) ExitUnion(ctx *UnionContext) {}

// EnterDictionary is called when production dictionary is entered.
func (s *BaseQueryLanguageListener) EnterDictionary(ctx *DictionaryContext) {}

// ExitDictionary is called when production dictionary is exited.
func (s *BaseQueryLanguageListener) ExitDictionary(ctx *DictionaryContext) {}

// EnterList is called when production list is entered.
func (s *BaseQueryLanguageListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseQueryLanguageListener) ExitList(ctx *ListContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseQueryLanguageListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseQueryLanguageListener) ExitStruct(ctx *StructContext) {}

// EnterFixedSizeBinary is called when production fixedSizeBinary is entered.
func (s *BaseQueryLanguageListener) EnterFixedSizeBinary(ctx *FixedSizeBinaryContext) {}

// ExitFixedSizeBinary is called when production fixedSizeBinary is exited.
func (s *BaseQueryLanguageListener) ExitFixedSizeBinary(ctx *FixedSizeBinaryContext) {}

// EnterDecimal128 is called when production decimal128 is entered.
func (s *BaseQueryLanguageListener) EnterDecimal128(ctx *Decimal128Context) {}

// ExitDecimal128 is called when production decimal128 is exited.
func (s *BaseQueryLanguageListener) ExitDecimal128(ctx *Decimal128Context) {}

// EnterDecimal256 is called when production decimal256 is entered.
func (s *BaseQueryLanguageListener) EnterDecimal256(ctx *Decimal256Context) {}

// ExitDecimal256 is called when production decimal256 is exited.
func (s *BaseQueryLanguageListener) ExitDecimal256(ctx *Decimal256Context) {}

// EnterRunEndEncoded is called when production runEndEncoded is entered.
func (s *BaseQueryLanguageListener) EnterRunEndEncoded(ctx *RunEndEncodedContext) {}

// ExitRunEndEncoded is called when production runEndEncoded is exited.
func (s *BaseQueryLanguageListener) ExitRunEndEncoded(ctx *RunEndEncodedContext) {}

// EnterSimpleTypes is called when production simpleTypes is entered.
func (s *BaseQueryLanguageListener) EnterSimpleTypes(ctx *SimpleTypesContext) {}

// ExitSimpleTypes is called when production simpleTypes is exited.
func (s *BaseQueryLanguageListener) ExitSimpleTypes(ctx *SimpleTypesContext) {}
