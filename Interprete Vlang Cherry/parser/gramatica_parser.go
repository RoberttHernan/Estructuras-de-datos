// Code generated from gramatica.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // gramatica

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type gramaticaParser struct {
	*antlr.BaseParser
}

var GramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.LiteralNames = []string{
		"", "'true'", "'false'", "'&'", "'func'", "'import'", "'var'", "'mut'",
		"'int'", "'float64'", "'string'", "'bool'", "'rune'", "'slices'", "'indexOf'",
		"'strings'", "'join'", "'len'", "'append'", "'struct'", "'nil'", "'if'",
		"'else'", "'for'", "'switch'", "'case'", "'default'", "'return'", "'break'",
		"'continue'", "'in'", "'fmt'", "", "'Atoi'", "'ParseFloat'", "'typeOf'",
		"", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"';'", "':'", "'*'", "'/'", "'+='", "'-='", "'++'", "'--'", "'+'", "'-'",
		"'&&'", "'||'", "'!='", "'!'", "'=='", "'='", "'<='", "'<'", "'>='",
		"'>'", "','", "'%'", "':='", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "RESFUNC", "RESIMPORT", "RESVAR", "RESMUT", "RESINT",
		"RESFLOAT", "RESSTRING", "RESBOOL", "RESRUNE", "RESSLICE", "RESINDEX",
		"RESSTRINGS", "RESJOIN", "RESLEN", "RESAPPEND", "RESSTRUCT", "RESNIL",
		"RESIF", "RESELSE", "RESFOR", "RESSWITCH", "RESCASE", "RESDEFAULT",
		"RESRETURN", "RESBREAK", "RESCONTINUE", "RESIN", "RESFMT", "RESPRINTLN",
		"RESATOI", "RESPARSEFLOAT", "RESTYPEOF", "IDENTIFICADOR", "INT", "FLOAT",
		"CADENA", "RUNE", "ESPACIO", "COMENTARIO", "PARABRE", "PARCIERRA", "LLAVEABRE",
		"LLAVECIERRA", "CORABRE", "CORCIERRA", "PUNTO_COMA", "DOS_PUNTOS", "SIGNO_MULTI",
		"SIGNO_DIV", "SIGNO_ASIGNACION_INCREMENTO", "SIGNO_ASIGNACION_DECREMENTO",
		"SIGNO_INCREMENTO", "SIGNO_DECREMENTO", "SIGNO_MAS", "SIGNO_MENOS",
		"SIGNO_AND", "SIGNO_OR", "SIGNO_NO_IGUAL", "SIGNO_NOT", "SIGNO_IGUALDAD",
		"SIGNO_ASIGNACION", "SIGNO_MENOSQUE_IGUAL", "SIGNO_MENOSQUE", "SIGNO_MASQUE_IGUAL",
		"SIGNO_MASQUE", "SIGNO_COMA", "SIGNO_MODULO", "ASIGNACION_INFERIDA",
		"PUNTO",
	}
	staticData.RuleNames = []string{
		"prog", "declaracionStruct", "funcionStruct_declaracion", "funcion_declaracion",
		"lista_parametros", "parametro", "lista_sentencias", "sentencia", "sentenciaBreak",
		"sentenciaReturn", "llamadafuncion", "sentencia_transferencia", "sentenciaIf",
		"sentenciaSwitch", "bloqueCase", "bloqueDefault", "sentenciaPrintln",
		"asignacion", "declaracion_slice", "asignacion_acceso_slice", "declaracion_variable",
		"expresion", "atributosStruct", "inicializacion_slice", "lista_expresion",
		"lista_slices", "tipo", "bloque", "importStmt", "creacion_instancia_struct",
		"lista_parametros_func", "param_func", "sentenciaFor", "sentencia_init",
		"sentencia_update", "incremento_variable", "llamadaIndexOf", "llamadaJoin",
		"llamadaLen", "llamadaAppend", "accesoSlice", "declaracion_matriz",
		"lista_arreglos", "asignacion_acceso_matriz", "accesoMatriz", "llamadaParseFloat",
		"llamadaTypeOf", "listaCampos", "campoStruct", "accesoInstaciaAtributo",
		"asignacion_acceso_atributo", "llamada_funcion_struct",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 695, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1, 0, 1,
		0, 1, 0, 1, 0, 4, 0, 109, 8, 0, 11, 0, 12, 0, 110, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 4, 1, 119, 8, 1, 11, 1, 12, 1, 120, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 134, 8, 2, 1, 2, 1,
		2, 3, 2, 138, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 146, 8, 3,
		1, 3, 1, 3, 3, 3, 150, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 157, 8,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 162, 8, 3, 10, 3, 12, 3, 165, 9, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 170, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 175, 8, 4, 10, 4, 12,
		4, 178, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 184, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 5, 6, 190, 8, 6, 10, 6, 12, 6, 193, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 213, 8, 7, 1, 8, 1, 8, 3, 8, 217, 8, 8, 1, 9, 1, 9, 1,
		9, 3, 9, 222, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 227, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 3, 11, 233, 8, 11, 1, 11, 1, 11, 3, 11, 237, 8, 11, 1, 11,
		1, 11, 3, 11, 241, 8, 11, 1, 11, 3, 11, 244, 8, 11, 3, 11, 246, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 264, 8, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 4, 13, 270, 8, 13, 11, 13, 12, 13, 271, 1, 13, 3, 13, 275,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 283, 8, 14, 10,
		14, 12, 14, 286, 9, 14, 1, 14, 3, 14, 289, 8, 14, 1, 15, 1, 15, 1, 15,
		3, 15, 294, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 299, 8, 16, 1, 16, 1, 16,
		3, 16, 303, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 309, 8, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 315, 8, 17, 1, 17, 3, 17, 318, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 326, 8, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 337, 8, 18, 3, 18,
		339, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 348,
		8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 353, 8, 20, 1, 20, 1, 20, 3, 20, 357,
		8, 20, 1, 20, 3, 20, 360, 8, 20, 1, 20, 3, 20, 363, 8, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 3, 20, 369, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3,
		20, 376, 8, 20, 3, 20, 378, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 436,
		8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 456,
		8, 21, 10, 21, 12, 21, 459, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 464, 8,
		22, 1, 23, 1, 23, 3, 23, 468, 8, 23, 1, 23, 1, 23, 1, 23, 3, 23, 473, 8,
		23, 1, 23, 3, 23, 476, 8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 481, 8, 24, 10,
		24, 12, 24, 484, 9, 24, 1, 25, 1, 25, 1, 25, 5, 25, 489, 8, 25, 10, 25,
		12, 25, 492, 9, 25, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 498, 8, 27, 10,
		27, 12, 27, 501, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 4, 28,
		509, 8, 28, 11, 28, 12, 28, 510, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 3, 29, 519, 8, 29, 1, 29, 1, 29, 3, 29, 523, 8, 29, 1, 30, 1, 30, 1,
		30, 5, 30, 528, 8, 30, 10, 30, 12, 30, 531, 9, 30, 1, 31, 3, 31, 534, 8,
		31, 1, 31, 1, 31, 3, 31, 538, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 554,
		8, 32, 1, 33, 1, 33, 3, 33, 558, 8, 33, 1, 34, 1, 34, 3, 34, 562, 8, 34,
		1, 35, 1, 35, 1, 35, 3, 35, 567, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 3, 41, 608, 8, 41, 1, 41, 3, 41, 611, 8, 41, 1, 42,
		1, 42, 1, 42, 3, 42, 616, 8, 42, 1, 42, 1, 42, 1, 42, 3, 42, 621, 8, 42,
		5, 42, 623, 8, 42, 10, 42, 12, 42, 626, 9, 42, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 640,
		8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 5, 47, 663, 8, 47, 10, 47, 12, 47, 666, 9, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 3, 48, 672, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 684, 8, 50, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 3, 51, 691, 8, 51, 1, 51, 1, 51, 1, 51, 0, 1, 42, 52, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 0, 9, 1, 0, 53, 54,
		2, 0, 51, 52, 70, 70, 1, 0, 57, 58, 2, 0, 61, 61, 63, 63, 1, 0, 65, 66,
		1, 0, 67, 68, 1, 0, 59, 60, 2, 0, 8, 12, 36, 36, 1, 0, 55, 56, 764, 0,
		108, 1, 0, 0, 0, 2, 114, 1, 0, 0, 0, 4, 124, 1, 0, 0, 0, 6, 169, 1, 0,
		0, 0, 8, 171, 1, 0, 0, 0, 10, 179, 1, 0, 0, 0, 12, 187, 1, 0, 0, 0, 14,
		212, 1, 0, 0, 0, 16, 214, 1, 0, 0, 0, 18, 218, 1, 0, 0, 0, 20, 223, 1,
		0, 0, 0, 22, 245, 1, 0, 0, 0, 24, 263, 1, 0, 0, 0, 26, 265, 1, 0, 0, 0,
		28, 278, 1, 0, 0, 0, 30, 290, 1, 0, 0, 0, 32, 295, 1, 0, 0, 0, 34, 317,
		1, 0, 0, 0, 36, 338, 1, 0, 0, 0, 38, 340, 1, 0, 0, 0, 40, 377, 1, 0, 0,
		0, 42, 435, 1, 0, 0, 0, 44, 460, 1, 0, 0, 0, 46, 475, 1, 0, 0, 0, 48, 477,
		1, 0, 0, 0, 50, 485, 1, 0, 0, 0, 52, 493, 1, 0, 0, 0, 54, 495, 1, 0, 0,
		0, 56, 504, 1, 0, 0, 0, 58, 512, 1, 0, 0, 0, 60, 524, 1, 0, 0, 0, 62, 537,
		1, 0, 0, 0, 64, 553, 1, 0, 0, 0, 66, 557, 1, 0, 0, 0, 68, 561, 1, 0, 0,
		0, 70, 563, 1, 0, 0, 0, 72, 568, 1, 0, 0, 0, 74, 575, 1, 0, 0, 0, 76, 582,
		1, 0, 0, 0, 78, 587, 1, 0, 0, 0, 80, 594, 1, 0, 0, 0, 82, 599, 1, 0, 0,
		0, 84, 612, 1, 0, 0, 0, 86, 629, 1, 0, 0, 0, 88, 641, 1, 0, 0, 0, 90, 649,
		1, 0, 0, 0, 92, 654, 1, 0, 0, 0, 94, 659, 1, 0, 0, 0, 96, 667, 1, 0, 0,
		0, 98, 673, 1, 0, 0, 0, 100, 677, 1, 0, 0, 0, 102, 685, 1, 0, 0, 0, 104,
		109, 3, 2, 1, 0, 105, 109, 3, 4, 2, 0, 106, 109, 3, 6, 3, 0, 107, 109,
		3, 54, 27, 0, 108, 104, 1, 0, 0, 0, 108, 105, 1, 0, 0, 0, 108, 106, 1,
		0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0,
		0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 0, 0, 1, 113,
		1, 1, 0, 0, 0, 114, 115, 5, 19, 0, 0, 115, 116, 5, 36, 0, 0, 116, 118,
		5, 45, 0, 0, 117, 119, 3, 44, 22, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1,
		0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0,
		0, 122, 123, 5, 46, 0, 0, 123, 3, 1, 0, 0, 0, 124, 125, 5, 4, 0, 0, 125,
		126, 5, 43, 0, 0, 126, 127, 5, 36, 0, 0, 127, 128, 5, 51, 0, 0, 128, 129,
		5, 36, 0, 0, 129, 130, 5, 44, 0, 0, 130, 131, 5, 36, 0, 0, 131, 133, 5,
		43, 0, 0, 132, 134, 3, 8, 4, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 135, 1, 0, 0, 0, 135, 137, 5, 44, 0, 0, 136, 138, 3, 52, 26, 0,
		137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 3, 54, 27, 0, 140, 5, 1, 0, 0, 0, 141, 142, 5, 4, 0, 0, 142, 143,
		5, 36, 0, 0, 143, 145, 5, 43, 0, 0, 144, 146, 3, 8, 4, 0, 145, 144, 1,
		0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 5, 44, 0,
		0, 148, 150, 3, 52, 26, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0,
		150, 151, 1, 0, 0, 0, 151, 170, 3, 54, 27, 0, 152, 153, 5, 4, 0, 0, 153,
		154, 5, 36, 0, 0, 154, 156, 5, 43, 0, 0, 155, 157, 3, 8, 4, 0, 156, 155,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 163, 5, 44,
		0, 0, 159, 160, 5, 47, 0, 0, 160, 162, 5, 48, 0, 0, 161, 159, 1, 0, 0,
		0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 3, 52, 26, 0, 167, 168,
		3, 54, 27, 0, 168, 170, 1, 0, 0, 0, 169, 141, 1, 0, 0, 0, 169, 152, 1,
		0, 0, 0, 170, 7, 1, 0, 0, 0, 171, 176, 3, 10, 5, 0, 172, 173, 5, 69, 0,
		0, 173, 175, 3, 10, 5, 0, 174, 172, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 9, 1, 0, 0, 0, 178, 176, 1,
		0, 0, 0, 179, 183, 5, 36, 0, 0, 180, 184, 5, 51, 0, 0, 181, 182, 5, 47,
		0, 0, 182, 184, 5, 48, 0, 0, 183, 180, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 3, 52, 26, 0, 186,
		11, 1, 0, 0, 0, 187, 191, 3, 14, 7, 0, 188, 190, 3, 14, 7, 0, 189, 188,
		1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0,
		0, 0, 192, 13, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 213, 3, 54, 27, 0,
		195, 213, 3, 40, 20, 0, 196, 213, 3, 34, 17, 0, 197, 213, 3, 36, 18, 0,
		198, 213, 3, 38, 19, 0, 199, 213, 3, 24, 12, 0, 200, 213, 3, 26, 13, 0,
		201, 213, 3, 22, 11, 0, 202, 213, 3, 32, 16, 0, 203, 213, 3, 58, 29, 0,
		204, 213, 3, 100, 50, 0, 205, 213, 3, 102, 51, 0, 206, 213, 3, 20, 10,
		0, 207, 213, 3, 64, 32, 0, 208, 213, 3, 82, 41, 0, 209, 213, 3, 86, 43,
		0, 210, 213, 3, 18, 9, 0, 211, 213, 3, 16, 8, 0, 212, 194, 1, 0, 0, 0,
		212, 195, 1, 0, 0, 0, 212, 196, 1, 0, 0, 0, 212, 197, 1, 0, 0, 0, 212,
		198, 1, 0, 0, 0, 212, 199, 1, 0, 0, 0, 212, 200, 1, 0, 0, 0, 212, 201,
		1, 0, 0, 0, 212, 202, 1, 0, 0, 0, 212, 203, 1, 0, 0, 0, 212, 204, 1, 0,
		0, 0, 212, 205, 1, 0, 0, 0, 212, 206, 1, 0, 0, 0, 212, 207, 1, 0, 0, 0,
		212, 208, 1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212,
		211, 1, 0, 0, 0, 213, 15, 1, 0, 0, 0, 214, 216, 5, 28, 0, 0, 215, 217,
		5, 49, 0, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 17, 1, 0,
		0, 0, 218, 219, 5, 27, 0, 0, 219, 221, 3, 42, 21, 0, 220, 222, 5, 49, 0,
		0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 19, 1, 0, 0, 0, 223,
		224, 5, 36, 0, 0, 224, 226, 5, 43, 0, 0, 225, 227, 3, 60, 30, 0, 226, 225,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 5, 44,
		0, 0, 229, 21, 1, 0, 0, 0, 230, 232, 5, 28, 0, 0, 231, 233, 5, 49, 0, 0,
		232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 246, 1, 0, 0, 0, 234,
		236, 5, 29, 0, 0, 235, 237, 5, 49, 0, 0, 236, 235, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 246, 1, 0, 0, 0, 238, 240, 5, 27, 0, 0, 239, 241, 3, 42,
		21, 0, 240, 239, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0,
		242, 244, 5, 49, 0, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		246, 1, 0, 0, 0, 245, 230, 1, 0, 0, 0, 245, 234, 1, 0, 0, 0, 245, 238,
		1, 0, 0, 0, 246, 23, 1, 0, 0, 0, 247, 248, 5, 21, 0, 0, 248, 249, 3, 42,
		21, 0, 249, 250, 3, 54, 27, 0, 250, 264, 1, 0, 0, 0, 251, 252, 5, 21, 0,
		0, 252, 253, 3, 42, 21, 0, 253, 254, 3, 54, 27, 0, 254, 255, 5, 22, 0,
		0, 255, 256, 3, 54, 27, 0, 256, 264, 1, 0, 0, 0, 257, 258, 5, 21, 0, 0,
		258, 259, 3, 42, 21, 0, 259, 260, 3, 54, 27, 0, 260, 261, 5, 22, 0, 0,
		261, 262, 3, 24, 12, 0, 262, 264, 1, 0, 0, 0, 263, 247, 1, 0, 0, 0, 263,
		251, 1, 0, 0, 0, 263, 257, 1, 0, 0, 0, 264, 25, 1, 0, 0, 0, 265, 266, 5,
		24, 0, 0, 266, 267, 3, 42, 21, 0, 267, 269, 5, 45, 0, 0, 268, 270, 3, 28,
		14, 0, 269, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0,
		271, 272, 1, 0, 0, 0, 272, 274, 1, 0, 0, 0, 273, 275, 3, 30, 15, 0, 274,
		273, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277,
		5, 46, 0, 0, 277, 27, 1, 0, 0, 0, 278, 279, 5, 25, 0, 0, 279, 280, 3, 42,
		21, 0, 280, 284, 5, 50, 0, 0, 281, 283, 3, 12, 6, 0, 282, 281, 1, 0, 0,
		0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285,
		288, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 289, 5, 49, 0, 0, 288, 287,
		1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 29, 1, 0, 0, 0, 290, 291, 5, 26,
		0, 0, 291, 293, 5, 50, 0, 0, 292, 294, 3, 12, 6, 0, 293, 292, 1, 0, 0,
		0, 293, 294, 1, 0, 0, 0, 294, 31, 1, 0, 0, 0, 295, 296, 5, 32, 0, 0, 296,
		298, 5, 43, 0, 0, 297, 299, 3, 48, 24, 0, 298, 297, 1, 0, 0, 0, 298, 299,
		1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 302, 5, 44, 0, 0, 301, 303, 5, 49,
		0, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 33, 1, 0, 0, 0,
		304, 305, 5, 36, 0, 0, 305, 306, 5, 64, 0, 0, 306, 308, 3, 42, 21, 0, 307,
		309, 5, 49, 0, 0, 308, 307, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 318,
		1, 0, 0, 0, 310, 311, 5, 36, 0, 0, 311, 312, 7, 0, 0, 0, 312, 314, 3, 42,
		21, 0, 313, 315, 5, 49, 0, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 318, 1, 0, 0, 0, 316, 318, 3, 38, 19, 0, 317, 304, 1, 0, 0, 0,
		317, 310, 1, 0, 0, 0, 317, 316, 1, 0, 0, 0, 318, 35, 1, 0, 0, 0, 319, 320,
		5, 7, 0, 0, 320, 321, 5, 36, 0, 0, 321, 322, 5, 47, 0, 0, 322, 323, 5,
		48, 0, 0, 323, 325, 3, 52, 26, 0, 324, 326, 5, 49, 0, 0, 325, 324, 1, 0,
		0, 0, 325, 326, 1, 0, 0, 0, 326, 339, 1, 0, 0, 0, 327, 328, 5, 36, 0, 0,
		328, 329, 5, 64, 0, 0, 329, 330, 5, 47, 0, 0, 330, 331, 5, 48, 0, 0, 331,
		332, 3, 52, 26, 0, 332, 333, 5, 45, 0, 0, 333, 334, 3, 48, 24, 0, 334,
		336, 5, 46, 0, 0, 335, 337, 5, 49, 0, 0, 336, 335, 1, 0, 0, 0, 336, 337,
		1, 0, 0, 0, 337, 339, 1, 0, 0, 0, 338, 319, 1, 0, 0, 0, 338, 327, 1, 0,
		0, 0, 339, 37, 1, 0, 0, 0, 340, 341, 5, 36, 0, 0, 341, 342, 5, 47, 0, 0,
		342, 343, 3, 42, 21, 0, 343, 344, 5, 48, 0, 0, 344, 345, 5, 64, 0, 0, 345,
		347, 3, 42, 21, 0, 346, 348, 5, 49, 0, 0, 347, 346, 1, 0, 0, 0, 347, 348,
		1, 0, 0, 0, 348, 39, 1, 0, 0, 0, 349, 350, 5, 7, 0, 0, 350, 352, 5, 36,
		0, 0, 351, 353, 3, 52, 26, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0, 0,
		0, 353, 356, 1, 0, 0, 0, 354, 355, 5, 64, 0, 0, 355, 357, 3, 42, 21, 0,
		356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0, 358,
		360, 5, 49, 0, 0, 359, 358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 378,
		1, 0, 0, 0, 361, 363, 5, 7, 0, 0, 362, 361, 1, 0, 0, 0, 362, 363, 1, 0,
		0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 5, 36, 0, 0, 365, 366, 5, 71, 0,
		0, 366, 368, 3, 42, 21, 0, 367, 369, 5, 49, 0, 0, 368, 367, 1, 0, 0, 0,
		368, 369, 1, 0, 0, 0, 369, 378, 1, 0, 0, 0, 370, 371, 5, 36, 0, 0, 371,
		372, 3, 52, 26, 0, 372, 373, 5, 64, 0, 0, 373, 375, 3, 42, 21, 0, 374,
		376, 5, 49, 0, 0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378,
		1, 0, 0, 0, 377, 349, 1, 0, 0, 0, 377, 362, 1, 0, 0, 0, 377, 370, 1, 0,
		0, 0, 378, 41, 1, 0, 0, 0, 379, 380, 6, 21, -1, 0, 380, 381, 5, 58, 0,
		0, 381, 382, 5, 43, 0, 0, 382, 383, 3, 42, 21, 0, 383, 384, 5, 44, 0, 0,
		384, 436, 1, 0, 0, 0, 385, 386, 5, 43, 0, 0, 386, 387, 3, 42, 21, 0, 387,
		388, 5, 44, 0, 0, 388, 436, 1, 0, 0, 0, 389, 390, 5, 47, 0, 0, 390, 391,
		3, 48, 24, 0, 391, 392, 5, 48, 0, 0, 392, 436, 1, 0, 0, 0, 393, 394, 5,
		62, 0, 0, 394, 436, 3, 42, 21, 21, 395, 436, 5, 37, 0, 0, 396, 436, 5,
		38, 0, 0, 397, 436, 5, 1, 0, 0, 398, 436, 5, 2, 0, 0, 399, 436, 5, 39,
		0, 0, 400, 436, 5, 40, 0, 0, 401, 436, 5, 20, 0, 0, 402, 436, 5, 36, 0,
		0, 403, 404, 5, 36, 0, 0, 404, 405, 5, 72, 0, 0, 405, 406, 5, 36, 0, 0,
		406, 407, 5, 43, 0, 0, 407, 436, 5, 44, 0, 0, 408, 436, 3, 20, 10, 0, 409,
		436, 3, 72, 36, 0, 410, 436, 3, 74, 37, 0, 411, 436, 3, 76, 38, 0, 412,
		436, 3, 78, 39, 0, 413, 436, 3, 88, 44, 0, 414, 436, 3, 98, 49, 0, 415,
		416, 5, 36, 0, 0, 416, 417, 5, 47, 0, 0, 417, 418, 3, 42, 21, 0, 418, 419,
		5, 48, 0, 0, 419, 436, 1, 0, 0, 0, 420, 421, 5, 36, 0, 0, 421, 422, 5,
		47, 0, 0, 422, 423, 3, 42, 21, 0, 423, 424, 5, 48, 0, 0, 424, 425, 5, 47,
		0, 0, 425, 426, 3, 42, 21, 0, 426, 427, 5, 48, 0, 0, 427, 436, 1, 0, 0,
		0, 428, 429, 5, 47, 0, 0, 429, 430, 5, 48, 0, 0, 430, 431, 3, 52, 26, 0,
		431, 432, 5, 45, 0, 0, 432, 433, 3, 48, 24, 0, 433, 434, 5, 46, 0, 0, 434,
		436, 1, 0, 0, 0, 435, 379, 1, 0, 0, 0, 435, 385, 1, 0, 0, 0, 435, 389,
		1, 0, 0, 0, 435, 393, 1, 0, 0, 0, 435, 395, 1, 0, 0, 0, 435, 396, 1, 0,
		0, 0, 435, 397, 1, 0, 0, 0, 435, 398, 1, 0, 0, 0, 435, 399, 1, 0, 0, 0,
		435, 400, 1, 0, 0, 0, 435, 401, 1, 0, 0, 0, 435, 402, 1, 0, 0, 0, 435,
		403, 1, 0, 0, 0, 435, 408, 1, 0, 0, 0, 435, 409, 1, 0, 0, 0, 435, 410,
		1, 0, 0, 0, 435, 411, 1, 0, 0, 0, 435, 412, 1, 0, 0, 0, 435, 413, 1, 0,
		0, 0, 435, 414, 1, 0, 0, 0, 435, 415, 1, 0, 0, 0, 435, 420, 1, 0, 0, 0,
		435, 428, 1, 0, 0, 0, 436, 457, 1, 0, 0, 0, 437, 438, 10, 28, 0, 0, 438,
		439, 7, 1, 0, 0, 439, 456, 3, 42, 21, 29, 440, 441, 10, 27, 0, 0, 441,
		442, 7, 2, 0, 0, 442, 456, 3, 42, 21, 28, 443, 444, 10, 24, 0, 0, 444,
		445, 7, 3, 0, 0, 445, 456, 3, 42, 21, 25, 446, 447, 10, 23, 0, 0, 447,
		448, 7, 4, 0, 0, 448, 456, 3, 42, 21, 24, 449, 450, 10, 22, 0, 0, 450,
		451, 7, 5, 0, 0, 451, 456, 3, 42, 21, 23, 452, 453, 10, 20, 0, 0, 453,
		454, 7, 6, 0, 0, 454, 456, 3, 42, 21, 21, 455, 437, 1, 0, 0, 0, 455, 440,
		1, 0, 0, 0, 455, 443, 1, 0, 0, 0, 455, 446, 1, 0, 0, 0, 455, 449, 1, 0,
		0, 0, 455, 452, 1, 0, 0, 0, 456, 459, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0,
		457, 458, 1, 0, 0, 0, 458, 43, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 460, 461,
		3, 52, 26, 0, 461, 463, 5, 36, 0, 0, 462, 464, 5, 49, 0, 0, 463, 462, 1,
		0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 45, 1, 0, 0, 0, 465, 467, 5, 47, 0,
		0, 466, 468, 3, 48, 24, 0, 467, 466, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0,
		468, 469, 1, 0, 0, 0, 469, 476, 5, 48, 0, 0, 470, 472, 5, 47, 0, 0, 471,
		473, 3, 50, 25, 0, 472, 471, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 474,
		1, 0, 0, 0, 474, 476, 5, 48, 0, 0, 475, 465, 1, 0, 0, 0, 475, 470, 1, 0,
		0, 0, 476, 47, 1, 0, 0, 0, 477, 482, 3, 42, 21, 0, 478, 479, 5, 69, 0,
		0, 479, 481, 3, 42, 21, 0, 480, 478, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0,
		482, 480, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 49, 1, 0, 0, 0, 484, 482,
		1, 0, 0, 0, 485, 490, 3, 46, 23, 0, 486, 487, 5, 69, 0, 0, 487, 489, 3,
		46, 23, 0, 488, 486, 1, 0, 0, 0, 489, 492, 1, 0, 0, 0, 490, 488, 1, 0,
		0, 0, 490, 491, 1, 0, 0, 0, 491, 51, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0,
		493, 494, 7, 7, 0, 0, 494, 53, 1, 0, 0, 0, 495, 499, 5, 45, 0, 0, 496,
		498, 3, 12, 6, 0, 497, 496, 1, 0, 0, 0, 498, 501, 1, 0, 0, 0, 499, 497,
		1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 502, 1, 0, 0, 0, 501, 499, 1, 0,
		0, 0, 502, 503, 5, 46, 0, 0, 503, 55, 1, 0, 0, 0, 504, 505, 5, 5, 0, 0,
		505, 506, 5, 39, 0, 0, 506, 508, 1, 0, 0, 0, 507, 509, 5, 49, 0, 0, 508,
		507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 510, 511,
		1, 0, 0, 0, 511, 57, 1, 0, 0, 0, 512, 513, 5, 7, 0, 0, 513, 514, 5, 36,
		0, 0, 514, 515, 5, 71, 0, 0, 515, 516, 5, 36, 0, 0, 516, 518, 5, 45, 0,
		0, 517, 519, 3, 94, 47, 0, 518, 517, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0,
		519, 520, 1, 0, 0, 0, 520, 522, 5, 46, 0, 0, 521, 523, 5, 49, 0, 0, 522,
		521, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 59, 1, 0, 0, 0, 524, 529, 3,
		42, 21, 0, 525, 526, 5, 69, 0, 0, 526, 528, 3, 42, 21, 0, 527, 525, 1,
		0, 0, 0, 528, 531, 1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 529, 530, 1, 0, 0,
		0, 530, 61, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 532, 534, 5, 3, 0, 0, 533,
		532, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 538,
		5, 36, 0, 0, 536, 538, 3, 42, 21, 0, 537, 533, 1, 0, 0, 0, 537, 536, 1,
		0, 0, 0, 538, 63, 1, 0, 0, 0, 539, 540, 5, 23, 0, 0, 540, 541, 3, 42, 21,
		0, 541, 542, 3, 54, 27, 0, 542, 554, 1, 0, 0, 0, 543, 544, 5, 23, 0, 0,
		544, 545, 5, 43, 0, 0, 545, 546, 3, 66, 33, 0, 546, 547, 5, 49, 0, 0, 547,
		548, 3, 42, 21, 0, 548, 549, 5, 49, 0, 0, 549, 550, 3, 68, 34, 0, 550,
		551, 5, 44, 0, 0, 551, 552, 3, 54, 27, 0, 552, 554, 1, 0, 0, 0, 553, 539,
		1, 0, 0, 0, 553, 543, 1, 0, 0, 0, 554, 65, 1, 0, 0, 0, 555, 558, 3, 40,
		20, 0, 556, 558, 3, 34, 17, 0, 557, 555, 1, 0, 0, 0, 557, 556, 1, 0, 0,
		0, 558, 67, 1, 0, 0, 0, 559, 562, 3, 34, 17, 0, 560, 562, 3, 70, 35, 0,
		561, 559, 1, 0, 0, 0, 561, 560, 1, 0, 0, 0, 562, 69, 1, 0, 0, 0, 563, 564,
		5, 36, 0, 0, 564, 566, 7, 8, 0, 0, 565, 567, 5, 49, 0, 0, 566, 565, 1,
		0, 0, 0, 566, 567, 1, 0, 0, 0, 567, 71, 1, 0, 0, 0, 568, 569, 5, 14, 0,
		0, 569, 570, 5, 43, 0, 0, 570, 571, 5, 36, 0, 0, 571, 572, 5, 69, 0, 0,
		572, 573, 3, 42, 21, 0, 573, 574, 5, 44, 0, 0, 574, 73, 1, 0, 0, 0, 575,
		576, 5, 16, 0, 0, 576, 577, 5, 43, 0, 0, 577, 578, 5, 36, 0, 0, 578, 579,
		5, 69, 0, 0, 579, 580, 3, 42, 21, 0, 580, 581, 5, 44, 0, 0, 581, 75, 1,
		0, 0, 0, 582, 583, 5, 17, 0, 0, 583, 584, 5, 43, 0, 0, 584, 585, 5, 36,
		0, 0, 585, 586, 5, 44, 0, 0, 586, 77, 1, 0, 0, 0, 587, 588, 5, 18, 0, 0,
		588, 589, 5, 43, 0, 0, 589, 590, 5, 36, 0, 0, 590, 591, 5, 69, 0, 0, 591,
		592, 3, 42, 21, 0, 592, 593, 5, 44, 0, 0, 593, 79, 1, 0, 0, 0, 594, 595,
		5, 36, 0, 0, 595, 596, 5, 47, 0, 0, 596, 597, 3, 42, 21, 0, 597, 598, 5,
		48, 0, 0, 598, 81, 1, 0, 0, 0, 599, 600, 5, 36, 0, 0, 600, 601, 5, 71,
		0, 0, 601, 602, 5, 45, 0, 0, 602, 603, 5, 46, 0, 0, 603, 604, 5, 45, 0,
		0, 604, 605, 5, 46, 0, 0, 605, 607, 3, 52, 26, 0, 606, 608, 3, 84, 42,
		0, 607, 606, 1, 0, 0, 0, 607, 608, 1, 0, 0, 0, 608, 610, 1, 0, 0, 0, 609,
		611, 5, 49, 0, 0, 610, 609, 1, 0, 0, 0, 610, 611, 1, 0, 0, 0, 611, 83,
		1, 0, 0, 0, 612, 615, 5, 45, 0, 0, 613, 616, 3, 84, 42, 0, 614, 616, 3,
		42, 21, 0, 615, 613, 1, 0, 0, 0, 615, 614, 1, 0, 0, 0, 616, 624, 1, 0,
		0, 0, 617, 620, 5, 69, 0, 0, 618, 621, 3, 84, 42, 0, 619, 621, 3, 42, 21,
		0, 620, 618, 1, 0, 0, 0, 620, 619, 1, 0, 0, 0, 621, 623, 1, 0, 0, 0, 622,
		617, 1, 0, 0, 0, 623, 626, 1, 0, 0, 0, 624, 622, 1, 0, 0, 0, 624, 625,
		1, 0, 0, 0, 625, 627, 1, 0, 0, 0, 626, 624, 1, 0, 0, 0, 627, 628, 5, 46,
		0, 0, 628, 85, 1, 0, 0, 0, 629, 630, 5, 36, 0, 0, 630, 631, 5, 47, 0, 0,
		631, 632, 3, 42, 21, 0, 632, 633, 5, 48, 0, 0, 633, 634, 5, 47, 0, 0, 634,
		635, 3, 42, 21, 0, 635, 636, 5, 48, 0, 0, 636, 637, 5, 64, 0, 0, 637, 639,
		3, 42, 21, 0, 638, 640, 5, 49, 0, 0, 639, 638, 1, 0, 0, 0, 639, 640, 1,
		0, 0, 0, 640, 87, 1, 0, 0, 0, 641, 642, 5, 36, 0, 0, 642, 643, 5, 47, 0,
		0, 643, 644, 3, 42, 21, 0, 644, 645, 5, 48, 0, 0, 645, 646, 5, 47, 0, 0,
		646, 647, 3, 42, 21, 0, 647, 648, 5, 48, 0, 0, 648, 89, 1, 0, 0, 0, 649,
		650, 5, 34, 0, 0, 650, 651, 5, 43, 0, 0, 651, 652, 3, 42, 21, 0, 652, 653,
		5, 44, 0, 0, 653, 91, 1, 0, 0, 0, 654, 655, 5, 35, 0, 0, 655, 656, 5, 43,
		0, 0, 656, 657, 3, 42, 21, 0, 657, 658, 5, 44, 0, 0, 658, 93, 1, 0, 0,
		0, 659, 664, 3, 96, 48, 0, 660, 661, 5, 69, 0, 0, 661, 663, 3, 96, 48,
		0, 662, 660, 1, 0, 0, 0, 663, 666, 1, 0, 0, 0, 664, 662, 1, 0, 0, 0, 664,
		665, 1, 0, 0, 0, 665, 95, 1, 0, 0, 0, 666, 664, 1, 0, 0, 0, 667, 668, 5,
		36, 0, 0, 668, 669, 5, 50, 0, 0, 669, 671, 3, 42, 21, 0, 670, 672, 5, 49,
		0, 0, 671, 670, 1, 0, 0, 0, 671, 672, 1, 0, 0, 0, 672, 97, 1, 0, 0, 0,
		673, 674, 5, 36, 0, 0, 674, 675, 5, 72, 0, 0, 675, 676, 5, 36, 0, 0, 676,
		99, 1, 0, 0, 0, 677, 678, 5, 36, 0, 0, 678, 679, 5, 72, 0, 0, 679, 680,
		5, 36, 0, 0, 680, 681, 5, 64, 0, 0, 681, 683, 3, 42, 21, 0, 682, 684, 5,
		49, 0, 0, 683, 682, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 101, 1, 0, 0,
		0, 685, 686, 5, 36, 0, 0, 686, 687, 5, 72, 0, 0, 687, 688, 5, 36, 0, 0,
		688, 690, 5, 43, 0, 0, 689, 691, 3, 60, 30, 0, 690, 689, 1, 0, 0, 0, 690,
		691, 1, 0, 0, 0, 691, 692, 1, 0, 0, 0, 692, 693, 5, 44, 0, 0, 693, 103,
		1, 0, 0, 0, 74, 108, 110, 120, 133, 137, 145, 149, 156, 163, 169, 176,
		183, 191, 212, 216, 221, 226, 232, 236, 240, 243, 245, 263, 271, 274, 284,
		288, 293, 298, 302, 308, 314, 317, 325, 336, 338, 347, 352, 356, 359, 362,
		368, 375, 377, 435, 455, 457, 463, 467, 472, 475, 482, 490, 499, 510, 518,
		522, 529, 533, 537, 553, 557, 561, 566, 607, 610, 615, 620, 624, 639, 664,
		671, 683, 690,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// gramaticaParserInit initializes any static state used to implement gramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewgramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgramaticaParser(input antlr.TokenStream) *gramaticaParser {
	GramaticaParserInit()
	this := new(gramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gramatica.g4"

	return this
}

// gramaticaParser tokens.
const (
	gramaticaParserEOF                         = antlr.TokenEOF
	gramaticaParserT__0                        = 1
	gramaticaParserT__1                        = 2
	gramaticaParserT__2                        = 3
	gramaticaParserRESFUNC                     = 4
	gramaticaParserRESIMPORT                   = 5
	gramaticaParserRESVAR                      = 6
	gramaticaParserRESMUT                      = 7
	gramaticaParserRESINT                      = 8
	gramaticaParserRESFLOAT                    = 9
	gramaticaParserRESSTRING                   = 10
	gramaticaParserRESBOOL                     = 11
	gramaticaParserRESRUNE                     = 12
	gramaticaParserRESSLICE                    = 13
	gramaticaParserRESINDEX                    = 14
	gramaticaParserRESSTRINGS                  = 15
	gramaticaParserRESJOIN                     = 16
	gramaticaParserRESLEN                      = 17
	gramaticaParserRESAPPEND                   = 18
	gramaticaParserRESSTRUCT                   = 19
	gramaticaParserRESNIL                      = 20
	gramaticaParserRESIF                       = 21
	gramaticaParserRESELSE                     = 22
	gramaticaParserRESFOR                      = 23
	gramaticaParserRESSWITCH                   = 24
	gramaticaParserRESCASE                     = 25
	gramaticaParserRESDEFAULT                  = 26
	gramaticaParserRESRETURN                   = 27
	gramaticaParserRESBREAK                    = 28
	gramaticaParserRESCONTINUE                 = 29
	gramaticaParserRESIN                       = 30
	gramaticaParserRESFMT                      = 31
	gramaticaParserRESPRINTLN                  = 32
	gramaticaParserRESATOI                     = 33
	gramaticaParserRESPARSEFLOAT               = 34
	gramaticaParserRESTYPEOF                   = 35
	gramaticaParserIDENTIFICADOR               = 36
	gramaticaParserINT                         = 37
	gramaticaParserFLOAT                       = 38
	gramaticaParserCADENA                      = 39
	gramaticaParserRUNE                        = 40
	gramaticaParserESPACIO                     = 41
	gramaticaParserCOMENTARIO                  = 42
	gramaticaParserPARABRE                     = 43
	gramaticaParserPARCIERRA                   = 44
	gramaticaParserLLAVEABRE                   = 45
	gramaticaParserLLAVECIERRA                 = 46
	gramaticaParserCORABRE                     = 47
	gramaticaParserCORCIERRA                   = 48
	gramaticaParserPUNTO_COMA                  = 49
	gramaticaParserDOS_PUNTOS                  = 50
	gramaticaParserSIGNO_MULTI                 = 51
	gramaticaParserSIGNO_DIV                   = 52
	gramaticaParserSIGNO_ASIGNACION_INCREMENTO = 53
	gramaticaParserSIGNO_ASIGNACION_DECREMENTO = 54
	gramaticaParserSIGNO_INCREMENTO            = 55
	gramaticaParserSIGNO_DECREMENTO            = 56
	gramaticaParserSIGNO_MAS                   = 57
	gramaticaParserSIGNO_MENOS                 = 58
	gramaticaParserSIGNO_AND                   = 59
	gramaticaParserSIGNO_OR                    = 60
	gramaticaParserSIGNO_NO_IGUAL              = 61
	gramaticaParserSIGNO_NOT                   = 62
	gramaticaParserSIGNO_IGUALDAD              = 63
	gramaticaParserSIGNO_ASIGNACION            = 64
	gramaticaParserSIGNO_MENOSQUE_IGUAL        = 65
	gramaticaParserSIGNO_MENOSQUE              = 66
	gramaticaParserSIGNO_MASQUE_IGUAL          = 67
	gramaticaParserSIGNO_MASQUE                = 68
	gramaticaParserSIGNO_COMA                  = 69
	gramaticaParserSIGNO_MODULO                = 70
	gramaticaParserASIGNACION_INFERIDA         = 71
	gramaticaParserPUNTO                       = 72
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_prog                       = 0
	gramaticaParserRULE_declaracionStruct          = 1
	gramaticaParserRULE_funcionStruct_declaracion  = 2
	gramaticaParserRULE_funcion_declaracion        = 3
	gramaticaParserRULE_lista_parametros           = 4
	gramaticaParserRULE_parametro                  = 5
	gramaticaParserRULE_lista_sentencias           = 6
	gramaticaParserRULE_sentencia                  = 7
	gramaticaParserRULE_sentenciaBreak             = 8
	gramaticaParserRULE_sentenciaReturn            = 9
	gramaticaParserRULE_llamadafuncion             = 10
	gramaticaParserRULE_sentencia_transferencia    = 11
	gramaticaParserRULE_sentenciaIf                = 12
	gramaticaParserRULE_sentenciaSwitch            = 13
	gramaticaParserRULE_bloqueCase                 = 14
	gramaticaParserRULE_bloqueDefault              = 15
	gramaticaParserRULE_sentenciaPrintln           = 16
	gramaticaParserRULE_asignacion                 = 17
	gramaticaParserRULE_declaracion_slice          = 18
	gramaticaParserRULE_asignacion_acceso_slice    = 19
	gramaticaParserRULE_declaracion_variable       = 20
	gramaticaParserRULE_expresion                  = 21
	gramaticaParserRULE_atributosStruct            = 22
	gramaticaParserRULE_inicializacion_slice       = 23
	gramaticaParserRULE_lista_expresion            = 24
	gramaticaParserRULE_lista_slices               = 25
	gramaticaParserRULE_tipo                       = 26
	gramaticaParserRULE_bloque                     = 27
	gramaticaParserRULE_importStmt                 = 28
	gramaticaParserRULE_creacion_instancia_struct  = 29
	gramaticaParserRULE_lista_parametros_func      = 30
	gramaticaParserRULE_param_func                 = 31
	gramaticaParserRULE_sentenciaFor               = 32
	gramaticaParserRULE_sentencia_init             = 33
	gramaticaParserRULE_sentencia_update           = 34
	gramaticaParserRULE_incremento_variable        = 35
	gramaticaParserRULE_llamadaIndexOf             = 36
	gramaticaParserRULE_llamadaJoin                = 37
	gramaticaParserRULE_llamadaLen                 = 38
	gramaticaParserRULE_llamadaAppend              = 39
	gramaticaParserRULE_accesoSlice                = 40
	gramaticaParserRULE_declaracion_matriz         = 41
	gramaticaParserRULE_lista_arreglos             = 42
	gramaticaParserRULE_asignacion_acceso_matriz   = 43
	gramaticaParserRULE_accesoMatriz               = 44
	gramaticaParserRULE_llamadaParseFloat          = 45
	gramaticaParserRULE_llamadaTypeOf              = 46
	gramaticaParserRULE_listaCampos                = 47
	gramaticaParserRULE_campoStruct                = 48
	gramaticaParserRULE_accesoInstaciaAtributo     = 49
	gramaticaParserRULE_asignacion_acceso_atributo = 50
	gramaticaParserRULE_llamada_funcion_struct     = 51
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaracionStruct() []IDeclaracionStructContext
	DeclaracionStruct(i int) IDeclaracionStructContext
	AllFuncionStruct_declaracion() []IFuncionStruct_declaracionContext
	FuncionStruct_declaracion(i int) IFuncionStruct_declaracionContext
	AllFuncion_declaracion() []IFuncion_declaracionContext
	Funcion_declaracion(i int) IFuncion_declaracionContext
	AllBloque() []IBloqueContext
	Bloque(i int) IBloqueContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserEOF, 0)
}

