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
	"encoding/binary"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/decimal128"
	"github.com/apache/arrow/go/v17/arrow/decimal256"
	"github.com/apache/arrow/go/v17/arrow/float16"
	"github.com/apache/arrow/go/v17/arrow/memory"
)

var handlerForType map[int]func(field arrow.Field, rows int) arrow.Array

func init(){
	handlerForType = map[int]func(field arrow.Field, rows int) arrow.Array{
		int(arrow.NULL):              mockNull,
		int(arrow.BOOL):              mockBool,
		int(arrow.UINT8):             mockUint8,
		int(arrow.INT8):              mockInt8,
		int(arrow.UINT16):            mockUint16,
		int(arrow.INT16):             mockInt16,
		int(arrow.UINT32):            mockUint32,
		int(arrow.INT32):             mockInt32,
		int(arrow.UINT64):            mockUint64,
		int(arrow.INT64):             mockInt64,
		int(arrow.FLOAT16):           mockFloat16,
		int(arrow.FLOAT32):           mockFloat32,
		int(arrow.FLOAT64):           mockFloat64,
		int(arrow.STRING):            mockString,
		int(arrow.BINARY):            mockBinary,
		int(arrow.FIXED_SIZE_BINARY): mockFixedSizeBinary,
		int(arrow.DATE32):            mockDate32,
		int(arrow.DATE64):            mockDate64,
		//
		// commented out until mock function implemented
		//
		int(arrow.TIMESTAMP): mockTimestamp,
		int(arrow.TIME32):   mockTime32,
		int(arrow.TIME64):   mockTime64,
		int(arrow.INTERVAL_MONTHS): mockIntervalMonths,
		int(arrow.INTERVAL_DAY_TIME): mockIntervalDays,
		int(arrow.DECIMAL128): mockDecimal128,
		int(arrow.DECIMAL256): mockDecimal256,
		int(arrow.LIST):      mockList,
		int(arrow.STRUCT):    mockStruct,
		// int(arrow.SPARSE_UNION): mockSparseUnion,
		// int(arrow.DENSE_UNION):  mockDenseUnion,
		// int(arrow.DICTIONARY):   mockDictionary,
		// int(arrow.MAP):          mockMap,
		//
		// TODO: add other types
		int(arrow.DURATION): mockDuration,
	}
}

func mockNull(field arrow.Field, rows int) arrow.Array {
	return array.NewNull(rows)
}

func mockBool(field arrow.Field, rows int) arrow.Array {
	builder := array.NewBooleanBuilder(memory.DefaultAllocator)
	
	for i := 0; i < rows; i++ {
		builder.Append(i % 2 == 0)
	}

	return builder.NewArray()
}

func mockUint8(field arrow.Field, rows int) arrow.Array {
	builder := array.NewUint8Builder(memory.DefaultAllocator)
	
	for i := 0; i < rows; i++ {
		builder.Append(uint8(i))
	}

	return builder.NewArray()
}

func mockInt8(field arrow.Field, rows int) arrow.Array {
	builder := array.NewInt8Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(int8(i * []int{-1, 1}[i%2]))
	}

	return builder.NewArray()
}

func mockUint16(field arrow.Field, rows int) arrow.Array {
	builder := array.NewUint16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(uint16(i))
	}

	return builder.NewArray()
}

func mockInt16(field arrow.Field, rows int) arrow.Array {
	builder := array.NewInt16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(int16(i * []int{-1, 1}[i%2]))
	}

	return builder.NewArray()
}

func mockUint32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewUint32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(uint32(i))
	}

	return builder.NewArray()
}

func mockInt32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewInt32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(int32(i * []int{-1, 1}[i%2]))
	}	

	return builder.NewArray()
}

func mockUint64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewUint64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(uint64(i))
	}

	return builder.NewArray()
}

func mockInt64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewInt64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(int64(i * []int{-1, 1}[i%2]))
	}

	return builder.NewArray()
}

func mockFloat16(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFloat16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(float16.New(float32(i * []int{-1, 1}[i%2])/10))
	}

	return builder.NewArray()
}

func mockFloat32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFloat32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(float32(i * []int{-1, 1}[i%2])/10)
	}

	return builder.NewArray()
}

func mockFloat64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFloat64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(float64(i * []int{-1, 1}[i%2])/10)
	}

	return builder.NewArray()
}

func mockString(field arrow.Field, rows int) arrow.Array {
	builder := array.NewStringBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(strconv.Itoa(i))
	}

	return builder.NewArray()
}

func mockBinary(field arrow.Field, rows int) arrow.Array {
	builder := array.NewBinaryBuilder(memory.DefaultAllocator, arrow.BinaryTypes.Binary)

	for i := 0; i < rows; i++ {
		builder.Append([]byte(strconv.Itoa(i)))
	}

	return builder.NewArray()
}

