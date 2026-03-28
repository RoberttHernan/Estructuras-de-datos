/*
 * Lista Simple (Singly Linked List)
 *
 * Lista enlazada simple de proposito general utilizada en el proyecto para:
 *   - Almacenar las casillas especiales cargadas desde el JSON (listaCasillas).
 *   - Guardar las fichas colocadas en el turno actual antes de validarlas (listaCasillasTemp).
 *   - Mantener el historial de puntajes de cada jugador (puntaje_usuario en Jugador).
 *   - Construir el ranking de jugadores en el reporte de puntuaciones totales.
 *
 * Soporta insercion al inicio, al final e insercion ordenada ascendente (AddSort).
 */
#ifndef LISTASIMPLE_H_
#define LISTASIMPLE_H_

#include <iostream>
#include <string>
#include <stdlib.h>

#include "Nodo.h"
#include "Nodo.cpp"

using namespace std;
template <class T>
class listaSimple
{
private:
    Nodo<T> *head;  // Primer nodo de la lista
    Nodo<T> *tail;  // Ultimo nodo de la lista
    int index;      // Numero de elementos
public:
    listaSimple();
    ~listaSimple();

    // Inserta un nodo al inicio
    void AddHead(T);
    // Inserta un nodo al final
    void AddEnd(T);
    // Inserta un nodo manteniendo orden ascendente (usado para historial de puntajes)
    void AddSort(T);
    void Concat(listaSimple);
    // Libera todos los nodos de la lista
    void DeleteAll();
    void DeleteByData(T);
    void DeleteByPosition(int);
    void FillByUser(int);
    void FillRandom(int);
    void Intersection(listaSimple);
    void Invert();
    void LoadFile();
    void Print();
    void SaveFile(string);
    void search(T);
    void sort();

    Nodo<T> *getNodo();
    void SetNodo(Nodo<T>);
    Nodo<T> *getHead();
    void setHead(Nodo<T>*);
    int getIndex();
    void setIndex(int);
    Nodo<T> *getTail();
    void setTail(Nodo<T> *);
};












#endif