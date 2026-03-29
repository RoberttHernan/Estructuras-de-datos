
# 🧠 Reporte Técnico

## "Registro de Variables y Funciones en el Intérprete VLangCherry"

---

##  1. ¿Dónde y cómo se crean las variables?

### 🔧 En `ast_builder.go`

Las variables se transforman en nodos `*ast.DeclarationNode` a través de los métodos:

| Método Visitor | Descripción |
|----------------|-------------|
| `VisitDec1`    | `mut x int = 5` – Declaración con tipo y valor |
| `VisitDec2`    | `x := 5` – Declaración por inferencia de tipo |
| `VisitDec3`    | `x int = 5` – Declaración sin `mut` (inmutable) |

###  Nodo generado:
Cada uno retorna un nodo de tipo `*ast.DeclarationNode`:

```go
&ast.DeclarationNode{
    Identifier: &ast.IdentifierNode{Name: id},
    Type:       tipo,
    Value:      valor,
    Mutable:    true/false,
    Line:       ctx.GetStart().GetLine(),
    Col:        ctx.GetStart().GetColumn(),
    Scope:      v.CurrentScope,
}
```

📌 **Importante:** el campo `Scope` debe ser dinámico usando `v.CurrentScope`.

---

##  2. ¿Cómo se evalúan y guardan las variables?

### 🔧 En `interpreter.go`

Cuando se encuentra un nodo `*ast.DeclarationNode`:

```go
case *ast.DeclarationNode:
    name := n.Identifier.Name
    value := i.Eval(n.Value)
    varType := n.Type

    if varType == "" {
        varType = inferType(value)
    }

    err := i.global.DeclareFull(
        name, value, varType,
        n.Mutable, n.Line, n.Col, n.Scope,
    )
```

###  Lógica del registro

- Se evalúa el valor de la variable.
- Se infiere el tipo si no se especificó.
- Se guarda en el entorno actual con `DeclareFull`.

---

##  Método: `DeclareFull(...)`

```go
func (env *Environment) DeclareFull(id string, value interface{}, varType string, mutable bool, line int, column int, scope string) error
```

Este método:

- Verifica si la variable ya fue declarada.
- Guarda su valor, tipo, mutabilidad y metadata (`línea`, `columna`, `scope`) en `env.Variables`.

---

##  3. ¿Dónde y cómo se crean las funciones?

### 🔧 En `ast_builder.go`

Las funciones se reconocen en:

| Método Visitor  | Descripción |
|-----------------|-------------|
| `VisitDecFunc1` | `func nombre(params...) tipo { ... }` |
| `VisitDecFunc2` | `func nombre(params...) { ... }` |

Ambos devuelven un nodo `*ast.FunctionDeclarationNode`:

```go
&ast.FunctionDeclarationNode{
    Name:       name,
    Params:     []*ast.ParameterNode,
    ReturnType: returnType,
    Body:       *ast.BlockNode,
}
```

### 🧩Scope del cuerpo

```go
v.CurrentScope = name
```

---

##  4. ¿Cómo se registran y ejecutan las funciones?

### 🔧 En `interpreter.go`

En `Evaluate(...)`, se registran así:

```go
for _, smt := range node.Elementos {
    if fn, ok := smt.(*ast.FunctionDeclarationNode); ok {
        i.global.Functions[fn.Name] = &UserFunction{
            Params: fn.Params,
            Body:   fn.Body,
        }
    }
}
```

###  `UserFunction` struct:

```go
type UserFunction struct {
    Params []*ast.ParameterNode
    Body   *ast.BlockNode
}
```

---

## 🧮 5. ¿Cómo se ejecutan funciones?

Al encontrar `*ast.FunctionCallNode`:

```go
fn, ok := i.global.GetFunction(n.Name)
i.global = NewEnvironment(i.global)
// asigna parámetros
i.Eval(fn.Body)
```
 Se ejecuta en un nuevo entorno local.

---

## 📤 6. ¿Cómo se genera el reporte de símbolos?

### 🔧 Método `GenerarReporteTablaSimbolos(path string)`

```go
for env != nil {
    for id, v := range env.Variables {
        fmt.Fprintf(file, "%s,Variable,%s,%s,%d,%d\n", id, v.Type, v.Scope, v.Line, v.Col)
    }
    for name := range env.Functions {
        fmt.Fprintf(file, "%s,Función,void,Global,0,0\n", name)
    }
    env = env.parent
}
```

---

## Recomendaciones

| Aspecto | Sugerencia |
|---------|------------|
| Scope estático | Usa `v.CurrentScope` dinámicamente |
| Validación | Imprime entornos anidados también |
| Declaraciones en bloques | Opcional: scopes tipo `main_block1` |

---