func (s *ProgContext) AllDeclaracionStruct() []IDeclaracionStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracionStructContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracionStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracionStructContext); ok {
			tst[i] = t.(IDeclaracionStructContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) DeclaracionStruct(i int) IDeclaracionStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionStructContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionStructContext)
}

func (s *ProgContext) AllFuncionStruct_declaracion() []IFuncionStruct_declaracionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncionStruct_declaracionContext); ok {
			len++
		}
	}

	tst := make([]IFuncionStruct_declaracionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncionStruct_declaracionContext); ok {
			tst[i] = t.(IFuncionStruct_declaracionContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) FuncionStruct_declaracion(i int) IFuncionStruct_declaracionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionStruct_declaracionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionStruct_declaracionContext)
}

func (s *ProgContext) AllFuncion_declaracion() []IFuncion_declaracionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncion_declaracionContext); ok {
			len++
		}
	}

	tst := make([]IFuncion_declaracionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncion_declaracionContext); ok {
			tst[i] = t.(IFuncion_declaracionContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Funcion_declaracion(i int) IFuncion_declaracionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_declaracionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_declaracionContext)
}

func (s *ProgContext) AllBloque() []IBloqueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueContext); ok {
			len++
		}
	}

	tst := make([]IBloqueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueContext); ok {
			tst[i] = t.(IBloqueContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Bloque(i int) IBloqueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gramaticaParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35184372613136) != 0) {
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(104)
				p.DeclaracionStruct()
			}

		case 2:
			{
				p.SetState(105)
				p.FuncionStruct_declaracion()
			}

		case 3:
			{
				p.SetState(106)
				p.Funcion_declaracion()
			}

		case 4:
			{
				p.SetState(107)
				p.Bloque()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Match(gramaticaParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionStructContext is an interface to support dynamic dispatch.
type IDeclaracionStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomStruct returns the nomStruct token.
	GetNomStruct() antlr.Token

	// SetNomStruct sets the nomStruct token.
	SetNomStruct(antlr.Token)

	// Getter signatures
	RESSTRUCT() antlr.TerminalNode
	LLAVEABRE() antlr.TerminalNode
	LLAVECIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	AllAtributosStruct() []IAtributosStructContext
	AtributosStruct(i int) IAtributosStructContext

	// IsDeclaracionStructContext differentiates from other interfaces.
	IsDeclaracionStructContext()
}

type DeclaracionStructContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	nomStruct antlr.Token
}

