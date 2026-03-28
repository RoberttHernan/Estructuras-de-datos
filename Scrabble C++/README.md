# Scrabble en C++ - Implementacion con Estructuras de Datos

Implementacion del juego de mesa Scrabble en C++ utilizando estructuras de datos personalizadas. El proyecto fue desarrollado como parte de un portafolio de estructuras de datos, con el objetivo de demostrar la aplicacion practica de estructuras lineales y no lineales en un contexto real.

## Descripcion del Proyecto

El juego corre en consola y soporta dos jugadores en modo por turnos. La configuracion del tablero y el diccionario de palabras se cargan desde un archivo JSON externo. El sistema genera visualizaciones de las estructuras de datos internas mediante Graphviz.

### Funcionalidades principales

- Carga de configuracion desde archivo JSON (dimension del tablero, casillas especiales, diccionario).
- Registro de jugadores con historial de puntuaciones.
- Partida entre dos jugadores con validacion de palabras contra el diccionario.
- Sistema de puntuacion con multiplicadores de casillas (doble y triple).
- Cambio de fichas durante el turno.
- Reportes graficos de todas las estructuras de datos generados con Graphviz.
- Ranking de jugadores por mayor puntuacion.

---

## Estructuras de Datos Implementadas

### Matriz Dispersa (`Matriz Dispersa/`)
Representa el tablero de juego. Solo almacena las casillas que contienen datos, ahorrando memoria para tableros grandes con pocas fichas colocadas. Cada nodo tiene cuatro punteros (`arriba`, `abajo`, `izquierda`, `derecha`) para navegacion ortogonal. Un nodo raiz con coordenada `(-1, -1)` actua como punto de entrada, con cabeceras de columna (y=-1) y de fila (x=-1).

Operaciones clave: `AddElement`, `buscarPalabraMatriz_x`, `buscarPalabraMatriz_y`, `borrarNodo`, `textoGraphviz`.

### Arbol Binario de Busqueda (`ABB/`)
Almacena los jugadores registrados ordenados lexicograficamente por nombre de usuario. Permite busqueda eficiente para verificar duplicados y recuperar el perfil de un jugador. Soporta los tres recorridos (inorden, preorden, postorden) con generacion de imagenes para cada uno.

Operaciones clave: `Add`, `buscar`, `buscarJugador`, `GraficarInorden`, `GraficarPreorden`, `GraficarPostorden`, `GraficareportePuntajeJugador`.

### Cola (FIFO) (`Cola/`)
Simula la bolsa de fichas del juego. Las fichas se distribuyen aleatoriamente respetando la distribucion oficial del Scrabble en espanol (A=12, E=12, O=9, etc.). Cada ficha es un objeto `Letra` con caracter y valor en puntos.

Operaciones clave: `Queue` (encolar), `InQueue` (desencolar desde el frente), `buscar`.

### Lista Doblemente Enlazada Circular (`Lista doblemente enlazada Circular/`)
Representa la mano de fichas de cada jugador (7 fichas). La estructura circular permite iterar continuamente sobre las fichas sin perder el extremo. Soporta eliminacion de un nodo especifico al momento de colocar una ficha en el tablero.

Operaciones clave: `AddHead`, `borrar`, `borrarNodo`, `borrarLista`, `buscar`.

### Lista Simple (`lista Simple/`)
Estructura auxiliar de proposito general. Se usa para almacenar las casillas especiales cargadas desde el JSON, las casillas temporales de la jugada en curso, y el historial de puntuaciones de cada jugador. Soporta insercion ordenada.

Operaciones clave: `AddHead`, `AddEnd`, `AddSort`, `DeleteByData`, `DeleteAll`.

---

## Clases del Dominio

| Clase | Archivo | Descripcion |
|---|---|---|
| `Jugador` | `Otra Clases/Jugador.h` | Nombre de usuario y puntero al historial de puntajes (`listaSimple<int>`). |
| `Letra` | `Otra Clases/Letra.h` | Ficha del juego: caracter (`char`) y valor en puntos (`int`). |
| `Casillas` | `Otra Clases/Casillas.h` | Casilla del tablero: tipo (`Normal`, `Doble`, `Triple`) y coordenadas (x, y). |

