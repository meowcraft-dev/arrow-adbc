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

// EnterFields is called when production fields is entered.
func (s *BaseQueryLanguageListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseQueryLanguageListener) ExitFields(ctx *FieldsContext) {}

// EnterStructFields is called when production structFields is entered.
func (s *BaseQueryLanguageListener) EnterStructFields(ctx *StructFieldsContext) {}

// ExitStructFields is called when production structFields is exited.
func (s *BaseQueryLanguageListener) ExitStructFields(ctx *StructFieldsContext) {}

// EnterInnerType is called when production innerType is entered.
func (s *BaseQueryLanguageListener) EnterInnerType(ctx *InnerTypeContext) {}

// ExitInnerType is called when production innerType is exited.
func (s *BaseQueryLanguageListener) ExitInnerType(ctx *InnerTypeContext) {}

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

// EnterSimpleTypes is called when production simpleTypes is entered.
func (s *BaseQueryLanguageListener) EnterSimpleTypes(ctx *SimpleTypesContext) {}

// ExitSimpleTypes is called when production simpleTypes is exited.
func (s *BaseQueryLanguageListener) ExitSimpleTypes(ctx *SimpleTypesContext) {}
