grammar flowDsl;
options { tokenVocab=flowDslLexer; }

flowFile: AllImports=imports? Flows+=flow+ EOF;

imports: IMPORT LBRACE (Imports+=STRING SEMI)+ RBRACE;

flow: FLOW Name=ID LBRACE (Statements+=statement)+ RBRACE;

statement:
	statementStart? statementMiddle statementEnd? SEMIF;

statementStart:
	StartPort=PORT AllArrData=allData ARROW DstPort=PORT?;

statementMiddle:
	FirstComponent=component ArrowComponents+=arrowComponent*?;

statementEnd: SrcPort=PORT? AllArrData=allData? ARROW EndPort=PORT;

arrowComponent: SrcPort=PORT? AllArrData=allData? ARROW DstPort=PORT? DstComponent=component;

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
