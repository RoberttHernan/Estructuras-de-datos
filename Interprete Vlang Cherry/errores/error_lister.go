// Package errores provee un listener de errores personalizado para ANTLR4.
// Reemplaza el DefaultErrorListener de ANTLR para acumular los mensajes de
// error lexico y sintactico en una lista en lugar de imprimirlos directamente,
// permitiendo exportarlos a un archivo al final del analisis.
package errores

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

// CustomErrorListener acumula los errores sintacticos y lexicos detectados por ANTLR.
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errores []string // Lista de mensajes de error en formato "Linea L:C -> descripcion"
}

// NewCustomErrorListener crea un listener de errores vacio listo para usar.
func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		Errores: []string{},
	}
}

// SyntaxError se invoca automaticamente por ANTLR cuando detecta un error.
// Formatea el mensaje con linea y columna y lo agrega a la lista de errores.
func (el *CustomErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	errorMsg := fmt.Sprintf("Linea %d:%d -> %s", line, column, msg)
	fmt.Println("Error detectado:", errorMsg)
	el.Errores = append(el.Errores, errorMsg)
}

// ExportToFile escribe todos los errores acumulados en el archivo indicado,
// uno por linea. Crea el archivo si no existe o lo sobreescribe.
func (el *CustomErrorListener) ExportToFile(path string) {
	f, _ := os.Create(path)
	defer f.Close()

	for _, err := range el.Errores {
		f.WriteString(err + "\n")
	}
}
