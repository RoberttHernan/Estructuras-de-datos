#include "analizador.h"

#include "iterator"
#include "cstring"
using namespace N;

bool comparador(string cadena1, string cadena2);
list<token> N::analizador::analizarCadena(string entrada)
{

    entrada += "#";

    estado = 0;

    int columna = 1;
    int fila = 1;

    lexema = "";
    char c;

    int size = entrada.length();

    for (int i = 0; i <= size - 1; i++)
    {
        c = entrada[i];

        columna++;

        switch (estado)
        {

        case 0:
            if (isalpha(c))
            {
                estado = 1;
                lexema += c;
            }
            else if (isdigit(c))
            {
                estado = 2;
                lexema += c;
            }
            else if (c == ' ' || c == '\t')
            {
                estado = 0;
            }
            else if (c == '\n')
            {
                estado = 0;
                fila++;
                columna = 1;
                AddToken("SL");
            }
            else if (c == '-')
            {
                lexema += c;
                AddToken("GUION");
            }
            else if (c == '=')
            {
                lexema += c;
                AddToken("IGUAL");
            }
            else if (c == '/')
            {
                estado = 3;
                lexema += c;
            }
            else if (c == '"')
            {
                estado = 4;
            }
            else if (c == '#' && i != entrada.length() - 1)
            {
                estado = 5;
            }
            else if (c == '#' && i == entrada.length() - 1)
            {

                cout << "Sin errores :)" << endl;
                lexema += c;
                AddToken("EOF");
            }
            else
            {
                estado = -99;
            }
            break;

        case 1:
            if (isalpha(c) || isdigit(c) || c == '_')
            {
                lexema += c;
                estado = 1;
            }
            else
            {

                if (comparador(lexema, "mkdisk"))
                {
                    AddToken("RESMKDISK");
                }
                else if (comparador(lexema, "size"))
                {

                    AddToken("RESSIZE");
                }
                else if (comparador(lexema, "u"))
                {
                    AddToken("RESUNIT");
                }
                else if (comparador(lexema, "f"))
                {
                    AddToken("RESFIT");
                }
                else if (comparador(lexema, "path"))
                {
                    AddToken("RESPATH");
                }
                else if (comparador(lexema, "rmdisk"))
                {

                    AddToken("RESRMDISK");
                }
                else if (comparador(lexema, "mkdir"))
                {

                    AddToken("RESMKDIR");
                }
                else if (comparador(lexema, "fdisk"))
                {
                    AddToken("RESFDISK");
                }
                else if (comparador(lexema, "type"))
                {
                    AddToken("RESTYPE");
                }
                else if (comparador(lexema, "delete"))
                {
                    AddToken("RESDELETE");
                }
                else if (comparador(lexema, "name"))
                {
                    AddToken("RESNAME");
                }
                else if (comparador(lexema, "add"))
                {
                    AddToken("RESADD");
                }
                else if (comparador(lexema, "rep"))
                {
                    AddToken("RESREP");
                }
                else if (comparador(lexema, "id"))
                {
                    AddToken("RESID");
                }
                else if (comparador(lexema, "ruta"))
                {
                    AddToken("RESRUTA");
                }
                else if (comparador(lexema, "mbr"))
                {
                    AddToken("RESMBR");
                }
                else if (comparador(lexema, "disk"))
                {
                    AddToken("RESDISK");
                }
                else if (comparador(lexema, "mkfs"))
                {
                    AddToken("RESMKFS");
                }
                else if (comparador(lexema, "sb"))
                {
                    AddToken("RESSB");
                }
                else if (comparador(lexema, "inode"))
                {
                    AddToken("RESINODE");
                }
                else if (comparador(lexema, "fs"))
                {
                    AddToken("RESFS");
                }
                else if (comparador(lexema, "login"))
                {
                    AddToken("RESLOGIN");
                }
                 else if (comparador(lexema, "usr"))
                {
                    AddToken("RESUSR");
                }
                 else if (comparador(lexema, "pwd"))
                {
                    AddToken("RESPWD");
                }
                else if (comparador(lexema, "logout"))
                {
                    AddToken("RESLOGOUT");
                }
                else if (comparador(lexema, "mount"))
                {
                    AddToken("RESMOUNT");
                }
                else if (comparador(lexema, "unmount"))
                {
                    AddToken("RESUNMOUNT");
                }
                else if (comparador(lexema, "block"))
                {
                    AddToken("RESBLOCK");
                }
                else if (comparador(lexema, "journaling"))
                {
                    AddToken("RESJOURNALING");
                }
                 else if (comparador(lexema, "tree"))
                {
                    AddToken("RESTREE");
                }
                else if (comparador(lexema, "cont"))
                {
                    AddToken("RESCONT");
                }
                else if (comparador(lexema, "pause"))
                {
                    AddToken("RESPAUSE");
                }
                else if (comparador(lexema, "bm_inode"))
                {
                    AddToken("RESBMINODE");
                }
                else if (comparador(lexema, "mkfile"))
                {
                    AddToken("RESMKFILE");
                }
                else
                {
                    AddToken("IDENTIFICADOR");
                }

                lexema = "";
                estado = 0;
                i--;
            }

            break;

        case 2:
            if (isdigit(c))
            {
                lexema += c;
                estado = 2;
            }else if (isalpha (c)){
                lexema += c; 
                AddToken ("IDPART"); 
            }
            else
            {
                AddToken("ENTERO");
                i--;
            }

            break;
        case 3:
            if (c == ' ' || c == '\t' || c == '\n')
            {
                AddToken("RUTA_SIMPLE");
            }
            else
            {
                lexema += c;
                estado = 3;
            }
            break;

        case 4:
            if (c != '"')
            {
                lexema += c;
                estado = 4;
            }
            else
            {
                AddToken("RUTA_COMILLAS");
            }
            break;
        case 5:

            if (c != '\n')
            {
                lexema += c;
                estado = 5;
            }
            else
            {
                AddToken("COMENTARIO");
            }

            break;
        case -99:
            cout << "Ha ocurrido un error lexico en la linea " << endl;
            cout << fila << endl;
            i = size;
            salida.clear();
            break;
        default:
            break;
        }
    }

    return salida;
}

void analizador ::printTokens(list<token> lista)
{
    list<token>::iterator it;
    for (it = lista.begin(); it != lista.end(); it++)
    {
        cout << "Tipo Token: " << it->getTipoToken() << " -Valor: " << it->getValor() << endl;
    }
}

string toLower(string entrada)
{
    int n = entrada.length();
    string salida;
    for (int i = 0; i < entrada.length(); i++)
    {
        salida += tolower(entrada[i]);
    }

    return salida;
}

bool comparador(string cadena1, string cadena2)
{

    int n1 = cadena1.length();
    int n2 = cadena2.length();

    char array1[n1 + 1];
    char array2[n2 + 1];

    strcpy(array1, toLower(cadena1).c_str());
    strcpy(array2, toLower(cadena2).c_str());
    if (strcmp(array1, array2) == 0)
    {

        return true;
    }
    else
    {

        return false;
    }
}