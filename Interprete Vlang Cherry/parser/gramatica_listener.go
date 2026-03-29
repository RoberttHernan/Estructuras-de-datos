// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gramatica

import "github.com/antlr4-go/antlr/v4"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDeclaracionStruct is called when entering the declaracionStruct production.
	EnterDeclaracionStruct(c *DeclaracionStructContext)

	// EnterFuncionStruct_declaracion is called when entering the funcionStruct_declaracion production.
	EnterFuncionStruct_declaracion(c *FuncionStruct_declaracionContext)

	// EnterDecFunc1 is called when entering the DecFunc1 production.
	EnterDecFunc1(c *DecFunc1Context)

	// EnterDecFunc2 is called when entering the DecFunc2 production.
	EnterDecFunc2(c *DecFunc2Context)

	// EnterLista_parametros is called when entering the lista_parametros production.
	EnterLista_parametros(c *Lista_parametrosContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterLista_sentencias is called when entering the lista_sentencias production.
	EnterLista_sentencias(c *Lista_sentenciasContext)

	// EnterSentencia is called when entering the sentencia production.
	EnterSentencia(c *SentenciaContext)

	// EnterSentenciaBreak is called when entering the sentenciaBreak production.
	EnterSentenciaBreak(c *SentenciaBreakContext)

	// EnterSentenciaReturn is called when entering the sentenciaReturn production.
	EnterSentenciaReturn(c *SentenciaReturnContext)

	// EnterLlamadafuncion is called when entering the llamadafuncion production.
	EnterLlamadafuncion(c *LlamadafuncionContext)

	// EnterSentBreak is called when entering the SentBreak production.
	EnterSentBreak(c *SentBreakContext)

	// EnterSentContinue is called when entering the SentContinue production.
	EnterSentContinue(c *SentContinueContext)

	// EnterSentReturn is called when entering the SentReturn production.
	EnterSentReturn(c *SentReturnContext)

	// EnterSentIf1 is called when entering the SentIf1 production.
	EnterSentIf1(c *SentIf1Context)

	// EnterSentIf2 is called when entering the SentIf2 production.
	EnterSentIf2(c *SentIf2Context)

	// EnterSentIf3 is called when entering the SentIf3 production.
	EnterSentIf3(c *SentIf3Context)

	// EnterSentenciaSwitch is called when entering the sentenciaSwitch production.
	EnterSentenciaSwitch(c *SentenciaSwitchContext)

	// EnterBloqueCase is called when entering the bloqueCase production.
	EnterBloqueCase(c *BloqueCaseContext)

	// EnterBloqueDefault is called when entering the bloqueDefault production.
	EnterBloqueDefault(c *BloqueDefaultContext)

	// EnterSentenciaPrintln is called when entering the sentenciaPrintln production.
	EnterSentenciaPrintln(c *SentenciaPrintlnContext)

	// EnterAsigSimple is called when entering the AsigSimple production.
	EnterAsigSimple(c *AsigSimpleContext)

	// EnterAsigSumRest is called when entering the AsigSumRest production.
	EnterAsigSumRest(c *AsigSumRestContext)

	// EnterAsignacion_acceso_slice1 is called when entering the asignacion_acceso_slice1 production.
	EnterAsignacion_acceso_slice1(c *Asignacion_acceso_slice1Context)

	// EnterDecSlice1 is called when entering the DecSlice1 production.
	EnterDecSlice1(c *DecSlice1Context)

	// EnterDecSlice2 is called when entering the DecSlice2 production.
	EnterDecSlice2(c *DecSlice2Context)

	// EnterAsigacionSlice is called when entering the AsigacionSlice production.
	EnterAsigacionSlice(c *AsigacionSliceContext)

	// EnterDec1 is called when entering the Dec1 production.
	EnterDec1(c *Dec1Context)

	// EnterDec2 is called when entering the Dec2 production.
	EnterDec2(c *Dec2Context)

	// EnterDec3 is called when entering the Dec3 production.
	EnterDec3(c *Dec3Context)

	// EnterExprOpAritmetica is called when entering the ExprOpAritmetica production.
	EnterExprOpAritmetica(c *ExprOpAritmeticaContext)

	// EnterExprBoolFalse is called when entering the ExprBoolFalse production.
	EnterExprBoolFalse(c *ExprBoolFalseContext)

	// EnterExprIdentificador is called when entering the ExprIdentificador production.
	EnterExprIdentificador(c *ExprIdentificadorContext)

	// EnterExprRune is called when entering the ExprRune production.
	EnterExprRune(c *ExprRuneContext)

	// EnterExprFloat is called when entering the ExprFloat production.
	EnterExprFloat(c *ExprFloatContext)

	// EnterExprNot is called when entering the ExprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprNegUnaria is called when entering the ExprNegUnaria production.
	EnterExprNegUnaria(c *ExprNegUnariaContext)

	// EnterExprNil is called when entering the ExprNil production.
	EnterExprNil(c *ExprNilContext)

	// EnterExprMultDivMod is called when entering the ExprMultDivMod production.
	EnterExprMultDivMod(c *ExprMultDivModContext)

	// EnterExprAccesoMatriz is called when entering the ExprAccesoMatriz production.
	EnterExprAccesoMatriz(c *ExprAccesoMatrizContext)

	// EnterExprAccesoSlice is called when entering the ExprAccesoSlice production.
	EnterExprAccesoSlice(c *ExprAccesoSliceContext)

	// EnterExprSliceLiteral is called when entering the ExprSliceLiteral production.
	EnterExprSliceLiteral(c *ExprSliceLiteralContext)

	// EnterExprBoolTrue is called when entering the ExprBoolTrue production.
	EnterExprBoolTrue(c *ExprBoolTrueContext)

	// EnterExprLlamadaFuncion is called when entering the ExprLlamadaFuncion production.
	EnterExprLlamadaFuncion(c *ExprLlamadaFuncionContext)

	// EnterExprLogica is called when entering the ExprLogica production.
	EnterExprLogica(c *ExprLogicaContext)

	// EnterExprIndexOf is called when entering the ExprIndexOf production.
	EnterExprIndexOf(c *ExprIndexOfContext)

	// EnterExprAccesoInstanciaAtributo is called when entering the ExprAccesoInstanciaAtributo production.
	EnterExprAccesoInstanciaAtributo(c *ExprAccesoInstanciaAtributoContext)

	// EnterExprRelacional is called when entering the ExprRelacional production.
	EnterExprRelacional(c *ExprRelacionalContext)

	// EnterExprString is called when entering the ExprString production.
	EnterExprString(c *ExprStringContext)

	// EnterExprInt is called when entering the ExprInt production.
	EnterExprInt(c *ExprIntContext)

	// EnterExprParentesis is called when entering the ExprParentesis production.
	EnterExprParentesis(c *ExprParentesisContext)

	// EnterExprLen is called when entering the ExprLen production.
	EnterExprLen(c *ExprLenContext)

	// EnterExprSlice is called when entering the ExprSlice production.
	EnterExprSlice(c *ExprSliceContext)

	// EnterExprJoin is called when entering the ExprJoin production.
	EnterExprJoin(c *ExprJoinContext)

	// EnterExprAppend is called when entering the ExprAppend production.
	EnterExprAppend(c *ExprAppendContext)

	// EnterExprLlamadaFuncionStruct is called when entering the ExprLlamadaFuncionStruct production.
	EnterExprLlamadaFuncionStruct(c *ExprLlamadaFuncionStructContext)

	// EnterAtributosStruct is called when entering the atributosStruct production.
	EnterAtributosStruct(c *AtributosStructContext)

	// EnterInicializacion_slice is called when entering the inicializacion_slice production.
	EnterInicializacion_slice(c *Inicializacion_sliceContext)

	// EnterLista_expresion is called when entering the lista_expresion production.
	EnterLista_expresion(c *Lista_expresionContext)

	// EnterLista_slices is called when entering the lista_slices production.
	EnterLista_slices(c *Lista_slicesContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterCreacion_instancia_struct is called when entering the creacion_instancia_struct production.
	EnterCreacion_instancia_struct(c *Creacion_instancia_structContext)

	// EnterLista_parametros_func is called when entering the lista_parametros_func production.
	EnterLista_parametros_func(c *Lista_parametros_funcContext)

	// EnterParam_func is called when entering the param_func production.
	EnterParam_func(c *Param_funcContext)

	// EnterSentFor1 is called when entering the SentFor1 production.
	EnterSentFor1(c *SentFor1Context)

	// EnterSentForClasico is called when entering the SentForClasico production.
	EnterSentForClasico(c *SentForClasicoContext)

	// EnterSentencia_init is called when entering the sentencia_init production.
	EnterSentencia_init(c *Sentencia_initContext)

	// EnterSentencia_update is called when entering the sentencia_update production.
	EnterSentencia_update(c *Sentencia_updateContext)

	// EnterIncremento_variable is called when entering the incremento_variable production.
	EnterIncremento_variable(c *Incremento_variableContext)

	// EnterLlamadaIndexOf is called when entering the llamadaIndexOf production.
	EnterLlamadaIndexOf(c *LlamadaIndexOfContext)

	// EnterLlamadaJoin is called when entering the llamadaJoin production.
	EnterLlamadaJoin(c *LlamadaJoinContext)

	// EnterLlamadaLen is called when entering the llamadaLen production.
	EnterLlamadaLen(c *LlamadaLenContext)

	// EnterLlamadaAppend is called when entering the llamadaAppend production.
	EnterLlamadaAppend(c *LlamadaAppendContext)

	// EnterAccesoSlice is called when entering the accesoSlice production.
	EnterAccesoSlice(c *AccesoSliceContext)

	// EnterDeclaracion_matriz is called when entering the declaracion_matriz production.
	EnterDeclaracion_matriz(c *Declaracion_matrizContext)

	// EnterLista_arreglos is called when entering the lista_arreglos production.
	EnterLista_arreglos(c *Lista_arreglosContext)

	// EnterAsignacion_acceso_matriz is called when entering the asignacion_acceso_matriz production.
	EnterAsignacion_acceso_matriz(c *Asignacion_acceso_matrizContext)

	// EnterAccesoMatriz is called when entering the accesoMatriz production.
	EnterAccesoMatriz(c *AccesoMatrizContext)

	// EnterLlamadaParseFloat is called when entering the llamadaParseFloat production.
	EnterLlamadaParseFloat(c *LlamadaParseFloatContext)

	// EnterLlamadaTypeOf is called when entering the llamadaTypeOf production.
	EnterLlamadaTypeOf(c *LlamadaTypeOfContext)

	// EnterListaCampos is called when entering the listaCampos production.
	EnterListaCampos(c *ListaCamposContext)

	// EnterCampoStruct is called when entering the campoStruct production.
	EnterCampoStruct(c *CampoStructContext)

	// EnterAccesoInstaciaAtributo is called when entering the accesoInstaciaAtributo production.
	EnterAccesoInstaciaAtributo(c *AccesoInstaciaAtributoContext)

	// EnterAsignacion_acceso_atributo is called when entering the asignacion_acceso_atributo production.
	EnterAsignacion_acceso_atributo(c *Asignacion_acceso_atributoContext)

	// EnterLlamada_funcion_struct is called when entering the llamada_funcion_struct production.
	EnterLlamada_funcion_struct(c *Llamada_funcion_structContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDeclaracionStruct is called when exiting the declaracionStruct production.
	ExitDeclaracionStruct(c *DeclaracionStructContext)

	// ExitFuncionStruct_declaracion is called when exiting the funcionStruct_declaracion production.
	ExitFuncionStruct_declaracion(c *FuncionStruct_declaracionContext)

	// ExitDecFunc1 is called when exiting the DecFunc1 production.
	ExitDecFunc1(c *DecFunc1Context)

	// ExitDecFunc2 is called when exiting the DecFunc2 production.
	ExitDecFunc2(c *DecFunc2Context)

	// ExitLista_parametros is called when exiting the lista_parametros production.
	ExitLista_parametros(c *Lista_parametrosContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitLista_sentencias is called when exiting the lista_sentencias production.
	ExitLista_sentencias(c *Lista_sentenciasContext)

	// ExitSentencia is called when exiting the sentencia production.
	ExitSentencia(c *SentenciaContext)

	// ExitSentenciaBreak is called when exiting the sentenciaBreak production.
	ExitSentenciaBreak(c *SentenciaBreakContext)

	// ExitSentenciaReturn is called when exiting the sentenciaReturn production.
	ExitSentenciaReturn(c *SentenciaReturnContext)

	// ExitLlamadafuncion is called when exiting the llamadafuncion production.
	ExitLlamadafuncion(c *LlamadafuncionContext)

	// ExitSentBreak is called when exiting the SentBreak production.
	ExitSentBreak(c *SentBreakContext)

	// ExitSentContinue is called when exiting the SentContinue production.
	ExitSentContinue(c *SentContinueContext)

	// ExitSentReturn is called when exiting the SentReturn production.
	ExitSentReturn(c *SentReturnContext)

	// ExitSentIf1 is called when exiting the SentIf1 production.
	ExitSentIf1(c *SentIf1Context)

	// ExitSentIf2 is called when exiting the SentIf2 production.
	ExitSentIf2(c *SentIf2Context)

	// ExitSentIf3 is called when exiting the SentIf3 production.
	ExitSentIf3(c *SentIf3Context)

	// ExitSentenciaSwitch is called when exiting the sentenciaSwitch production.
	ExitSentenciaSwitch(c *SentenciaSwitchContext)

	// ExitBloqueCase is called when exiting the bloqueCase production.
	ExitBloqueCase(c *BloqueCaseContext)

	// ExitBloqueDefault is called when exiting the bloqueDefault production.
	ExitBloqueDefault(c *BloqueDefaultContext)

	// ExitSentenciaPrintln is called when exiting the sentenciaPrintln production.
	ExitSentenciaPrintln(c *SentenciaPrintlnContext)

	// ExitAsigSimple is called when exiting the AsigSimple production.
	ExitAsigSimple(c *AsigSimpleContext)

	// ExitAsigSumRest is called when exiting the AsigSumRest production.
	ExitAsigSumRest(c *AsigSumRestContext)

	// ExitAsignacion_acceso_slice1 is called when exiting the asignacion_acceso_slice1 production.
	ExitAsignacion_acceso_slice1(c *Asignacion_acceso_slice1Context)

	// ExitDecSlice1 is called when exiting the DecSlice1 production.
	ExitDecSlice1(c *DecSlice1Context)

	// ExitDecSlice2 is called when exiting the DecSlice2 production.
	ExitDecSlice2(c *DecSlice2Context)

	// ExitAsigacionSlice is called when exiting the AsigacionSlice production.
	ExitAsigacionSlice(c *AsigacionSliceContext)

	// ExitDec1 is called when exiting the Dec1 production.
	ExitDec1(c *Dec1Context)

	// ExitDec2 is called when exiting the Dec2 production.
	ExitDec2(c *Dec2Context)

	// ExitDec3 is called when exiting the Dec3 production.
	ExitDec3(c *Dec3Context)

	// ExitExprOpAritmetica is called when exiting the ExprOpAritmetica production.
	ExitExprOpAritmetica(c *ExprOpAritmeticaContext)

	// ExitExprBoolFalse is called when exiting the ExprBoolFalse production.
	ExitExprBoolFalse(c *ExprBoolFalseContext)

	// ExitExprIdentificador is called when exiting the ExprIdentificador production.
	ExitExprIdentificador(c *ExprIdentificadorContext)

	// ExitExprRune is called when exiting the ExprRune production.
	ExitExprRune(c *ExprRuneContext)

	// ExitExprFloat is called when exiting the ExprFloat production.
	ExitExprFloat(c *ExprFloatContext)

	// ExitExprNot is called when exiting the ExprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprNegUnaria is called when exiting the ExprNegUnaria production.
	ExitExprNegUnaria(c *ExprNegUnariaContext)

	// ExitExprNil is called when exiting the ExprNil production.
	ExitExprNil(c *ExprNilContext)

	// ExitExprMultDivMod is called when exiting the ExprMultDivMod production.
	ExitExprMultDivMod(c *ExprMultDivModContext)

	// ExitExprAccesoMatriz is called when exiting the ExprAccesoMatriz production.
	ExitExprAccesoMatriz(c *ExprAccesoMatrizContext)

	// ExitExprAccesoSlice is called when exiting the ExprAccesoSlice production.
	ExitExprAccesoSlice(c *ExprAccesoSliceContext)

	// ExitExprSliceLiteral is called when exiting the ExprSliceLiteral production.
	ExitExprSliceLiteral(c *ExprSliceLiteralContext)

	// ExitExprBoolTrue is called when exiting the ExprBoolTrue production.
	ExitExprBoolTrue(c *ExprBoolTrueContext)

	// ExitExprLlamadaFuncion is called when exiting the ExprLlamadaFuncion production.
	ExitExprLlamadaFuncion(c *ExprLlamadaFuncionContext)

	// ExitExprLogica is called when exiting the ExprLogica production.
	ExitExprLogica(c *ExprLogicaContext)

	// ExitExprIndexOf is called when exiting the ExprIndexOf production.
	ExitExprIndexOf(c *ExprIndexOfContext)

	// ExitExprAccesoInstanciaAtributo is called when exiting the ExprAccesoInstanciaAtributo production.
	ExitExprAccesoInstanciaAtributo(c *ExprAccesoInstanciaAtributoContext)

	// ExitExprRelacional is called when exiting the ExprRelacional production.
	ExitExprRelacional(c *ExprRelacionalContext)

	// ExitExprString is called when exiting the ExprString production.
	ExitExprString(c *ExprStringContext)

	// ExitExprInt is called when exiting the ExprInt production.
	ExitExprInt(c *ExprIntContext)

	// ExitExprParentesis is called when exiting the ExprParentesis production.
	ExitExprParentesis(c *ExprParentesisContext)

	// ExitExprLen is called when exiting the ExprLen production.
	ExitExprLen(c *ExprLenContext)

	// ExitExprSlice is called when exiting the ExprSlice production.
	ExitExprSlice(c *ExprSliceContext)

	// ExitExprJoin is called when exiting the ExprJoin production.
	ExitExprJoin(c *ExprJoinContext)

	// ExitExprAppend is called when exiting the ExprAppend production.
	ExitExprAppend(c *ExprAppendContext)

	// ExitExprLlamadaFuncionStruct is called when exiting the ExprLlamadaFuncionStruct production.
	ExitExprLlamadaFuncionStruct(c *ExprLlamadaFuncionStructContext)

	// ExitAtributosStruct is called when exiting the atributosStruct production.
	ExitAtributosStruct(c *AtributosStructContext)

	// ExitInicializacion_slice is called when exiting the inicializacion_slice production.
	ExitInicializacion_slice(c *Inicializacion_sliceContext)

	// ExitLista_expresion is called when exiting the lista_expresion production.
	ExitLista_expresion(c *Lista_expresionContext)

	// ExitLista_slices is called when exiting the lista_slices production.
	ExitLista_slices(c *Lista_slicesContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitCreacion_instancia_struct is called when exiting the creacion_instancia_struct production.
	ExitCreacion_instancia_struct(c *Creacion_instancia_structContext)

	// ExitLista_parametros_func is called when exiting the lista_parametros_func production.
	ExitLista_parametros_func(c *Lista_parametros_funcContext)

	// ExitParam_func is called when exiting the param_func production.
	ExitParam_func(c *Param_funcContext)

	// ExitSentFor1 is called when exiting the SentFor1 production.
	ExitSentFor1(c *SentFor1Context)

	// ExitSentForClasico is called when exiting the SentForClasico production.
	ExitSentForClasico(c *SentForClasicoContext)

	// ExitSentencia_init is called when exiting the sentencia_init production.
	ExitSentencia_init(c *Sentencia_initContext)

	// ExitSentencia_update is called when exiting the sentencia_update production.
	ExitSentencia_update(c *Sentencia_updateContext)

	// ExitIncremento_variable is called when exiting the incremento_variable production.
	ExitIncremento_variable(c *Incremento_variableContext)

	// ExitLlamadaIndexOf is called when exiting the llamadaIndexOf production.
	ExitLlamadaIndexOf(c *LlamadaIndexOfContext)

	// ExitLlamadaJoin is called when exiting the llamadaJoin production.
	ExitLlamadaJoin(c *LlamadaJoinContext)

	// ExitLlamadaLen is called when exiting the llamadaLen production.
	ExitLlamadaLen(c *LlamadaLenContext)

	// ExitLlamadaAppend is called when exiting the llamadaAppend production.
	ExitLlamadaAppend(c *LlamadaAppendContext)

	// ExitAccesoSlice is called when exiting the accesoSlice production.
	ExitAccesoSlice(c *AccesoSliceContext)

	// ExitDeclaracion_matriz is called when exiting the declaracion_matriz production.
	ExitDeclaracion_matriz(c *Declaracion_matrizContext)

	// ExitLista_arreglos is called when exiting the lista_arreglos production.
	ExitLista_arreglos(c *Lista_arreglosContext)

	// ExitAsignacion_acceso_matriz is called when exiting the asignacion_acceso_matriz production.
	ExitAsignacion_acceso_matriz(c *Asignacion_acceso_matrizContext)

	// ExitAccesoMatriz is called when exiting the accesoMatriz production.
	ExitAccesoMatriz(c *AccesoMatrizContext)

	// ExitLlamadaParseFloat is called when exiting the llamadaParseFloat production.
	ExitLlamadaParseFloat(c *LlamadaParseFloatContext)

	// ExitLlamadaTypeOf is called when exiting the llamadaTypeOf production.
	ExitLlamadaTypeOf(c *LlamadaTypeOfContext)

	// ExitListaCampos is called when exiting the listaCampos production.
	ExitListaCampos(c *ListaCamposContext)

	// ExitCampoStruct is called when exiting the campoStruct production.
	ExitCampoStruct(c *CampoStructContext)

	// ExitAccesoInstaciaAtributo is called when exiting the accesoInstaciaAtributo production.
	ExitAccesoInstaciaAtributo(c *AccesoInstaciaAtributoContext)

	// ExitAsignacion_acceso_atributo is called when exiting the asignacion_acceso_atributo production.
	ExitAsignacion_acceso_atributo(c *Asignacion_acceso_atributoContext)

	// ExitLlamada_funcion_struct is called when exiting the llamada_funcion_struct production.
	ExitLlamada_funcion_struct(c *Llamada_funcion_structContext)
}