func mockFixedSizeBinary(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFixedSizeBinaryBuilder(memory.DefaultAllocator,field.Type.(*arrow.FixedSizeBinaryType))
	byteWidth := field.Type.(*arrow.FixedSizeBinaryType).ByteWidth

	for i := 0; i < rows; i++ {
		numStr := strconv.Itoa(i)
		strBytes := []byte(numStr)
		paddingBytes := make([]byte, byteWidth-len(strBytes))
		builder.Append(append(paddingBytes, strBytes...))
	}

	return builder.NewArray()
}

func mockDate32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDate32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Date32FromTime(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i)))
	}

	return builder.NewArray()
}

func mockDate64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDate64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Date64FromTime(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i)))
	}

	return builder.NewArray()
}

func mockTimestamp(field arrow.Field, rows int) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, field.Type.(*arrow.TimestampType))

	for i := 0; i < rows; i++ {
		timestamp,_ := arrow.TimestampFromTime(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i), field.Type.(*arrow.TimestampType).TimeUnit())
		builder.Append(timestamp)
	}

	return builder.NewArray()
}

func mockTime32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewTime32Builder(memory.DefaultAllocator, field.Type.(*arrow.Time32Type))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Time32(i))
	}

	return builder.NewArray()
}

func mockTime64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewTime64Builder(memory.DefaultAllocator, field.Type.(*arrow.Time64Type))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Time64(i))
	}

	return builder.NewArray()
}

func mockIntervalMonths(field arrow.Field, rows int) arrow.Array {
	builder := array.NewMonthIntervalBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.MonthInterval(i))
	}

	return builder.NewArray()
}

func mockIntervalDays(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDayTimeIntervalBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.DayTimeInterval{Days: int32(i), Milliseconds: int32(i)})
	}

	return builder.NewArray()
}

func mockDecimal128(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDecimal128Builder(memory.DefaultAllocator, field.Type.(*arrow.Decimal128Type))
	number := big.NewInt(math.MaxInt64)
	for i := 0; i < rows; i++ {
		builder.Append(decimal128.FromBigInt(number.Add(number, big.NewInt(int64(i)))))
	}

	return builder.NewArray()
}

func mockDecimal256(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDecimal256Builder(memory.DefaultAllocator, field.Type.(*arrow.Decimal256Type))
	number := big.NewInt(math.MaxInt64)
	for i := 0; i < rows; i++ {
		builder.Append(decimal256.FromBigInt(number.Add(number, big.NewInt(int64(i)))))
	}

	return builder.NewArray()
}

func mockList(field arrow.Field, rows int) arrow.Array {
	listType := field.Type.(*arrow.ListType)
	innerField := listType.ElemField()
	listLength := 1
	if listLengthStr, ok := innerField.Metadata.GetValue("list_length"); ok {
		listLength, _ = strconv.Atoi(listLengthStr)
	}

	innerValue := handlerForType[int(innerField.Type.ID())](innerField, listLength*rows)

	offsets := make([]int32, rows+1)

	for i := 0; i < rows; i++ {
		offsets[i] = int32(i * listLength)
	}
	offsets[rows] = int32(innerValue.Len())

	offsetBytes := make([]byte, (rows+1)*4)

	for i,v := range offsets {
		binary.LittleEndian.PutUint32(offsetBytes[i*4:], uint32(v))
	}

	offsetsBuffer := memory.NewBufferBytes(offsetBytes)

	arrData := array.NewData(field.Type,rows,[]*memory.Buffer{nil,offsetsBuffer },[]arrow.ArrayData{innerValue.Data()},0,0)

	return array.NewListData(arrData)
}

func mockStruct(field arrow.Field, rows int) arrow.Array {

	structType := field.Type.(*arrow.StructType)
	log.Printf("Struct type: %v", structType)
	log.Printf("Fields: %v", structType.Fields())

	innerDatas := make([]arrow.ArrayData, structType.NumFields())

	for i := 0; i < structType.NumFields(); i++ {
		innerField := structType.Field(i)
		innerDatas[i] = handlerForType[int(innerField.Type.ID())](innerField, rows).Data()
	}

	structData := array.NewData(field.Type, rows, []*memory.Buffer{nil}, innerDatas, 0, 0)

	return array.NewStructData(structData)
}

func mockDuration(field arrow.Field, rows int) arrow.Array {
	builder := array.NewDurationBuilder(memory.DefaultAllocator, field.Type.(*arrow.DurationType))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Duration(i))
	}

	return builder.NewArray()
}

func PopulateSchema(schema *arrow.Schema, rows int) arrow.Record {

	cols := make([]arrow.Array, len(schema.Fields()))

	for i, field := range schema.Fields() {
		cols[i] = handlerForType[int(field.Type.ID())](field, rows)
	}

	return array.NewRecord(schema, cols, int64(rows))
}