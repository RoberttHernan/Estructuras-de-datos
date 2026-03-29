//vamos a generar un visitor para recorrer el AST mediante el ctx := context.Background() y vamos a transformar los nodos que ya creamos

package visitor

import (
	"fmt"
	"proyecto1/ast"
	"proyecto1/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

//implementaos el visitor y construimos el ast mediante los nodos que ya tenemos

type ASTBuilder struct {
	parser.BasegramaticaVisitor        // hereda de BasegramaticaVisitor para implementar los métodos de visita
	CurrentScope                string // para manejar el contexto actual del scope
}

func NewASTBuilder() *ASTBuilder {
	fmt.Println("Creando ASTBuilder")
	return &ASTBuilder{
		BasegramaticaVisitor: parser.BasegramaticaVisitor{},
	}
}
func (v *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	//fmt.Println("Visitando el arbol")
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

// para regla prog inicial
func (v *ASTBuilder) VisitProg(ctx *parser.ProgContext) interface{} {
	fmt.Println("Visitando Prog")
	v.CurrentScope = "global" // inicia en global

	nodes := []ast.Node{}
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			result := v.Visit(node)
			if result != nil {
				nodes = append(nodes, result.(ast.Node))
			}
		}
	}
	return &ast.Prog{Elementos: nodes}
}

// metodo para visitar declaraciones de funcion
// dec1
func (v *ASTBuilder) VisitDecFunc1(ctx *parser.DecFunc1Context) interface{} {
	fmt.Println("Visitando Declaration funcion 1")
	name := ctx.IDENTIFICADOR().GetText()

	var params []*ast.ParameterNode

	if ctx.Lista_parametros() != nil {
		result := v.Visit(ctx.Lista_parametros())
		if list, ok := result.([]*ast.ParameterNode); ok {
			params = list
		}

	} else {
		params = []*ast.ParameterNode{} // Si no hay parámetros, inicializamos como slice vacío
	}

	//tipo de retorno
	var returnType string
	if ctx.GetTipoRetorno() != nil {
		returnType = ctx.GetTipoRetorno().GetText()
	}

	prev := v.CurrentScope
	v.CurrentScope = name
	defer func() { v.CurrentScope = prev }()
	//bñloque de sentencias
	block := v.Visit(ctx.Bloque()).(*ast.BlockNode)
	return &ast.FunctionDeclarationNode{
		Name:       name,
		Params:     params,
		ReturnType: returnType,
		Body:       block,
	}
}

// dec2
func (v *ASTBuilder) VisitDecFunc2(ctx *parser.DecFunc2Context) interface{} {
	fmt.Println("Visitando Declaration2")

	name := ctx.IDENTIFICADOR().GetText()

	var params []*ast.ParameterNode

	if ctx.Lista_parametros() != nil {
		params = v.Visit(ctx.Lista_parametros()).([]*ast.ParameterNode)

	} else {
		params = []*ast.ParameterNode{} // Si no hay parámetros, inicializamos como slice vacío
	}

	//tipo de retorno
	var returnType string
	if ctx.GetTipoRetorno() != nil {
		returnType = ctx.GetTipoRetorno().GetText()
	}

	prev := v.CurrentScope
	v.CurrentScope = name
	defer func() { v.CurrentScope = prev }()

	body := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.FunctionDeclarationNode{
		Name:       name,
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}

}

// Llamada a funcion
func (v *ASTBuilder) VisitLlamadafuncion(ctx *parser.LlamadafuncionContext) interface{} {
	name := ctx.IDENTIFICADOR().GetText()

	var args []ast.Node
	if ctx.Lista_parametros_func() != nil {
		result := v.Visit(ctx.Lista_parametros_func())
		if list, ok := result.([]ast.Node); ok {
			args = list
		}
	}

	switch name {
	case "len", "append", "cap", "make", "copy", "Atoi", "parseFloat", "typeOf":
		fmt.Printf("➡️ Llamando a función embebida '%s' con %d argumentos\n", name, len(args))
		return &ast.BuiltinFunCallNode{
			FuncName:  name,
			Arguments: args,
		}
	default:
		fmt.Printf("➡️ Llamando a función normal '%s'\n", name)
		return &ast.FunctionCallNode{
			Name:      name,
			Arguments: args,
		}
	}

}

