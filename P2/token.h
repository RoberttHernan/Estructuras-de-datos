#ifndef TOKEN_H
#define TOKEN_H

#include <string>
using namespace std;

namespace N {

    /**
     * Representa un token producido por el analizador lexico.
     * Cada token tiene un tipo (e.g., "RESMKDISK", "ENTERO", "RUTA_SIMPLE")
     * y el valor del lexema que lo origino.
     */
    class token
    {
    private:
        string tipo_token; // Categoria del token (nombre de la produccion)
        string valor;      // Texto original del lexema
        int numToken;      // Numero de orden del token en la lista

    public:
        token(string tipo_token, string valor);

        string getTipoToken();           // Retorna el tipo del token
        string getValor();               // Retorna el lexema original
        void setTipoToken(string tipo_token); // Asigna el tipo del token
        void setValor(string valor);     // Asigna el valor del lexema
    };
    
   /* token::token()
    {
    }
    
    token::~token()
    {
    }*/
    
}
#endif