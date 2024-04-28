// Code generated by _tmpl/utils.h.tmpl. DO NOT EDIT.

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

// clang-format off
//go:build driverlib
// clang-format on

#pragma once

#include <stdlib.h>
#include "../../drivermgr/adbc.h"

struct AdbcError* MocksErrorFromArrayStream(struct ArrowArrayStream*, AdbcStatusCode*);
AdbcStatusCode MocksDatabaseGetOption(struct AdbcDatabase*, const char*, char*, size_t*,
                                      struct AdbcError*);
AdbcStatusCode MocksDatabaseGetOptionBytes(struct AdbcDatabase*, const char*, uint8_t*,
                                           size_t*, struct AdbcError*);
AdbcStatusCode MocksDatabaseGetOptionDouble(struct AdbcDatabase*, const char*, double*,
                                            struct AdbcError*);
AdbcStatusCode MocksDatabaseGetOptionInt(struct AdbcDatabase*, const char*, int64_t*,
                                         struct AdbcError*);
AdbcStatusCode MocksDatabaseInit(struct AdbcDatabase* db, struct AdbcError* err);
AdbcStatusCode MocksDatabaseNew(struct AdbcDatabase* db, struct AdbcError* err);
AdbcStatusCode MocksDatabaseRelease(struct AdbcDatabase* db, struct AdbcError* err);
AdbcStatusCode MocksDatabaseSetOption(struct AdbcDatabase* db, const char* key,
                                      const char* value, struct AdbcError* err);
AdbcStatusCode MocksDatabaseSetOptionBytes(struct AdbcDatabase*, const char*,
                                           const uint8_t*, size_t, struct AdbcError*);
AdbcStatusCode MocksDatabaseSetOptionDouble(struct AdbcDatabase*, const char*, double,
                                            struct AdbcError*);
AdbcStatusCode MocksDatabaseSetOptionInt(struct AdbcDatabase*, const char*, int64_t,
                                         struct AdbcError*);

AdbcStatusCode MocksConnectionCancel(struct AdbcConnection*, struct AdbcError*);
AdbcStatusCode MocksConnectionCommit(struct AdbcConnection* cnxn, struct AdbcError* err);
AdbcStatusCode MocksConnectionGetInfo(struct AdbcConnection* cnxn, const uint32_t* codes,
                                      size_t len, struct ArrowArrayStream* out,
                                      struct AdbcError* err);
AdbcStatusCode MocksConnectionGetObjects(struct AdbcConnection* cnxn, int depth,
                                         const char* catalog, const char* dbSchema,
                                         const char* tableName, const char** tableType,
                                         const char* columnName,
                                         struct ArrowArrayStream* out,
                                         struct AdbcError* err);
AdbcStatusCode MocksConnectionGetOption(struct AdbcConnection*, const char*, char*,
                                        size_t*, struct AdbcError*);
AdbcStatusCode MocksConnectionGetOptionBytes(struct AdbcConnection*, const char*,
                                             uint8_t*, size_t*, struct AdbcError*);
AdbcStatusCode MocksConnectionGetOptionDouble(struct AdbcConnection*, const char*,
                                              double*, struct AdbcError*);
AdbcStatusCode MocksConnectionGetOptionInt(struct AdbcConnection*, const char*, int64_t*,
                                           struct AdbcError*);
AdbcStatusCode MocksConnectionGetStatistics(struct AdbcConnection*, const char*,
                                            const char*, const char*, char,
                                            struct ArrowArrayStream*, struct AdbcError*);
AdbcStatusCode MocksConnectionGetStatisticNames(struct AdbcConnection*,
                                                struct ArrowArrayStream*,
                                                struct AdbcError*);
AdbcStatusCode MocksConnectionGetTableSchema(struct AdbcConnection* cnxn,
                                             const char* catalog, const char* dbSchema,
                                             const char* tableName,
                                             struct ArrowSchema* schema,
                                             struct AdbcError* err);
AdbcStatusCode MocksConnectionGetTableTypes(struct AdbcConnection* cnxn,
                                            struct ArrowArrayStream* out,
                                            struct AdbcError* err);
AdbcStatusCode MocksConnectionInit(struct AdbcConnection* cnxn, struct AdbcDatabase* db,
                                   struct AdbcError* err);
AdbcStatusCode MocksConnectionNew(struct AdbcConnection* cnxn, struct AdbcError* err);
AdbcStatusCode MocksConnectionReadPartition(struct AdbcConnection* cnxn,
                                            const uint8_t* serialized,
                                            size_t serializedLen,
                                            struct ArrowArrayStream* out,
                                            struct AdbcError* err);
AdbcStatusCode MocksConnectionRelease(struct AdbcConnection* cnxn, struct AdbcError* err);
AdbcStatusCode MocksConnectionRollback(struct AdbcConnection* cnxn,
                                       struct AdbcError* err);
