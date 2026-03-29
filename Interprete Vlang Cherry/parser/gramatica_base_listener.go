// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr4-go/antlr/v4"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasegramaticaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasegramaticaListener) ExitProg(ctx *ProgContext) {}

// EnterDeclaracionStruct is called when production declaracionStruct is entered.
func (s *BasegramaticaListener) EnterDeclaracionStruct(ctx *DeclaracionStructContext) {}

// ExitDeclaracionStruct is called when production declaracionStruct is exited.
func (s *BasegramaticaListener) ExitDeclaracionStruct(ctx *DeclaracionStructContext) {}

// EnterFuncionStruct_declaracion is called when production funcionStruct_declaracion is entered.
func (s *BasegramaticaListener) EnterFuncionStruct_declaracion(ctx *FuncionStruct_declaracionContext) {
}

// ExitFuncionStruct_declaracion is called when production funcionStruct_declaracion is exited.
func (s *BasegramaticaListener) ExitFuncionStruct_declaracion(ctx *FuncionStruct_declaracionContext) {
}

// EnterDecFunc1 is called when production DecFunc1 is entered.
func (s *BasegramaticaListener) EnterDecFunc1(ctx *DecFunc1Context) {}

// ExitDecFunc1 is called when production DecFunc1 is exited.
func (s *BasegramaticaListener) ExitDecFunc1(ctx *DecFunc1Context) {}

// EnterDecFunc2 is called when production DecFunc2 is entered.
func (s *BasegramaticaListener) EnterDecFunc2(ctx *DecFunc2Context) {}

// ExitDecFunc2 is called when production DecFunc2 is exited.
func (s *BasegramaticaListener) ExitDecFunc2(ctx *DecFunc2Context) {}

// EnterLista_parametros is called when production lista_parametros is entered.
func (s *BasegramaticaListener) EnterLista_parametros(ctx *Lista_parametrosContext) {}

// ExitLista_parametros is called when production lista_parametros is exited.
func (s *BasegramaticaListener) ExitLista_parametros(ctx *Lista_parametrosContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BasegramaticaListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BasegramaticaListener) ExitParametro(ctx *ParametroContext) {}

// EnterLista_sentencias is called when production lista_sentencias is entered.
func (s *BasegramaticaListener) EnterLista_sentencias(ctx *Lista_sentenciasContext) {}

// ExitLista_sentencias is called when production lista_sentencias is exited.
func (s *BasegramaticaListener) ExitLista_sentencias(ctx *Lista_sentenciasContext) {}

// EnterSentencia is called when production sentencia is entered.
func (s *BasegramaticaListener) EnterSentencia(ctx *SentenciaContext) {}

// ExitSentencia is called when production sentencia is exited.
func (s *BasegramaticaListener) ExitSentencia(ctx *SentenciaContext) {}

// EnterSentenciaBreak is called when production sentenciaBreak is entered.
func (s *BasegramaticaListener) EnterSentenciaBreak(ctx *SentenciaBreakContext) {}

// ExitSentenciaBreak is called when production sentenciaBreak is exited.
func (s *BasegramaticaListener) ExitSentenciaBreak(ctx *SentenciaBreakContext) {}

// EnterSentenciaReturn is called when production sentenciaReturn is entered.
func (s *BasegramaticaListener) EnterSentenciaReturn(ctx *SentenciaReturnContext) {}

// ExitSentenciaReturn is called when production sentenciaReturn is exited.
func (s *BasegramaticaListener) ExitSentenciaReturn(ctx *SentenciaReturnContext) {}

// EnterLlamadafuncion is called when production llamadafuncion is entered.
func (s *BasegramaticaListener) EnterLlamadafuncion(ctx *LlamadafuncionContext) {}

// ExitLlamadafuncion is called when production llamadafuncion is exited.
func (s *BasegramaticaListener) ExitLlamadafuncion(ctx *LlamadafuncionContext) {}

// EnterSentBreak is called when production SentBreak is entered.
func (s *BasegramaticaListener) EnterSentBreak(ctx *SentBreakContext) {}

// ExitSentBreak is called when production SentBreak is exited.
func (s *BasegramaticaListener) ExitSentBreak(ctx *SentBreakContext) {}

