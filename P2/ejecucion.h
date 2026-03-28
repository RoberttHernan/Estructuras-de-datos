#ifndef EJECUCION_H
#define EJECUCION_H

#include "token.h"
#include <list>
#include <iterator>
#include <iostream>
#include <algorithm>
namespace ejc {

    /**
     * Parser y ejecutor de comandos.
     * Recibe la lista de tokens producida por el analizador lexico y los
     * despacha al objeto disk correspondiente segun el comando reconocido.
     * Comandos soportados: mkdisk, rmdisk, fdisk, mount, unmount, mkfs,
     *                      login, logout, mkdir, mkfile, rep.
     */
    class ejecucion
    {
    private:

    public:
        // Itera sobre la lista de tokens, identifica cada comando y
        // recopila sus parametros antes de invocar la operacion en disco
        void ejecutarComandos(list<N::token> lista);
    };

}


#endif