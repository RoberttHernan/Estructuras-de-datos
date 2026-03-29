// base general para los nodos del AST

package ast

import (
	"fmt"
	generadorcodigo "proyecto1/GeneradorCodigo" // importa el generador de código
	// importa el paquete generado por ANTLR
)

// -------------------------------------------------------------------------------------
// interfaz parai implementar en todos los nodos del AST
type Node interface {
	String() string
	GenerateCode(cb *generadorcodigo.CodeBuilder) // método para generar código
}

type LiteralNode struct {
	// valor del literal
	Value interface{}
}

func (p *Prog) String() string {
	return "Prog"
}
func (p *Prog) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	//cb.AddText(".global _start")
	//cb.AddText("_start:")

	// Generar código para cada nodo
	for _, nodo := range p.Elementos {
		if _, ok := nodo.(*FunctionDeclarationNode); ok {
			nodo.GenerateCode(cb)
		}
	}

	cb.AddText("bl main")
	// Finalizar programa con syscall exit
	cb.AddText("mov x0, #0")
	cb.AddText("mov x8, #93")
	cb.AddText("svc #0")

	// Depuración: mostrar si se usará print_int
	fmt.Println("¿Se marcó UsesPrintInt?", cb.UsesPrintInt)

	// Agregar rutina print_int si se usó
	if cb.UsesPrintInt {
		cb.AddText("\n# Rutina print_int")
		cb.AddText("print_int:")
		cb.AddText("    mov x1, #0")
		cb.AddText("    adr x2, int_buf_end")
		cb.AddText(".loop:")
		cb.AddText("    mov x3, #10")
		cb.AddText("    udiv x4, x0, x3")
		cb.AddText("    msub x5, x4, x3, x0")
		cb.AddText("    add x5, x5, #'0'")
		cb.AddText("    sub x2, x2, #1")
		cb.AddText("    strb w5, [x2]")
		cb.AddText("    mov x0, x4")
		cb.AddText("    add x1, x1, #1")
		cb.AddText("    cbnz x0, .loop")
		cb.AddText("    mov x6, x1")  // guardar longitud
		cb.AddText("    mov x1, x2")  // puntero al inicio del número
		cb.AddText("    mov x2, x6")  // longitud
		cb.AddText("    mov x0, #1")  // stdout
		cb.AddText("    mov x8, #64") // syscall write
		cb.AddText("    svc #0")
		cb.AddText("    ret")

		// Agregar buffer en la sección .data
		cb.AddData("int_buf: .space 20")
		cb.AddData("int_buf_end: .byte 0") // ← asegura que el símbolo tenga valor
	}
}

func (c *CaseNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("⚠️ CaseNode no se ejecuta directamente")
}

//literales del tipo 10, true, "hola"

func (l *LiteralNode) String() string {
	return fmt.Sprintf("Literal(%v)", l.Value)
}

// funcion que permite generar el código arm64 para los nodos de tipo LiteralNode
func (l *LiteralNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	switch v := l.Value.(type) {
	case int:
		cb.AddComment(fmt.Sprintf("Literal int: %d", v))
		cb.AddLine(fmt.Sprintf("mov x0, #%d", v))

	case float64:
		cb.AddComment(fmt.Sprintf("Literal float64: %f", v))
		label := cb.NewLabel("float")
		cb.AddData(fmt.Sprintf("%s: .double %f", label, v))
		cb.AddText(fmt.Sprintf("ldr d0, =%s", label))

	case string:
		label := cb.NewLabel("str")
		cb.AddData(fmt.Sprintf("%s: .asciz %q", label, v))
		cb.AddText(fmt.Sprintf("ldr x1, =%s", label))
		cb.AddText(fmt.Sprintf("mov x2, #%d", len(v)))
		cb.AddText("mov x0, #1")
		cb.AddText("mov x8, #64")
		cb.AddText("svc #0")

	case bool:
		val := 0
		if v {
			val = 1
		}
		cb.AddComment(fmt.Sprintf("Literal bool: %v", v))
		cb.AddLine(fmt.Sprintf("mov x0, #%d", val))
	}
}

type Prog struct {
	Elementos []Node // lista de nodos que representan las sentencias del programa

}

// nodo para operaciones binarias
type BinaryOperationNode struct {
	Left     Node   // operando izquierdo
	Right    Node   // operando derecho
	Operator string // operador de la operación, por ejemplo "+", "-", "*", "/", "==", etc.
}

