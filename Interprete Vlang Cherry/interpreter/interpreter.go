// Package interpreter implementa un interprete de recorrido de arbol (tree-walking interpreter)
// para el lenguaje V-Lang Cherry.
//
// El interprete opera sobre el AST producido por visitor/ast_builder.go.
// El estado de ejecucion se gestiona mediante una cadena de entornos (Environment)
// que implementa el alcance lexico: cada bloque, funcion o ciclo crea un nuevo
// Environment que apunta al entorno padre.
//
// El control de flujo (return, break, continue) se implementa mediante panic/recover:
//   - panic(ReturnValue{...}) indica un return de funcion
//   - panic("break_signal")    indica un break dentro de un ciclo
//   - panic("continue_signal") indica un continue dentro de un ciclo
//
// Esta es una tecnica comun en interpretes Go de cero-dependencias que evita
// propagar valores de retorno por todos los metodos del recorrido.
package interpreter

import (
	"fmt"
	"io"
	"os"
	"proyecto1/ast"
	"strconv"
	"strings"
)

// Environment representa un ambito de ejecucion (scope).
// Cada llamada a funcion o bloque de control crea un Environment nuevo cuyo
// campo parent apunta al ambito contenedor, formando una cadena de alcance lexico.
type Environment struct {
	Variables map[string]Variable  // variables declaradas en este ambito
	parent    *Environment         // ambito contenedor (nil en el ambito global)
	Functions map[string]*UserFunction // funciones registradas en este ambito
}

// NewEnvironment crea un Environment vacio con el padre indicado.
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		Variables: make(map[string]Variable),
		parent:    parent,
		Functions: make(map[string]*UserFunction),
	}
}

// Variable almacena el valor y los metadatos de una variable en el entorno.
type Variable struct {
	Value   interface{} // valor actual de la variable
	Type    string      // tipo declarado: "int", "float", "string", "bool", etc.
	Mutable bool        // true si fue declarada con 'mut'; false si es inmutable

	Line  int    // linea de declaracion (para la tabla de simbolos)
	Col   int    // columna de declaracion
	Scope string // nombre del ambito donde fue declarada
}

// Interpreter es el evaluador principal del AST.
// Mantiene un puntero al entorno activo (global o el de la funcion en ejecucion)
// y un escritor de salida configurable para redirigir el output a la GUI o consola.
type Interpreter struct {
	global *Environment // entorno activo durante la ejecucion
	Output io.Writer    // destino de salida para print()
}

// Global retorna el entorno global del interprete.
func (i *Interpreter) Global() *Environment {
	return i.global
}

// NewInterpreter crea un Interpreter con un entorno global vacio y salida en stdout.
func NewInterpreter() *Interpreter {
	return &Interpreter{
		global: NewEnvironment(nil),
		Output: os.Stdout,
	}
}

// Evaluate es el punto de entrada del interprete.
// Primero registra todas las declaraciones de funciones en el entorno global,
// luego busca y ejecuta la funcion 'main'.
func (i *Interpreter) Evaluate(node *ast.Prog) {

	for _, smt := range node.Elementos {
		if fn, ok := smt.(*ast.FunctionDeclarationNode); ok {
			fmt.Println("📘 Registrando función:", fn.Name)
			i.global.Functions[fn.Name] = &UserFunction{
				Params: fn.Params,
				Body:   fn.Body,
			}
		}
	}

	// 2. Ejecutar main si existe
	if mainFn, ok := i.global.Functions["main"]; ok {
		fmt.Printf("🔍 main tiene %d sentencias\n", len(mainFn.Body.Statements))
		i.Eval(mainFn.Body)
	} else {
		panic("❌ No se encontró la función 'main'")
	}
}

