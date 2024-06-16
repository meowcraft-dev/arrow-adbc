// Code generated from QueryLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // QueryLanguage

import "github.com/antlr4-go/antlr/v4"

// QueryLanguageListener is a complete listener for a parse tree produced by QueryLanguageParser.
type QueryLanguageListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterTopLevelField is called when entering the topLevelField production.
	EnterTopLevelField(c *TopLevelFieldContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterUnionField is called when entering the unionField production.
	EnterUnionField(c *UnionFieldContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterDictionary is called when entering the dictionary production.
	EnterDictionary(c *DictionaryContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterFixedSizeBinary is called when entering the fixedSizeBinary production.
	EnterFixedSizeBinary(c *FixedSizeBinaryContext)

	// EnterDecimal128 is called when entering the decimal128 production.
	EnterDecimal128(c *Decimal128Context)

	// EnterDecimal256 is called when entering the decimal256 production.
	EnterDecimal256(c *Decimal256Context)

	// EnterRunEndEncoded is called when entering the runEndEncoded production.
	EnterRunEndEncoded(c *RunEndEncodedContext)

	// EnterUnionValue is called when entering the unionValue production.
	EnterUnionValue(c *UnionValueContext)

	// EnterDictEntry is called when entering the dictEntry production.
	EnterDictEntry(c *DictEntryContext)

	// EnterSimpleTypes is called when entering the simpleTypes production.
	EnterSimpleTypes(c *SimpleTypesContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitTopLevelField is called when exiting the topLevelField production.
	ExitTopLevelField(c *TopLevelFieldContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitUnionField is called when exiting the unionField production.
	ExitUnionField(c *UnionFieldContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitDictionary is called when exiting the dictionary production.
	ExitDictionary(c *DictionaryContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitFixedSizeBinary is called when exiting the fixedSizeBinary production.
	ExitFixedSizeBinary(c *FixedSizeBinaryContext)

	// ExitDecimal128 is called when exiting the decimal128 production.
	ExitDecimal128(c *Decimal128Context)

	// ExitDecimal256 is called when exiting the decimal256 production.
	ExitDecimal256(c *Decimal256Context)

	// ExitRunEndEncoded is called when exiting the runEndEncoded production.
	ExitRunEndEncoded(c *RunEndEncodedContext)

	// ExitUnionValue is called when exiting the unionValue production.
	ExitUnionValue(c *UnionValueContext)

	// ExitDictEntry is called when exiting the dictEntry production.
	ExitDictEntry(c *DictEntryContext)

	// ExitSimpleTypes is called when exiting the simpleTypes production.
	ExitSimpleTypes(c *SimpleTypesContext)
}