func NewEmptyDeclaracionStructContext() *DeclaracionStructContext {
	var p = new(DeclaracionStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionStruct
	return p
}

func InitEmptyDeclaracionStructContext(p *DeclaracionStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionStruct
}

func (*DeclaracionStructContext) IsDeclaracionStructContext() {}

func NewDeclaracionStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionStructContext {
	var p = new(DeclaracionStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracionStruct

	return p
}

func (s *DeclaracionStructContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionStructContext) GetNomStruct() antlr.Token { return s.nomStruct }

func (s *DeclaracionStructContext) SetNomStruct(v antlr.Token) { s.nomStruct = v }

func (s *DeclaracionStructContext) RESSTRUCT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESSTRUCT, 0)
}

func (s *DeclaracionStructContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *DeclaracionStructContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *DeclaracionStructContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DeclaracionStructContext) AllAtributosStruct() []IAtributosStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtributosStructContext); ok {
			len++
		}
	}

	tst := make([]IAtributosStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtributosStructContext); ok {
			tst[i] = t.(IAtributosStructContext)
			i++
		}
	}

	return tst
}

func (s *DeclaracionStructContext) AtributosStruct(i int) IAtributosStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosStructContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosStructContext)
}

func (s *DeclaracionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracionStruct(s)
	}
}

func (s *DeclaracionStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracionStruct(s)
	}
}

func (s *DeclaracionStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracionStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DeclaracionStruct() (localctx IDeclaracionStructContext) {
	localctx = NewDeclaracionStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_declaracionStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(gramaticaParserRESSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*DeclaracionStructContext).nomStruct = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719484672) != 0) {
		{
			p.SetState(117)
			p.AtributosStruct()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncionStruct_declaracionContext is an interface to support dynamic dispatch.
type IFuncionStruct_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetNomStruct returns the nomStruct token.
	GetNomStruct() antlr.Token

	// GetNomFunc returns the nomFunc token.
	GetNomFunc() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetNomStruct sets the nomStruct token.
	SetNomStruct(antlr.Token)

	// SetNomFunc sets the nomFunc token.
	SetNomFunc(antlr.Token)

	// Getter signatures
	RESFUNC() antlr.TerminalNode
	AllPARABRE() []antlr.TerminalNode
	PARABRE(i int) antlr.TerminalNode
	SIGNO_MULTI() antlr.TerminalNode
	AllPARCIERRA() []antlr.TerminalNode
	PARCIERRA(i int) antlr.TerminalNode
	Bloque() IBloqueContext
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	Lista_parametros() ILista_parametrosContext
	Tipo() ITipoContext

	// IsFuncionStruct_declaracionContext differentiates from other interfaces.
	IsFuncionStruct_declaracionContext()
}

type FuncionStruct_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	id        antlr.Token
	nomStruct antlr.Token
	nomFunc   antlr.Token
}

func NewEmptyFuncionStruct_declaracionContext() *FuncionStruct_declaracionContext {
	var p = new(FuncionStruct_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcionStruct_declaracion
	return p
}

func InitEmptyFuncionStruct_declaracionContext(p *FuncionStruct_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcionStruct_declaracion
}

func (*FuncionStruct_declaracionContext) IsFuncionStruct_declaracionContext() {}

func NewFuncionStruct_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionStruct_declaracionContext {
	var p = new(FuncionStruct_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcionStruct_declaracion

	return p
}

func (s *FuncionStruct_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionStruct_declaracionContext) GetId() antlr.Token { return s.id }

func (s *FuncionStruct_declaracionContext) GetNomStruct() antlr.Token { return s.nomStruct }

func (s *FuncionStruct_declaracionContext) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *FuncionStruct_declaracionContext) SetId(v antlr.Token) { s.id = v }

func (s *FuncionStruct_declaracionContext) SetNomStruct(v antlr.Token) { s.nomStruct = v }

func (s *FuncionStruct_declaracionContext) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *FuncionStruct_declaracionContext) RESFUNC() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFUNC, 0)
}

func (s *FuncionStruct_declaracionContext) AllPARABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPARABRE)
}

func (s *FuncionStruct_declaracionContext) PARABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, i)
}

func (s *FuncionStruct_declaracionContext) SIGNO_MULTI() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MULTI, 0)
}

func (s *FuncionStruct_declaracionContext) AllPARCIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPARCIERRA)
}

func (s *FuncionStruct_declaracionContext) PARCIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, i)
}

func (s *FuncionStruct_declaracionContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionStruct_declaracionContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *FuncionStruct_declaracionContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *FuncionStruct_declaracionContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *FuncionStruct_declaracionContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionStruct_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionStruct_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionStruct_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncionStruct_declaracion(s)
	}
}

func (s *FuncionStruct_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncionStruct_declaracion(s)
	}
}

func (s *FuncionStruct_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFuncionStruct_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FuncionStruct_declaracion() (localctx IFuncionStruct_declaracionContext) {
	localctx = NewFuncionStruct_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gramaticaParserRULE_funcionStruct_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(gramaticaParserRESFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*FuncionStruct_declaracionContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(gramaticaParserSIGNO_MULTI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*FuncionStruct_declaracionContext).nomStruct = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*FuncionStruct_declaracionContext).nomFunc = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserIDENTIFICADOR {
		{
			p.SetState(132)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(135)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719484672) != 0 {
		{
			p.SetState(136)
			p.Tipo()
		}

	}
	{
		p.SetState(139)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_declaracionContext is an interface to support dynamic dispatch.
type IFuncion_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncion_declaracionContext differentiates from other interfaces.
	IsFuncion_declaracionContext()
}

type Funcion_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_declaracionContext() *Funcion_declaracionContext {
	var p = new(Funcion_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcion_declaracion
	return p
}

func InitEmptyFuncion_declaracionContext(p *Funcion_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcion_declaracion
}

func (*Funcion_declaracionContext) IsFuncion_declaracionContext() {}

func NewFuncion_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_declaracionContext {
	var p = new(Funcion_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcion_declaracion

	return p
}

func (s *Funcion_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_declaracionContext) CopyAll(ctx *Funcion_declaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Funcion_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecFunc2Context struct {
	Funcion_declaracionContext
	nomFunc     antlr.Token
	tipoRetorno ITipoContext
}

func NewDecFunc2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecFunc2Context {
	var p = new(DecFunc2Context)

	InitEmptyFuncion_declaracionContext(&p.Funcion_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Funcion_declaracionContext))

	return p
}

func (s *DecFunc2Context) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *DecFunc2Context) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *DecFunc2Context) GetTipoRetorno() ITipoContext { return s.tipoRetorno }

func (s *DecFunc2Context) SetTipoRetorno(v ITipoContext) { s.tipoRetorno = v }

func (s *DecFunc2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecFunc2Context) RESFUNC() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFUNC, 0)
}

func (s *DecFunc2Context) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *DecFunc2Context) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *DecFunc2Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *DecFunc2Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DecFunc2Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DecFunc2Context) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *DecFunc2Context) AllCORABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORABRE)
}

func (s *DecFunc2Context) CORABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, i)
}

func (s *DecFunc2Context) AllCORCIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORCIERRA)
}

func (s *DecFunc2Context) CORCIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, i)
}

func (s *DecFunc2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDecFunc2(s)
	}
}

func (s *DecFunc2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDecFunc2(s)
	}
}

func (s *DecFunc2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDecFunc2(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecFunc1Context struct {
	Funcion_declaracionContext
	nomFunc     antlr.Token
	tipoRetorno ITipoContext
}

func NewDecFunc1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecFunc1Context {
	var p = new(DecFunc1Context)

	InitEmptyFuncion_declaracionContext(&p.Funcion_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Funcion_declaracionContext))

	return p
}

func (s *DecFunc1Context) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *DecFunc1Context) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *DecFunc1Context) GetTipoRetorno() ITipoContext { return s.tipoRetorno }

func (s *DecFunc1Context) SetTipoRetorno(v ITipoContext) { s.tipoRetorno = v }

func (s *DecFunc1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecFunc1Context) RESFUNC() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFUNC, 0)
}

func (s *DecFunc1Context) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *DecFunc1Context) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *DecFunc1Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *DecFunc1Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DecFunc1Context) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *DecFunc1Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DecFunc1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDecFunc1(s)
	}
}

func (s *DecFunc1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDecFunc1(s)
	}
}

func (s *DecFunc1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDecFunc1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Funcion_declaracion() (localctx IFuncion_declaracionContext) {
	localctx = NewFuncion_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_funcion_declaracion)
	var _la int

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDecFunc1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(gramaticaParserRESFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*DecFunc1Context).nomFunc = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserIDENTIFICADOR {
			{
				p.SetState(144)
				p.Lista_parametros()
			}

		}
		{
			p.SetState(147)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719484672) != 0 {
			{
				p.SetState(148)

				var _x = p.Tipo()

				localctx.(*DecFunc1Context).tipoRetorno = _x
			}

		}
		{
			p.SetState(151)
			p.Bloque()
		}

	case 2:
		localctx = NewDecFunc2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(gramaticaParserRESFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*DecFunc2Context).nomFunc = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserIDENTIFICADOR {
			{
				p.SetState(155)
				p.Lista_parametros()
			}

		}
		{
			p.SetState(158)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == gramaticaParserCORABRE {
			{
				p.SetState(159)
				p.Match(gramaticaParserCORABRE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(160)
				p.Match(gramaticaParserCORCIERRA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)

			var _x = p.Tipo()

			localctx.(*DecFunc2Context).tipoRetorno = _x
		}
		{
			p.SetState(167)
			p.Bloque()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametrosContext is an interface to support dynamic dispatch.
type ILista_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametro() []IParametroContext
	Parametro(i int) IParametroContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsLista_parametrosContext differentiates from other interfaces.
	IsLista_parametrosContext()
}

type Lista_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametrosContext() *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_parametros
	return p
}

func InitEmptyLista_parametrosContext(p *Lista_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_parametros
}

func (*Lista_parametrosContext) IsLista_parametrosContext() {}

func NewLista_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_parametros

	return p
}

func (s *Lista_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametrosContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Lista_parametrosContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Lista_parametrosContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *Lista_parametrosContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *Lista_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_parametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_parametros(s)
	}
}

func (s *Lista_parametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_parametros(s)
	}
}

func (s *Lista_parametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_parametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_parametros() (localctx ILista_parametrosContext) {
	localctx = NewLista_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_lista_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Parametro()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(172)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Parametro()
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdParametro returns the idParametro token.
	GetIdParametro() antlr.Token

	// SetIdParametro sets the idParametro token.
	SetIdParametro(antlr.Token)

	// Getter signatures
	Tipo() ITipoContext
	IDENTIFICADOR() antlr.TerminalNode
	SIGNO_MULTI() antlr.TerminalNode
	CORABRE() antlr.TerminalNode
	CORCIERRA() antlr.TerminalNode

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	idParametro antlr.Token
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) GetIdParametro() antlr.Token { return s.idParametro }

func (s *ParametroContext) SetIdParametro(v antlr.Token) { s.idParametro = v }

func (s *ParametroContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ParametroContext) SIGNO_MULTI() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MULTI, 0)
}

func (s *ParametroContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *ParametroContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (s *ParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gramaticaParserRULE_parametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*ParametroContext).idParametro = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case gramaticaParserSIGNO_MULTI:
		{
			p.SetState(180)
			p.Match(gramaticaParserSIGNO_MULTI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case gramaticaParserCORABRE:
		{
			p.SetState(181)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case gramaticaParserRESINT, gramaticaParserRESFLOAT, gramaticaParserRESSTRING, gramaticaParserRESBOOL, gramaticaParserRESRUNE, gramaticaParserIDENTIFICADOR:

	default:
	}
	{
		p.SetState(185)
		p.Tipo()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_sentenciasContext is an interface to support dynamic dispatch.
type ILista_sentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSentencia() []ISentenciaContext
	Sentencia(i int) ISentenciaContext

	// IsLista_sentenciasContext differentiates from other interfaces.
	IsLista_sentenciasContext()
}

type Lista_sentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_sentenciasContext() *Lista_sentenciasContext {
	var p = new(Lista_sentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_sentencias
	return p
}

func InitEmptyLista_sentenciasContext(p *Lista_sentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_sentencias
}

func (*Lista_sentenciasContext) IsLista_sentenciasContext() {}

func NewLista_sentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_sentenciasContext {
	var p = new(Lista_sentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_sentencias

	return p
}

func (s *Lista_sentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_sentenciasContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *Lista_sentenciasContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaContext)
}

func (s *Lista_sentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_sentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_sentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_sentencias(s)
	}
}

func (s *Lista_sentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_sentencias(s)
	}
}

func (s *Lista_sentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_sentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_sentencias() (localctx ILista_sentenciasContext) {
	localctx = NewLista_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_lista_sentencias)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Sentencia()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(188)
				p.Sentencia()
			}

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bloque() IBloqueContext
	Declaracion_variable() IDeclaracion_variableContext
	Asignacion() IAsignacionContext
	Declaracion_slice() IDeclaracion_sliceContext
	Asignacion_acceso_slice() IAsignacion_acceso_sliceContext
	SentenciaIf() ISentenciaIfContext
	SentenciaSwitch() ISentenciaSwitchContext
	Sentencia_transferencia() ISentencia_transferenciaContext
	SentenciaPrintln() ISentenciaPrintlnContext
	Creacion_instancia_struct() ICreacion_instancia_structContext
	Asignacion_acceso_atributo() IAsignacion_acceso_atributoContext
	Llamada_funcion_struct() ILlamada_funcion_structContext
	Llamadafuncion() ILlamadafuncionContext
	SentenciaFor() ISentenciaForContext
	Declaracion_matriz() IDeclaracion_matrizContext
	Asignacion_acceso_matriz() IAsignacion_acceso_matrizContext
	SentenciaReturn() ISentenciaReturnContext
	SentenciaBreak() ISentenciaBreakContext

	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentenciaContext) Declaracion_variable() IDeclaracion_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_variableContext)
}

func (s *SentenciaContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *SentenciaContext) Declaracion_slice() IDeclaracion_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_sliceContext)
}

func (s *SentenciaContext) Asignacion_acceso_slice() IAsignacion_acceso_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_acceso_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_acceso_sliceContext)
}

func (s *SentenciaContext) SentenciaIf() ISentenciaIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaIfContext)
}

func (s *SentenciaContext) SentenciaSwitch() ISentenciaSwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaSwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaSwitchContext)
}

func (s *SentenciaContext) Sentencia_transferencia() ISentencia_transferenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencia_transferenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencia_transferenciaContext)
}

func (s *SentenciaContext) SentenciaPrintln() ISentenciaPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaPrintlnContext)
}

func (s *SentenciaContext) Creacion_instancia_struct() ICreacion_instancia_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreacion_instancia_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreacion_instancia_structContext)
}

func (s *SentenciaContext) Asignacion_acceso_atributo() IAsignacion_acceso_atributoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_acceso_atributoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_acceso_atributoContext)
}

func (s *SentenciaContext) Llamada_funcion_struct() ILlamada_funcion_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_funcion_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcion_structContext)
}

func (s *SentenciaContext) Llamadafuncion() ILlamadafuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadafuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadafuncionContext)
}

func (s *SentenciaContext) SentenciaFor() ISentenciaForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaForContext)
}

func (s *SentenciaContext) Declaracion_matriz() IDeclaracion_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_matrizContext)
}

func (s *SentenciaContext) Asignacion_acceso_matriz() IAsignacion_acceso_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_acceso_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_acceso_matrizContext)
}

func (s *SentenciaContext) SentenciaReturn() ISentenciaReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaReturnContext)
}

func (s *SentenciaContext) SentenciaBreak() ISentenciaBreakContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaBreakContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaBreakContext)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentencia(s)
	}
}

func (s *SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentencia(s)
	}
}

func (s *SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gramaticaParserRULE_sentencia)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Bloque()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Declaracion_variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.Asignacion()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)
			p.Declaracion_slice()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(198)
			p.Asignacion_acceso_slice()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(199)
			p.SentenciaIf()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(200)
			p.SentenciaSwitch()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(201)
			p.Sentencia_transferencia()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(202)
			p.SentenciaPrintln()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(203)
			p.Creacion_instancia_struct()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(204)
			p.Asignacion_acceso_atributo()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(205)
			p.Llamada_funcion_struct()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(206)
			p.Llamadafuncion()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(207)
			p.SentenciaFor()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(208)
			p.Declaracion_matriz()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(209)
			p.Asignacion_acceso_matriz()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(210)
			p.SentenciaReturn()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(211)
			p.SentenciaBreak()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaBreakContext is an interface to support dynamic dispatch.
type ISentenciaBreakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESBREAK() antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsSentenciaBreakContext differentiates from other interfaces.
	IsSentenciaBreakContext()
}

type SentenciaBreakContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaBreakContext() *SentenciaBreakContext {
	var p = new(SentenciaBreakContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaBreak
	return p
}

func InitEmptySentenciaBreakContext(p *SentenciaBreakContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaBreak
}

func (*SentenciaBreakContext) IsSentenciaBreakContext() {}

func NewSentenciaBreakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaBreakContext {
	var p = new(SentenciaBreakContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaBreak

	return p
}

func (s *SentenciaBreakContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaBreakContext) RESBREAK() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESBREAK, 0)
}

