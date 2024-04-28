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
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/float16"
	"github.com/apache/arrow/go/v16/arrow/memory"
)

type mockReader struct {
	haveNext bool
	record   arrow.Record
}

func (r *mockReader) Retain() {}
func (r *mockReader) Release() {
	if r.record != nil {
		r.record.Release()
	}
}
func (r *mockReader) Schema() *arrow.Schema {
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "int8", Type: arrow.PrimitiveTypes.Int8},
		{Name: "int16", Type: arrow.PrimitiveTypes.Int16},
		{Name: "int32", Type: arrow.PrimitiveTypes.Int32},
		{Name: "int64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "uint8", Type: arrow.PrimitiveTypes.Uint8},
		{Name: "uint16", Type: arrow.PrimitiveTypes.Uint16},
		{Name: "uint32", Type: arrow.PrimitiveTypes.Uint32},
		{Name: "uint64", Type: arrow.PrimitiveTypes.Uint64},
		{Name: "float32", Type: arrow.PrimitiveTypes.Float32},
		{Name: "float64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "date32", Type: arrow.PrimitiveTypes.Date32},
		{Name: "date64", Type: arrow.PrimitiveTypes.Date64},
		{Name: "binary", Type: arrow.BinaryTypes.Binary},
		{Name: "string", Type: arrow.BinaryTypes.String},
		{Name: "daytimeinterval", Type: arrow.FixedWidthTypes.DayTimeInterval},
		{Name: "duration_s", Type: arrow.FixedWidthTypes.Duration_s},
		{Name: "duration_ms", Type: arrow.FixedWidthTypes.Duration_ms},
		{Name: "duration_us", Type: arrow.FixedWidthTypes.Duration_us},
		{Name: "duration_ns", Type: arrow.FixedWidthTypes.Duration_ns},
		{Name: "float16", Type: arrow.FixedWidthTypes.Float16},
		{Name: "monthInterval", Type: arrow.FixedWidthTypes.MonthInterval},
		{Name: "time32s", Type: arrow.FixedWidthTypes.Time32s},
		{Name: "time32ms", Type: arrow.FixedWidthTypes.Time32ms},
		{Name: "time64us", Type: arrow.FixedWidthTypes.Time64us},
		{Name: "time64ns", Type: arrow.FixedWidthTypes.Time64ns},
		{Name: "timestamp_s", Type: arrow.FixedWidthTypes.Timestamp_s},
		{Name: "timestamp_ms", Type: arrow.FixedWidthTypes.Timestamp_ms},
		{Name: "timestamp_us", Type: arrow.FixedWidthTypes.Timestamp_us},
		{Name: "timestamp_ns", Type: arrow.FixedWidthTypes.Timestamp_ns},
	}, nil)
	return schema
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
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Second,TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Second),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Second),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ms(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Millisecond,TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Millisecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Millisecond),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_us(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Microsecond,TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Microsecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Microsecond),
	}, nil)
	return ib.NewTimestampArray()
}

func mockTimestamp_ns(mem memory.Allocator) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Nanosecond,TimeZone: "UTC"})
	defer ib.Release()
	ib.AppendValues([]arrow.Timestamp{tsFromStr("1970-01-01 00:00:01Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:02Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:03Z", arrow.Nanosecond),
		tsFromStr("1970-01-01 00:00:04Z", arrow.Nanosecond),
	}, nil)
	return ib.NewTimestampArray()
}

func (r *mockReader) Record() arrow.Record {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	cols := []arrow.Array{mockInt8(mem), mockInt16(mem), mockInt32(mem), mockInt64(mem),
		mockUint8(mem), mockUint16(mem), mockUint32(mem), mockUint64(mem),
		mockFloat32(mem), mockFloat64(mem),
		mockDate32(mem), mockDate64(mem),
		mockBinary(mem), mockString(mem),
		mockDayTimeInterval(mem),
		mockDuration_s(mem), mockDuration_ms(mem), mockDuration_us(mem), mockDuration_ns(mem),
		mockFloat16(mem),
		mockMonthInterval(mem),
		mockTime32s(mem), mockTime32ms(mem),
		mockTime64us(mem), mockTime64ns(mem),
		mockTimestamp_s(mem), mockTimestamp_ms(mem), mockTimestamp_us(mem), mockTimestamp_ns(mem),}
	rec := array.NewRecord(r.Schema(), cols, 4)
	rec.Retain()
	r.record = rec
	return rec
}
func (r *mockReader) Err() error {
	return nil
}