// EnterSentContinue is called when production SentContinue is entered.
func (s *BasegramaticaListener) EnterSentContinue(ctx *SentContinueContext) {}

// ExitSentContinue is called when production SentContinue is exited.
func (s *BasegramaticaListener) ExitSentContinue(ctx *SentContinueContext) {}

// EnterSentReturn is called when production SentReturn is entered.
func (s *BasegramaticaListener) EnterSentReturn(ctx *SentReturnContext) {}

// ExitSentReturn is called when production SentReturn is exited.
func (s *BasegramaticaListener) ExitSentReturn(ctx *SentReturnContext) {}

// EnterSentIf1 is called when production SentIf1 is entered.
func (s *BasegramaticaListener) EnterSentIf1(ctx *SentIf1Context) {}

// ExitSentIf1 is called when production SentIf1 is exited.
func (s *BasegramaticaListener) ExitSentIf1(ctx *SentIf1Context) {}

// EnterSentIf2 is called when production SentIf2 is entered.
func (s *BasegramaticaListener) EnterSentIf2(ctx *SentIf2Context) {}

// ExitSentIf2 is called when production SentIf2 is exited.
func (s *BasegramaticaListener) ExitSentIf2(ctx *SentIf2Context) {}

// EnterSentIf3 is called when production SentIf3 is entered.
func (s *BasegramaticaListener) EnterSentIf3(ctx *SentIf3Context) {}

// ExitSentIf3 is called when production SentIf3 is exited.
func (s *BasegramaticaListener) ExitSentIf3(ctx *SentIf3Context) {}

// EnterSentenciaSwitch is called when production sentenciaSwitch is entered.
func (s *BasegramaticaListener) EnterSentenciaSwitch(ctx *SentenciaSwitchContext) {}

// ExitSentenciaSwitch is called when production sentenciaSwitch is exited.
func (s *BasegramaticaListener) ExitSentenciaSwitch(ctx *SentenciaSwitchContext) {}

// EnterBloqueCase is called when production bloqueCase is entered.
func (s *BasegramaticaListener) EnterBloqueCase(ctx *BloqueCaseContext) {}

// ExitBloqueCase is called when production bloqueCase is exited.
func (s *BasegramaticaListener) ExitBloqueCase(ctx *BloqueCaseContext) {}

// EnterBloqueDefault is called when production bloqueDefault is entered.
func (s *BasegramaticaListener) EnterBloqueDefault(ctx *BloqueDefaultContext) {}

// ExitBloqueDefault is called when production bloqueDefault is exited.
func (s *BasegramaticaListener) ExitBloqueDefault(ctx *BloqueDefaultContext) {}

// EnterSentenciaPrintln is called when production sentenciaPrintln is entered.
func (s *BasegramaticaListener) EnterSentenciaPrintln(ctx *SentenciaPrintlnContext) {}

// ExitSentenciaPrintln is called when production sentenciaPrintln is exited.
func (s *BasegramaticaListener) ExitSentenciaPrintln(ctx *SentenciaPrintlnContext) {}

// EnterAsigSimple is called when production AsigSimple is entered.
func (s *BasegramaticaListener) EnterAsigSimple(ctx *AsigSimpleContext) {}

// ExitAsigSimple is called when production AsigSimple is exited.
func (s *BasegramaticaListener) ExitAsigSimple(ctx *AsigSimpleContext) {}

// EnterAsigSumRest is called when production AsigSumRest is entered.
func (s *BasegramaticaListener) EnterAsigSumRest(ctx *AsigSumRestContext) {}

// ExitAsigSumRest is called when production AsigSumRest is exited.
func (s *BasegramaticaListener) ExitAsigSumRest(ctx *AsigSumRestContext) {}

// EnterAsignacion_acceso_slice1 is called when production asignacion_acceso_slice1 is entered.
func (s *BasegramaticaListener) EnterAsignacion_acceso_slice1(ctx *Asignacion_acceso_slice1Context) {}

// ExitAsignacion_acceso_slice1 is called when production asignacion_acceso_slice1 is exited.
func (s *BasegramaticaListener) ExitAsignacion_acceso_slice1(ctx *Asignacion_acceso_slice1Context) {}

// EnterDecSlice1 is called when production DecSlice1 is entered.
func (s *BasegramaticaListener) EnterDecSlice1(ctx *DecSlice1Context) {}

