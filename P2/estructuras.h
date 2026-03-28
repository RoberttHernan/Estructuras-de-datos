#ifndef ESTRUCTURAS_H
#define ESTRUCTURAS_H

/**
 * Definicion de las estructuras binarias del sistema de archivos.
 * Estas estructuras se escriben y leen directamente sobre archivos de disco
 * virtuales usando fread/fwrite, por lo que su layout en memoria es fijo.
 *
 * Jerarquia en disco:
 *   MBR -> particion[] (hasta 4)
 *     Particion extendida -> EBR -> EBR -> ... (cadena de particiones logicas)
 *     Particion formateada -> SuperBloque -> bitmap inodos -> bitmap bloques
 *                          -> TablaInodo[] -> Bloque[]
 */

#include <string>
#include <ctime>

// Extended Boot Record: encabezado de cada particion logica dentro de la extendida
typedef struct
{
    char estado_particion;      // '1' = activa, '0' = libre
    char fit_particion;         // Algoritmo de ajuste: 'f' (first), 'b' (best), 'w' (worst)
    int inicio_particion;       // Byte de inicio de la particion en el disco
    int tamanio_particion;      // Tamanio en bytes
    int sig_particion;          // Byte donde empieza el siguiente EBR (-1 si es el ultimo)
    char nombre_particion[16];  // Nombre de la particion

} ebr;

// Entrada de particion en el MBR (hasta 4 por disco)
typedef struct
{
    char estado_particion;      // '1' = activa, '0' = libre
    char tipo_particion;        // 'P' = primaria, 'E' = extendida, 'L' = logica
    char fit_particion;         // Algoritmo de ajuste
    int inicio_particion;       // Byte de inicio en el disco
    int size_particion;         // Tamanio en bytes
    char nombre_particion[16];  // Nombre de la particion
} particion;

// Master Boot Record: primer bloque de cada disco virtual
typedef struct
{
    int tamanio_mbr;             // Tamanio total del disco en bytes
    char fecha_creacion_mbr[16]; // Fecha de creacion del disco
    int signature_mbr;           // Numero aleatorio que identifica el disco
    char fit_disco;              // Fit por defecto para nuevas particiones
    particion particiones_mbr[4]; // Tabla de particiones (maximo 4)

} mbr;

// Superbloque: metadatos de la particion formateada (ext2 o ext3)
typedef struct
{
    int s_filesystem_type;    // 2 = ext2, 3 = ext3
    int s_inodes_count;       // Total de inodos en la particion
    int s_blocks_count;       // Total de bloques en la particion
    int s_free_blocks_count;  // Bloques libres disponibles
    int s_free_inodes_count;  // Inodos libres disponibles
    char s_mtime[16];         // Fecha del ultimo montaje
    char s_umtime[16];        // Fecha del ultimo desmontaje
    int s_mnt_count;          // Numero de veces que se ha montado
    int s_magic;              // Numero magico de identificacion (0xEF53)
    int s_inode_size;         // Tamanio de cada inodo en bytes
    int s_block_size;         // Tamanio de cada bloque en bytes
    int s_first_ino;          // Byte de inicio del primer inodo
    int s_first_blo;          // Byte de inicio del primer bloque
    int s_bm_inode_start;     // Byte de inicio del bitmap de inodos
    int s_bm_block_start;     // Byte de inicio del bitmap de bloques
    int s_inode_start;        // Byte de inicio de la tabla de inodos
    int s_block_start;        // Byte de inicio de la zona de bloques
} SuperBloque;

typedef struct
{
    int i_uid;       // id del usuario propietario del fichero o carpeta
    int i_gid;       //Id del grupo propietario del fichero o carpeta
    int i_size;      //tamanio de archivos en bytes
    int i_block[15]; //Lista de bloques
    char i_type;     // (0 carpeta) (1 archivo)
    int i_perm;      // permisos del archivo o carpeta
    time_t i_atime;  // fecha en que se leyo el archivo sin modicar
    time_t i_ctime;  //fecha en que se creo en inodo
    time_t i_mtime;  // fecha en que se modifico
} TablaInodo;

typedef struct
{
    char journal_op_type[10];  //Tipo de operacion que guarda el journal
    int journal_type;          // Si es archivo o carpeta
    char journal_name[100];    //Nombre del Journal
    char journal_content[100]; //Contenido del Journal
    time_t journal_date;       //Fecha de creacion del Journal
    int journal_owner;         //Propietario del Journal
    int journal_permissions;   //Permisos de Journal
} Journal;

typedef struct
{
    int id_user;
    int id_grp;
    int inicioSuper;
    int inicioJournal;
    int tipo_sistema;
    std::string direccion;
    char fit;
} Sesion;
typedef struct
{
    char b_name[12]; //Nombre carpeta/archivo
    int b_inodo;     //Apuntador hacia un inodo asociado al archivo o carpeta
} Content;

typedef struct
{
    Content b_content[4];
} BloqueCarpeta;

typedef struct
{
    int b_pointer[16];
} BloqueApuntador;
typedef struct
{
    char b_content[64];
} BloqueArchivo;

class estructuras
{
private:
public:
    estructuras();
};

#endif