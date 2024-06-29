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
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/decimal128"
	"github.com/apache/arrow/go/v17/arrow/decimal256"
	"github.com/apache/arrow/go/v17/arrow/float16"
	"github.com/apache/arrow/go/v17/arrow/memory"
)

var handlerForType map[int]func(field arrow.Field, rows int, level int) arrow.Array

func init() {
	handlerForType = map[int]func(field arrow.Field, rows int, level int) arrow.Array{
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
		int(arrow.TIMESTAMP):               mockTimestamp,
		int(arrow.TIME32):                  mockTime32,
		int(arrow.TIME64):                  mockTime64,
		int(arrow.INTERVAL_MONTHS):         mockIntervalMonths,
		int(arrow.INTERVAL_DAY_TIME):       mockIntervalDays,
		int(arrow.INTERVAL_MONTH_DAY_NANO): mockIntervalMonthDayNano,
		int(arrow.DECIMAL128):              mockDecimal128,
		int(arrow.DECIMAL256):              mockDecimal256,
		int(arrow.LIST):                    mockList,
		int(arrow.STRUCT):                  mockStruct,
		int(arrow.SPARSE_UNION):            mockSparseUnion,
		int(arrow.DENSE_UNION):             mockDenseUnion,
		int(arrow.DICTIONARY):              mockDictionary,
		int(arrow.DURATION):                mockDuration,
		int(arrow.RUN_END_ENCODED):         mockRunEndEncoded,
	}
}

func mockNull(field arrow.Field, rows int, level int) arrow.Array {
	return array.NewNull(rows)
}

func mockBool(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewBooleanBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(i%2 == 0)
	}

	return builder.NewArray()
}

func mockUint8(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewUint8Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 1:
				builder.Append(uint8(math.MaxUint8))
			default:
				builder.Append(uint8(i))
			}
		} else {
			builder.Append(uint8(i))
		}
	}

	return builder.NewArray()
}

func mockInt8(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewInt8Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(int8(math.MinInt8))
			case 1:
				builder.Append(int8(math.MaxInt8))
			default:
				builder.Append(int8(i * []int{-1, 1}[i%2]))
			}
		} else {
			builder.Append(int8(i * []int{-1, 1}[i%2]))
		}
	}

	return builder.NewArray()
}

func mockUint16(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewUint16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 1:
				builder.Append(uint16(math.MaxUint16))
			default:
				builder.Append(uint16(i))
			}
		} else {
			builder.Append(uint16(i))
		}
	}

	return builder.NewArray()
}

func mockInt16(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewInt16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(int16(math.MinInt16))
			case 1:
				builder.Append(int16(math.MaxInt16))
			default:
				builder.Append(int16(i * []int{-1, 1}[i%2]))
			}
		} else {
			builder.Append(int16(i * []int{-1, 1}[i%2]))
		}
	}

	return builder.NewArray()
}

func mockUint32(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewUint32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 1:
				builder.Append(uint32(math.MaxUint32))
			default:
				builder.Append(uint32(i))
			}
		} else {
			builder.Append(uint32(i))
		}
	}

	return builder.NewArray()
}

func mockInt32(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewInt32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(int32(math.MinInt32))
			case 1:
				builder.Append(int32(math.MaxInt32))
			default:
				builder.Append(int32(i * []int{-1, 1}[i%2]))
			}
		} else {
			builder.Append(int32(i * []int{-1, 1}[i%2]))
		}
	}

	return builder.NewArray()
}

func mockUint64(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewUint64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 1:
				builder.Append(uint64(math.MaxUint64))
			default:
				builder.Append(uint64(i))
			}
		} else {
			builder.Append(uint64(i))
		}
	}

	return builder.NewArray()
}

func mockInt64(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewInt64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(int64(math.MinInt64))
			case 1:
				builder.Append(int64(math.MaxInt64))
			default:
				builder.Append(int64(i * []int{-1, 1}[i%2]))
			}
		} else {
			builder.Append(int64(i * []int{-1, 1}[i%2]))
		}
	}

	return builder.NewArray()
}