func (s *SentenciaBreakContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentenciaBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaBreakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaBreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentenciaBreak(s)
	}
}

func (s *SentenciaBreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentenciaBreak(s)
	}
}

func (s *SentenciaBreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentenciaBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaBreak() (localctx ISentenciaBreakContext) {
	localctx = NewSentenciaBreakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gramaticaParserRULE_sentenciaBreak)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(gramaticaParserRESBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(215)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaReturnContext is an interface to support dynamic dispatch.
type ISentenciaReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESRETURN() antlr.TerminalNode
	Expresion() IExpresionContext
	PUNTO_COMA() antlr.TerminalNode

	// IsSentenciaReturnContext differentiates from other interfaces.
	IsSentenciaReturnContext()
}

type SentenciaReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaReturnContext() *SentenciaReturnContext {
	var p = new(SentenciaReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaReturn
	return p
}

func InitEmptySentenciaReturnContext(p *SentenciaReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaReturn
}

func (*SentenciaReturnContext) IsSentenciaReturnContext() {}

func NewSentenciaReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaReturnContext {
	var p = new(SentenciaReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaReturn

	return p
}

func (s *SentenciaReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaReturnContext) RESRETURN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESRETURN, 0)
}

func (s *SentenciaReturnContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentenciaReturnContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentenciaReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentenciaReturn(s)
	}
}

func (s *SentenciaReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentenciaReturn(s)
	}
}

func (s *SentenciaReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentenciaReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaReturn() (localctx ISentenciaReturnContext) {
	localctx = NewSentenciaReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gramaticaParserRULE_sentenciaReturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(gramaticaParserRESRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.expresion(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(220)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadafuncionContext is an interface to support dynamic dispatch.
type ILlamadafuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomFunc returns the nomFunc token.
	GetNomFunc() antlr.Token

	// SetNomFunc sets the nomFunc token.
	SetNomFunc(antlr.Token)

	// Getter signatures
	PARABRE() antlr.TerminalNode
	PARCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	Lista_parametros_func() ILista_parametros_funcContext

	// IsLlamadafuncionContext differentiates from other interfaces.
	IsLlamadafuncionContext()
}

type LlamadafuncionContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	nomFunc antlr.Token
}

func NewEmptyLlamadafuncionContext() *LlamadafuncionContext {
	var p = new(LlamadafuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadafuncion
	return p
}

func InitEmptyLlamadafuncionContext(p *LlamadafuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadafuncion
}

func (*LlamadafuncionContext) IsLlamadafuncionContext() {}

func NewLlamadafuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadafuncionContext {
	var p = new(LlamadafuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadafuncion

	return p
}

func (s *LlamadafuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadafuncionContext) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *LlamadafuncionContext) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *LlamadafuncionContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadafuncionContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadafuncionContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadafuncionContext) Lista_parametros_func() ILista_parametros_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametros_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametros_funcContext)
}

func (s *LlamadafuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadafuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadafuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadafuncion(s)
	}
}

func (s *LlamadafuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadafuncion(s)
	}
}

func (s *LlamadafuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadafuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Llamadafuncion() (localctx ILlamadafuncionContext) {
	localctx = NewLlamadafuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gramaticaParserRULE_llamadafuncion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*LlamadafuncionContext).nomFunc = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4900068058465779718) != 0 {
		{
			p.SetState(225)
			p.Lista_parametros_func()
		}

	}
	{
		p.SetState(228)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentencia_transferenciaContext is an interface to support dynamic dispatch.
type ISentencia_transferenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentencia_transferenciaContext differentiates from other interfaces.
	IsSentencia_transferenciaContext()
}

type Sentencia_transferenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencia_transferenciaContext() *Sentencia_transferenciaContext {
	var p = new(Sentencia_transferenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_transferencia
	return p
}

func InitEmptySentencia_transferenciaContext(p *Sentencia_transferenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_transferencia
}

func (*Sentencia_transferenciaContext) IsSentencia_transferenciaContext() {}

func NewSentencia_transferenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencia_transferenciaContext {
	var p = new(Sentencia_transferenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentencia_transferencia

	return p
}

func (s *Sentencia_transferenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencia_transferenciaContext) CopyAll(ctx *Sentencia_transferenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Sentencia_transferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencia_transferenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SentContinueContext struct {
	Sentencia_transferenciaContext
}

func NewSentContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentContinueContext {
	var p = new(SentContinueContext)

	InitEmptySentencia_transferenciaContext(&p.Sentencia_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencia_transferenciaContext))

	return p
}

func (s *SentContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentContinueContext) RESCONTINUE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESCONTINUE, 0)
}

func (s *SentContinueContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentContinue(s)
	}
}

func (s *SentContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentContinue(s)
	}
}

func (s *SentContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentBreakContext struct {
	Sentencia_transferenciaContext
}

func NewSentBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentBreakContext {
	var p = new(SentBreakContext)

	InitEmptySentencia_transferenciaContext(&p.Sentencia_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencia_transferenciaContext))

	return p
}

func (s *SentBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentBreakContext) RESBREAK() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESBREAK, 0)
}

func (s *SentBreakContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentBreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentBreak(s)
	}
}

func (s *SentBreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentBreak(s)
	}
}

func (s *SentBreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentReturnContext struct {
	Sentencia_transferenciaContext
}

func NewSentReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentReturnContext {
	var p = new(SentReturnContext)

	InitEmptySentencia_transferenciaContext(&p.Sentencia_transferenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Sentencia_transferenciaContext))

	return p
}

func (s *SentReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentReturnContext) RESRETURN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESRETURN, 0)
}

func (s *SentReturnContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentReturnContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentReturn(s)
	}
}

func (s *SentReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentReturn(s)
	}
}

func (s *SentReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sentencia_transferencia() (localctx ISentencia_transferenciaContext) {
	localctx = NewSentencia_transferenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gramaticaParserRULE_sentencia_transferencia)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserRESBREAK:
		localctx = NewSentBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Match(gramaticaParserRESBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(231)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case gramaticaParserRESCONTINUE:
		localctx = NewSentContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(gramaticaParserRESCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(235)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case gramaticaParserRESRETURN:
		localctx = NewSentReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)
			p.Match(gramaticaParserRESRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(239)
				p.expresion(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(242)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaIfContext is an interface to support dynamic dispatch.
type ISentenciaIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaIfContext differentiates from other interfaces.
	IsSentenciaIfContext()
}

type SentenciaIfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaIfContext() *SentenciaIfContext {
	var p = new(SentenciaIfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaIf
	return p
}

func InitEmptySentenciaIfContext(p *SentenciaIfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaIf
}

func (*SentenciaIfContext) IsSentenciaIfContext() {}

func NewSentenciaIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaIfContext {
	var p = new(SentenciaIfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaIf

	return p
}

func (s *SentenciaIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaIfContext) CopyAll(ctx *SentenciaIfContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SentIf3Context struct {
	SentenciaIfContext
}

func NewSentIf3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentIf3Context {
	var p = new(SentIf3Context)

	InitEmptySentenciaIfContext(&p.SentenciaIfContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaIfContext))

	return p
}

func (s *SentIf3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentIf3Context) RESIF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESIF, 0)
}

func (s *SentIf3Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentIf3Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentIf3Context) RESELSE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESELSE, 0)
}

func (s *SentIf3Context) SentenciaIf() ISentenciaIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaIfContext)
}

func (s *SentIf3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentIf3(s)
	}
}

func (s *SentIf3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentIf3(s)
	}
}

func (s *SentIf3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentIf3(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentIf1Context struct {
	SentenciaIfContext
}

func NewSentIf1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentIf1Context {
	var p = new(SentIf1Context)

	InitEmptySentenciaIfContext(&p.SentenciaIfContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaIfContext))

	return p
}

func (s *SentIf1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentIf1Context) RESIF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESIF, 0)
}

func (s *SentIf1Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentIf1Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentIf1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentIf1(s)
	}
}

func (s *SentIf1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentIf1(s)
	}
}

func (s *SentIf1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentIf1(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentIf2Context struct {
	SentenciaIfContext
}

func NewSentIf2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentIf2Context {
	var p = new(SentIf2Context)

	InitEmptySentenciaIfContext(&p.SentenciaIfContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaIfContext))

	return p
}

func (s *SentIf2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentIf2Context) RESIF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESIF, 0)
}

func (s *SentIf2Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentIf2Context) AllBloque() []IBloqueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueContext); ok {
			len++
		}
	}

	tst := make([]IBloqueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueContext); ok {
			tst[i] = t.(IBloqueContext)
			i++
		}
	}

	return tst
}

func (s *SentIf2Context) Bloque(i int) IBloqueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentIf2Context) RESELSE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESELSE, 0)
}

func (s *SentIf2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentIf2(s)
	}
}

func (s *SentIf2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentIf2(s)
	}
}

func (s *SentIf2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentIf2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaIf() (localctx ISentenciaIfContext) {
	localctx = NewSentenciaIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gramaticaParserRULE_sentenciaIf)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSentIf1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(gramaticaParserRESIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.expresion(0)
		}
		{
			p.SetState(249)
			p.Bloque()
		}

	case 2:
		localctx = NewSentIf2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(gramaticaParserRESIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.expresion(0)
		}
		{
			p.SetState(253)
			p.Bloque()
		}
		{
			p.SetState(254)
			p.Match(gramaticaParserRESELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Bloque()
		}

	case 3:
		localctx = NewSentIf3Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(257)
			p.Match(gramaticaParserRESIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.expresion(0)
		}
		{
			p.SetState(259)
			p.Bloque()
		}
		{
			p.SetState(260)
			p.Match(gramaticaParserRESELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.SentenciaIf()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaSwitchContext is an interface to support dynamic dispatch.
type ISentenciaSwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESSWITCH() antlr.TerminalNode
	Expresion() IExpresionContext
	LLAVEABRE() antlr.TerminalNode
	LLAVECIERRA() antlr.TerminalNode
	AllBloqueCase() []IBloqueCaseContext
	BloqueCase(i int) IBloqueCaseContext
	BloqueDefault() IBloqueDefaultContext

	// IsSentenciaSwitchContext differentiates from other interfaces.
	IsSentenciaSwitchContext()
}

type SentenciaSwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaSwitchContext() *SentenciaSwitchContext {
	var p = new(SentenciaSwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaSwitch
	return p
}

func InitEmptySentenciaSwitchContext(p *SentenciaSwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaSwitch
}

func (*SentenciaSwitchContext) IsSentenciaSwitchContext() {}

func NewSentenciaSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaSwitchContext {
	var p = new(SentenciaSwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaSwitch

	return p
}

func (s *SentenciaSwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaSwitchContext) RESSWITCH() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESSWITCH, 0)
}

func (s *SentenciaSwitchContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentenciaSwitchContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *SentenciaSwitchContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *SentenciaSwitchContext) AllBloqueCase() []IBloqueCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueCaseContext); ok {
			len++
		}
	}

	tst := make([]IBloqueCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueCaseContext); ok {
			tst[i] = t.(IBloqueCaseContext)
			i++
		}
	}

	return tst
}

func (s *SentenciaSwitchContext) BloqueCase(i int) IBloqueCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueCaseContext)
}

func (s *SentenciaSwitchContext) BloqueDefault() IBloqueDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueDefaultContext)
}

func (s *SentenciaSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaSwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentenciaSwitch(s)
	}
}

func (s *SentenciaSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentenciaSwitch(s)
	}
}

func (s *SentenciaSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentenciaSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaSwitch() (localctx ISentenciaSwitchContext) {
	localctx = NewSentenciaSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gramaticaParserRULE_sentenciaSwitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(gramaticaParserRESSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.expresion(0)
	}
	{
		p.SetState(267)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gramaticaParserRESCASE {
		{
			p.SetState(268)
			p.BloqueCase()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserRESDEFAULT {
		{
			p.SetState(273)
			p.BloqueDefault()
		}

	}
	{
		p.SetState(276)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueCaseContext is an interface to support dynamic dispatch.
type IBloqueCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESCASE() antlr.TerminalNode
	Expresion() IExpresionContext
	DOS_PUNTOS() antlr.TerminalNode
	AllLista_sentencias() []ILista_sentenciasContext
	Lista_sentencias(i int) ILista_sentenciasContext
	PUNTO_COMA() antlr.TerminalNode

	// IsBloqueCaseContext differentiates from other interfaces.
	IsBloqueCaseContext()
}

type BloqueCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueCaseContext() *BloqueCaseContext {
	var p = new(BloqueCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueCase
	return p
}

func InitEmptyBloqueCaseContext(p *BloqueCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueCase
}

func (*BloqueCaseContext) IsBloqueCaseContext() {}

func NewBloqueCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueCaseContext {
	var p = new(BloqueCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloqueCase

	return p
}

func (s *BloqueCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueCaseContext) RESCASE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESCASE, 0)
}

func (s *BloqueCaseContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *BloqueCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOS_PUNTOS, 0)
}

func (s *BloqueCaseContext) AllLista_sentencias() []ILista_sentenciasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_sentenciasContext); ok {
			len++
		}
	}

	tst := make([]ILista_sentenciasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_sentenciasContext); ok {
			tst[i] = t.(ILista_sentenciasContext)
			i++
		}
	}

	return tst
}

func (s *BloqueCaseContext) Lista_sentencias(i int) ILista_sentenciasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_sentenciasContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_sentenciasContext)
}

func (s *BloqueCaseContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *BloqueCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloqueCase(s)
	}
}

func (s *BloqueCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloqueCase(s)
	}
}

func (s *BloqueCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBloqueCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) BloqueCase() (localctx IBloqueCaseContext) {
	localctx = NewBloqueCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gramaticaParserRULE_bloqueCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(gramaticaParserRESCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.expresion(0)
	}
	{
		p.SetState(280)
		p.Match(gramaticaParserDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35258353320064) != 0 {
		{
			p.SetState(281)
			p.Lista_sentencias()
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserPUNTO_COMA {
		{
			p.SetState(287)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueDefaultContext is an interface to support dynamic dispatch.
type IBloqueDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESDEFAULT() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	Lista_sentencias() ILista_sentenciasContext

	// IsBloqueDefaultContext differentiates from other interfaces.
	IsBloqueDefaultContext()
}

type BloqueDefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueDefaultContext() *BloqueDefaultContext {
	var p = new(BloqueDefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueDefault
	return p
}

func InitEmptyBloqueDefaultContext(p *BloqueDefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueDefault
}

func (*BloqueDefaultContext) IsBloqueDefaultContext() {}

func NewBloqueDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueDefaultContext {
	var p = new(BloqueDefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloqueDefault

	return p
}

func (s *BloqueDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueDefaultContext) RESDEFAULT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESDEFAULT, 0)
}

func (s *BloqueDefaultContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOS_PUNTOS, 0)
}

func (s *BloqueDefaultContext) Lista_sentencias() ILista_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_sentenciasContext)
}

func (s *BloqueDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloqueDefault(s)
	}
}

func (s *BloqueDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloqueDefault(s)
	}
}

func (s *BloqueDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBloqueDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) BloqueDefault() (localctx IBloqueDefaultContext) {
	localctx = NewBloqueDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gramaticaParserRULE_bloqueDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(gramaticaParserRESDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(gramaticaParserDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35258353320064) != 0 {
		{
			p.SetState(292)
			p.Lista_sentencias()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaPrintlnContext is an interface to support dynamic dispatch.
type ISentenciaPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESPRINTLN() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	PARCIERRA() antlr.TerminalNode
	Lista_expresion() ILista_expresionContext
	PUNTO_COMA() antlr.TerminalNode

	// IsSentenciaPrintlnContext differentiates from other interfaces.
	IsSentenciaPrintlnContext()
}

type SentenciaPrintlnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaPrintlnContext() *SentenciaPrintlnContext {
	var p = new(SentenciaPrintlnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaPrintln
	return p
}

func InitEmptySentenciaPrintlnContext(p *SentenciaPrintlnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaPrintln
}

func (*SentenciaPrintlnContext) IsSentenciaPrintlnContext() {}

func NewSentenciaPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaPrintlnContext {
	var p = new(SentenciaPrintlnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaPrintln

	return p
}

func (s *SentenciaPrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaPrintlnContext) RESPRINTLN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESPRINTLN, 0)
}

func (s *SentenciaPrintlnContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *SentenciaPrintlnContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *SentenciaPrintlnContext) Lista_expresion() ILista_expresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionContext)
}

func (s *SentenciaPrintlnContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *SentenciaPrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaPrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaPrintlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentenciaPrintln(s)
	}
}

func (s *SentenciaPrintlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentenciaPrintln(s)
	}
}

func (s *SentenciaPrintlnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentenciaPrintln(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaPrintln() (localctx ISentenciaPrintlnContext) {
	localctx = NewSentenciaPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gramaticaParserRULE_sentenciaPrintln)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(gramaticaParserRESPRINTLN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4900068058465779718) != 0 {
		{
			p.SetState(297)
			p.Lista_expresion()
		}

	}
	{
		p.SetState(300)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(301)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsigSumRestContext struct {
	AsignacionContext
	id antlr.Token
	op antlr.Token
}

func NewAsigSumRestContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsigSumRestContext {
	var p = new(AsigSumRestContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsigSumRestContext) GetId() antlr.Token { return s.id }

func (s *AsigSumRestContext) GetOp() antlr.Token { return s.op }

func (s *AsigSumRestContext) SetId(v antlr.Token) { s.id = v }

func (s *AsigSumRestContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsigSumRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigSumRestContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigSumRestContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsigSumRestContext) SIGNO_ASIGNACION_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION_INCREMENTO, 0)
}

func (s *AsigSumRestContext) SIGNO_ASIGNACION_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION_DECREMENTO, 0)
}

func (s *AsigSumRestContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *AsigSumRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigSumRest(s)
	}
}

func (s *AsigSumRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigSumRest(s)
	}
}

func (s *AsigSumRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigSumRest(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_acceso_slice1Context struct {
	AsignacionContext
}

func NewAsignacion_acceso_slice1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_acceso_slice1Context {
	var p = new(Asignacion_acceso_slice1Context)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_acceso_slice1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_acceso_slice1Context) Asignacion_acceso_slice() IAsignacion_acceso_sliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_acceso_sliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_acceso_sliceContext)
}

func (s *Asignacion_acceso_slice1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacion_acceso_slice1(s)
	}
}

func (s *Asignacion_acceso_slice1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacion_acceso_slice1(s)
	}
}

func (s *Asignacion_acceso_slice1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacion_acceso_slice1(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsigSimpleContext struct {
	AsignacionContext
	id antlr.Token
}

func NewAsigSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsigSimpleContext {
	var p = new(AsigSimpleContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsigSimpleContext) GetId() antlr.Token { return s.id }

func (s *AsigSimpleContext) SetId(v antlr.Token) { s.id = v }

func (s *AsigSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigSimpleContext) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *AsigSimpleContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigSimpleContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsigSimpleContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *AsigSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigSimple(s)
	}
}

func (s *AsigSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigSimple(s)
	}
}

func (s *AsigSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gramaticaParserRULE_asignacion)
	var _la int

	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsigSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*AsigSimpleContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(gramaticaParserSIGNO_ASIGNACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.expresion(0)
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(307)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewAsigSumRestContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*AsigSumRestContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsigSumRestContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == gramaticaParserSIGNO_ASIGNACION_INCREMENTO || _la == gramaticaParserSIGNO_ASIGNACION_DECREMENTO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsigSumRestContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(312)
			p.expresion(0)
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(313)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		localctx = NewAsignacion_acceso_slice1Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.Asignacion_acceso_slice()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_sliceContext is an interface to support dynamic dispatch.
type IDeclaracion_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_sliceContext differentiates from other interfaces.
	IsDeclaracion_sliceContext()
}

type Declaracion_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_sliceContext() *Declaracion_sliceContext {
	var p = new(Declaracion_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_slice
	return p
}

func InitEmptyDeclaracion_sliceContext(p *Declaracion_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_slice
}

func (*Declaracion_sliceContext) IsDeclaracion_sliceContext() {}

func NewDeclaracion_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_sliceContext {
	var p = new(Declaracion_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracion_slice

	return p
}

func (s *Declaracion_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_sliceContext) CopyAll(ctx *Declaracion_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecSlice2Context struct {
	Declaracion_sliceContext
	nomSlice antlr.Token
}

func NewDecSlice2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecSlice2Context {
	var p = new(DecSlice2Context)

	InitEmptyDeclaracion_sliceContext(&p.Declaracion_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_sliceContext))

	return p
}

func (s *DecSlice2Context) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *DecSlice2Context) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *DecSlice2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecSlice2Context) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *DecSlice2Context) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *DecSlice2Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DecSlice2Context) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *DecSlice2Context) Lista_expresion() ILista_expresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionContext)
}

func (s *DecSlice2Context) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *DecSlice2Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DecSlice2Context) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *DecSlice2Context) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *DecSlice2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDecSlice2(s)
	}
}

func (s *DecSlice2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDecSlice2(s)
	}
}

func (s *DecSlice2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDecSlice2(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecSlice1Context struct {
	Declaracion_sliceContext
	nomSlice antlr.Token
}

func NewDecSlice1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecSlice1Context {
	var p = new(DecSlice1Context)

	InitEmptyDeclaracion_sliceContext(&p.Declaracion_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_sliceContext))

	return p
}

func (s *DecSlice1Context) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *DecSlice1Context) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *DecSlice1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecSlice1Context) RESMUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESMUT, 0)
}

