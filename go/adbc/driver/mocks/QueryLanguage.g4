grammar QueryLanguage;

query: (COUNT)? topLevelField (',' topLevelField)* EOF;

topLevelField: type (FIELD_NAME)?;
structField: type (FIELD_NAME)?;
unionField: type (FIELD_NAME)?;

type:
	simpleTypes
	| list
	| struct
	| dictionary
	| fixedSizeBinary
	| decimal128
	| decimal256
	| union
	| runEndEncoded;

union: (SPARSE_UNION | DENSE_UNION) '<' unionField (
		',' unionField
	)* '>';
dictionary: ('dict' | 'dictionary') '<' type '>';
list: 'list' '<' (COUNT)? type '>';
struct: 'struct' '<' structField (',' structField)* '>';
fixedSizeBinary: 'fixed_size_binary' '<' BYTE_WIDTH '>';
decimal128: 'decimal128' '<' DECIMAL_PS '>';
decimal256: 'decimal256' '<' DECIMAL_PS '>';
runEndEncoded: ('ree' | 'run_end_encoded') '<' type '>';

simpleTypes:
	UINT8
	| INT8
	| UINT16
	| INT16
	| UINT32
	| INT32
	| UINT64
	| INT64
	| FLOAT16
	| FLOAT32
	| FLOAT64
	| BINARY
	| STRING
	| DATE32
	| DATE64
	| BOOL
	| TIME32S
	| TIME32MS
	| TIME64US
	| TIME64NS
	| TIMESTAMP_S
	| TIMESTAMP_MS
	| TIMESTAMP_US
	| TIMESTAMP_NS
	| DURATION_S
	| DURATION_MS
	| DURATION_US
	| DURATION_NS
	| INTERVAL_MONTH
	| INTERVAL_DAYTIME
	| INTERVAL_MONTHDAYNANO
	| NULL;
NULL: 'null' | 'n';
BOOL: 'bool' | 'boolean' | 'b';
UINT8: 'uint8' | 'u8' | 'C';
INT8: 'int8' | 'i8' | 'c';
UINT16: 'uint16' | 'u16' | 'S';
INT16: 'int16' | 'i16' | 's';
UINT32: 'uint32' | 'u32' | 'I';
INT32: 'int32' | 'i32' | 'i';
UINT64: 'uint64' | 'u64' | 'L';
INT64: 'int64' | 'i64' | 'l';
FLOAT16: 'float16' | 'f16' | 'e';
FLOAT32: 'float32' | 'f32' | 'f';
FLOAT64: 'float64' | 'f64' | 'g';
BINARY: 'binary' | 'z';
STRING: 'string' | 'str';
DATE32: 'date32' | 'd32' | 'tdD';
DATE64: 'date64' | 'd64' | 'tdm';
TIME32S: 'time32s' | 't32s' | 'tts';
TIME32MS: 'time32ms' | 't32ms' | 'ttm';
TIME64US: 'time64us' | 't64us' | 'ttu';
TIME64NS: 'time64ns' | 't64ns' | 'ttn';
TIMESTAMP_S: 'timestamp_s';
TIMESTAMP_MS: 'timestamp_ms';
TIMESTAMP_US: 'timestamp_us';
TIMESTAMP_NS: 'timestamp_ns';
DURATION_S: 'duration_s';
DURATION_MS: 'duration_ms';
DURATION_US: 'duration_us';
DURATION_NS: 'duration_ns';
INTERVAL_MONTH: 'interval_month';
INTERVAL_DAYTIME: 'interval_daytime';
INTERVAL_MONTHDAYNANO: 'interval_monthdaynano';

SPARSE_UNION: 'sparse_union';
DENSE_UNION: 'dense_union';
DECIMAL_PS: INTEGER ':' INTEGER;
COUNT: INTEGER ':';
FIELD_NAME: '#' IDENTIFIER;
UNION_VALUE_NAME: IDENTIFIER;
BYTE_WIDTH: INTEGER;
DICT_ENTRY: INTEGER;

WS: [ \t\r\n]+ -> skip;

fragment IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
fragment INTEGER: [1-9][0-9]*;