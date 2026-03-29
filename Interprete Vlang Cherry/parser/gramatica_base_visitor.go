// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr4-go/antlr/v4"

type BasegramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasegramaticaVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracionStruct(ctx *DeclaracionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFuncionStruct_declaracion(ctx *FuncionStruct_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDecFunc1(ctx *DecFunc1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDecFunc2(ctx *DecFunc2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_parametros(ctx *Lista_parametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_sentencias(ctx *Lista_sentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentencia(ctx *SentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentenciaBreak(ctx *SentenciaBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentenciaReturn(ctx *SentenciaReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadafuncion(ctx *LlamadafuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentBreak(ctx *SentBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentContinue(ctx *SentContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentReturn(ctx *SentReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentIf1(ctx *SentIf1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentIf2(ctx *SentIf2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentIf3(ctx *SentIf3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentenciaSwitch(ctx *SentenciaSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBloqueCase(ctx *BloqueCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBloqueDefault(ctx *BloqueDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentenciaPrintln(ctx *SentenciaPrintlnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsigSimple(ctx *AsigSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsigSumRest(ctx *AsigSumRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacion_acceso_slice1(ctx *Asignacion_acceso_slice1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDecSlice1(ctx *DecSlice1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDecSlice2(ctx *DecSlice2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsigacionSlice(ctx *AsigacionSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDec1(ctx *Dec1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDec2(ctx *Dec2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDec3(ctx *Dec3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprOpAritmetica(ctx *ExprOpAritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprBoolFalse(ctx *ExprBoolFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprIdentificador(ctx *ExprIdentificadorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprRune(ctx *ExprRuneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprFloat(ctx *ExprFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprNot(ctx *ExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprNegUnaria(ctx *ExprNegUnariaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprNil(ctx *ExprNilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprMultDivMod(ctx *ExprMultDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprAccesoMatriz(ctx *ExprAccesoMatrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprAccesoSlice(ctx *ExprAccesoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprSliceLiteral(ctx *ExprSliceLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprBoolTrue(ctx *ExprBoolTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprLlamadaFuncion(ctx *ExprLlamadaFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprLogica(ctx *ExprLogicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprIndexOf(ctx *ExprIndexOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprAccesoInstanciaAtributo(ctx *ExprAccesoInstanciaAtributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprRelacional(ctx *ExprRelacionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprString(ctx *ExprStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprInt(ctx *ExprIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprParentesis(ctx *ExprParentesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprLen(ctx *ExprLenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprSlice(ctx *ExprSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprJoin(ctx *ExprJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprAppend(ctx *ExprAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExprLlamadaFuncionStruct(ctx *ExprLlamadaFuncionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAtributosStruct(ctx *AtributosStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitInicializacion_slice(ctx *Inicializacion_sliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_expresion(ctx *Lista_expresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_slices(ctx *Lista_slicesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCreacion_instancia_struct(ctx *Creacion_instancia_structContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_parametros_func(ctx *Lista_parametros_funcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitParam_func(ctx *Param_funcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentFor1(ctx *SentFor1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentForClasico(ctx *SentForClasicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentencia_init(ctx *Sentencia_initContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSentencia_update(ctx *Sentencia_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitIncremento_variable(ctx *Incremento_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaIndexOf(ctx *LlamadaIndexOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaJoin(ctx *LlamadaJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaLen(ctx *LlamadaLenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaAppend(ctx *LlamadaAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAccesoSlice(ctx *AccesoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracion_matriz(ctx *Declaracion_matrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLista_arreglos(ctx *Lista_arreglosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacion_acceso_matriz(ctx *Asignacion_acceso_matrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAccesoMatriz(ctx *AccesoMatrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaParseFloat(ctx *LlamadaParseFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaTypeOf(ctx *LlamadaTypeOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitListaCampos(ctx *ListaCamposContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCampoStruct(ctx *CampoStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAccesoInstaciaAtributo(ctx *AccesoInstaciaAtributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacion_acceso_atributo(ctx *Asignacion_acceso_atributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamada_funcion_struct(ctx *Llamada_funcion_structContext) interface{} {
	return v.VisitChildren(ctx)
}
