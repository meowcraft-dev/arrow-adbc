package mocks_test

import (
	"bytes"
	"context"
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
	query := "5:int8"
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
	suite.Equal(recv, int64(5))
}

func (suite *MocksDriverTests) TestIntegers() {
	query := "2:int8,uint8,int16,uint16,int32,uint32,int64,uint64"
	suite.Require().NoError(suite.stmt.SetSqlQuery(query))
	rdr, n, err := suite.stmt.ExecuteQuery(suite.ctx)
	suite.Require().NoError(err)
	defer rdr.Release()

	result := rdr.Record()

	suite.EqualValues(2, n)
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