// Eval evalua un nodo individual del AST y retorna su valor.
// Es el nucleo del interprete: cada tipo de nodo tiene su caso en el switch.
// El control de flujo no-local (return, break, continue) se senaliza con panic.
func (i *Interpreter) Eval(node ast.Node) interface{} {

	switch n := node.(type) {
	case *ast.PrintNode:
		i.evalPrint(n)
		return nil

	case *ast.FunctionDeclarationNode:
		fmt.Println("Registrando función:", n.Name)
		i.global.Functions[n.Name] = &UserFunction{
			Params: n.Params,
			Body:   n.Body,
		}
		return nil
	case *ast.LiteralNode:
		return n.Value
	case *ast.IdentifierNode:
		fmt.Printf("Identificador: %s\n", n.Name)
		val, err := i.global.Get(n.Name)
		if err != nil {
			panic(err.Error())
		}
		return val
	case *ast.BlockNode:
		oldEnv := i.global
		i.global = NewEnvironment(oldEnv)

		defer func() {
			if r := recover(); r != nil {
				i.global = oldEnv // restaurar siempre

				// retransmitir ReturnValue o errores
				panic(r)
			}
			i.global = oldEnv
		}()

		for _, stmt := range n.Statements {
			fmt.Printf("🔄 Ejecutando stmt en bloque: %T\n", stmt)
			i.Eval(stmt)
		}
		return nil

	case *ast.FunctionCallNode:
		fmt.Printf("➡️ Llamando a functionCallNode '%s' con %d argumentos\n", n.Name, len(n.Arguments))

		fn, ok := i.global.GetFunction(n.Name)
		if !ok {
			panic("❌ Función no definida: " + n.Name)
		}

		if len(n.Arguments) != len(fn.Params) {
			panic("❌ Número incorrecto de argumentos")
		}

		// Crear nuevo entorno para la función
		newEnv := NewEnvironment(i.global)

		// Asignar argumentos a los parámetros
		for idx := range n.Arguments {
			val := i.Eval(n.Arguments[idx])
			paramName := fn.Params[idx].Name
			paramType := fn.Params[idx].Type

			fmt.Printf("🧩 Asignando argumento '%v' a parámetro '%s'\n", val, paramName)

			newEnv.Variables[paramName] = Variable{
				Value:   val,
				Type:    paramType,
				Mutable: false,
			}
		}

		// Guardar entorno anterior
		old := i.global
		i.global = newEnv

		var result interface{}
		defer func() {
			i.global = old // Restaurar entorno anterior
			if r := recover(); r != nil {
				if ret, ok := r.(ReturnValue); ok {
					result = ret.Value
				} else {
					panic(r)
				}
			}
		}()

		// Ejecutar el cuerpo de la función
		result = i.Eval(fn.Body)
		fmt.Printf("⚠️ Función '%s' terminó sin ejecutar return\n", n.Name)

		return result

	case *ast.BuiltinFunCallNode:
		fmt.Println("Ejecutando función embebida:", n.FuncName)

		// Evalúa los argumentos
		var args []interface{}
		for _, arg := range n.Arguments {
			args = append(args, i.Eval(arg))
		}

		switch n.FuncName {
		case "len":
			if len(args) != 1 {
				panic("❌ len() espera exactamente 1 argumento")
			}
			switch v := args[0].(type) {
			case string:
				return len(v)
			case []interface{}:
				return len(v)
			default:
				panic("❌ len() solo funciona con strings o slices")
			}

		case "typeOf":
			if len(args) != 1 {
				panic("❌ typeOf() espera exactamente 1 argumento")
			}
			switch args[0].(type) {
			case int:
				return "int"
			case float64:
				return "float64"
			case bool:
				return "bool"
			case string:
				return "string"
			default:
				return "unknown"
			}

		case "Atoi":
			if len(args) != 1 {
				panic("❌ Atoi() espera un string")
			}
			if str, ok := args[0].(string); ok {
				val, err := strconv.Atoi(str)
				if err != nil {
					panic("❌ Error en Atoi: " + err.Error())
				}
				return val
			}
			panic("❌ Atoi() requiere string como argumento")

		default:
			panic("❌ Función embebida no soportada: " + n.FuncName)
		}

	case *ast.BinaryOperationNode:
		left := i.Eval(n.Left)
		right := i.Eval(n.Right)

		switch n.Operator {
		case "+":
			return sumar(left, right)
		case "-":
			return restar(left, right)
		case "*":
			return multiplicar(left, right)
		case "/":
			//fmt.Println("-----------------Dividiendo:", left, "y", right, " pos: ")
			//if right == 0 || right == 0.0 {
			//	//se guarda el error en una lista y ya no entra el la funcion dividir
			//	return nil
			//}
			return dividir(left, right)
		case "==":
			return left == right
		case "!=":
			return left != right
		case "&&":
			return left.(bool) && right.(bool)
		case "||":
			return left.(bool) || right.(bool)
		case "<":
			return menor(left, right)
		case "<=":
			return menorIgual(left, right)
		case ">":
			return mayor(left, right)
		case ">=":
			return mayorIgual(left, right)
		case "%":
			return modulo(left, right)
		default:
			panic("Operador no soportado: " + n.Operator)
		}

	case *ast.AssignmentNode:
		fmt.Printf("Asignación: %s %s ...\n", n.Variable.Name, n.Operator)
		value := i.Eval(n.Value)

		switch n.Operator {
		case "=":
			if err := i.global.Assign(n.Variable.Name, value); err != nil {
				panic(err.Error())
			}
			return nil
		case "+=", "-=":
			currentVal, err := i.global.Get(n.Variable.Name)
			if err != nil {
				panic(err.Error())
			}
			var result interface{}
			switch n.Operator {
			case "+=":
				result = sumar(currentVal, value)
			case "-=":
				result = restar(currentVal, value)
			}
			if err := i.global.Assign(n.Variable.Name, result); err != nil {
				panic(err.Error())
			}
			return nil

		default:
			panic("⚠️ Operador de asignación no soportado: " + n.Operator)
		}
	case *ast.IfNode:
		fmt.Println("Evaluando IfNode")
		condition := i.Eval(n.Condition)
		fmt.Printf("🧪 Condición del if: %v\n", condition)

		if b, ok := condition.(bool); ok {
			if b {
				return i.Eval(n.IfBlock)
			} else if n.ElseIf != nil {
				return i.Eval(n.ElseIf) // else if como otro if
			} else if n.ElseBlock != nil {
				return i.Eval(n.ElseBlock)
			}
		} else {
			panic("Condición en if no es booleana")
		}
		return nil
	case *ast.UnaryOperationNode:
		operand := i.Eval(n.Operand)

		switch n.Operator {
		case "-":
			switch v := operand.(type) {
			case int:
				return -v
			case float64:
				return -v
			default:
				panic("❌ Error: operador '-' solo se puede usar con números")
			}
		case "!":
			if b, ok := operand.(bool); ok {
				return !b
			}
			panic("❌ Error: operador '!' solo se puede usar con booleanos")
		default:
			panic("❌ Operador unario no soportado: " + n.Operator)
		}
	case *ast.ForNode:
		for {
			cond := i.Eval(n.Condition)
			if b, ok := cond.(bool); ok {
				if !b {
					break
				}
			} else {
				panic("❌ Condición en for no booleana")
			}

			// Ejecutar cuerpo del for y capturar break
			broken := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						switch r {
						case "break_signal":
							broken = true
						case "continue_signal":
							// simplemente salta a la siguiente iteración
						default:
							panic(r)
						}
					}
				}()
				i.Eval(n.Body)
			}()

			if broken {
				break
			}
		}
		return nil

	case *ast.SwtchNode:
		match := i.Eval(n.Condition)

		for _, caseNode := range n.Cases {
			caseVal := i.Eval(caseNode.Value)
			if match == caseVal {
				defer func() {
					if r := recover(); r != nil {
						if r != "break_signal" {
							panic(r) // solo ignora break
						}
					}
				}()
				return i.Eval(caseNode.Body)
			}

		}

		if n.DefaultCase != nil {
			defer func() {
				if r := recover(); r != nil && r != "break_signal" {
					panic(r)
				}
			}()
			return i.Eval(n.DefaultCase)
		}

		return nil
	case *ast.ReturnNode:
		fmt.Println("Evaluando ReturnNode")

		val := i.Eval(n.Value)
		// Forzar salida con valor
		fmt.Printf("📤 Valor retornado: %v\n", val)
		panic(ReturnValue{Value: val})

	case *ast.ForClassicNode:
		newEnv := NewEnvironment(i.global)
		oldEnv := i.global
		i.global = newEnv

		defer func() {
			if r := recover(); r != nil {
				i.global = oldEnv
				panic(r)
			}
			i.global = oldEnv
		}()

		for {
			cond := i.Eval(n.Condition)
			if b, ok := cond.(bool); ok && !b {
				break
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						if r != "break_signal" && r != "continue_signal" {
							panic(r)
						}
					}
				}()
				i.Eval(n.Body)
			}()

			if n.Update != nil {
				i.Eval(n.Update)
			}
		}
		return nil

	case *ast.BreakNode:
		panic("break_signal") // señal especial para detener el ciclo
	case *ast.DeclarationNode:

		fmt.Printf("📦 Declaración de variable: %s\n", n.Identifier.Name)
		name := n.Identifier.Name
		value := i.Eval(n.Value)
		varType := n.Type

		if varType == "" {
			varType = inferType(value) // Inferir tipo si no se especifica
		}

		fmt.Printf("🧾 Registrando símbolo: %s tipo=%s scope=%s línea=%d\n",
			name, varType, n.Scope, n.Line)

		err := i.global.DeclareFull(
			name,
			value,
			varType,
			n.Mutable,
			n.Line,
			n.Col,
			n.Scope,
		)
		if err != nil {
			panic(fmt.Sprintf("❌ Error al declarar variable '%s': %s", name, err))
		}
		return nil

	case *ast.ContinueNode:
		fmt.Println("🟡 Continue ejecutado")
		panic("continue_signal") // señal especial para saltar a la siguiente iteración

	/*case *ast.SliceDeclarationNode:
	fmt.Printf("📦 Declarando slice: %s\n", n.Name)

	// Evaluar los elementos si vienen con valores
	var elementos []interface{}
	for _, elem := range n.Elements {
		elementos = append(elementos, i.Eval(elem))
	}

	err := i.global.Declare(n.Name, elementos, "slice", n.Mutable, )
	if err != nil {
		panic(fmt.Sprintf("❌ Error al declarar slice '%s': %s", n.Name, err))
	}
	return nil
	*/
	case *ast.SliceAccessNode:
		idNode, ok := n.Collection.(*ast.IdentifierNode)
		if !ok {
			panic("❌ Error: la colección no es un identificador")
		}

		sliceRaw, err := i.global.Get(idNode.Name)
		if err != nil {
			panic(err.Error())
		}

		slice, ok := sliceRaw.([]interface{})
		if !ok {
			panic("❌ Error: variable no es un slice")
		}

		indexVal := i.Eval(n.Index)
		index, ok := indexVal.(int)
		if !ok {
			panic("❌ Error: el índice no es entero")
		}

		if index < 0 || index >= len(slice) {
			panic("❌ Error: índice fuera de rango")
		}

		return slice[index]

	case *ast.SliceAssignmentNode:
		idNode, ok := n.Colection.(*ast.IdentifierNode)
		if !ok {
			panic("❌ Error: la colección no es un identificador")
		}

		sliceVar, ok := i.global.lookup(idNode.Name)
		if !ok {
			panic(fmt.Sprintf("❌ Slice '%s' no declarado", idNode.Name))
		}

		slice, ok := sliceVar.Value.([]interface{})
		if !ok {
			panic("❌ Error: variable no es un slice")
		}

		indexVal := i.Eval(n.Index)
		index, ok := indexVal.(int)
		if !ok {
			panic("❌ Índice no es entero")
		}

		if index < 0 || index >= len(slice) {
			panic("❌ Índice fuera de rango en asignación")
		}

		newVal := i.Eval(n.Value)
		slice[index] = newVal

		err := i.global.Assign(idNode.Name, slice)
		if err != nil {
			panic(err.Error())
		}

		return nil

	default:
		fmt.Printf("No se puede evaluar el nodo de tipo %T\n", n)
		return nil
	}
}

