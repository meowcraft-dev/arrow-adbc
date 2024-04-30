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
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/float16"
	"github.com/apache/arrow/go/v16/arrow/memory"
)

type mockReader struct {
	refCount int64
	haveNext bool
	record   arrow.Record
	schema   *arrow.Schema
}

type typeBuilder struct {
	field   arrow.Field
	builder func(memory.Allocator,int64) arrow.Array
}

var (
	rowsRegex      = regexp.MustCompile(`^(?P<rows>\d+)[\s]*:`)
	availableTypes = map[string]typeBuilder{
		"int8": typeBuilder{
			field:   arrow.Field{Name: "int8", Type: arrow.PrimitiveTypes.Int8},
			builder: mockInt8,
		},
		"int16": typeBuilder{
			field:   arrow.Field{Name: "int16", Type: arrow.PrimitiveTypes.Int16},
			builder: mockInt16,
		},
		"int32": typeBuilder{
			field:   arrow.Field{Name: "int32", Type: arrow.PrimitiveTypes.Int32},
			builder: mockInt32,
		},
		"int64": typeBuilder{
			field:   arrow.Field{Name: "int64", Type: arrow.PrimitiveTypes.Int64},
			builder: mockInt64,
		},
		"uint8": typeBuilder{
			field:   arrow.Field{Name: "uint8", Type: arrow.PrimitiveTypes.Uint8},
			builder: mockUint8,
		},
		"uint16": typeBuilder{
			field:   arrow.Field{Name: "uint16", Type: arrow.PrimitiveTypes.Uint16},
			builder: mockUint16,
		},
		"uint32": typeBuilder{
			field:   arrow.Field{Name: "uint32", Type: arrow.PrimitiveTypes.Uint32},
			builder: mockUint32,
		},
		"uint64": typeBuilder{
			field:   arrow.Field{Name: "uint64", Type: arrow.PrimitiveTypes.Uint64},
			builder: mockUint64,
		},
		"float32": typeBuilder{
			field:   arrow.Field{Name: "float32", Type: arrow.PrimitiveTypes.Float32},
			builder: mockFloat32,
		},
		"float64": typeBuilder{
			field:   arrow.Field{Name: "float64", Type: arrow.PrimitiveTypes.Float64},
			builder: mockFloat64,
		},
		"date32": typeBuilder{
			field:   arrow.Field{Name: "date32", Type: arrow.PrimitiveTypes.Date32},
			builder: mockDate32,
		},
		"date64": typeBuilder{
			field:   arrow.Field{Name: "date64", Type: arrow.PrimitiveTypes.Date64},
			builder: mockDate64,
		},
		"binary": typeBuilder{
			field:   arrow.Field{Name: "binary", Type: arrow.BinaryTypes.Binary},
			builder: mockBinary,
		},
		"string": typeBuilder{
			field:   arrow.Field{Name: "string", Type: arrow.BinaryTypes.String},
			builder: mockString,
		},
		"daytimeinterval": typeBuilder{
			field:   arrow.Field{Name: "daytimeinterval", Type: arrow.FixedWidthTypes.DayTimeInterval},
			builder: mockDayTimeInterval,
		},
		"duration_s": typeBuilder{
			field:   arrow.Field{Name: "duration_s", Type: arrow.FixedWidthTypes.Duration_s},
			builder: mockDuration_s,
		},
		"duration_ms": typeBuilder{
			field:   arrow.Field{Name: "duration_ms", Type: arrow.FixedWidthTypes.Duration_ms},
			builder: mockDuration_ms,
		},
		"duration_us": typeBuilder{
			field:   arrow.Field{Name: "duration_us", Type: arrow.FixedWidthTypes.Duration_us},
			builder: mockDuration_us,
		},
		"duration_ns": typeBuilder{
			field:   arrow.Field{Name: "duration_ns", Type: arrow.FixedWidthTypes.Duration_ns},
			builder: mockDuration_ns,
		},
		"float16": typeBuilder{
			field:   arrow.Field{Name: "float16", Type: arrow.FixedWidthTypes.Float16},
			builder: mockFloat16,
		},
		"monthInterval": typeBuilder{
			field:   arrow.Field{Name: "monthInterval", Type: arrow.FixedWidthTypes.MonthInterval},
			builder: mockMonthInterval,
		},
		"time32s": typeBuilder{
			field:   arrow.Field{Name: "time32s", Type: arrow.FixedWidthTypes.Time32s},
			builder: mockTime32s,
		},
		"time32ms": typeBuilder{
			field:   arrow.Field{Name: "time32ms", Type: arrow.FixedWidthTypes.Time32ms},
			builder: mockTime32ms,
		},
		"time64us": typeBuilder{
			field:   arrow.Field{Name: "time64us", Type: arrow.FixedWidthTypes.Time64us},
			builder: mockTime64us,
		},
		"time64ns": typeBuilder{
			field:   arrow.Field{Name: "time64ns", Type: arrow.FixedWidthTypes.Time64ns},
			builder: mockTime64ns,
		},
		"timestamp_s": typeBuilder{
			field:   arrow.Field{Name: "timestamp_s", Type: arrow.FixedWidthTypes.Timestamp_s},
			builder: mockTimestamp_s,
		},
		"timestamp_ms": typeBuilder{
			field:   arrow.Field{Name: "timestamp_ms", Type: arrow.FixedWidthTypes.Timestamp_ms},
			builder: mockTimestamp_ms,
		},
		"timestamp_us": typeBuilder{
			field:   arrow.Field{Name: "timestamp_us", Type: arrow.FixedWidthTypes.Timestamp_us},
			builder: mockTimestamp_us,
		},
		"timestamp_ns": typeBuilder{
			field:   arrow.Field{Name: "timestamp_ns", Type: arrow.FixedWidthTypes.Timestamp_ns},
			builder: mockTimestamp_ns,
		},
	}
)

