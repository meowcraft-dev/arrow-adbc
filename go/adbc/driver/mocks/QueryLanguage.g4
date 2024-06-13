grammar QueryLanguage;

query: (ROWCOUNT)? fields (',' fields)* EOF;

fields: (simpleTypes | list | struct) (FIELDNAME)?;

innerType: simpleTypes | list | struct;

list: 'list' '<' (ROWCOUNT)? innerType '>';
struct: 'struct' '<' innerType (',' innerType)* '>';

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

WS: [ \t\r\n]+ -> skip;

ROWCOUNT: [1-9][0-9]* ':';
FIELDNAME: '#' IDENTIFIER;

fragment IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;