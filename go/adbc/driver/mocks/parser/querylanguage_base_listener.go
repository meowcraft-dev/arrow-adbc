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

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseQueryLanguageListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseQueryLanguageListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterList is called when production list is entered.
func (s *BaseQueryLanguageListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseQueryLanguageListener) ExitList(ctx *ListContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseQueryLanguageListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseQueryLanguageListener) ExitStruct(ctx *StructContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseQueryLanguageListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseQueryLanguageListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}