func mockFloat16(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewFloat16Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(float16.MinNum)
			case 1:
				builder.Append(float16.MaxNum)
			case 2:
				builder.Append(float16.Inf())
			case 3:
				builder.Append(float16.Inf().Negate())
			case 4:
				builder.Append(float16.NaN())
			case 5:
				builder.Append(float16.New(0))
			case 6:
				builder.Append(float16.New(float32(math.Copysign(0, -1))))
			case 7:
				builder.Append(float16.FromBits(1))
			case 8:
				builder.Append(float16.FromBits(1).Negate())
			default:
				builder.Append(float16.New(float32(i*[]int{-1, 1}[i%2]) / 10))
			}
		} else {
			builder.Append(float16.New(float32(i*[]int{-1, 1}[i%2]) / 10))
		}
	}

	return builder.NewArray()
}

func mockFloat32(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewFloat32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(float32(-math.MaxFloat32))
			case 1:
				builder.Append(float32(math.MaxFloat32))
			case 2:
				builder.Append(float32(math.Inf(1)))
			case 3:
				builder.Append(float32(math.Inf(-1)))
			case 4:
				builder.Append(float32(math.NaN()))
			case 5:
				builder.Append(float32(math.Copysign(0, 1)))
			case 6:
				builder.Append(float32(math.Copysign(0, -1)))
			case 7:
				builder.Append(float32(math.SmallestNonzeroFloat32))
			case 8:
				builder.Append(float32(-math.SmallestNonzeroFloat32))
			default:
				builder.Append(float32(i*[]int{-1, 1}[i%2]) / 10)
			}
		} else {
			builder.Append(float32(i*[]int{-1, 1}[i%2]) / 10)
		}
	}

	return builder.NewArray()
}

func mockFloat64(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewFloat64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		if level == 0 {
			switch i {
			case 0:
				builder.Append(float64(-math.MaxFloat64))
			case 1:
				builder.Append(float64(math.MaxFloat64))
			case 2:
				builder.Append(float64(math.Inf(1)))
			case 3:
				builder.Append(float64(math.Inf(-1)))
			case 4:
				builder.Append(float64(math.NaN()))
			case 5:
				builder.Append(float64(math.Copysign(0, 1)))
			case 6:
				builder.Append(float64(math.Copysign(0, -1)))
			case 7:
				builder.Append(float64(math.SmallestNonzeroFloat64))
			case 8:
				builder.Append(float64(-math.SmallestNonzeroFloat64))
			default:
				builder.Append(float64(i*[]int{-1, 1}[i%2]) / 10)
			}
		} else {
			builder.Append(float64(i*[]int{-1, 1}[i%2]) / 10)
		}
	}

	return builder.NewArray()
}

func mockString(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewStringBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		value := strconv.Itoa(i)
		builder.Append(fmt.Sprintf("%s%s", strings.Repeat("0", i+1-len(value)), value))
	}

	return builder.NewArray()
}

func mockBinary(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewBinaryBuilder(memory.DefaultAllocator, arrow.BinaryTypes.Binary)

	for i := 0; i < rows; i++ {
		byteArray := []byte{uint8(i)}
		builder.Append(bytes.Repeat(byteArray, i+1))
	}

	return builder.NewArray()
}

func mockFixedSizeBinary(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewFixedSizeBinaryBuilder(memory.DefaultAllocator, field.Type.(*arrow.FixedSizeBinaryType))
	byteWidth := field.Type.(*arrow.FixedSizeBinaryType).ByteWidth

	for i := 0; i < rows; i++ {
		numStr := strconv.Itoa(i)
		strBytes := []byte(numStr)
		paddingBytes := make([]byte, byteWidth-len(strBytes))
		builder.Append(append(paddingBytes, strBytes...))
	}

	return builder.NewArray()
}

func mockDate32(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDate32Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Date32(i))
	}

	return builder.NewArray()
}

func mockDate64(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDate64Builder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Date64(1000*60*60*24*i + i))
	}

	return builder.NewArray()
}

