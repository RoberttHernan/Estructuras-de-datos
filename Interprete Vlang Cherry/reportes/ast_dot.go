package reportes

import (
	"fmt"
	"os"
	"os/exec"
	"proyecto1/ast"
	"strings"
)

var contador int

func GenerarDotAST(nodo ast.Node, archivo string) {
	contador = 0
	f, err := os.Create(archivo)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("digraph AST {\n")
	f.WriteString("  node [shape=box, style=filled, fillcolor=lightgray];\n")

	raiz := generarNodoDot(f, nodo)

	f.WriteString(fmt.Sprintf("  %s;\n", raiz))
	f.WriteString("}\n")

}

func generarNodoDot(f *os.File, nodo ast.Node) string {
	if nodo == nil {
		return ""
	}

	id := fmt.Sprintf("n%d", contador)
	contador++

	label := fmt.Sprintf("%T", nodo)
	if str := nodo.String(); str != "" {
		escaped := strings.ReplaceAll(nodo.String(), "\"", "\\\"")
		label = fmt.Sprintf("%s\\n%s", label, escaped)
	}

	fmt.Fprintf(f, "  %s [label=\"%s\"];\n", id, label)

	switch n := nodo.(type) {
	case *ast.Prog:
		for _, stmt := range n.Elementos {
			hijo := generarNodoDot(f, stmt)
			fmt.Fprintf(f, "  %s -> %s;\n", id, hijo)
		}
	case *ast.BlockNode:
		for _, stmt := range n.Statements {
			hijo := generarNodoDot(f, stmt)
			fmt.Fprintf(f, "  %s -> %s;\n", id, hijo)
		}
	case *ast.FunctionDeclarationNode:
		hijo := generarNodoDot(f, n.Body)
		fmt.Fprintf(f, "  %s -> %s;\n", id, hijo)
	case *ast.IfNode:
		fmt.Fprintf(f, "  %s_cond [label=\"Condicion\"];\n", id)
		fmt.Fprintf(f, "  %s -> %s_cond;\n", id, id)
		h1 := generarNodoDot(f, n.Condition)
		fmt.Fprintf(f, "  %s_cond -> %s;\n", id, h1)
		h2 := generarNodoDot(f, n.IfBlock)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h2)
		if n.ElseIf != nil {
			h3 := generarNodoDot(f, n.ElseIf)
			fmt.Fprintf(f, "  %s -> %s;\n", id, h3)
		}
		if n.ElseBlock != nil {
			h4 := generarNodoDot(f, n.ElseBlock)
			fmt.Fprintf(f, "  %s -> %s;\n", id, h4)
		}
	case *ast.DeclarationNode:
		h1 := generarNodoDot(f, n.Identifier)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h1)
		h2 := generarNodoDot(f, n.Value)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h2)
	case *ast.AssignmentNode:
		h1 := generarNodoDot(f, n.Variable)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h1)
		h2 := generarNodoDot(f, n.Value)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h2)
	case *ast.PrintNode:
		for _, expr := range n.Expressions {
			h := generarNodoDot(f, expr)
			fmt.Fprintf(f, "  %s -> %s;\n", id, h)
		}
	case *ast.FunctionCallNode:
		for _, expr := range n.Arguments {
			h := generarNodoDot(f, expr)
			fmt.Fprintf(f, "  %s -> %s;\n", id, h)
		}
	case *ast.BuiltinFunCallNode:
		for _, expr := range n.Arguments {
			h := generarNodoDot(f, expr)
			fmt.Fprintf(f, "  %s -> %s;\n", id, h)
		}
	case *ast.BinaryOperationNode:
		h1 := generarNodoDot(f, n.Left)
		h2 := generarNodoDot(f, n.Right)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h1)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h2)
	case *ast.UnaryOperationNode:
		h := generarNodoDot(f, n.Operand)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h)
	case *ast.ReturnNode:
		h := generarNodoDot(f, n.Value)
		fmt.Fprintf(f, "  %s -> %s;\n", id, h)
	}

	return id
}

func GenerarImagenDot(dotFile, outputFile string) {
	cmd := exec.Command("dot", "-Tpng", dotFile, "-o", outputFile)
	err := cmd.Run()
	if err != nil {
		panic("❌ Error al generar imagen .png del AST: " + err.Error())
	}
}