func parseQuery(query string) ([]arrow.Field, []arrow.Array, int64, error) {
	query = strings.TrimSpace(query)
	matches := rowsRegex.FindStringSubmatch(query)
	rows := int64(1)
	rowsString := ""
	if len(matches) == 2 {
		var err error
		rowsString = matches[rowsRegex.SubexpIndex("rows")]
		if rows, err = strconv.ParseInt(rowsString, 10, 64); err != nil {
			return nil, nil, -1, err
		}
	}
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	fields := make([]arrow.Field, 0)
	fieldValues := make([]arrow.Array, 0)

	query = strings.TrimSpace(query[len(rowsString):])
	if query[0] == ':' {
		query = query[1:]
	}

	var t string
	for {
		var requestTypes []string
		if strings.HasPrefix(query, "list") {
			// TODO: check if this is a valid list type
			// query = strings.TrimSpace(query[4:])
			// if query[0] == '<' && query[len(query)-1] == '>' {
			// 	query = query[1 : len(query)-1]
			// }
			
			break
		}
		requestTypes = strings.SplitN(query, ",", 2)

		if len(requestTypes) == 2 {
			t = strings.TrimSpace(requestTypes[0])
			query = strings.TrimSpace(requestTypes[1])
		} else if len(requestTypes) == 1 {
			t = strings.TrimSpace(requestTypes[0])
			query = ""
		}
		if builder, ok := availableTypes[t]; ok {
			fields = append(fields, builder.field)
			fieldValues = append(fieldValues, builder.builder(mem, rows))
		} else if t == "list" {
			fields = append(fields, arrow.Field{Name: "list", Type: arrow.ListOf(arrow.PrimitiveTypes.Int8)})
			fieldValues = append(fieldValues, mockListInt8(mem))
		} else {
			return nil, nil, -1, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("unknown type %s", t),
			}
		}

		if len(query) == 0 {
			break
		}
	}

	return fields, fieldValues, rows, nil
}