// expresion llamada funcion
func (v *ASTBuilder) VisitExprLlamadaFuncion(ctx *parser.ExprLlamadaFuncionContext) interface{} {
	fmt.Println("Visitando llamada a función como ex")

	name := ctx.Llamadafuncion().IDENTIFICADOR().GetText()
	var args []ast.Node

	if ctx.Llamadafuncion().Lista_parametros_func() != nil {
		result := v.Visit(ctx.Llamadafuncion().Lista_parametros_func())
		if list, ok := result.([]ast.Node); ok {
			args = list
		}
	}

	switch name {
	case "len", "append", "cap", "make", "copy", "Atoi", "parseFloat", "typeOf":
		fmt.Printf("➡️ Llamando a función embebida '%s' con %d argumentos\n", name, len(args))
		return &ast.BuiltinFunCallNode{
			FuncName:  name,
			Arguments: args,
		}
	default:
		fmt.Printf("➡️ Llamando a función normal '%s'\n", name)
		return &ast.FunctionCallNode{
			Name:      name,
			Arguments: args,
		}
	}

}

func (v *ASTBuilder) VisitBloque(ctx *parser.BloqueContext) interface{} {
	fmt.Println("Visitando Bloque")
	staments := []ast.Node{}

	for _, stmt := range ctx.AllLista_sentencias() {
		result := v.Visit(stmt)
		if result != nil {
			if node, ok := result.([]ast.Node); ok {
				for _, sub := range node {
					fmt.Printf("🧱 Sentencia agregada al bloque: %T\n", sub)
				}
				staments = append(staments, node...)
			}
		}
	}

	return &ast.BlockNode{
		Statements: staments,
	}
}

// motodo para lista de sentencias
func (v *ASTBuilder) VisitLista_sentencias(ctx *parser.Lista_sentenciasContext) interface{} {
	fmt.Println("Visitando Lista de Sentencias")

	nodes := []ast.Node{}
	for _, stmt := range ctx.AllSentencia() {
		if node, ok := stmt.(antlr.ParseTree); ok {
			result := v.Visit(node)
			if astNode, ok := result.(ast.Node); ok {
				nodes = append(nodes, astNode)
			}
		}
	}
	return nodes

}

// metodo para expresiones
// identificador

// funcion print
func (v *ASTBuilder) VisitSentenciaPrintln(ctx *parser.SentenciaPrintlnContext) interface{} {

	fmt.Println("Visitando Print")

	var expressions []ast.Node

	if ctx.Lista_expresion() != nil {
		result := v.Visit(ctx.Lista_expresion())

		if lista, ok := result.([]ast.Node); ok {
			expressions = lista
		}
	}
	return &ast.PrintNode{
		Expressions: expressions,
	}
}

// visit Sentencia
func (v *ASTBuilder) VisitSentencia(ctx *parser.SentenciaContext) interface{} {
	fmt.Println("Visitando Sentencia")
	for _, child := range ctx.GetChildren() {
		if sub, ok := child.(antlr.ParseTree); ok {
			result := v.Visit(sub)
			if node, ok := result.(ast.Node); ok {
				fmt.Printf("🔎 Sentencia reconocida: %T\n", node)
				return ast.Node(node) // Retorna el nodo AST correspondiente
			}
		}
	}

	//fmt.Println("⚠️ No se reconoció ninguna sub-sentencia.")
	return nil
}

// lista de parametros
func (v *ASTBuilder) VisitLista_parametros(ctx *parser.Lista_parametrosContext) interface{} {
	fmt.Println("Visitando Lista de Parametros")
	var params []*ast.ParameterNode

	for _, p := range ctx.AllParametro() {
		param := v.Visit(p).(*ast.ParameterNode)
		params = append(params, param)
	}
	return params

}

// parametroq
func (v *ASTBuilder) VisitParametro(ctx *parser.ParametroContext) interface{} {
	fmt.Println("Visitando Parametro")

	id := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipo().GetText()

	return &ast.ParameterNode{
		Name: id,
		Type: tipo,
	}
}

// lista expresion
func (v *ASTBuilder) VisitLista_expresion(ctx *parser.Lista_expresionContext) interface{} {
	fmt.Println("Visitando Lista de Expresiones")
	var expressions []ast.Node

	for _, expr := range ctx.AllExpresion() {
		result := v.Visit(expr)
		if node, ok := result.(ast.Node); ok {
			expressions = append(expressions, node)
		}
	}
	return expressions
}

