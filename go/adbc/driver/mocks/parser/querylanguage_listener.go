// Code generated from QueryLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // QueryLanguage

import "github.com/antlr4-go/antlr/v4"

// QueryLanguageListener is a complete listener for a parse tree produced by QueryLanguageParser.
type QueryLanguageListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterStructFields is called when entering the structFields production.
	EnterStructFields(c *StructFieldsContext)

	// EnterInnerType is called when entering the innerType production.
	EnterInnerType(c *InnerTypeContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterSimpleTypes is called when entering the simpleTypes production.
	EnterSimpleTypes(c *SimpleTypesContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitStructFields is called when exiting the structFields production.
	ExitStructFields(c *StructFieldsContext)

	// ExitInnerType is called when exiting the innerType production.
	ExitInnerType(c *InnerTypeContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitSimpleTypes is called when exiting the simpleTypes production.
	ExitSimpleTypes(c *SimpleTypesContext)
}