---

## Algoritmos Destacados

- **`llenarColaAleatoria()`**: Llena la cola con la distribucion oficial de fichas del Scrabble en espanol usando numeros aleatorios para determinar el orden de insercion.
- **`match()`**: Valida si las fichas colocadas en el turno actual forman alguna palabra valida del diccionario, buscando en la columna y fila de cada ficha colocada.
- **`getPuntuacionTotal()`**: Calcula el puntaje de un turno recorriendo la columna o fila de la palabra y aplicando el multiplicador de la casilla (`BuscarDoble_Triple()`).
- **`config()`**: Parsea el archivo JSON con la libreria `nlohmann/json` para cargar la dimension, casillas especiales y el diccionario.

---

## Dependencias

- **Compilador C++**: g++ con soporte para C++11 o superior.
- **Graphviz**: Para generar los reportes graficos (`.dot` -> `.png`). Debe estar instalado y accesible en el `PATH`.
- **nlohmann/json**: Libreria de cabecera unica incluida en `configJson/json.hpp`. No requiere instalacion adicional.

---

## Compilacion y Ejecucion

```bash
# Compilar
g++ -std=c++11 main.cpp -o scrabble

# Ejecutar
./scrabble
```

---

## Formato del Archivo de Configuracion JSON

```json
{
  "dimension": 10,
  "casillas": {
    "dobles": [
      { "x": 2, "y": 2 },
      { "x": 8, "y": 8 }
    ],
    "triples": [
      { "x": 5, "y": 5 }
    ]
  },
  "diccionario": [
    { "palabra": "perro" },
    { "palabra": "gato" },
    { "palabra": "casa" }
  ]
}
```

- `dimension`: Tamanio N del tablero (N x N).
- `casillas.dobles`: Lista de coordenadas con multiplicador x2.
- `casillas.triples`: Lista de coordenadas con multiplicador x3.
- `diccionario`: Lista de palabras validas para el juego.

---

## Flujo del Menu

```
1. Agregar configuraciones  -> Carga el JSON e inicializa el tablero y la bolsa de fichas
2. Agregar jugadores        -> Registra un jugador en el ABB
3. Seleccionar jugadores    -> Asigna jugador 1 y jugador 2 para la partida
4. Jugar                    -> Inicia la partida (repartir fichas, turnos, terminar)
5. Reportes                 -> Genera imagenes Graphviz de las estructuras
10. Salir
```

---

## Estructura del Repositorio

```
P1/
+-- ABB/                              # Arbol Binario de Busqueda
|   +-- ABB.h / ABB.cpp
|   +-- NodoArbol.h / NodoArbol.cpp
+-- Cola/                             # Cola FIFO para la bolsa de fichas
|   +-- Cola.h / Cola.cpp
|   +-- NodoCola.h / NodoCola.cpp
+-- Lista doblemente enlazada Circular/
|   +-- listaEnlazadaDobleCircular.h / .cpp
|   +-- NodoDCLL.h / NodoDCLL.cpp
+-- lista Simple/
|   +-- listaSimple.h / listaSimple.cpp
|   +-- Nodo.h / Nodo.cpp
+-- Matriz Dispersa/                  # Tablero de juego
|   +-- Matriz.h / Matriz.cpp
|   +-- NodeMatriz.h / NodeMatriz.cpp
+-- Otra Clases/                      # Entidades del dominio
|   +-- Jugador.h / Jugador.cpp
|   +-- Letra.h / Letra.cpp
|   +-- Casillas.h / Casillas.cpp
+-- configJson/
|   +-- json.hpp                      # nlohmann/json (cabecera unica)
+-- main.cpp                          # Controlador principal y logica del juego
+-- .gitignore
+-- README.md
```
