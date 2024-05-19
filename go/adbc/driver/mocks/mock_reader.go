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
	rowsRegex      = regexp.MustCompile(`^(?P<rows>\d+)[\s]*:`)
	listRegex      = regexp.MustCompile(`^list<(?P<len>\d*):?(?P<typename>.+)>`)
	structRegex    = regexp.MustCompile(`^struct\<(?P<struct>.+)\>`)
	availableTypes = map[string]typeBuilder{
		"int8": {
			field:   arrow.Field{Name: "int8", Type: arrow.PrimitiveTypes.Int8},
			builder: mockInt8,
		},
		"int16": {
			field:   arrow.Field{Name: "int16", Type: arrow.PrimitiveTypes.Int16},
			builder: mockInt16,
		},
		"int32": {
			field:   arrow.Field{Name: "int32", Type: arrow.PrimitiveTypes.Int32},
			builder: mockInt32,
		},
		"int64": {
			field:   arrow.Field{Name: "int64", Type: arrow.PrimitiveTypes.Int64},
			builder: mockInt64,
		},
		"uint8": {
			field:   arrow.Field{Name: "uint8", Type: arrow.PrimitiveTypes.Uint8},
			builder: mockUint8,
		},
		"uint16": {
			field:   arrow.Field{Name: "uint16", Type: arrow.PrimitiveTypes.Uint16},
			builder: mockUint16,
		},
		"uint32": {
			field:   arrow.Field{Name: "uint32", Type: arrow.PrimitiveTypes.Uint32},
			builder: mockUint32,
		},
		"uint64": {
			field:   arrow.Field{Name: "uint64", Type: arrow.PrimitiveTypes.Uint64},
			builder: mockUint64,
		},
		"float32": {
			field:   arrow.Field{Name: "float32", Type: arrow.PrimitiveTypes.Float32},
			builder: mockFloat32,
		},
		"float64": {
			field:   arrow.Field{Name: "float64", Type: arrow.PrimitiveTypes.Float64},
			builder: mockFloat64,
		},
		"date32": {
			field:   arrow.Field{Name: "date32", Type: arrow.PrimitiveTypes.Date32},
			builder: mockDate32,
		},
		"date64": {
			field:   arrow.Field{Name: "date64", Type: arrow.PrimitiveTypes.Date64},
			builder: mockDate64,
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
		"float16": {
			field:   arrow.Field{Name: "float16", Type: arrow.FixedWidthTypes.Float16},
			builder: mockFloat16,
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
			field:   arrow.Field{Name: "interval_monthdaynano", Type: arrow.FixedWidthTypes.MonthInterval},
			builder: mockMonthDayNanoInterval,
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
	}
)

func parseQuery(query string, innerRows int) ([]arrow.Field, []arrow.Array, int, error) {
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
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
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
			if t == "list" {
				fields = append(fields, arrow.Field{Name: "list", Type: arrow.ListOf(arrow.PrimitiveTypes.Int8)})
				fieldValues = append(fieldValues, mockArray(mem, rows, arrow.PrimitiveTypes.Int8))
			} else if listMatch := listRegex.FindStringSubmatch(t); len(listMatch) > 1 {
				listLen := 1
				if lenStr := listMatch[listRegex.SubexpIndex("len")]; len(lenStr) > 0 {
					listLen, _ = strconv.Atoi(lenStr)
				}
				elmTypeStr := listMatch[listRegex.SubexpIndex("typename")]
				if innerType, ok := availableTypes[elmTypeStr]; ok {
					fields = append(fields, arrow.Field{Name: "list", Type: arrow.ListOf(innerType.field.Type)})
					fieldValues = append(fieldValues, mockArray(mem, listLen, innerType.field.Type))
				} else {
					// array.NewListBuilder(mem,)
					// innerFields, innerValues, _, err := parseQuery(elmTypeStr,listLen)
					// if err != nil {
					// 	return nil, nil, -1, err
					// }
					// fields = append(fields, arrow.Field{Name: "list", Type: arrow.ListOf(innerFields[0].Type)})
					// fieldValues = append(fieldValues, mockList(innerValues[0]))
					return nil, nil, -1, adbc.Error{
						Code: adbc.StatusInvalidArgument,
						Msg:  fmt.Sprintf("unknown type %s", t),
					}
				}
			} else {
				return nil, nil, -1, adbc.Error{
					Code: adbc.StatusInvalidArgument,
					Msg:  fmt.Sprintf("unknown type %s", t),
				}
			}
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