// nodo para negación unaria
type UnaryOperationNode struct {
	Operator string // operador de la operación, por ejemplo "!", "-", etc.
	Operand  Node   // operando de la operación, puede ser un literal, una variable, etc.
}

//nodo para identificadores y variables
/**
para variable como
x
miVariable
*/
type IdentifierNode struct {
	Name string
}

func (i *IdentifierNode) String() string {
	return fmt.Sprintf("Identifier(%s)", i.Name)
}
func (i *IdentifierNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Cargar variable: " + i.Name)
	cb.AddLine(fmt.Sprintf("ldr x0, =_%s", i.Name)) // ← usa _ para coincidir con .data
	cb.AddLine("ldr x0, [x0]")
}

// nodo para asignaciones de variables tipo x=5 o x+=5
type AssignmentNode struct {
	//nombre de la variable a asignar
	Variable *IdentifierNode
	//valor a asignar
	Value Node
	//operador de asignación
	Operator string // puede ser =, +=, -=, *=, /=
}

func (n *AssignmentNode) String() string {
	return fmt.Sprintf("Assignment(%s %s %s)", n.Variable.String(), n.Operator, n.Value.String())
}

func (n *AssignmentNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Asignación: " + n.Variable.Name + " " + n.Operator)

	// Evaluar el valor de la expresión (queda en x0)
	n.Value.GenerateCode(cb)

	switch n.Operator {
	case "=":
		cb.AddText(fmt.Sprintf("ldr x1, =_%s", n.Variable.Name)) // Dirección con _
		cb.AddText("str x0, [x1]")                               // Guardar valor

	case "+=":
		cb.AddText(fmt.Sprintf("ldr x1, =_%s", n.Variable.Name))
		cb.AddText("ldr x2, [x1]")
		cb.AddText("add x0, x2, x0")
		cb.AddText("str x0, [x1]")

	case "-=":
		cb.AddText(fmt.Sprintf("ldr x1, =_%s", n.Variable.Name))
		cb.AddText("ldr x2, [x1]")
		cb.AddText("sub x0, x2, x0")
		cb.AddText("str x0, [x1]")

	default:
		cb.AddComment("⚠️ Operador de asignación no soportado: " + n.Operator)
	}
}

// nodo para declaraciones de variables tipo mut  x int = 5
type DeclarationNode struct {
	Mutable bool // si la variable es mutable o no, por ejemplo: mut x int = 5
	//nombre de la variable a declara
	Identifier *IdentifierNode

	Type  string // tipo de la variable, por ejemplo "int", "string", etc.
	Value Node   // valor inicial de la variable, puede ser nil si no se inicializa

	Line  int    // número de línea donde se declara la variable, para propósitos de depuración
	Col   int    // número de columna donde se declara la variable, para propósitos de depuración
	Scope string // ámbito de la variable, por ejemplo "global", "local", etc.

}

func (d *DeclarationNode) String() string {

	//si no hay valor asignado
	if d.Value != nil {
		return fmt.Sprintf("Declare(%v %s %s = %s)", d.Mutable, d.Identifier.Name, d.Type, d.Value.String())
	}

	return fmt.Sprintf("Declare(%v %s %s)", d.Mutable, d.Identifier.Name, d.Type)

}

func (d *DeclarationNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment(fmt.Sprintf("Declaración: %s", d.Identifier.Name))

	// Detectar tipo (explícito o inferido)
	tipo := d.Type
	if tipo == "" {
		if lit, ok := d.Value.(*LiteralNode); ok {
			switch lit.Value.(type) {
			case int:
				tipo = "int"
			case float64:
				tipo = "float64"
			case string:
				tipo = "string"
			case bool:
				tipo = "bool"
			default:
				tipo = "int" // fallback
			}
		}
	}

	// Declarar variable en .data con guion bajo
	switch tipo {
	case "int", "bool":
		cb.AddData(fmt.Sprintf("_%s: .quad 0", d.Identifier.Name))
	case "float64":
		cb.AddData(fmt.Sprintf("_%s: .quad 0", d.Identifier.Name))
	case "string":
		cb.AddData(fmt.Sprintf("_%s: .quad 0", d.Identifier.Name))
	default:
		cb.AddComment(fmt.Sprintf("⚠️ Tipo desconocido: %s", tipo))
		cb.AddData(fmt.Sprintf("_%s: .quad 0", d.Identifier.Name))
	}

	// Si hay valor inicial, asignarlo
	if d.Value != nil {
		d.Value.GenerateCode(cb) // resultado queda en x0
		cb.AddText(fmt.Sprintf("ldr x1, =_%s", d.Identifier.Name))
		cb.AddText("str x0, [x1]")
	}
}