// Create a mockReader according to the query
// The query should be a list of types separated by commas
// The returned mockReader will have the types in the same order
func NewMockReader(query string) (*mockReader, error) {
	fields, fieldValues, rows, err := parseQuery(query)
	if err != nil {
		return nil, err
	}
	schema := arrow.NewSchema(fields, nil)
	rec := array.NewRecord(schema, fieldValues, rows)
	rec.Retain()
	return &mockReader{
		refCount: 1,
		haveNext: true,
		schema:   schema,
		record:   rec,
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
	if r.haveNext {
		r.haveNext = false
		return true
	}
	return false
}

func mockListInt8(mem memory.Allocator) arrow.Array {
	listBuilder := array.NewListBuilder(mem, arrow.PrimitiveTypes.Int8)
	defer listBuilder.Release()

	int8Builder := listBuilder.ValueBuilder().(*array.Int8Builder)

	// Create a list of int8 values. For example, create 3 lists with varying lengths.
	lists := [][]int8{
		{1, 2, 3},
		{4, 5},
		{6},
		{},
	}

	for _, vals := range lists {
		listBuilder.Append(true) // Append true to indicate a valid list item
		for _, val := range vals {
			int8Builder.Append(val)
		}
	}

	return listBuilder.NewArray()
}

func mockInt8(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewInt8Builder(mem)
	defer ib.Release()
	values := make([]int8, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = int8(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewInt8Array()
}

func mockInt16(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewInt16Builder(mem)
	defer ib.Release()
	values := make([]int16, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = int16(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewInt16Array()
}


func mockInt32(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewInt32Builder(mem)
	defer ib.Release()
	values := make([]int32, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = int32(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewInt32Array()
}

func mockInt64(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewInt64Builder(mem)
	defer ib.Release()
	values := make([]int64, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = int64(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewInt64Array()
}

func mockUint8(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewUint8Builder(mem)
	defer ib.Release()
	values := make([]uint8, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = uint8(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewUint8Array()
}

func mockUint16(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewUint16Builder(mem)
	defer ib.Release()
	values := make([]uint16, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = uint16(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewUint16Array()
}

func mockUint32(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewUint32Builder(mem)
	defer ib.Release()
	values := make([]uint32, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = uint32(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewUint32Array()
}

func mockUint64(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewUint64Builder(mem)
	defer ib.Release()
	values := make([]uint64, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = uint64(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewUint64Array()
}

func mockFloat32(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewFloat32Builder(mem)
	defer ib.Release()
	values := make([]float32, 0, rows)
	if rows > 3 {
		values = append(values, math.SmallestNonzeroFloat32, -1, 0, math.MaxFloat32)
	}
	for i := int64(len(values)); i < rows; i++ {
		values = append(values, float32(i))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat32Array()
}

func mockFloat64(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewFloat64Builder(mem)
	defer ib.Release()
	values := make([]float64, 0, rows)
	if rows > 3 {
		values = append(values, math.SmallestNonzeroFloat64, -1, 0, math.MaxFloat64)
	}
	for i := int64(len(values)); i < rows; i++ {
		values = append(values, float64(i))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat64Array()
}

func mockDate32(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDate32Builder(mem)
	defer ib.Release()
	values := make([]arrow.Date32, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Date32(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDate32Array()
}

func mockDate64(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDate64Builder(mem)
	defer ib.Release()
	values := make([]arrow.Date64, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Date64(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDate64Array()
}

func mockBinary(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewBinaryBuilder(mem, arrow.BinaryTypes.Binary)
	defer ib.Release()
	values := make([][]byte, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = []byte{byte(i)}
	}
	ib.AppendValues(values, nil)
	return ib.NewBinaryArray()
}

func mockString(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewStringBuilder(mem)
	defer ib.Release()
	values := make([]string, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = string(65 + int(i) % 26)
	}
	ib.AppendValues(values, nil)
	return ib.NewStringArray()
}


func mockDayTimeInterval(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDayTimeIntervalBuilder(mem)
	defer ib.Release()
	values := make([]arrow.DayTimeInterval, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.DayTimeInterval{Days: int32(i), Milliseconds: int32(i + 1)}
	}
	ib.AppendValues(values, nil)
	return ib.NewDayTimeIntervalArray()
}

func mockDuration_s(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Second})
	defer ib.Release()
	values := make([]arrow.Duration, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Duration(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDurationArray()
}

func mockDuration_ms(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Millisecond})
	defer ib.Release()
	values := make([]arrow.Duration, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Duration(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDurationArray()
}

func mockDuration_us(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Microsecond})
	defer ib.Release()
	values := make([]arrow.Duration, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Duration(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDurationArray()
}

func mockDuration_ns(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Nanosecond})
	defer ib.Release()
	values := make([]arrow.Duration, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Duration(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewDurationArray()
}

func mockFloat16(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewFloat16Builder(mem)
	defer ib.Release()
	values := make([]float16.Num, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = float16.New(float32(i + 1))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat16Array()
}

func mockMonthInterval(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewMonthIntervalBuilder(mem)
	defer ib.Release()
	values := make([]arrow.MonthInterval, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.MonthInterval(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewMonthIntervalArray()
}

func mockTime32s(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Second})
	defer ib.Release()
	values := make([]arrow.Time32, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Time32(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewTime32Array()
}

func mockTime32ms(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Millisecond})
	defer ib.Release()
	values := make([]arrow.Time32, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Time32(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewTime32Array()
}

func mockTime64us(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Microsecond})
	defer ib.Release()
	values := make([]arrow.Time64, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Time64(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewTime64Array()
}

func mockTime64ns(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Nanosecond})
	defer ib.Release()
	values := make([]arrow.Time64, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Time64(i + 1)
	}
	ib.AppendValues(values, nil)
	return ib.NewTime64Array()
}


func mockTimestamp_s(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Second, TimeZone: "UTC"})
	defer ib.Release()
	values := make([]arrow.Timestamp, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Timestamp(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ms(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: "UTC"})
	defer ib.Release()
	values := make([]arrow.Timestamp, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Timestamp(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_us(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	defer ib.Release()
	values := make([]arrow.Timestamp, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Timestamp(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ns(mem memory.Allocator, rows int64) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: "UTC"})
	defer ib.Release()
	values := make([]arrow.Timestamp, rows)
	for i := int64(0); i < rows; i++ {
		values[i] = arrow.Timestamp(i)
	}
	ib.AppendValues(values, nil)
	return ib.NewTimestampArray()
}
