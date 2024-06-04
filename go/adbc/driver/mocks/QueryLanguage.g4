grammar QueryLanguage;

query: typeSpec (',' typeSpec)* EOF;

typeSpec: primitiveType | list | struct;

list: 'list' '<' typeSpec '>';
struct: 'struct' '<' typeSpec (',' typeSpec)* '>';

primitiveType: 'int8' | 'string' | 'bool';

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

WS: [ \t\r\n]+ -> skip;