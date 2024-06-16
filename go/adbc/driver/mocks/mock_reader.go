// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mocks

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
	"github.com/apache/arrow-adbc/go/adbc/driver/mocks/parser"
	"github.com/apache/arrow/go/v17/arrow"
)

type mockReader struct {
	refCount           int64
	numRecords         int64
	currentRecordIndex int64
	record             arrow.Record
	schema             *arrow.Schema
}

var (
	// https://arrow.apache.org/docs/format/CDataInterface.html
	availableTypes = map[string]arrow.Field{
		"int8":                  {Name: "int8", Type: arrow.PrimitiveTypes.Int8},
		"uint8":                 {Name: "uint8", Type: arrow.PrimitiveTypes.Uint8},
		"int16":                 {Name: "int16", Type: arrow.PrimitiveTypes.Int16},
		"uint16":                {Name: "uint16", Type: arrow.PrimitiveTypes.Uint16},
		"int32":                 {Name: "int32", Type: arrow.PrimitiveTypes.Int32},
		"uint32":                {Name: "uint32", Type: arrow.PrimitiveTypes.Uint32},
		"int64":                 {Name: "int64", Type: arrow.PrimitiveTypes.Int64},
		"uint64":                {Name: "uint64", Type: arrow.PrimitiveTypes.Uint64},
		"float16":               {Name: "float16", Type: arrow.FixedWidthTypes.Float16},
		"float32":               {Name: "float32", Type: arrow.PrimitiveTypes.Float32},
		"float64":               {Name: "float64", Type: arrow.PrimitiveTypes.Float64},
		"binary":                {Name: "binary", Type: arrow.BinaryTypes.Binary},
		"string":                {Name: "string", Type: arrow.BinaryTypes.String},
		"date32":                {Name: "date32", Type: arrow.PrimitiveTypes.Date32},
		"date64":                {Name: "date64", Type: arrow.PrimitiveTypes.Date64},
		"bool":                  {Name: "bool", Type: arrow.FixedWidthTypes.Boolean},
		"time32s":               {Name: "time32s", Type: arrow.FixedWidthTypes.Time32s},
		"time32ms":              {Name: "time32ms", Type: arrow.FixedWidthTypes.Time32ms},
		"time64us":              {Name: "time64us", Type: arrow.FixedWidthTypes.Time64us},
		"time64ns":              {Name: "time64ns", Type: arrow.FixedWidthTypes.Time64ns},
		"timestamp_s":           {Name: "timestamp_s", Type: arrow.FixedWidthTypes.Timestamp_s},
		"timestamp_ms":          {Name: "timestamp_ms", Type: arrow.FixedWidthTypes.Timestamp_ms},
		"timestamp_us":          {Name: "timestamp_us", Type: arrow.FixedWidthTypes.Timestamp_us},
		"timestamp_ns":          {Name: "timestamp_ns", Type: arrow.FixedWidthTypes.Timestamp_ns},
		"duration_s":            {Name: "duration_s", Type: arrow.FixedWidthTypes.Duration_s},
		"duration_ms":           {Name: "duration_ms", Type: arrow.FixedWidthTypes.Duration_ms},
		"duration_us":           {Name: "duration_us", Type: arrow.FixedWidthTypes.Duration_us},
		"duration_ns":           {Name: "duration_ns", Type: arrow.FixedWidthTypes.Duration_ns},
		"interval_month":        {Name: "interval_month", Type: arrow.FixedWidthTypes.MonthInterval},
		"interval_daytime":      {Name: "interval_daytime", Type: arrow.FixedWidthTypes.DayTimeInterval},
		"interval_monthdaynano": {Name: "interval_monthdaynano", Type: arrow.FixedWidthTypes.MonthDayNanoInterval},
		"null": {Name: "null", Type: arrow.Null},
	}
)

type QueryListener struct {
	*parser.BaseQueryLanguageListener
	lexer *parser.QueryLanguageLexer
	typeStack []arrow.DataType
	fields []*arrow.Field
	rows      int
	nameIdCounter int
}