func (n *FunctionDeclarationNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio de función: " + n.Name)
	cb.AddText(n.Name + ":")
	// ---- Prólogo ----
	cb.AddText("stp x29, x30, [sp, #-16]!")
	cb.AddText("mov x29, sp")
	// ---- Cuerpo ----
	n.Body.GenerateCode(cb)
	// ---- Retorno implícito (si no hay return) ----
	if n.ReturnType == "" {
		cb.AddComment("Retorno implícito")
		cb.AddText("mov x0, #0")
	}
	// ---- Epílogo ----
	cb.AddText("ldp x29, x30, [sp], #16")
	cb.AddText("ret")
	cb.AddComment("Fin de función: " + n.Name)
}

//nodo para bloque dentro de {}

type BlockNode struct {
	//lista de statements dentro del bloque que representa las sentencias dentro del bloque
	Statements []Node
}

func (b *BlockNode) String() string {
	result := "Block("

	//indenta cada sentencia
	for _, stmt := range b.Statements {
		result += "  " + stmt.String()

	}

	result += ")"
	return result

}
func (b *BlockNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio de bloque")
	for _, stmt := range b.Statements {
		stmt.GenerateCode(cb)
	}
	cb.AddComment("Fin de bloque")
}

// nodos para bloque if
type IfNode struct {
	Condition Node // condición BOOLEANA a evaluar

	IfBlock   *BlockNode // bloque de código a ejecutar si la condición es verdadera
	ElseIf    *IfNode    // bloque de código a ejecutar si la condición es falsa (opcional)
	ElseBlock *BlockNode // bloque de código a ejecutar si la condición es falsa (opcional)

}

func (i *IfNode) String() string {
	return "IfNode"
}

func (n *IfNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	labelElse := cb.NewLabel("else")
	labelEnd := cb.NewLabel("endif")

	cb.AddComment("Evaluando condición del if")
	n.Condition.GenerateCode(cb)
	cb.AddLine("cmp x0, #0")

	if n.ElseIf != nil || n.ElseBlock != nil {
		cb.AddLine("beq " + labelElse)
	} else {
		cb.AddLine("beq " + labelEnd)
	}

	cb.AddComment("Bloque IF")
	n.IfBlock.GenerateCode(cb)

	if n.ElseIf != nil || n.ElseBlock != nil {
		cb.AddLine("b " + labelEnd)
		cb.AddLine(labelElse + ":")

		if n.ElseIf != nil {
			cb.AddComment("Else if")
			n.ElseIf.GenerateCode(cb)
		} else {
			cb.AddComment("Bloque ELSE")
			n.ElseBlock.GenerateCode(cb)
		}
	}

	cb.AddLine(labelEnd + ":")
}

type ElseIfBlock struct {
	Condition Node       // condición BOOLEANA a evaluar
	Block     *BlockNode // bloque de código a ejecutar si la condición es verdadera

}

/*// string que devulve una representación del nodo if completo para el AST
func (n *IfNode) String() string {
	result := fmt.Sprintf("If(%s) %s", n.Condition.String(), n.ThenBlock.String())

	for _, elseif := range n.ElseIf {
		result += fmt.Sprintf("\nElseIf(%s) %s", elseif.Condition.String(), elseif.Block.String())
	}

	if n.ElseBlock != nil {
		result += fmt.Sprintf("\nElse %s", n.ElseBlock.String())
	}
	return result
}
*/
// nodo para funcion print
type PrintNode struct {
	//lista de expresiones a imprimir
	Expressions []Node // puede ser un literal, una variable, etc.

}

func (p *PrintNode) String() string {
	result := "Print("
	for i, expr := range p.Expressions {
		result += expr.String()
		if i < len(p.Expressions)-1 {
			result += ", "
		}
	}
	result += ")"
	return result

}

