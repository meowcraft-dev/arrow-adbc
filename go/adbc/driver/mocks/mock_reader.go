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
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
	"github.com/apache/arrow-adbc/go/adbc/driver/mocks/parser"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/memory"
)

type mockReader struct {
	refCount           int64
	numRecords         int64
	currentRecordIndex int64
	record             arrow.Record
	schema             *arrow.Schema
}

type typeBuilder struct {
	field   arrow.Field
	builder func(memory.Allocator, int) arrow.Array
}

var (
	// https://arrow.apache.org/docs/format/CDataInterface.html
	availableTypes = map[string]typeBuilder{
		"int8": {
			field:   arrow.Field{Name: "int8", Type: arrow.PrimitiveTypes.Int8},
			builder: mockInt8,
		},
		"uint8": {
			field:   arrow.Field{Name: "uint8", Type: arrow.PrimitiveTypes.Uint8},
			builder: mockUint8,
		},
		"int16": {
			field:   arrow.Field{Name: "int16", Type: arrow.PrimitiveTypes.Int16},
			builder: mockInt16,
		},
		"uint16": {
			field:   arrow.Field{Name: "uint16", Type: arrow.PrimitiveTypes.Uint16},
			builder: mockUint16,
		},
		"int32": {
			field:   arrow.Field{Name: "int32", Type: arrow.PrimitiveTypes.Int32},
			builder: mockInt32,
		},
		"uint32": {
			field:   arrow.Field{Name: "uint32", Type: arrow.PrimitiveTypes.Uint32},
			builder: mockUint32,
		},
		"int64": {
			field:   arrow.Field{Name: "int64", Type: arrow.PrimitiveTypes.Int64},
			builder: mockInt64,
		},
		"uint64": {
			field:   arrow.Field{Name: "uint64", Type: arrow.PrimitiveTypes.Uint64},
			builder: mockUint64,
		},
		"float16": {
			field:   arrow.Field{Name: "float16", Type: arrow.FixedWidthTypes.Float16},
			builder: mockFloat16,
		},
		"float32": {
			field:   arrow.Field{Name: "float32", Type: arrow.PrimitiveTypes.Float32},
			builder: mockFloat32,
		},
		"float64": {
			field:   arrow.Field{Name: "float64", Type: arrow.PrimitiveTypes.Float64},
			builder: mockFloat64,
		},
		"binary": {
			field:   arrow.Field{Name: "binary", Type: arrow.BinaryTypes.Binary},
			builder: mockBinary,
		},
		"string": {
			field:   arrow.Field{Name: "string", Type: arrow.BinaryTypes.String},
			builder: mockString,
		},
		"fixed_size_binary": {
			field:   arrow.Field{Name: "fixed_size_binary", Type: &arrow.FixedSizeBinaryType{ByteWidth: 5}},
			builder: mockFixedSizeBinary,
		},
		"date32": {
			field:   arrow.Field{Name: "date32", Type: arrow.PrimitiveTypes.Date32},
			builder: mockDate32,
		},
		"date64": {
			field:   arrow.Field{Name: "date64", Type: arrow.PrimitiveTypes.Date64},
			builder: mockDate64,
		},
		"bool": {
			field:   arrow.Field{Name: "bool", Type: arrow.FixedWidthTypes.Boolean},
			builder: mockBool,
		},
		"time32s": {
			field:   arrow.Field{Name: "time32s", Type: arrow.FixedWidthTypes.Time32s},
			builder: mockTime32s,
		},
		"time32ms": {
			field:   arrow.Field{Name: "time32ms", Type: arrow.FixedWidthTypes.Time32ms},
			builder: mockTime32ms,
		},
		"time64us": {
			field:   arrow.Field{Name: "time64us", Type: arrow.FixedWidthTypes.Time64us},
			builder: mockTime64us,
		},
		"time64ns": {
			field:   arrow.Field{Name: "time64ns", Type: arrow.FixedWidthTypes.Time64ns},
			builder: mockTime64ns,
		},
		"timestamp_s": {
			field:   arrow.Field{Name: "timestamp_s", Type: arrow.FixedWidthTypes.Timestamp_s},
			builder: mockTimestamp_s,
		},
		"timestamp_ms": {
			field:   arrow.Field{Name: "timestamp_ms", Type: arrow.FixedWidthTypes.Timestamp_ms},
			builder: mockTimestamp_ms,
		},
		"timestamp_us": {
			field:   arrow.Field{Name: "timestamp_us", Type: arrow.FixedWidthTypes.Timestamp_us},
			builder: mockTimestamp_us,
		},
		"timestamp_ns": {
			field:   arrow.Field{Name: "timestamp_ns", Type: arrow.FixedWidthTypes.Timestamp_ns},
			builder: mockTimestamp_ns,
		},
		"duration_s": {
			field:   arrow.Field{Name: "duration_s", Type: arrow.FixedWidthTypes.Duration_s},
			builder: mockDuration_s,
		},
		"duration_ms": {
			field:   arrow.Field{Name: "duration_ms", Type: arrow.FixedWidthTypes.Duration_ms},
			builder: mockDuration_ms,
		},
		"duration_us": {
			field:   arrow.Field{Name: "duration_us", Type: arrow.FixedWidthTypes.Duration_us},
			builder: mockDuration_us,
		},
		"duration_ns": {
			field:   arrow.Field{Name: "duration_ns", Type: arrow.FixedWidthTypes.Duration_ns},
			builder: mockDuration_ns,
		},
		"interval_month": {
			field:   arrow.Field{Name: "interval_month", Type: arrow.FixedWidthTypes.MonthInterval},
			builder: mockMonthInterval,
		},
		"interval_daytime": {
			field:   arrow.Field{Name: "interval_daytime", Type: arrow.FixedWidthTypes.DayTimeInterval},
			builder: mockDayTimeInterval,
		},
		"interval_monthdaynano": {
			field:   arrow.Field{Name: "interval_monthdaynano", Type: arrow.FixedWidthTypes.MonthDayNanoInterval},
			builder: mockMonthDayNanoInterval,
		},
		"decimal128": {
			field:   arrow.Field{Name: "decimal128", Type: &arrow.Decimal128Type{Precision: 37, Scale: 2}},
			builder: mockDecimal128,
		},
		"decimal256": {
			field:   arrow.Field{Name: "decimal256", Type: &arrow.Decimal256Type{Precision: 76, Scale: 4}},
			builder: mockDecimal256,
		},
		"sample_list": {
			field:   arrow.Field{Name: "sample_list", Type: arrow.ListOf(arrow.PrimitiveTypes.Int32)},
			builder: mockSampleList,
		},
		"sample_list_with_struct": {
			field: arrow.Field{Name: "sample_list_with_struct", Type: arrow.ListOf(arrow.StructOf(
				arrow.Field{Name: "start_time", Type: arrow.FixedWidthTypes.Timestamp_s},
				arrow.Field{Name: "end_time", Type: arrow.FixedWidthTypes.Timestamp_s},
				arrow.Field{Name: "data_points", Type: arrow.PrimitiveTypes.Int32},
			))},
			builder: mockSampleListWithStruct,
		},
		"sample_nested_list": {
			field:   arrow.Field{Name: "sample_nested_list", Type: arrow.ListOf(arrow.ListOf(arrow.PrimitiveTypes.Int32))},
			builder: mockSampleNestedList,
		},
		"sample_list_view": {
			field:   arrow.Field{Name: "sample_list_view", Type: arrow.ListViewOf(arrow.PrimitiveTypes.Int32)},
			builder: mockSampleListView,
		},
		"sample_large_list_view": {
			field: arrow.Field{
				Name: "sample_large_list_view",
				Type: arrow.LargeListViewOf(arrow.PrimitiveTypes.Int32),
			},
			builder: mockSampleLargeListView,
		},
		"sample_fixed_size_list": {
			field: arrow.Field{
				Name: "sample_fixed_size_list",
				Type: arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32),
			},
			builder: mockSampleFixedSizeList,
		},
		"sample_nested_fixed_size_list": {
			field: arrow.Field{
				Name: "sample_nested_fixed_size_list",
				Type: arrow.FixedSizeListOf(3, arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32)),
			},
			builder: mockSampleNestedFixedSizeList,
		},
		"sample_run_end_encoded_array": {
			field: arrow.Field{
				Name: "sample_run_end_encoded_array",
				Type: arrow.RunEndEncodedOf(arrow.PrimitiveTypes.Int32, arrow.PrimitiveTypes.Float32),
			},
			builder: mockSampleRunEndEncodedArray,
		},
		"sample_dictionary_encoded_array": {
			field: arrow.Field{
				Name: "sample_dictionary_encoded_array",
				Type: &arrow.DictionaryType{
					IndexType: arrow.PrimitiveTypes.Int32,
					ValueType: arrow.BinaryTypes.String,
				},
			},
			builder: mockSampleDictionaryEncodedArray,
		},
		"sample_dense_union": {
			field: arrow.Field{
				Name: "sample_dense_union",
				Type: arrow.DenseUnionOf(
					[]arrow.Field{
						{Name: "a", Type: arrow.PrimitiveTypes.Int32},
						{Name: "b", Type: arrow.BinaryTypes.String},
					},
					[]arrow.UnionTypeCode{0, 1},
				),
			},
			builder: mockSampleDenseUnion,
		},
		"sample_sparse_union": {
			field: arrow.Field{
				Name: "sample_sparse_union",
				Type: arrow.SparseUnionOf(
					[]arrow.Field{
						{Name: "a", Type: arrow.PrimitiveTypes.Int32},
						{Name: "b", Type: arrow.BinaryTypes.String},
					},
					[]arrow.UnionTypeCode{0, 1},
				),
			},
			builder: mockSampleSparseUnion,
		},
		"null": {
			field:   arrow.Field{Name: "null", Type: arrow.Null},
			builder: mockNull,
		},
	}
)

