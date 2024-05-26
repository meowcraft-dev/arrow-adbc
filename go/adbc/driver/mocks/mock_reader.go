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
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
)

type mockReader struct {
	refCount   int64
	rows       int64
	currentRow int64
	record     arrow.Record
	schema     *arrow.Schema
}

type typeBuilder struct {
	field   arrow.Field
	builder func(memory.Allocator, int) arrow.Array
}

var (
	rowsRegex = regexp.MustCompile(`^(?P<rows>\d+)[\s]*:`)
	//listRegex   = regexp.MustCompile(`^list[\s]*<[\s]*(?P<len>\d*):?(?P<typename>.+)>`)
	structRegex = regexp.MustCompile(`^struct\<(?P<struct>.+)\>`)

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
	}
)

func parseListType(query string, mem memory.Allocator) (arrow.DataType, arrow.Array, int, error) {
	rows := 1
	matches := rowsRegex.FindStringSubmatch(query)
	rowsString := ""
	if len(matches) == 2 {
		var err error
		rowsString = matches[rowsRegex.SubexpIndex("rows")]
		if rows, err = strconv.Atoi(rowsString); err != nil {
			return nil, nil, -1, err
		}
	}
	query = query[len(rowsString):]
	if query[0] == ':' {
		query = query[1:]
	}

	var t string
	requestTypes := strings.SplitN(query, ",", 2)

	if len(requestTypes) == 2 {
		t = requestTypes[0]
		query = requestTypes[1]
	} else if len(requestTypes) == 1 {
		t = requestTypes[0]
		query = ""
	}

	if builder, ok := availableTypes[t]; ok {
		return builder.field.Type, builder.builder(mem, rows), rows, nil
	} else if strings.HasPrefix(t, "list") {
		if len(t) > 4 {
			expected := 0
			startIndex := -1
			endIndex := -1
			for i := 4; i < len(t); i++ {
				if t[i] == '<' {
					if startIndex == -1 {
						startIndex = i
					}
					expected += 1
				} else if t[i] == '>' {
					expected -= 1
					if expected == 0 {
						endIndex = i
						break
					}
				}
			}

			if expected != 0 {
				return nil, nil, rows, adbc.Error{
					Code: adbc.StatusInvalidArgument,
					Msg:  fmt.Sprintf("unmatched brackets in query: %s", query),
				}
			}

			if startIndex != -1 && endIndex != -1 {

			}
			fmt.Printf("list inner type: %s\n", t[startIndex:endIndex])
			innerType, innerValue, _, err := parseListType(t[5:len(t)-1], mem)
			if err != nil {
				return nil, nil, rows, err
			}
			return arrow.ListOf(innerType), innerValue, rows, nil
		} else {
			return nil, nil, rows, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("invalid type: `%s`", t),
			}
		}
	} else {
		return nil, nil, rows, adbc.Error{
			Code: adbc.StatusInvalidArgument,
			Msg:  fmt.Sprintf("type `%s` not supported yet", t),
		}
	}
}

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

	typeEnd := -1
	for j := index + 1; j < queryLen; j++ {
		if (query[j] >= '0' && query[j] <= '9') || (query[j] >= 'a' && query[j] <= 'z') || (query[j] >= 'A' && query[j] <= 'Z') {
			continue
		} else {
			typeEnd = j
			break
		}
	}

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
		fmt.Printf("[debug:parseQueryToFields] currentType=%v, err=%v\n", currentType, err)
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
				fmt.Printf("[debug:expectOneComma] queryLen=%d, typeEnd=%d", queryLen, typeEnd)
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
	currentTypes, currentFields, currentArrays, rowsForType, typeEnd, err := parseQueryToFields(mem, query, len(query), 0, Start, 1, true)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, nil, 0, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("eof"),
			}
		}
		return nil, nil, 0, adbc.Error{}
	}
	fmt.Printf("[parseQueryToFields] query: %s\n", query)
	fmt.Printf("currentTypes: %v\n", currentTypes)
	fmt.Printf("currentFields: %v\n", currentFields)
	fmt.Printf("currentArrays: %v\n", currentArrays)
	fmt.Printf("rowsForType: %v\n", rowsForType)
	fmt.Printf("typeEnd: %v\n", typeEnd)

	return currentFields, currentArrays, int(rowsForType[0]), nil

	matches := rowsRegex.FindStringSubmatch(query)
	rows := innerRows
	rowsString := ""
	if len(matches) == 2 {
		var err error
		rowsString = matches[rowsRegex.SubexpIndex("rows")]
		if rows, err = strconv.Atoi(rowsString); err != nil {
			return nil, nil, -1, err
		}
	}

	fields := make([]arrow.Field, 0)
	fieldValues := make([]arrow.Array, 0)

	query = query[len(rowsString):]
	if query[0] == ':' {
		query = query[1:]
	}

	var t string
	for {
		requestTypes := strings.SplitN(query, ",", 2)

		if len(requestTypes) == 2 {
			t = requestTypes[0]
			query = requestTypes[1]
		} else if len(requestTypes) == 1 {
			t = requestTypes[0]
			query = ""
		}
		if builder, ok := availableTypes[t]; ok {
			fields = append(fields, builder.field)
			fieldValues = append(fieldValues, builder.builder(mem, rows))
		} else if strings.HasPrefix(t, "list") {
			listType, innerValue, _, err := parseListType(t, mem)
			if err != nil {
				return nil, nil, -1, err
			}
			fields = append(fields, arrow.Field{Name: "list", Type: listType})
			fieldValues = append(fieldValues, innerValue)
		} else if strings.HasPrefix(t, "struct") {
			fmt.Println("oh it's a struct!")
			structMatch := structRegex.FindStringSubmatch(t)
			fmt.Printf("structMatch: %v\n", structMatch)
			if structMatch := structRegex.FindStringSubmatch(t); structMatch != nil {
				structString := strings.Split(structMatch[structRegex.SubexpIndex("struct")], ",")
				fmt.Printf("structString: %v\n", structString)
				innerTypeFields := make([]arrow.Field, 0)
				for _, innerType := range structString {
					innerTypeFields = append(innerTypeFields, availableTypes[innerType].field)
				}
				fields = append(fields, arrow.Field{Name: "struct", Type: arrow.StructOf(innerTypeFields...)})
				fieldValues = append(fieldValues, mockStruct(mem, rows, arrow.StructOf(innerTypeFields...)))
			} else {
				return nil, nil, -1, adbc.Error{
					Code: adbc.StatusInvalidArgument,
					Msg:  fmt.Sprintf("unknown type %s", t),
				}
			}
		}
		if len(query) == 0 {
			break
		}
	}

	return fields, fieldValues, rows, nil
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
		refCount:   1,
		rows:       1,
		currentRow: 0,
		schema:     schema,
		record:     rec,
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
	hasNext := r.currentRow < r.rows
	r.currentRow += 1
	return hasNext
}
