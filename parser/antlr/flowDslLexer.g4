lexer grammar flowDslLexer;

IMPORT: 'import';

FLOW: 'flow';

LBRACE: '{' -> pushMode(F);

ID: [a-zA-Z_] [a-zA-Z0-9_]*;

BLOCK_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);

LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);

WS:
	[\p{White_Space}] -> channel(HIDDEN); // match all Unicode whitespace


//
// mode F (for parsing flows and imports)

mode F;

ARROW: '->';

SEMI: ';';

COLON: ':';

STRING: '"' .*? '"';

NAME: [a-z] [a-zA-Z0-9]*;

INT: [0-9]+;

LPAREN: '(' -> pushMode(I);

LBRACK: '[' -> pushMode(I);

RBRACE: '}' -> popMode;

BLOCK_COMMENTF: '/*' .*? '*/' -> channel(HIDDEN);

LINE_COMMENTF: '//' ~[\r\n]* -> channel(HIDDEN);

WSF:
	[\p{White_Space}] -> channel(HIDDEN); // match all Unicode whitespace


//
// mode I (for parsing inner constructs: components and data)

mode I;

COMMA: ',';

PIPE: '|';

DOTI: '.';

RPAREN: ')' -> popMode;

LBRACKI: '[' -> pushMode(P);

RBRACK: ']' -> popMode;

IDI: [a-zA-Z_] [a-zA-Z0-9_]*;

BLOCK_COMMENTI: '/*' .*? '*/' -> channel(HIDDEN);

LINE_COMMENTI: '//' ~[\r\n]* -> channel(HIDDEN);

WSI:
	[\p{White_Space}] -> channel(HIDDEN); // match all Unicode whitespace


//
// mode P (for parsing plugins)

mode P;

ASSIGN: '=';

COMMAP: ',';

DOTP: '.';

PIPEP: '|';

RBRACKP: ']' -> popMode;

IDP: [a-zA-Z_] [a-zA-Z0-9_]*;

BLOCK_COMMENTP: '/*' .*? '*/' -> channel(HIDDEN);

LINE_COMMENTP: '//' ~[\r\n]* -> channel(HIDDEN);

WSP:
	[\p{White_Space}] -> channel(HIDDEN); // match all Unicode whitespace