func (l *QueryListener) ExitStructField(ctx *parser.StructFieldContext) {
	// Currently is the same as fields
	log.Printf("Exiting struct fields: %s", ctx.GetText())
	fieldName := "struct#" + strconv.Itoa(l.nameIdCounter)
	l.nameIdCounter++
	fieldNameNode := ctx.FIELD_NAME()

	if fieldNameNode != nil {
		fieldName = fieldNameNode.GetText()[1:]
	}

	l.fields = append(l.fields, &arrow.Field{
		Name: fieldName,
		Type: l.typeStack[len(l.typeStack)-1],
	})
	log.Printf("Created struct field: %v", l.fields[len(l.fields)-1])

	l.typeStack = l.typeStack[:len(l.typeStack)-1]
}

func (l *QueryListener) ExitTopLevelField(ctx *parser.TopLevelFieldContext) {
	log.Printf("Exiting fields: %s", ctx.GetText())
	fieldName := l.typeStack[len(l.typeStack)-1].Name() + "#" + strconv.Itoa(l.nameIdCounter)
	l.nameIdCounter++
	fieldNameNode := ctx.FIELD_NAME()

	if fieldNameNode != nil {
		fieldName = fieldNameNode.GetText()[1:]
	}

	l.fields = append(l.fields, &arrow.Field{
		Name: fieldName,
		Type: l.typeStack[len(l.typeStack)-1],
	})
	log.Printf("Created field: %v", l.fields[len(l.fields)-1])

	l.typeStack = l.typeStack[:len(l.typeStack)-1]
}

func (l *QueryListener) EnterSimpleTypes(ctx *parser.SimpleTypesContext) {
	tokenType := ctx.GetStart().GetTokenType();
	typeTokenName := l.lexer.GetSymbolicNames()[tokenType]
	typeName := strings.ToLower(typeTokenName);
	log.Printf("Entering simple type: %s", typeName)
	if dataType, ok := availableTypes[typeName]; ok {
		log.Printf("Adding data type: %v", dataType.Type)
		l.typeStack = append(l.typeStack, dataType.Type)
	} else {
		panic(fmt.Sprintf("Unknown data type: %s", typeName))
	}
}

func (l *QueryListener) ExitFixedSizeBinary(ctx *parser.FixedSizeBinaryContext) {
	byteWidthStr := ctx.BYTE_WIDTH()
	if byteWidth, err := strconv.Atoi(byteWidthStr.GetText()); err != nil {
		panic(fmt.Sprintf("Invalid byte width: %s, check parser", byteWidthStr.GetText()))
	} else {
		l.typeStack = append(l.typeStack, &arrow.FixedSizeBinaryType{ByteWidth: byteWidth})
	}
}

