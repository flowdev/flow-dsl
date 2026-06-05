grammar flowDsl;
options { tokenVocab=flowDslLexer; }

flows: Import=imp? Flows+=flow+ EOF;

imp: IMPORT LBRACE (Imports+=STRING SEMI)+ RBRACE;

flow: FLOW Name=ID LBRACE (Statements+=statement)+ RBRACE;

statement:
	statementStart? statementMiddle statementEnd? SEMI;
//	statementStart statementMiddle statementEnd
//	| statementStart statementMiddle SEMI
//	| statementMiddle statementEnd
//	| statementMiddle SEMI;

statementMiddle:
	Component=component ArrowComponents+=arrowComponent*?;

statementStart:
	StartPort=port AllStartData=data ARROW DstPort=port?;

statementEnd: SrcPort=port? AllEndData=data? ARROW EndPort=port; // SEMI;

arrowComponent: SrcPort=port? AllData=data? ARROW DstPort=port? DstComponent=component;

component: LBRACK Core=componentTypeName AllPlugins=plugin? RBRACK;

componentTypeName: Name=IDI? Typ=packageIDI;

plugin: LBRACKI PluginGroups+=pluginPart (PIPEP PluginGroups+=pluginPart)* RBRACKP;

pluginPart:
	Interface=packageIDP
	| Interface=packageIDP ASSIGN Plugins+=packageIDP (COMMAP Plugins+=packageIDP)*;

data: LPAREN Datas+=packageIDI? (dataSep Datas+=packageIDI)* RPAREN;

packageIDI: ID1=IDI (DOTI ID2=IDI)?;

packageIDP: ID1=IDP (DOTP ID2=IDP)?;

port: Name=NAME (COLON Num=INT)?;

dataSep: PIPE | COMMA;