func mockTimestamp(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewTimestampBuilder(memory.DefaultAllocator, field.Type.(*arrow.TimestampType))

	for i := 0; i < rows; i++ {
		timestamp, _ := arrow.TimestampFromTime(time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i), field.Type.(*arrow.TimestampType).TimeUnit())
		builder.Append(timestamp)
	}

	return builder.NewArray()
}

func mockTime32(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewTime32Builder(memory.DefaultAllocator, field.Type.(*arrow.Time32Type))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Time32(i))
	}

	return builder.NewArray()
}

func mockTime64(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewTime64Builder(memory.DefaultAllocator, field.Type.(*arrow.Time64Type))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Time64(i))
	}

	return builder.NewArray()
}

func mockIntervalMonths(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewMonthIntervalBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.MonthInterval(i))
	}

	return builder.NewArray()
}

func mockIntervalDays(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDayTimeIntervalBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.DayTimeInterval{Days: int32(i), Milliseconds: int32(i)})
	}

	return builder.NewArray()
}

func mockIntervalMonthDayNano(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewMonthDayNanoIntervalBuilder(memory.DefaultAllocator)

	for i := 0; i < rows; i++ {
		builder.Append(arrow.MonthDayNanoInterval{Months: int32(i), Days: int32(i), Nanoseconds: int64(i)})
	}

	return builder.NewArray()
}

func mockDecimal128(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDecimal128Builder(memory.DefaultAllocator, field.Type.(*arrow.Decimal128Type))
	number := big.NewInt(math.MaxInt64)
	for i := 0; i < rows; i++ {
		builder.Append(decimal128.FromBigInt(number.Add(number, big.NewInt(int64(i)))))
	}

	return builder.NewArray()
}

func mockDecimal256(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDecimal256Builder(memory.DefaultAllocator, field.Type.(*arrow.Decimal256Type))
	number := big.NewInt(math.MaxInt64)
	for i := 0; i < rows; i++ {
		builder.Append(decimal256.FromBigInt(number.Add(number, big.NewInt(int64(i)))))
	}

	return builder.NewArray()
}

func mockList(field arrow.Field, rows int, level int) arrow.Array {
	listType := field.Type.(*arrow.ListType)
	innerField := listType.ElemField()
	listLength := 1
	if listLengthStr, ok := innerField.Metadata.GetValue("list_length"); ok {
		listLength, _ = strconv.Atoi(listLengthStr)
	}

	innerValue := handlerForType[int(innerField.Type.ID())](innerField, listLength*rows, level+1)

	offsets := make([]int32, rows+1)

	for i := 0; i < rows; i++ {
		offsets[i] = int32(i * listLength)
	}
	offsets[rows] = int32(innerValue.Len())

	offsetBytes := make([]byte, (rows+1)*4)

	for i, v := range offsets {
		binary.LittleEndian.PutUint32(offsetBytes[i*4:], uint32(v))
	}

	offsetsBuffer := memory.NewBufferBytes(offsetBytes)

	arrData := array.NewData(field.Type, rows, []*memory.Buffer{nil, offsetsBuffer}, []arrow.ArrayData{innerValue.Data()}, 0, 0)

	return array.NewListData(arrData)
}

func mockStruct(field arrow.Field, rows int, level int) arrow.Array {

	structType := field.Type.(*arrow.StructType)
	log.Printf("Struct type: %v", structType)
	log.Printf("Fields: %v", structType.Fields())

	innerDatas := make([]arrow.ArrayData, structType.NumFields())

	for i := 0; i < structType.NumFields(); i++ {
		innerField := structType.Field(i)
		innerDatas[i] = handlerForType[int(innerField.Type.ID())](innerField, rows, level+1).Data()
	}

	structData := array.NewData(field.Type, rows, []*memory.Buffer{nil}, innerDatas, 0, 0)

	return array.NewStructData(structData)
}