func parsePercisionAndScale(scaleStr string) (int32, int32) {
	split := strings.Split(scaleStr, ":")
	precision, err := strconv.Atoi(split[0])
	if err != nil {
		panic(fmt.Sprintf("Invalid precision: %s, check parser", split[0]))
	}
	scale, err := strconv.Atoi(split[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid scale: %s, check parser", split[1]))
	}
	return int32(precision), int32(scale)
}

func (l *QueryListener) ExitDecimal128(ctx *parser.Decimal128Context) {
	precisionScaleStr := ctx.DECIMAL_PS()
	precision, scale := parsePercisionAndScale(precisionScaleStr.GetText())
	// TODO check if presicion and scale are valid
	l.typeStack = append(l.typeStack, &arrow.Decimal128Type{Precision: precision, Scale: scale})
}

func (l *QueryListener) ExitDecimal256(ctx *parser.Decimal256Context) {
	precisionScaleStr := ctx.DECIMAL_PS()
	precision, scale := parsePercisionAndScale(precisionScaleStr.GetText())
	// TODO check if presicion and scale are valid
	l.typeStack = append(l.typeStack, &arrow.Decimal256Type{Precision: precision, Scale: scale})
}

func (l *QueryListener) EnterList(ctx *parser.ListContext) {
	log.Printf("Entering list")
}

func (l *QueryListener) EnterStruct(ctx *parser.StructContext) {
	log.Printf("Entering struct")
	l.fields = append(l.fields, nil)
}

func (l *QueryListener) ExitList(ctx *parser.ListContext) {
	innerType := l.typeStack[len(l.typeStack)-1]
	log.Printf("Exiting list, Inner type: %v", innerType)
	l.typeStack = l.typeStack[:len(l.typeStack)-1]

	listLength := 1
	if ctx.COUNT() != nil {
		rowCountStr := strings.Split(ctx.COUNT().GetText(),":")[0]
		rowCount, err := strconv.Atoi(rowCountStr)
		if err != nil {
			panic(fmt.Sprintf("Invalid row count in list: %s, check parser", rowCountStr))
		}
		listLength = rowCount
	}

	thisList := arrow.ListOf(innerType)
	md := arrow.NewMetadata([]string{"list_length"}, []string{fmt.Sprintf("%d", listLength)})
	thisList.SetElemMetadata(md)
	l.typeStack = append(l.typeStack, thisList)
	log.Printf("Added list: %v", thisList)
}

func (l *QueryListener) ExitStruct(ctx *parser.StructContext) {
	log.Printf("Exiting struct")
	structFields := make([]arrow.Field, 0);
	// struct can have multiple fields, so we keep popping until reaching struct start(nil)
	for l.fields[len(l.fields)-1] != nil {
		structFields = append(structFields, *l.fields[len(l.fields)-1])
		l.fields = l.fields[:len(l.fields)-1]
	}
	l.fields = l.fields[:len(l.fields)-1]

	//reverse the structFields since we added them in reverse order
	for i, j := 0, len(structFields)-1; i < j; i, j = i+1, j-1 {
		structFields[i], structFields[j] = structFields[j], structFields[i]
	}

	thisStruct := arrow.StructOf(structFields...)
	l.typeStack = append(l.typeStack, thisStruct)
	log.Printf("Added struct: %v", thisStruct)
}

func (l* QueryListener)ExitQuery(ctx *parser.QueryContext) {
	// Typestack should be empty now if everything went well
	if len(l.typeStack) != 0 {
		panic(fmt.Sprintf("Typestack not empty: %v", l.typeStack))
	}

	log.Printf("Query finished, types: %v", l.typeStack)
	rowCountNode := ctx.COUNT()
	if rowCountNode != nil {
		rowCountStr := strings.Split(rowCountNode.GetText(),":")[0]
		rowCount, err := strconv.Atoi(rowCountStr)
		if err != nil {
			panic(fmt.Sprintf("Invalid row count in query: %s, check parser", rowCountStr))
		}
		l.rows = rowCount
	} else {
		l.rows = 1
	}
	log.Printf("Query finished, rows: %d", l.rows)
}

// NewMockReader Creates a mockReader according to the query.
// The query should be a list of types separated by commas
// The returned mockReader will have the types in the same order
func NewMockReader(query string) (*mockReader,int64, error) {
	log.Printf("Creating mock reader for query: %s", query)

	inputStream := antlr.NewInputStream(query)
	lexer := parser.NewQueryLanguageLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewQueryLanguageParser(tokenStream)

	listener := &QueryListener{nameIdCounter: 0,lexer: lexer}
	log.Println("Walking query tree")

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Query())

	if len(listener.fields) == 0 {
		return nil, -1, fmt.Errorf("the query does not contain any fields")
	}

	log.Printf("Parsed query into types: %v", listener.typeStack)

	fieldsCopy := make([]arrow.Field, len(listener.fields))
	for i, field := range listener.fields {
		fieldsCopy[i] = *field 
	}
	schema := arrow.NewSchema(fieldsCopy, nil)

	log.Printf("Created schema: %v", schema)

	mockData := PopulateSchema(schema, listener.rows)

	log.Printf("Populated schema with data: %v", mockData)

	return &mockReader{
		refCount:           1,
		numRecords:         1,
		currentRecordIndex: 0,
		schema:             schema,
		record:             mockData,
	}, int64(listener.rows),nil
}

func (r *mockReader) Retain() {
	atomic.AddInt64(&r.refCount, 1)
}

func (r *mockReader) Release() {
	if atomic.AddInt64(&r.refCount, -1) == 0 {
		if r.record != nil {
			r.record.Release()
		}
	}
}

func (r *mockReader) Schema() *arrow.Schema {
	return r.schema
}

func (r *mockReader) Record() arrow.Record {
	return r.record
}

func (r *mockReader) Err() error {
	return nil
}

func (r *mockReader) Next() bool {
	hasNext := r.currentRecordIndex < r.numRecords
	r.currentRecordIndex += 1
	return hasNext
}
