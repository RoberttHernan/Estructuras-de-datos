# Sistema de Archivos en C++

Simulador de sistema de archivos implementado en C++ que opera sobre discos virtuales (archivos binarios). El proyecto implementa los formatos ext2 y ext3, administracion de particiones mediante MBR y EBR, y un interprete de comandos propio con analizador lexico y parser.

## Descripcion del Proyecto

El simulador permite crear y gestionar discos virtuales con particiones primarias, extendidas y logicas. Las particiones pueden formatearse con ext2 (sin journaling) o ext3 (con journaling). El sistema soporta operaciones de directorio y archivo, control de acceso por usuario y grupo, y generacion de reportes graficos de las estructuras internas mediante Graphviz.

El interprete acepta comandos de forma interactiva o desde un archivo de script.

### Funcionalidades

- Crear y eliminar discos virtuales de tamanio configurable.
- Gestionar particiones primarias, extendidas y logicas con algoritmos de ajuste First-Fit, Best-Fit y Worst-Fit.
- Formatear particiones con ext2 o ext3.
- Montar y desmontar particiones, asignandoles un ID (e.g. `701a`).
- Crear directorios y archivos dentro de una particion montada.
- Control de acceso: login/logout de usuarios con permisos por propietario, grupo y otros.
- Journaling de operaciones en particiones ext3.
- Reportes graficos: MBR, disco, superbloque, inodos, bloques, bitmap, arbol de directorios, journaling.

---

## Arquitectura

### Analizador Lexico (`analizador.h / analizador.cpp`)

Maquina de estados finita con 6 estados que tokeniza la entrada del usuario. Reconoce palabras reservadas (comandos y parametros), enteros, rutas de archivo (con y sin comillas) y comentarios (`#`). Produce una lista de tokens que el parser consume.

### Parser y Ejecutor (`ejecucion.h / ejecucion.cpp`)

Itera sobre la lista de tokens, identifica el comando activo y recopila sus parametros antes de invocar la operacion correspondiente en el objeto `disk`. Maneja 11 comandos distintos.

### Controlador de Disco (`disk.h / disk.cpp`)

Nucleo del simulador. Implementa todas las operaciones sobre el archivo binario que representa el disco:

- Escritura del MBR y particiones con `fwrite/fread`.
- Asignacion de espacio con algoritmos First-Fit, Best-Fit y Worst-Fit sobre los bitmaps de inodos y bloques.
- Navegacion del arbol de directorios mediante inodos y bloques de carpeta.
- Verificacion de permisos (`WriteRights`) con notacion octal (e.g. `664`, `755`).
- Generacion de archivos `.dot` para Graphviz.

### Estructuras Binarias (`estructuras.h`)

Structs de tamanio fijo que se serializan directamente en disco con `fwrite`:

| Estructura | Descripcion |
|---|---|
| `mbr` | Master Boot Record: tamanio, firma, fit, tabla de 4 particiones |
| `particion` | Entrada en el MBR: tipo (P/E/L), fit, inicio, tamanio, nombre |
| `ebr` | Extended Boot Record: encabezado de cada particion logica |
| `SuperBloque` | Metadatos ext2/ext3: contadores, magic number, posiciones de bitmap e inodos |
| `TablaInodo` | Inodo: uid, gid, tamanio, arreglo de 15 punteros a bloques, permisos, timestamps |
| `BloqueCarpeta` | Bloque de directorio: arreglo de 4 entradas (nombre + numero de inodo) |
| `BloqueApuntador` | Bloque de apuntadores indirectos: 16 punteros a otros bloques |
| `BloqueArchivo` | Bloque de datos: 64 bytes de contenido |
| `Journal` | Entrada de journal: tipo de operacion, nombre, contenido, fecha, propietario |
| `Sesion` | Estado de la sesion activa: id de usuario/grupo, inicio del superbloque, fit |

### Lista de Particiones Montadas (`mount/`)

Lista simplemente enlazada que registra las particiones montadas durante la sesion. Cada nodo almacena la ruta del disco fisico, el nombre de la particion y el ID de montaje asignado.

---

## Comandos del Interprete

| Comando | Descripcion |
|---|---|
| `mkdisk -size=N -u=[k/m] -f=[bf/ff/wf] -path="ruta"` | Crea un nuevo disco virtual |
| `rmdisk -path="ruta"` | Elimina un disco virtual |
| `fdisk -size=N -u=[k/m] -type=[P/E/L] -f=[bf/ff/wf] -name=X -path="ruta"` | Crea una particion |
| `mount -path="ruta" -name=X` | Monta una particion y le asigna un ID |
| `unmount -id=XX` | Desmonta la particion con el ID indicado |
| `mkfs -id=XX -type=[full] -fs=[2/3]` | Formatea una particion montada |
| `login -usr=X -pwd=X -id=XX` | Inicia sesion en la particion montada |
| `logout` | Cierra la sesion activa |
| `mkdir -path="ruta" [-r]` | Crea un directorio (con `-r` crea padres faltantes) |
| `mkfile -path="ruta" [-r] [-size=N] [-cont="texto"]` | Crea un archivo |
| `rep -name=X -path="ruta" -id=XX` | Genera un reporte Graphviz |
| `exec -path="ruta"` | Ejecuta un script de comandos desde archivo |
| `pause` | Pausa la ejecucion hasta que el usuario presione Enter |
| `# comentario` | Linea de comentario, ignorada por el analizador |

Reportes disponibles (`-name`): `mbr`, `disk`, `sb`, `inode`, `block`, `bm_inode`, `journaling`, `tree`.

---

## Dependencias

- **Compilador C++**: g++ con soporte para C++11 o superior.
- **Graphviz**: Para generar los reportes graficos (`.dot` -> `.png`). Debe estar instalado y accesible en el `PATH`.

---

## Compilacion y Ejecucion

```bash
# Compilar
g++ -std=c++11 main.cpp analizador.cpp token.cpp ejecucion.cpp disk.cpp estructuras.cpp mount/listaSimple.cpp mount/nodoLista.cpp -o sistema_archivos

# Ejecutar en modo interactivo
./sistema_archivos

# Ejecutar un script (dentro del interprete)
# exec -path="script.txt"
```

---

## Estructura del Repositorio

```
P2/
+-- main.cpp                        # Punto de entrada y bucle interactivo
+-- token.h / token.cpp             # Clase Token del analizador lexico
+-- analizador.h / analizador.cpp   # Analizador lexico (maquina de estados)
+-- ejecucion.h / ejecucion.cpp     # Parser y ejecutor de comandos
+-- estructuras.h / estructuras.cpp # Estructuras binarias del sistema de archivos
+-- disk.h / disk.cpp               # Controlador de disco y operaciones de filesystem
+-- mount/                          # Lista de particiones montadas
|   +-- listaSimple.h / listaSimple.cpp
|   +-- nodoLista.h / nodoLista.cpp
+-- .gitignore
+-- README.md
```
