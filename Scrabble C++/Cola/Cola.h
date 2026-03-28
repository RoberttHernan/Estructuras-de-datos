/*
 * Cola (Queue) - Estructura FIFO
 *
 * Simula la bolsa de fichas del juego. Las fichas se insertan por el tail
 * y se extraen desde el head, garantizando el orden de llegada (FIFO).
 *
 * Se usa para:
 *   - Almacenar las fichas disponibles del juego (bolsa inicial ~95 fichas).
 *   - Distribuir 7 fichas a cada jugador al inicio de la partida.
 *   - Recibir las fichas devueltas cuando un jugador cambia su mano.
 */
#ifndef COLA_H_
#define COLA_H_

#include <string>
#include <iostream>
#include "NodoCola.cpp"

template <class T>
class Cola
{
private:
    NodoCola<T> *head;  // Frente de la cola (siguiente en salir)
    NodoCola<T> *tail;  // Final de la cola (ultimo en entrar)
    int size;           // Numero de elementos en la cola
public:
    Cola();
    ~Cola();
    // Encola un nuevo elemento al final (tail)
    void Queue(T);
    // Desencola el elemento del frente (head)
    void InQueue();
    void Print();

    NodoCola<T> *getHead();
    void SetHead(NodoCola<T>*);
    NodoCola<T> *getTail();
    void SetTail(NodoCola<T>*);
    int getSize();
    void setSize(int);
    // Retorna true si el elemento existe en la cola
    bool buscar(T);
    void DeleteAll();
};





#endif