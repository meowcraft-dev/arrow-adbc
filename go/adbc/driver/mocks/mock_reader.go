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
	builder func(memory.Allocator) arrow.Array
}

var (
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

// Create a mockReader according to the query
// The query should be a list of types separated by commas
// The returned mockReader will have the types in the same order
func NewMockReader(query string) (*mockReader, error) {
	requestTypes := strings.Split(query, ",")
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	fields := make([]arrow.Field, 0)
	fieldValues := make([]arrow.Array, 0)
	for _, t := range requestTypes {
		t = strings.TrimSpace(t)
		if builder, ok := availableTypes[t]; ok {
			fields = append(fields, builder.field)
			fieldValues = append(fieldValues, builder.builder(mem))
		} else {
			return nil, adbc.Error{
				Code: adbc.StatusInvalidArgument,
				Msg:  fmt.Sprintf("unknown type %s", t),
			}
		}
	}

	schema := arrow.NewSchema(fields, nil)
	rec := array.NewRecord(schema, fieldValues, 4)
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

func mockInt8(mem memory.Allocator) arrow.Array {
	ib := array.NewInt8Builder(mem)
	defer ib.Release()
	ib.AppendValues([]int8{1, 2, 3, 4}, nil)
	return ib.NewInt8Array()
}

func mockInt16(mem memory.Allocator) arrow.Array {
	ib := array.NewInt16Builder(mem)
	defer ib.Release()
	ib.AppendValues([]int16{1, 2, 3, 4}, nil)
	return ib.NewInt16Array()
}

func mockInt32(mem memory.Allocator) arrow.Array {
	ib := array.NewInt32Builder(mem)
	defer ib.Release()
	ib.AppendValues([]int32{1, 2, 3, 4}, nil)
	return ib.NewInt32Array()
}

func mockInt64(mem memory.Allocator) arrow.Array {
	ib := array.NewInt64Builder(mem)
	defer ib.Release()
	ib.AppendValues([]int64{1, 2, 3, 4}, nil)
	return ib.NewInt64Array()
}

func mockUint8(mem memory.Allocator) arrow.Array {
	ib := array.NewUint8Builder(mem)
	defer ib.Release()
	ib.AppendValues([]uint8{1, 2, 3, 4}, nil)
	return ib.NewUint8Array()
}

func mockUint16(mem memory.Allocator) arrow.Array {
	ib := array.NewUint16Builder(mem)
	defer ib.Release()
	ib.AppendValues([]uint16{1, 2, 3, 4}, nil)
	return ib.NewUint16Array()
}

func mockUint32(mem memory.Allocator) arrow.Array {
	ib := array.NewUint32Builder(mem)
	defer ib.Release()
	ib.AppendValues([]uint32{1, 2, 3, 4}, nil)
	return ib.NewUint32Array()
}

func mockUint64(mem memory.Allocator) arrow.Array {
	ib := array.NewUint64Builder(mem)
	defer ib.Release()
	ib.AppendValues([]uint64{1, 2, 3, 4}, nil)
	return ib.NewUint64Array()
}

func mockFloat32(mem memory.Allocator) arrow.Array {
	ib := array.NewFloat32Builder(mem)
	defer ib.Release()
	ib.AppendValues([]float32{1, 2, 3, 4}, nil)
	return ib.NewFloat32Array()
}

func mockFloat64(mem memory.Allocator) arrow.Array {
	ib := array.NewFloat64Builder(mem)
	defer ib.Release()
	ib.AppendValues([]float64{1, 2, 3, 4}, nil)
	return ib.NewFloat64Array()
}

func mockDate32(mem memory.Allocator) arrow.Array {
	ib := array.NewDate32Builder(mem)
	defer ib.Release()
	ib.AppendValues([]arrow.Date32{1, 2, 3, 4}, nil)
	return ib.NewDate32Array()
}

func mockDate64(mem memory.Allocator) arrow.Array {
	ib := array.NewDate64Builder(mem)
	defer ib.Release()
	ib.AppendValues([]arrow.Date64{1, 2, 3, 4}, nil)
	return ib.NewDate64Array()
}

func mockBinary(mem memory.Allocator) arrow.Array {
	ib := array.NewBinaryBuilder(mem, arrow.BinaryTypes.Binary)
	defer ib.Release()
	ib.AppendValues([][]byte{{1}, {2}, {3}, {4}}, nil)
	return ib.NewBinaryArray()
}

func mockString(mem memory.Allocator) arrow.Array {
	ib := array.NewStringBuilder(mem)
	defer ib.Release()
	ib.AppendValues([]string{"a", "b", "c", "d"}, nil)
	return ib.NewStringArray()
}

func mockDayTimeInterval(mem memory.Allocator) arrow.Array {
	ib := array.NewDayTimeIntervalBuilder(mem)
	defer ib.Release()
	ib.AppendValues([]arrow.DayTimeInterval{{Days: 1, Milliseconds: 2}, {Days: 3, Milliseconds: 4}, {Days: 5, Milliseconds: 6}, {Days: 7, Milliseconds: 8}}, nil)
	return ib.NewDayTimeIntervalArray()
}

func mockDuration_s(mem memory.Allocator) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Second})
	defer ib.Release()
	ib.AppendValues([]arrow.Duration{1, 2, 3, 4}, nil)
	return ib.NewDurationArray()
}

func mockDuration_ms(mem memory.Allocator) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Millisecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Duration{1, 2, 3, 4}, nil)
	return ib.NewDurationArray()
}

func mockDuration_us(mem memory.Allocator) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Microsecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Duration{1, 2, 3, 4}, nil)
	return ib.NewDurationArray()
}

func mockDuration_ns(mem memory.Allocator) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Nanosecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Duration{1, 2, 3, 4}, nil)
	return ib.NewDurationArray()
}

func mockFloat16(mem memory.Allocator) arrow.Array {
	ib := array.NewFloat16Builder(mem)
	defer ib.Release()
	ib.AppendValues([]float16.Num{float16.New(1), float16.New(2), float16.New(3), float16.New(4)}, nil)
	return ib.NewFloat16Array()
}

func mockMonthInterval(mem memory.Allocator) arrow.Array {
	ib := array.NewMonthIntervalBuilder(mem)
	defer ib.Release()
	ib.AppendValues([]arrow.MonthInterval{1, 2, 3, 4}, nil)
	return ib.NewMonthIntervalArray()
}

func mockTime32s(mem memory.Allocator) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Second})
	defer ib.Release()
	ib.AppendValues([]arrow.Time32{1, 2, 3, 4}, nil)
	return ib.NewTime32Array()
}

func mockTime32ms(mem memory.Allocator) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Millisecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Time32{1, 2, 3, 4}, nil)
	return ib.NewTime32Array()
}

func mockTime64us(mem memory.Allocator) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Microsecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Time64{1, 2, 3, 4}, nil)
	return ib.NewTime64Array()
}

func mockTime64ns(mem memory.Allocator) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Nanosecond})
	defer ib.Release()
	ib.AppendValues([]arrow.Time64{1, 2, 3, 4}, nil)
	return ib.NewTime64Array()
}

func tsFromStr(str string, unit arrow.TimeUnit) arrow.Timestamp {
	ts, _ := arrow.TimestampFromString(str, unit)
	return ts
}

func mockTimestamp_s(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Second, TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Second),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ms(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Millisecond),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_us(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Microsecond),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ns(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Nanosecond),
	}, nil)
	return ib.NewTimestampArray()
}