// ExitDecSlice1 is called when production DecSlice1 is exited.
func (s *BasegramaticaListener) ExitDecSlice1(ctx *DecSlice1Context) {}

// EnterDecSlice2 is called when production DecSlice2 is entered.
func (s *BasegramaticaListener) EnterDecSlice2(ctx *DecSlice2Context) {}

// ExitDecSlice2 is called when production DecSlice2 is exited.
func (s *BasegramaticaListener) ExitDecSlice2(ctx *DecSlice2Context) {}

// EnterAsigacionSlice is called when production AsigacionSlice is entered.
func (s *BasegramaticaListener) EnterAsigacionSlice(ctx *AsigacionSliceContext) {}

// ExitAsigacionSlice is called when production AsigacionSlice is exited.
func (s *BasegramaticaListener) ExitAsigacionSlice(ctx *AsigacionSliceContext) {}

// EnterDec1 is called when production Dec1 is entered.
func (s *BasegramaticaListener) EnterDec1(ctx *Dec1Context) {}

// ExitDec1 is called when production Dec1 is exited.
func (s *BasegramaticaListener) ExitDec1(ctx *Dec1Context) {}

// EnterDec2 is called when production Dec2 is entered.
func (s *BasegramaticaListener) EnterDec2(ctx *Dec2Context) {}

// ExitDec2 is called when production Dec2 is exited.
func (s *BasegramaticaListener) ExitDec2(ctx *Dec2Context) {}

// EnterDec3 is called when production Dec3 is entered.
func (s *BasegramaticaListener) EnterDec3(ctx *Dec3Context) {}

// ExitDec3 is called when production Dec3 is exited.
func (s *BasegramaticaListener) ExitDec3(ctx *Dec3Context) {}

// EnterExprOpAritmetica is called when production ExprOpAritmetica is entered.
func (s *BasegramaticaListener) EnterExprOpAritmetica(ctx *ExprOpAritmeticaContext) {}

// ExitExprOpAritmetica is called when production ExprOpAritmetica is exited.
func (s *BasegramaticaListener) ExitExprOpAritmetica(ctx *ExprOpAritmeticaContext) {}

// EnterExprBoolFalse is called when production ExprBoolFalse is entered.
func (s *BasegramaticaListener) EnterExprBoolFalse(ctx *ExprBoolFalseContext) {}

// ExitExprBoolFalse is called when production ExprBoolFalse is exited.
func (s *BasegramaticaListener) ExitExprBoolFalse(ctx *ExprBoolFalseContext) {}

// EnterExprIdentificador is called when production ExprIdentificador is entered.
func (s *BasegramaticaListener) EnterExprIdentificador(ctx *ExprIdentificadorContext) {}

// ExitExprIdentificador is called when production ExprIdentificador is exited.
func (s *BasegramaticaListener) ExitExprIdentificador(ctx *ExprIdentificadorContext) {}

// EnterExprRune is called when production ExprRune is entered.
func (s *BasegramaticaListener) EnterExprRune(ctx *ExprRuneContext) {}

// ExitExprRune is called when production ExprRune is exited.
func (s *BasegramaticaListener) ExitExprRune(ctx *ExprRuneContext) {}

// EnterExprFloat is called when production ExprFloat is entered.
func (s *BasegramaticaListener) EnterExprFloat(ctx *ExprFloatContext) {}

// ExitExprFloat is called when production ExprFloat is exited.
func (s *BasegramaticaListener) ExitExprFloat(ctx *ExprFloatContext) {}