type QueryListener struct {
	*parser.BaseQueryLanguageListener
	typeStack []arrow.DataType
}

func (l *QueryListener)EnterPrimitiveType(ctx *parser.PrimitiveTypeContext) {
	typeName := ctx.GetText()
	if dataType, ok := availableTypes[typeName]; ok {
		l.typeStack = append(l.typeStack, dataType.field.Type)
	}
}

func (l *QueryListener)ExitList(ctx *parser.ListContext) {
	innerType := l.typeStack[len(l.typeStack)-1]
	l.typeStack = l.typeStack[:len(l.typeStack)-1]

	thisList := arrow.ListOf(innerType)
	l.typeStack = append(l.typeStack, thisList)
}

// NewMockReader Creates a mockReader according to the query.
// The query should be a list of types separated by commas
// The returned mockReader will have the types in the same order
func NewMockReader(query string) (*mockReader, error) {
	inputStream := antlr.NewInputStream(query)
	lexer := parser.NewQueryLanguageLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewQueryLanguageParser(tokenStream)

	listener := &QueryListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Query())

	fields := make([]arrow.Field, len(listener.typeStack))

	for i, t := range listener.typeStack {
		fields[i] = arrow.Field{
			Name: fmt.Sprintf("field_%d", i),
			Type: t,
		}
	}
	schema := arrow.NewSchema(fields, nil)

	var mockData arrow.Record
	// TODO mockData := populate(schema)

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