// ReturnValue es el tipo usado como valor de panic para implementar 'return'.
// El caso FunctionCallNode captura este tipo con recover() para obtener el
// valor de retorno de la funcion invocada.
type ReturnValue struct {
	Value interface{}
}

// evalprint evalua nodos de tipo PrintNode
func (i *Interpreter) evalPrint(node *ast.PrintNode) {
	var salida []interface{}
	for _, value := range node.Expressions {
		val := i.Eval(value)
		salida = append(salida, val)
	}
	i.Println(salida...) // Solo esto se redirige a la consola de la UI
}

// funcnion para inferir tipos
func inferType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func sumar(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a + b
		case float64:
			return float64(a) + b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a + float64(b)
		case float64:
			return a + b
		}
	case string:
		switch b := b.(type) {
		case string:
			a = strings.ReplaceAll(a, "\"", "")
			b = strings.ReplaceAll(b, "\"", "")
			return a + b
		default:
			panic("❌ Error en suma: no se puede concatenar string con tipo no string")
		}
	}
	panic("❌ Error en suma: tipos incompatibles")
}

func restar(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a - b
		case float64:
			return float64(a) - b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a - float64(b)
		case float64:
			return a - b
		}
	}
	panic("Error en resta: tipos incompatibles")
}
func multiplicar(a, b interface{}) interface{} {
	fmt.Printf("🧮 Multiplicando: %v (%T) * %v (%T)\n", a, a, b, b)

	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a * b
		case float64:
			return float64(a) * b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a * float64(b)
		case float64:
			return a * b
		}
	}
	panic("Error en multiplicación: tipos incompatibles")
}