func (s *DecSlice1Context) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *DecSlice1Context) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *DecSlice1Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DecSlice1Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DecSlice1Context) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *DecSlice1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDecSlice1(s)
	}
}

func (s *DecSlice1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDecSlice1(s)
	}
}

func (s *DecSlice1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDecSlice1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Declaracion_slice() (localctx IDeclaracion_sliceContext) {
	localctx = NewDeclaracion_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gramaticaParserRULE_declaracion_slice)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserRESMUT:
		localctx = NewDecSlice1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(gramaticaParserRESMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*DecSlice1Context).nomSlice = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Tipo()
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(324)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case gramaticaParserIDENTIFICADOR:
		localctx = NewDecSlice2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*DecSlice2Context).nomSlice = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(328)
			p.Match(gramaticaParserSIGNO_ASIGNACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(329)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Tipo()
		}
		{
			p.SetState(332)
			p.Match(gramaticaParserLLAVEABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Lista_expresion()
		}
		{
			p.SetState(334)
			p.Match(gramaticaParserLLAVECIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(335)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacion_acceso_sliceContext is an interface to support dynamic dispatch.
type IAsignacion_acceso_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacion_acceso_sliceContext differentiates from other interfaces.
	IsAsignacion_acceso_sliceContext()
}

type Asignacion_acceso_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacion_acceso_sliceContext() *Asignacion_acceso_sliceContext {
	var p = new(Asignacion_acceso_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_slice
	return p
}

func InitEmptyAsignacion_acceso_sliceContext(p *Asignacion_acceso_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_slice
}

func (*Asignacion_acceso_sliceContext) IsAsignacion_acceso_sliceContext() {}

func NewAsignacion_acceso_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_acceso_sliceContext {
	var p = new(Asignacion_acceso_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_slice

	return p
}

func (s *Asignacion_acceso_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_acceso_sliceContext) CopyAll(ctx *Asignacion_acceso_sliceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asignacion_acceso_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_acceso_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsigacionSliceContext struct {
	Asignacion_acceso_sliceContext
	index IExpresionContext
}

func NewAsigacionSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsigacionSliceContext {
	var p = new(AsigacionSliceContext)

	InitEmptyAsignacion_acceso_sliceContext(&p.Asignacion_acceso_sliceContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignacion_acceso_sliceContext))

	return p
}

func (s *AsigacionSliceContext) GetIndex() IExpresionContext { return s.index }

func (s *AsigacionSliceContext) SetIndex(v IExpresionContext) { s.index = v }

func (s *AsigacionSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigacionSliceContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsigacionSliceContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *AsigacionSliceContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *AsigacionSliceContext) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *AsigacionSliceContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *AsigacionSliceContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigacionSliceContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *AsigacionSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigacionSlice(s)
	}
}

func (s *AsigacionSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigacionSlice(s)
	}
}

func (s *AsigacionSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigacionSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Asignacion_acceso_slice() (localctx IAsignacion_acceso_sliceContext) {
	localctx = NewAsignacion_acceso_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gramaticaParserRULE_asignacion_acceso_slice)
	localctx = NewAsigacionSliceContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)

		var _x = p.expresion(0)

		localctx.(*AsigacionSliceContext).index = _x
	}
	{
		p.SetState(343)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(gramaticaParserSIGNO_ASIGNACION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.expresion(0)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(346)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_variableContext is an interface to support dynamic dispatch.
type IDeclaracion_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_variableContext differentiates from other interfaces.
	IsDeclaracion_variableContext()
}

type Declaracion_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_variableContext() *Declaracion_variableContext {
	var p = new(Declaracion_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_variable
	return p
}

func InitEmptyDeclaracion_variableContext(p *Declaracion_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_variable
}

func (*Declaracion_variableContext) IsDeclaracion_variableContext() {}

func NewDeclaracion_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_variableContext {
	var p = new(Declaracion_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracion_variable

	return p
}

func (s *Declaracion_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_variableContext) CopyAll(ctx *Declaracion_variableContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Dec2Context struct {
	Declaracion_variableContext
	id antlr.Token
}

func NewDec2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Dec2Context {
	var p = new(Dec2Context)

	InitEmptyDeclaracion_variableContext(&p.Declaracion_variableContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_variableContext))

	return p
}

func (s *Dec2Context) GetId() antlr.Token { return s.id }

func (s *Dec2Context) SetId(v antlr.Token) { s.id = v }

func (s *Dec2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec2Context) ASIGNACION_INFERIDA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserASIGNACION_INFERIDA, 0)
}

func (s *Dec2Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Dec2Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Dec2Context) RESMUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESMUT, 0)
}

func (s *Dec2Context) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Dec2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDec2(s)
	}
}

func (s *Dec2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDec2(s)
	}
}

func (s *Dec2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDec2(s)

	default:
		return t.VisitChildren(s)
	}
}

type Dec3Context struct {
	Declaracion_variableContext
	id antlr.Token
}

func NewDec3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Dec3Context {
	var p = new(Dec3Context)

	InitEmptyDeclaracion_variableContext(&p.Declaracion_variableContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_variableContext))

	return p
}

func (s *Dec3Context) GetId() antlr.Token { return s.id }

func (s *Dec3Context) SetId(v antlr.Token) { s.id = v }

func (s *Dec3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec3Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Dec3Context) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *Dec3Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Dec3Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Dec3Context) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Dec3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDec3(s)
	}
}

func (s *Dec3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDec3(s)
	}
}

func (s *Dec3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDec3(s)

	default:
		return t.VisitChildren(s)
	}
}

type Dec1Context struct {
	Declaracion_variableContext
	id  antlr.Token
	tip ITipoContext
	val IExpresionContext
}

func NewDec1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Dec1Context {
	var p = new(Dec1Context)

	InitEmptyDeclaracion_variableContext(&p.Declaracion_variableContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_variableContext))

	return p
}

func (s *Dec1Context) GetId() antlr.Token { return s.id }

func (s *Dec1Context) SetId(v antlr.Token) { s.id = v }

func (s *Dec1Context) GetTip() ITipoContext { return s.tip }

func (s *Dec1Context) GetVal() IExpresionContext { return s.val }

func (s *Dec1Context) SetTip(v ITipoContext) { s.tip = v }

func (s *Dec1Context) SetVal(v IExpresionContext) { s.val = v }

func (s *Dec1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec1Context) RESMUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESMUT, 0)
}

func (s *Dec1Context) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Dec1Context) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *Dec1Context) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Dec1Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Dec1Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Dec1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDec1(s)
	}
}

func (s *Dec1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDec1(s)
	}
}

func (s *Dec1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDec1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Declaracion_variable() (localctx IDeclaracion_variableContext) {
	localctx = NewDeclaracion_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gramaticaParserRULE_declaracion_variable)
	var _la int

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDec1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(gramaticaParserRESMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*Dec1Context).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(351)

				var _x = p.Tipo()

				localctx.(*Dec1Context).tip = _x
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserSIGNO_ASIGNACION {
			{
				p.SetState(354)
				p.Match(gramaticaParserSIGNO_ASIGNACION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(355)

				var _x = p.expresion(0)

				localctx.(*Dec1Context).val = _x
			}

		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(358)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewDec2Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserRESMUT {
			{
				p.SetState(361)
				p.Match(gramaticaParserRESMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(364)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*Dec2Context).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(gramaticaParserASIGNACION_INFERIDA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.expresion(0)
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(367)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		localctx = NewDec3Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(370)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*Dec3Context).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Tipo()
		}
		{
			p.SetState(372)
			p.Match(gramaticaParserSIGNO_ASIGNACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.expresion(0)
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(374)
				p.Match(gramaticaParserPUNTO_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprOpAritmeticaContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewExprOpAritmeticaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprOpAritmeticaContext {
	var p = new(ExprOpAritmeticaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprOpAritmeticaContext) GetOp() antlr.Token { return s.op }

func (s *ExprOpAritmeticaContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprOpAritmeticaContext) GetLeft() IExpresionContext { return s.left }

func (s *ExprOpAritmeticaContext) GetRight() IExpresionContext { return s.right }

func (s *ExprOpAritmeticaContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExprOpAritmeticaContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExprOpAritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOpAritmeticaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExprOpAritmeticaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprOpAritmeticaContext) SIGNO_MAS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MAS, 0)
}

func (s *ExprOpAritmeticaContext) SIGNO_MENOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MENOS, 0)
}

func (s *ExprOpAritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprOpAritmetica(s)
	}
}

func (s *ExprOpAritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprOpAritmetica(s)
	}
}

func (s *ExprOpAritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprOpAritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBoolFalseContext struct {
	ExpresionContext
	valor antlr.Token
}

func NewExprBoolFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBoolFalseContext {
	var p = new(ExprBoolFalseContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprBoolFalseContext) GetValor() antlr.Token { return s.valor }

func (s *ExprBoolFalseContext) SetValor(v antlr.Token) { s.valor = v }

func (s *ExprBoolFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBoolFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprBoolFalse(s)
	}
}

func (s *ExprBoolFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprBoolFalse(s)
	}
}

func (s *ExprBoolFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprBoolFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIdentificadorContext struct {
	ExpresionContext
}

func NewExprIdentificadorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIdentificadorContext {
	var p = new(ExprIdentificadorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprIdentificadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIdentificadorContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ExprIdentificadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprIdentificador(s)
	}
}

func (s *ExprIdentificadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprIdentificador(s)
	}
}

func (s *ExprIdentificadorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprIdentificador(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprRuneContext struct {
	ExpresionContext
}

func NewExprRuneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprRuneContext {
	var p = new(ExprRuneContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprRuneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRuneContext) RUNE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRUNE, 0)
}

func (s *ExprRuneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprRune(s)
	}
}

func (s *ExprRuneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprRune(s)
	}
}

func (s *ExprRuneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprRune(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprFloatContext struct {
	ExpresionContext
}

func NewExprFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFloatContext {
	var p = new(ExprFloatContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFLOAT, 0)
}

func (s *ExprFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprFloat(s)
	}
}

func (s *ExprFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprFloat(s)
	}
}

func (s *ExprFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNotContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExprNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNotContext {
	var p = new(ExprNotContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprNotContext) GetOp() antlr.Token { return s.op }

func (s *ExprNotContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNotContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprNotContext) SIGNO_NOT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_NOT, 0)
}

func (s *ExprNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprNot(s)
	}
}

func (s *ExprNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprNot(s)
	}
}

func (s *ExprNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNegUnariaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExprNegUnariaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNegUnariaContext {
	var p = new(ExprNegUnariaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprNegUnariaContext) GetOp() antlr.Token { return s.op }

func (s *ExprNegUnariaContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprNegUnariaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNegUnariaContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *ExprNegUnariaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprNegUnariaContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *ExprNegUnariaContext) SIGNO_MENOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MENOS, 0)
}

func (s *ExprNegUnariaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprNegUnaria(s)
	}
}

func (s *ExprNegUnariaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprNegUnaria(s)
	}
}

func (s *ExprNegUnariaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprNegUnaria(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNilContext struct {
	ExpresionContext
}

func NewExprNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNilContext {
	var p = new(ExprNilContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNilContext) RESNIL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESNIL, 0)
}

func (s *ExprNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprNil(s)
	}
}

func (s *ExprNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprNil(s)
	}
}

func (s *ExprNilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprMultDivModContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewExprMultDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMultDivModContext {
	var p = new(ExprMultDivModContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprMultDivModContext) GetOp() antlr.Token { return s.op }

func (s *ExprMultDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprMultDivModContext) GetLeft() IExpresionContext { return s.left }

func (s *ExprMultDivModContext) GetRight() IExpresionContext { return s.right }

func (s *ExprMultDivModContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExprMultDivModContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExprMultDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMultDivModContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExprMultDivModContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprMultDivModContext) SIGNO_MULTI() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MULTI, 0)
}

func (s *ExprMultDivModContext) SIGNO_DIV() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_DIV, 0)
}

func (s *ExprMultDivModContext) SIGNO_MODULO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MODULO, 0)
}

func (s *ExprMultDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprMultDivMod(s)
	}
}

func (s *ExprMultDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprMultDivMod(s)
	}
}

func (s *ExprMultDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprMultDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAccesoMatrizContext struct {
	ExpresionContext
}

func NewExprAccesoMatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAccesoMatrizContext {
	var p = new(ExprAccesoMatrizContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprAccesoMatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAccesoMatrizContext) AccesoMatriz() IAccesoMatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesoMatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesoMatrizContext)
}

func (s *ExprAccesoMatrizContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ExprAccesoMatrizContext) AllCORABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORABRE)
}

func (s *ExprAccesoMatrizContext) CORABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, i)
}

func (s *ExprAccesoMatrizContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExprAccesoMatrizContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprAccesoMatrizContext) AllCORCIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORCIERRA)
}

func (s *ExprAccesoMatrizContext) CORCIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, i)
}

func (s *ExprAccesoMatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprAccesoMatriz(s)
	}
}

func (s *ExprAccesoMatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprAccesoMatriz(s)
	}
}

func (s *ExprAccesoMatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprAccesoMatriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAccesoSliceContext struct {
	ExpresionContext
}

func NewExprAccesoSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAccesoSliceContext {
	var p = new(ExprAccesoSliceContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprAccesoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAccesoSliceContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ExprAccesoSliceContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *ExprAccesoSliceContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprAccesoSliceContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *ExprAccesoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprAccesoSlice(s)
	}
}

func (s *ExprAccesoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprAccesoSlice(s)
	}
}

func (s *ExprAccesoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprAccesoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprSliceLiteralContext struct {
	ExpresionContext
}

func NewExprSliceLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprSliceLiteralContext {
	var p = new(ExprSliceLiteralContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprSliceLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSliceLiteralContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *ExprSliceLiteralContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *ExprSliceLiteralContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ExprSliceLiteralContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *ExprSliceLiteralContext) Lista_expresion() ILista_expresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionContext)
}

func (s *ExprSliceLiteralContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *ExprSliceLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprSliceLiteral(s)
	}
}

func (s *ExprSliceLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprSliceLiteral(s)
	}
}

func (s *ExprSliceLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprSliceLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBoolTrueContext struct {
	ExpresionContext
	valor antlr.Token
}

func NewExprBoolTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBoolTrueContext {
	var p = new(ExprBoolTrueContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprBoolTrueContext) GetValor() antlr.Token { return s.valor }

func (s *ExprBoolTrueContext) SetValor(v antlr.Token) { s.valor = v }

func (s *ExprBoolTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBoolTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprBoolTrue(s)
	}
}

func (s *ExprBoolTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprBoolTrue(s)
	}
}

func (s *ExprBoolTrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprBoolTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLlamadaFuncionContext struct {
	ExpresionContext
}

func NewExprLlamadaFuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLlamadaFuncionContext {
	var p = new(ExprLlamadaFuncionContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprLlamadaFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLlamadaFuncionContext) Llamadafuncion() ILlamadafuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadafuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadafuncionContext)
}

func (s *ExprLlamadaFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprLlamadaFuncion(s)
	}
}

func (s *ExprLlamadaFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprLlamadaFuncion(s)
	}
}

func (s *ExprLlamadaFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprLlamadaFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLogicaContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewExprLogicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLogicaContext {
	var p = new(ExprLogicaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprLogicaContext) GetOp() antlr.Token { return s.op }

func (s *ExprLogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprLogicaContext) GetLeft() IExpresionContext { return s.left }

func (s *ExprLogicaContext) GetRight() IExpresionContext { return s.right }

func (s *ExprLogicaContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExprLogicaContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExprLogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLogicaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExprLogicaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprLogicaContext) SIGNO_AND() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_AND, 0)
}

func (s *ExprLogicaContext) SIGNO_OR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_OR, 0)
}

func (s *ExprLogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprLogica(s)
	}
}

func (s *ExprLogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprLogica(s)
	}
}

func (s *ExprLogicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprLogica(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIndexOfContext struct {
	ExpresionContext
}

func NewExprIndexOfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIndexOfContext {
	var p = new(ExprIndexOfContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprIndexOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIndexOfContext) LlamadaIndexOf() ILlamadaIndexOfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaIndexOfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaIndexOfContext)
}

func (s *ExprIndexOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprIndexOf(s)
	}
}

func (s *ExprIndexOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprIndexOf(s)
	}
}

func (s *ExprIndexOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprIndexOf(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAccesoInstanciaAtributoContext struct {
	ExpresionContext
}

func NewExprAccesoInstanciaAtributoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAccesoInstanciaAtributoContext {
	var p = new(ExprAccesoInstanciaAtributoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprAccesoInstanciaAtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAccesoInstanciaAtributoContext) AccesoInstaciaAtributo() IAccesoInstaciaAtributoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesoInstaciaAtributoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesoInstaciaAtributoContext)
}

func (s *ExprAccesoInstanciaAtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprAccesoInstanciaAtributo(s)
	}
}

func (s *ExprAccesoInstanciaAtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprAccesoInstanciaAtributo(s)
	}
}

func (s *ExprAccesoInstanciaAtributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprAccesoInstanciaAtributo(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprRelacionalContext struct {
	ExpresionContext
	left  IExpresionContext
	op    antlr.Token
	right IExpresionContext
}

func NewExprRelacionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprRelacionalContext {
	var p = new(ExprRelacionalContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprRelacionalContext) GetOp() antlr.Token { return s.op }

func (s *ExprRelacionalContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprRelacionalContext) GetLeft() IExpresionContext { return s.left }

func (s *ExprRelacionalContext) GetRight() IExpresionContext { return s.right }

func (s *ExprRelacionalContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExprRelacionalContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExprRelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRelacionalContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExprRelacionalContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprRelacionalContext) SIGNO_IGUALDAD() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_IGUALDAD, 0)
}

func (s *ExprRelacionalContext) SIGNO_NO_IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_NO_IGUAL, 0)
}

func (s *ExprRelacionalContext) SIGNO_MENOSQUE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MENOSQUE, 0)
}

func (s *ExprRelacionalContext) SIGNO_MENOSQUE_IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MENOSQUE_IGUAL, 0)
}

func (s *ExprRelacionalContext) SIGNO_MASQUE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MASQUE, 0)
}

func (s *ExprRelacionalContext) SIGNO_MASQUE_IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_MASQUE_IGUAL, 0)
}

func (s *ExprRelacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprRelacional(s)
	}
}

func (s *ExprRelacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprRelacional(s)
	}
}

func (s *ExprRelacionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprRelacional(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStringContext struct {
	ExpresionContext
}

func NewExprStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStringContext {
	var p = new(ExprStringContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStringContext) CADENA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCADENA, 0)
}

func (s *ExprStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprString(s)
	}
}

func (s *ExprStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprString(s)
	}
}

func (s *ExprStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIntContext struct {
	ExpresionContext
}

func NewExprIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIntContext {
	var p = new(ExprIntContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIntContext) INT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINT, 0)
}

func (s *ExprIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprInt(s)
	}
}

func (s *ExprIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprInt(s)
	}
}

func (s *ExprIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprParentesisContext struct {
	ExpresionContext
}

func NewExprParentesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParentesisContext {
	var p = new(ExprParentesisContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprParentesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParentesisContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *ExprParentesisContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExprParentesisContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *ExprParentesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprParentesis(s)
	}
}

func (s *ExprParentesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprParentesis(s)
	}
}

func (s *ExprParentesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprParentesis(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLenContext struct {
	ExpresionContext
}

func NewExprLenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLenContext {
	var p = new(ExprLenContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprLenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLenContext) LlamadaLen() ILlamadaLenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaLenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaLenContext)
}

func (s *ExprLenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprLen(s)
	}
}

func (s *ExprLenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprLen(s)
	}
}

func (s *ExprLenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprLen(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprSliceContext struct {
	ExpresionContext
}

func NewExprSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprSliceContext {
	var p = new(ExprSliceContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSliceContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *ExprSliceContext) Lista_expresion() ILista_expresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionContext)
}

func (s *ExprSliceContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *ExprSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprSlice(s)
	}
}

func (s *ExprSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprSlice(s)
	}
}

func (s *ExprSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprJoinContext struct {
	ExpresionContext
}

func NewExprJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprJoinContext {
	var p = new(ExprJoinContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprJoinContext) LlamadaJoin() ILlamadaJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaJoinContext)
}

func (s *ExprJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprJoin(s)
	}
}

func (s *ExprJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprJoin(s)
	}
}

func (s *ExprJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAppendContext struct {
	ExpresionContext
}

func NewExprAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAppendContext {
	var p = new(ExprAppendContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAppendContext) LlamadaAppend() ILlamadaAppendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaAppendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaAppendContext)
}

func (s *ExprAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprAppend(s)
	}
}

func (s *ExprAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprAppend(s)
	}
}

func (s *ExprAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLlamadaFuncionStructContext struct {
	ExpresionContext
	nomInstancia antlr.Token
	nomFunc      antlr.Token
}

func NewExprLlamadaFuncionStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLlamadaFuncionStructContext {
	var p = new(ExprLlamadaFuncionStructContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *ExprLlamadaFuncionStructContext) GetNomInstancia() antlr.Token { return s.nomInstancia }

func (s *ExprLlamadaFuncionStructContext) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *ExprLlamadaFuncionStructContext) SetNomInstancia(v antlr.Token) { s.nomInstancia = v }

func (s *ExprLlamadaFuncionStructContext) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *ExprLlamadaFuncionStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLlamadaFuncionStructContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO, 0)
}

func (s *ExprLlamadaFuncionStructContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *ExprLlamadaFuncionStructContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *ExprLlamadaFuncionStructContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *ExprLlamadaFuncionStructContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *ExprLlamadaFuncionStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExprLlamadaFuncionStruct(s)
	}
}

func (s *ExprLlamadaFuncionStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExprLlamadaFuncionStruct(s)
	}
}

func (s *ExprLlamadaFuncionStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExprLlamadaFuncionStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *gramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, gramaticaParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprNegUnariaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(380)

			var _m = p.Match(gramaticaParserSIGNO_MENOS)

			localctx.(*ExprNegUnariaContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.expresion(0)
		}
		{
			p.SetState(383)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprParentesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(385)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.expresion(0)
		}
		{
			p.SetState(387)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExprSliceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(389)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Lista_expresion()
		}
		{
			p.SetState(391)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExprNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)

			var _m = p.Match(gramaticaParserSIGNO_NOT)

			localctx.(*ExprNotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.expresion(21)
		}

	case 5:
		localctx = NewExprIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(395)
			p.Match(gramaticaParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExprFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(396)
			p.Match(gramaticaParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExprBoolTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(397)

			var _m = p.Match(gramaticaParserT__0)

			localctx.(*ExprBoolTrueContext).valor = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExprBoolFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(398)

			var _m = p.Match(gramaticaParserT__1)

			localctx.(*ExprBoolFalseContext).valor = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewExprStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(399)
			p.Match(gramaticaParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExprRuneContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(400)
			p.Match(gramaticaParserRUNE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewExprNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(401)
			p.Match(gramaticaParserRESNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewExprIdentificadorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(402)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewExprLlamadaFuncionStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(403)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*ExprLlamadaFuncionStructContext).nomInstancia = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(gramaticaParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*ExprLlamadaFuncionStructContext).nomFunc = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewExprLlamadaFuncionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(408)
			p.Llamadafuncion()
		}

	case 15:
		localctx = NewExprIndexOfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(409)
			p.LlamadaIndexOf()
		}

	case 16:
		localctx = NewExprJoinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(410)
			p.LlamadaJoin()
		}

	case 17:
		localctx = NewExprLenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(411)
			p.LlamadaLen()
		}

	case 18:
		localctx = NewExprAppendContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			p.LlamadaAppend()
		}

	case 19:
		localctx = NewExprAccesoMatrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.AccesoMatriz()
		}

	case 20:
		localctx = NewExprAccesoInstanciaAtributoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.AccesoInstaciaAtributo()
		}

	case 21:
		localctx = NewExprAccesoSliceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(415)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expresion(0)
		}
		{
			p.SetState(418)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 22:
		localctx = NewExprAccesoMatrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.expresion(0)
		}
		{
			p.SetState(423)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.expresion(0)
		}
		{
			p.SetState(426)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 23:
		localctx = NewExprSliceLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(428)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Tipo()
		}
		{
			p.SetState(431)
			p.Match(gramaticaParserLLAVEABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Lista_expresion()
		}
		{
			p.SetState(433)
			p.Match(gramaticaParserLLAVECIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprMultDivModContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprMultDivModContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(438)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprMultDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-51)) & ^0x3f) == 0 && ((int64(1)<<(_la-51))&524291) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprMultDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(439)

					var _x = p.expresion(29)

					localctx.(*ExprMultDivModContext).right = _x
				}

			case 2:
				localctx = NewExprOpAritmeticaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprOpAritmeticaContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(441)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprOpAritmeticaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserSIGNO_MAS || _la == gramaticaParserSIGNO_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprOpAritmeticaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(442)

					var _x = p.expresion(28)

					localctx.(*ExprOpAritmeticaContext).right = _x
				}

			case 3:
				localctx = NewExprRelacionalContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprRelacionalContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(444)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprRelacionalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserSIGNO_NO_IGUAL || _la == gramaticaParserSIGNO_IGUALDAD) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprRelacionalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(445)

					var _x = p.expresion(25)

					localctx.(*ExprRelacionalContext).right = _x
				}

			case 4:
				localctx = NewExprRelacionalContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprRelacionalContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(447)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprRelacionalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserSIGNO_MENOSQUE_IGUAL || _la == gramaticaParserSIGNO_MENOSQUE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprRelacionalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(448)

					var _x = p.expresion(24)

					localctx.(*ExprRelacionalContext).right = _x
				}

			case 5:
				localctx = NewExprRelacionalContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprRelacionalContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(450)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprRelacionalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserSIGNO_MASQUE_IGUAL || _la == gramaticaParserSIGNO_MASQUE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprRelacionalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(451)

					var _x = p.expresion(23)

					localctx.(*ExprRelacionalContext).right = _x
				}

			case 6:
				localctx = NewExprLogicaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				localctx.(*ExprLogicaContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(452)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(453)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprLogicaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == gramaticaParserSIGNO_AND || _la == gramaticaParserSIGNO_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprLogicaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(454)

					var _x = p.expresion(21)

					localctx.(*ExprLogicaContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtributosStructContext is an interface to support dynamic dispatch.
type IAtributosStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomAtributo returns the nomAtributo token.
	GetNomAtributo() antlr.Token

	// SetNomAtributo sets the nomAtributo token.
	SetNomAtributo(antlr.Token)

	// Getter signatures
	Tipo() ITipoContext
	IDENTIFICADOR() antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsAtributosStructContext differentiates from other interfaces.
	IsAtributosStructContext()
}

type AtributosStructContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	nomAtributo antlr.Token
}

func NewEmptyAtributosStructContext() *AtributosStructContext {
	var p = new(AtributosStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_atributosStruct
	return p
}

func InitEmptyAtributosStructContext(p *AtributosStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_atributosStruct
}

func (*AtributosStructContext) IsAtributosStructContext() {}

func NewAtributosStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosStructContext {
	var p = new(AtributosStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_atributosStruct

	return p
}

func (s *AtributosStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosStructContext) GetNomAtributo() antlr.Token { return s.nomAtributo }

func (s *AtributosStructContext) SetNomAtributo(v antlr.Token) { s.nomAtributo = v }

func (s *AtributosStructContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *AtributosStructContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AtributosStructContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *AtributosStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributosStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAtributosStruct(s)
	}
}

func (s *AtributosStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAtributosStruct(s)
	}
}

func (s *AtributosStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAtributosStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AtributosStruct() (localctx IAtributosStructContext) {
	localctx = NewAtributosStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gramaticaParserRULE_atributosStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Tipo()
	}
	{
		p.SetState(461)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*AtributosStructContext).nomAtributo = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserPUNTO_COMA {
		{
			p.SetState(462)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInicializacion_sliceContext is an interface to support dynamic dispatch.
type IInicializacion_sliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CORABRE() antlr.TerminalNode
	CORCIERRA() antlr.TerminalNode
	Lista_expresion() ILista_expresionContext
	Lista_slices() ILista_slicesContext

	// IsInicializacion_sliceContext differentiates from other interfaces.
	IsInicializacion_sliceContext()
}

type Inicializacion_sliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicializacion_sliceContext() *Inicializacion_sliceContext {
	var p = new(Inicializacion_sliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_inicializacion_slice
	return p
}

func InitEmptyInicializacion_sliceContext(p *Inicializacion_sliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_inicializacion_slice
}

func (*Inicializacion_sliceContext) IsInicializacion_sliceContext() {}

func NewInicializacion_sliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inicializacion_sliceContext {
	var p = new(Inicializacion_sliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_inicializacion_slice

	return p
}

func (s *Inicializacion_sliceContext) GetParser() antlr.Parser { return s.parser }

func (s *Inicializacion_sliceContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *Inicializacion_sliceContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *Inicializacion_sliceContext) Lista_expresion() ILista_expresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionContext)
}

func (s *Inicializacion_sliceContext) Lista_slices() ILista_slicesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_slicesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_slicesContext)
}

func (s *Inicializacion_sliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inicializacion_sliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inicializacion_sliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInicializacion_slice(s)
	}
}

func (s *Inicializacion_sliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInicializacion_slice(s)
	}
}

func (s *Inicializacion_sliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInicializacion_slice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Inicializacion_slice() (localctx IInicializacion_sliceContext) {
	localctx = NewInicializacion_sliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gramaticaParserRULE_inicializacion_slice)
	var _la int

	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4900068058465779718) != 0 {
			{
				p.SetState(466)
				p.Lista_expresion()
			}

		}
		{
			p.SetState(469)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)
			p.Match(gramaticaParserCORABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserCORABRE {
			{
				p.SetState(471)
				p.Lista_slices()
			}

		}
		{
			p.SetState(474)
			p.Match(gramaticaParserCORCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_expresionContext is an interface to support dynamic dispatch.
type ILista_expresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsLista_expresionContext differentiates from other interfaces.
	IsLista_expresionContext()
}

type Lista_expresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_expresionContext() *Lista_expresionContext {
	var p = new(Lista_expresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_expresion
	return p
}

func InitEmptyLista_expresionContext(p *Lista_expresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_expresion
}

func (*Lista_expresionContext) IsLista_expresionContext() {}

func NewLista_expresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_expresionContext {
	var p = new(Lista_expresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_expresion

	return p
}

func (s *Lista_expresionContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_expresionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Lista_expresionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_expresionContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *Lista_expresionContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *Lista_expresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_expresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_expresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_expresion(s)
	}
}

func (s *Lista_expresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_expresion(s)
	}
}

func (s *Lista_expresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_expresion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_expresion() (localctx ILista_expresionContext) {
	localctx = NewLista_expresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gramaticaParserRULE_lista_expresion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.expresion(0)
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(478)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.expresion(0)
		}

		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_slicesContext is an interface to support dynamic dispatch.
type ILista_slicesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInicializacion_slice() []IInicializacion_sliceContext
	Inicializacion_slice(i int) IInicializacion_sliceContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsLista_slicesContext differentiates from other interfaces.
	IsLista_slicesContext()
}

type Lista_slicesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_slicesContext() *Lista_slicesContext {
	var p = new(Lista_slicesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_slices
	return p
}

func InitEmptyLista_slicesContext(p *Lista_slicesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_slices
}

func (*Lista_slicesContext) IsLista_slicesContext() {}

func NewLista_slicesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_slicesContext {
	var p = new(Lista_slicesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_slices

	return p
}

func (s *Lista_slicesContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_slicesContext) AllInicializacion_slice() []IInicializacion_sliceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInicializacion_sliceContext); ok {
			len++
		}
	}

	tst := make([]IInicializacion_sliceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInicializacion_sliceContext); ok {
			tst[i] = t.(IInicializacion_sliceContext)
			i++
		}
	}

	return tst
}

func (s *Lista_slicesContext) Inicializacion_slice(i int) IInicializacion_sliceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInicializacion_sliceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInicializacion_sliceContext)
}

func (s *Lista_slicesContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *Lista_slicesContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *Lista_slicesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_slicesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_slicesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_slices(s)
	}
}

func (s *Lista_slicesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_slices(s)
	}
}

func (s *Lista_slicesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_slices(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_slices() (localctx ILista_slicesContext) {
	localctx = NewLista_slicesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gramaticaParserRULE_lista_slices)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Inicializacion_slice()
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(486)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)
			p.Inicializacion_slice()
		}

		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESINT() antlr.TerminalNode
	RESFLOAT() antlr.TerminalNode
	RESSTRING() antlr.TerminalNode
	RESBOOL() antlr.TerminalNode
	RESRUNE() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) RESINT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESINT, 0)
}

func (s *TipoContext) RESFLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFLOAT, 0)
}

func (s *TipoContext) RESSTRING() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESSTRING, 0)
}

func (s *TipoContext) RESBOOL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESBOOL, 0)
}

func (s *TipoContext) RESRUNE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESRUNE, 0)
}

func (s *TipoContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gramaticaParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719484672) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAVEABRE() antlr.TerminalNode
	LLAVECIERRA() antlr.TerminalNode
	AllLista_sentencias() []ILista_sentenciasContext
	Lista_sentencias(i int) ILista_sentenciasContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *BloqueContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *BloqueContext) AllLista_sentencias() []ILista_sentenciasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_sentenciasContext); ok {
			len++
		}
	}

	tst := make([]ILista_sentenciasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_sentenciasContext); ok {
			tst[i] = t.(ILista_sentenciasContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Lista_sentencias(i int) ILista_sentenciasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_sentenciasContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_sentenciasContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gramaticaParserRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35258353320064) != 0 {
		{
			p.SetState(496)
			p.Lista_sentencias()
		}

		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(502)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESIMPORT() antlr.TerminalNode
	CADENA() antlr.TerminalNode
	AllPUNTO_COMA() []antlr.TerminalNode
	PUNTO_COMA(i int) antlr.TerminalNode

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) RESIMPORT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESIMPORT, 0)
}

func (s *ImportStmtContext) CADENA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCADENA, 0)
}

func (s *ImportStmtContext) AllPUNTO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPUNTO_COMA)
}

func (s *ImportStmtContext) PUNTO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, i)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gramaticaParserRULE_importStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.Match(gramaticaParserRESIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(505)
		p.Match(gramaticaParserCADENA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gramaticaParserPUNTO_COMA {
		{
			p.SetState(507)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreacion_instancia_structContext is an interface to support dynamic dispatch.
type ICreacion_instancia_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomInstancia returns the nomInstancia token.
	GetNomInstancia() antlr.Token

	// GetNomStruct returns the nomStruct token.
	GetNomStruct() antlr.Token

	// SetNomInstancia sets the nomInstancia token.
	SetNomInstancia(antlr.Token)

	// SetNomStruct sets the nomStruct token.
	SetNomStruct(antlr.Token)

	// Getter signatures
	RESMUT() antlr.TerminalNode
	ASIGNACION_INFERIDA() antlr.TerminalNode
	LLAVEABRE() antlr.TerminalNode
	LLAVECIERRA() antlr.TerminalNode
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	ListaCampos() IListaCamposContext
	PUNTO_COMA() antlr.TerminalNode

	// IsCreacion_instancia_structContext differentiates from other interfaces.
	IsCreacion_instancia_structContext()
}

type Creacion_instancia_structContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	nomInstancia antlr.Token
	nomStruct    antlr.Token
}

func NewEmptyCreacion_instancia_structContext() *Creacion_instancia_structContext {
	var p = new(Creacion_instancia_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_creacion_instancia_struct
	return p
}

func InitEmptyCreacion_instancia_structContext(p *Creacion_instancia_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_creacion_instancia_struct
}

func (*Creacion_instancia_structContext) IsCreacion_instancia_structContext() {}

func NewCreacion_instancia_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Creacion_instancia_structContext {
	var p = new(Creacion_instancia_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_creacion_instancia_struct

	return p
}

func (s *Creacion_instancia_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Creacion_instancia_structContext) GetNomInstancia() antlr.Token { return s.nomInstancia }

func (s *Creacion_instancia_structContext) GetNomStruct() antlr.Token { return s.nomStruct }

func (s *Creacion_instancia_structContext) SetNomInstancia(v antlr.Token) { s.nomInstancia = v }

func (s *Creacion_instancia_structContext) SetNomStruct(v antlr.Token) { s.nomStruct = v }

func (s *Creacion_instancia_structContext) RESMUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESMUT, 0)
}

func (s *Creacion_instancia_structContext) ASIGNACION_INFERIDA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserASIGNACION_INFERIDA, 0)
}

func (s *Creacion_instancia_structContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *Creacion_instancia_structContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *Creacion_instancia_structContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *Creacion_instancia_structContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *Creacion_instancia_structContext) ListaCampos() IListaCamposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaCamposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaCamposContext)
}

func (s *Creacion_instancia_structContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Creacion_instancia_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Creacion_instancia_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Creacion_instancia_structContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCreacion_instancia_struct(s)
	}
}

func (s *Creacion_instancia_structContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCreacion_instancia_struct(s)
	}
}

func (s *Creacion_instancia_structContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCreacion_instancia_struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Creacion_instancia_struct() (localctx ICreacion_instancia_structContext) {
	localctx = NewCreacion_instancia_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gramaticaParserRULE_creacion_instancia_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(gramaticaParserRESMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(513)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Creacion_instancia_structContext).nomInstancia = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(gramaticaParserASIGNACION_INFERIDA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Creacion_instancia_structContext).nomStruct = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(516)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserIDENTIFICADOR {
		{
			p.SetState(517)
			p.ListaCampos()
		}

	}
	{
		p.SetState(520)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(521)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametros_funcContext is an interface to support dynamic dispatch.
type ILista_parametros_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsLista_parametros_funcContext differentiates from other interfaces.
	IsLista_parametros_funcContext()
}

type Lista_parametros_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametros_funcContext() *Lista_parametros_funcContext {
	var p = new(Lista_parametros_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_parametros_func
	return p
}

func InitEmptyLista_parametros_funcContext(p *Lista_parametros_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_parametros_func
}

func (*Lista_parametros_funcContext) IsLista_parametros_funcContext() {}

func NewLista_parametros_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametros_funcContext {
	var p = new(Lista_parametros_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_parametros_func

	return p
}

func (s *Lista_parametros_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametros_funcContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Lista_parametros_funcContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_parametros_funcContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *Lista_parametros_funcContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *Lista_parametros_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametros_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_parametros_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_parametros_func(s)
	}
}

func (s *Lista_parametros_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_parametros_func(s)
	}
}

func (s *Lista_parametros_funcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_parametros_func(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_parametros_func() (localctx ILista_parametros_funcContext) {
	localctx = NewLista_parametros_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gramaticaParserRULE_lista_parametros_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.expresion(0)
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(525)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.expresion(0)
		}

		p.SetState(531)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParam_funcContext is an interface to support dynamic dispatch.
type IParam_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdInferencia returns the idInferencia token.
	GetIdInferencia() antlr.Token

	// SetIdInferencia sets the idInferencia token.
	SetIdInferencia(antlr.Token)

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsParam_funcContext differentiates from other interfaces.
	IsParam_funcContext()
}

type Param_funcContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	idInferencia antlr.Token
}

func NewEmptyParam_funcContext() *Param_funcContext {
	var p = new(Param_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_param_func
	return p
}

func InitEmptyParam_funcContext(p *Param_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_param_func
}

func (*Param_funcContext) IsParam_funcContext() {}

func NewParam_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_funcContext {
	var p = new(Param_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_param_func

	return p
}

func (s *Param_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_funcContext) GetIdInferencia() antlr.Token { return s.idInferencia }

func (s *Param_funcContext) SetIdInferencia(v antlr.Token) { s.idInferencia = v }

func (s *Param_funcContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Param_funcContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Param_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterParam_func(s)
	}
}

func (s *Param_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitParam_func(s)
	}
}

func (s *Param_funcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitParam_func(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Param_func() (localctx IParam_funcContext) {
	localctx = NewParam_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gramaticaParserRULE_param_func)
	var _la int

	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserT__2 {
			{
				p.SetState(532)
				p.Match(gramaticaParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(535)

			var _m = p.Match(gramaticaParserIDENTIFICADOR)

			localctx.(*Param_funcContext).idInferencia = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(536)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaForContext is an interface to support dynamic dispatch.
type ISentenciaForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaForContext differentiates from other interfaces.
	IsSentenciaForContext()
}

type SentenciaForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaForContext() *SentenciaForContext {
	var p = new(SentenciaForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaFor
	return p
}

func InitEmptySentenciaForContext(p *SentenciaForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentenciaFor
}

func (*SentenciaForContext) IsSentenciaForContext() {}

func NewSentenciaForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaForContext {
	var p = new(SentenciaForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentenciaFor

	return p
}

func (s *SentenciaForContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaForContext) CopyAll(ctx *SentenciaForContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SentFor1Context struct {
	SentenciaForContext
}

func NewSentFor1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentFor1Context {
	var p = new(SentFor1Context)

	InitEmptySentenciaForContext(&p.SentenciaForContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaForContext))

	return p
}

func (s *SentFor1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentFor1Context) RESFOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFOR, 0)
}

func (s *SentFor1Context) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentFor1Context) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentFor1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentFor1(s)
	}
}

func (s *SentFor1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentFor1(s)
	}
}

func (s *SentFor1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentFor1(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentForClasicoContext struct {
	SentenciaForContext
}

func NewSentForClasicoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentForClasicoContext {
	var p = new(SentForClasicoContext)

	InitEmptySentenciaForContext(&p.SentenciaForContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaForContext))

	return p
}

func (s *SentForClasicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentForClasicoContext) RESFOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESFOR, 0)
}

func (s *SentForClasicoContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *SentForClasicoContext) Sentencia_init() ISentencia_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencia_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencia_initContext)
}

func (s *SentForClasicoContext) AllPUNTO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPUNTO_COMA)
}

func (s *SentForClasicoContext) PUNTO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, i)
}

func (s *SentForClasicoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentForClasicoContext) Sentencia_update() ISentencia_updateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentencia_updateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentencia_updateContext)
}

func (s *SentForClasicoContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *SentForClasicoContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SentForClasicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentForClasico(s)
	}
}

func (s *SentForClasicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentForClasico(s)
	}
}