func mockSparseUnion(field arrow.Field, rows int, level int) arrow.Array {
	// until we figure out how to get values from the metadata
	// we'll just create 1 value for each field
	// so union<bool,string,int8> will be [true, "0", 0]

	unionType := field.Type.(*arrow.SparseUnionType)
	innerFields := unionType.Fields()

	typeIdsBuilder := array.NewInt8Builder(memory.DefaultAllocator)
	for i := 0; i < rows; i++ {
		typeIdsBuilder.Append(int8(i % len(innerFields)))
	}
	typeIds := typeIdsBuilder.NewArray()

	innerDatas := make([]arrow.Array, len(innerFields))
	for i := 0; i < len(innerFields); i++ {
		innerField := innerFields[i]
		innerDatas[i] = handlerForType[int(innerField.Type.ID())](innerField, rows, level+1)
	}

	names := make([]string, len(innerFields))
	for i, f := range innerFields {
		names[i] = f.Name
	}

	thisUnion, err := array.NewSparseUnionFromArraysWithFields(typeIds, innerDatas, names)

	if err != nil {
		panic(err)
	}
	return thisUnion
}

func mockDenseUnion(field arrow.Field, rows int, level int) arrow.Array {
	unionType := field.Type.(*arrow.DenseUnionType)
	innerFields := unionType.Fields()

	typeIdsBuilder := array.NewInt8Builder(memory.DefaultAllocator)
	for i := 0; i < rows; i++ {
		typeIdsBuilder.Append(int8(i % len(innerFields)))
	}
	typeIds := typeIdsBuilder.NewArray()

	innerDatas := make([]arrow.Array, len(innerFields))
	for i := 0; i < len(innerFields); i++ {
		innerField := innerFields[i]
		innerDatas[i] = handlerForType[int(innerField.Type.ID())](innerField, rows, level+1)
	}

	names := make([]string, len(innerFields))
	for i, f := range innerFields {
		names[i] = f.Name
	}

	offsetsBuilder := array.NewInt32Builder(memory.DefaultAllocator)
	for i := 0; i < rows; i++ {
		offsetsBuilder.Append(int32(i / len(innerFields)))
	}

	thisUnion, err := array.NewDenseUnionFromArraysWithFields(typeIds, offsetsBuilder.NewArray(), innerDatas, names)

	if err != nil {
		panic(err)
	}
	return thisUnion
}

func mockDictionary(field arrow.Field, rows int, level int) arrow.Array {
	// indexType := arrow.PrimitiveTypes.Int32
	valueType := field.Type.(*arrow.DictionaryType).ValueType
	values := handlerForType[int(valueType.ID())](arrow.Field{Type: valueType}, rows, level+1)

	indicesBuilder := array.NewInt32Builder(memory.DefaultAllocator)
	for i := 0; i < rows; i++ {
		indicesBuilder.Append(int32(i))
	}
	indices := indicesBuilder.NewArray()

	return array.NewDictionaryArray(field.Type, indices, values)
}

func mockDuration(field arrow.Field, rows int, level int) arrow.Array {
	builder := array.NewDurationBuilder(memory.DefaultAllocator, field.Type.(*arrow.DurationType))

	for i := 0; i < rows; i++ {
		builder.Append(arrow.Duration(i))
	}

	return builder.NewArray()
}

func mockRunEndEncoded(field arrow.Field, rows int, level int) arrow.Array {
	innerType := field.Type.(*arrow.RunEndEncodedType).Encoded()
	innerValue := handlerForType[int(innerType.ID())](arrow.Field{Type: innerType}, rows, level+1)

	runEndsBuilder := array.NewInt32Builder(memory.DefaultAllocator)
	for i := 0; i < rows; i++ {
		runEndsBuilder.Append(int32(i) + 1)
	}
	runEnds := runEndsBuilder.NewArray()

	return array.NewRunEndEncodedArray(runEnds, innerValue, rows, 0)
}

func PopulateSchema(schema *arrow.Schema, rows int) arrow.Record {

	cols := make([]arrow.Array, len(schema.Fields()))

	for i, field := range schema.Fields() {
		cols[i] = handlerForType[int(field.Type.ID())](field, rows, 0)
	}

	return array.NewRecord(schema, cols, int64(rows))
}
