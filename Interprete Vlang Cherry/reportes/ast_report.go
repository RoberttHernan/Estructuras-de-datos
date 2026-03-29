package reportes

import (
	"fmt"
	"os"
	"proyecto1/ast"
)

func GenerarReporteAST(nodo ast.Node, archivo string) {
	f, err := os.Create(archivo)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writeNodo(f, nodo, 0)
}

func writeNodo(f *os.File, nodo ast.Node, nivel int) {
	if nodo == nil {
		return
	}

	indent := ""
	for i := 0; i < nivel; i++ {
		indent += "  "
	}

	fmt.Fprintf(f, "%s- %T: %s\n", indent, nodo, nodo.String())

	switch n := nodo.(type) {
	case *ast.Prog:
		for _, stmt := range n.Elementos {
			writeNodo(f, stmt, nivel+1)
		}

	case *ast.BlockNode:
		for _, stmt := range n.Statements {
			writeNodo(f, stmt, nivel+1)
		}

	case *ast.FunctionDeclarationNode:
		writeNodo(f, n.Body, nivel+1)

	case *ast.IfNode:
		writeNodo(f, n.Condition, nivel+1)
		writeNodo(f, n.IfBlock, nivel+1)
		if n.ElseIf != nil {
			writeNodo(f, n.ElseIf, nivel+1)
		}
		if n.ElseBlock != nil {
			writeNodo(f, n.ElseBlock, nivel+1)
		}

	case *ast.ForNode:
		writeNodo(f, n.Init, nivel+1)
		writeNodo(f, n.Condition, nivel+1)
		writeNodo(f, n.Post, nivel+1)
		writeNodo(f, n.Body, nivel+1)

	case *ast.ForClassicNode:
		writeNodo(f, n.Init, nivel+1)
		writeNodo(f, n.Condition, nivel+1)
		writeNodo(f, n.Update, nivel+1)
		writeNodo(f, n.Body, nivel+1)

	case *ast.PrintNode:
		for _, expr := range n.Expressions {
			writeNodo(f, expr, nivel+1)
		}

	case *ast.AssignmentNode:
		writeNodo(f, n.Variable, nivel+1)
		writeNodo(f, n.Value, nivel+1)

	case *ast.DeclarationNode:
		writeNodo(f, n.Identifier, nivel+1)
		writeNodo(f, n.Value, nivel+1)

	case *ast.FunctionCallNode:
		for _, arg := range n.Arguments {
			writeNodo(f, arg, nivel+1)
		}

	case *ast.BuiltinFunCallNode:
		for _, arg := range n.Arguments {
			writeNodo(f, arg, nivel+1)
		}

	case *ast.BinaryOperationNode:
		writeNodo(f, n.Left, nivel+1)
		writeNodo(f, n.Right, nivel+1)

	case *ast.UnaryOperationNode:
		writeNodo(f, n.Operand, nivel+1)

	case *ast.ReturnNode:
		writeNodo(f, n.Value, nivel+1)

	case *ast.SwtchNode:
		writeNodo(f, n.Condition, nivel+1)
		for _, cs := range n.Cases {
			writeNodo(f, cs, nivel+1)
		}
		writeNodo(f, n.DefaultCase, nivel+1)

	case *ast.CaseNode:
		writeNodo(f, n.Value, nivel+1)
		writeNodo(f, n.Body, nivel+1)

	case *ast.SliceAccessNode:
		writeNodo(f, n.Collection, nivel+1)
		writeNodo(f, n.Index, nivel+1)

	case *ast.SliceAssignmentNode:
		writeNodo(f, n.Colection, nivel+1)
		writeNodo(f, n.Index, nivel+1)
		writeNodo(f, n.Value, nivel+1)

	}
}