// EnterExprNot is called when production ExprNot is entered.
func (s *BasegramaticaListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production ExprNot is exited.
func (s *BasegramaticaListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprNegUnaria is called when production ExprNegUnaria is entered.
func (s *BasegramaticaListener) EnterExprNegUnaria(ctx *ExprNegUnariaContext) {}

// ExitExprNegUnaria is called when production ExprNegUnaria is exited.
func (s *BasegramaticaListener) ExitExprNegUnaria(ctx *ExprNegUnariaContext) {}

// EnterExprNil is called when production ExprNil is entered.
func (s *BasegramaticaListener) EnterExprNil(ctx *ExprNilContext) {}

// ExitExprNil is called when production ExprNil is exited.
func (s *BasegramaticaListener) ExitExprNil(ctx *ExprNilContext) {}

// EnterExprMultDivMod is called when production ExprMultDivMod is entered.
func (s *BasegramaticaListener) EnterExprMultDivMod(ctx *ExprMultDivModContext) {}

// ExitExprMultDivMod is called when production ExprMultDivMod is exited.
func (s *BasegramaticaListener) ExitExprMultDivMod(ctx *ExprMultDivModContext) {}

// EnterExprAccesoMatriz is called when production ExprAccesoMatriz is entered.
func (s *BasegramaticaListener) EnterExprAccesoMatriz(ctx *ExprAccesoMatrizContext) {}

// ExitExprAccesoMatriz is called when production ExprAccesoMatriz is exited.
func (s *BasegramaticaListener) ExitExprAccesoMatriz(ctx *ExprAccesoMatrizContext) {}

// EnterExprAccesoSlice is called when production ExprAccesoSlice is entered.
func (s *BasegramaticaListener) EnterExprAccesoSlice(ctx *ExprAccesoSliceContext) {}

// ExitExprAccesoSlice is called when production ExprAccesoSlice is exited.
func (s *BasegramaticaListener) ExitExprAccesoSlice(ctx *ExprAccesoSliceContext) {}

// EnterExprSliceLiteral is called when production ExprSliceLiteral is entered.
func (s *BasegramaticaListener) EnterExprSliceLiteral(ctx *ExprSliceLiteralContext) {}

// ExitExprSliceLiteral is called when production ExprSliceLiteral is exited.
func (s *BasegramaticaListener) ExitExprSliceLiteral(ctx *ExprSliceLiteralContext) {}

// EnterExprBoolTrue is called when production ExprBoolTrue is entered.
func (s *BasegramaticaListener) EnterExprBoolTrue(ctx *ExprBoolTrueContext) {}

// ExitExprBoolTrue is called when production ExprBoolTrue is exited.
func (s *BasegramaticaListener) ExitExprBoolTrue(ctx *ExprBoolTrueContext) {}

// EnterExprLlamadaFuncion is called when production ExprLlamadaFuncion is entered.
func (s *BasegramaticaListener) EnterExprLlamadaFuncion(ctx *ExprLlamadaFuncionContext) {}

// ExitExprLlamadaFuncion is called when production ExprLlamadaFuncion is exited.
func (s *BasegramaticaListener) ExitExprLlamadaFuncion(ctx *ExprLlamadaFuncionContext) {}

// EnterExprLogica is called when production ExprLogica is entered.
func (s *BasegramaticaListener) EnterExprLogica(ctx *ExprLogicaContext) {}

// ExitExprLogica is called when production ExprLogica is exited.
func (s *BasegramaticaListener) ExitExprLogica(ctx *ExprLogicaContext) {}

// EnterExprIndexOf is called when production ExprIndexOf is entered.
func (s *BasegramaticaListener) EnterExprIndexOf(ctx *ExprIndexOfContext) {}

// ExitExprIndexOf is called when production ExprIndexOf is exited.
func (s *BasegramaticaListener) ExitExprIndexOf(ctx *ExprIndexOfContext) {}

// EnterExprAccesoInstanciaAtributo is called when production ExprAccesoInstanciaAtributo is entered.
func (s *BasegramaticaListener) EnterExprAccesoInstanciaAtributo(ctx *ExprAccesoInstanciaAtributoContext) {
}

// ExitExprAccesoInstanciaAtributo is called when production ExprAccesoInstanciaAtributo is exited.
func (s *BasegramaticaListener) ExitExprAccesoInstanciaAtributo(ctx *ExprAccesoInstanciaAtributoContext) {
}

// EnterExprRelacional is called when production ExprRelacional is entered.
func (s *BasegramaticaListener) EnterExprRelacional(ctx *ExprRelacionalContext) {}

// ExitExprRelacional is called when production ExprRelacional is exited.
func (s *BasegramaticaListener) ExitExprRelacional(ctx *ExprRelacionalContext) {}

// EnterExprString is called when production ExprString is entered.
func (s *BasegramaticaListener) EnterExprString(ctx *ExprStringContext) {}

// ExitExprString is called when production ExprString is exited.
func (s *BasegramaticaListener) ExitExprString(ctx *ExprStringContext) {}

// EnterExprInt is called when production ExprInt is entered.
func (s *BasegramaticaListener) EnterExprInt(ctx *ExprIntContext) {}

// ExitExprInt is called when production ExprInt is exited.
func (s *BasegramaticaListener) ExitExprInt(ctx *ExprIntContext) {}

// EnterExprParentesis is called when production ExprParentesis is entered.
func (s *BasegramaticaListener) EnterExprParentesis(ctx *ExprParentesisContext) {}

// ExitExprParentesis is called when production ExprParentesis is exited.
func (s *BasegramaticaListener) ExitExprParentesis(ctx *ExprParentesisContext) {}

// EnterExprLen is called when production ExprLen is entered.
func (s *BasegramaticaListener) EnterExprLen(ctx *ExprLenContext) {}

// ExitExprLen is called when production ExprLen is exited.
func (s *BasegramaticaListener) ExitExprLen(ctx *ExprLenContext) {}

// EnterExprSlice is called when production ExprSlice is entered.
func (s *BasegramaticaListener) EnterExprSlice(ctx *ExprSliceContext) {}

// ExitExprSlice is called when production ExprSlice is exited.
func (s *BasegramaticaListener) ExitExprSlice(ctx *ExprSliceContext) {}

// EnterExprJoin is called when production ExprJoin is entered.
func (s *BasegramaticaListener) EnterExprJoin(ctx *ExprJoinContext) {}

// ExitExprJoin is called when production ExprJoin is exited.
func (s *BasegramaticaListener) ExitExprJoin(ctx *ExprJoinContext) {}

// EnterExprAppend is called when production ExprAppend is entered.
func (s *BasegramaticaListener) EnterExprAppend(ctx *ExprAppendContext) {}

// ExitExprAppend is called when production ExprAppend is exited.
func (s *BasegramaticaListener) ExitExprAppend(ctx *ExprAppendContext) {}

// EnterExprLlamadaFuncionStruct is called when production ExprLlamadaFuncionStruct is entered.
func (s *BasegramaticaListener) EnterExprLlamadaFuncionStruct(ctx *ExprLlamadaFuncionStructContext) {}

// ExitExprLlamadaFuncionStruct is called when production ExprLlamadaFuncionStruct is exited.
func (s *BasegramaticaListener) ExitExprLlamadaFuncionStruct(ctx *ExprLlamadaFuncionStructContext) {}

// EnterAtributosStruct is called when production atributosStruct is entered.
func (s *BasegramaticaListener) EnterAtributosStruct(ctx *AtributosStructContext) {}

// ExitAtributosStruct is called when production atributosStruct is exited.
func (s *BasegramaticaListener) ExitAtributosStruct(ctx *AtributosStructContext) {}

// EnterInicializacion_slice is called when production inicializacion_slice is entered.
func (s *BasegramaticaListener) EnterInicializacion_slice(ctx *Inicializacion_sliceContext) {}

// ExitInicializacion_slice is called when production inicializacion_slice is exited.
func (s *BasegramaticaListener) ExitInicializacion_slice(ctx *Inicializacion_sliceContext) {}

// EnterLista_expresion is called when production lista_expresion is entered.
func (s *BasegramaticaListener) EnterLista_expresion(ctx *Lista_expresionContext) {}

// ExitLista_expresion is called when production lista_expresion is exited.
func (s *BasegramaticaListener) ExitLista_expresion(ctx *Lista_expresionContext) {}

// EnterLista_slices is called when production lista_slices is entered.
func (s *BasegramaticaListener) EnterLista_slices(ctx *Lista_slicesContext) {}

// ExitLista_slices is called when production lista_slices is exited.
func (s *BasegramaticaListener) ExitLista_slices(ctx *Lista_slicesContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BasegramaticaListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BasegramaticaListener) ExitTipo(ctx *TipoContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BasegramaticaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BasegramaticaListener) ExitBloque(ctx *BloqueContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BasegramaticaListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BasegramaticaListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterCreacion_instancia_struct is called when production creacion_instancia_struct is entered.
func (s *BasegramaticaListener) EnterCreacion_instancia_struct(ctx *Creacion_instancia_structContext) {
}

// ExitCreacion_instancia_struct is called when production creacion_instancia_struct is exited.
func (s *BasegramaticaListener) ExitCreacion_instancia_struct(ctx *Creacion_instancia_structContext) {
}

// EnterLista_parametros_func is called when production lista_parametros_func is entered.
func (s *BasegramaticaListener) EnterLista_parametros_func(ctx *Lista_parametros_funcContext) {}

// ExitLista_parametros_func is called when production lista_parametros_func is exited.
func (s *BasegramaticaListener) ExitLista_parametros_func(ctx *Lista_parametros_funcContext) {}

// EnterParam_func is called when production param_func is entered.
func (s *BasegramaticaListener) EnterParam_func(ctx *Param_funcContext) {}

// ExitParam_func is called when production param_func is exited.
func (s *BasegramaticaListener) ExitParam_func(ctx *Param_funcContext) {}

// EnterSentFor1 is called when production SentFor1 is entered.
func (s *BasegramaticaListener) EnterSentFor1(ctx *SentFor1Context) {}

// ExitSentFor1 is called when production SentFor1 is exited.
func (s *BasegramaticaListener) ExitSentFor1(ctx *SentFor1Context) {}

// EnterSentForClasico is called when production SentForClasico is entered.
func (s *BasegramaticaListener) EnterSentForClasico(ctx *SentForClasicoContext) {}

// ExitSentForClasico is called when production SentForClasico is exited.
func (s *BasegramaticaListener) ExitSentForClasico(ctx *SentForClasicoContext) {}

// EnterSentencia_init is called when production sentencia_init is entered.
func (s *BasegramaticaListener) EnterSentencia_init(ctx *Sentencia_initContext) {}

// ExitSentencia_init is called when production sentencia_init is exited.
func (s *BasegramaticaListener) ExitSentencia_init(ctx *Sentencia_initContext) {}

// EnterSentencia_update is called when production sentencia_update is entered.
func (s *BasegramaticaListener) EnterSentencia_update(ctx *Sentencia_updateContext) {}

// ExitSentencia_update is called when production sentencia_update is exited.
func (s *BasegramaticaListener) ExitSentencia_update(ctx *Sentencia_updateContext) {}

// EnterIncremento_variable is called when production incremento_variable is entered.
func (s *BasegramaticaListener) EnterIncremento_variable(ctx *Incremento_variableContext) {}

// ExitIncremento_variable is called when production incremento_variable is exited.
func (s *BasegramaticaListener) ExitIncremento_variable(ctx *Incremento_variableContext) {}

// EnterLlamadaIndexOf is called when production llamadaIndexOf is entered.
func (s *BasegramaticaListener) EnterLlamadaIndexOf(ctx *LlamadaIndexOfContext) {}

// ExitLlamadaIndexOf is called when production llamadaIndexOf is exited.
func (s *BasegramaticaListener) ExitLlamadaIndexOf(ctx *LlamadaIndexOfContext) {}

// EnterLlamadaJoin is called when production llamadaJoin is entered.
func (s *BasegramaticaListener) EnterLlamadaJoin(ctx *LlamadaJoinContext) {}

// ExitLlamadaJoin is called when production llamadaJoin is exited.
func (s *BasegramaticaListener) ExitLlamadaJoin(ctx *LlamadaJoinContext) {}

// EnterLlamadaLen is called when production llamadaLen is entered.
func (s *BasegramaticaListener) EnterLlamadaLen(ctx *LlamadaLenContext) {}

// ExitLlamadaLen is called when production llamadaLen is exited.
func (s *BasegramaticaListener) ExitLlamadaLen(ctx *LlamadaLenContext) {}

// EnterLlamadaAppend is called when production llamadaAppend is entered.
func (s *BasegramaticaListener) EnterLlamadaAppend(ctx *LlamadaAppendContext) {}

// ExitLlamadaAppend is called when production llamadaAppend is exited.
func (s *BasegramaticaListener) ExitLlamadaAppend(ctx *LlamadaAppendContext) {}

// EnterAccesoSlice is called when production accesoSlice is entered.
func (s *BasegramaticaListener) EnterAccesoSlice(ctx *AccesoSliceContext) {}

// ExitAccesoSlice is called when production accesoSlice is exited.
func (s *BasegramaticaListener) ExitAccesoSlice(ctx *AccesoSliceContext) {}

// EnterDeclaracion_matriz is called when production declaracion_matriz is entered.
func (s *BasegramaticaListener) EnterDeclaracion_matriz(ctx *Declaracion_matrizContext) {}

// ExitDeclaracion_matriz is called when production declaracion_matriz is exited.
func (s *BasegramaticaListener) ExitDeclaracion_matriz(ctx *Declaracion_matrizContext) {}

// EnterLista_arreglos is called when production lista_arreglos is entered.
func (s *BasegramaticaListener) EnterLista_arreglos(ctx *Lista_arreglosContext) {}

// ExitLista_arreglos is called when production lista_arreglos is exited.
func (s *BasegramaticaListener) ExitLista_arreglos(ctx *Lista_arreglosContext) {}

// EnterAsignacion_acceso_matriz is called when production asignacion_acceso_matriz is entered.
func (s *BasegramaticaListener) EnterAsignacion_acceso_matriz(ctx *Asignacion_acceso_matrizContext) {}

// ExitAsignacion_acceso_matriz is called when production asignacion_acceso_matriz is exited.
func (s *BasegramaticaListener) ExitAsignacion_acceso_matriz(ctx *Asignacion_acceso_matrizContext) {}

// EnterAccesoMatriz is called when production accesoMatriz is entered.
func (s *BasegramaticaListener) EnterAccesoMatriz(ctx *AccesoMatrizContext) {}

// ExitAccesoMatriz is called when production accesoMatriz is exited.
func (s *BasegramaticaListener) ExitAccesoMatriz(ctx *AccesoMatrizContext) {}

// EnterLlamadaParseFloat is called when production llamadaParseFloat is entered.
func (s *BasegramaticaListener) EnterLlamadaParseFloat(ctx *LlamadaParseFloatContext) {}

// ExitLlamadaParseFloat is called when production llamadaParseFloat is exited.
func (s *BasegramaticaListener) ExitLlamadaParseFloat(ctx *LlamadaParseFloatContext) {}

// EnterLlamadaTypeOf is called when production llamadaTypeOf is entered.
func (s *BasegramaticaListener) EnterLlamadaTypeOf(ctx *LlamadaTypeOfContext) {}

// ExitLlamadaTypeOf is called when production llamadaTypeOf is exited.
func (s *BasegramaticaListener) ExitLlamadaTypeOf(ctx *LlamadaTypeOfContext) {}

// EnterListaCampos is called when production listaCampos is entered.
func (s *BasegramaticaListener) EnterListaCampos(ctx *ListaCamposContext) {}

// ExitListaCampos is called when production listaCampos is exited.
func (s *BasegramaticaListener) ExitListaCampos(ctx *ListaCamposContext) {}

// EnterCampoStruct is called when production campoStruct is entered.
func (s *BasegramaticaListener) EnterCampoStruct(ctx *CampoStructContext) {}

// ExitCampoStruct is called when production campoStruct is exited.
func (s *BasegramaticaListener) ExitCampoStruct(ctx *CampoStructContext) {}

// EnterAccesoInstaciaAtributo is called when production accesoInstaciaAtributo is entered.
func (s *BasegramaticaListener) EnterAccesoInstaciaAtributo(ctx *AccesoInstaciaAtributoContext) {}

// ExitAccesoInstaciaAtributo is called when production accesoInstaciaAtributo is exited.
func (s *BasegramaticaListener) ExitAccesoInstaciaAtributo(ctx *AccesoInstaciaAtributoContext) {}

// EnterAsignacion_acceso_atributo is called when production asignacion_acceso_atributo is entered.
func (s *BasegramaticaListener) EnterAsignacion_acceso_atributo(ctx *Asignacion_acceso_atributoContext) {
}

// ExitAsignacion_acceso_atributo is called when production asignacion_acceso_atributo is exited.
func (s *BasegramaticaListener) ExitAsignacion_acceso_atributo(ctx *Asignacion_acceso_atributoContext) {
}

// EnterLlamada_funcion_struct is called when production llamada_funcion_struct is entered.
func (s *BasegramaticaListener) EnterLlamada_funcion_struct(ctx *Llamada_funcion_structContext) {}

// ExitLlamada_funcion_struct is called when production llamada_funcion_struct is exited.
func (s *BasegramaticaListener) ExitLlamada_funcion_struct(ctx *Llamada_funcion_structContext) {}
