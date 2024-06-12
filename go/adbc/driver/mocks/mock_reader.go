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
		"fixed_size_binary":     {Name: "fixed_size_binary", Type: &arrow.FixedSizeBinaryType{ByteWidth: 5}},
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
		"decimal128":            {Name: "decimal128", Type: &arrow.Decimal128Type{Precision: 37, Scale: 2}},
		"decimal256":            {Name: "decimal256", Type: &arrow.Decimal256Type{Precision: 76, Scale: 4}},
		"sample_list":           {Name: "sample_list", Type: arrow.ListOf(arrow.PrimitiveTypes.Int32)},
		"sample_list_with_struct": {Name: "sample_list_with_struct", Type: arrow.ListOf(arrow.StructOf(
			arrow.Field{Name: "start_time", Type: arrow.FixedWidthTypes.Timestamp_s},
			arrow.Field{Name: "end_time", Type: arrow.FixedWidthTypes.Timestamp_s},
			arrow.Field{Name: "data_points", Type: arrow.PrimitiveTypes.Int32},
		))},
		"sample_nested_list": {Name: "sample_nested_list", Type: arrow.ListOf(arrow.ListOf(arrow.PrimitiveTypes.Int32))},
		"sample_list_view":   {Name: "sample_list_view", Type: arrow.ListViewOf(arrow.PrimitiveTypes.Int32)},
		"sample_large_list_view": {
			Name: "sample_large_list_view",
			Type: arrow.LargeListViewOf(arrow.PrimitiveTypes.Int32),
		},
		"sample_fixed_size_list": {
			Name: "sample_fixed_size_list",
			Type: arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32),
		},
		"sample_nested_fixed_size_list": {
			Name: "sample_nested_fixed_size_list",
			Type: arrow.FixedSizeListOf(3, arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32)),
		},
		"sample_run_end_encoded_array": {
			Name: "sample_run_end_encoded_array",
			Type: arrow.RunEndEncodedOf(arrow.PrimitiveTypes.Int32, arrow.PrimitiveTypes.Float32),
		},
		"sample_dictionary_encoded_array": {
			Name: "sample_dictionary_encoded_array",
			Type: &arrow.DictionaryType{
				IndexType: arrow.PrimitiveTypes.Int32,
				ValueType: arrow.BinaryTypes.String,
			},
		},
		"sample_dense_union": {
			Name: "sample_dense_union",
			Type: arrow.DenseUnionOf(
				[]arrow.Field{
					{Name: "a", Type: arrow.PrimitiveTypes.Int32},
					{Name: "b", Type: arrow.BinaryTypes.String},
				},
				[]arrow.UnionTypeCode{0, 1},
			),
		},
		"sample_sparse_union": {
			Name: "sample_sparse_union",
			Type: arrow.SparseUnionOf(
				[]arrow.Field{
					{Name: "a", Type: arrow.PrimitiveTypes.Int32},
					{Name: "b", Type: arrow.BinaryTypes.String},
				},
				[]arrow.UnionTypeCode{0, 1},
			),
		},
		"null": {Name: "null", Type: arrow.Null},
	}
)

type QueryListener struct {
	*parser.BaseQueryLanguageListener
	typeStack []arrow.DataType
	rows      int
}

func (l *QueryListener) EnterTypeSpec(ctx *parser.TypeSpecContext) {
	log.Printf("TypeSpec: %s", ctx.GetText())
}

func (l *QueryListener) EnterSimpleTypes(ctx *parser.SimpleTypesContext) {
	typeName := ctx.GetText()
	log.Printf("Entering simple type: %s", typeName)
	if dataType, ok := availableTypes[typeName]; ok {
		log.Printf("Adding data type: %v", dataType.Type)
		l.typeStack = append(l.typeStack, dataType.Type)
	} else {
		log.Printf("Unknown data type: %s", typeName)
	}
}

func (l *QueryListener) ExitList(ctx *parser.ListContext) {
	innerType := l.typeStack[len(l.typeStack)-1]
	l.typeStack = l.typeStack[:len(l.typeStack)-1]

	thisList := arrow.ListOf(innerType)
	l.typeStack = append(l.typeStack, thisList)
}

func (l* QueryListener)ExitQuery(ctx *parser.QueryContext) {
	log.Printf("Query finished, types: %v", l.typeStack)
	rowCountNode := ctx.ROWCOUNT()
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
func NewMockReader(query string) (*mockReader, error) {
	log.Printf("Creating mock reader for query: %s", query)

	inputStream := antlr.NewInputStream(query)
	lexer := parser.NewQueryLanguageLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewQueryLanguageParser(tokenStream)

	listener := &QueryListener{}
	log.Println("Walking query tree")

	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Query())

	if len(listener.typeStack) == 0 {
		return nil, fmt.Errorf("no types found in query")
	}

	log.Printf("Parsed query into types: %v", listener.typeStack)

	fields := make([]arrow.Field, len(listener.typeStack))

	for i, t := range listener.typeStack {
		log.Printf("Creating field %d with type %v", i, t)
		fields[i] = arrow.Field{
			Name: fmt.Sprintf("field_%d", i),
			Type: t,
		}
	}
	schema := arrow.NewSchema(fields, nil)

	log.Printf("Created schema: %v", schema)

	mockData := PopulateSchema(schema, listener.rows)

	log.Printf("Populated schema with data: %v", mockData)

	return &mockReader{
		refCount:           1,
		numRecords:         1,
		currentRecordIndex: 0,
		schema:             schema,
		record:             mockData,
	}, nil
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
