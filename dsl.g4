// dsl.g4 non-recursive grammar
grammar dsl;

//WHITESPACE: [ \r\n\t]+ -> skip;

// Rules start : expression EOF;

typesDeclaration: 'types:' type (',' type)*;
program: decleration (decleration)* EOF;

tuple: '(' tuplelist ')';

tuplelist: (identifier ',')+ ;


decleration:
	actionDeclaration
	| typesDeclaration
	| monitorDeclaration
	| messageDeclaration ;

actionDeclaration:
	'action' type identifier '{' statementList '}';

monitorDeclaration:
	'monitor' type identifier '{' statementList '}';

messageDeclaration: identifier ':' message;

varDeclaration: 'var' identifier '=' (tuple | DigitSeq | func);

builtins: 'rand' | 'tag' | 'store';
func: builtins '(' args ')';
args: 'type' type | identifier;

statementList: (
		(
			assignmentExpression
			| sendExpression
			| recvExpression
			| varDeclaration
			| barrier
		)
	)+;

assignmentExpression: variable Assign Wildcard;

sendExpression: (message | identifier) '=>' identifier;

recvExpression: (message | identifier) '<=' identifier;

barrier: '.';

Wildcard: '_' (DigitSeq);

Assign: '=';

message: '"' StringLiteral '"';

type: StringLiteral | primitive;
primitive: 'int';

variable: type identifier;

identifier: StringLiteral;

StringLiteral: Char+ ;

fragment Char: [a-zA-Z0-9_] ;

DigitSeq: Digit+ ;

fragment Digit: [0-9];



Comment : '//' ~[\r\n]*  -> skip ;

Whitespace
    : [ \t]+ -> channel(HIDDEN)
    ;

Newline
    : ('\r' '\n'? | '\n') -> channel(HIDDEN)
    ;