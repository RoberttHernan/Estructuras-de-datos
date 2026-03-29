# Interprete V-Lang Cherry en Go

Interprete de recorrido de arbol (tree-walking interpreter) para el lenguaje V-Lang Cherry, implementado en Go. El lenguaje tiene una sintaxis inspirada en V-Lang y esta orientado a demostrar los conceptos fundamentales de la construccion de un compilador/interprete: analisis lexico, sintactico y semantico, construccion de AST, evaluacion y generacion de codigo.

## Descripcion del Proyecto

El interprete toma codigo fuente en V-Lang Cherry, lo analiza con un lexer/parser generado por ANTLR4, construye un Arbol de Sintaxis Abstracta (AST) mediante el patron Visitor, y luego evalua ese AST directamente. Adicionalmente genera codigo ensamblador ARM64 a partir del AST.

La interfaz grafica (Fyne) permite editar codigo, ejecutarlo, y visualizar el AST y la tabla de simbolos.

---

## Arquitectura

```
Entrada (texto fuente)
       |
  [gramatica.g4]  --ANTLR4--> parser/ (lexer + parser en Go)
       |
  visitor/ast_builder.go       (CST -> AST usando patron Visitor)
       |
  interpreter/interpreter.go   (evaluacion del AST por recorrido de arbol)
       |
  GeneradorCodigo/              (generacion de ensamblador ARM64 desde el AST)
```

### Analisis Lexico y Sintactico (`parser/`, `gramatica.g4`)

Generados automaticamente por ANTLR4 a partir de la gramatica `gramatica.g4`. El lexer reconoce tokens de palabras reservadas, literales, operadores y rutas. El parser genera un CST (Concrete Syntax Tree).

Un `CustomErrorListener` (`errores/error_lister.go`) reemplaza el listener por defecto de ANTLR para acumular errores lexicos y sintacticos en una lista exportable.

### Construccion del AST (`visitor/ast_builder.go`)

Implementa el patron Visitor sobre el CST de ANTLR. Transforma cada regla gramatical en un nodo del AST definido en `ast/nodes.go`. Soporta todos los constructos del lenguaje: declaraciones, expresiones, estructuras de control, funciones, structs y slices.

### Interprete (`interpreter/interpreter.go`)

Evalua el AST mediante un switch sobre el tipo de nodo. Gestiona el estado con una cadena de entornos (`Environment`) que implementa alcance lexico estatico.

El control de flujo no-local se implementa con `panic/recover`:
- `panic(ReturnValue{...})` senaliza un `return` de funcion
- `panic("break_signal")` senaliza un `break` de ciclo
- `panic("continue_signal")` senaliza un `continue`

Las operaciones aritmeticas soportan conversion implicita `int <-> float64`.

### Nodos del AST (`ast/nodes.go`)

Cada nodo implementa la interfaz `Node`:
- `String() string` — representacion textual para reportes
- `GenerateCode(*CodeBuilder)` — emite instrucciones ARM64

| Nodo | Descripcion |
|---|---|
| `Prog` | Raiz del programa; lista de sentencias de nivel superior |
| `FunctionDeclarationNode` | Declaracion de funcion con parametros y cuerpo |
| `DeclarationNode` | Declaracion de variable (`mut`, `:=`, tipada) |
| `AssignmentNode` | Asignacion simple o compuesta (`=`, `+=`, `-=`) |
| `BinaryOperationNode` | Expresion binaria (`+`, `-`, `*`, `/`, `%`, comparaciones, logicos) |
| `UnaryOperationNode` | Negacion aritmetica o logica (`-`, `!`) |
| `IfNode` | Condicional con ramas `else if` y `else` opcionales |
| `ForNode` | Ciclo `while` (condicion booleana) |
| `ForClassicNode` | Ciclo `for` clasico con inicializacion, condicion y actualizacion |
| `SwtchNode` | `switch/case/default` |
| `FunctionCallNode` | Llamada a funcion definida por el usuario |
| `BuiltinFunCallNode` | Llamada a funcion nativa (`len`, `Atoi`, `typeOf`, `append`) |
| `ReturnNode` | Instruccion `return` |
| `BreakNode` / `ContinueNode` | Control de ciclos |
| `SliceAccessNode` | Acceso a elemento de slice por indice |
| `SliceAssignmentNode` | Asignacion a elemento de slice |
| `PrintNode` | Instruccion `print` |
| `LiteralNode` | Literal: entero, flotante, string, bool |
| `IdentifierNode` | Referencia a variable |
| `BlockNode` | Bloque de sentencias con su propio ambito |

