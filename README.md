# Estructuras de Datos

Coleccion de proyectos que implementan estructuras de datos fundamentales en C++. Cada proyecto aplica multiples estructuras en un contexto practico y funcional, demostrando su uso real mas alla de implementaciones aisladas. El objetivo del repositorio es servir como portafolio tecnico que evidencie el dominio de estructuras lineales, no lineales y sus algoritmos asociados.

## Proyectos

### Sistema de Archivos en C++ (`P2/`)

Simulador de sistema de archivos que opera sobre discos virtuales (archivos binarios). Implementa los formatos ext2 y ext3 con particiones primarias, extendidas y logicas administradas mediante MBR y EBR. Incluye un interprete de comandos propio con analizador lexico (maquina de estados finita) y parser.

Estructuras y conceptos implementados:
- **MBR / EBR**: structs serializados directamente en disco con `fwrite/fread`.
- **SuperBloque, Inodos y Bloques**: tabla de inodos, bloques de carpeta, archivo y apuntador.
- **Bitmaps**: asignacion de inodos y bloques con algoritmos First-Fit, Best-Fit y Worst-Fit.
- **Journaling**: registro de operaciones para particiones ext3.
- **Lista simplemente enlazada**: registro de particiones montadas en sesion.
- **Analizador lexico**: maquina de estados finita de 6 estados para tokenizar comandos.

Lenguaje: C++11 | Dependencias: Graphviz

---

### Scrabble en C++ (`Scrabble C++/`)

Implementacion del juego de mesa Scrabble para dos jugadores en consola. El proyecto integra cinco estructuras de datos personalizadas para representar cada componente del juego:

- **Matriz Dispersa** (lista ortogonal enlazada): tablero de juego NxN que solo almacena casillas con contenido.
- **Arbol Binario de Busqueda (ABB)**: registro de jugadores ordenado lexicograficamente por nombre de usuario.
- **Cola FIFO**: bolsa de fichas con la distribucion oficial del Scrabble en espanol (~95 fichas).
- **Lista Doblemente Enlazada Circular**: mano de cada jugador (7 fichas), con iteracion continua.
- **Lista Simple**: almacenamiento auxiliar para casillas especiales, jugadas en curso e historial de puntajes.

La configuracion del tablero (dimension, casillas especiales, diccionario) se carga desde un archivo JSON externo usando la libreria `nlohmann/json`. El sistema genera reportes graficos de las estructuras internas en formato Graphviz.

Lenguaje: C++11 | Dependencias: Graphviz, nlohmann/json (incluida)
