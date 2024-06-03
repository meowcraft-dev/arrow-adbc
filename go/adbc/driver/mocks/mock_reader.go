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
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync/atomic"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
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
				Type: arrow.FixedSizeListOf(3,arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32)),
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
	}
)

type ParseStatus int

const (
	Start ParseStatus = iota
	ExpectType
	ENd
)

func parseRows(query string, queryLen, index int) (uint64, int, error) {
	rowEnd := -1
	for j := index + 1; j < queryLen; j++ {
		if query[j] >= '0' && query[j] <= '9' {
			continue
		} else {
			rowEnd = j
			break
		}
	}
	parsedRows, err := strconv.Atoi(query[index:rowEnd])
	if err != nil {
		return 0, 0, adbc.Error{
			Code: adbc.StatusInvalidArgument,
			Msg:  fmt.Sprintf("invalid query: `%s`: %s", query, err.Error()),
		}
	}
	return uint64(parsedRows), rowEnd, nil
}

func parseList(mem memory.Allocator, query string, queryLen, startIndex int) (arrow.DataType, arrow.Field, arrow.Array, uint64, int, error) {
	leftEnd, err := expectLeftAngleBracket(query, queryLen, startIndex)
	if err != nil {
		return nil, arrow.Field{}, nil, 0, 0, err
	}
	listInnerType, _, innerArray, innerRows, listEnd, err := parseQueryToFields(mem, query, queryLen, leftEnd, Start, 1, false)
	if err != nil {
		return nil, arrow.Field{}, nil, 0, 0, err
	}
	rightEnd, err := expectRightAngleBracket(query, queryLen, listEnd)
	if err != nil {
		return nil, arrow.Field{}, nil, 0, 0, err
	}
	listType := arrow.ListOf(listInnerType[0])

	return listType, arrow.Field{Name: "list", Type: listType}, innerArray[0], innerRows[0], rightEnd, nil
}

func parseType(mem memory.Allocator, query string, queryLen, index int, rows uint64) (arrow.DataType, arrow.Field, arrow.Array, uint64, int, error) {
	if index == queryLen {
		return nil, arrow.Field{}, nil, 0, 0, io.EOF
	}

	typeEnd := -2
	for j := index + 1; j < queryLen; j++ {
		if (query[j] >= '0' && query[j] <= '9') || (query[j] >= 'a' && query[j] <= 'z') || (query[j] >= 'A' && query[j] <= 'Z') || query[j] == '_' {
			typeEnd = j
		} else {
			break
		}
	}
	typeEnd += 1

	if typeEnd == -1 {
		return nil, arrow.Field{}, nil, 0, 0, adbc.Error{
			Code: adbc.StatusInvalidArgument,
			Msg:  fmt.Sprintf("[parseType] invalid query: `%s`", query),
		}
	}

	typeString := query[index:typeEnd]
	if builder, ok := availableTypes[typeString]; ok {
		return builder.field.Type, builder.field, builder.builder(mem, int(rows)), rows, typeEnd, nil
	} else if typeString == "list" {
		return parseList(mem, query, queryLen, typeEnd)
	} else {
		return nil, arrow.Field{}, nil, 0, 0, adbc.Error{
			Code: adbc.StatusInvalidArgument,
			Msg:  fmt.Sprintf("invalid query [parseType]: `%s`", query),
		}
	}
}

func expectLeftAngleBracket(query string, queryLen, index int) (int, error) {
	if index == queryLen {
		return 0, io.EOF
	}

	for i := index; i < queryLen; i++ {
		if query[i] == '<' {
			return i + 1, nil
		} else if query[i] == ' ' || query[i] == '\t' || query[i] == '\n' {
			continue
		} else {
			return 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("expecting a `<` but got: `%c` at %d", query[i], i),
			}
		}
	}
	return queryLen, io.EOF
}

func expectRightAngleBracket(query string, queryLen, index int) (int, error) {
	if index == queryLen {
		return 0, io.EOF
	}

	for i := index; i < queryLen; i++ {
		if query[i] == '>' {
			return i + 1, nil
		} else if query[i] == ' ' || query[i] == '\t' || query[i] == '\n' {
			continue
		} else {
			return 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("expecting a `<` but got: `%c` at %d", query[i], i),
			}
		}
	}
	return queryLen, io.EOF
}

func expectOneComma(query string, queryLen, index int) (int, error) {
	if index == queryLen {
		return 0, io.EOF
	}

	for i := index; i < queryLen; i++ {
		if query[i] == ',' {
			return i + 1, nil
		} else if query[i] == ' ' || query[i] == '\t' || query[i] == '\n' {
			continue
		} else {
			return 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("expecting a `,` but got: `%c` at %d", query[i], i),
			}
		}
	}
	return queryLen, io.EOF
}