type CaseNode struct {
	// valor a comparar en el caso, por ejemplo: case 1:
	Value Node
	// bloque de código a ejecutar si se cumple el caso
	Body *BlockNode
}

type ParameterNode struct {
	Name string // nombre del parámetro
	Type string // tipo (int, float, string, etc.)
}

func (n *ParameterNode) String() string {
	return fmt.Sprintf("%s %s", n.Name, n.Type)
}

func (n *ParameterNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	// No genera código directamente, pero esto permite que implemente Node
}

type SwtchNode struct {
	Condition Node

	Cases       []*CaseNode // lista de casos a evaluar
	DefaultCase *BlockNode  // bloque de código a ejecutar si no se cumple ningún caso
}

func (s *SwtchNode) String() string {
	return "SwitchNode"
}

func (s *SwtchNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio de switch")

	switchEnd := cb.NewLabel("switch_end")
	s.Condition.GenerateCode(cb) // el valor queda en x0
	cb.AddLine("mov x9, x0")     // guardar el valor a comparar en x9

	cb.PushBreakLabel(switchEnd)

	for _, caseNode := range s.Cases {
		caseLabel := cb.NewLabel("case")
		endCaseLabel := cb.NewLabel("end_case")

		caseNode.Value.GenerateCode(cb) // en x0
		cb.AddLine("cmp x9, x0")
		cb.AddLine("bne " + endCaseLabel) // solo entra si coincide
		cb.AddLine(caseLabel + ":")
		caseNode.Body.GenerateCode(cb)
		cb.AddLine("b " + switchEnd) // salir del switch
		cb.AddLine(endCaseLabel + ":")
	}

	if s.DefaultCase != nil {
		cb.AddComment("Bloque default")
		s.DefaultCase.GenerateCode(cb)
	}

	cb.PopBreakLabel()

	cb.AddLine(switchEnd + ":")
	cb.AddComment("Fin del switch")
}

func (c *CaseNode) String() string {
	return "CaseNode"
}

// nodo para bucles for
type ForNode struct {
	Init      Node // inicialización del bucle, por ejemplo: i := 0
	Condition Node // condición del bucle, por ejemplo: i < 10
	Post      Node // actualización del bucle, por ejemplo: i++

	RangeId1 *IdentifierNode // identificador de rango, si es un bucle for range //indice
	RangeId2 *IdentifierNode // expresión de rango, si es un bucle for range //valor

	RangeSlice Node // expresión de rango, si es un bucle for range //slice

	Body *BlockNode // bloque de código a ejecutar en cada iteración del bucle

	Type string // tipo de bucle, puede ser "for", "for range", etc.

}

func (f *ForNode) String() string {
	switch f.Type {
	case "classic":
		return fmt.Sprintf("For(classic: Init=%v; Cond=%v; Post=%v) %s",
			printNode(f.Init), printNode(f.Condition), printNode(f.Post), f.Body.String())
	case "condition":
		return fmt.Sprintf("For(condition: %v) %s", printNode(f.Condition), f.Body.String())
	case "range":
		return fmt.Sprintf("For(range: Id1=%s; Id2=%s; Slice=%s) %s",
			f.RangeId1.String(), f.RangeId2.String(), printNode(f.RangeSlice), f.Body.String())
	default:
		return "For(unknown type) " + f.Body.String()

	}
}
func (n *ForNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio del for")

	// Inicialización
	n.Init.GenerateCode(cb)

	condLabel := cb.NewLabel("cond_for")
	endLabel := cb.NewLabel("end_for")

	cb.AddLine(condLabel + ":")
	n.Condition.GenerateCode(cb)
	cb.AddText("cmp x0, #0")
	cb.AddText(fmt.Sprintf("beq %s", endLabel))

	cb.PushBreakLabel(endLabel)
	n.Body.GenerateCode(cb)
	cb.PopBreakLabel()

	cb.AddText(fmt.Sprintf("b %s", condLabel))
	cb.AddLine(endLabel + ":")

	cb.AddComment("Fin del for")
}

func printNode(n Node) string {
	if n == nil {
		return "nil"
	}
	return n.String()
}

