/*
 * Matriz Dispersa (Sparse Matrix)
 *
 * Representa el tablero de juego como una lista ortogonal de listas enlazadas.
 * Solo almacena los nodos que contienen datos, ahorrando memoria cuando el tablero
 * tiene pocas fichas colocadas.
 *
 * Estructura interna:
 *   - El nodo raiz tiene coordenadas (-1, -1).
 *   - Los nodos cabecera de columna tienen y = -1.
 *   - Los nodos cabecera de fila tienen x = -1.
 *   - Cada nodo de dato tiene coordenadas (x, y) reales y punteros a los cuatro vecinos.
 *
 * Navegacion:
 *   - Horizontal (izquierda/derecha): recorre una fila.
 *   - Vertical (arriba/abajo): recorre una columna.
 */
#ifndef MATRIZ_H_
#define MATRIZ_H_

#include <string>
#include "NodeMatriz.cpp"

using namespace std;
template <class T>
class Matriz
{
private:
    // Nodo raiz con coordenada (-1, -1), punto de entrada a toda la estructura
    NodeMatriz<T> *root = new NodeMatriz<T>("root", -1, -1);
public:
    Matriz();
    ~Matriz();
    NodeMatriz<T>* getRoot();
    void setRoot(NodeMatriz<T>*);
    // Busca y retorna el nodo cabecera de la columna x; retorna NULL si no existe
    NodeMatriz<T>* searchCol(int x);
    // Busca y retorna el nodo cabecera de la fila y; retorna NULL si no existe
    NodeMatriz<T>* searchRow(int y);
    // Crea e inserta el nodo cabecera de la columna x de forma ordenada
    NodeMatriz<T>* AddCol(int x);
    // Crea e inserta el nodo cabecera de la fila y de forma ordenada
    NodeMatriz<T>* AddRow(int y);
    // Inserta un elemento en (x, y), creando fila y columna si no existen
    void AddElement(T data, int x, int y);
    // Inserta un nodo de forma ordenada en la lista horizontal (recorre por right)
    NodeMatriz<T>* AddSortCol(NodeMatriz<T>*, NodeMatriz<T>*);
    // Inserta un nodo de forma ordenada en la lista vertical (recorre por down)
    NodeMatriz<T>* AddSortRow(NodeMatriz<T>*, NodeMatriz<T>*);
    // Genera el archivo .dot y la imagen .png de la matriz para visualizacion
    void textoGraphviz();
    // Retorna como string la concatenacion de letras en la columna x (palabra vertical)
    string buscarPalabraMatriz_x(int x);
    // Retorna como string la concatenacion de letras en la fila y (palabra horizontal)
    string buscarPalabraMatriz_y(int y);
    // Elimina el nodo en (x, y) y borra la fila/columna si quedan sin elementos
    void borrarNodo(int, int);
    // Elimina el nodo cabecera de la fila y si ya no tiene elementos
    void borrarFila(int);
    // Elimina el nodo cabecera de la columna x si ya no tiene elementos
    void borrarColumna(int);
    // Libera todos los nodos de datos dejando solo la raiz
    void borrarMatriz();
};






#endif

