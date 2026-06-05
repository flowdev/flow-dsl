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
	FirstComponent=component ArrowComponents+=arrowComponent*?;

statementStart:
	StartPort=port AllArrData=allData ARROW DstPort=port?;

statementEnd: SrcPort=port? AllArrData=allData? ARROW EndPort=port; // SEMI;

arrowComponent: SrcPort=port? AllArrData=allData? ARROW DstPort=port? DstComponent=component;

component: LBRACK Core=componentTypeName AllPlugins=plugin? RBRACK;

componentTypeName: Name=IDI TypPack=IDI? (DOTI TypName=IDI)?;

plugin: LBRACKI PluginGroups+=pluginPart (PIPEP PluginGroups+=pluginPart)* RBRACKP;

pluginPart:
	Interface=packageIDP
	| Interface=packageIDP ASSIGN Plugins+=packageIDP (COMMAP Plugins+=packageIDP)*;

allData: LPAREN Datas+=data? (COMMA Datas+=data)* RPAREN;

data: Name=IDI TypPack=IDI (DOTI TypName=IDI)?;

packageIDI: ID1=IDI;

packageIDP: ID1=IDP (DOTP ID2=IDP)?;

port: Name=NAME (COLON Num=INT)?;