func (p *PrintNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Instrucción print")

	for _, expr := range p.Expressions {
		switch e := expr.(type) {
		case *LiteralNode:
			switch v := e.Value.(type) {
			case string:
				label := cb.NewLabel("str")
				cb.AddData(fmt.Sprintf("%s: .asciz %q", label, v))
				cb.AddText(fmt.Sprintf("ldr x1, =%s", label)) // ← esta es la forma segura
				cb.AddText(fmt.Sprintf("mov x2, #%d", len(v)))
				cb.AddText("mov x0, #1")
				cb.AddText("mov x8, #64")
				cb.AddText("svc #0")

			case int:
				cb.AddComment(fmt.Sprintf("Print literal int: %d", v))
				cb.UsesPrintInt = true
				cb.AddText(fmt.Sprintf("mov x0, #%d", v))
				cb.AddText("bl print_int")

			case bool:
				boolStr := "false"
				if v {
					boolStr = "true"
				}
				label := cb.NewLabel("bool")
				cb.AddData(fmt.Sprintf("%s: .asciz %q", label, boolStr))
				cb.AddText(fmt.Sprintf("adr x1, %s", label))
				cb.AddText(fmt.Sprintf("mov x2, #%d", len(boolStr)))
				cb.AddText("mov x0, #1")
				cb.AddText("mov x8, #64")
				cb.AddText("svc #0")

			default:
				cb.AddComment("⚠️ tipo no soportado en print")
			}

		case *IdentifierNode:
			cb.AddComment("Print variable int: " + e.Name)
			cb.UsesPrintInt = true
			cb.AddText(fmt.Sprintf("ldr x0, =_%s", e.Name)) // ← Aquí está el fix
			cb.AddText("ldr x0, [x0]")
			cb.AddText("bl print_int")

		default:
			cb.AddComment("⚠️ solo se permite print de literales o identificadores por ahora")
		}
	}

	// Salto de línea automático
	label := cb.NewLabel("newline")
	cb.AddData(fmt.Sprintf("%s: .asciz \"\\n\"", label))
	cb.AddText("ldr x1, =" + label)
	cb.AddText("mov x2, #1")
	cb.AddText("mov x0, #1")
	cb.AddText("mov x8, #64")
	cb.AddText("svc #0")

	cb.AddComment("Fin print")
}

// nodo para declaracion de funciones
// Nodo que representa una declaración de función
type FunctionDeclarationNode struct {
	Name       string           // nombre de la función
	Params     []*ParameterNode // lista de parámetros de la función
	ReturnType string           // tipo de retorno de la función, puede ser "" si no tiene retorno
	Body       *BlockNode       // bloque de código que representa el cuerpo de la función

}

type FunctionParam struct {
	Name string // nombre del parámetro
	Type string // tipo del parámetro

}

func (n *FunctionDeclarationNode) String() string {
	paramList := ""
	for i, p := range n.Params {
		paramList += fmt.Sprintf("%s %s", p.Name, p.Type)
		if i > 0 {
			paramList += ", "
		}
		paramList += fmt.Sprintf("%s %s", p.Name, p.Type)
	}
	return fmt.Sprintf("Function(%s(%s) %s) %s", n.Name, paramList, n.ReturnType, n.Body.String())
}

// nodo para llamada a funciones
type FunctionCallNode struct {
	Name      string // nombre de la función a llamar
	Arguments []Node // lista de argumentos a pasar a la función, puede ser un literal, una variable, etc.

}

func (f *FunctionCallNode) String() string {
	args := ""
	for i, arg := range f.Arguments {
		args += arg.String()
		if i < len(f.Arguments)-1 {
			args += ", "
		}
	}

	/*
		for i, arg := range n.Arguments {
			if i > 0 {
				args += ", "
			}
			args += arg.String()
		}
	*/

	return fmt.Sprintf("FunctionCall(%s(%s))", f.Name, args)

}
func (n *FunctionCallNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment(fmt.Sprintf("Llamada a función: %s", n.Name))

	// Generar código para cada argumento y colocarlos en x0, x1, ..., sin redundancia
	for i, arg := range n.Arguments {
		arg.GenerateCode(cb) // El valor queda en x0

		if i == 0 {
			// Si es el primer argumento, ya está en x0
			continue
		}
		cb.AddText(fmt.Sprintf("mov x%d, x0", i)) // Mover x0 → x1, x2, ...
	}

	// Llamar a la función
	cb.AddText(fmt.Sprintf("bl %s", n.Name))

	// Resultado queda en x0
}