func dividir(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			if b == 0 {
				panic("Error: División por cero")
			}
			return a / b
		case float64:
			if b == 0 {
				panic("Error: División por cero")
			}
			return float64(a) / b
		}
	case float64:
		switch b := b.(type) {
		case int:
			if b == 0 {
				panic("Error: División por cero")
			}
			return a / float64(b)
		case float64:
			if b == 0 {
				panic("Error: División por cero")
			}
			return a / b
		}
	}
	panic("Error en división: tipos incompatibles")
}

func menor(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a < b
		case float64:
			return float64(a) < b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a < float64(b)
		case float64:
			return a < b
		}
	}
	panic("Error: tipos incompatibles en operador '<'")
}

func menorIgual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a <= b
		case float64:
			return float64(a) <= b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a <= float64(b)
		case float64:
			return a <= b
		}
	}
	panic("Error: tipos incompatibles en operador '<='")
}

func mayor(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a > b
		case float64:
			return float64(a) > b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a > float64(b)
		case float64:
			return a > b
		}
	}
	panic("Error: tipos incompatibles en operador '>'")
}

func mayorIgual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a >= b
		case float64:
			return float64(a) >= b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a >= float64(b)
		case float64:
			return a >= b
		}
	}
	panic("Error: tipos incompatibles en operador '>='")
}