func (s *SentForClasicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentForClasico(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SentenciaFor() (localctx ISentenciaForContext) {
	localctx = NewSentenciaForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gramaticaParserRULE_sentenciaFor)
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSentFor1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)
			p.Match(gramaticaParserRESFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.expresion(0)
		}
		{
			p.SetState(541)
			p.Bloque()
		}

	case 2:
		localctx = NewSentForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.Match(gramaticaParserRESFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.Match(gramaticaParserPARABRE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Sentencia_init()
		}
		{
			p.SetState(546)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.expresion(0)
		}
		{
			p.SetState(548)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.Sentencia_update()
		}
		{
			p.SetState(550)
			p.Match(gramaticaParserPARCIERRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Bloque()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentencia_initContext is an interface to support dynamic dispatch.
type ISentencia_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion_variable() IDeclaracion_variableContext
	Asignacion() IAsignacionContext

	// IsSentencia_initContext differentiates from other interfaces.
	IsSentencia_initContext()
}

type Sentencia_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencia_initContext() *Sentencia_initContext {
	var p = new(Sentencia_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_init
	return p
}

func InitEmptySentencia_initContext(p *Sentencia_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_init
}

func (*Sentencia_initContext) IsSentencia_initContext() {}

func NewSentencia_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencia_initContext {
	var p = new(Sentencia_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentencia_init

	return p
}

func (s *Sentencia_initContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencia_initContext) Declaracion_variable() IDeclaracion_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_variableContext)
}

func (s *Sentencia_initContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Sentencia_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencia_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sentencia_initContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentencia_init(s)
	}
}

func (s *Sentencia_initContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentencia_init(s)
	}
}

func (s *Sentencia_initContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentencia_init(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sentencia_init() (localctx ISentencia_initContext) {
	localctx = NewSentencia_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gramaticaParserRULE_sentencia_init)
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.Declaracion_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)
			p.Asignacion()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentencia_updateContext is an interface to support dynamic dispatch.
type ISentencia_updateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Asignacion() IAsignacionContext
	Incremento_variable() IIncremento_variableContext

	// IsSentencia_updateContext differentiates from other interfaces.
	IsSentencia_updateContext()
}

type Sentencia_updateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentencia_updateContext() *Sentencia_updateContext {
	var p = new(Sentencia_updateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_update
	return p
}

func InitEmptySentencia_updateContext(p *Sentencia_updateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sentencia_update
}

func (*Sentencia_updateContext) IsSentencia_updateContext() {}

func NewSentencia_updateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sentencia_updateContext {
	var p = new(Sentencia_updateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sentencia_update

	return p
}

func (s *Sentencia_updateContext) GetParser() antlr.Parser { return s.parser }

func (s *Sentencia_updateContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Sentencia_updateContext) Incremento_variable() IIncremento_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncremento_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncremento_variableContext)
}

func (s *Sentencia_updateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sentencia_updateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sentencia_updateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSentencia_update(s)
	}
}

func (s *Sentencia_updateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSentencia_update(s)
	}
}

func (s *Sentencia_updateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSentencia_update(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sentencia_update() (localctx ISentencia_updateContext) {
	localctx = NewSentencia_updateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gramaticaParserRULE_sentencia_update)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(559)
			p.Asignacion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(560)
			p.Incremento_variable()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncremento_variableContext is an interface to support dynamic dispatch.
type IIncremento_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	SIGNO_INCREMENTO() antlr.TerminalNode
	SIGNO_DECREMENTO() antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsIncremento_variableContext differentiates from other interfaces.
	IsIncremento_variableContext()
}

type Incremento_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyIncremento_variableContext() *Incremento_variableContext {
	var p = new(Incremento_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_incremento_variable
	return p
}

func InitEmptyIncremento_variableContext(p *Incremento_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_incremento_variable
}

func (*Incremento_variableContext) IsIncremento_variableContext() {}

func NewIncremento_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Incremento_variableContext {
	var p = new(Incremento_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_incremento_variable

	return p
}

func (s *Incremento_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Incremento_variableContext) GetOp() antlr.Token { return s.op }

func (s *Incremento_variableContext) SetOp(v antlr.Token) { s.op = v }

func (s *Incremento_variableContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Incremento_variableContext) SIGNO_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_INCREMENTO, 0)
}

func (s *Incremento_variableContext) SIGNO_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_DECREMENTO, 0)
}

func (s *Incremento_variableContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Incremento_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Incremento_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Incremento_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIncremento_variable(s)
	}
}

func (s *Incremento_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIncremento_variable(s)
	}
}

func (s *Incremento_variableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitIncremento_variable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Incremento_variable() (localctx IIncremento_variableContext) {
	localctx = NewIncremento_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gramaticaParserRULE_incremento_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(564)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Incremento_variableContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == gramaticaParserSIGNO_INCREMENTO || _la == gramaticaParserSIGNO_DECREMENTO) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Incremento_variableContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserPUNTO_COMA {
		{
			p.SetState(565)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaIndexOfContext is an interface to support dynamic dispatch.
type ILlamadaIndexOfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomSlice returns the nomSlice token.
	GetNomSlice() antlr.Token

	// SetNomSlice sets the nomSlice token.
	SetNomSlice(antlr.Token)

	// Getter signatures
	RESINDEX() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	SIGNO_COMA() antlr.TerminalNode
	Expresion() IExpresionContext
	PARCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsLlamadaIndexOfContext differentiates from other interfaces.
	IsLlamadaIndexOfContext()
}

type LlamadaIndexOfContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	nomSlice antlr.Token
}

func NewEmptyLlamadaIndexOfContext() *LlamadaIndexOfContext {
	var p = new(LlamadaIndexOfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaIndexOf
	return p
}

func InitEmptyLlamadaIndexOfContext(p *LlamadaIndexOfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaIndexOf
}

func (*LlamadaIndexOfContext) IsLlamadaIndexOfContext() {}

func NewLlamadaIndexOfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaIndexOfContext {
	var p = new(LlamadaIndexOfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaIndexOf

	return p
}

func (s *LlamadaIndexOfContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaIndexOfContext) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *LlamadaIndexOfContext) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *LlamadaIndexOfContext) RESINDEX() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESINDEX, 0)
}

func (s *LlamadaIndexOfContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaIndexOfContext) SIGNO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, 0)
}

func (s *LlamadaIndexOfContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaIndexOfContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaIndexOfContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaIndexOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaIndexOfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaIndexOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaIndexOf(s)
	}
}

func (s *LlamadaIndexOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaIndexOf(s)
	}
}

func (s *LlamadaIndexOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaIndexOf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaIndexOf() (localctx ILlamadaIndexOfContext) {
	localctx = NewLlamadaIndexOfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gramaticaParserRULE_llamadaIndexOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(568)
		p.Match(gramaticaParserRESINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(569)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(570)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*LlamadaIndexOfContext).nomSlice = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(571)
		p.Match(gramaticaParserSIGNO_COMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(572)
		p.expresion(0)
	}
	{
		p.SetState(573)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaJoinContext is an interface to support dynamic dispatch.
type ILlamadaJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomSlice returns the nomSlice token.
	GetNomSlice() antlr.Token

	// SetNomSlice sets the nomSlice token.
	SetNomSlice(antlr.Token)

	// Getter signatures
	RESJOIN() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	SIGNO_COMA() antlr.TerminalNode
	Expresion() IExpresionContext
	PARCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsLlamadaJoinContext differentiates from other interfaces.
	IsLlamadaJoinContext()
}

type LlamadaJoinContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	nomSlice antlr.Token
}

func NewEmptyLlamadaJoinContext() *LlamadaJoinContext {
	var p = new(LlamadaJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaJoin
	return p
}

func InitEmptyLlamadaJoinContext(p *LlamadaJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaJoin
}

func (*LlamadaJoinContext) IsLlamadaJoinContext() {}

func NewLlamadaJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaJoinContext {
	var p = new(LlamadaJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaJoin

	return p
}

func (s *LlamadaJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaJoinContext) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *LlamadaJoinContext) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *LlamadaJoinContext) RESJOIN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESJOIN, 0)
}

func (s *LlamadaJoinContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaJoinContext) SIGNO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, 0)
}

func (s *LlamadaJoinContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaJoinContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaJoinContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaJoin(s)
	}
}

func (s *LlamadaJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaJoin(s)
	}
}

func (s *LlamadaJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaJoin() (localctx ILlamadaJoinContext) {
	localctx = NewLlamadaJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, gramaticaParserRULE_llamadaJoin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(575)
		p.Match(gramaticaParserRESJOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(576)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(577)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*LlamadaJoinContext).nomSlice = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(578)
		p.Match(gramaticaParserSIGNO_COMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(579)
		p.expresion(0)
	}
	{
		p.SetState(580)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaLenContext is an interface to support dynamic dispatch.
type ILlamadaLenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomSlice returns the nomSlice token.
	GetNomSlice() antlr.Token

	// SetNomSlice sets the nomSlice token.
	SetNomSlice(antlr.Token)

	// Getter signatures
	RESLEN() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	PARCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsLlamadaLenContext differentiates from other interfaces.
	IsLlamadaLenContext()
}

type LlamadaLenContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	nomSlice antlr.Token
}

func NewEmptyLlamadaLenContext() *LlamadaLenContext {
	var p = new(LlamadaLenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaLen
	return p
}

func InitEmptyLlamadaLenContext(p *LlamadaLenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaLen
}

func (*LlamadaLenContext) IsLlamadaLenContext() {}

func NewLlamadaLenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaLenContext {
	var p = new(LlamadaLenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaLen

	return p
}

func (s *LlamadaLenContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaLenContext) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *LlamadaLenContext) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *LlamadaLenContext) RESLEN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESLEN, 0)
}

func (s *LlamadaLenContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaLenContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaLenContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaLenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaLenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaLenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaLen(s)
	}
}

func (s *LlamadaLenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaLen(s)
	}
}

func (s *LlamadaLenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaLen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaLen() (localctx ILlamadaLenContext) {
	localctx = NewLlamadaLenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, gramaticaParserRULE_llamadaLen)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(582)
		p.Match(gramaticaParserRESLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(583)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(584)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*LlamadaLenContext).nomSlice = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(585)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaAppendContext is an interface to support dynamic dispatch.
type ILlamadaAppendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomSlice returns the nomSlice token.
	GetNomSlice() antlr.Token

	// SetNomSlice sets the nomSlice token.
	SetNomSlice(antlr.Token)

	// Getter signatures
	RESAPPEND() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	SIGNO_COMA() antlr.TerminalNode
	Expresion() IExpresionContext
	PARCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsLlamadaAppendContext differentiates from other interfaces.
	IsLlamadaAppendContext()
}

type LlamadaAppendContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	nomSlice antlr.Token
}

func NewEmptyLlamadaAppendContext() *LlamadaAppendContext {
	var p = new(LlamadaAppendContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaAppend
	return p
}

func InitEmptyLlamadaAppendContext(p *LlamadaAppendContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaAppend
}

func (*LlamadaAppendContext) IsLlamadaAppendContext() {}

func NewLlamadaAppendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaAppendContext {
	var p = new(LlamadaAppendContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaAppend

	return p
}

func (s *LlamadaAppendContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaAppendContext) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *LlamadaAppendContext) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *LlamadaAppendContext) RESAPPEND() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESAPPEND, 0)
}

func (s *LlamadaAppendContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaAppendContext) SIGNO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, 0)
}

func (s *LlamadaAppendContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaAppendContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaAppendContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaAppendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaAppend(s)
	}
}

func (s *LlamadaAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaAppend(s)
	}
}

func (s *LlamadaAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaAppend() (localctx ILlamadaAppendContext) {
	localctx = NewLlamadaAppendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, gramaticaParserRULE_llamadaAppend)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		p.Match(gramaticaParserRESAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(588)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(589)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*LlamadaAppendContext).nomSlice = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(590)
		p.Match(gramaticaParserSIGNO_COMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(591)
		p.expresion(0)
	}
	{
		p.SetState(592)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccesoSliceContext is an interface to support dynamic dispatch.
type IAccesoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomSlice returns the nomSlice token.
	GetNomSlice() antlr.Token

	// SetNomSlice sets the nomSlice token.
	SetNomSlice(antlr.Token)

	// GetIndex returns the index rule contexts.
	GetIndex() IExpresionContext

	// SetIndex sets the index rule contexts.
	SetIndex(IExpresionContext)

	// Getter signatures
	CORABRE() antlr.TerminalNode
	CORCIERRA() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAccesoSliceContext differentiates from other interfaces.
	IsAccesoSliceContext()
}

type AccesoSliceContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	nomSlice antlr.Token
	index    IExpresionContext
}

func NewEmptyAccesoSliceContext() *AccesoSliceContext {
	var p = new(AccesoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoSlice
	return p
}

func InitEmptyAccesoSliceContext(p *AccesoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoSlice
}

func (*AccesoSliceContext) IsAccesoSliceContext() {}

func NewAccesoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoSliceContext {
	var p = new(AccesoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_accesoSlice

	return p
}

func (s *AccesoSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoSliceContext) GetNomSlice() antlr.Token { return s.nomSlice }

func (s *AccesoSliceContext) SetNomSlice(v antlr.Token) { s.nomSlice = v }

func (s *AccesoSliceContext) GetIndex() IExpresionContext { return s.index }

func (s *AccesoSliceContext) SetIndex(v IExpresionContext) { s.index = v }

func (s *AccesoSliceContext) CORABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, 0)
}

func (s *AccesoSliceContext) CORCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, 0)
}

func (s *AccesoSliceContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AccesoSliceContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AccesoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAccesoSlice(s)
	}
}

func (s *AccesoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAccesoSlice(s)
	}
}

func (s *AccesoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAccesoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AccesoSlice() (localctx IAccesoSliceContext) {
	localctx = NewAccesoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, gramaticaParserRULE_accesoSlice)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*AccesoSliceContext).nomSlice = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(595)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(596)

		var _x = p.expresion(0)

		localctx.(*AccesoSliceContext).index = _x
	}
	{
		p.SetState(597)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_matrizContext is an interface to support dynamic dispatch.
type IDeclaracion_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdMatriz returns the idMatriz token.
	GetIdMatriz() antlr.Token

	// SetIdMatriz sets the idMatriz token.
	SetIdMatriz(antlr.Token)

	// Getter signatures
	ASIGNACION_INFERIDA() antlr.TerminalNode
	AllLLAVEABRE() []antlr.TerminalNode
	LLAVEABRE(i int) antlr.TerminalNode
	AllLLAVECIERRA() []antlr.TerminalNode
	LLAVECIERRA(i int) antlr.TerminalNode
	Tipo() ITipoContext
	IDENTIFICADOR() antlr.TerminalNode
	Lista_arreglos() ILista_arreglosContext
	PUNTO_COMA() antlr.TerminalNode

	// IsDeclaracion_matrizContext differentiates from other interfaces.
	IsDeclaracion_matrizContext()
}

type Declaracion_matrizContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	idMatriz antlr.Token
}

func NewEmptyDeclaracion_matrizContext() *Declaracion_matrizContext {
	var p = new(Declaracion_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_matriz
	return p
}

func InitEmptyDeclaracion_matrizContext(p *Declaracion_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracion_matriz
}

func (*Declaracion_matrizContext) IsDeclaracion_matrizContext() {}

func NewDeclaracion_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_matrizContext {
	var p = new(Declaracion_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracion_matriz

	return p
}

func (s *Declaracion_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_matrizContext) GetIdMatriz() antlr.Token { return s.idMatriz }

func (s *Declaracion_matrizContext) SetIdMatriz(v antlr.Token) { s.idMatriz = v }

func (s *Declaracion_matrizContext) ASIGNACION_INFERIDA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserASIGNACION_INFERIDA, 0)
}

func (s *Declaracion_matrizContext) AllLLAVEABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserLLAVEABRE)
}

func (s *Declaracion_matrizContext) LLAVEABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, i)
}

func (s *Declaracion_matrizContext) AllLLAVECIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserLLAVECIERRA)
}

func (s *Declaracion_matrizContext) LLAVECIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, i)
}

func (s *Declaracion_matrizContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_matrizContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Declaracion_matrizContext) Lista_arreglos() ILista_arreglosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_arreglosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_arreglosContext)
}

func (s *Declaracion_matrizContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Declaracion_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_matrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracion_matriz(s)
	}
}

func (s *Declaracion_matrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracion_matriz(s)
	}
}

func (s *Declaracion_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracion_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Declaracion_matriz() (localctx IDeclaracion_matrizContext) {
	localctx = NewDeclaracion_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, gramaticaParserRULE_declaracion_matriz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(599)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Declaracion_matrizContext).idMatriz = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(600)
		p.Match(gramaticaParserASIGNACION_INFERIDA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(601)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(602)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(603)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(604)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(605)
		p.Tipo()
	}
	p.SetState(607)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(606)
			p.Lista_arreglos()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(610)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(609)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_arreglosContext is an interface to support dynamic dispatch.
type ILista_arreglosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAVEABRE() antlr.TerminalNode
	LLAVECIERRA() antlr.TerminalNode
	AllLista_arreglos() []ILista_arreglosContext
	Lista_arreglos(i int) ILista_arreglosContext
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsLista_arreglosContext differentiates from other interfaces.
	IsLista_arreglosContext()
}

type Lista_arreglosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_arreglosContext() *Lista_arreglosContext {
	var p = new(Lista_arreglosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_arreglos
	return p
}

func InitEmptyLista_arreglosContext(p *Lista_arreglosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_lista_arreglos
}

func (*Lista_arreglosContext) IsLista_arreglosContext() {}

func NewLista_arreglosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_arreglosContext {
	var p = new(Lista_arreglosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_lista_arreglos

	return p
}

func (s *Lista_arreglosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_arreglosContext) LLAVEABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVEABRE, 0)
}

func (s *Lista_arreglosContext) LLAVECIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAVECIERRA, 0)
}

func (s *Lista_arreglosContext) AllLista_arreglos() []ILista_arreglosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_arreglosContext); ok {
			len++
		}
	}

	tst := make([]ILista_arreglosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_arreglosContext); ok {
			tst[i] = t.(ILista_arreglosContext)
			i++
		}
	}

	return tst
}

func (s *Lista_arreglosContext) Lista_arreglos(i int) ILista_arreglosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_arreglosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_arreglosContext)
}

func (s *Lista_arreglosContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Lista_arreglosContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_arreglosContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *Lista_arreglosContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *Lista_arreglosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_arreglosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_arreglosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLista_arreglos(s)
	}
}

func (s *Lista_arreglosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLista_arreglos(s)
	}
}

func (s *Lista_arreglosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLista_arreglos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Lista_arreglos() (localctx ILista_arreglosContext) {
	localctx = NewLista_arreglosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, gramaticaParserRULE_lista_arreglos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(612)
		p.Match(gramaticaParserLLAVEABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserLLAVEABRE:
		{
			p.SetState(613)
			p.Lista_arreglos()
		}

	case gramaticaParserT__0, gramaticaParserT__1, gramaticaParserRESINDEX, gramaticaParserRESJOIN, gramaticaParserRESLEN, gramaticaParserRESAPPEND, gramaticaParserRESNIL, gramaticaParserIDENTIFICADOR, gramaticaParserINT, gramaticaParserFLOAT, gramaticaParserCADENA, gramaticaParserRUNE, gramaticaParserPARABRE, gramaticaParserCORABRE, gramaticaParserSIGNO_MENOS, gramaticaParserSIGNO_NOT:
		{
			p.SetState(614)
			p.expresion(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(617)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(620)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case gramaticaParserLLAVEABRE:
			{
				p.SetState(618)
				p.Lista_arreglos()
			}

		case gramaticaParserT__0, gramaticaParserT__1, gramaticaParserRESINDEX, gramaticaParserRESJOIN, gramaticaParserRESLEN, gramaticaParserRESAPPEND, gramaticaParserRESNIL, gramaticaParserIDENTIFICADOR, gramaticaParserINT, gramaticaParserFLOAT, gramaticaParserCADENA, gramaticaParserRUNE, gramaticaParserPARABRE, gramaticaParserCORABRE, gramaticaParserSIGNO_MENOS, gramaticaParserSIGNO_NOT:
			{
				p.SetState(619)
				p.expresion(0)
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(626)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(627)
		p.Match(gramaticaParserLLAVECIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacion_acceso_matrizContext is an interface to support dynamic dispatch.
type IAsignacion_acceso_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdMatriz returns the idMatriz token.
	GetIdMatriz() antlr.Token

	// SetIdMatriz sets the idMatriz token.
	SetIdMatriz(antlr.Token)

	// GetIndex1 returns the index1 rule contexts.
	GetIndex1() IExpresionContext

	// GetIndex2 returns the index2 rule contexts.
	GetIndex2() IExpresionContext

	// SetIndex1 sets the index1 rule contexts.
	SetIndex1(IExpresionContext)

	// SetIndex2 sets the index2 rule contexts.
	SetIndex2(IExpresionContext)

	// Getter signatures
	AllCORABRE() []antlr.TerminalNode
	CORABRE(i int) antlr.TerminalNode
	AllCORCIERRA() []antlr.TerminalNode
	CORCIERRA(i int) antlr.TerminalNode
	SIGNO_ASIGNACION() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	IDENTIFICADOR() antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsAsignacion_acceso_matrizContext differentiates from other interfaces.
	IsAsignacion_acceso_matrizContext()
}

type Asignacion_acceso_matrizContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	idMatriz antlr.Token
	index1   IExpresionContext
	index2   IExpresionContext
}

func NewEmptyAsignacion_acceso_matrizContext() *Asignacion_acceso_matrizContext {
	var p = new(Asignacion_acceso_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_matriz
	return p
}

func InitEmptyAsignacion_acceso_matrizContext(p *Asignacion_acceso_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_matriz
}

func (*Asignacion_acceso_matrizContext) IsAsignacion_acceso_matrizContext() {}

func NewAsignacion_acceso_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_acceso_matrizContext {
	var p = new(Asignacion_acceso_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_matriz

	return p
}

func (s *Asignacion_acceso_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_acceso_matrizContext) GetIdMatriz() antlr.Token { return s.idMatriz }

func (s *Asignacion_acceso_matrizContext) SetIdMatriz(v antlr.Token) { s.idMatriz = v }

func (s *Asignacion_acceso_matrizContext) GetIndex1() IExpresionContext { return s.index1 }

func (s *Asignacion_acceso_matrizContext) GetIndex2() IExpresionContext { return s.index2 }

func (s *Asignacion_acceso_matrizContext) SetIndex1(v IExpresionContext) { s.index1 = v }

func (s *Asignacion_acceso_matrizContext) SetIndex2(v IExpresionContext) { s.index2 = v }

func (s *Asignacion_acceso_matrizContext) AllCORABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORABRE)
}

func (s *Asignacion_acceso_matrizContext) CORABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, i)
}

func (s *Asignacion_acceso_matrizContext) AllCORCIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORCIERRA)
}

func (s *Asignacion_acceso_matrizContext) CORCIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, i)
}