func expectOneColumn(query string, queryLen, index int) (int, error) {
	if index == queryLen {
		return 0, io.EOF
	}

	for i := index; i < queryLen; i++ {
		if query[i] == ':' {
			return i + 1, nil
		} else if query[i] == ' ' || query[i] == '\t' || query[i] == '\n' {
			continue
		} else {
			return 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("expecting a `:` but got: `%c` at %d", query[i], i),
			}
		}
	}
	return queryLen, io.EOF
}

func parseQueryToFields(mem memory.Allocator, query string, queryLen, index int, status ParseStatus, rows uint64, allowMultipleTypes bool) ([]arrow.DataType, []arrow.Field, []arrow.Array, []uint64, int, error) {
	switch status {
	case Start:
		for i := index; i < queryLen; i++ {
			if query[i] == ' ' || query[i] == '\t' || query[i] == '\n' {
				continue
			} else if query[i] >= '0' && query[i] <= '9' {
				parsedRows, rowEnd, err := parseRows(query, queryLen, i)
				if err != nil {
					return nil, nil, nil, nil, 0, err
				}

				columnEnd, err := expectOneColumn(query, queryLen, rowEnd)
				if err != nil {
					return nil, nil, nil, nil, 0, err
				}
				return parseQueryToFields(mem, query, queryLen, columnEnd, ExpectType, parsedRows, allowMultipleTypes)
			} else {
				return parseQueryToFields(mem, query, queryLen, i, ExpectType, 1, allowMultipleTypes)
			}
		}
	case ExpectType:
		types := make([]arrow.DataType, 0)
		typeRows := make([]uint64, 0)
		typeFields := make([]arrow.Field, 0)
		typeArrays := make([]arrow.Array, 0)
		currentType, typeField, typeArray, rowsForType, typeEnd, err := parseType(mem, query, queryLen, index, rows)
		if err != nil {
			if errors.Is(err, io.EOF) {
				types = append(types, currentType)
				typeRows = append(typeRows, rowsForType)
				typeFields = append(typeFields, typeField)
				typeArrays = append(typeArrays, typeArray)
				return types, typeFields, typeArrays, typeRows, typeEnd, nil
			}
			return nil, nil, nil, nil, 0, err
		}
		types = append(types, currentType)
		typeRows = append(typeRows, rowsForType)
		typeFields = append(typeFields, typeField)
		typeArrays = append(typeArrays, typeArray)
		if allowMultipleTypes {
			for {
				//fmt.Printf("[debug:expectOneComma] queryLen=%d, typeEnd=%d", queryLen, typeEnd)
				commaEnd, err := expectOneComma(query, queryLen, typeEnd)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return nil, nil, nil, nil, 0, err
				}

				nextType, nextField, nextArray, nextRows, nextEnd, err := parseType(mem, query, queryLen, commaEnd, rows)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return nil, nil, nil, nil, 0, err
				}
				types = append(types, nextType)
				typeRows = append(typeRows, nextRows)
				typeFields = append(typeFields, nextField)
				typeArrays = append(typeArrays, nextArray)
				typeEnd = nextEnd
			}
			return types, typeFields, typeArrays, typeRows, typeEnd, nil
		} else {
			return types, typeFields, typeArrays, typeRows, typeEnd, nil
		}
	default:
		return nil, nil, nil, nil, 0, adbc.Error{
			Code: adbc.StatusInvalidArgument,
			Msg:  fmt.Sprintf("[default] invalid query: `%s`", query),
		}
	}

	return nil, nil, nil, nil, 0, adbc.Error{
		Code: adbc.StatusInvalidArgument,
		Msg:  fmt.Sprintf("[return] invalid query: `%s`", query),
	}
}

func parseQuery(query string, innerRows int) ([]arrow.Field, []arrow.Array, int, error) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	_, currentFields, currentArrays, rowsForType, _, err := parseQueryToFields(mem, query, len(query), 0, Start, 1, true)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, nil, 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("eof"),
			}
		}
		return nil, nil, 0, adbc.Error{}
	}
	//fmt.Printf("[parseQueryToFields] query: %s\n", query)
	//fmt.Printf("currentTypes: %v\n", currentTypes)
	//fmt.Printf("currentFields: %v\n", currentFields)
	//fmt.Printf("currentArrays: %v\n", currentArrays)
	//fmt.Printf("rowsForType: %v\n", rowsForType)
	//fmt.Printf("typeEnd: %v\n", typeEnd)

	return currentFields, currentArrays, int(rowsForType[0]), nil
}

// NewMockReader Creates a mockReader according to the query.
// The query should be a list of types separated by commas
// The returned mockReader will have the types in the same order
func NewMockReader(query string) (*mockReader, error) {
	fields, fieldValues, rows, err := parseQuery(query, 1)
	if err != nil {
		return nil, err
	}
	schema := arrow.NewSchema(fields, nil)
	rec := array.NewRecord(schema, fieldValues, int64(rows))
	rec.Retain()
	return &mockReader{
		refCount:           1,
		numRecords:         1,
		currentRecordIndex: 0,
		schema:             schema,
		record:             rec,
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
