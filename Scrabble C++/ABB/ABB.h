/*
 * Arbol Binario de Busqueda (ABB)
 *
 * Almacena los jugadores registrados ordenados lexicograficamente por nombre de usuario.
 * Permite busqueda, insercion y recorridos en O(log n) promedio.
 *
 * Cada nodo contiene un objeto Jugador con su nombre y su historial de puntajes.
 * Los tres recorridos (inorden, preorden, postorden) generan visualizaciones Graphviz.
 */
#ifndef ABB_H_
#define ABB_H_
#include "NodoArbol.cpp"

using namespace std;
class ABB
{
private:
    NodoArbol *root;    // Raiz del arbol
    NodoArbol *actual;  // Puntero auxiliar usado durante las operaciones de busqueda
    int contador;       // Numero de jugadores insertados
    int altura;         // Altura actual del arbol

public:
    ABB();
    ~ABB();

    // Getters y Setters
    NodoArbol* getRoot();
    void setRoot(NodoArbol *root);
    NodoArbol *getActual();
    void setActual(NodoArbol *actual);
    int getContador();
    void setContador(int contador);
    int getAltura();
    void setAltura(int altura);

    // Inserta un jugador manteniendo el orden del arbol; ignora duplicados
    void Add(Jugador*);
    // Retorna true si existe un jugador con ese nombre de usuario
    bool buscar(string);
    // Retorna el objeto Jugador con ese nombre; retorna un Jugador vacio si no existe
    Jugador buscarJugador(string);
    // Retorna true si el nodo es NULL (arbol o subarbol vacio)
    bool IsEmpty(NodoArbol*);

    // Recorrido Inorden (izquierda, raiz, derecha) - genera texto Graphviz
    string InOrden(NodoArbol *);
    // Genera la imagen Graphviz del recorrido inorden
    void GraficarInorden();
    // Recorrido Preorden (raiz, izquierda, derecha) - genera texto Graphviz
    string Preorden(NodoArbol *);
    // Genera la imagen Graphviz del recorrido preorden
    void GraficarPreorden();
    // Recorrido Postorden (izquierda, derecha, raiz) - genera texto Graphviz
    string PostOrden(NodoArbol *);
    // Genera la imagen Graphviz del recorrido postorden
    void GraficarPostorden();

    // Genera la imagen Graphviz del historial de puntajes de un jugador especifico
    void GraficareportePuntajeJugador(string usuario);

    // Genera la imagen Graphviz con la estructura del arbol (vista simple de jugadores)
    void mostrarJugadoresSimple();
    // Construye recursivamente el texto DOT para graficarArbolSimple
    string graphvizArbolSimple(NodoArbol *nodo);
};






#endif