// exprsiones
func (f *BinaryOperationNode) String() string {
	return fmt.Sprintf("BinaryOperation(%s %s %s)", f.Left.String(), f.Operator, f.Right.String())
}

func (n *BinaryOperationNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio operación binaria: " + n.Operator)

	n.Left.GenerateCode(cb)  // resultado queda en x0
	cb.AddText("mov x9, x0") // guarda resultado izquierdo en x9

	n.Right.GenerateCode(cb)  // resultado derecho queda en x0
	cb.AddText("mov x10, x0") // guarda resultado derecho en x10

	switch n.Operator {
	case "+":
		cb.AddText("add x0, x9, x10")
	case "-":
		cb.AddText("sub x0, x9, x10")
	case "*":
		cb.AddText("mul x0, x9, x10")
	case "/":
		cb.AddText("sdiv x0, x9, x10")
	case "%":
		cb.AddText("udiv x11, x9, x10")     // x11 = x9 / x10
		cb.AddText("msub x0, x11, x10, x9") // x0 = x9 - (x11 * x10) = x9 % x10

	// Comparaciones
	case "==":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, eq") // x0 = 1 si x9 == x10
	case "!=":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, ne")
	case "<":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, lt")
	case "<=":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, le")
	case ">":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, gt")
	case ">=":
		cb.AddText("cmp x9, x10")
		cb.AddText("cset x0, ge")

	default:
		cb.AddComment("Operador binario no soportado: " + n.Operator)
	}

	cb.AddComment("Fin operación binaria: " + n.Operator)
}

func (f *UnaryOperationNode) String() string {
	return fmt.Sprintf("UnaryOperation(%s %v)", f.Operator, f.Operand.String())
}

func (n *UnaryOperationNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Operación unaria: " + n.Operator)

	n.Operand.GenerateCode(cb)

	switch n.Operator {
	case "-":
		cb.AddLine("neg x0, x0") // negación numérica
	case "!":
		// Lógica booleana NOT: si x0 == 0 → 1, si x0 != 0 → 0
		cb.AddLine("cmp x0, #0")
		cb.AddLine("cset x0, eq")
	default:
		cb.AddComment("⚠️ Operador unario no soportado: " + n.Operator)
	}
}

// sentencias de retorno
type BreakNode struct {
}

func (b *BreakNode) String() string {
	return "BreakNode"
}

func (b *BreakNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Instrucción break")
	cb.AddLine("b " + cb.CurrentBreakLabel())

}

type ContinueNode struct {
}

type ReturnNode struct {
	Value Node // valor a retornar, puede ser nil si no se retorna nada
}

func (r *ReturnNode) String() string {
	return "ReturnNode"
}

func (n *ReturnNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio return")
	if n.Value != nil {
		n.Value.GenerateCode(cb) // deja 42 en x0
	} else {
		cb.AddText("mov x0, #0")
	}
	cb.AddText("ldp x29, x30, [sp], #16")
	cb.AddText("ret")
	cb.AddComment("Fin return")
}

// nodos para los structs
type StructDeclarationNode struct {
	Name   string         // nombre del struct
	Fields []*StructField // lista de campos del struct, cada campo es una declaración de variable
}

type StructField struct {
	Name string // nombre del campo
	Type string // tipo del campo
}

type StructInstanceNode struct {
	InstanceName string // nombre del tipo de struct
	StrcutName   string // nombre del struct
	Values       []*DeclarationNode
}

func (s *StructDeclarationNode) String() string {
	return "StructDeclaration(" + s.Name + ")"
}
func (s *StructInstanceNode) String() string {
	return "StructInstance(" + s.InstanceName + ":" + s.StrcutName + ")"
}

// nodos para creacion y acceso a colecciones
type SliceAccessNode struct {
	Collection Node // colección a la que se accede, puede ser un slice, un array, etc.
	Index      Node // índice al que se accede, puede ser un literal, una variable, etc.
}

func (n *SliceAccessNode) String() string {
	return n.Collection.String() + "[...]"
}

type MatrixAccessNode struct {
	Collection Node // colección a la que se accede, puede ser una matriz, un slice de slices, etc.
	RowIndex   Node // índice de la fila a la que se accede, puede ser un literal, una variable, etc.
	ColIndex   Node // índice de la columna a la que se accede, puede ser un literal, una variable, etc.
}

