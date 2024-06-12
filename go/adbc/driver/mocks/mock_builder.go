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
	"math/big"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/decimal128"
	"github.com/apache/arrow/go/v17/arrow/decimal256"
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
		case arrow.FixedWidthTypes.Boolean:
			fillBoolValue(fieldBuilder.(*array.BooleanBuilder).AppendValues, rows)
		case arrow.PrimitiveTypes.Int8:
			fillValue(fieldBuilder.(*array.Int8Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint8:
			fillValue(fieldBuilder.(*array.Uint8Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int16:
			fillValue(fieldBuilder.(*array.Int16Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint16:
			fillValue(fieldBuilder.(*array.Uint16Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int32:
			fillValue(fieldBuilder.(*array.Int32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint32:
			fillValue(fieldBuilder.(*array.Uint32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Int64:
			fillValue(fieldBuilder.(*array.Int64Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Uint64:
			fillValue(fieldBuilder.(*array.Uint64Builder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.Float16:
			fillFloat16Value(fieldBuilder.(*array.Float16Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Float32:
			fillValue(fieldBuilder.(*array.Float32Builder).AppendValues, rows, 0)
		case arrow.PrimitiveTypes.Float64:
			fillValue(fieldBuilder.(*array.Float64Builder).AppendValues, rows, 0)
		case arrow.BinaryTypes.Binary:
			fillBinaryValue(fieldBuilder.(*array.BinaryBuilder).AppendValues, rows, 0)
		case arrow.BinaryTypes.String:
			fillStringValue(fieldBuilder.(*array.StringBuilder).AppendValues, rows, 0)
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
			fillIntervalMonthValue(fieldBuilder.(*array.MonthIntervalBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.DayTimeInterval:
			fillIntervalDayTimeValue(fieldBuilder.(*array.DayTimeIntervalBuilder).AppendValues, rows, 0)
		case arrow.FixedWidthTypes.MonthDayNanoInterval:
			fillIntervalMonthDayNanoValue(fieldBuilder.(*array.MonthDayNanoIntervalBuilder).AppendValues, rows, 0)
		}
	}
	return structBuilder.NewArray()
}

func mockArray(mem memory.Allocator, rows int, elemType arrow.DataType) arrow.Array {
	listBuilder := array.NewListBuilder(mem, elemType)
	defer listBuilder.Release()
	listBuilder.Append(true)

	switch elemType {
	case arrow.FixedWidthTypes.Boolean:
		fillBoolValue(listBuilder.ValueBuilder().(*array.BooleanBuilder).AppendValues, rows)
	case arrow.PrimitiveTypes.Int8:
		fillValue(listBuilder.ValueBuilder().(*array.Int8Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint8:
		fillValue(listBuilder.ValueBuilder().(*array.Uint8Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int16:
		fillValue(listBuilder.ValueBuilder().(*array.Int16Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint16:
		fillValue(listBuilder.ValueBuilder().(*array.Uint16Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int32:
		fillValue(listBuilder.ValueBuilder().(*array.Int32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint32:
		fillValue(listBuilder.ValueBuilder().(*array.Uint32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Int64:
		fillValue(listBuilder.ValueBuilder().(*array.Int64Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Uint64:
		fillValue(listBuilder.ValueBuilder().(*array.Uint64Builder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.Float16:
		fillFloat16Value(listBuilder.ValueBuilder().(*array.Float16Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Float32:
		fillValue(listBuilder.ValueBuilder().(*array.Float32Builder).AppendValues, rows, 0)
	case arrow.PrimitiveTypes.Float64:
		fillValue(listBuilder.ValueBuilder().(*array.Float64Builder).AppendValues, rows, 0)
	case arrow.BinaryTypes.Binary:
		fillBinaryValue(listBuilder.ValueBuilder().(*array.BinaryBuilder).AppendValues, rows, 0)
	case arrow.BinaryTypes.String:
		fillStringValue(listBuilder.ValueBuilder().(*array.StringBuilder).AppendValues, rows, 0)
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
		fillIntervalMonthValue(listBuilder.ValueBuilder().(*array.MonthIntervalBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.DayTimeInterval:
		fillIntervalDayTimeValue(listBuilder.ValueBuilder().(*array.DayTimeIntervalBuilder).AppendValues, rows, 0)
	case arrow.FixedWidthTypes.MonthDayNanoInterval:
		fillIntervalMonthDayNanoValue(listBuilder.ValueBuilder().(*array.MonthDayNanoIntervalBuilder).AppendValues, rows, 0)
	}

	return listBuilder.NewArray()
}

func fillBoolValue(append func(value []bool, valid []bool), rows int) {
	append(getBoolSlice(rows), nil)
}

func fillValue[T constraints.Integer | constraints.Float](append func(value []T, valid []bool), rows int, start int) {
	append(getSlice(rows, T(start)), nil)
}

func fillFloat16Value(append func(value []float16.Num, valid []bool), rows int, start int) {
	append(getFloat16Slice(rows, float32(start)), nil)
}

func fillBinaryValue(append func(value [][]byte, valid []bool), rows int, start int) {
	value, valid := getBinarySlice(rows, start)
	append(value, valid)
}

func fillStringValue(append func(value []string, valid []bool), rows int, start int) {
	append(getStringSlice(rows, start), nil)
}

func fillIntervalMonthValue(append func(value []arrow.MonthInterval, valid []bool), rows int, start int) {
	append(getIntervalMonthSlice(rows, start), nil)
}

func fillIntervalDayTimeValue(append func(value []arrow.DayTimeInterval, valid []bool), rows int, start int) {
	append(getIntervalDayTimeSlice(rows, start), nil)
}

func fillIntervalMonthDayNanoValue(append func(value []arrow.MonthDayNanoInterval, valid []bool), rows int, start int) {
	append(getIntervalMonthDayNanoSlice(rows, start), nil)
}

func getBoolSlice(rows int) []bool {
	slice := make([]bool, rows)
	for i := 0; i < rows; i++ {
		slice[i] = i%2 == 0
	}
	return slice
}

func getSlice[T constraints.Integer | constraints.Float](rows int, start T) []T {
	slice := make([]T, rows)
	for i := int64(0); i < int64(rows); i++ {
		slice[i] = T(i) + start
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

func getBinarySlice(rows int, start int) ([][]byte, []bool) {
	slice := make([][]byte, rows)
	valid := make([]bool, rows)
	j := 0
	for i := 0; i < rows; i++ {
		valid[i] = i%2 == 0
		if valid[i] {
			slice[i] = []byte{byte(j + start), byte(j + start + 1), byte(j + start + 2), byte(j + start + 3)}
			j++
		}
	}
	return slice, valid
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

func getIntervalMonthSlice(rows int, start int) []arrow.MonthInterval {
	slice := make([]arrow.MonthInterval, rows)
	for i := start; i < start+rows; i++ {
		slice[i] = arrow.MonthInterval(i)
	}
	return slice
}

func getIntervalDayTimeSlice(rows int, start int) []arrow.DayTimeInterval {
	slice := make([]arrow.DayTimeInterval, rows)
	for i := start; i < start+rows; i++ {
		slice[i] = arrow.DayTimeInterval{Days: int32(i), Milliseconds: int32(1000 + i)}
	}
	return slice
}

func getIntervalMonthDayNanoSlice(rows int, start int) []arrow.MonthDayNanoInterval {
	slice := make([]arrow.MonthDayNanoInterval, rows)
	for i := start; i < start+rows; i++ {
		slice[i] = arrow.MonthDayNanoInterval{Months: int32(i), Days: int32(100 + i), Nanoseconds: int64(10000 + i)}
	}
	return slice
}

func mockInt8(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt8Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt8Array()
}

func mockUint8(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint8Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint8Array()
}

func mockInt16(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt16Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt16Array()
}

func mockUint16(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint16Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint16Array()
}

func mockInt32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt32Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt32Array()
}

func mockUint32(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint32Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint32Array()
}

func mockInt64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewInt64Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewInt64Array()
}

func mockUint64(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewUint64Builder(mem)
	defer ib.Release()
	fillValue(ib.AppendValues, rows, 0)
	return ib.NewUint64Array()
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

func mockFixedSizeBinary(mem memory.Allocator, rows int) arrow.Array {
	nbytes := 5
	ib := array.NewFixedSizeBinaryBuilder(mem, &arrow.FixedSizeBinaryType{ByteWidth: nbytes})
	defer ib.Release()
	start := 0
	slice := make([][]byte, rows)
	valid := make([]bool, rows)
	j := 0
	for i := 0; i < rows; i++ {
		valid[i] = i%2 == 0
		if valid[i] {
			pos := (j + start) % len(stringData)
			slice[i] = []byte(stringData[pos : pos+nbytes])
			j += 1
		}
	}
	ib.AppendValues(slice, valid)
	return ib.NewFixedSizeBinaryArray()
}

func mockBool(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewBooleanBuilder(mem)
	defer ib.Release()
	for i := 0; i < rows; i++ {
		ib.Append(i%2 == 0)
	}
	return ib.NewBooleanArray()
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

func mockMonthInterval(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewMonthIntervalBuilder(mem)
	defer ib.Release()
	fillIntervalMonthValue(ib.AppendValues, rows, 0)
	return ib.NewMonthIntervalArray()
}

func mockDayTimeInterval(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDayTimeIntervalBuilder(mem)
	defer ib.Release()
	fillIntervalDayTimeValue(ib.AppendValues, rows, 0)
	return ib.NewDayTimeIntervalArray()
}

func mockMonthDayNanoInterval(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewMonthDayNanoIntervalBuilder(mem)
	defer ib.Release()
	fillIntervalMonthDayNanoValue(ib.AppendValues, rows, 0)
	return ib.NewMonthDayNanoIntervalArray()
}

func mockDecimal128(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDecimal128Builder(mem, &arrow.Decimal128Type{Precision: 37, Scale: 2})
	defer ib.Release()

	for i := 0; i < rows; i++ {
		v:= new(big.Int).SetUint64(uint64(math.Pow(2, 64)-1)) 
		v = v.Add(v, big.NewInt(int64(i)))
		ib.Append(decimal128.FromBigInt(v))
	}
	
	return ib.NewDecimal128Array()
}

func mockDecimal256 (mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDecimal256Builder(mem, &arrow.Decimal256Type{Precision: 76, Scale: 4})
	defer ib.Release()

	for i := 0; i < rows; i++ {
		v:= new(big.Int).SetUint64(uint64(math.Pow(2, 64)-1))
		v = v.Add(v, big.NewInt(int64(i)))
		ib.Append(decimal256.FromBigInt(v))
	}

	return ib.NewDecimal256Array()
}

func mockSampleList(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewListBuilder(mem, arrow.PrimitiveTypes.Int32)
	defer ib.Release()

	const subListLength = 2
	// this should fill `rows/3*subListLength` int32 in the value array
	// rows = 1, length = 2
	// rows = 2, length = 2
	// rows = 3, length = 2
	// rows = 4, length = 4
	valueBuilder := ib.ValueBuilder().(*array.Int32Builder)
	fillValue(valueBuilder.AppendValues, max(int(math.Ceil(float64(rows)/3.0)), 1)*subListLength, 0)

	// let's make every 2 of them as a sub-list
	// with `nil` and `[]` in between
	// when `rows == 4`, we're expecting the following results
	//   [[0,1], nil, [], [2,3]]
	// therefore, the offsets should be
	//   [0, 2, 2, 2, 4]
	offsets := make([]int32, rows+1)
	valid := make([]bool, rows)
	offset := int32(0)
	for i := 0; i < rows+1; i++ {
		offsets[i] = offset
		rem := i % 3
		switch rem {
		case 0:
			offset += subListLength
			fallthrough
		case 2:
			if i < rows {
				valid[i] = true
			}
		case 1:
			if i < rows {
				valid[i] = false
			}
		}
	}
	ib.AppendValues(offsets, valid)

	return ib.NewListArray()
}

func mockSampleListWithStruct(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewListBuilder(mem, arrow.StructOf(
		arrow.Field{Name: "start_time", Type: arrow.FixedWidthTypes.Timestamp_s},
		arrow.Field{Name: "end_time", Type: arrow.FixedWidthTypes.Timestamp_s},
		arrow.Field{Name: "data_points", Type: arrow.PrimitiveTypes.Int32},
	))
	defer ib.Release()

	structBuilder := ib.ValueBuilder().(*array.StructBuilder)
	startTimeBuilder := structBuilder.FieldBuilder(0).(*array.TimestampBuilder)
	endTimeBuilder := structBuilder.FieldBuilder(1).(*array.TimestampBuilder)
	dataPointsBuilder := structBuilder.FieldBuilder(2).(*array.Int32Builder)

	fillValue(startTimeBuilder.AppendValues, rows, 0)
	fillValue(endTimeBuilder.AppendValues, rows, 100)
	fillValue(dataPointsBuilder.AppendValues, rows, 0)
	structValid := make([]bool, rows)
	for i := 0; i < rows; i++ {
		structValid[i] = true
	}
	structBuilder.AppendValues(structValid)

	// one data point each row
	offsets := make([]int32, rows+1)
	valid := make([]bool, rows)
	for i := 0; i < rows+1; i++ {
		offsets[i] = int32(i)
		if i < rows {
			valid[i] = true
		}
	}

	ib.AppendValues(offsets, valid)

	return ib.NewListArray()
}

func mockSampleNestedList(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewListBuilder(mem, arrow.ListOf(arrow.PrimitiveTypes.Int32))
	defer ib.Release()

	innerListBuilder := ib.ValueBuilder().(*array.ListBuilder)
	valueBuilder := innerListBuilder.ValueBuilder().(*array.Int32Builder)

	// for nested lists, lets generated the following pattern
	// [
	//  [[1], nil, []],    <- row 0,
	//  [[2,3], nil, []]   <- row 1,
	//  [[4,5,6], nil, []] <- row 2, ...
	// ]
	// so we need `rows*(rows+1)/2` values, where `n == rows`
	// rows = 1, values = 1
	// rows = 2, values = 3
	// rows = 3, values = 6
	fillValue(valueBuilder.AppendValues, rows*(rows+1)/2, 0)
	// for all inner lists
	// the offsets for the inner list are
	//   [0, 1, 1, 1, 3, 3, 3, 6, 6, 6]
	//   rows = 1, [0, 1, 1, 1]
	//   rows = 2, [0, 1, 1, 1, 3, 3, 3]
	//   rows = 3, [0, 1, 1, 1, 3, 3, 3, 6, 6, 6]
	// the validity is
	//   rows = 1, [true, false, true]
	//   rows = 2, [true, false, true, true, false, true]
	//   rows = 2, [true, false, true, true, false, true, true, false, true]
	innerOffsets := make([]int32, rows*3+1)
	innerValid := make([]bool, rows*3)
	innerOffsets[0] = 0
	innerValid[0] = true
	for i := 1; i < rows*3+1; i++ {
		j := int32(math.Ceil(float64(i) / 3.0))
		innerOffsets[i] = j * (j + 1) / 2
		if i < rows*3 {
			if i%3 == 1 {
				innerValid[i] = false
			} else {
				innerValid[i] = true
			}
		}
	}
	innerListBuilder.AppendValues(innerOffsets, innerValid)
	// for the outer list
	// they're all valid, and the offsets are
	//   [0, 1, 2, 3]
	//   rows = 1, [0, 1]
	//   rows = 2, [0, 1, 2]
	outerOffsets := make([]int32, rows+1)
	outerValid := make([]bool, rows)
	for i := 0; i < rows+1; i++ {
		outerOffsets[i] = int32(i * 3)
		if i < rows {
			outerValid[i] = true
		}
	}
	ib.AppendValues(outerOffsets, outerValid)

	return ib.NewListArray()
}

func mockSampleListView(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewListViewBuilder(mem, arrow.PrimitiveTypes.Int32)
	defer ib.Release()

	const listSize = 5
	offsets := make([]int32, rows)
	sizes := make([]int32, rows)
	valid := make([]bool, rows)
	values := make([]int32, rows*3+2)
	valuesValid := make([]bool, rows*3+2)
	for i := 0; i < rows*3+2; i++ {
		values[i] = int32(i)
		valuesValid[i] = true
	}
	for i := 0; i < rows; i++ {
		offsets[i] = int32(i * 3)
		sizes[i] = listSize
		valid[i] = i%2 == 0
	}
	ib.AppendValuesWithSizes(offsets, sizes, valid)
	ib.ValueBuilder().(*array.Int32Builder).AppendValues(values, valuesValid)

	return ib.NewListViewArray()
}

func mockSampleLargeListView(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewLargeListViewBuilder(mem, arrow.PrimitiveTypes.Int32)
	defer ib.Release()

	const listSize = 5
	offsets := make([]int64, rows)
	sizes := make([]int64, rows)
	valid := make([]bool, rows)
	values := make([]int32, rows*3+2)
	valuesValid := make([]bool, rows*3+2)
	for i := 0; i < rows*3+2; i++ {
		values[i] = int32(i)
		valuesValid[i] = true
	}
	for i := 0; i < rows; i++ {
		offsets[i] = int64(i * 3)
		sizes[i] = listSize
		valid[i] = i%2 == 0
	}
	ib.AppendValuesWithSizes(offsets, sizes, valid)
	ib.ValueBuilder().(*array.Int32Builder).AppendValues(values, valuesValid)

	return ib.NewLargeListViewArray()
}

func mockSampleFixedSizeList(mem memory.Allocator, rows int) arrow.Array {
	fixedSize := int32(3)
	ib := array.NewFixedSizeListBuilder(mem, fixedSize, arrow.PrimitiveTypes.Int32)
	defer ib.Release()

	values := make([]int32, fixedSize)
	valid := make([]bool, fixedSize)

	increment := 1


	for i := 0; i < rows; i++ {
		for j := 0; j < int(fixedSize); j++ {
			values[j] = int32(increment)
			valid[j] = j%2 == 0
			increment++
		}
		ib.Append(true)
		ib.ValueBuilder().(*array.Int32Builder).AppendValues(values, valid)
	}

	return ib.NewListArray()
}

func mockSampleNestedFixedSizeList(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewFixedSizeListBuilder(mem, 3, arrow.FixedSizeListOf(3, arrow.PrimitiveTypes.Int32))
	defer ib.Release()

	innerBuilder := ib.ValueBuilder().(*array.FixedSizeListBuilder)
	intBuilder := innerBuilder.ValueBuilder().(*array.Int32Builder)

	values := make([]int32, 3)
	valid := make([]bool, 3)

	increment := 1

	for i := 0; i < rows; i++ {
		for j := 0; j < 3; j++ {
			values[j] = int32(increment)
			valid[j] = j%2 == 0
			increment++
		}
		ib.Append(true)
		for j := 0; j < 3; j++ {
			innerBuilder.Append(true)
			intBuilder.AppendValues(values, valid)
		}
	}

	return ib.NewListArray()
}

func mockSampleRunEndEncodedArray(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewRunEndEncodedBuilder(mem, arrow.PrimitiveTypes.Int32,arrow.PrimitiveTypes.Float32)
	defer ib.Release()

	runs := make([]uint64, rows)
	for i := 0; i < rows; i++ {
		runs[i] = uint64(i)
	}

	values := make([]float32, rows)
	for i := 0; i < rows; i++ {
		values[i] = float32(i+1)
	}

	valid:= make([]bool, rows)

	for i := 0; i < rows; i++ {
		valid[i] = i%2 == 0
	}

	ib.AppendRuns(runs)
	ib.ValueBuilder().(*array.Float32Builder).AppendValues(values,valid)

	return ib.NewRunEndEncodedArray()
}

func mockSampleDictionaryEncodedArray(mem memory.Allocator, rows int) arrow.Array {
	ib := array.NewDictionaryBuilder(mem,&arrow.DictionaryType{
		IndexType: arrow.PrimitiveTypes.Int32,
		ValueType: arrow.BinaryTypes.String,
	}).(*array.BinaryDictionaryBuilder)

	ib.AppendString("hello")
	ib.AppendString("goodbye")

	values := make([]int, rows)
	for i := 0; i < rows; i++ {
		values[i] = i%2
	}

	valid := make([]bool, rows)
	for i := 0; i < rows; i++ {
		valid[i] = true
	}
	
	ib.AppendIndices(values,valid)

	defer ib.Release()
	return ib.NewDictionaryArray()
}

func mockSampleDenseUnion(mem memory.Allocator, rows int) arrow.Array {

	intBuilder := array.NewInt32Builder(mem)
	defer intBuilder.Release()
	strBuilder := array.NewStringBuilder(mem)
	defer strBuilder.Release()

	ib := array.NewDenseUnionBuilderWithBuilders(mem, arrow.DenseUnionOf(
		[]arrow.Field{
			{Name: "a", Type: arrow.PrimitiveTypes.Int32},
			{Name: "b", Type: arrow.BinaryTypes.String},
		},
		[]arrow.UnionTypeCode{0, 1},
	), []array.Builder{ intBuilder,strBuilder })

	defer ib.Release()

	for i := 0; i < rows; i++ {
		ib.Append(0)
		intBuilder.Append(int32(i))

		ib.Append(1)
		strBuilder.Append(fmt.Sprint("str%d",i))
	}

	return ib.NewDenseUnionArray()
}

func mockSampleSparseUnion(mem memory.Allocator, rows int) arrow.Array {

	intBuilder := array.NewInt32Builder(mem)
	defer intBuilder.Release()
	strBuilder := array.NewStringBuilder(mem)
	defer strBuilder.Release()

	ib := array.NewSparseUnionBuilderWithBuilders(mem, arrow.SparseUnionOf(
		[]arrow.Field{
			{Name: "a", Type: arrow.PrimitiveTypes.Int32},
			{Name: "b", Type: arrow.BinaryTypes.String},
		},
		[]arrow.UnionTypeCode{0, 1},
	), []array.Builder{ intBuilder,strBuilder })

	defer ib.Release()

	for i := 0; i < rows; i++ {
		ib.Append(0)
		intBuilder.Append(int32(i))
		strBuilder.AppendNull()

		ib.Append(1)
		strBuilder.Append(fmt.Sprint("str%d",i))
		intBuilder.AppendNull()
	}

	return ib.NewSparseUnionArray()
}

func mockNull(mem memory.Allocator, rows int) arrow.Array {
	return array.NewNull(rows)
}

// func mockList(mem memory.Allocator, rows, length int64, innerList arrow.Array) arrow.Array {
// 	lb := array.NewListBuilder(mem, arrow.ListOf(innerList.DataType()))
// 	defer lb.Release()
// 	// vb := lb.ValueBuilder().(*array.ListBuilder).Append()
// }