### Generacion de Codigo ARM64 (`GeneradorCodigo/codigo_generator.go`)

`CodeBuilder` acumula instrucciones de texto, comentarios, datos y etiquetas para producir un archivo `.s` ensamblable con `as` en arquitecturas ARM64. Incluye una rutina `print_int` para imprimir enteros en consola via syscall.

### Reportes (`reportes/`)

- `ast_report.go`: genera un reporte textual del AST (recorrido pre-orden)
- `ast_dot.go`: genera un archivo `.dot` para visualizar el AST con Graphviz

---

## Lenguaje V-Lang Cherry

### Tipos de datos

| Tipo | Descripcion |
|---|---|
| `int` | Entero con signo |
| `float64` | Numero flotante de doble precision |
| `string` | Cadena de caracteres |
| `bool` | Valor booleano (`true` / `false`) |
| `rune` | Caracter individual |
| `[]tipo` | Slice (arreglo dinamico) |

### Declaracion de variables

```
mut x int = 5        // mutable, con tipo y valor
x := 10              // inmutable, con inferencia de tipo
var x int = 5        // inmutable, con tipo explicito
var x int ?          // sin valor inicial
```

### Funciones

```
func nombre(param1 tipo1, param2 tipo2) -> tipoRetorno {
    // cuerpo
    return valor
}
```

### Estructuras de control

```
// If
if condicion {
    // ...
} else if condicion2 {
    // ...
} else {
    // ...
}

// Switch
switch expresion {
    case valor1:
        // ...
    default:
        // ...
}

// While
while condicion {
    // ...
}

// For clasico
for i := 0; i < 10; i += 1 {
    // ...
}
```

### Funciones nativas

| Funcion | Descripcion |
|---|---|
| `print(...)` | Imprime en consola |
| `len(x)` | Longitud de un string o slice |
| `Atoi(s)` | Convierte string a int |
| `typeOf(x)` | Retorna el tipo como string |
| `append(slice, val)` | Agrega un elemento a un slice |

### Comentarios

```
// Comentario de una linea
/* Comentario
   de multiples lineas */
```

---

## Dependencias

- **Go 1.22+**
- **ANTLR4 runtime para Go**: `github.com/antlr4-go/antlr/v4`
- **Fyne v2** (GUI): `fyne.io/fyne/v2`
- **gofpdf** (reportes PDF): `github.com/jung-kurt/gofpdf`
- **Graphviz** (opcional, para visualizar el AST desde archivos `.dot`)

---

## Compilacion y Ejecucion

```bash
# Instalar dependencias
go mod tidy

# Compilar y ejecutar
go run .

# El interprete lee el archivo entrada.vhc en el directorio actual
```

Para usar la interfaz grafica, el modo GUI debe estar habilitado en `main.go` (ver `ui/main_ui.go`).

---

## Estructura del Repositorio

```
Interprete Vlang Cherry/
+-- ast/
|   +-- nodes.go                  # Definiciones de todos los nodos del AST
+-- visitor/
|   +-- ast_builder.go            # Patron Visitor: CST (ANTLR) -> AST
+-- interpreter/
|   +-- interpreter.go            # Evaluador del AST por recorrido de arbol
+-- errores/
|   +-- error_lister.go           # Listener de errores ANTLR personalizado
+-- parser/                       # Lexer y parser generados por ANTLR4 (Go)
+-- GeneradorCodigo/
|   +-- codigo_generator.go       # Generador de ensamblador ARM64
+-- reportes/
|   +-- ast_report.go             # Reporte textual del AST
|   +-- ast_dot.go                # Reporte Graphviz del AST
+-- ui/
|   +-- main_ui.go                # Interfaz grafica con Fyne
+-- docs/
|   +-- gramatica/readme.md       # Especificacion de la gramatica ANTLR4
|   +-- tecnico/readme.md         # Manual tecnico
|   +-- usuario/readme.md         # Manual de usuario con especificacion del lenguaje
+-- entradas/                     # Archivos de prueba (.vhc, .txt)
+-- gramatica.g4                  # Gramatica ANTLR4 del lenguaje
+-- main.go                       # Punto de entrada
+-- interfaz_adapter.go           # Adaptador entre GUI e interprete
+-- go.mod / go.sum               # Modulo y dependencias Go
+-- .gitignore
+-- README.md
```
