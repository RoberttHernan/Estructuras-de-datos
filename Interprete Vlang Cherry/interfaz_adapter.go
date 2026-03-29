package main

import (
	"bytes"
	"fmt"
	"os"
	generadorcodigo "proyecto1/GeneradorCodigo"
	"proyecto1/ast"
	"proyecto1/errores"
	"proyecto1/parser"
	"proyecto1/reportes"
	"proyecto1/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func EjecutarCodigo(codigo string) string {
	defer func() string {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("💥 Error de ejecución: %v", r)
			_ = escribirArchivo("errores_ejecucion.txt", msg)
			return msg
		}
		return ""
	}()

	var salida bytes.Buffer

	input := antlr.NewInputStream(codigo)
	lexer := parser.NewgramaticaLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgramaticaParser(tokens)

	errorListener := errores.NewCustomErrorListener()
	lexer.RemoveErrorListeners()
	p.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	p.AddErrorListener(errorListener)

	tree := p.Prog()

	if len(errorListener.Errores) > 0 {
		for _, err := range errorListener.Errores {
			salida.WriteString(fmt.Sprintf("❌ %s\n", err))
		}
		errorListener.ExportToFile("errores_sintacticos.txt")
		return salida.String()
	}

	visitor := visitor.NewASTBuilder()
	result := visitor.Visit(tree)
	program := result.(ast.Node)

	//inter := interpreter.NewInterpreter()
	//inter.Output = &salida // 🔁 Redirige los print de VLangCherry aquí
	//inter.Evaluate(result.(*ast.Prog))
	//inter.Global().GenerarReporteTablaSimbolos("tabla_simbolos.txt")

	builder := generadorcodigo.NewCodeBuilder()
	builder.AddLine(".text")
	builder.AddLine(".global _start")
	builder.AddLine("_start:")
	program.GenerateCode(builder)

	err := escribirArchivo("salida.s", builder.EmitAll())
	if err != nil {
		salida.WriteString(fmt.Sprintf("❌ Error al guardar archivo salida.s: %v\n", err))
	} else {
		salida.WriteString("✅ Archivo ensamblador generado correctamente: salida.s\n")
	}
	reportes.GenerarReporteAST(result.(ast.Node), "ast.txt")
	reportes.GenerarDotAST(result.(ast.Node), "ast.dot")
	//_ = inter.GenerarReporteTablaSimbolos("tabla_simbolos.txt")

	return salida.String()
}

func escribirArchivo(nombre string, contenido string) error {
	return os.WriteFile(nombre, []byte(contenido), 0644)
}