// expresiones
// identificador
func (v *ASTBuilder) VisitExpresionIdentificador(ctx *parser.ExprIdentificadorContext) interface{} {
	fmt.Println("Visitando Expresion Identificador")
	return &ast.IdentifierNode{
		Name: ctx.GetText(),
	}
}

// cadena
func (v *ASTBuilder) VisitExprString(ctx *parser.ExprStringContext) interface{} {
	fmt.Println("Visitando Expresion Cadena")

	value := ctx.CADENA().GetText()

	return &ast.LiteralNode{
		Value: value,
	}
}

// int
func (v *ASTBuilder) VisitExprInt(ctx *parser.ExprIntContext) interface{} {
	fmt.Println("Visitando Expresion Int")
	valueStr := ctx.INT().GetText()

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error converting string to int: %s\n", err)
		return nil
	}
	return &ast.LiteralNode{
		Value: value,
	}

}

// expresiones aritmeticas
// sumas y restas
func (v *ASTBuilder) VisitExprOpAritmetica(ctx *parser.ExprOpAritmeticaContext) interface{} {
	fmt.Println("Visitando Expresion Aritmetica")

	left := v.Visit(ctx.GetLeft()).(ast.Node)
	right := v.Visit(ctx.GetRight()).(ast.Node)

	operator := ctx.GetOp().GetText()

	return &ast.BinaryOperationNode{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

// negacion unaria
func (v *ASTBuilder) VisitExprNegUnaria(ctx *parser.ExprNegUnariaContext) interface{} {
	fmt.Println("Visitando Expresion Negacion Unaria")

	expr := v.Visit(ctx.Expresion()).(ast.Node)

	return &ast.UnaryOperationNode{
		Operator: "-",
		Operand:  expr,
	}
}

// operacion de multiplicacion y division
func (v *ASTBuilder) VisitExprMultDivMod(ctx *parser.ExprMultDivModContext) interface{} {
	fmt.Println("Visitando Expresion Multiplicacion Division Modulo")
	left := v.Visit(ctx.Expresion(0)).(ast.Node)
	right := v.Visit(ctx.Expresion(1)).(ast.Node)
	operator := ctx.GetOp().GetText()
	//izq := v.Visit(ctx.Expresion(0))
	//der := v.Visit(ctx.Expresion(1))

	fmt.Println("left: ", left, "right: ", right, "operacion: ", operator)

	return &ast.BinaryOperationNode{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

// expresion parentesis
func (v *ASTBuilder) VisitExprParentesis(ctx *parser.ExprParentesisContext) interface{} {
	fmt.Println("Visitando Expresion Parentesis")
	expr := v.Visit(ctx.Expresion()).(ast.Node)
	return expr

}

// expresion relacional
func (v *ASTBuilder) VisitExprRelacional(ctx *parser.ExprRelacionalContext) interface{} {
	fmt.Println("Visitando Expresion Relacional")

	leftRaw := v.Visit(ctx.Expresion(0))
	if leftRaw == nil {
		fmt.Println("⚠️ Error: expresión izquierda es nil")
		return nil
	}

	rightRaw := v.Visit(ctx.Expresion(1))
	if rightRaw == nil {
		fmt.Println("⚠️ Error: expresión derecha es nil")
		return nil
	}

	left, ok1 := leftRaw.(ast.Node)
	right, ok2 := rightRaw.(ast.Node)

	if !ok1 || !ok2 {
		fmt.Println("⚠️ Error: conversión de interface a ast.Node falló")
		return nil
	}

	operator := ctx.GetOp().GetText()

	return &ast.BinaryOperationNode{
		Left:     left,
		Right:    right,
		Operator: operator,
	}

}

// expresion logica
func (v *ASTBuilder) VisitExprLogica(ctx *parser.ExprLogicaContext) interface{} {
	fmt.Println("Visitando Expresion Logica")
	left := v.Visit(ctx.GetLeft()).(ast.Node)
	right := v.Visit(ctx.GetRight()).(ast.Node)
	operator := ctx.GetOp().GetText()

	return &ast.BinaryOperationNode{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

// expresion not
func (v *ASTBuilder) VisitExprNot(ctx *parser.ExprNotContext) interface{} {
	fmt.Println("Visitando Expresion Not")
	expr := v.Visit(ctx.Expresion()).(ast.Node)

	return &ast.UnaryOperationNode{
		Operator: "!",
		Operand:  expr,
	}
}

// expresion true
func (v *ASTBuilder) VisitExprBoolTrue(ctx *parser.ExprBoolTrueContext) interface{} {
	fmt.Println("Visitando Expresion Boolean True")
	return &ast.LiteralNode{
		Value: true,
	}
}

// expresion false
func (v *ASTBuilder) VisitExprBoolFalse(ctx *parser.ExprBoolFalseContext) interface{} {
	fmt.Println("Visitando Expresion Boolean False")
	return &ast.LiteralNode{
		Value: false,
	}
}

// expresion rune
func (v *ASTBuilder) VisitExprRune(ctx *parser.ExprRuneContext) interface{} {
	fmt.Println("Visitando Expresion Rune")
	value := ctx.RUNE().GetText()

	// Eliminar las comillas simples
	value = value[1 : len(value)-1]

	return &ast.LiteralNode{
		Value: value,
	}

}

// expresion identificador
func (v *ASTBuilder) VisitExprIdentificador(ctx *parser.ExprIdentificadorContext) interface{} {
	fmt.Println("Visitando Expresion Identificador")
	id := ctx.IDENTIFICADOR().GetText()

	if id == "" {
		fmt.Println("⚠️ Error: Identificador vacío en expresión")
		return nil
	}

	return &ast.IdentifierNode{
		Name: id,
	}
}

// expresion float
func (v *ASTBuilder) VisitExprFloat(ctx *parser.ExprFloatContext) interface{} {
	fmt.Println("Visitando Expresion Float")
	valueStr := ctx.FLOAT().GetText()

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Printf("Error converting string to float: %s\n", err)
		return nil
	}
	return &ast.LiteralNode{
		Value: value,
	}
}

// -------------------------------------asginaciones  de variables-------------------------------------
// dec1
func (v *ASTBuilder) VisitDec1(ctx *parser.Dec1Context) interface{} {
	fmt.Println("Visitando Declaracion de Variable tipo 1")

	var valor ast.Node
	var tipo string

	if ctx.Tipo() == nil {
		//fmt.Println("⚠️ Error: IDENTIFICADOR no encontrado en la declaración")
		//return nil

		//debemos buscar el tipo de expresion que viene con la declaracion para asirnar
		if ctx.Expresion() != nil {
			if result := v.Visit(ctx.Expresion()); result != nil {
				valor = result.(ast.Node)
			}
		}

		// revisamos de que tipo es el resultado
		if _, err := strconv.Atoi(valor.String()); err == nil {
			tipo = "int"
		} else if _, err := strconv.ParseFloat(valor.String(), 64); err == nil {
			tipo = "float"
		} else {
			// Verificamos si es un booleano
			if lit, ok := valor.(*ast.LiteralNode); ok {
				if b, ok := lit.Value.(bool); ok && (b == true || b == false) {
					tipo = "bool"
				} else {
					tipo = "string"
				}
			}
		}
	}

	id := ctx.IDENTIFICADOR().GetText()
	//tipo := ctx.Tipo().GetText()
	mutable := true

	//var valor ast.Node
	if ctx.Expresion() != nil {
		if result := v.Visit(ctx.Expresion()); result != nil {
			valor = result.(ast.Node)
		}
	}

	return &ast.DeclarationNode{
		Identifier: &ast.IdentifierNode{Name: id},
		Type:       tipo,
		Value:      valor,
		Mutable:    mutable,

		Line:  ctx.GetStart().GetLine(),
		Col:   ctx.GetStart().GetColumn(),
		Scope: v.CurrentScope,
	}

}

// dec2
func (v *ASTBuilder) VisitDec2(ctx *parser.Dec2Context) interface{} {
	fmt.Println("Visitando Declaracion de Variable tipo 2 ([:= expr])")
	id := ctx.IDENTIFICADOR().GetText()

	var valor ast.Node
	if result := v.Visit(ctx.Expresion()); result != nil {
		valor = result.(ast.Node)
	}

	return &ast.DeclarationNode{
		Identifier: &ast.IdentifierNode{Name: id},
		Type:       "", // No se especifica tipo en esta declaración
		Value:      valor,
		Mutable:    true,

		Line:  ctx.GetStart().GetLine(),
		Col:   ctx.GetStart().GetColumn(),
		Scope: v.CurrentScope,
	}

}

// dec3 x int = 20 (sin mut, tipo y valor explícitos)
func (v *ASTBuilder) VisitDec3(ctx *parser.Dec3Context) interface{} {
	fmt.Println("Visitando Declaracion de Variable tipo 3")

	id := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipo().GetText()
	mutable := false // No es mutable en este caso

	var valor ast.Node

	if result := v.Visit(ctx.Expresion()); result != nil {
		valor = result.(ast.Node)
	}

	return &ast.DeclarationNode{
		Identifier: &ast.IdentifierNode{Name: id},
		Type:       tipo,
		Value:      valor,
		Mutable:    mutable,

		Line:  ctx.GetStart().GetLine(),
		Col:   ctx.GetStart().GetColumn(),
		Scope: v.CurrentScope,
	}
}

// asignacion de variables x=5, x+=5
// asinacion simple
func (v *ASTBuilder) VisitAsigSimple(ctx *parser.AsigSimpleContext) interface{} {
	fmt.Println("Visitando Asignacion Simple")

	id := ctx.IDENTIFICADOR().GetText()
	valor := v.Visit(ctx.Expresion()).(ast.Node)

	return &ast.AssignmentNode{
		Variable: &ast.IdentifierNode{Name: id},
		Value:    valor,
		Operator: "=", // Asignación simple
	}
}

// asignacion compuesta
func (v *ASTBuilder) VisitAsigSumRest(ctx *parser.AsigSumRestContext) interface{} {
	fmt.Println("Visitando Asignacion Compuesta (Suma o Resta)")

	id := ctx.IDENTIFICADOR().GetText()
	op := ctx.GetOp().GetText() // Puede ser += o -=
	valor := v.Visit(ctx.Expresion()).(ast.Node)

	return &ast.AssignmentNode{
		Variable: &ast.IdentifierNode{Name: id},
		Value:    valor,
		Operator: op, // Asignación compuesta
	}
}

/*
*sentenciaIf: RESIF expresion LLAVEABRE sentencia bloque        #SentIf1 // if con una sola sentencia
           | RESIF expresion bloque RESELSE bloque              #SentIf2// if-else ambos bloques
           | RESIF expresion bloque RESELSE sentenciaIf         #SentIf3 // if con else if recursivo
;
*/

// sentIf1
func (v *ASTBuilder) VisitSentIf1(ctx *parser.SentIf1Context) interface{} {
	fmt.Println("Visitando Sentencia If 1")

	condicion := v.Visit(ctx.Expresion()).(ast.Node)
	ifBlock := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.IfNode{
		Condition: condicion,
		IfBlock:   ifBlock,
		ElseIf:    nil, // No hay else if en este caso
		ElseBlock: nil, // No hay bloque else en este caso
	}
}

// sentIf2
func (v *ASTBuilder) VisitSentIf2(ctx *parser.SentIf2Context) interface{} {
	fmt.Println("Visitando Sentencia If else")
	condicion := v.Visit(ctx.Expresion()).(ast.Node)

	ifBlock := v.Visit(ctx.Bloque(0)).(*ast.BlockNode)
	elseBlock := v.Visit(ctx.Bloque(1)).(*ast.BlockNode)

	return &ast.IfNode{
		Condition: condicion,
		IfBlock:   ifBlock,
		ElseIf:    nil, // No hay else if en este caso
		ElseBlock: elseBlock,
	}
}

// sentIf3
func (v *ASTBuilder) VisitSentIf3(ctx *parser.SentIf3Context) interface{} {
	fmt.Println("Visitando Sentencia If 3")
	condicion := v.Visit(ctx.Expresion()).(ast.Node)

	ifBlock := v.Visit(ctx.Bloque()).(*ast.BlockNode)
	elseIf := v.Visit(ctx.SentenciaIf()).(*ast.IfNode)

	return &ast.IfNode{
		Condition: condicion,
		IfBlock:   ifBlock,
		ElseIf:    elseIf, // Agrega el else if
		ElseBlock: elseIf.ElseBlock,
	}
}

// ciclos for
// for 1 --- RESFOR expresion bloque
func (v *ASTBuilder) VisitSentFor1(ctx *parser.SentFor1Context) interface{} {
	fmt.Println("Visitando Ciclo For 1")

	cond := v.Visit(ctx.Expresion()).(ast.Node)
	forBlock := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.ForNode{
		Init:      nil, // No hay inicialización explícita
		Condition: cond,
		Post:      nil, // No hay actualización explícita
		Body:      forBlock,
	}
}

/*
// SentFor2 RESFOR declaracion_for PUNTO_COMA expresion PUNTO_COMA incremento bloque
func (v *ASTBuilder) VisitSentFor2(ctx *parser.SentFor2Context) interface{} {
	fmt.Println("Visitando Ciclo For 2")

	init := v.Visit(ctx.Declaracion_for()).(*ast.DeclarationNode)
	condition := v.Visit(ctx.Expresion()).(ast.Node)
	post := v.Visit(ctx.Incremento()).(ast.Node)
	body := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.ForNode{
		Init:      init,
		Condition: condition,
		Post:      post,
		Body:      body,
	}
}

// SentFor3 RESFOR IDENTIFICADOR SIGNO_COMA IDENTIFICADOR RESIN IDENTIFICADOR bloque
func (v *ASTBuilder) VisitSentFor3(ctx *parser.SentFor3Context) interface{} {
	fmt.Println("Visitando Ciclo For 3")

	indexvar := v.Visit(ctx.IDENTIFICADOR(0)).(*ast.IdentifierNode)
	valuevar := v.Visit(ctx.IDENTIFICADOR(1)).(*ast.IdentifierNode)
	listName := v.Visit(ctx.IDENTIFICADOR(2)).(*ast.IdentifierNode)

	block := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.ForNode{
		RangeId1:   &ast.IdentifierNode{Name: indexvar.Name}, // Identificador de índice
		RangeId2:   &ast.IdentifierNode{Name: valuevar.Name}, // Identificador de valor
		RangeSlice: &ast.IdentifierNode{Name: listName.Name}, // Identificador del slice
		Body:       block,
	}
}
*/

// --------------------sentencia switch--------------------
// Sentencia switch
func (v *ASTBuilder) VisitSentenciaSwitch(ctx *parser.SentenciaSwitchContext) interface{} {
	fmt.Println("Visitando Sentencia Switch")

	condition := v.Visit(ctx.Expresion()).(ast.Node)

	var cases []*ast.CaseNode

	for _, c := range ctx.AllBloqueCase() {
		result := v.Visit(c)
		if node, ok := result.(*ast.CaseNode); ok {
			cases = append(cases, node)
		}
	}

	var defaultBlock *ast.BlockNode
	if ctx.BloqueDefault() != nil {
		lista := v.Visit(ctx.BloqueDefault().Lista_sentencias())
		stmsts := []ast.Node{}
		if lista != nil {
			stmsts = lista.([]ast.Node)
		}
		defaultBlock = &ast.BlockNode{
			Statements: stmsts,
		}
	}

	return &ast.SwtchNode{
		Condition:   condition,
		Cases:       cases,
		DefaultCase: defaultBlock,
	}

}

// bloque case
func (v *ASTBuilder) VisitBloqueCase(ctx *parser.BloqueCaseContext) interface{} {
	fmt.Println("Visitando Bloque Case")

	expr := v.Visit(ctx.Expresion()).(ast.Node)

	var cuerpo *ast.BlockNode

	if ctx.Lista_sentencias(0) != nil {
		result := v.Visit(ctx.Lista_sentencias(0))
		if stmts, ok := result.([]ast.Node); ok {
			cuerpo = &ast.BlockNode{
				Statements: stmts,
			}
		}
	} else {
		// Si por alguna razón no hay cuerpo, lo inicializamos vacío
		cuerpo = &ast.BlockNode{
			Statements: []ast.Node{},
		}
	}

	return &ast.CaseNode{
		Value: expr,
		Body:  cuerpo,
	}
}

// bloque default
func (v *ASTBuilder) VisitBloqueDefault(ctx *parser.BloqueDefaultContext) interface{} {
	fmt.Println("Visitando Bloque Default")

	var stmsts []ast.Node

	if ctx.Lista_sentencias() != nil {
		result := v.Visit(ctx.Lista_sentencias())
		if result != nil {
			stmsts = result.([]ast.Node)
		}
	}

	return &ast.BlockNode{
		Statements: stmsts,
	}
}

// sentencia de control
// sentencia break
func (v *ASTBuilder) VisitSentBreak(ctx *parser.SentBreakContext) interface{} {
	fmt.Println("Visitando Sentencia Break")
	return &ast.BreakNode{}
}

// sentencia continue
func (v *ASTBuilder) VisitSentContinue(ctx *parser.SentContinueContext) interface{} {
	fmt.Println("Visitando Sentencia Continue")
	return &ast.ContinueNode{}
}

// structs
// declaracion de struct
func (v *ASTBuilder) VisitDeclaracionStruct(ctx *parser.DeclaracionStructContext) interface{} {
	fmt.Println("Visitando Declaracion de Struct")
	name := ctx.IDENTIFICADOR().GetText()
	fields := []*ast.StructField{}

	for _, attr := range ctx.AllAtributosStruct() {
		field := v.Visit(attr)
		if field, ok := field.(*ast.StructField); ok {
			fields = append(fields, field)
		} else {
			fmt.Println("⚠️ Error: No se pudo convertir a DeclarationNode")
		}
	}

	return &ast.StructDeclarationNode{
		Name:   name,
		Fields: fields,
	}
}

// atributos de struct
func (v *ASTBuilder) VisitAtributosStruct(ctx *parser.AtributosStructContext) interface{} {
	fmt.Println("Visitando Atributos de Struct")
	tipo := ctx.Tipo().GetText()
	nombre := ctx.GetNomAtributo().GetText()

	return &ast.StructField{
		Name: nombre,
		Type: tipo,
	}
}

// creacion de instancia de struct
func (v *ASTBuilder) VisitCreacion_instancia_struct(ctx *parser.Creacion_instancia_structContext) interface{} {
	fmt.Println("Visitando Creacion de Instancia de Struct")

	instanceName := ctx.GetNomInstancia().GetText()
	struct_name := ctx.GetNomStruct().GetText()

	decls := []*ast.DeclarationNode{}

	if ctx.ListaCampos() != nil {
		result := v.Visit(ctx.ListaCampos())

		if lista, ok := result.([]*ast.DeclarationNode); ok {
			decls = lista
		}
	}

	return &ast.StructInstanceNode{
		InstanceName: instanceName, // Nombre de la instancia
		StrcutName:   struct_name,  // Nombre del tipo de struct
		Values:       decls,
	}

}

// lista de campos para instancias de struct
func (v *ASTBuilder) VisitLista_campos(ctx *parser.ListaCamposContext) interface{} {
	fmt.Println("Visitando Lista de Campos para Instancia de Struct")

	fields := []*ast.DeclarationNode{}

	for _, fieldCtx := range ctx.AllCampoStruct() {
		result := v.Visit(fieldCtx)

		if decl, ok := result.(*ast.DeclarationNode); ok {
			fields = append(fields, decl)
		}
	}

	return fields
}

// campo de struct
func (v *ASTBuilder) VisitCampoStruct(ctx *parser.CampoStructContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	value := v.Visit(ctx.Expresion()).(ast.Node)

	return &ast.DeclarationNode{
		Identifier: &ast.IdentifierNode{Name: id},
		Value:      value,
		Mutable:    true, // Asumimos que los campos son mutables
	}
}

// slices
// slice vacio
func (v *ASTBuilder) VisitDecSlice1(ctx *parser.DecSlice1Context) interface{} {
	fmt.Println("Visitando Slice Vacio")

	name := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipo().GetText()

	return &ast.SliceDeclarationNode{
		Name:     name,         // Nombre del slice
		Type:     tipo,         // Tipo de los elementos del slice
		Elements: []ast.Node{}, // Slice vacío
		Mutable:  true,         // Asumimos que el slice es mutable

	}

}

// slice con valores
/*func (v *ASTBuilder) VisitDecSlice2(ctx *parser.DecSlice2Context) interface{} {
	fmt.Println("Visitando Declaracion de Slice con valores")

	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()

	name := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipo().GetText()

	var elements []ast.Node
	if ctx.Lista_expresion() != nil {
		result := v.Visit(ctx.Lista_expresion())
		if list, ok := result.([]ast.Node); ok {
			elements = list
		}
	}

	return &ast.SliceDeclarationNode{
		Name:     name,
		Type:     tipo,
		Elements: elements,
		Mutable:  true,

		Line:   line,
		Column: col,
	}
}
*/

// asignacion de slice-----------------------
/*func (v *ASTBuilder) VisitAsigacionAccesoSlice1(ctx *parser.AsigacionAccesoSlice1Context) interface{} {
	fmt.Println("Visitando Acceso a Slice como Expresion")

	collection := &ast.IdentifierNode{Name: ctx.IDENTIFICADOR().GetText()}
	index := v.Visit(ctx.GetIndex()).(ast.Node)

	return &ast.SliceAccessNode{
		Collection: collection,
		Index:      index,
	}
}
*/
/*
func (v *ASTBuilder) VisitAsigacionAccesoSlice(ctx *parser.AsigacionAccesoSliceContext) interface{} {
	 fmt.Println("Visitando Asignación de Slice completa (x = []int{...})")

    name := ctx.IDENTIFICADOR().GetText()
    tipo := ctx.Tipo().GetText()

    var elementos []ast.Node
    if ctx.Lista_expresion() != nil {
        result := v.Visit(ctx.Lista_expresion())
        if list, ok := result.([]ast.Node); ok {
            elementos = list
        }
    }

    return &ast.SliceAssignmentNode{
        Colection: &ast.IdentifierNode{Name: name},
        Type:      tipo,
        Elements:  elementos,
    }
}

// expresion de acceso a slice
func (v *ASTBuilder) VisitAsignacion_acceso_slice1(ctx *parser.Asignacion_acceso_slice1Context) interface{} {
    fmt.Println("Visitando Acceso a Slice como Expresion")

    collection := &ast.IdentifierNode{Name: ctx.IDENTIFICADOR().GetText()}
    index := v.Visit(ctx.GetIndex()).(ast.Node)

    return &ast.SliceAccessNode{
        Collection: collection,
        Index:      index,
    }
}
*/

// acceso a matrices
func (v *ASTBuilder) VisitExprAccesoMatriz(ctx *parser.ExprAccesoMatrizContext) interface{} {
	fmt.Println("Visitando Acceso a Matriz como Expresion")

	collection := &ast.IdentifierNode{Name: ctx.IDENTIFICADOR().GetText()}
	row := v.Visit(ctx.Expresion(0)).(ast.Node)
	col := v.Visit(ctx.Expresion(1)).(ast.Node)

	return &ast.MatrixAccessNode{
		Collection: collection,
		RowIndex:   row,
		ColIndex:   col,
	}
}

// sentencia return
func (v *ASTBuilder) VisitSentReturn(ctx *parser.SentReturnContext) interface{} {
	fmt.Println("Visitando Return")

	if ctx.Expresion() != nil {
		value := v.Visit(ctx.Expresion()).(ast.Node)
		return &ast.ReturnNode{
			Value: value,
		}
	}
	return &ast.ReturnNode{
		Value: nil,
	}

}

// ciclos for
func (v *ASTBuilder) VisitSentForClasico(ctx *parser.SentForClasicoContext) interface{} {
	fmt.Println("Visitando For tipo clásico")

	init := v.Visit(ctx.Sentencia_init()).(ast.Node)
	cond := v.Visit(ctx.Expresion()).(ast.Node)
	update := v.Visit(ctx.Sentencia_update()).(ast.Node)
	body := v.Visit(ctx.Bloque()).(*ast.BlockNode)

	return &ast.ForClassicNode{
		Init:      init,
		Condition: cond,
		Update:    update,
		Body:      body,
	}
}

func (v *ASTBuilder) VisitSentencia_init(ctx *parser.Sentencia_initContext) interface{} {
	if ctx.Declaracion_variable() != nil {
		return v.Visit(ctx.Declaracion_variable())
	}
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())
	}
	return nil
}
func (v *ASTBuilder) VisitSentencia_update(ctx *parser.Sentencia_updateContext) interface{} {
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())
	}
	if ctx.Incremento_variable() != nil {
		return v.Visit(ctx.Incremento_variable())
	}
	return nil
}

func (v *ASTBuilder) VisitIncremento_variable(ctx *parser.Incremento_variableContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	op := ctx.GetOp().GetText() // '++' o '--'

	var operando ast.Node = &ast.IdentifierNode{Name: id}
	var uno = &ast.LiteralNode{Value: 1}

	var operador string
	if op == "++" {
		operador = "+"
	} else {
		operador = "-"
	}

	return &ast.AssignmentNode{
		Variable: &ast.IdentifierNode{Name: id},
		Value: &ast.BinaryOperationNode{
			Left:     operando,
			Right:    uno,
			Operator: operador,
		},
	}
}

// sentencia break
func (v *ASTBuilder) VisitSentenciaBreak(ctx *parser.SentenciaBreakContext) interface{} {
	fmt.Println("Visitando sentencia break")
	return &ast.BreakNode{}
}

func (v *ASTBuilder) VisitLista_parametros_func(ctx *parser.Lista_parametros_funcContext) interface{} {
	fmt.Println("Visitando Lista de Parámetros de Función")

	var args []ast.Node
	for _, expr := range ctx.AllExpresion() {
		result := v.Visit(expr)
		if node, ok := result.(ast.Node); ok {
			args = append(args, node)
		}
	}
	return args
}
