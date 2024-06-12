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
	"strconv"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/float16"
	"github.com/apache/arrow/go/v17/arrow/memory"
)

var (
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
		// int(arrow.FIXED_SIZE_BINARY): mockFixedSizeBinary,
		// int(arrow.DATE32):            mockDate32,
		// int(arrow.DATE64):            mockDate64,
		//
		// commented out until mock function implemented
		//
		// int(arrow.TIMESTAMP): mockTimestamp,
		// int(arrow.TIME32):   mockTime32,
		// int(arrow.TIME64):   mockTime64,
		// int(arrow.INTERVAL_MONTHS): mockIntervalMonth,
		// int(arrow.INTERVAL_DAY_TIME): mockIntervalDayTime,
		// int(arrow.DECIMAL128): mockDecimal128,
		// int(arrow.DECIMAL256): mockDecimal256,
		// int(arrow.LIST):      mockList,
		// int(arrow.STRUCT):    mockStruct,
		// int(arrow.SPARSE_UNION): mockSparseUnion,
		// int(arrow.DENSE_UNION):  mockDenseUnion,
		// int(arrow.DICTIONARY):   mockDictionary,
		// int(arrow.MAP):          mockMap,
		//
		// TODO: add other types
	}
)

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
		builder.Append(float16.New(float32(i * []int{-1, 1}[i%2])))
	}

	return builder.NewArray()
}

func mockFloat32(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFloat32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(float32(i * []int{-1, 1}[i%2]))
	}

	return builder.NewArray()
}

func mockFloat64(field arrow.Field, rows int) arrow.Array {
	builder := array.NewFloat64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(float64(i * []int{-1, 1}[i%2]))
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

func PopulateSchema(schema *arrow.Schema, rows int) arrow.Record {

	cols := make([]arrow.Array, len(schema.Fields()))

	for i, field := range schema.Fields() {
		cols[i] = handlerForType[int(field.Type.ID())](field, rows)
	}

	return array.NewRecord(schema, cols, int64(rows))
}