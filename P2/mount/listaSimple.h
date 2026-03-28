#pragma once

#include "nodoLista.h"

/**
 * Lista simplemente enlazada de particiones montadas.
 * Cada nodo almacena la ruta del disco, el nombre de la particion,
 * y el ID asignado al montaje (formato: "70" + numero + letra, e.g. "701a").
 * Se usa para resolver el ID de particion al ejecutar comandos como mkfs, login y rep.
 */
class listaSimple
{
private:
    nodoLista *head; // Puntero al primer nodo de la lista

public:
    listaSimple();

    // Retorna el numero del siguiente ID disponible para el path y nombre dados
    int findNumber(std::string, std::string);
    // Retorna el caracter (letra) del siguiente ID disponible para el path y nombre dados
    int findLetter(std::string, std::string);
    // Imprime en consola todas las particiones actualmente montadas
    void showList();
    // Inserta un nodo al final de la lista
    void insertNode(nodoLista *nodo_nuevo);
    // Retorna la ruta del disco fisico asociado al ID de montaje
    std::string getPathDisk(std::string id_param);
    // Elimina de la lista la particion con el ID indicado; retorna 1 si tuvo exito
    int deleteNode(std::string id);
    // Retorna true si existe una particion montada con el path y nombre dados
    bool findNode(std::string, std::string);
    // Retorna el puntero al primer nodo de la lista
    nodoLista *getHead();
    // Retorna el nodo cuyo ID coincide con id; nullptr si no existe
    nodoLista *getNodo(std::string id);
};
