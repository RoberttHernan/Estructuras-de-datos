# **Manual Técnico**

## **Tabla de Contenido**  
- [Descripción General](#descripción-general)  
- [Descripción de la solución](#descripción-de-la-solución)  
- [Estructura del intérprete](#estructura-del-intérprete)  
- [Análisis Léxico](#análisis-léxico)  
- [Análisis Sintáctico](#análisis-sintáctico)  
- [Análisis Semántico](#análisis-semántico)  
- [Manejo de errores](#manejo-de-errores)  
- [Consideraciones del lenguaje](#consideraciones-del-lenguaje)  
- [Patrón Visitor](#patrón-visitor)  
- [Estructuras de datos relevantes](#estructuras-de-datos-relevantes)  
- [Interfaz gráfica](#interfaz-gráfica)  

## **Descripción General**  
El objetivo del proyecto es desarrollar un intérprete para el lenguaje de programación V-lang Cherry, un lenguaje inspirado en VLang pero adaptado para entornos con bajos recursos. El intérprete incluye análisis léxico, sintáctico y semántico, generación de AST y un sistema de manejo de errores.

## **Descripción de la solución**  
El proyecto se implementó en Go utilizando ANTLR para la generación del lexer y parser. La interfaz gráfica se desarrolló con Fyne, permitiendo la edición, ejecución y visualización de reportes. La arquitectura sigue un modelo de visitante para el recorrido del AST.

## **Estructura del intérprete**  
El intérprete consta de tres componentes principales:  
1. **Frontend**: Interfaz gráfica para edición de código  
2. **Analizador**: Léxico, sintáctico y semántico  
3. **Backend**: Ejecución e interpretación del código  

## **Análisis Léxico**  
Implementado con ANTLR mediante el archivo `VLangCherryLexer.g4`. Los tokens reconocidos incluyen:  
- Palabras reservadas: `mut, func, struct, if, else, switch, case, for, while`  
- Literales: números, strings, booleanos  
- Operadores: `+, -, *, /, %, ==, !=, >, <, &&, ||`  

## **Análisis Sintáctico**  
Gramática definida en `VLangCherryParser.g4` que genera el CST. Reglas principales:  
```antlr
program : (stmt)* EOF;  
stmt : decl_stmt | if_stmt | for_stmt | func_dcl | struct_dcl;  
expr : literal | id | expr op expr | func_call;  
```

## **Análisis Semántico**  
Verifica:  
- Tipos compatibles en operaciones  
- Variables declaradas antes de su uso  
- Ámbitos válidos  
- Correcto número de parámetros en funciones  

## **Manejo de errores**  
Clasificación de errores:  
- **Léxicos**: Símbolos no reconocidos  
- **Sintácticos**: Estructuras inválidas  
- **Semánticos**: Tipos incompatibles, variables no declaradas  

## **Consideraciones del lenguaje**  
- Tipado estático con inferencia  
- Ámbitos anidados  
- Structs y funciones como tipos compuestos  
- Manejo de slices y matrices  

## **Patrón Visitor**  
Implementado para recorrer el AST y ejecutar las instrucciones. Métodos principales:
```go
func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{}  
func (v *Visitor) VisitDeclaration(ctx *parser.Decl_stmtContext) interface{}  
func (v *Visitor) VisitFunctionCall(ctx *parser.Func_callContext) interface{}  
```

## **Estructuras de datos relevantes**  
- `Environment`: Manejo de ámbitos y tablas de símbolos  
- `SymbolTable`: Registro de variables y funciones  
- `AST Nodes`: Representación de estructuras del lenguaje  

## **Interfaz gráfica**  
Componentes:  
- Editor de código con resaltado de sintaxis  
- Consola de salida  
- Visualización de reportes (AST, Tabla de Símbolos)  
- Barra de herramientas para ejecución y análisis  
