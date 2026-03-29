// Punto de entrada del interprete V-Lang Cherry.
// Flujo principal:
//  1. Lee el archivo de entrada entrada.vhc
//  2. Tokeniza y parsea con el lexer/parser generado por ANTLR4
//  3. Construye el AST mediante el patron Visitor (ast_builder.go)
//  4. Evalua el AST con el interprete de recorrido de arbol (interpreter.go)
//
// Los errores sintacticos se exportan a errores_sintacticos.txt.
// Los panics de ejecucion se recuperan y se exportan a errores_ejecucion.txt.

package main

import (
	"fmt"
	"os"
	"proyecto1/ast"
	"proyecto1/errores"
	"proyecto1/interpreter"
	"proyecto1/parser"
	"proyecto1/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func main() {

	defer ManejarErroresDeEjecucion()
	file, err := os.ReadFile("entrada.vhc")
	if err != nil {

		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	input := antlr.NewInputStream(string(file))
	lexer := parser.NewgramaticaLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgramaticaParser(tokens)

	//fmt.Println(tree.ToStringTree(nil, p))

	// Manejo de errores
	errorListener := errores.NewCustomErrorListener()
	lexer.RemoveErrorListeners() // Elimina los listeners por defecto
	p.RemoveErrorListeners()     // Elimina los listeners por defecto

	lexer.AddErrorListener(errorListener) // Añade tu listener personalizado
	p.AddErrorListener(errorListener)     // Añade tu listener personalizado

	tree := p.Prog() // Cambia "Start" por tu regla inicial

	if len(errorListener.Errores) > 0 {
		fmt.Println("Errores detectados:")
		for _, err := range errorListener.Errores {
			fmt.Println(err)
		}

		// Guardar en archivo
		errorListener.ExportToFile("errores_sintacticos.txt")
		fmt.Println("Errores guardados en errores_sintacticos.txt")

		os.Exit(1) // Salir si hay errores
	}

	// de aqui para abajo es el visito y el ast
	visitor := visitor.NewASTBuilder()
	result := visitor.Visit(tree) // Asegúrate de que el resultado sea del tipo correcto

	interpreter := interpreter.NewInterpreter()
	interpreter.Evaluate(result.(*ast.Prog)) // Asegúrate de que el resultado sea del tipo correcto
	//interpreter.Global().GenerarReporteTablaSimbolos("tabla_simbolos.txt")

}

// ManejarErroresDeEjecucion recupera cualquier panic que no haya sido capturado
// durante la evaluacion del AST y lo escribe en errores_ejecucion.txt.
func ManejarErroresDeEjecucion() {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Error de ejecucion: %v", r)
		fmt.Println(msg)
		os.WriteFile("errores_ejecucion.txt", []byte(msg), 0644)
	}
}

/*func EjecutarCodigo() {

}*/
