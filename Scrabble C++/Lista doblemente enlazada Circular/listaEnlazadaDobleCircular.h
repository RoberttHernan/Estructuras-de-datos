/*
 * Lista Doblemente Enlazada Circular (DCLL - Doubly Circular Linked List)
 *
 * Representa la mano de fichas de cada jugador (7 fichas por defecto).
 * Es circular: el ultimo nodo (tail) apunta de vuelta al head, y el head
 * apunta hacia atras al tail. Esto permite iterar continuamente por la mano
 * sin perder el extremo de la lista.
 *
 * Se usa como diccionario de palabras validas (string) y como mano de fichas (Letra).
 */
#ifndef LISTAENLAZADADOBLECIRCULAR_H_
#define LISTAENLAZADADOBLECIRCULAR_H_

#include <iostream>
#include <string>
#include <stdlib.h>

#include "NodoDCLL.cpp"

using namespace std;
template <class T>
class listaEnlazadaDobleCircular
{
private:
    NodoDCLL<T> *head;  // Primer nodo de la lista
    NodoDCLL<T> *tail;  // Ultimo nodo de la lista (apunta de vuelta al head)
    int index;          // Numero de elementos en la lista
public:
    listaEnlazadaDobleCircular();
    ~listaEnlazadaDobleCircular();

    // Inserta un nodo al inicio de la lista y actualiza los punteros circulares
    void AddHead(T);
    void Print();
    // Busca un dato y retorna true si lo encuentra
    bool buscar(T);
    // Elimina la primera ocurrencia del dato en la lista
    void borrar(T data);
    // Elimina un nodo especifico dado su puntero (usado al jugar una ficha)
    void borrarNodo(NodoDCLL<T> *nodo);
    // Libera todos los nodos de la lista
    void borrarLista();

    NodoDCLL<T> *getNodo();
    void SetNodo(NodoDCLL<T>);
    int getIndex();
    void setIndex(int);
    NodoDCLL<T> *getHead();
    void setHead(NodoDCLL<T> *);
    NodoDCLL<T> *getTail();
};












#endif