AdbcStatusCode MocksConnectionSetOption(struct AdbcConnection* cnxn, const char* key,
                                        const char* val, struct AdbcError* err);
AdbcStatusCode MocksConnectionSetOptionBytes(struct AdbcConnection*, const char*,
                                             const uint8_t*, size_t, struct AdbcError*);
AdbcStatusCode MocksConnectionSetOptionDouble(struct AdbcConnection*, const char*, double,
                                              struct AdbcError*);
AdbcStatusCode MocksConnectionSetOptionInt(struct AdbcConnection*, const char*, int64_t,
                                           struct AdbcError*);

AdbcStatusCode MocksStatementBind(struct AdbcStatement* stmt, struct ArrowArray* values,
                                  struct ArrowSchema* schema, struct AdbcError* err);
AdbcStatusCode MocksStatementBindStream(struct AdbcStatement* stmt,
                                        struct ArrowArrayStream* stream,
                                        struct AdbcError* err);
AdbcStatusCode MocksStatementCancel(struct AdbcStatement*, struct AdbcError*);
AdbcStatusCode MocksStatementExecuteQuery(struct AdbcStatement* stmt,
                                          struct ArrowArrayStream* out, int64_t* affected,
                                          struct AdbcError* err);
AdbcStatusCode MocksStatementExecutePartitions(struct AdbcStatement* stmt,
                                               struct ArrowSchema* schema,
                                               struct AdbcPartitions* partitions,
                                               int64_t* affected, struct AdbcError* err);
AdbcStatusCode MocksStatementExecuteSchema(struct AdbcStatement*, struct ArrowSchema*,
                                           struct AdbcError*);
AdbcStatusCode MocksStatementGetOption(struct AdbcStatement*, const char*, char*, size_t*,
                                       struct AdbcError*);
AdbcStatusCode MocksStatementGetOptionBytes(struct AdbcStatement*, const char*, uint8_t*,
                                            size_t*, struct AdbcError*);
AdbcStatusCode MocksStatementGetOptionDouble(struct AdbcStatement*, const char*, double*,
                                             struct AdbcError*);
AdbcStatusCode MocksStatementGetOptionInt(struct AdbcStatement*, const char*, int64_t*,
                                          struct AdbcError*);
AdbcStatusCode MocksStatementGetParameterSchema(struct AdbcStatement* stmt,
                                                struct ArrowSchema* schema,
                                                struct AdbcError* err);
AdbcStatusCode MocksStatementNew(struct AdbcConnection* cnxn, struct AdbcStatement* stmt,
                                 struct AdbcError* err);
AdbcStatusCode MocksStatementPrepare(struct AdbcStatement* stmt, struct AdbcError* err);
AdbcStatusCode MocksStatementRelease(struct AdbcStatement* stmt, struct AdbcError* err);
AdbcStatusCode MocksStatementSetOption(struct AdbcStatement* stmt, const char* key,
                                       const char* value, struct AdbcError* err);
AdbcStatusCode MocksStatementSetOptionBytes(struct AdbcStatement*, const char*,
                                            const uint8_t*, size_t, struct AdbcError*);
AdbcStatusCode MocksStatementSetOptionDouble(struct AdbcStatement*, const char*, double,
                                             struct AdbcError*);
AdbcStatusCode MocksStatementSetOptionInt(struct AdbcStatement*, const char*, int64_t,
                                          struct AdbcError*);
AdbcStatusCode MocksStatementSetSqlQuery(struct AdbcStatement* stmt, const char* query,
                                         struct AdbcError* err);
AdbcStatusCode MocksStatementSetSubstraitPlan(struct AdbcStatement* stmt,
                                              const uint8_t* plan, size_t length,
                                              struct AdbcError* err);

AdbcStatusCode MocksDriverInit(int version, void* rawDriver, struct AdbcError* err);

static inline void MockserrRelease(struct AdbcError* error) {
  if (error->release) {
    error->release(error);
    error->release = NULL;
  }
}

void Mocks_release_error(struct AdbcError* error);

struct MocksError {
  char* message;
  char** keys;
  uint8_t** values;
  size_t* lengths;
  int count;
};

void MocksReleaseErrWithDetails(struct AdbcError* error);

int MocksErrorGetDetailCount(const struct AdbcError* error);
struct AdbcErrorDetail MocksErrorGetDetail(const struct AdbcError* error, int index);

int MocksArrayStreamGetSchemaTrampoline(struct ArrowArrayStream* stream,
                                        struct ArrowSchema* out);
int MocksArrayStreamGetNextTrampoline(struct ArrowArrayStream* stream,
                                      struct ArrowArray* out);
