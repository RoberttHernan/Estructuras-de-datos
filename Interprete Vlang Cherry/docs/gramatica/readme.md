```g4

grammar gramatica;

//// antes de ejecutar el main siempre ejecutar el comando;                    ////
////    antlr4 -Dlanguage=Go -visitor -package parser -o parser gramatica.g4   ////


//--palabras reservadas-------------------

RESFUNC : 'func'; // 'fn'
RESIMPORT : 'import' ;


RESVAR : 'var';
RESMUT: 'mut' ;
RESINT : 'int';
RESFLOAT : 'float64';
RESSTRING : 'string';

RESBOOL :   'bool';
RESRUNE : 'rune';
RESSLICE : 'slices';
RESINDEX : 'indexOf'; //'index'
RESSTRINGS : 'strings';
RESJOIN: 'join'; //'Join'
RESLEN : 'len';
RESAPPEND: 'append';
RESSTRUCT : 'struct';
// RESVOID :   'void'; no se usa void
RESNIL : 'nil'

;

RESIF : 'if';
RESELSE : 'else';
RESFOR : 'for';
RESSWITCH : 'switch';
RESCASE : 'case';
RESDEFAULT : 'default';
// RESRANGE : 'range'; no se usa
RESRETURN : 'return';
RESBREAK : 'break';
RESCONTINUE : 'continue';
RESIN : 'in';

RESFMT : 'fmt';
RESPRINTLN : 'print';  // 'Println' confirme con el aux que es solo 'print'
// RESSTR: 'strconv'; no se usa 
RESATOI : 'Atoi';
RESPARSEFLOAT : 'ParseFloat';
// RESREFLECT : 'reflect'; no se usa
RESTYPEOF : 'typeOf'; // 'Typeof'



//-------tokens para lexer----------------

IDENTIFICADOR : ('_')?[a-zA-Z_][a-zA-Z_0-9]*;
INT :  [0-9]+;
FLOAT : [0-9]+'.'[0-9]+;
CADENA : '"' (~["\r\n])* '"';
RUNE :  '\'' . '\'';


ESPACIO : [ \t\r\n]+ -> skip;
COMENTARIO : ('//' ~[\r\n]* 
            | '/*' .*? '*/') 
            -> skip;





//-----tokens para signos-------------------
PARABRE : '(';
PARCIERRA : ')';
LLAVEABRE : '{';
LLAVECIERRA : '}';
CORABRE :'[';
CORCIERRA :']';
PUNTO_COMA : ';';
DOS_PUNTOS : ':';
SIGNO_MULTI : '*';
SIGNO_DIV : '/';
SIGNO_ASIGNACION_INCREMENTO : '+=';
SIGNO_ASIGNACION_DECREMENTO : '-=';
SIGNO_INCREMENTO : '++';
SIGNO_DECREMENTO : '--';
SIGNO_MAS : '+';
SIGNO_MENOS : '-';
SIGNO_AND : '&&';
SIGNO_OR : '||';
SIGNO_NO_IGUAL : '!=';
SIGNO_NOT : '!';
SIGNO_IGUALDAD : '==';
SIGNO_ASIGNACION : '=';
SIGNO_MENOSQUE_IGUAL: '<=';
SIGNO_MENOSQUE : '<';
SIGNO_MASQUE_IGUAL : '>=';
SIGNO_MASQUE : '>';
SIGNO_COMA : ',';
SIGNO_MODULO : '%';
ASIGNACION_INFERIDA :  ':='; 
PUNTO : '.';







//--inicio del programa

// prog : importStmt* (declaracionStruct|funcion_declaracion| bloque| sentencia) EOF;
prog :  (declaracionStruct | funcionStruct_declaracion | funcion_declaracion | bloque)+ EOF;

//declaracionStruct : RESSTRUCT IDENTIFICADOR LLAVEABRE atributosStruct+ LLAVECIERRA ;
declaracionStruct : RESSTRUCT nomStruct= IDENTIFICADOR LLAVEABRE atributosStruct+ LLAVECIERRA ;

funcionStruct_declaracion : 
        RESFUNC PARABRE  id= IDENTIFICADOR nomStruct= IDENTIFICADOR PARCIERRA nomFunc= IDENTIFICADOR PARABRE lista_parametros? PARCIERRA tipo? bloque;

funcion_declaracion : 
        RESFUNC nomFunc= IDENTIFICADOR PARABRE lista_parametros? PARCIERRA tipoRetorno= tipo? bloque                              #DecFunc1
      | RESFUNC nomFunc= IDENTIFICADOR PARABRE lista_parametros? PARCIERRA (LLAVEABRE LLAVECIERRA)* tipoRetorno= tipo bloque      #DecFunc2
; 

lista_parametros : parametro (SIGNO_COMA parametro)* ;
parametro : idParametro= IDENTIFICADOR (SIGNO_MULTI | LLAVEABRE LLAVECIERRA)? tipo; // tipo IDENTIFICADOR


lista_sentencias : sentencia (sentencia)*;

sentencia: //bloque
         declaracion_variable
        |asignacion
        |declaracion_slice // asignacion_slice
        |asignacion_acceso_slice
        |sentenciaIf
        |sentenciaSwitch
        |sentencia_transferencia
        |llamadafuncion 
        |sentenciaPrintln
        |creacion_instancia_struct
        |sentenciaFor
        |declaracion_matriz
        |asignacion_acceso_matriz
        |sentenciaReturn
        |sentenciaBreak
        ; 
sentenciaBreak : RESBREAK PUNTO_COMA? ;

sentenciaReturn : RESRETURN expresion PUNTO_COMA? ; 

llamadafuncion : IDENTIFICADOR PARABRE lista_expresion? PARCIERRA PUNTO_COMA ;

sentencia_transferencia : RESBREAK PUNTO_COMA?                  #SentBreak
                        | RESCONTINUE PUNTO_COMA?               #SentContinue
; 



// sentenciaIf :
//         RESIF expresion bloque (RESELSE RESIF expresion bloque)* (RESELSE bloque)?
// ;
sentenciaIf: RESIF PARABRE expresion PARCIERRA   bloque                             #SentIf1
           | RESIF PARABRE expresion PARCIERRA bloque RESELSE bloque              #SentIf2
           | RESIF PARABRE expresion PARCIERRA bloque RESELSE sentenciaIf         #SentIf3
;

//sentenciaSwitch : RESSWITCH expresion LLAVEABRE (bloqueCase)* bloqueDefault? LLAVECIERRA;
//bloqueCase : RESCASE expresion bloque;
//bloqueDefault : RESDEFAULT bloque;
sentenciaSwitch : RESSWITCH expresion LLAVEABRE (bloqueCase)+ bloqueDefault? LLAVECIERRA;
bloqueCase : RESCASE expresion DOS_PUNTOS bloque;
bloqueDefault : RESDEFAULT DOS_PUNTOS lista_sentencias? ;

//sentenciaPrintln : RESPRINTLN PARABRE lista_parametros? PARCIERRA PUNTO_COMA  ;
sentenciaPrintln : RESPRINTLN PARABRE lista_expresion? PARCIERRA PUNTO_COMA?  ;





// asignacion : IDENTIFICADOR SIGNO_ASIGNACION expresion PUNTO_COMA ; 
asignacion : id= IDENTIFICADOR SIGNO_ASIGNACION expresion PUNTO_COMA?                                                       #AsigSimple
           | id= IDENTIFICADOR op= (SIGNO_ASIGNACION_INCREMENTO | SIGNO_ASIGNACION_DECREMENTO) expresion PUNTO_COMA?        #AsigSumRest
;

declaracion_slice : RESMUT nomSlice= IDENTIFICADOR CORABRE CORCIERRA tipo PUNTO_COMA?                                                     #DecSlice1 //vacio
                | nomSlice= IDENTIFICADOR SIGNO_IGUALDAD  CORABRE CORCIERRA tipo LLAVEABRE lista_expresion LLAVECIERRA PUNTO_COMA?        #DecSlice2 //con valores
;

// asignacion_slice : accesoSlice SIGNO_ASIGNACION inicializacion_slice PUNTO_COMA ;

// asignacion_acceso_slice : accesoSlice SIGNO_ASIGNACION expresion PUNTO_COMA ; 
asignacion_acceso_slice : nomSlice= IDENTIFICADOR (CORABRE index= expresion CORCIERRA) SIGNO_ASIGNACION expresion PUNTO_COMA? ; 


// accesoSlice : IDENTIFICADOR (CORABRE expresion CORCIERRA)+ ;

/*
Es posible declarar variables de la siguiente manera:
        mut IDENTIFICADOR tipo = expresion PUNTO_COMA? ---> Dec1
        mut IDENTIFICADOR tipo PUNTO_COMA? ---> Dec1
        mut? IDENTIFICADOR := expresion PUNTO_COMA?  ---> Dec2 confirmado por el aux que el mut puede o no venir
        IDENTIFICADOR tipo = expresion PUNTO_COMA? --> Dec3
*/

// declaracion_variable: RESMUT IDENTIFICADOR tipo? (SIGNO_ASIGNACION expresion)? PUNTO_COMA+ 
declaracion_variable: RESMUT id= IDENTIFICADOR tipo (SIGNO_ASIGNACION expresion) PUNTO_COMA?      #Dec1
                     |RESMUT id= IDENTIFICADOR ASIGNACION_INFERIDA expresion PUNTO_COMA?          #Dec2
                    | id= IDENTIFICADOR tipo SIGNO_ASIGNACION expresion PUNTO_COMA?                  #Dec3
                    ;  

// expresion: expresion op=(SIGNO_MULTI | SIGNO_DIV | SIGNO_MODULO) expresion
//         |expresion op= (SIGNO_MAS | SIGNO_MENOS) expresion
//         | expresion op= (SIGNO_IGUALDAD| SIGNO_NO_IGUAL | SIGNO_MENOSQUE | SIGNO_MENOSQUE_IGUAL | SIGNO_MASQUE | SIGNO_MASQUE_IGUAL) expresion
//         |expresion op= (SIGNO_AND | SIGNO_OR) expresion
//         |SIGNO_NOT expresion
//         |SIGNO_MENOS expresion
//         |PARABRE expresion PARCIERRA
//         |CADENA
//         |IDENTIFICADOR
//         |literal
//         |inicializacion_slice

// ; 
expresion: op= SIGNO_MENOS PARABRE expresion PARCIERRA                                       #ExprNegUnaria
         | left= expresion op=(SIGNO_MULTI | SIGNO_DIV | SIGNO_MODULO) right= expresion      #ExprMultDivMod
         | left= expresion op=(SIGNO_MAS | SIGNO_MENOS) right= expresion                     #ExprOpAritmetica
         | PARABRE expresion PARCIERRA                                                       #ExprParentesis
         | left= expresion op= (SIGNO_IGUALDAD| SIGNO_NO_IGUAL) right= expresion             #ExprRelacional
         | left= expresion op= (SIGNO_MENOSQUE | SIGNO_MENOSQUE_IGUAL) right= expresion      #ExprRelacional
         | left= expresion op= (SIGNO_MASQUE | SIGNO_MASQUE_IGUAL) right= expresion          #ExprRelacional
         | op = SIGNO_NOT expresion                                                          #ExprNot
         | left= expresion op= (SIGNO_AND | SIGNO_OR) right= expresion                       #ExprLogica
         | INT                                                                               #ExprInt
         | FLOAT                                                                             #ExprFloat
         | valor= 'true'                                                                     #ExprBoolTrue    
         | valor= 'false'                                                                    #ExprBoolFalse
         | CADENA                                                                            #ExprString
         | RUNE                                                                              #ExprRune
         | RESNIL                                                                            #ExprNil
         | IDENTIFICADOR                                                                     #ExprIdentificador
         | nomInstancia= IDENTIFICADOR PUNTO nomFunc= IDENTIFICADOR PARABRE PARCIERRA        #ExprLlamadaFuncionStruct
         | llamadafuncion                                                                    #ExprLlamadaFuncion
         | llamadaIndexOf                                                                    #ExprIndexOf
         | llamadaJoin                                                                       #ExprJoin
         | llamadaLen                                                                        #ExprLen
         | llamadaAppend                                                                     #ExprAppend
         | accesoSlice                                                                       #ExprAccesoSlice
         | accesoMatriz                                                                      #ExprAccesoMatriz
         | llamadaAtoi                                                                       #ExprAtoi
         | llamadaParseFloat                                                                 #ExprParseFloat
         | llamadaTypeOf                                                                     #ExprTypeOf
         |IDENTIFICADOR CORABRE expresion CORCIERRA                                          #ExprAccesoSlice
         |IDENTIFICADOR CORABRE expresion CORCIERRA CORABRE expresion CORCIERRA              #ExprAccesoMatriz

;



//atributosStruct : tipo  IDENTIFICADOR PUNTO_COMA ;
atributosStruct : tipo nomAtributo= IDENTIFICADOR PUNTO_COMA? ;

inicializacion_slice : CORABRE lista_expresion? CORCIERRA
                |CORABRE lista_slices? CORCIERRA
                ;
lista_expresion : expresion (SIGNO_COMA expresion)* ;

lista_slices : inicializacion_slice ( SIGNO_COMA inicializacion_slice)*; 


// literal: INT            #LiteralInt
//         |FLOAT          #LiteralFloat
//         |BOOL           #LiteralBool    
//         |CADENA         #LiteralString
//         |RUNE           #LiteralRune
//         |RESNIL         #LiteralNil
//         ;

// tipo: RESINT
//     |RESFLOAT
//     |RESSTRING
//     |RESBOOL
//     |RESRUNE
//     |'[]' tipo
//     |IDENTIFICADOR
//     ;

tipo: RESINT
    |RESFLOAT
    |RESSTRING
    |RESBOOL
    |RESRUNE
    |IDENTIFICADOR
    ;

// BOOL: 'true' | 'false';

//bloque : LLAVEABRE sentencia* LLAVECIERRA; 
bloque : LLAVEABRE lista_sentencias* LLAVECIERRA; 

importStmt : (RESIMPORT CADENA) PUNTO_COMA+ ;


creacion_instancia_struct: nomInstancia= IDENTIFICADOR ASIGNACION_INFERIDA nomStruct= IDENTIFICADOR LLAVEABRE listaCampos? LLAVECIERRA PUNTO_COMA? ;

listaCampos : campoStruct (SIGNO_COMA campoStruct)* ;

campoStruct :
            idCampo= IDENTIFICADOR DOS_PUNTOS expresion PUNTO_COMA? 
            ;

llamada_funcion : nomFunc= IDENTIFICADOR PARABRE (lista_parametros_func)? PARCIERRA;

lista_parametros_func :  param_func (SIGNO_COMA param_func)* ;

param_func: ('&'? idInferencia= IDENTIFICADOR) 
          | expresion
          ;

sentenciaFor : RESFOR expresion bloque                                                          #SentFor1
              |RESFOR PARABRE sentencia_init PUNTO_COMA expresion PUNTO_COMA sentencia_update PARCIERRA bloque #SentForClasico
             //| RESFOR IDENTIFICADOR (SIGNO_COMA IDENTIFICADOR) RESIN IDENTIFICADOR bloque         #SentFor3
            ;
sentencia_init : declaracion_variable
               | asignacion
               ;
sentencia_update : asignacion
                 | incremento_variable
                 ;

incremento_variable : IDENTIFICADOR op= (SIGNO_INCREMENTO|SIGNO_DECREMENTO) PUNTO_COMA? 
;

declaracion_for : IDENTIFICADOR ASIGNACION_INFERIDA expresion;


llamadaIndexOf : RESINDEX PARABRE nomSlice= IDENTIFICADOR SIGNO_COMA expresion PARCIERRA;

llamadaJoin : RESJOIN PARABRE nomSlice= IDENTIFICADOR SIGNO_COMA expresion PARCIERRA;

llamadaLen : RESLEN PARABRE nomSlice= IDENTIFICADOR PARCIERRA;

llamadaAppend : RESAPPEND PARABRE nomSlice= IDENTIFICADOR SIGNO_COMA expresion PARCIERRA;

accesoSlice : nomSlice= IDENTIFICADOR CORABRE index= expresion CORCIERRA;

declaracion_matriz : idMatriz= IDENTIFICADOR ASIGNACION_INFERIDA LLAVEABRE LLAVECIERRA LLAVEABRE LLAVECIERRA tipo lista_arreglos? PUNTO_COMA? ;

lista_arreglos : LLAVEABRE (lista_arreglos| expresion) (SIGNO_COMA (lista_arreglos | expresion))* LLAVECIERRA;

asignacion_acceso_matriz : 
        idMatriz= IDENTIFICADOR CORABRE index1= expresion CORCIERRA CORABRE index2= expresion CORCIERRA SIGNO_ASIGNACION expresion PUNTO_COMA? ;

accesoMatriz : idMatriz= IDENTIFICADOR CORABRE index1= expresion CORCIERRA CORABRE index2= expresion CORCIERRA;

llamadaAtoi : RESATOI PARABRE expresion PARCIERRA;

llamadaParseFloat : RESPARSEFLOAT PARABRE expresion PARCIERRA;

llamadaTypeOf : RESTYPEOF PARABRE expresion PARCIERRA;


```