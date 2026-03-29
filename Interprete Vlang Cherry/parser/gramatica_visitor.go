// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by gramaticaParser.
type gramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by gramaticaParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracionStruct.
	VisitDeclaracionStruct(ctx *DeclaracionStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#funcionStruct_declaracion.
	VisitFuncionStruct_declaracion(ctx *FuncionStruct_declaracionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#DecFunc1.
	VisitDecFunc1(ctx *DecFunc1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#DecFunc2.
	VisitDecFunc2(ctx *DecFunc2Context) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_parametros.
	VisitLista_parametros(ctx *Lista_parametrosContext) interface{}

	// Visit a parse tree produced by gramaticaParser#parametro.
	VisitParametro(ctx *ParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_sentencias.
	VisitLista_sentencias(ctx *Lista_sentenciasContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentencia.
	VisitSentencia(ctx *SentenciaContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentenciaBreak.
	VisitSentenciaBreak(ctx *SentenciaBreakContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentenciaReturn.
	VisitSentenciaReturn(ctx *SentenciaReturnContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadafuncion.
	VisitLlamadafuncion(ctx *LlamadafuncionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SentBreak.
	VisitSentBreak(ctx *SentBreakContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SentContinue.
	VisitSentContinue(ctx *SentContinueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SentReturn.
	VisitSentReturn(ctx *SentReturnContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SentIf1.
	VisitSentIf1(ctx *SentIf1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#SentIf2.
	VisitSentIf2(ctx *SentIf2Context) interface{}

	// Visit a parse tree produced by gramaticaParser#SentIf3.
	VisitSentIf3(ctx *SentIf3Context) interface{}

	// Visit a parse tree produced by gramaticaParser#sentenciaSwitch.
	VisitSentenciaSwitch(ctx *SentenciaSwitchContext) interface{}

	// Visit a parse tree produced by gramaticaParser#bloqueCase.
	VisitBloqueCase(ctx *BloqueCaseContext) interface{}

	// Visit a parse tree produced by gramaticaParser#bloqueDefault.
	VisitBloqueDefault(ctx *BloqueDefaultContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentenciaPrintln.
	VisitSentenciaPrintln(ctx *SentenciaPrintlnContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsigSimple.
	VisitAsigSimple(ctx *AsigSimpleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsigSumRest.
	VisitAsigSumRest(ctx *AsigSumRestContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asignacion_acceso_slice1.
	VisitAsignacion_acceso_slice1(ctx *Asignacion_acceso_slice1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#DecSlice1.
	VisitDecSlice1(ctx *DecSlice1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#DecSlice2.
	VisitDecSlice2(ctx *DecSlice2Context) interface{}

	// Visit a parse tree produced by gramaticaParser#AsigacionSlice.
	VisitAsigacionSlice(ctx *AsigacionSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#Dec1.
	VisitDec1(ctx *Dec1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#Dec2.
	VisitDec2(ctx *Dec2Context) interface{}

	// Visit a parse tree produced by gramaticaParser#Dec3.
	VisitDec3(ctx *Dec3Context) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprOpAritmetica.
	VisitExprOpAritmetica(ctx *ExprOpAritmeticaContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprBoolFalse.
	VisitExprBoolFalse(ctx *ExprBoolFalseContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprIdentificador.
	VisitExprIdentificador(ctx *ExprIdentificadorContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprRune.
	VisitExprRune(ctx *ExprRuneContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprFloat.
	VisitExprFloat(ctx *ExprFloatContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprNot.
	VisitExprNot(ctx *ExprNotContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprNegUnaria.
	VisitExprNegUnaria(ctx *ExprNegUnariaContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprNil.
	VisitExprNil(ctx *ExprNilContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprMultDivMod.
	VisitExprMultDivMod(ctx *ExprMultDivModContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprAccesoMatriz.
	VisitExprAccesoMatriz(ctx *ExprAccesoMatrizContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprAccesoSlice.
	VisitExprAccesoSlice(ctx *ExprAccesoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprSliceLiteral.
	VisitExprSliceLiteral(ctx *ExprSliceLiteralContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprBoolTrue.
	VisitExprBoolTrue(ctx *ExprBoolTrueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprLlamadaFuncion.
	VisitExprLlamadaFuncion(ctx *ExprLlamadaFuncionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprLogica.
	VisitExprLogica(ctx *ExprLogicaContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprIndexOf.
	VisitExprIndexOf(ctx *ExprIndexOfContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprAccesoInstanciaAtributo.
	VisitExprAccesoInstanciaAtributo(ctx *ExprAccesoInstanciaAtributoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprRelacional.
	VisitExprRelacional(ctx *ExprRelacionalContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprString.
	VisitExprString(ctx *ExprStringContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprInt.
	VisitExprInt(ctx *ExprIntContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprParentesis.
	VisitExprParentesis(ctx *ExprParentesisContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprLen.
	VisitExprLen(ctx *ExprLenContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprSlice.
	VisitExprSlice(ctx *ExprSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprJoin.
	VisitExprJoin(ctx *ExprJoinContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprAppend.
	VisitExprAppend(ctx *ExprAppendContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ExprLlamadaFuncionStruct.
	VisitExprLlamadaFuncionStruct(ctx *ExprLlamadaFuncionStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#atributosStruct.
	VisitAtributosStruct(ctx *AtributosStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#inicializacion_slice.
	VisitInicializacion_slice(ctx *Inicializacion_sliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_expresion.
	VisitLista_expresion(ctx *Lista_expresionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_slices.
	VisitLista_slices(ctx *Lista_slicesContext) interface{}

	// Visit a parse tree produced by gramaticaParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by gramaticaParser#creacion_instancia_struct.
	VisitCreacion_instancia_struct(ctx *Creacion_instancia_structContext) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_parametros_func.
	VisitLista_parametros_func(ctx *Lista_parametros_funcContext) interface{}

	// Visit a parse tree produced by gramaticaParser#param_func.
	VisitParam_func(ctx *Param_funcContext) interface{}

	// Visit a parse tree produced by gramaticaParser#SentFor1.
	VisitSentFor1(ctx *SentFor1Context) interface{}

	// Visit a parse tree produced by gramaticaParser#SentForClasico.
	VisitSentForClasico(ctx *SentForClasicoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentencia_init.
	VisitSentencia_init(ctx *Sentencia_initContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sentencia_update.
	VisitSentencia_update(ctx *Sentencia_updateContext) interface{}

	// Visit a parse tree produced by gramaticaParser#incremento_variable.
	VisitIncremento_variable(ctx *Incremento_variableContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaIndexOf.
	VisitLlamadaIndexOf(ctx *LlamadaIndexOfContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaJoin.
	VisitLlamadaJoin(ctx *LlamadaJoinContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaLen.
	VisitLlamadaLen(ctx *LlamadaLenContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaAppend.
	VisitLlamadaAppend(ctx *LlamadaAppendContext) interface{}

	// Visit a parse tree produced by gramaticaParser#accesoSlice.
	VisitAccesoSlice(ctx *AccesoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracion_matriz.
	VisitDeclaracion_matriz(ctx *Declaracion_matrizContext) interface{}

	// Visit a parse tree produced by gramaticaParser#lista_arreglos.
	VisitLista_arreglos(ctx *Lista_arreglosContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asignacion_acceso_matriz.
	VisitAsignacion_acceso_matriz(ctx *Asignacion_acceso_matrizContext) interface{}

	// Visit a parse tree produced by gramaticaParser#accesoMatriz.
	VisitAccesoMatriz(ctx *AccesoMatrizContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaParseFloat.
	VisitLlamadaParseFloat(ctx *LlamadaParseFloatContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaTypeOf.
	VisitLlamadaTypeOf(ctx *LlamadaTypeOfContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaCampos.
	VisitListaCampos(ctx *ListaCamposContext) interface{}

	// Visit a parse tree produced by gramaticaParser#campoStruct.
	VisitCampoStruct(ctx *CampoStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#accesoInstaciaAtributo.
	VisitAccesoInstaciaAtributo(ctx *AccesoInstaciaAtributoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asignacion_acceso_atributo.
	VisitAsignacion_acceso_atributo(ctx *Asignacion_acceso_atributoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamada_funcion_struct.
	VisitLlamada_funcion_struct(ctx *Llamada_funcion_structContext) interface{}
}
