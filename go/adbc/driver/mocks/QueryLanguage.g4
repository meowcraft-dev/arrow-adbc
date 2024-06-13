grammar QueryLanguage;

query: (ROWCOUNT)? fields (',' fields)* EOF;

fields: (
		simpleTypes
		| list
		| struct
		| fixedSizeBinary
		| decimal128
		| decimal256
	) (FIELDNAME)?;
structFields: (
		simpleTypes
		| list
		| struct
		| fixedSizeBinary
		| decimal128
		| decimal256
	) (FIELDNAME)?;

innerType: simpleTypes | list | struct;

list: 'list' '<' (ROWCOUNT)? innerType '>';
struct: 'struct' '<' structFields (',' structFields)* '>';
fixedSizeBinary: 'fixed_size_binary' '<' BYTEWIDTH '>';
decimal128: 'decimal128' '<' DECIMALPS '>';
decimal256: 'decimal128' '<' DECIMALPS '>';

simpleTypes:
	'uint8'
	| 'int8'
	| 'uint16'
	| 'int16'
	| 'uint32'
	| 'int32'
	| 'uint64'
	| 'int64'
	| 'float16'
	| 'float32'
	| 'float64'
	| 'binary'
	| 'string'
	| 'date32'
	| 'date64'
	| 'bool'
	| 'time32s'
	| 'time32ms'
	| 'time64us'
	| 'time64ns'
	| 'timestamp_s'
	| 'timestamp_ms'
	| 'timestamp_us'
	| 'timestamp_ns'
	| 'duration_s'
	| 'duration_ms'
	| 'duration_us'
	| 'duration_ns'
	| 'interval_month'
	| 'interval_daytime'
	| 'interval_monthdaynano'
	| 'null';

DECIMALPS: POSITIVE_INT ':' POSITIVE_INT;

ROWCOUNT: POSITIVE_INT ':';
BYTEWIDTH: POSITIVE_INT;
FIELDNAME: '#' IDENTIFIER;

WS: [ \t\r\n]+ -> skip;

fragment IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
fragment POSITIVE_INT: [1-9][0-9]*;