func (s *Asignacion_acceso_matrizContext) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *Asignacion_acceso_matrizContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_acceso_matrizContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_acceso_matrizContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *Asignacion_acceso_matrizContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Asignacion_acceso_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_acceso_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asignacion_acceso_matrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacion_acceso_matriz(s)
	}
}

func (s *Asignacion_acceso_matrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacion_acceso_matriz(s)
	}
}

func (s *Asignacion_acceso_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacion_acceso_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Asignacion_acceso_matriz() (localctx IAsignacion_acceso_matrizContext) {
	localctx = NewAsignacion_acceso_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, gramaticaParserRULE_asignacion_acceso_matriz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Asignacion_acceso_matrizContext).idMatriz = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(630)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(631)

		var _x = p.expresion(0)

		localctx.(*Asignacion_acceso_matrizContext).index1 = _x
	}
	{
		p.SetState(632)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(633)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(634)

		var _x = p.expresion(0)

		localctx.(*Asignacion_acceso_matrizContext).index2 = _x
	}
	{
		p.SetState(635)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(636)
		p.Match(gramaticaParserSIGNO_ASIGNACION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(637)
		p.expresion(0)
	}
	p.SetState(639)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(638)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccesoMatrizContext is an interface to support dynamic dispatch.
type IAccesoMatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdMatriz returns the idMatriz token.
	GetIdMatriz() antlr.Token

	// SetIdMatriz sets the idMatriz token.
	SetIdMatriz(antlr.Token)

	// GetIndex1 returns the index1 rule contexts.
	GetIndex1() IExpresionContext

	// GetIndex2 returns the index2 rule contexts.
	GetIndex2() IExpresionContext

	// SetIndex1 sets the index1 rule contexts.
	SetIndex1(IExpresionContext)

	// SetIndex2 sets the index2 rule contexts.
	SetIndex2(IExpresionContext)

	// Getter signatures
	AllCORABRE() []antlr.TerminalNode
	CORABRE(i int) antlr.TerminalNode
	AllCORCIERRA() []antlr.TerminalNode
	CORCIERRA(i int) antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsAccesoMatrizContext differentiates from other interfaces.
	IsAccesoMatrizContext()
}

type AccesoMatrizContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	idMatriz antlr.Token
	index1   IExpresionContext
	index2   IExpresionContext
}

func NewEmptyAccesoMatrizContext() *AccesoMatrizContext {
	var p = new(AccesoMatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoMatriz
	return p
}

func InitEmptyAccesoMatrizContext(p *AccesoMatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoMatriz
}

func (*AccesoMatrizContext) IsAccesoMatrizContext() {}

func NewAccesoMatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoMatrizContext {
	var p = new(AccesoMatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_accesoMatriz

	return p
}

func (s *AccesoMatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoMatrizContext) GetIdMatriz() antlr.Token { return s.idMatriz }

func (s *AccesoMatrizContext) SetIdMatriz(v antlr.Token) { s.idMatriz = v }

func (s *AccesoMatrizContext) GetIndex1() IExpresionContext { return s.index1 }

func (s *AccesoMatrizContext) GetIndex2() IExpresionContext { return s.index2 }

func (s *AccesoMatrizContext) SetIndex1(v IExpresionContext) { s.index1 = v }

func (s *AccesoMatrizContext) SetIndex2(v IExpresionContext) { s.index2 = v }

func (s *AccesoMatrizContext) AllCORABRE() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORABRE)
}

func (s *AccesoMatrizContext) CORABRE(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORABRE, i)
}

func (s *AccesoMatrizContext) AllCORCIERRA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCORCIERRA)
}

func (s *AccesoMatrizContext) CORCIERRA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCORCIERRA, i)
}

func (s *AccesoMatrizContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AccesoMatrizContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *AccesoMatrizContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AccesoMatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoMatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoMatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAccesoMatriz(s)
	}
}

func (s *AccesoMatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAccesoMatriz(s)
	}
}

func (s *AccesoMatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAccesoMatriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AccesoMatriz() (localctx IAccesoMatrizContext) {
	localctx = NewAccesoMatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, gramaticaParserRULE_accesoMatriz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*AccesoMatrizContext).idMatriz = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(642)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(643)

		var _x = p.expresion(0)

		localctx.(*AccesoMatrizContext).index1 = _x
	}
	{
		p.SetState(644)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(645)
		p.Match(gramaticaParserCORABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(646)

		var _x = p.expresion(0)

		localctx.(*AccesoMatrizContext).index2 = _x
	}
	{
		p.SetState(647)
		p.Match(gramaticaParserCORCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaParseFloatContext is an interface to support dynamic dispatch.
type ILlamadaParseFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESPARSEFLOAT() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	Expresion() IExpresionContext
	PARCIERRA() antlr.TerminalNode

	// IsLlamadaParseFloatContext differentiates from other interfaces.
	IsLlamadaParseFloatContext()
}

type LlamadaParseFloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaParseFloatContext() *LlamadaParseFloatContext {
	var p = new(LlamadaParseFloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaParseFloat
	return p
}

func InitEmptyLlamadaParseFloatContext(p *LlamadaParseFloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaParseFloat
}

func (*LlamadaParseFloatContext) IsLlamadaParseFloatContext() {}

func NewLlamadaParseFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaParseFloatContext {
	var p = new(LlamadaParseFloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaParseFloat

	return p
}

func (s *LlamadaParseFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaParseFloatContext) RESPARSEFLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESPARSEFLOAT, 0)
}

func (s *LlamadaParseFloatContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaParseFloatContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaParseFloatContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaParseFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaParseFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaParseFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaParseFloat(s)
	}
}

func (s *LlamadaParseFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaParseFloat(s)
	}
}

func (s *LlamadaParseFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaParseFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaParseFloat() (localctx ILlamadaParseFloatContext) {
	localctx = NewLlamadaParseFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, gramaticaParserRULE_llamadaParseFloat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(649)
		p.Match(gramaticaParserRESPARSEFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(650)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(651)
		p.expresion(0)
	}
	{
		p.SetState(652)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaTypeOfContext is an interface to support dynamic dispatch.
type ILlamadaTypeOfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESTYPEOF() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	Expresion() IExpresionContext
	PARCIERRA() antlr.TerminalNode

	// IsLlamadaTypeOfContext differentiates from other interfaces.
	IsLlamadaTypeOfContext()
}

type LlamadaTypeOfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaTypeOfContext() *LlamadaTypeOfContext {
	var p = new(LlamadaTypeOfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaTypeOf
	return p
}

func InitEmptyLlamadaTypeOfContext(p *LlamadaTypeOfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaTypeOf
}

func (*LlamadaTypeOfContext) IsLlamadaTypeOfContext() {}

func NewLlamadaTypeOfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaTypeOfContext {
	var p = new(LlamadaTypeOfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaTypeOf

	return p
}

func (s *LlamadaTypeOfContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaTypeOfContext) RESTYPEOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRESTYPEOF, 0)
}

func (s *LlamadaTypeOfContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *LlamadaTypeOfContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaTypeOfContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *LlamadaTypeOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaTypeOfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaTypeOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaTypeOf(s)
	}
}

func (s *LlamadaTypeOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaTypeOf(s)
	}
}

func (s *LlamadaTypeOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaTypeOf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaTypeOf() (localctx ILlamadaTypeOfContext) {
	localctx = NewLlamadaTypeOfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, gramaticaParserRULE_llamadaTypeOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(654)
		p.Match(gramaticaParserRESTYPEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(655)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(656)
		p.expresion(0)
	}
	{
		p.SetState(657)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaCamposContext is an interface to support dynamic dispatch.
type IListaCamposContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCampoStruct() []ICampoStructContext
	CampoStruct(i int) ICampoStructContext
	AllSIGNO_COMA() []antlr.TerminalNode
	SIGNO_COMA(i int) antlr.TerminalNode

	// IsListaCamposContext differentiates from other interfaces.
	IsListaCamposContext()
}

type ListaCamposContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaCamposContext() *ListaCamposContext {
	var p = new(ListaCamposContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaCampos
	return p
}

func InitEmptyListaCamposContext(p *ListaCamposContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaCampos
}

func (*ListaCamposContext) IsListaCamposContext() {}

func NewListaCamposContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaCamposContext {
	var p = new(ListaCamposContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaCampos

	return p
}

func (s *ListaCamposContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaCamposContext) AllCampoStruct() []ICampoStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICampoStructContext); ok {
			len++
		}
	}

	tst := make([]ICampoStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICampoStructContext); ok {
			tst[i] = t.(ICampoStructContext)
			i++
		}
	}

	return tst
}

func (s *ListaCamposContext) CampoStruct(i int) ICampoStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICampoStructContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICampoStructContext)
}

func (s *ListaCamposContext) AllSIGNO_COMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserSIGNO_COMA)
}

func (s *ListaCamposContext) SIGNO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_COMA, i)
}

func (s *ListaCamposContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaCamposContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaCamposContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaCampos(s)
	}
}

func (s *ListaCamposContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaCampos(s)
	}
}

func (s *ListaCamposContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaCampos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaCampos() (localctx IListaCamposContext) {
	localctx = NewListaCamposContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, gramaticaParserRULE_listaCampos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(659)
		p.CampoStruct()
	}
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserSIGNO_COMA {
		{
			p.SetState(660)
			p.Match(gramaticaParserSIGNO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(661)
			p.CampoStruct()
		}

		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICampoStructContext is an interface to support dynamic dispatch.
type ICampoStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdCampo returns the idCampo token.
	GetIdCampo() antlr.Token

	// SetIdCampo sets the idCampo token.
	SetIdCampo(antlr.Token)

	// Getter signatures
	DOS_PUNTOS() antlr.TerminalNode
	Expresion() IExpresionContext
	IDENTIFICADOR() antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsCampoStructContext differentiates from other interfaces.
	IsCampoStructContext()
}

type CampoStructContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	idCampo antlr.Token
}

func NewEmptyCampoStructContext() *CampoStructContext {
	var p = new(CampoStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_campoStruct
	return p
}

func InitEmptyCampoStructContext(p *CampoStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_campoStruct
}

func (*CampoStructContext) IsCampoStructContext() {}

func NewCampoStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CampoStructContext {
	var p = new(CampoStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_campoStruct

	return p
}

func (s *CampoStructContext) GetParser() antlr.Parser { return s.parser }

func (s *CampoStructContext) GetIdCampo() antlr.Token { return s.idCampo }

func (s *CampoStructContext) SetIdCampo(v antlr.Token) { s.idCampo = v }

func (s *CampoStructContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOS_PUNTOS, 0)
}

func (s *CampoStructContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CampoStructContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *CampoStructContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *CampoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CampoStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CampoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCampoStruct(s)
	}
}

func (s *CampoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCampoStruct(s)
	}
}

func (s *CampoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCampoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) CampoStruct() (localctx ICampoStructContext) {
	localctx = NewCampoStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, gramaticaParserRULE_campoStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(667)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*CampoStructContext).idCampo = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(668)
		p.Match(gramaticaParserDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(669)
		p.expresion(0)
	}
	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserPUNTO_COMA {
		{
			p.SetState(670)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccesoInstaciaAtributoContext is an interface to support dynamic dispatch.
type IAccesoInstaciaAtributoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomInstancia returns the nomInstancia token.
	GetNomInstancia() antlr.Token

	// GetNomAtributo returns the nomAtributo token.
	GetNomAtributo() antlr.Token

	// SetNomInstancia sets the nomInstancia token.
	SetNomInstancia(antlr.Token)

	// SetNomAtributo sets the nomAtributo token.
	SetNomAtributo(antlr.Token)

	// Getter signatures
	PUNTO() antlr.TerminalNode
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode

	// IsAccesoInstaciaAtributoContext differentiates from other interfaces.
	IsAccesoInstaciaAtributoContext()
}

type AccesoInstaciaAtributoContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	nomInstancia antlr.Token
	nomAtributo  antlr.Token
}

func NewEmptyAccesoInstaciaAtributoContext() *AccesoInstaciaAtributoContext {
	var p = new(AccesoInstaciaAtributoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoInstaciaAtributo
	return p
}

func InitEmptyAccesoInstaciaAtributoContext(p *AccesoInstaciaAtributoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoInstaciaAtributo
}

func (*AccesoInstaciaAtributoContext) IsAccesoInstaciaAtributoContext() {}

func NewAccesoInstaciaAtributoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoInstaciaAtributoContext {
	var p = new(AccesoInstaciaAtributoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_accesoInstaciaAtributo

	return p
}

func (s *AccesoInstaciaAtributoContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoInstaciaAtributoContext) GetNomInstancia() antlr.Token { return s.nomInstancia }

func (s *AccesoInstaciaAtributoContext) GetNomAtributo() antlr.Token { return s.nomAtributo }

func (s *AccesoInstaciaAtributoContext) SetNomInstancia(v antlr.Token) { s.nomInstancia = v }

func (s *AccesoInstaciaAtributoContext) SetNomAtributo(v antlr.Token) { s.nomAtributo = v }

func (s *AccesoInstaciaAtributoContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO, 0)
}

func (s *AccesoInstaciaAtributoContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *AccesoInstaciaAtributoContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *AccesoInstaciaAtributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoInstaciaAtributoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoInstaciaAtributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAccesoInstaciaAtributo(s)
	}
}

func (s *AccesoInstaciaAtributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAccesoInstaciaAtributo(s)
	}
}

func (s *AccesoInstaciaAtributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAccesoInstaciaAtributo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AccesoInstaciaAtributo() (localctx IAccesoInstaciaAtributoContext) {
	localctx = NewAccesoInstaciaAtributoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, gramaticaParserRULE_accesoInstaciaAtributo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(673)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*AccesoInstaciaAtributoContext).nomInstancia = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(674)
		p.Match(gramaticaParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(675)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*AccesoInstaciaAtributoContext).nomAtributo = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacion_acceso_atributoContext is an interface to support dynamic dispatch.
type IAsignacion_acceso_atributoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomInstancia returns the nomInstancia token.
	GetNomInstancia() antlr.Token

	// GetNomAtributo returns the nomAtributo token.
	GetNomAtributo() antlr.Token

	// SetNomInstancia sets the nomInstancia token.
	SetNomInstancia(antlr.Token)

	// SetNomAtributo sets the nomAtributo token.
	SetNomAtributo(antlr.Token)

	// Getter signatures
	PUNTO() antlr.TerminalNode
	SIGNO_ASIGNACION() antlr.TerminalNode
	Expresion() IExpresionContext
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	PUNTO_COMA() antlr.TerminalNode

	// IsAsignacion_acceso_atributoContext differentiates from other interfaces.
	IsAsignacion_acceso_atributoContext()
}

type Asignacion_acceso_atributoContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	nomInstancia antlr.Token
	nomAtributo  antlr.Token
}

func NewEmptyAsignacion_acceso_atributoContext() *Asignacion_acceso_atributoContext {
	var p = new(Asignacion_acceso_atributoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_atributo
	return p
}

func InitEmptyAsignacion_acceso_atributoContext(p *Asignacion_acceso_atributoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_atributo
}

func (*Asignacion_acceso_atributoContext) IsAsignacion_acceso_atributoContext() {}

func NewAsignacion_acceso_atributoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_acceso_atributoContext {
	var p = new(Asignacion_acceso_atributoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion_acceso_atributo

	return p
}

func (s *Asignacion_acceso_atributoContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_acceso_atributoContext) GetNomInstancia() antlr.Token { return s.nomInstancia }

func (s *Asignacion_acceso_atributoContext) GetNomAtributo() antlr.Token { return s.nomAtributo }

func (s *Asignacion_acceso_atributoContext) SetNomInstancia(v antlr.Token) { s.nomInstancia = v }

func (s *Asignacion_acceso_atributoContext) SetNomAtributo(v antlr.Token) { s.nomAtributo = v }

func (s *Asignacion_acceso_atributoContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO, 0)
}

func (s *Asignacion_acceso_atributoContext) SIGNO_ASIGNACION() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSIGNO_ASIGNACION, 0)
}

func (s *Asignacion_acceso_atributoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_acceso_atributoContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *Asignacion_acceso_atributoContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *Asignacion_acceso_atributoContext) PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO_COMA, 0)
}

func (s *Asignacion_acceso_atributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_acceso_atributoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asignacion_acceso_atributoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacion_acceso_atributo(s)
	}
}

func (s *Asignacion_acceso_atributoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacion_acceso_atributo(s)
	}
}

func (s *Asignacion_acceso_atributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacion_acceso_atributo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Asignacion_acceso_atributo() (localctx IAsignacion_acceso_atributoContext) {
	localctx = NewAsignacion_acceso_atributoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, gramaticaParserRULE_asignacion_acceso_atributo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(677)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Asignacion_acceso_atributoContext).nomInstancia = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(678)
		p.Match(gramaticaParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(679)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Asignacion_acceso_atributoContext).nomAtributo = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(680)
		p.Match(gramaticaParserSIGNO_ASIGNACION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(681)
		p.expresion(0)
	}
	p.SetState(683)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(682)
			p.Match(gramaticaParserPUNTO_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamada_funcion_structContext is an interface to support dynamic dispatch.
type ILlamada_funcion_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNomInstancia returns the nomInstancia token.
	GetNomInstancia() antlr.Token

	// GetNomFunc returns the nomFunc token.
	GetNomFunc() antlr.Token

	// SetNomInstancia sets the nomInstancia token.
	SetNomInstancia(antlr.Token)

	// SetNomFunc sets the nomFunc token.
	SetNomFunc(antlr.Token)

	// Getter signatures
	PUNTO() antlr.TerminalNode
	PARABRE() antlr.TerminalNode
	PARCIERRA() antlr.TerminalNode
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	Lista_parametros_func() ILista_parametros_funcContext

	// IsLlamada_funcion_structContext differentiates from other interfaces.
	IsLlamada_funcion_structContext()
}

type Llamada_funcion_structContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	nomInstancia antlr.Token
	nomFunc      antlr.Token
}

func NewEmptyLlamada_funcion_structContext() *Llamada_funcion_structContext {
	var p = new(Llamada_funcion_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamada_funcion_struct
	return p
}

func InitEmptyLlamada_funcion_structContext(p *Llamada_funcion_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamada_funcion_struct
}

func (*Llamada_funcion_structContext) IsLlamada_funcion_structContext() {}

func NewLlamada_funcion_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_funcion_structContext {
	var p = new(Llamada_funcion_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamada_funcion_struct

	return p
}

func (s *Llamada_funcion_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_funcion_structContext) GetNomInstancia() antlr.Token { return s.nomInstancia }

func (s *Llamada_funcion_structContext) GetNomFunc() antlr.Token { return s.nomFunc }

func (s *Llamada_funcion_structContext) SetNomInstancia(v antlr.Token) { s.nomInstancia = v }

func (s *Llamada_funcion_structContext) SetNomFunc(v antlr.Token) { s.nomFunc = v }

func (s *Llamada_funcion_structContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPUNTO, 0)
}

func (s *Llamada_funcion_structContext) PARABRE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARABRE, 0)
}

func (s *Llamada_funcion_structContext) PARCIERRA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARCIERRA, 0)
}

func (s *Llamada_funcion_structContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *Llamada_funcion_structContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *Llamada_funcion_structContext) Lista_parametros_func() ILista_parametros_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametros_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametros_funcContext)
}

func (s *Llamada_funcion_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_funcion_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_funcion_structContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamada_funcion_struct(s)
	}
}

func (s *Llamada_funcion_structContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamada_funcion_struct(s)
	}
}

func (s *Llamada_funcion_structContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamada_funcion_struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Llamada_funcion_struct() (localctx ILlamada_funcion_structContext) {
	localctx = NewLlamada_funcion_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, gramaticaParserRULE_llamada_funcion_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(685)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Llamada_funcion_structContext).nomInstancia = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(686)
		p.Match(gramaticaParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(687)

		var _m = p.Match(gramaticaParserIDENTIFICADOR)

		localctx.(*Llamada_funcion_structContext).nomFunc = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(688)
		p.Match(gramaticaParserPARABRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4900068058465779718) != 0 {
		{
			p.SetState(689)
			p.Lista_parametros_func()
		}

	}
	{
		p.SetState(692)
		p.Match(gramaticaParserPARCIERRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 20)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
