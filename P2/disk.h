#ifndef DISK_H
#define DISK_H

#include <string>
#include <stdio.h>
#include "estructuras.h"
#include <string.h>
#include <iterator>
#include <tgmath.h>
#include <libgen.h>
using namespace std;

namespace dsk
{

    /**
     * Controlador principal del sistema de archivos simulado.
     * Administra discos virtuales (archivos binarios) que contienen particiones
     * con formato ext2 o ext3. Soporta inodos, bloques de directorio/archivo/apuntador,
     * bitmaps de asignacion, journaling (ext3) y reportes graficos via Graphviz.
     *
     * Algoritmos de ajuste soportados: First-Fit ('f'), Best-Fit ('b'), Worst-Fit ('w').
     * Unidades de tamanio: kilobytes ('k'), megabytes ('m').
     */
    class disk
    {
    private:
    public:
        bool flag_login = false; // Indica si hay una sesion de usuario activa
        Sesion sesion_activa;    // Datos de la sesion en curso (usuario, grupo, particion)

        // Crea un nuevo disco virtual (archivo binario) con el tamanio y fit indicados
        void createDisk(int size, std::string fit, char unit, std::string path);
        // Ejecuta operaciones sobre el disco: crear/eliminar/ampliar particion logica
        void fkDisk(int size, char unit, std::string path, char type, char fit, std::string delete_param, int add, string name);
        void deleteDisk(std::string path);
        void imprimirDatosDisco(string path);
        void createPart(string path, int size, char unit, char type, char fit, string name);
        int createCarpet (std::string path_carpeta, bool p_param); 
        bool existPart(std::string path, std::string name);
        void createPrimaryPart(string path, int size, char unit, char fit, string name, char type);
        void createExtendedPart(std::string path, int size, char unit, std::string name, char fit);
        void createLogicPart(std::string path, int size, char unit, char fit, std::string name);
        void ReporteMBR(std::string path_disco, std::string path_destino, std::string extension);
        void ReporteDisk(std::string path_disco, std::string path_destino, std::string extension);
        void deletePart(std::string path, std::string name, std::string type_delete);
        void Add(std::string path, std::string name, char unit, int add_param);
        int FindLogicPart(std::string path, std::string name);
        int FindPrimExtPart(std::string path, std::string name);
        void mountPart(std::string path, std::string name);
        void unmount(std::string id);
        void rep(std::string name, std::string path, std::string id_param, std::string ruta);
        void mkfs(std::string id, std::string type, int fs);
        void FormateoExt2(int tamanio_particion, int inicio_particion, std::string path);
        void ReporteSuper(std::string path_disco, std::string path_destino, int inicio_super);
        void ReporteBlock (std::string path_disco, std::string path_destino, int bm_block_start, int block_start, int inode_start); 
        void ReporteBitmap (std::string path_disco, std::string path_destino, int start_bm, int n); 
        void ReporteArbol (std::string path_disco, std::string path_destino, int inicio_super); 
        void ReporteJournaling (std::string path_disco, std::string path_destion, int inicio_sp);
        void FormateoExt3(int tamanio_particion, int inicio_particon, std::string path);
        void login(std::string id, std::string usuario, std::string password);
        void Logout();
        void ComandoMKDIR(std::string path_carpeta, bool flag_p ); 
        void SaveOpJournal (char * op, int type, int rights, char * name, char * content); 
        /**
         * Funcion que devuelve el numero de un apuntador libre en bloque inodo
         * @param *arch: Archivo fisico en el que se encuentra el disco del sistema
         * @param numero_inodo: Numero de inodo en donde se encuentra la carpeta
         * Los otros parametro son las referencias de los bloques en los cuales se buscara espacio
         * 
         * */
        int LookFreePointerInode(FILE *arch, int numero_inodo, TablaInodo &tablaInodo, BloqueCarpeta &bloque_carpeta, BloqueApuntador &bloque_apuntador, int &contenido, int &bloque, int &apuntero);

        /*
        *Genera el reporte de los inodos en un particion
        *@param path_disco: Ruta del disco en donde se encuentra la particion
        *@param path_destino: Ruta en donde se generara el reporte
        *@param bm_inode_start: inicio del bitmap de inodos (en byte)
        *@param inode_start: inicio de la tabla de inodos (byte)
        *@param bm_block_start: inicio del bitmap de bloques 
        */
        void ReporteInodo(std::string path_disco, std::string path_destino, int bm_inode_start, int inode_start, int bm_block_start);
        /*
        *Funcion que verifica los datos de un usuario que quiera logearse
        *@param usuario: nombre de usuario a logear
        *@param password: contraseña de usuario
        *@param path_disco: ruta del disco en donde se encuentra la particion
        */
        int ComprobarDatos(std::string usuario, std::string password, std::string path_disco);

        /**
         *Función Que crea una nueva carpeta en el sistema de archivos
         *@param arch: Archivo Fisico en donde se encuentra en disco 
         *@param fit: Fit del la particion en donde se quiere crear la carpeta
         *@param flag_p: Recursividad en la creacion de carpetas padre
         *@param path*: Direccion de la carpeta en el sistema
         *@param num_inodo: numero de inodo en donde se debe crear la carpeta
         @return 1: Carpeta creada  2:Error en permisos de escritura 3: No existe la carpeta padre y no hay parametro p
         **/
        int MakeNewDIr(FILE *arch, char fit, bool flag_p, char *path, int num_inodo);
        /*
        *Funcion que retorna el numero de inodo en donde se encuentra una carpeta o fichero
        *@param *arch: Archivos que contiene la particion (archivo fisico)
        *@param *path: Direccion de la carpeta o archivo (en sistema de archivos)
        */
        int existFileorDir(FILE *arch, char *path);
        int FindGroup(std::string name);

        /*
        *Funcion que retorna el byte en donde empieza uno inodo o bloque
        *@param *arch: Archivo en el que se encuentra la partición (archivo fisico)
        *@param num_inodo: numero de inodo o bloque a buscar
        *@param type: tipo de elemento a buscar (1= inodo), (2= bloque); 
        */
        int inicioInodoBloque(FILE *arch, int num_inodo, char type);

        /**
         * Funcion que verifica el permiso de escritura del usuario activo
         * @param pass: Permisos del Bloque Inodo
         * @param flag_user: Indica si el usuario es el propietario  de la carpeta padre
         * @param flag_group: Indica si el grupo es el propietario de la carpeta padre
         * return true= Con permisos activos | false= Sin permisos activos
         * */
        bool WriteRights(int pass, bool flag_user, bool flag_group);
        
        /**
         * Funcion que devuelve el bit libre en el bitmap de inodos o bloques segun el fit de la particion
         * @param *archivo fisico del disco que se esta leyendo
         * @param type: Tipo de bit a buscar (inodo ó bloque)
         * @param fit: Fit de la particion
         * @return -1 cuando no existen mas bloques libres
         * 
         * */
        int FindFreeBit (FILE *arch, char type, char fit); 
        TablaInodo InodoCreate(int size_i, char type_i, int perms_i); 
        BloqueCarpeta createDirBlock(); 



        
    };

}

#endif