type SliceAssignmentNode struct {
	Colection Node // colección a la que se asigna, puede ser un slice, un array, etc.
	Index     Node // índice al que se asigna, puede ser un literal, una variable, etc.
	Value     Node // valor a asignar, puede ser un literal, una variable, etc.
}

func (n *SliceAssignmentNode) String() string {
	return "SliceAssignment(" + n.Colection.String() + "[" + n.Index.String() + "] = ...)"
}

func (s *SliceAssignmentNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("⚠️ SliceAssignment ignorado (no se evalúa en este proyecto)")
}

func (s *SliceAccessNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("⚠️ SliceAccess ignorado (no soportado en este proyecto)")
}

type MatrixAssignmentNode struct {
	Coollection Node // colección a la que se asigna, puede ser una matriz, un slice de slices, etc.
	RowIndex    Node // índice de la fila a la que se asigna, puede ser un literal, una variable, etc.
	ColIndex    Node // índice de la columna a la que se asigna, puede ser un literal, una variable, etc.
	Value       Node // valor a asignar, puede ser un literal, una variable, etc.
}

// creacion de slices
type SliceLiteralNode struct {
	Elements []Node // lista de elementos del slice, puede ser un literal, una variable, etc.
}

// funcion que nos ayudara a implementar las funciones incorporadas
type BuiltinFunCallNode struct {
	FuncName  string // nombre de la función incorporada a llamar, por ejemplo "len", "append", etc.
	Arguments []Node // lista de argumentos a pasar a la función, puede ser un literal, una variable, etc.
}

func (b *BuiltinFunCallNode) String() string {
	return "BuiltinFunCallNode"
}

func (b *BuiltinFunCallNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Función embebida: " + b.FuncName)

	switch b.FuncName {
	case "len":
		// Suponemos que solo se permite len(slice) o len(string)
		b.Arguments[0].GenerateCode(cb)
		cb.AddLine("# Aquí iría código específico para obtener longitud de string o slice")
		// Deja valor en x0 (debes definir comportamiento según el tipo)

	case "typeof":
		b.Arguments[0].GenerateCode(cb)
		cb.AddLine("# typeof no implementado: requiere metainformación de tipos")

	case "atoi":
		b.Arguments[0].GenerateCode(cb)
		cb.AddLine("# atoi no implementado: requiere convertir cadena a int")

	default:
		cb.AddComment("Función embebida no soportada aún: " + b.FuncName)
	}
}

type SliceDeclarationNode struct {
	Name     string
	Type     string
	Elements []Node
	Mutable  bool
}

func (n *SliceDeclarationNode) String() string {
	return "SliceDeclaration(" + n.Name + ")"
}

// cliclos for
type ForClassicNode struct {
	Init      Node
	Condition Node
	Update    Node
	Body      *BlockNode
}

func (f *ForClassicNode) String() string {
	return "ForClassicNode"
}

func (f *ForClassicNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Inicio for clásico")

	startLabel := cb.NewLabel("for_start")
	endLabel := cb.NewLabel("for_end")

	// Inicialización
	if f.Init != nil {
		cb.AddComment("Inicialización del for")
		f.Init.GenerateCode(cb)
	}

	cb.AddLine(startLabel + ":")

	// Condición
	if f.Condition != nil {
		cb.AddComment("Condición del for")
		f.Condition.GenerateCode(cb)
		cb.AddLine("cmp x0, #0")
		cb.AddLine("beq " + endLabel)
	}

	// Cuerpo del for
	cb.AddComment("Cuerpo del for")
	cb.PushBreakLabel(endLabel)
	f.Body.GenerateCode(cb)
	cb.PopBreakLabel()

	// Incremento
	if f.Update != nil {
		cb.AddComment("Actualización del for")
		f.Update.GenerateCode(cb)
	}

	cb.AddLine("b " + startLabel)
	cb.AddLine(endLabel + ":")
	cb.AddComment("Fin for clásico")
}

func (c *ContinueNode) NodeType() string {
	return "ContinueNode"
}

func (n *ContinueNode) String() string {
	return "Continue"
}

func (c *ContinueNode) GenerateCode(cb *generadorcodigo.CodeBuilder) {
	cb.AddComment("Instrucción continue")
	cb.AddLine("b loop_start") // ← etiqueta del comienzo del ciclo
}

//slice