func modulo(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			if b == 0 {
				panic("❌ Error: módulo por cero")
			}
			return a % b
		}
	}
	panic("❌ Error: operador '%' solo válido entre enteros")
}

// Declare registra una nueva variable en el entorno actual.
// Retorna error si el identificador ya existe en este ambito (no en padres).
func (env *Environment) Declare(id string, value interface{}, varType string, mutable bool) error {
	if _, exists := env.Variables[id]; exists {
		return fmt.Errorf("❌ Error: variable '%s' ya fue declarada en este ámbito", id)
	}
	env.Variables[id] = Variable{
		Value:   value,
		Type:    varType,
		Mutable: mutable,
	}
	return nil
}

// DeclareFull registra una nueva variable con todos sus metadatos (linea, columna, scope).
// Usada principalmente por el nodo DeclarationNode para llenar la tabla de simbolos.
func (env *Environment) DeclareFull(
	id string,
	value interface{},
	varType string,
	mutable bool,
	line int,
	column int,
	scope string,
) error {
	if _, exists := env.Variables[id]; exists {
		return fmt.Errorf("❌ Variable '%s' ya fue declarada", id)
	}
	env.Variables[id] = Variable{
		Value:   value,
		Type:    varType,
		Mutable: mutable,
		Line:    line,
		Col:     column,
		Scope:   scope,
	}
	return nil
}

