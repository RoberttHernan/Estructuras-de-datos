#ifndef ANALIZADOR_H
#define ANALIZADOR_H

#include "token.h"
#include <list>
#include <bits/stdc++.h>
using namespace std;

namespace N
{
    /**
     * Analizador lexico (scanner) para el interprete de comandos.
     * Implementa una maquina de estados finita con 6 estados:
     *   0 = inicial, 1 = identificador/palabra reservada, 2 = entero,
     *   3 = ruta sin comillas, 4 = ruta con comillas, 5 = comentario.
     * Produce una lista de tokens que el parser (ejecucion) consume.
     */
    class analizador
    {
        string lexema;        // Lexema acumulado durante el escaneo
        int estado;           // Estado actual de la maquina de estados
        list<token> salida;   // Lista de tokens producidos
        list<string> errores; // Mensajes de error lexico detectados

    private:

    public:
        // Escanea la cadena de entrada y retorna la lista de tokens reconocidos
        list<token> analizarCadena(string entrada);

        // Imprime en consola el tipo y valor de cada token de la lista
        void printTokens(list<token>);

       /* analizador::analizador()
        {
        }

        analizador::~analizador()
        {
        }*/
        // Crea un token con el tipo y lexema actuales y lo agrega a la lista de salida
        void AddToken(string tipoToken)
        {
            salida.push_back(token(tipoToken, lexema));
            lexema = "";
            estado = 0;
        }
        
    };


}


#endif