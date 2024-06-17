package mocks_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/apache/arrow-adbc/go/adbc"
	driver "github.com/apache/arrow-adbc/go/adbc/driver/mocks"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/stretchr/testify/suite"
)

type mockDriverQuirks struct {
	mem *memory.CheckedAllocator
}

func (q *mockDriverQuirks) SetupDriver(t *testing.T) adbc.Driver {
	q.mem = memory.NewCheckedAllocator(memory.DefaultAllocator)
	return driver.NewDriver(q.mem)
}

func (q *mockDriverQuirks) TearDownDriver(t *testing.T, _ adbc.Driver) {
	q.mem.AssertSize(t, 0)
}

func (q *mockDriverQuirks) DatabaseOptions() map[string]string {
	return map[string]string{}
}

func (q *mockDriverQuirks) Alloc() memory.Allocator                     { return q.mem }
func (q *mockDriverQuirks) BindParameter(_ int) string                  { return "?" }
func (q *mockDriverQuirks) SupportsBulkIngest(string) bool              { return true }
func (q *mockDriverQuirks) SupportsConcurrentStatements() bool          { return true }
func (q *mockDriverQuirks) SupportsCurrentCatalogSchema() bool          { return true }
func (q *mockDriverQuirks) SupportsExecuteSchema() bool                 { return true }
func (q *mockDriverQuirks) SupportsGetSetOptions() bool                 { return true }
func (q *mockDriverQuirks) SupportsPartitionedData() bool               { return false }
func (q *mockDriverQuirks) SupportsStatistics() bool                    { return false }
func (q *mockDriverQuirks) SupportsTransactions() bool                  { return true }
func (q *mockDriverQuirks) SupportsGetParameterSchema() bool            { return false }
func (q *mockDriverQuirks) SupportsDynamicParameterBinding() bool       { return false }
func (q *mockDriverQuirks) SupportsErrorIngestIncompatibleSchema() bool { return false }
func (q *mockDriverQuirks) Catalog() string                             { return "catalog" }
func (q *mockDriverQuirks) DBSchema() string                            { return "schema" }

func (q *mockDriverQuirks) GetMetadata(code adbc.InfoCode) interface{} {
	switch code {
	case adbc.InfoDriverName:
		return "ADBC Mock Driver - Go"
		// runtime/debug.ReadBuildInfo doesn't currently work for tests
	// github.com/golang/go/issues/33976
	case adbc.InfoDriverVersion:
		return "(unknown or development build)"
	case adbc.InfoDriverArrowVersion:
		return "(unknown or development build)"
	case adbc.InfoVendorVersion:
		return "(unknown or development build)"
	case adbc.InfoVendorArrowVersion:
		return "(unknown or development build)"
	case adbc.InfoDriverADBCVersion:
		return adbc.AdbcVersion1_1_0
	case adbc.InfoVendorName:
		return "Mock"
	}

	return nil
}

func withQuirks(t *testing.T, fn func(quirks *mockDriverQuirks)) {
	quirks := &mockDriverQuirks{}
	fn(quirks)
}

func TestMocks(t *testing.T) {
	withQuirks(t, func(quirks *mockDriverQuirks) {
		suite.Run(t, &MocksDriverTests{Quirks: quirks})
	})
}

type MocksDriverTests struct {
	suite.Suite

	Quirks *mockDriverQuirks

	ctx    context.Context
	driver adbc.Driver
	db     adbc.Database
	cnxn   adbc.Connection
	stmt   adbc.Statement
}

func (suite *MocksDriverTests) SetupTest() {
	var err error
	suite.ctx = context.Background()
	suite.driver = suite.Quirks.SetupDriver(suite.T())
	suite.db, err = suite.driver.NewDatabase(suite.Quirks.DatabaseOptions())
	suite.NoError(err)
	suite.cnxn, err = suite.db.Open(suite.ctx)
	suite.NoError(err)
	suite.stmt, err = suite.cnxn.NewStatement()
	suite.NoError(err)
}

func (suite *MocksDriverTests) TearDownTest() {
	suite.NoError(suite.stmt.Close())
	suite.NoError(suite.cnxn.Close())
	suite.Quirks.TearDownDriver(suite.T(), suite.driver)
	suite.cnxn = nil
	suite.NoError(suite.db.Close())
	suite.db = nil
	suite.driver = nil
}

func (suite *MocksDriverTests) TestNewDatabaseGetSetOptions() {
	key1, val1 := driver.DatabaseOptionStringOption1, "val1"
	key2, val2 := driver.DatabaseOptionStringOption2, "val2"

	db, err := suite.driver.NewDatabase(map[string]string{
		key1: val1,
		key2: val2,
	})
	suite.NoError(err)
	suite.NotNil(db)
	defer suite.NoError(db.Close())

	getSetDB, ok := db.(adbc.GetSetOptions)
	suite.True(ok)

	optVal1, err := getSetDB.GetOption(key1)
	suite.NoError(err)
	suite.Equal(optVal1, val1)
	optVal2, err := getSetDB.GetOption(key2)
	suite.NoError(err)
	suite.Equal(optVal2, val2)
}