// Assign actualiza el valor de una variable ya declarada en cualquier ambito de la cadena.
// Verifica que la variable sea mutable y que el tipo del nuevo valor sea compatible.
func (env *Environment) Assign(id string, value interface{}) error {
	if variable, ok := env.lookup(id); ok {
		if !variable.Mutable {
			return fmt.Errorf("❌ Error: variable '%s' no es mutable", id)
		}
		expectedType := variable.Type
		actualType := inferType(value)

		if expectedType != actualType {
			return fmt.Errorf("❌ Error: tipo incompatible en asignación a '%s': esperado %s, recibido %s", id, expectedType, actualType)
		}
		env.set(id, value)
		return nil
	}
	return fmt.Errorf("❌ Error: variable '%s' no está definida", id)
}

// Get busca una variable por nombre en la cadena de ambitos y retorna su valor.
func (env *Environment) Get(id string) (interface{}, error) {
	if variable, ok := env.lookup(id); ok {
		return variable.Value, nil
	}
	return nil, fmt.Errorf("❌ Error: variable '%s' no está definida", id)
}

// lookup busca un identificador en el ambiente actual y en todos sus padres.
// Retorna la Variable y true si se encuentra; Variable{} y false si no.
func (env *Environment) lookup(id string) (Variable, bool) {
	current := env
	for current != nil {
		if val, ok := current.Variables[id]; ok {
			return val, true
		}
		current = current.parent
	}
	return Variable{}, false
}

func (env *Environment) set(id string, value interface{}) {
	current := env
	for current != nil {
		if _, ok := current.Variables[id]; ok {
			v := current.Variables[id]
			v.Value = value
			current.Variables[id] = v
			return
		}
		current = current.parent
	}
}

// UserFunction representa una funcion definida por el usuario.
// Se registra en el Environment cuando se encuentra un FunctionDeclarationNode.
type UserFunction struct {
	Params []*ast.ParameterNode // lista de parametros formales
	Body   *ast.BlockNode       // cuerpo de la funcion
}

// GetFunction busca una funcion por nombre en el entorno actual y en los padres.
func (e *Environment) GetFunction(name string) (*UserFunction, bool) {
	if fn, ok := e.Functions[name]; ok {
		return fn, true
	}
	if e.parent != nil {
		return e.parent.GetFunction(name)
	}
	return nil, false
}

func (i *Interpreter) Print(args ...interface{}) {
	fmt.Fprint(i.Output, joinArgs(args...))
}

func (i *Interpreter) Println(args ...interface{}) {
	fmt.Fprintln(i.Output, joinArgs(args...))
}

func joinArgs(args ...interface{}) string {
	var out string
	for idx, a := range args {
		if idx > 0 {
			out += " "
		}
		out += fmt.Sprintf("%v", a)
	}
	return out
}

func (i *Interpreter) GenerarReporteTablaSimbolos(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("❌ No se pudo crear el archivo: %s", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file, "ID,Tipo símbolo,Tipo dato,Ámbito,Línea,Columna\n")

	visited := make(map[*Environment]bool)
	env := i.global

	for env != nil {
		if visited[env] {
			break
		}
		visited[env] = true

		// Variables
		for id, v := range env.Variables {
			tipo := "Variable"
			fmt.Fprintf(file, "%s,%s,%s,%s,%d,%d\n",
				id, tipo, v.Type, v.Scope, v.Line, v.Col)
		}
		// Funciones
		for name, fn := range env.Functions {
			// Mejor información para funciones
			returnType := "void"
			if len(fn.Params) > 0 {
				returnType = fn.Params[0].Type // O puedes mejorar esto según tu implementación
			}
			fmt.Fprintf(file, "%s,Función,%s,global,0,0\n", // Changed v.CurrentScope to "global"
				name, returnType)
		}

		env = env.parent
	}

	fmt.Println("✅ Reporte de tabla de símbolos generado:", path)
	return nil
}
