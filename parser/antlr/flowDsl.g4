grammar flowDsl;
options { tokenVocab=flowDslLexer; }

flows: flow+ EOF;

flow: FLOW ID LBRACE statement+ RBRACE;

statement:
	statementStart statementMiddle statementEnd
	| statementStart statementMiddle SEMI
	| statementMiddle statementEnd;

statementMiddle:
	component (port? data? ARROW port? component)*?;

statementStart:
	port data ARROW port?
	| CONTINUATION data? ARROW port?;

statementEnd: port? data? ARROW (port | CONTINUATION) SEMI;

component: LBRACK componentTypeName plugin? RBRACK;

componentTypeName: IDI? packageIDI;

plugin: LBRACKI pluginPart (PIPEP pluginPart)* RBRACKP;

pluginPart:
	packageIDP
	| packageIDP ASSIGN packageIDP (COMMAP packageIDP)*;

data: LPAREN packageIDI (dataSep packageIDI)* RPAREN;

packageIDI: IDI (DOTI IDI)?;

packageIDP: IDP (DOTP IDP)?;

port: NAME (COLON INT)?;

dataSep: PIPE | COMMA;