func (suite *MocksDriverTests) TestRowCount() {
	expectedRows := 5
	query := fmt.Sprintf("%d:int8", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	recv := int64(0)
	for rdr.Next() {
		recv += rdr.Record().NumRows()
	}

	// verify that we got the expected number of rows if we sum up
	// all the rows from each record in the stream.
	suite.Equal(n, recv)
	suite.Equal(recv, int64(expectedRows))
}

func (suite *MocksDriverTests) TestIntegers() {
	expectedRows := 2
	query := fmt.Sprintf("%d:int8,uint8,int16,uint16,int32,uint32,int64,uint64", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{Name: "int8#0", Type: arrow.PrimitiveTypes.Int8},
		{Name: "uint8#1", Type: arrow.PrimitiveTypes.Uint8},
		{Name: "int16#2", Type: arrow.PrimitiveTypes.Int16},
		{Name: "uint16#3", Type: arrow.PrimitiveTypes.Uint16},
		{Name: "int32#4", Type: arrow.PrimitiveTypes.Int32},
		{Name: "uint32#5", Type: arrow.PrimitiveTypes.Uint32},
		{Name: "int64#6", Type: arrow.PrimitiveTypes.Int64},
		{Name: "uint64#7", Type: arrow.PrimitiveTypes.Uint64},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`
		[
			{
				"int16#2": 0,
				"int32#4": 0,
				"int64#6": 0,
				"int8#0": 0,
				"uint16#3": 0,
				"uint32#5": 0,
				"uint64#7": 0,
				"uint8#1": 0
			},
			{
				"int16#2": 1,
				"int32#4": 1,
				"int64#6": 1,
				"int8#0": 1,
				"uint16#3": 1,
				"uint32#5": 1,
				"uint64#7": 1,
				"uint8#1": 1
			}
		]`)),
	)

	suite.Equal(expectedSchema, rdr.Schema())
	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestSimpleList() {
	expectedRows := 1
	query := fmt.Sprintf("%d:list<int8>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "list#0",
			Type: arrow.ListOf(arrow.PrimitiveTypes.Int8),
			// Metadata: arrow.MetadataFrom(map[string]string{"list_length": "1"}),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[{"list#0":[0]}]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestNestedLists() {
	expectedRows := 2
	query := fmt.Sprintf("%d:list<3:list<4:int16>>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "list#0",
			Type: arrow.ListOf(arrow.ListOf(arrow.PrimitiveTypes.Int16)),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`
		[
			{
				"list#0": [
				[0, 1, -2, 3],
				[-4, 5, -6, 7],
				[-8, 9, -10, 11]
				]
			},
			{
				"list#0": [
				[-12, 13, -14, 15],
				[-16, 17, -18, 19],
				[-20, 21, -22, 23]
				]
			}
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestStructs() {
	expectedRows := 2
	query := fmt.Sprintf("%d:struct<list<4:int16>,int8,struct<float32>#inner>#outer", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "outer",
			Type: arrow.StructOf(
				arrow.Field{Name: "list#0", Type: arrow.ListOf(arrow.PrimitiveTypes.Int16)},
				arrow.Field{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8},
				arrow.Field{Name: "inner", Type: arrow.StructOf(
					arrow.Field{Name: "float32#2", Type: arrow.PrimitiveTypes.Float32},
				)},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{
				"outer": {
				"inner": { "float32#2": 0 },
				"list#0": [0, 1, -2, 3],
				"int8#1": 0
				}
			},
			{
				"outer": {
				"inner": { "float32#2": 0.1 },
				"list#0": [-4, 5, -6, 7],
				"int8#1": 1
				}
			}
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestAlias() {
	expectedRows := 2
	query := fmt.Sprintf("%d:i8,f,tdD", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "int8#0",
			Type: arrow.PrimitiveTypes.Int8,
		},
		{
			Name: "float32#1",
			Type: arrow.PrimitiveTypes.Float32,
		},
		{
			Name: "date32#2",
			Type: arrow.FixedWidthTypes.Date32,
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "date32#2": "1984-01-01", "float32#1": 0, "int8#0": 0 },
			{ "date32#2": "1984-01-02", "float32#1": 0.1, "int8#0": 1 }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestSparseUnion1() {
	expectedRows := 3
	query := fmt.Sprintf("%d:sparse_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "sparse_union#3",
			Type: arrow.SparseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "sparse_union#3": [0, true] },
			{ "sparse_union#3": [1, 1] },
			{ "sparse_union#3": [2, "2"] }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestSparseUnion2() {
	expectedRows := 1
	query := fmt.Sprintf("%d:sparse_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "sparse_union#3",
			Type: arrow.SparseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[{"sparse_union#3":[0,true]}]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestSparseUnion3() {
	expectedRows := 5
	query := fmt.Sprintf("%d:sparse_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "sparse_union#3",
			Type: arrow.SparseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "sparse_union#3": [0, true] },
			{ "sparse_union#3": [1, 1] },
			{ "sparse_union#3": [2, "2"] },
			{ "sparse_union#3": [0, false] },
			{ "sparse_union#3": [1, -4] }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestDenseUnion1() {
	expectedRows := 3
	query := fmt.Sprintf("%d:dense_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "dense_union#3",
			Type: arrow.DenseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "dense_union#3": [0, true] },
			{ "dense_union#3": [1, 0] },
			{ "dense_union#3": [2, "0"] }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestDenseUnion2() {
	expectedRows := 1
	query := fmt.Sprintf("%d:dense_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "dense_union#3",
			Type: arrow.DenseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[{"dense_union#3":[0,true]}]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestDenseUnion3() {
	expectedRows := 6
	query := fmt.Sprintf("%d:dense_union<bool,int8,string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{
		{
			Name: "dense_union#3",
			Type: arrow.DenseUnionOf(
				[]arrow.Field{
					{Name: "bool#0", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
					{Name: "int8#1", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
					{Name: "utf8#2", Type: arrow.BinaryTypes.String, Nullable: true},
				},
				[]arrow.UnionTypeCode{0, 1, 2},
			),
		},
	}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "dense_union#3": [0, true] },
			{ "dense_union#3": [1, 0] },
			{ "dense_union#3": [2, "0"] },
			{ "dense_union#3": [0, false] },
			{ "dense_union#3": [1, 1] },
			{ "dense_union#3": [2, "1"] }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestDictionary() {
	expectedRows := 3
	query := fmt.Sprintf("%d:dict<string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{{
		Name: "dictionary#0",
		Type: &arrow.DictionaryType{
			IndexType: arrow.PrimitiveTypes.Int32,
			ValueType: arrow.BinaryTypes.String,
		},
	}}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{"dictionary#0":"0"},
			{"dictionary#0":"1"},
			{"dictionary#0":"2"}
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestRee() {
	expectedRows := 3
	query := fmt.Sprintf("%d:ree<string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{{
		Name: "run_end_encoded#0",
		Type: arrow.RunEndEncodedOf(arrow.PrimitiveTypes.Int32,arrow.BinaryTypes.String),
	}}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "run_end_encoded#0": "1" },
			{ "run_end_encoded#0": "2" },
			{ "run_end_encoded#0": "3" }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) Test() {
	expectedRows := 3
	query := fmt.Sprintf("%d:ree<string>", expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	expectedSchema := arrow.NewSchema([]arrow.Field{{
		Name: "run_end_encoded#0",
		Type: arrow.RunEndEncodedOf(arrow.PrimitiveTypes.Int32, arrow.BinaryTypes.String),
	}}, nil)

	expectedRecords, _, err := array.RecordFromJSON(
		suite.Quirks.Alloc(),
		expectedSchema,
		bytes.NewReader([]byte(`[
			{ "run_end_encoded#0": "1" },
			{ "run_end_encoded#0": "2" },
			{ "run_end_encoded#0": "3" }
		]`)),
	)

	suite.Require().NoError(err)
	defer expectedRecords.Release()

	suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}

func (suite *MocksDriverTests) TestEverything() {
	suite.T().Skip("TODO")
	expectedRows := 6
	query := fmt.Sprintf(`%d:
		list<struct<u8,u16,u32,u64,i8,i16,i32,i64>>,
		dense_union<f16,f32,f64,bool,dict<string>#foo,list<3:z>>,
		ree<time32s>,
		t32ms, ttu, ttn,
		fixed_size_binary<5>,
		timestamp_s, timestamp_ms, timestamp_us, timestamp_ns,
		duration_s, duration_ms, duration_us, duration_ns,
		decimal128<38:10>, decimal256<76:20>,
		interval_month, interval_daytime, interval_monthdaynano
	`, expectedRows)
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	j, _ := json.Marshal(result)
	log.Println(string(j))

	suite.EqualValues(expectedRows, n)
	suite.True(rdr.Next())

	// expectedSchema := arrow.NewSchema([]arrow.Field{
	// 	{
	// 		Name: "list#8",
	// 		Type: arrow.StructOf(
	// 			arrow.Field{Name: "uint8#0", Type: arrow.PrimitiveTypes.Uint8, Nullable: true},
	// 			arrow.Field{Name: "uint16#1", Type: arrow.PrimitiveTypes.Uint16, Nullable: true},
	// 			arrow.Field{Name: "uint32#2", Type: arrow.PrimitiveTypes.Uint32, Nullable: true},
	// 			arrow.Field{Name: "uint64#3", Type: arrow.PrimitiveTypes.Uint64, Nullable: true},
	// 			arrow.Field{Name: "int8#4", Type: arrow.PrimitiveTypes.Int8, Nullable: true},
	// 			arrow.Field{Name: "int16#5", Type: arrow.PrimitiveTypes.Int16, Nullable: true},
	// 			arrow.Field{Name: "int32#6", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	// 			arrow.Field{Name: "int64#7", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	// 		),
	// 	},
	// 	{
	// 		Name: "dense_union#15",
	// 		Type: arrow.DenseUnionOf(
	// 			[]arrow.Field{
	// 				{Name: "float16#9", Type: arrow.FixedWidthTypes.Float16, Nullable: true},
	// 				{Name: "float32#10", Type: arrow.PrimitiveTypes.Float32, Nullable: true},
	// 				{Name: "float64#11", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	// 				{Name: "bool#12", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	// 				{Name: "foo", Type: &arrow.DictionaryType{
	// 					IndexType:   arrow.PrimitiveTypes.Int32,
	// 					ValueType:   arrow.BinaryTypes.String,
	// 				}},
	// 				{Name: "list#14", Type: arrow.ListOf(arrow.BinaryTypes.Binary), Nullable: true},
	// 			},
	// 			[]arrow.UnionTypeCode{0, 1, 2, 3, 4, 5},
	// 		),
	// 	},
	// 	{
	// 		Name: "run_end_encoded#16",
	// 		Type: arrow.RunEndEncodedOf(arrow.PrimitiveTypes.Int32, arrow.FixedWidthTypes.Time32s),
	// 	},
	// 	{
	// 		Name: "time32#17",
	// 		Type: arrow.FixedWidthTypes.Time32ms,
	// 	},
	// 	{
	// 		Name: "time64#18",
	// 		Type: arrow.FixedWidthTypes.Time64us,
	// 	},
	// 	{
	// 		Name: "time64#19",
	// 		Type: arrow.FixedWidthTypes.Time64ns,
	// 	},
	// 	{
	// 		Name: "fixed_size_binary#20",
	// 		Type: &arrow.FixedSizeBinaryType{ByteWidth: 5},
	// 	},
	// 	{
	// 		Name: "timestamp#21",
	// 		Type: arrow.FixedWidthTypes.Timestamp_s,
	// 	},
	// 	{
	// 		Name: "timestamp#22",
	// 		Type: arrow.FixedWidthTypes.Timestamp_ms,
	// 	},
	// 	{
	// 		Name: "timestamp#23",
	// 		Type: arrow.FixedWidthTypes.Timestamp_us,
	// 	},
	// 	{
	// 		Name: "timestamp#24",
	// 		Type: arrow.FixedWidthTypes.Timestamp_ns,
	// 	},
	// 	{
	// 		Name: "duration#25",
	// 		Type: arrow.FixedWidthTypes.Duration_s,
	// 	},
	// 	{
	// 		Name: "duration#26",
	// 		Type: arrow.FixedWidthTypes.Duration_ms,
	// 	},
	// 	{
	// 		Name: "duration#27",
	// 		Type: arrow.FixedWidthTypes.Duration_us,
	// 	},
	// 	{
	// 		Name: "duration#28",
	// 		Type: arrow.FixedWidthTypes.Duration_ns,
	// 	},
	// 	{
	// 		Name: "decimal#29",
	// 		Type: &arrow.Decimal128Type{Precision: 38, Scale: 10},
	// 	},
	// 	{
	// 		Name: "decimal256#30",
	// 		Type: &arrow.Decimal256Type{Precision: 76, Scale: 20},
	// 	},
	// 	{
	// 		Name: "month_interval#31",
	// 		Type: arrow.FixedWidthTypes.MonthInterval,
	// 	},
	// 	{
	// 		Name: "day_time_interval#32",
	// 		Type: arrow.FixedWidthTypes.DayTimeInterval,
	// 	},
	// 	{
	// 		Name: "month_day_nano_interval#33",
	// 		Type: arrow.FixedWidthTypes.MonthDayNanoInterval,
	// 	},
	// }, nil)

	// The json doesn't work :(
	// expectedRecords, _, err := array.RecordFromJSON(
	// 	suite.Quirks.Alloc(),
	// 	expectedSchema,
	// 	bytes.NewReader([]byte(``)),
	// )

	suite.Require().NoError(err)
	// defer expectedRecords.Release()

	// suite.Truef(array.RecordEqual(expectedRecords, result), "expected: %s\ngot: %s", expectedRecords, result)

	suite.False(rdr.Next())
	suite.Require().NoError(rdr.Err())
}
