package main

/*
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"proyecto1/ast"
	"proyecto1/errores"
	"proyecto1/interpreter"
	"proyecto1/parser"
	"proyecto1/visitor"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jung-kurt/gofpdf"
)

type API struct {
	interpreter   *interpreter.Interpreter
	lastAST       ast.Node
	lastErrors    []ErrorReport
	lastSymbolTable []SymbolEntry
}

func NewAPI() *API {
	return &API{
		interpreter: interpreter.NewInterpreter(),
	}
}

func (api *API) StartServer() {
	http.HandleFunc("/execute", api.handleExecute)
	http.HandleFunc("/ast-visual", api.handleASTVisualization)
	http.HandleFunc("/export/pdf", api.handleExportPDF)
	http.HandleFunc("/export/html", api.handleExportHTML)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Servidor backend iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (api *API) handleExecute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req ExecutionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al leer la solicitud", http.StatusBadRequest)
		return
	}

	// Procesar el código
	input := antlr.NewInputStream(req.Code)
	lexer := parser.NewgramaticaLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgramaticaParser(tokens)

	// Manejo de errores
	errorListener := errores.NewCustomErrorListener()
	lexer.RemoveErrorListeners()
	p.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	p.AddErrorListener(errorListener)

	tree := p.Prog()

	// Construir AST
	visitor := visitor.NewASTBuilder()
	astResult := visitor.Visit(tree).(*ast.Prog)
	api.lastAST = astResult

	// Interpretar
	api.interpreter.Evaluate(astResult)

	// Convertir errores
	api.lastErrors = convertErrors(errorListener.Errores)
	api.lastSymbolTable = api.interpreter.GetSymbolTable()

	// Preparar respuesta
	response := ExecutionResponse{
		Output:       captureInterpreterOutput(),
		Errors:       api.lastErrors,
		SymbolTable:  api.lastSymbolTable,
		AST:          astResult.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (api *API) handleASTVisualization(w http.ResponseWriter, r *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>AST Visualization</title>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/vis/4.21.0/vis.min.js"></script>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/vis/4.21.0/vis.min.css" rel="stylesheet">
		<style>
			#mynetwork { width: 100%; height: 800px; border: 1px solid #444; }
			.details { margin: 20px; padding: 10px; background: #f5f5f5; }
		</style>
	</head>
	<body>
		<h1>Abstract Syntax Tree</h1>
		<div class="details">
			<button onclick="generateImage()">Exportar como PNG</button>
			<button onclick="window.print()">Imprimir</button>
		</div>
		<div id="mynetwork"></div>
		<script>
			var nodes = new vis.DataSet({{.Nodes}});
			var edges = new vis.DataSet({{.Edges}});
			var container = document.getElementById('mynetwork');
			var data = { nodes: nodes, edges: edges };
			var options = {
				layout: { hierarchical: { direction: "UD", sortMethod: "directed" } },
				physics: false,
				nodes: { shape: 'box', margin: 10, font: { size: 12, face: 'monospace' } }
			};
			var network = new vis.Network(container, data, options);

			function generateImage() {
				var dataURL = network.toImage({ type: 'png', quality: 1.0 });
				var link = document.createElement('a');
				link.download = 'ast-diagram.png';
				link.href = dataURL;
				link.click();
			}
		</script>
	</body>
	</html>`

	nodes, edges := convertASTToVisJS(api.lastAST)

	tmpl, err := template.New("ast").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Nodes string
		Edges string
	}{
		Nodes: nodes,
		Edges: edges,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

func (api *API) handleExportPDF(w http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Reporte VLangCherry")
	pdf.Ln(12)

	// Encabezado de errores
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Errores")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)

	// Tabla de errores
	for _, err := range api.lastErrors {
		pdf.Cell(40, 6, fmt.Sprintf("Línea %d:%d - %s (%s)",
			err.Line, err.Column, err.Message, err.ErrorType))
		pdf.Ln(6)
	}

	pdf.Ln(10)

	// Tabla de símbolos
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Tabla de Símbolos")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)

	// Encabezados de tabla
	pdf.Cell(40, 6, "ID")
	pdf.Cell(30, 6, "Tipo")
	pdf.Cell(30, 6, "Tipo Dato")
	pdf.Cell(30, 6, "Ámbito")
	pdf.Cell(20, 6, "Línea")
	pdf.Ln(6)

	// Contenido de tabla
	for _, sym := range api.lastSymbolTable {
		pdf.Cell(40, 6, sym.ID)
		pdf.Cell(30, 6, sym.Type)
		pdf.Cell(30, 6, sym.DataType)
		pdf.Cell(30, 6, sym.Scope)
		pdf.Cell(20, 6, fmt.Sprintf("%d", sym.Line))
		pdf.Ln(6)
	}

	pdf.Ln(10)

	// AST
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Árbol de Sintaxis Abstracta (AST)")
	pdf.Ln(8)
	pdf.SetFont("Courier", "", 8)
	pdf.MultiCell(0, 5, api.lastAST.String(), "0", "L", false)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	pdf.Output(w)
}

func (api *API) handleExportHTML(w http.ResponseWriter, r *http.Request) {
	const tpl = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>VLangCherry Report</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			h1 { color: #2c3e50; }
			h2 { color: #3498db; margin-top: 30px; }
			table { border-collapse: collapse; width: 100%; margin-bottom: 20px; }
			th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
			th { background-color: #f2f2f2; }
			pre { background: #f4f4f4; padding: 10px; border-radius: 5px; overflow-x: auto; }
			.timestamp { color: #7f8c8d; font-size: 0.9em; }
			.print-button { margin: 20px 0; padding: 10px 15px; background: #3498db; color: white; border: none; border-radius: 4px; cursor: pointer; }
		</style>
	</head>
	<body>
		<h1>VLangCherry Report <span class="timestamp">{{.Timestamp}}</span></h1>

		<button class="print-button" onclick="window.print()">Imprimir Reporte</button>

		<h2>Errores</h2>
		{{if .Errors}}
		<table>
			<tr><th>Línea</th><th>Columna</th><th>Tipo</th><th>Mensaje</th></tr>
			{{range .Errors}}
			<tr>
				<td>{{.Line}}</td><td>{{.Column}}</td><td>{{.ErrorType}}</td><td>{{.Message}}</td>
			</tr>
			{{end}}
		</table>
		{{else}}
		<p>No se encontraron errores.</p>
		{{end}}

		<h2>Tabla de Símbolos</h2>
		{{if .SymbolTable}}
		<table>
			<tr><th>ID</th><th>Tipo</th><th>Tipo Dato</th><th>Ámbito</th><th>Línea</th></tr>
			{{range .SymbolTable}}
			<tr>
				<td>{{.ID}}</td><td>{{.Type}}</td><td>{{.DataType}}</td><td>{{.Scope}}</td><td>{{.Line}}</td>
			</tr>
			{{end}}
		</table>
		{{else}}
		<p>No se encontraron símbolos.</p>
		{{end}}

		<h2>AST</h2>
		<pre>{{.AST}}</pre>

		<script>
			document.addEventListener('DOMContentLoaded', function() {
				document.querySelector('.timestamp').textContent = new Date().toLocaleString();
			});
		</script>
	</body>
	</html>`

	data := struct {
		Timestamp   string
		Errors      []ErrorReport
		SymbolTable []SymbolEntry
		AST         string
	}{
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Errors:      api.lastErrors,
		SymbolTable: api.lastSymbolTable,
		AST:         api.lastAST.String(),
	}

	tmpl, err := template.New("report").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Disposition", "attachment; filename=report.html")
	tmpl.Execute(w, data)
}

// Funciones auxiliares
func convertASTToVisJS(astNode ast.Node) (string, string) {
	// Implementación real necesitaría recorrer el AST completo
	// Esta es una implementación simplificada para el ejemplo
	nodes := []map[string]interface{}{
		{"id": 1, "label": "Program", "shape": "box", "color": "#97C2FC"},
		{"id": 2, "label": "Function: main", "shape": "box", "color": "#FFA07A"},
	}
	edges := []map[string]interface{}{
		{"from": 1, "to": 2, "arrows": "to"},
	}

	nodesJSON, _ := json.Marshal(nodes)
	edgesJSON, _ := json.Marshal(edges)
	return string(nodesJSON), string(edgesJSON)
}

func captureInterpreterOutput() string {
	// Implementar captura de salida del intérprete
	return "Salida de ejecución:\n- Función main ejecutada\n- 5 declaraciones procesadas"
}

func convertErrors(errors []string) []ErrorReport {
	// Implementar conversión de errores del parser
	return []ErrorReport{
		{Line: 1, Column: 5, Message: "Falta punto y coma", ErrorType: "Sintáctico"},
	}
}*/
