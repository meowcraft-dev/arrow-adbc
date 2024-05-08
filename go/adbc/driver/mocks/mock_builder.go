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
	"math"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/float16"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"golang.org/x/exp/constraints"
)

func mockStruct(mem memory.Allocator, rows int, structType *arrow.StructType) arrow.Array {
	structBuilder := array.NewStructBuilder(mem, structType)
	defer structBuilder.Release()
	structBuilder.Reserve(structType.NumFields())
	for i := 0; i < structType.NumFields(); i++ {
		fieldBuilder := structBuilder.FieldBuilder(i)
		fieldBuilder.Resize(rows)
		valid := make([]bool, rows)
		for i := 0; i < rows; i++ {
			valid[i] = true
		}
		structBuilder.AppendValues(valid)
		switch fieldBuilder.Type() {
		case arrow.PrimitiveTypes.Int8:
			fillValue(fieldBuilder.(*array.Int8Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int16:
			fillValue(fieldBuilder.(*array.Int16Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int32:
			fillValue(fieldBuilder.(*array.Int32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int64:
			fillValue(fieldBuilder.(*array.Int64Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint8:
			fillValue(fieldBuilder.(*array.Uint8Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint16:
			fillValue(fieldBuilder.(*array.Uint16Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint32:
			fillValue(fieldBuilder.(*array.Uint32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint64:
			fillValue(fieldBuilder.(*array.Uint64Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Float32:
			fillValue(fieldBuilder.(*array.Float32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Float64:
			fillValue(fieldBuilder.(*array.Float64Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Date32, arrow.FixedWidthTypes.Date32:
			fillValue(fieldBuilder.(*array.Date32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Date64, arrow.FixedWidthTypes.Date64:
			fillValue(fieldBuilder.(*array.Date64Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Time32s:
			fillValue(fieldBuilder.(*array.Time32Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Time32ms:
			fillValue(fieldBuilder.(*array.Time32Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Time64us:
			fillValue(fieldBuilder.(*array.Time64Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Time64ns:
			fillValue(fieldBuilder.(*array.Time64Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Timestamp_s:
			fillValue(fieldBuilder.(*array.TimestampBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Timestamp_ms:
			fillValue(fieldBuilder.(*array.TimestampBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Timestamp_us:
			fillValue(fieldBuilder.(*array.TimestampBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Timestamp_ns:
			fillValue(fieldBuilder.(*array.TimestampBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Duration_s:
			fillValue(fieldBuilder.(*array.DurationBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Duration_ms:
			fillValue(fieldBuilder.(*array.DurationBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Duration_us:
			fillValue(fieldBuilder.(*array.DurationBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Duration_ns:
			fillValue(fieldBuilder.(*array.DurationBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.MonthInterval:
			fillValue(fieldBuilder.(*array.MonthIntervalBuilder).AppendValues, rows, 0)

		case arrow.FixedWidthTypes.Boolean:
			fillBoolValue(fieldBuilder.(*array.BooleanBuilder).AppendValues, rows)
		case arrow.FixedWidthTypes.DayTimeInterval:
			fillDayTimeIntervalValue(fieldBuilder.(*array.DayTimeIntervalBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.MonthDayNanoInterval:
			fillMonthDayNanoIntervalValue(fieldBuilder.(*array.MonthDayNanoIntervalBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Float16:
			fillFloat16Value(fieldBuilder.(*array.Float16Builder).AppendValues, rows, 0)
		case arrow.BinaryTypes.Binary:
			fillBinaryValue(fieldBuilder.(*array.BinaryBuilder).AppendValues, rows, 0)
		case arrow.BinaryTypes.String:
			fillStringValue(fieldBuilder.(*array.StringBuilder).AppendValues, rows, 0)
		}
	}
	return structBuilder.NewArray()
}

func mockArray(mem memory.Allocator, rows int, elemType arrow.DataType) arrow.Array {
	listBuilder := array.NewListBuilder(mem, elemType)
	defer listBuilder.Release()
	listBuilder.Append(true)

	switch elemType {
	case arrow.PrimitiveTypes.Int8:
		fillValue(listBuilder.ValueBuilder().(*array.Int8Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int16:
		fillValue(listBuilder.ValueBuilder().(*array.Int16Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int32:
		fillValue(listBuilder.ValueBuilder().(*array.Int32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int64:
		fillValue(listBuilder.ValueBuilder().(*array.Int64Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint8:
		fillValue(listBuilder.ValueBuilder().(*array.Uint8Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint16:
		fillValue(listBuilder.ValueBuilder().(*array.Uint16Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint32:
		fillValue(listBuilder.ValueBuilder().(*array.Uint32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint64:
		fillValue(listBuilder.ValueBuilder().(*array.Uint64Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Float32:
		fillValue(listBuilder.ValueBuilder().(*array.Float32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Float64:
		fillValue(listBuilder.ValueBuilder().(*array.Float64Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Date32, arrow.FixedWidthTypes.Date32:
		fillValue(listBuilder.ValueBuilder().(*array.Date32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Date64, arrow.FixedWidthTypes.Date64:
		fillValue(listBuilder.ValueBuilder().(*array.Date64Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Time32s:
		fillValue(listBuilder.ValueBuilder().(*array.Time32Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Time32ms:
		fillValue(listBuilder.ValueBuilder().(*array.Time32Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Time64us:
		fillValue(listBuilder.ValueBuilder().(*array.Time64Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Time64ns:
		fillValue(listBuilder.ValueBuilder().(*array.Time64Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Timestamp_s:
		fillValue(listBuilder.ValueBuilder().(*array.TimestampBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Timestamp_ms:
		fillValue(listBuilder.ValueBuilder().(*array.TimestampBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Timestamp_us:
		fillValue(listBuilder.ValueBuilder().(*array.TimestampBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Timestamp_ns:
		fillValue(listBuilder.ValueBuilder().(*array.TimestampBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Duration_s:
		fillValue(listBuilder.ValueBuilder().(*array.DurationBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Duration_ms:
		fillValue(listBuilder.ValueBuilder().(*array.DurationBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Duration_us:
		fillValue(listBuilder.ValueBuilder().(*array.DurationBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Duration_ns:
		fillValue(listBuilder.ValueBuilder().(*array.DurationBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.MonthInterval:
		fillValue(listBuilder.ValueBuilder().(*array.MonthIntervalBuilder).AppendValues, rows, 0)

	case arrow.FixedWidthTypes.Boolean:
		fillBoolValue(listBuilder.ValueBuilder().(*array.BooleanBuilder).AppendValues, rows)
	case arrow.FixedWidthTypes.DayTimeInterval:
		fillDayTimeIntervalValue(listBuilder.ValueBuilder().(*array.DayTimeIntervalBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.MonthDayNanoInterval:
		fillMonthDayNanoIntervalValue(listBuilder.ValueBuilder().(*array.MonthDayNanoIntervalBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Float16:
		fillFloat16Value(listBuilder.ValueBuilder().(*array.Float16Builder).AppendValues, rows, 0)
	case arrow.BinaryTypes.Binary:
		fillBinaryValue(listBuilder.ValueBuilder().(*array.BinaryBuilder).AppendValues, rows, 0)
	case arrow.BinaryTypes.String:
		fillStringValue(listBuilder.ValueBuilder().(*array.StringBuilder).AppendValues, rows, 0)
	}

	return listBuilder.NewArray()
}

func fillValue[T constraints.Integer | constraints.Float](append func(value []T, valid []bool), rows int, start int) {
	append(getSlice(rows, T(start)), nil)
}

func fillBoolValue(append func(value []bool, valid []bool), rows int) {
	append(getBoolSlice(rows), nil)
}

func fillDayTimeIntervalValue(append func(value []arrow.DayTimeInterval, valid []bool), rows int, start int) {
	append(getDayTimeIntervalSlice(rows, start), nil)
}

func fillMonthDayNanoIntervalValue(append func(value []arrow.MonthDayNanoInterval, valid []bool), rows int, start int) {
	append(getMonthDayNanoIntervalSlice(rows, start), nil)
}

func fillFloat16Value(append func(value []float16.Num, valid []bool), rows int, start int) {
	append(getFloat16Slice(rows, float32(start)), nil)
}

func fillBinaryValue(append func(value [][]byte, valid []bool), rows int, start int) {
	append(getBinarySlice(rows, start), nil)
}

func fillStringValue(append func(value []string, valid []bool), rows int, start int) {
	append(getStringSlice(rows, start), nil)
}

func getSlice[T constraints.Integer | constraints.Float](rows int, start T) []T {
	slice := make([]T, rows)
	for i := int64(0); i < int64(rows); i++ {
		slice[i] = T(i) + start
	}
	return slice
}

func getBoolSlice(rows int) []bool {
	slice := make([]bool, rows)
	for i := 0; i < rows; i++ {
		slice[i] = i%2 == 0
	}
	return slice
}

func getDayTimeIntervalSlice(rows int, start int) []arrow.DayTimeInterval {
	slice := make([]arrow.DayTimeInterval, rows)
	for i := start; i < start+rows; i++ {
		slice[i] = arrow.DayTimeInterval{Days: int32(i), Milliseconds: int32(i)}
	}
	return slice
}

func getMonthDayNanoIntervalSlice(rows int, start int) []arrow.MonthDayNanoInterval {
	slice := make([]arrow.MonthDayNanoInterval, rows)
	for i := start; i < start+rows; i++ {
		slice[i] = arrow.MonthDayNanoInterval{Months: int32(i), Days: int32(i), Nanoseconds: int64(i)}
	}
	return slice
}

func getFloat16Slice(rows int, start float32) []float16.Num {
	slice := make([]float16.Num, rows)
	for i := 0; i < rows; i++ {
		slice[i] = float16.New(start + float32(i))
	}
	return slice
}

func getBinarySlice(rows int, start int) [][]byte {
	slice := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		slice[i] = []byte{byte(i + start), byte(i + start + 1), byte(i + start + 2), byte(i + start + 3)}
	}
	return slice
}

var stringData = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getStringSlice(rows int, start int) []string {
	slice := make([]string, rows)
	for i := 0; i < rows; i++ {
		pos := (i + start) % len(stringData)
		slice[i] = stringData[pos : pos+10]
	}
	return slice
}

func mockInt8(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt8Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt8Array()
}

func mockInt16(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt16Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt16Array()
}

func mockInt32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt32Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt32Array()
}

func mockInt64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt64Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt64Array()
}

func mockUint8(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint8Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint8Array()
}

func mockUint16(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint16Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint16Array()
}

func mockUint32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint32Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint32Array()
}

func mockUint64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint64Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint64Array()
}

func mockFloat32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewFloat32Builder(mem)
	defer ib.Release()
	values := make([]float32, 0, rows)
	if rows > 3 {
		values = append(values, math.SmallestNonzeroFloat32, -1, 0, math.MaxFloat32)
	}
	for i := len(values); i < rows; i++ {
		values = append(values, float32(i))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat32Array()
}

func mockFloat64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewFloat64Builder(mem)
	defer ib.Release()
	values := make([]float64, 0, rows)
	if rows > 3 {
		values = append(values, math.SmallestNonzeroFloat64, -1, 0, math.MaxFloat64)
	}
	for i := len(values); i < rows; i++ {
		values = append(values, float64(i))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat64Array()
}

func mockDate32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDate32Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDate32Array()
}

func mockDate64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDate64Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDate64Array()
}

func mockBinary(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewBinaryBuilder(mem, arrow.BinaryTypes.Binary)
	defer ib.Release()
	fillBinaryValue(ib.AppendValues, rows, 0)
	return ib.NewBinaryArray()
}

func mockString(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewStringBuilder(mem)
	defer ib.Release()
	fillStringValue(ib.AppendValues, rows, 0)
	return ib.NewStringArray()
}

func mockDayTimeInterval(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDayTimeIntervalBuilder(mem)
	defer ib.Release()
	fillDayTimeIntervalValue(ib.AppendValues, rows, 0)
	return ib.NewDayTimeIntervalArray()
}

func mockDuration_s(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Second})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDurationArray()
}

func mockDuration_ms(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Millisecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDurationArray()
}

func mockDuration_us(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Microsecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDurationArray()
}

func mockDuration_ns(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDurationBuilder(mem, &arrow.DurationType{Unit: arrow.Nanosecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewDurationArray()
}

func mockFloat16(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewFloat16Builder(mem)
	defer ib.Release()
	values := make([]float16.Num, rows)
	for i := 0; i < rows; i++ {
		values[i] = float16.New(float32(i + 1))
	}
	ib.AppendValues(values, nil)
	return ib.NewFloat16Array()
}

func mockMonthInterval(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewMonthIntervalBuilder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewMonthIntervalArray()
}

func mockTime32s(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Second})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTime32Array()
}

func mockTime32ms(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTime32Builder(mem, &arrow.Time32Type{Unit: arrow.Millisecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTime32Array()
}

func mockTime64us(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Microsecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTime64Array()
}

func mockTime64ns(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTime64Builder(mem, &arrow.Time64Type{Unit: arrow.Nanosecond})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTime64Array()
}

func mockTimestamp_s(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Second, TimeZone: "UTC"})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTimestampArray()
}

func mockTimestamp_ms(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: "UTC"})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTimestampArray()
}

func mockTimestamp_us(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTimestampArray()
}

func mockTimestamp_ns(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewTimestampBuilder(mem, &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: "UTC"})
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewTimestampArray()
}

// func mockList(mem memory.Allocator, rows, length int64, innerList arrow.Array) arrow.Array {
// 	lb := array.NewListBuilder(mem, arrow.ListOf(innerList.DataType()))
// 	defer lb.Release()
// 	// vb := lb.ValueBuilder().(*array.ListBuilder).Append()
// }
