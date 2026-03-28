#include "ejecucion.h"
#include <stdio.h>
#include <string.h>
#include <thread>
#include <chrono>
#include "disk.h"
using namespace ejc;
using namespace std;
using namespace N;
using namespace dsk;

void msg_error(int fila, string error)
{
    std::cout << "ERROR, fila: " << fila << " " << error << endl;
}

void NextToken(list<token>::iterator &it, std::string &tipoToken)
{
    it++;
    tipoToken = it->getTipoToken();
}

disk disco_aux;

string toLower_(string entrada)
{
    std::for_each(entrada.begin(), entrada.end(), [](char &c)
                  { c = ::tolower(c); });
    return entrada;
}

void ejecucion::ejecutarComandos(list<token> lista)
{
    list<token>::iterator it;
    int idComando = 0;
    int idParametro = 0;
    int fila = -1;

    for (it = lista.begin(); it != lista.end(); it++)
    {
        fila++;
        std::string tipoComando = it->getTipoToken();

        if (tipoComando == "RESMKDISK")
        {

            idComando = 1;
        }
        else if (tipoComando == "RESRMDISK")
        {
            idComando = 2;
        }
        else if (tipoComando == "RESFDISK")
        {

            idComando = 3;
        }
        else if (tipoComando == "RESREP")
        {

            idComando = 4;
        }
        else if (tipoComando == "RESMOUNT")
        {
            idComando = 5;
        }
        else if (tipoComando == "RESUNMOUNT")
        {
            idComando = 6;
        }
        else if (tipoComando == "RESMKFS")
        {
            idComando = 7;
        }
        else if (tipoComando == "RESLOGIN")
        {
            idComando = 8;
        }
        else if (tipoComando == "RESLOGOUT")
        {
            idComando = 9;
        }
        else if (tipoComando == "RESMKDIR")
        {
            idComando = 10;
        }
        else if (tipoComando == "RESMKFILE")
        {
            idComando = 11;
        }
        else if (tipoComando == "RESPAUSE")
        {
            idComando = 0;
            int flag;
            std::cout << "Presione enter para continuar" << endl;
            flag = cin.get();
        }
        else if (tipoComando == "EOF" || tipoComando == "SL")
        {
            idComando = 0;
            fila++;
        }
        else if (tipoComando == "COMENTARIO")
        {
            idComando = 0;
        }
        else
        {
            cout << "Comando no reconocido, fila: " << fila << endl;
            idComando = 0;
            do
            {
                it++;
                tipoComando = it->getTipoToken();

            } while (tipoComando != "EOF" && tipoComando != "SL");
            std::cout << "Error en tipo de ruta, fila: " << fila << endl;
        }

        switch (idComando)
        {
        case 0:
            break;

        //Mkdisk
        case 1:
        {
            int size = 0;
            string fit = "FF";
            char unit = 'm';
            string path = "";
            bool flag_size = false;
            bool flag_path = false;
            bool flag_unit = true;
            bool flag_fit = true;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string guion = it->getTipoToken();

                if (guion == "GUION")
                {
                    it++;
                    string parametro = it->getTipoToken();

                    if (parametro == "RESSIZE")
                    {
                        it++;
                        parametro = it->getTipoToken();
                        if (parametro == "IGUAL")
                        {
                            it++;
                            parametro = it->getTipoToken();
                            if (parametro == "ENTERO")
                            {
                                size = stoi(it->getValor());

                                if (size > 0)
                                {
                                    flag_size = true;
                                }
                                else
                                {
                                    std::cout << "El size en MKDISK debe ser > 0, fila: " << fila << endl;
                                    std::exit(1);
                                }
                            }
                            else
                            {
                                std::cout << "Se esperaba un numero entero en size, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Se esperaba un = en size, fila: " << fila << endl;
                        }
                    }
                    else if (parametro == "RESUNIT")
                    {
                        it++;
                        parametro = it->getTipoToken();
                        if (parametro == "IGUAL")
                        {
                            it++;
                            parametro = it->getTipoToken();
                            if (parametro == "IDENTIFICADOR")
                            {

                                unit = tolower(it->getValor()[0]);
                                if (unit == 'k' || unit == 'm')
                                {
                                    flag_unit = true;
                                }
                                else
                                {
                                    flag_unit = false;
                                    std::cout << "Error en unit en Mkdisk, fila: " << fila << endl;
                                }
                            }
                            else
                            {
                                std::cout << "Error en parametro unit en fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Error, se esperaba igual en fila: " << fila << endl;
                        }
                    }
                    else if (parametro == "RESFIT")
                    {
                        it++;
                        parametro = it->getTipoToken();
                        if (parametro == "IGUAL")
                        {
                            it++;
                            parametro = it->getTipoToken();
                            if (parametro == "IDENTIFICADOR")
                            {
                                fit = toLower_(it->getValor());
                                if (fit == "bf" || fit == "ff" || fit == "wf")
                                {
                                    flag_fit = true;
                                }
                                else
                                {
                                    flag_fit = false;
                                    std::cout << "Error en fit en fila: " << fila << endl;
                                }
                            }
                            else
                            {
                                flag_fit = false;
                                std::cout << "Error en fit en fila: " << fila << endl;
                            }
                        }
                    }
                    else if (parametro == "RESPATH")
                    {
                        it++;
                        parametro = it->getTipoToken();
                        if (parametro == "IGUAL")
                        {
                            it++;
                            parametro = it->getTipoToken();
                            if (parametro == "RUTA_SIMPLE" || parametro == "RUTA_COMILLAS")

                            {
                                path = it->getValor();
                                flag_path = true;
                            }
                            else
                            {
                                std::cout << "Error en tipo de ruta, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Error, se esperaba igual en fila: " << fila << endl;
                        }
                    }
                    else
                    {
                        cout << "Parametro no reconocido, fila: " << fila << endl;

                        while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }

            if (flag_path == true && flag_fit == true && flag_size == true && flag_unit == true)
            {

                disco_aux.createDisk(size, fit, unit, path);
            }
        }
        break;
        //Rmdisk
        case 2:
        {
            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string stringAux = it->getTipoToken();

                if (stringAux == "GUION")
                {
                    it++;
                    stringAux = it->getTipoToken();

                    if (stringAux == "RESPATH")
                    {
                        it++;
                        stringAux = it->getTipoToken();

                        if (stringAux == "IGUAL")
                        {
                            it++;
                            stringAux = it->getTipoToken();

                            if (stringAux == "RUTA_SIMPLE" || stringAux == "RUTA_COMILLAS")
                            {
                                string path = it->getValor();
                                disco_aux.deleteDisk(path);
                                it++;
                            }
                        }
                        else
                        {
                            cout << "Error, se esperaba signo igual, RMDISK, fila: " << fila << endl;
                        }
                    }
                    else
                    {
                        cout << "Error, se esperaba parametro path en RMDISK, fila: " << fila << endl;
                         while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }
        }
        break;
        //Fdisk
        case 3:
        {

            int size;
            char unit = 'k';
            std::string path;
            char type = 'p';
            std::string fit = "wf";
            char char_fit = 'w';
            std::string name;
            int add = NULL;
            std::string delete_param = "";

            bool flag_size = false;
            bool flag_unit = true;
            bool flag_path = false;
            bool typeflag = true;
            bool flag_fit = true;
            bool delete_flag = false;

            bool flag_name = false;
            bool flag_add = false;
            fila++;
            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {

                it++;
                string tipoToken = it->getTipoToken();
                if (tipoToken == "GUION")
                {
                    it++;
                    string tipoToken = it->getTipoToken();

                    if (tipoToken == "RESSIZE")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "ENTERO")
                            {
                                size = stoi(it->getValor());

                                if (size > 0)
                                {
                                    flag_size = true;
                                }
                                else
                                {
                                    std::cout << "El size en MKDISK debe ser > 0, fila: " << fila << endl;
                                    std::exit(1);
                                }
                            }
                            else
                            {
                                std::cout << "Se esperaba un numero entero en size, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Se esperaba un = en size, fila: " << fila << endl;
                        }
                    }
                    else if (tipoToken == "RESUNIT")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR")
                            {

                                unit = tolower(it->getValor()[0]);
                                if (unit == 'k' || unit == 'm')
                                {
                                    flag_unit = true;
                                }
                                else
                                {
                                    flag_unit = false;
                                    std::cout << "Error en unit en Fdisk, fila: " << fila << endl;
                                }
                            }
                            else
                            {
                                std::cout << "Error en parametro unit en fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Error, se esperaba igual en fila: " << fila << endl;
                        }
                    }
                    else if (tipoToken == "RESPATH")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "RUTA_SIMPLE" || tipoToken == "RUTA_COMILLAS")

                            {
                                path = it->getValor();
                                flag_path = true;
                            }
                            else
                            {
                                std::cout << "Error en tipo de ruta, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            std::cout << "Error, se esperaba igual en fila: " << fila << endl;
                        }
                    }

                    else if (tipoToken == "RESTYPE")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();

                            if (tipoToken == "IDENTIFICADOR")
                            {
                                string valor_type = toLower_(it->getValor());
                                if (valor_type == "p" || valor_type == "e" || valor_type == "l")
                                {
                                    type = valor_type[0];
                                    typeflag = true;
                                }
                                else
                                {
                                    cout << "Error en type, en Fdisk, fila" << fila << endl;
                                    typeflag = false;
                                }
                            }
                            else
                            {
                                cout << "Error en type, en Fdisk, fila" << fila << endl;
                                typeflag = false;
                            }
                        }
                        else
                        {
                            cout << "Error en Fdisk, se esperaba signo igual, fila: " << fila << endl;
                            typeflag = false;
                        }
                    }
                    else if (tipoToken == "RESFIT")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR")
                            {
                                fit = toLower_(it->getValor());
                                if (fit == "bf" || fit == "ff" || fit == "wf")
                                {
                                    if (fit == "bf")
                                    {
                                        char_fit = 'b';
                                    }
                                    else if (fit == "wf")
                                    {
                                        char_fit = 'w';
                                    }
                                    else
                                    {
                                        char_fit = 'f';
                                    }
                                    flag_fit = true;
                                }
                                else
                                {
                                    flag_fit = false;
                                    std::cout << "Error en fit en fila: " << fila << endl;
                                }
                            }
                            else
                            {
                                flag_fit = false;
                                std::cout << "Error en fit en fila: " << fila << endl;
                            }
                        }
                    }
                    else if (tipoToken == "RESDELETE")
                    {
                        it++;
                        std::string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR")
                            {
                                string val = toLower_(it->getValor());

                                if (val == "fast" || val == "full")
                                {
                                    delete_param = val;
                                    delete_flag = true;
                                }
                                else
                                {
                                    std::cout << "Error en delete, FDISK, fila " << fila << endl;
                                    delete_flag = false;
                                }
                            }
                        }
                        else
                        {
                            cout << "Error, se esperaba igual en Delete, Fdisk, fila: " << fila << endl;
                            delete_flag = false;
                        }
                    }
                    else if (tipoToken == "RESNAME")
                    {
                        it++;
                        std::string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            std::string tipoToken = it->getTipoToken();

                            if (tipoToken == "RUTA_COMILLAS" || tipoToken == "IDENTIFICADOR")
                            {

                                name = it->getValor();
                                flag_name = true;
                            }
                            else
                            {
                                cout << "Nombre no valido en Fdisk, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            cout << "Error, se esperaba signo = en Fdisk, Name, fila: " << fila << endl;
                        }
                    }
                    else if (tipoToken == "RESADD")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();

                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();

                            if (tipoToken == "GUION") //signo menos y numero negativos
                            {
                                it++;
                                string tipoToken = it->getTipoToken();
                                if (tipoToken == "ENTERO")
                                {
                                    add = (-1) * stoi((it->getValor()));
                                    flag_add = true;
                                }
                            }
                            else if (tipoToken == "ENTERO")
                            {
                                add = stoi(it->getValor());
                                flag_add = true;
                            }
                            else
                            {
                                cout << "Error en Fdisk, ADD, fila: " << fila << endl;
                                flag_add = false;
                            }
                        }
                        else
                        {
                            cout << "Error en FDISK, ADD, fila: " << fila << endl;
                            flag_add = false;
                        }
                    }
                    else
                    {
                        cout << "Parametro no reconocido, FDisk, fila: " << fila << endl;
                         while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }
            /*cout << flag_size << endl;
            cout << flag_unit << endl;
            cout << flag_path << endl;
            cout << flag_fit << endl;
            cout << typeflag << endl;
            cout << delete_flag << endl;
            cout << flag_name << endl;
            cout << flag_add << endl;*/

            if (delete_flag == true && flag_name == true && flag_path == true && flag_add == false)
            {
                disco_aux.fkDisk(size, unit, path, type, char_fit, delete_param, add, name); //Eliminar Particion
            }
            else if (flag_size == true && flag_unit == true && flag_fit == true && typeflag == true && flag_name == true && flag_path == true && (flag_add == false && delete_flag == false))
            {
                disco_aux.fkDisk(size, unit, path, type, char_fit, delete_param, add, name);
            }
            else if (flag_path == true && flag_name == true && flag_unit == true && flag_add == true && delete_flag == false)
            {
                disco_aux.fkDisk(size, unit, path, type, char_fit, delete_param, add, name); //Añadir o quitar espacio
            }
        }
        break;
        //Rep
        case 4:
        {

            std::string name = "";
            std::string path = "";
            std::string id_param = "";
            std::string ruta = "";

            bool flag_name = false;
            bool flag_path = false;
            bool flag_id = false;
            bool flag_ruta = true;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string tipoToken = it->getTipoToken();

                if (tipoToken == "GUION")
                {
                    it++;
                    string tipoToken = it->getTipoToken();
                    if (tipoToken == "RESNAME")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "RESMBR" || tipoToken == "RESDISK" || tipoToken == "RESSB" || tipoToken == "RESINODE" || tipoToken == "RESBLOCK" || tipoToken == "RESJOURNALING" || tipoToken == "RESTREE" || tipoToken == "RESBMINODE")
                            {
                                name = it->getValor();
                                flag_name = true;
                            }
                            else
                            {
                                msg_error(fila, "Error de tipo de reporte en name");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en parametro name");
                        }
                    }
                    else if (tipoToken == "RESPATH")
                    {

                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "RUTA_SIMPLE" || tipoToken == "RUTA_COMILLAS")

                            {
                                path = it->getValor();
                                flag_path = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba una ruta en Respath");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en parametro name");
                        }
                    }
                    else if (tipoToken == "RESID")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "IDPART")
                            {
                                id_param = it->getValor();
                                flag_id = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba identificador de partición en parametro id");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba Signo igual, en parametro Id");
                        }
                    }
                    else
                    {
                        msg_error(fila, " Parametro no reconocido");
                         while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }

            if (flag_name == true && flag_id == true && flag_path == true && flag_ruta == true)
            {
                disco_aux.rep(toLower_(name), path, id_param, ruta);
            }
        }

        break;

        //mount
        case 5:
        {
            string path;
            string name;
            bool path_flag = false;
            bool name_flag = false;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string stringAux = it->getTipoToken();
                if (stringAux == "GUION")
                {
                    it++;
                    stringAux = it->getTipoToken();

                    if (stringAux == "RESPATH")
                    {
                        it++;
                        stringAux = it->getTipoToken();

                        if (stringAux == "IGUAL")
                        {
                            it++;
                            stringAux = it->getTipoToken();

                            if (stringAux == "RUTA_SIMPLE" || stringAux == "RUTA_COMILLAS")
                            {
                                path = it->getValor();
                                path_flag = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba una ruta en path, comando mount");
                            }
                        }
                        else
                        {
                            msg_error(fila, "se esperaba signo igual en comando mount, parametro path");
                        }
                    }
                    else if (stringAux == "RESNAME")
                    {
                        it++;
                        stringAux = it->getTipoToken();
                        if (stringAux == "IGUAL")
                        {
                            it++;
                            stringAux = it->getTipoToken();

                            if (stringAux == "RUTA_COMILLAS" || stringAux == "IDENTIFICADOR")
                            {

                                name = it->getValor();
                                name_flag = true;
                            }
                            else
                            {
                                cout << "Nombre no valido en mount, fila: " << fila << endl;
                            }
                        }
                        else
                        {
                            cout << "Error, se esperaba signo = en Mount, Name, fila: " << fila << endl;
                        }
                    }
                    else
                    {
                        cout << "Parametro no reconocido, fila: " << fila << endl;

                        while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }

            /*std::cout <<name<<endl; 
            std::cout <<path<<endl; */

            if (name_flag == true && path_flag == true)
            {

                disco_aux.mountPart(path, name);
            }
        }
        break;

        //unmount
        case 6:
        {

            std::string id_param = "";
            bool flag_id = false;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string stringAux = it->getTipoToken();

                if (stringAux == "GUION")
                {
                    it++;
                    stringAux = it->getTipoToken();

                    if (stringAux == "RESID")
                    {
                        it++;
                        stringAux = it->getTipoToken();

                        if (stringAux == "IGUAL")
                        {
                            it++;
                            stringAux = it->getTipoToken();

                            if (stringAux == "IDPART")
                            {
                                id_param = it->getValor();
                                flag_id = true;
                            }
                        }
                        else
                        {
                            msg_error(fila, "no se reconoce el id de partición");
                        }
                    }
                    else
                    {
                        msg_error(fila, "Error en unmount");
                         while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }

            if (flag_id == true)
            {
                disco_aux.unmount(id_param);
            }
        }
        break;

        //Mkfs
        case 7:
        {
            std::string id_param = "";
            std::string type_param = "full";
            int type_fs = 2;

            bool flag_id = false;
            bool flag_type = true;
            bool flag_type_fs = true;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string tipoToken = it->getTipoToken();

                if (tipoToken == "GUION")
                {
                    it++;
                    tipoToken = it->getTipoToken();
                    if (tipoToken == "RESID")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "IDPART")
                            {
                                id_param = it->getValor();
                                flag_id = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba identificador de partición en parametro id");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba Signo igual, en parametro Id");
                        }
                    }
                    else if (tipoToken == "RESTYPE")
                    {
                        it++;
                        std::string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR")
                            {
                                string val = toLower_(it->getValor());

                                if (val == "fast" || val == "full")
                                {
                                    type_param = val;
                                }
                                else
                                {
                                    std::cout << "Error en delete, MKFS, fila " << fila << endl;
                                    flag_type = false;
                                }
                            }
                        }
                        else
                        {
                            cout << "Error, se esperaba igual en type, MKFS, fila: " << fila << endl;
                            flag_type = false;
                        }
                    }
                    else if (tipoToken == "RESFS")
                    {
                        it++;
                        tipoToken = it->getTipoToken();

                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();

                            std::string aux_fs = "";
                            if (tipoToken == "IDPART")
                            {
                                aux_fs += it->getValor();
                                it++;
                                tipoToken = it->getTipoToken();
                                aux_fs += it->getValor();

                                if (aux_fs == "2fs" || aux_fs == "2Fs" || aux_fs == "2FS" || aux_fs == "2fS")
                                {
                                    type_fs = 2;
                                }
                                else if (aux_fs == "3fs" || aux_fs == "3Fs" || aux_fs == "3FS" || aux_fs == "3fS")
                                {
                                    type_fs = 3;
                                }
                                else
                                {
                                    msg_error(fila, "No se reconoce parametro fs en mkfs.fs");
                                    flag_type_fs = false;
                                }
                            }
                            else
                            {
                                msg_error(fila, "No se reconoce parametro fs en mkfs.fs");
                                flag_type_fs = false;
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en mkfs.fs");
                        }
                    }else {
                        msg_error(fila, "No re reconoce parametro en mkfs"); 
                         while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }

            if (flag_id == true && flag_type == true && flag_type_fs == true)
            {
                disco_aux.mkfs(id_param, type_param, type_fs);
            }
        }
        break;

        //Login
        case 8:
        {
            std::string usuario;
            std::string password;
            std::string id_param;

            bool flag_usuario = false;
            bool flag_password = false;
            bool flag_id = false;

            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string tipoToken = it->getTipoToken();

                if (tipoToken == "GUION")
                {
                    it++;
                    tipoToken = it->getTipoToken();

                    if (tipoToken == "RESUSR")
                    {
                        it++;
                        tipoToken = it->getTipoToken();

                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR" || tipoToken == "RUTA_COMILLAS")
                            {
                                usuario = it->getValor();
                                flag_usuario = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba un identificador en login.usr");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en login.usr");
                        }
                    }
                    else if (tipoToken == "RESPWD")
                    {
                        it++;
                        tipoToken = it->getTipoToken();

                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "IDENTIFICADOR" || tipoToken == "ENTERO" || tipoToken == "RUTA_COMILLAS")
                            {
                                password = it->getValor();
                                flag_password = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba contraseña en login.pwd");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en login.usr");
                        }
                    }
                    else if (tipoToken == "RESID")
                    {
                        it++;
                        string tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            string tipoToken = it->getTipoToken();
                            if (tipoToken == "IDPART")
                            {
                                id_param = it->getValor();
                                flag_id = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba identificador en login.id");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba Signo igual logic.id");
                        }
                    }
                    else
                    {
                        cout << "Parametro no reconocido, fila: " << fila << endl;

                        while (tipoComando != "EOF" && tipoComando != "SL")
                        {
                            it++;
                            tipoComando = it->getTipoToken();
                        }
                    }
                }
            }
            if (disco_aux.flag_login == true)
            {
                std::cout << "Hay una sesion iniciada con este usuario " << usuario << endl;
            }

            if (flag_usuario == true && flag_password == true && flag_usuario == true && disco_aux.flag_login == false)
            {
                disco_aux.login(id_param, usuario, password);
            }
        }

        break;

            //Logout

        case 9:
            disco_aux.Logout();
            break;

        //MKDIR
        case 10:
        {

            std::string path = "";
            bool flag_path = false;
            bool flag_p = false;
            bool flag_ok = true;
            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                string tipoToken = it->getTipoToken();

                if (tipoToken == "GUION")
                {
                    it++;
                    string tipoToken = it->getTipoToken();
                    if (tipoToken == "RESPATH")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "RUTA_SIMPLE" || tipoToken == "RUTA_COMILLAS")

                            {
                                path = it->getValor();
                                flag_path = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba ruta en mkdir.path");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo en mkdir.path");
                        }
                    }
                    else if (tipoToken == "IDENTIFICADOR")
                    {
                        if (it->getValor() == "p" || it->getValor() == "P")
                        {
                            flag_p = true;
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba parametro p en mkdir");
                            flag_ok = false;
                        }
                    }
                }
            }

            if (flag_path == true && flag_ok == true)
            {
                disco_aux.ComandoMKDIR(path, flag_p);
            }
        }
        break;

        //MKFILE
        case 11:
        {
            std::string path = "";
            int size = 0;
            std::string path_cont = "";

            bool flag_path = false;
            bool flag_size = true;
            bool flag_count = true;
            bool flag_r = true; 
            while (it->getTipoToken() != "SL" && it->getTipoToken() != "EOF")
            {
                it++;
                std::string tipoToken = it->getTipoToken();

                if (tipoToken == "GUION")
                {
                    it++;
                    tipoToken = it->getTipoToken();
                    if (tipoToken == "RESPATH")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "RUTA_SIMPLE" || tipoToken == "RUTA_COMILLAS")

                            {
                                path = it->getValor();
                                flag_path = true;
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba ruta en mkfile.path");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en mkfile.path");
                        }
                    }
                    else if (tipoToken == "RESSIZE")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "ENTERO")
                            {
                                size = stoi(it->getValor());

                                if (size > 0)
                                {
                                    flag_size = true;
                                }
                                else
                                {
                                    msg_error(fila, "El valor de size en mkfile.size de ser positivo");
                                }
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba numero entero en mkfile.size");
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en mkfile.size");
                        }
                    }
                    else if (tipoToken == "RESCONT")
                    {
                        it++;
                        tipoToken = it->getTipoToken();
                        if (tipoToken == "IGUAL")
                        {
                            it++;
                            tipoToken = it->getTipoToken();
                            if (tipoToken == "RUTA_SIMPLE" || tipoToken == "RUTA_COMILLAS")

                            {
                                path_cont = it->getValor();
                                
                            }
                            else
                            {
                                msg_error(fila, "Se esperaba ruta en mkfile.cont");
                                flag_count = false;
                            }
                        }
                        else
                        {
                            msg_error(fila, "Se esperaba signo igual en mkfile.cont");
                            flag_count = false;
                        }
                    }
                    else if (tipoToken =="IDENTIFICADOR"){
                        if (it->getValor() == "r"|| it->getValor() == "R"){
                            flag_r = true; 
                        }else {
                            flag_r = false; 
                            msg_error (fila, "Parametro no reconocido"); 
                        }
                    }else {
                        msg_error (fila, "Parametro no reconocido en mkfile");
                    }
                }
            }

            if (flag_count == true && flag_path == true && flag_size == true && flag_r == true){
                std::cout << path<<endl; 
                std::cout << path_cont<<endl;
                std::cout << size<<endl;
                std::cout << flag_r <<endl;   
            }
        }
        break;
        default:
            break;
        }
    }
}
