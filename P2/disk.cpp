#include <sys/stat.h>
#include <sys/types.h>
#include <fstream>
#include <vector>
#include <iostream>
#include "disk.h"
#include "mount/listaSimple.h"

using namespace dsk;

bool verificarPath(string path);
vector<string> dividirPath(string entrada);
std::string currentDateTime();
listaSimple *lista_mount = new listaSimple();

string convertToString(char *a, int size)
{
    int i;
    string s = "";
    for (i = 0; i < size; i++)
    {
        s = s + a[i];
    }
    return s;
}

void disk::createDisk(int size, std::string fit, char unit, string path)
{
    mbr prueba;
    string nombre_disco;
    string path_disco;
    FILE *arch;
    vector<string> path_vector = dividirPath(path);
    path_disco = path_vector[0];
    nombre_disco = path_vector[1];
    char buffer[1024];

    for (int i = 0; i < 1024; i++)
    {
        buffer[i] = '\0';
    }

    if (verificarPath(path_disco) == false)
    {
        if (mkdir(path_disco.c_str(), S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH == -1))
        {
            std::cout << "Error al crear carpeta en: " << path_disco << endl;
        }
    }
    arch = fopen(path.c_str(), "wb");
    if (arch == NULL)
    {
        exit(1);
    }

    if (unit == 'k')
    {
        prueba.tamanio_mbr = size * 1024;
        for (int i = 0; i < size; i++)
        {
            std::fwrite(&buffer, 1024, 1, arch);
        }
        std::fclose(arch);
    }
    else if (unit == 'm')
    {
        prueba.tamanio_mbr = size * 1024 * 1024;
        for (int i = 0; i < size * 1024; i++)
        {
            std::fwrite(&buffer, 1024, 1, arch);
        }
        std::fclose(arch);
    }

    string fecha = currentDateTime();

    prueba.signature_mbr = (rand() % 100);
    strcpy(prueba.fecha_creacion_mbr, fecha.c_str());

    if (fit == "bf")
    {
        prueba.fit_disco = 'b';
    }
    else if (fit == "ff")
    {
        prueba.fit_disco = 'f';
    }
    else
    {
        prueba.fit_disco = 'b';
    }

    particion particionAux;
    particionAux.estado_particion = '0';
    particionAux.fit_particion = '0';
    particionAux.tipo_particion = '0';
    particionAux.inicio_particion = -1;
    particionAux.size_particion = 0;
    particionAux.nombre_particion[0] = '\0';
    for (int i = 0; i < 4; i++)
    {
        prueba.particiones_mbr[i] = particionAux;
    }

    arch = fopen(path.c_str(), "rb+");

    if (arch != NULL)
    {
        std::fseek(arch, 0, SEEK_SET);
        std::fwrite(&prueba, sizeof(mbr), 1, arch);
        std::fclose(arch);
        cout << "Disco creado correctamente,: " << path << " ,MBR agreagado correctamente" << endl;
    }
    else
    {
        cout << "Error al acceder a disco: " << nombre_disco << "MBR no creado" << endl;
    }
}

void disk::deleteDisk(std::string path)
{
    ifstream arch(path);
    if (arch.good())
    {
        remove(path.c_str());
        cout << "Disco " << path << " Eliminado correctamente" << endl;
        arch.close();
    }
    else
    {
        cout << "Error al eliminar disco: " << path << endl;
    }
}

void disk::fkDisk(int size, char unit, std::string path, char type, char fit, std::string delete_param, int add, string name)
{
    if (delete_param != "")
    {
        deletePart(path, name, delete_param);
    }
    else if (add != NULL)
    {
        Add(path, name, unit, add);
    }
    else
    {
        createPart(path, size, unit, type, fit, name);
    }
}
void disk::createPart(string path, int size, char unit, char type, char fit, string name)
{
    if (verificarPath(path) == false)
    {
        std::cout << "Error al buscar disco " << path << endl;
        return;
    }
    if (existPart(path, name))
    {
        std::cout << "La particion: " << name << " En el disco: " << path << " Ya existe" << endl;
        return;
    }

    if (type == 'p')
    {
        createPrimaryPart(path, size, unit, fit, name, type);
    }
    else if (type == 'e')
    {
        createExtendedPart(path, size, unit, name, fit);
    }
    else if (type == 'l')
    {
        createLogicPart(path, size, unit, fit, name);
    }
}
bool disk::existPart(string path, std::string name)
{
    if (verificarPath(path) == false)
    {
        return false;
    }

    mbr temp_MBR;
    int extended = -1;
    FILE *arch = fopen(path.c_str(), "rb+");
    std::fseek(arch, 0, SEEK_SET);
    fread(&temp_MBR, sizeof(mbr), 1, arch);
    for (int i = 0; i < 4; i++)
    {
        if (temp_MBR.particiones_mbr[i].nombre_particion == name)
        {
            std::fclose(arch);
            return true;
        }
        else if (temp_MBR.particiones_mbr[i].tipo_particion == 'e')
        {
            extended = i;
        }
    }
    if (extended != -1)
    {
        std::fseek(arch, temp_MBR.particiones_mbr[extended].inicio_particion, SEEK_SET);
        ebr extBoot;
        while ((fread(&extBoot, sizeof(ebr), 1, arch)) != 0 && (ftell(arch) < (temp_MBR.particiones_mbr[extended].size_particion + temp_MBR.particiones_mbr[extended].inicio_particion)))
        {
            if (extBoot.nombre_particion == name)
            {
                std::fclose(arch);
                return true;
            }
            if (extBoot.sig_particion == -1)
            {
                std::fclose(arch);
                return false;
            }
        }
    }

    std::fclose(arch);
    return false;
}
void disk::createPrimaryPart(string path, int size, char unit, char fit, string name, char type)
{

    int size_bytes = 1024;
    char buffer = '1';

    if (unit == 'b')
    {
        size_bytes = size;
    }
    else if (unit == 'k')
    {
        size_bytes = size * 1024;
    }
    else
    {
        size_bytes = size * 1024 * 1024;
    }

    FILE *arch = fopen(path.c_str(), "rb+");
    mbr masterBoot;

    bool flag_particion = false;
    int num_particion = 0;

    std::fseek(arch, 0, SEEK_SET);
    fread(&masterBoot, sizeof(mbr), 1, arch);

    for (int i = 0; i < 4; i++)
    {

        if (masterBoot.particiones_mbr[i].inicio_particion == -1 || (masterBoot.particiones_mbr[i].estado_particion == '1' && masterBoot.particiones_mbr[i].size_particion >= size_bytes))
        {
            flag_particion = true;
            num_particion = i;
            break;
        }
    }

    if (flag_particion)
    {
        int espaciousado = 0;
        for (particion p : masterBoot.particiones_mbr)
        {
            if (p.estado_particion != '1')
            {
                espaciousado += p.size_particion;
            }
        }

        if ((masterBoot.tamanio_mbr - espaciousado) >= size_bytes)
        {
            if (!existPart(path, name))
            {
                //Primer ajuste
                if (masterBoot.fit_disco == 'f')
                {
                    masterBoot.particiones_mbr[num_particion].tipo_particion = 'p';
                    masterBoot.particiones_mbr[num_particion].fit_particion = fit;

                    if (num_particion == 0)
                    {
                        masterBoot.particiones_mbr[num_particion].inicio_particion = sizeof(masterBoot);
                    }
                    else
                    {
                        int inicio_particion_aux = masterBoot.particiones_mbr[num_particion - 1].inicio_particion + masterBoot.particiones_mbr[num_particion - 1].size_particion;
                        masterBoot.particiones_mbr[num_particion].inicio_particion = inicio_particion_aux;
                    }
                    masterBoot.particiones_mbr[num_particion].size_particion = size_bytes;
                    masterBoot.particiones_mbr[num_particion].estado_particion = '0';
                    strcpy(masterBoot.particiones_mbr[num_particion].nombre_particion, name.c_str());
                    //Guardando el nuevo MBR
                    std::fseek(arch, 0, SEEK_SET);
                    std::fwrite(&masterBoot, sizeof(mbr), 1, arch);
                    //Guardando lo bytes de la particion
                    std::fseek(arch, masterBoot.particiones_mbr[num_particion].inicio_particion, SEEK_SET);

                    for (int i = 0; i < size_bytes; i++)
                    {
                        std::fwrite(&buffer, 1, 1, arch);
                    }
                    std::cout << "Particion " << name << " Creada en " << path << endl;
                }
                //Mejor ajuste
                else if (masterBoot.fit_disco == 'b')
                {
                    int mejor = num_particion;
                    int n = 0;
                    for (particion p : masterBoot.particiones_mbr)
                    {
                        if (p.inicio_particion == -1 || (p.estado_particion == '1' && p.size_particion >= size_bytes))
                        {
                            if (n != num_particion)
                            {
                                if (masterBoot.particiones_mbr[mejor].size_particion > p.size_particion)
                                {
                                    mejor = n;
                                    break;
                                }
                            }
                        }
                        n++;
                    }

                    masterBoot.particiones_mbr[mejor].tipo_particion = 'p';
                    masterBoot.particiones_mbr[mejor].fit_particion = fit;

                    if (mejor == 0)
                    {
                        masterBoot.particiones_mbr[mejor].inicio_particion = sizeof(masterBoot);
                    }
                    else
                    {
                        int mbrStart = masterBoot.particiones_mbr[mejor - 1].inicio_particion + masterBoot.particiones_mbr[mejor - 1].size_particion;
                        masterBoot.particiones_mbr[mejor].inicio_particion = mbrStart;
                    }
                    masterBoot.particiones_mbr[mejor].size_particion = size_bytes;
                    masterBoot.particiones_mbr[mejor].estado_particion = '0';
                    strcpy(masterBoot.particiones_mbr[mejor].nombre_particion, name.c_str());
                    //Guardamos el mbr otra vez xd
                    std::fseek(arch, 0, SEEK_SET);
                    std::fwrite(&masterBoot, sizeof(mbr), 1, arch);
                    //Guardamos los bytes de la particion
                    std::fseek(arch, masterBoot.particiones_mbr[mejor].inicio_particion, SEEK_SET);
                    for (int i = 0; i < size_bytes; i++)
                    {
                        std::fwrite(&buffer, 1, 1, arch);
                    }
                    std::cout << "Particion " << name << "Creada en " << path << endl;
                }
                //Peor ajuste
                else if (masterBoot.fit_disco == 'w')
                {
                    int peor = num_particion;
                    int n = 0;
                    for (particion p : masterBoot.particiones_mbr)
                    {
                        if (p.inicio_particion == -1 || (p.estado_particion == '1' && p.size_particion >= size_bytes))
                        {
                            if (n != num_particion)
                            {
                                if (masterBoot.particiones_mbr[peor].size_particion < p.size_particion)
                                {
                                    peor = n;
                                    break;
                                }
                            }
                        }
                        n++;
                    }
                    masterBoot.particiones_mbr[peor].tipo_particion = 'p';
                    masterBoot.particiones_mbr[peor].fit_particion = fit;

                    if (peor == 0)
                    {
                        masterBoot.particiones_mbr[peor].inicio_particion = sizeof(masterBoot);
                    }
                    else
                    {
                        int mbrStart = masterBoot.particiones_mbr[peor - 1].inicio_particion + masterBoot.particiones_mbr[peor - 1].size_particion;
                        masterBoot.particiones_mbr[peor].inicio_particion = mbrStart;
                    }

                    masterBoot.particiones_mbr[peor].size_particion = size_bytes;
                    masterBoot.particiones_mbr[peor].estado_particion = '0';
                    strcpy(masterBoot.particiones_mbr[peor].nombre_particion, name.c_str());
                    //Guardamos el mbr otra vez xd
                    std::fseek(arch, 0, SEEK_SET);
                    std::fwrite(&masterBoot, sizeof(mbr), 1, arch);
                    //Guardamos los bytes de la particion
                    std::fseek(arch, masterBoot.particiones_mbr[peor].inicio_particion, SEEK_SET);
                    for (int i = 0; i < size_bytes; i++)
                    {
                        std::fwrite(&buffer, 1, 1, arch);
                    }
                    std::cout << "Particion " << name << "Creada en " << path << endl;
                }
            }
            else
            {
                std::cout << "Error, en " << path << "Ya existe la particion: " << name << endl;
            }
        }
        else
        {
            std::cout << "Error al crear particion: " << name << " en " << path << endl;
            std::cout << "No hay suficiente espacio disponible" << endl;
        }
    }
    else
    {
        std::cout << "ERROR, ya existen 4 particiones en el disco " << path << endl;
    }
    std::fclose(arch);
}
void disk::createExtendedPart(std::string path, int size, char unit, std::string name, char fit)
{
    int size_bytes = 1024;
    char buffer = '1';

    FILE *arch;
    arch = fopen(path.c_str(), "rb+");
    mbr temp_mbr;

    if (unit == 'b')
    {
        size_bytes = size;
    }
    else if (unit == 'k')
    {
        size_bytes = size * 1024;
    }
    else
    {
        size_bytes = size * 1024 * 1024;
    }

    bool flag_particion = false;
    bool flag_extendida = false;
    int num_particion = 0;

    std::fseek(arch, 0, SEEK_SET);
    fread(&temp_mbr, sizeof(mbr), 1, arch);

    for (int i = 0; i < 4; i++)
    {
        if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
        {
            flag_extendida = true;
            break;
        }
    }

    if (!flag_extendida)
    {
        for (int i = 0; i < 4; i++)
        {

            if (temp_mbr.particiones_mbr[i].inicio_particion == -1 || (temp_mbr.particiones_mbr[i].estado_particion == '1' && temp_mbr.particiones_mbr[i].size_particion >= size_bytes))
            {
                flag_particion = true;
                num_particion = i;
                break;
            }
        }
        if (flag_particion)
        {
            int espaciousado = 0;
            for (int i = 0; i < 4; i++)
            {
                if (temp_mbr.particiones_mbr[i].estado_particion != '1')
                {
                    espaciousado += temp_mbr.particiones_mbr[i].size_particion;
                }
            }

            if ((temp_mbr.tamanio_mbr - espaciousado) >= size_bytes)
            {
                if (existPart(path, name) == false)
                {
                    //primer ajuste
                    if (temp_mbr.fit_disco == 'f')
                    {
                        temp_mbr.particiones_mbr[num_particion].tipo_particion = 'e';
                        temp_mbr.particiones_mbr[num_particion].fit_particion = fit;

                        if (num_particion == 0)
                        {
                            temp_mbr.particiones_mbr[num_particion].inicio_particion = sizeof(temp_mbr);
                        }
                        else
                        {
                            temp_mbr.particiones_mbr[num_particion].inicio_particion = temp_mbr.particiones_mbr[num_particion - 1].inicio_particion + temp_mbr.particiones_mbr[num_particion - 1].size_particion;
                        }
                        temp_mbr.particiones_mbr[num_particion].size_particion = size_bytes;
                        temp_mbr.particiones_mbr[num_particion].estado_particion = '0';
                        strcpy(temp_mbr.particiones_mbr[num_particion].nombre_particion, name.c_str());
                        //Guardando el nuevo MBR
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        //Guardando lo bytes de la particion extendida
                        std::fseek(arch, temp_mbr.particiones_mbr[num_particion].inicio_particion, SEEK_SET);

                        ebr temp_eb;
                        temp_eb.fit_particion = fit;
                        temp_eb.estado_particion = '0';
                        temp_eb.inicio_particion = temp_mbr.particiones_mbr[num_particion].inicio_particion;
                        temp_eb.tamanio_particion = 0;
                        temp_eb.sig_particion = -1;
                        strcpy(temp_eb.nombre_particion, "");
                        std::fwrite(&temp_eb, sizeof(ebr), 1, arch);

                        for (int i = 0; i < (size_bytes - (int)sizeof(ebr)); i++)
                        {
                            std::fwrite(&buffer, 1, 1, arch);
                        }

                        std::cout << "Particion extendida " << name << " Creada en " << path << endl;
                    }
                    else if (temp_mbr.fit_disco == 'b')
                    {
                        int mejor = num_particion;
                        for (int i = 0; i < 4; i++)
                        {
                            if (temp_mbr.particiones_mbr[i].inicio_particion == -1 || (temp_mbr.particiones_mbr[i].estado_particion == '1' && temp_mbr.particiones_mbr[i].size_particion >= size_bytes))
                            {
                                if (i != num_particion)
                                {
                                    if (temp_mbr.particiones_mbr[mejor].size_particion > temp_mbr.particiones_mbr[i].size_particion)
                                    {
                                        mejor = i;
                                    }
                                }
                            }
                        }

                        temp_mbr.particiones_mbr[mejor].tipo_particion = 'e';
                        temp_mbr.particiones_mbr[mejor].fit_particion = fit;
                        if (mejor == 0)
                        {
                            temp_mbr.particiones_mbr[mejor].inicio_particion = sizeof(temp_mbr);
                        }
                        else
                        {
                            temp_mbr.particiones_mbr[mejor].inicio_particion = temp_mbr.particiones_mbr[mejor - 1].inicio_particion + temp_mbr.particiones_mbr[mejor - 1].size_particion;
                        }

                        temp_mbr.particiones_mbr[mejor].size_particion = size_bytes;
                        temp_mbr.particiones_mbr[mejor].estado_particion = '0';
                        strcpy(temp_mbr.particiones_mbr[mejor].nombre_particion, name.c_str());

                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        //Guardando lo bytes de la particion extendida
                        std::fseek(arch, temp_mbr.particiones_mbr[mejor].inicio_particion, SEEK_SET);

                        ebr temp_eb;
                        temp_eb.fit_particion = fit;
                        temp_eb.estado_particion = '0';
                        temp_eb.inicio_particion = temp_mbr.particiones_mbr[mejor].inicio_particion;
                        temp_eb.tamanio_particion = 0;
                        temp_eb.sig_particion = -1;
                        strcpy(temp_eb.nombre_particion, "");
                        std::fwrite(&temp_eb, sizeof(ebr), 1, arch);

                        for (int i = 0; i < (size_bytes - (int)sizeof(ebr)); i++)
                        {
                            std::fwrite(&buffer, 1, 1, arch);
                        }

                        std::cout << "Particion extendida " << name << " Creada en " << path << endl;
                    }
                    else if (temp_mbr.fit_disco == 'w')
                    {
                        int peor = num_particion;
                        int n = 0;
                        for (particion p : temp_mbr.particiones_mbr)
                        {
                            if (p.inicio_particion == -1 || (p.estado_particion == '1' && p.size_particion >= size_bytes))
                            {
                                if (n != num_particion)
                                {
                                    if (temp_mbr.particiones_mbr[peor].size_particion < p.size_particion)
                                    {
                                        peor = n;
                                        break;
                                    }
                                }
                            }
                            n++;
                        }
                        temp_mbr.particiones_mbr[peor].tipo_particion = 'e';
                        temp_mbr.particiones_mbr[peor].fit_particion = fit;

                        if (peor == 0)
                        {
                            temp_mbr.particiones_mbr[peor].inicio_particion = sizeof(temp_mbr);
                        }
                        else
                        {
                            int mbrStart = temp_mbr.particiones_mbr[peor - 1].inicio_particion + temp_mbr.particiones_mbr[peor - 1].size_particion;
                            temp_mbr.particiones_mbr[peor].inicio_particion = mbrStart;
                        }

                        temp_mbr.particiones_mbr[peor].size_particion = size_bytes;
                        temp_mbr.particiones_mbr[peor].estado_particion = '0';
                        strcpy(temp_mbr.particiones_mbr[peor].nombre_particion, name.c_str());
                        //Guardando el nuevo MBR
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        //Guardando lo bytes de la particion extendida
                        std::fseek(arch, temp_mbr.particiones_mbr[peor].inicio_particion, SEEK_SET);

                        ebr temp_eb;
                        temp_eb.fit_particion = fit;
                        temp_eb.estado_particion = '0';
                        temp_eb.inicio_particion = temp_mbr.particiones_mbr[peor].inicio_particion;
                        temp_eb.tamanio_particion = 0;
                        temp_eb.sig_particion = -1;
                        strcpy(temp_eb.nombre_particion, "");
                        std::fwrite(&temp_eb, sizeof(ebr), 1, arch);

                        for (int i = 0; i < (size_bytes - sizeof(ebr)); i++)
                        {
                            std::fwrite(&buffer, 1, 1, arch);
                        }

                        std::cout << "Particion extendida " << name << " Creada en " << path << endl;
                    }
                }
                else
                {
                    std::cout << "Error, en " << path << "Ya existe la particion: " << name << endl;
                }
            }
            else
            {
                std::cout << "Error al crear particion: " << name << " en " << path << endl;
                std::cout << "No hay suficiente espacio disponible" << endl;
            }
        }
        else
        {
            std::cout << "ERROR, ya existen 4 particiones en el disco " << path << endl;
        }
    }
    else
    {
        std::cout << "ERROR, Ya existe una particion extendida en " << path << endl;
    }
    std::fclose(arch);
}
void disk::createLogicPart(std::string path, int size, char unit, char fit, std::string name)
{
    int size_bytes = 1024;

    char buffer = '1';

    FILE *arch;
    arch = fopen(path.c_str(), "rb+");
    mbr temp_mbr;

    if (unit == 'b')
    {
        size_bytes = size;
    }
    else if (unit == 'k')
    {
        size_bytes = size * 1024;
    }
    else
    {
        size_bytes = size * 1024 * 1024;
    }

    int num_extended = -1;
    std::fseek(arch, 0, SEEK_SET);
    fread(&temp_mbr, sizeof(mbr), 1, arch);

    for (int i = 0; i < 4; i++)
    {
        if (temp_mbr.particiones_mbr[i].tipo_particion == 'e' || temp_mbr.particiones_mbr[i].tipo_particion == 'E')
        {
            num_extended = i;
            break;
        }
    }
    if (!existPart(path, name))
    {
        if (num_extended != -1)
        {
            ebr ext_boot;
            int aux = temp_mbr.particiones_mbr[num_extended].inicio_particion;
            std::fseek(arch, aux, SEEK_SET);
            fread(&ext_boot, sizeof(ebr), 1, arch);

            if (ext_boot.tamanio_particion == 0) //particion vacia
            {
                if (temp_mbr.particiones_mbr[num_extended].size_particion < size_bytes)
                {

                    std::cout << "Error, el tamaño de la particion logica " << name << "Excede el espacio disponible en la particion " << endl;
                }
                else
                {
                    ext_boot.estado_particion = '0';
                    ext_boot.fit_particion = fit;
                    ext_boot.inicio_particion = ftell(arch) - sizeof(ebr);
                    ext_boot.tamanio_particion = size_bytes;
                    ext_boot.sig_particion = -1;
                    strcpy(ext_boot.nombre_particion, name.c_str());
                    std::fseek(arch, temp_mbr.particiones_mbr[num_extended].inicio_particion, SEEK_SET);
                    std::fwrite(&ext_boot, sizeof(ebr), 1, arch);
                    cout << "Particion logica " << name << "Creada con exito en " << path << endl;
                }
            }
            else
            {
                while ((ext_boot.sig_particion != -1) && (ftell(arch) < (temp_mbr.particiones_mbr[num_extended].size_particion + temp_mbr.particiones_mbr[num_extended].inicio_particion)))
                {
                    std::fseek(arch, ext_boot.sig_particion, SEEK_SET);
                    fread(&ext_boot, sizeof(ebr), 1, arch);
                }
                int esp = ext_boot.inicio_particion + ext_boot.tamanio_particion + size_bytes;

                if (esp <= (temp_mbr.particiones_mbr[num_extended].size_particion + temp_mbr.particiones_mbr[num_extended].inicio_particion))
                {
                    ext_boot.sig_particion = ext_boot.inicio_particion + ext_boot.tamanio_particion;
                    std::fseek(arch, ftell(arch) - sizeof(ebr), SEEK_SET);
                    std::fwrite(&ext_boot, sizeof(ebr), 1, arch);
                    std::fseek(arch, ext_boot.inicio_particion + ext_boot.tamanio_particion, SEEK_SET);

                    ext_boot.estado_particion = '0';
                    ext_boot.fit_particion = fit;
                    ext_boot.inicio_particion = ftell(arch);
                    ext_boot.tamanio_particion = size_bytes;
                    ext_boot.sig_particion = -1;
                    strcpy(ext_boot.nombre_particion, name.c_str());
                    std::fwrite(&ext_boot, sizeof(ebr), 1, arch);

                    cout << "Particion logica " << name << "Creada con exito en " << path << endl;
                }
                else
                {

                    std::cout << "ERROR!, el tamaño de la particion logica " << name << " es mas grande que el espacio disponible en la particion extendida de " << path << endl;
                }
            }
        }
        else
        {

            std::cout << "ERROR, no existe un particion logica en donde guardar " << name << " en el disco" << path << endl;
        }
        std::fclose(arch);
    }

    else
    {

        std::cout << "ERROR, ya existe una particion con este nombre " << name << " en el disco " << path << endl;
    }
}
void disk::imprimirDatosDisco(string path)
{
    FILE *arch;
    arch = fopen(path.c_str(), "rb+");
    if (arch != NULL)
    {
        mbr MBR;
        std::fseek(arch, 0, SEEK_SET);
        fread(&MBR, sizeof(mbr), 1, arch);
        std::fclose(arch);

        std::cout << "-----------Datos del disco------------" << endl;
        std::cout << "Nombre MBR: " << MBR.signature_mbr << endl;
        std::cout << "Tamanio MBR: " << MBR.tamanio_mbr << endl;

        for (int i = 0; i < 4; i++)
        {
            if (MBR.particiones_mbr[i].estado_particion != '1')
            {
                std::cout << "Particion: " << i << endl;
                std::cout << "Estado Particion: " << MBR.particiones_mbr[i].estado_particion << endl;
                std::cout << "Tipo Particion: " << MBR.particiones_mbr[i].tipo_particion << endl;
                std::cout << "Fit Particion: " << MBR.particiones_mbr[i].fit_particion << endl;
                std::cout << "Inicio Particion: " << MBR.particiones_mbr[i].inicio_particion << endl;
                std::cout << "Tamanio Particion: " << MBR.particiones_mbr[i].size_particion << endl;
                std::cout << "Nombre Particion: " << MBR.particiones_mbr[i].nombre_particion << endl;
            }
        }
    }
}
void disk::deletePart(std::string path, std::string name, std::string type_delete)
{
    FILE *arch;
    arch = fopen(path.c_str(), "rb+");
    bool is_mount = lista_mount->findNode(path, name);
    if (is_mount == false)
    {
        mbr temp_mbr;
        std::fseek(arch, 0, SEEK_SET);
        fread(&temp_mbr, sizeof(mbr), 1, arch);
        int indice = -1;
        int indice_extendida = -1;
        bool flag_extendida = false;
        string opcion = "";
        char buffer = '\0';

        //Buscando la partición
        for (int i = 0; i < 4; i++)
        {
            if (temp_mbr.particiones_mbr[i].nombre_particion == name)
            {
                indice = i;
                if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
                {
                    flag_extendida = true;
                }
                break;
            }
            else if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
            {
                indice_extendida = i;
            }
        }
        std::cout << "Desea eliminar la particion " << name << " en " << path << " S/N" << endl;
        std::getline(cin, opcion);

        if (opcion == "S" || opcion == "s")
        {
            //es una particion principal
            if (indice != -1)
            {
                //Puedo eliminar el if xd probamos luego
                //Particion primaria
                if (!flag_extendida)
                {
                    if (type_delete == "fast")
                    {
                        temp_mbr.particiones_mbr[indice].estado_particion = '1';
                        strcpy(temp_mbr.particiones_mbr[indice].nombre_particion, "");
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        std::cout << "Particion primaria " << name << " eliminada en fast de " << path << endl;
                    }
                    else
                    {
                        temp_mbr.particiones_mbr[indice].estado_particion = '1';
                        strcpy(temp_mbr.particiones_mbr[indice].nombre_particion, "");
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        std::fseek(arch, temp_mbr.particiones_mbr[indice].inicio_particion, SEEK_SET);
                        std::fwrite(&buffer, 1, temp_mbr.particiones_mbr[indice].size_particion, arch);
                        std::cout << "Particion primaria " << name << " eliminada en full de " << path << endl;
                    }
                }
                //PArticion extendida
                else
                {

                    if (type_delete == "fast")
                    {
                        temp_mbr.particiones_mbr[indice].estado_particion = '1';
                        strcpy(temp_mbr.particiones_mbr[indice].nombre_particion, "");
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        std::cout << "Particion extendida " << name << " eliminada en fast de " << path << endl;
                    }
                    else
                    {
                        temp_mbr.particiones_mbr[indice].estado_particion = '1';
                        strcpy(temp_mbr.particiones_mbr[indice].nombre_particion, "");
                        std::fseek(arch, 0, SEEK_SET);
                        std::fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        std::fseek(arch, temp_mbr.particiones_mbr[indice].inicio_particion, SEEK_SET);
                        std::fwrite(&buffer, 1, temp_mbr.particiones_mbr[indice].size_particion, arch);
                        std::cout << "Particion extendida " << name << " eliminada en full de " << path << endl;
                    }
                }
            }
            //es una particion logica
            else
            {
                if (indice_extendida != -1)
                {
                    bool flag_exist = false;
                    ebr extender_boot;
                    std::fseek(arch, temp_mbr.particiones_mbr[indice_extendida].inicio_particion, SEEK_SET);
                    fread(&extender_boot, sizeof(ebr), 1, arch);
                    if (extender_boot.tamanio_particion != 0)
                    {
                        std::fseek(arch, temp_mbr.particiones_mbr[indice_extendida].inicio_particion, SEEK_SET);
                        while ((fread(&extender_boot, sizeof(ebr), 1, arch)) != 0 && (ftell(arch) < (temp_mbr.particiones_mbr[indice_extendida].inicio_particion + temp_mbr.particiones_mbr[indice_extendida].size_particion)))
                        {
                            if (extender_boot.nombre_particion == name && extender_boot.estado_particion != '1')
                            {
                                flag_exist = true;
                                break;
                            }
                            else if (extender_boot.sig_particion == -1)
                            {
                                break;
                            }
                        }
                    }

                    if (flag_exist)
                    {
                        if (type_delete == "fast")
                        {
                            extender_boot.estado_particion = '1';
                            strcpy(extender_boot.nombre_particion, "");
                            std::fseek(arch, ftell(arch) - sizeof(ebr), SEEK_SET);
                            std::fwrite(&extender_boot, sizeof(ebr), 1, arch);
                            std::cout << "Particion logica " << name << " eliminada en fast de " << path << endl;
                        }
                        else
                        {
                            extender_boot.estado_particion = '1';
                            strcpy(extender_boot.nombre_particion, "");
                            std::fseek(arch, ftell(arch) - sizeof(ebr), SEEK_SET);
                            std::fwrite(&extender_boot, sizeof(ebr), 1, arch);
                            std::fwrite(&buffer, 1, extender_boot.tamanio_particion, arch);
                            std::cout << "Particion logica " << name << " eliminada en full de " << path << endl;
                        }
                    }
                    else
                    {
                        std::cout << "ERROR, no se encontro la particon " << name << " en " << path << endl;
                    }
                }
                else
                {
                    std::cout << "ERROR, no se encontro la particon " << name << " en " << path << endl;
                }
            }
        }
        else if (opcion == "N" || opcion == "n")
        {
        }
        else
        {
            std::cout << "ERROR, opcion incorrecta" << endl;
        }
    }
    else
    {
        std::cout << "ERROR, debe desmontar " << name << " en " << path << " antes de eliminar" << std::endl;
    }
    std::fclose(arch);
}
bool verificarPath(string path)
{
    ifstream f(path.c_str());
    f.close();
    return f.good();
}
int disk::createCarpet(std::string path_carpeta, bool p_param)
{
    FILE *arch = fopen(sesion_activa.direccion.c_str(), "rb+");

    SuperBloque sp;
    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);

    string aux = path_carpeta;
    char auxPath[500];
    strcpy(auxPath, aux.c_str());
    int existe = existFileorDir(arch, auxPath);
    strcpy(auxPath, aux.c_str());
    int res = -1;
    if (existe != -1)
        res = 0;
    else
        res = MakeNewDIr(arch, sesion_activa.fit, p_param, auxPath, 0);

    fclose(arch);

    return res;
}
void disk::Add(std::string path, std::string name, char unit, int add_param)
{
    /*std::cout <<"Path:" <<path<<endl; 
    std::cout <<"Name:" <<name<<endl;
    std::cout <<"Units:" <<unit<<endl;
    std::cout <<"To Add:" <<add_param<<endl;*/

    int size_bytes = 0;
    FILE *arch;
    arch = fopen(path.c_str(), "rb+");
    string tipo = "";

    if (add_param > 0)
    {
        tipo = "add";
    }

    if (unit == 'b')
    {
        size_bytes = add_param;
    }
    else if (unit == 'k')
    {
        size_bytes = add_param * 1024;
    }
    else
    {
        size_bytes = add_param * 1024 * 1024;
    }

    bool is_mount = lista_mount->findNode(path, name);
    if (is_mount == false)
    {
        mbr temp_mbr;
        fseek(arch, 0, SEEK_SET);
        fread(&temp_mbr, sizeof(mbr), 1, arch);
        int indice = -1;
        int indice_extend = -1;
        bool flag_extendida = false;

        for (int i = 0; i < 4; i++)
        {
            if (temp_mbr.particiones_mbr[i].nombre_particion == name)
            {
                indice = i;
                if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
                {
                    flag_extendida = true;
                }
                break;
            }
            else if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
            {
                indice_extend = i;
            }
        }
        if (indice != -1)
        {
            if (flag_extendida == false)
            {
                //Particion Primaria
                if (tipo == "add")
                {
                    //espacio a la dercha
                    if (indice != 3)
                    {
                        int aux1 = temp_mbr.particiones_mbr[indice].inicio_particion + temp_mbr.particiones_mbr[indice].size_particion;
                        int aux2 = temp_mbr.particiones_mbr[indice + 1].inicio_particion;

                        if (aux2 - aux1 != 0)
                        {
                            if ((aux2 - aux1) >= size_bytes)
                            {
                                temp_mbr.particiones_mbr[indice].size_particion += size_bytes;

                                fseek(arch, 0, SEEK_SET);
                                fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                            }
                            else
                            {
                                std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                            }
                        }
                        else
                        {
                            //Espacio libre
                            if (temp_mbr.particiones_mbr[indice + 1].estado_particion == '1')
                            {
                                if (temp_mbr.particiones_mbr[indice + 1].size_particion >= size_bytes)
                                {
                                    temp_mbr.particiones_mbr[indice].size_particion += size_bytes;
                                    temp_mbr.particiones_mbr[indice + 1].size_particion -= size_bytes;
                                    temp_mbr.particiones_mbr[indice + 1].inicio_particion += size_bytes;
                                    fseek(arch, 0, SEEK_SET);
                                    fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                    std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                                }
                                else
                                {
                                    std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                                }
                            }
                        }
                    }
                    else
                    {
                        //No hay espacio a la deracha xd
                        int sp = temp_mbr.particiones_mbr[indice].inicio_particion + temp_mbr.particiones_mbr[indice].size_particion;
                        int total = temp_mbr.tamanio_mbr + (int)sizeof(mbr); //verificamos el espacio total
                        if ((total - sp) != 0)
                        {
                            if ((total - sp) >= size_bytes)
                            {
                                temp_mbr.particiones_mbr[indice].size_particion += size_bytes;
                                fseek(arch, 0, SEEK_SET);
                                fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                            }
                            else
                            {
                                std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                            }
                        }
                        else
                        {
                            std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                        }
                    }
                }
                else
                {

                    //Eliminar espacio sin eliminar la partición
                    if (size_bytes >= temp_mbr.particiones_mbr[indice].size_particion)
                    {
                        std::cout << "ERROR, el tamaño que desea borrar de " << name << " en " << path << " es mas grande que la particion" << endl;
                    }
                    else
                    {
                        temp_mbr.particiones_mbr[indice].size_particion -= size_bytes;
                        fseek(arch, 0, SEEK_SET);
                        fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                        std::cout << "Se redujo el espacio a la partición " << name << " en " << path << endl;
                    }
                }
            }
            else
            {

                //Ahora si particion extendida xd
                if (tipo == "add")
                {
                    if (indice != 3)
                    {
                        int aux1 = temp_mbr.particiones_mbr[indice].inicio_particion + temp_mbr.particiones_mbr[indice].size_particion;
                        int aux2 = temp_mbr.particiones_mbr[indice + 1].inicio_particion;

                        if (aux2 - aux1 != 0)
                        {
                            if ((aux2 - aux1) >= size_bytes)
                            {
                                temp_mbr.particiones_mbr[indice].size_particion += size_bytes;

                                fseek(arch, 0, SEEK_SET);
                                fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                            }
                            else
                            {
                                std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                            }
                        }
                        else
                        {
                            //Espacio libre
                            if (temp_mbr.particiones_mbr[indice + 1].estado_particion == '1')
                            {
                                if (temp_mbr.particiones_mbr[indice + 1].size_particion >= size_bytes)
                                {
                                    temp_mbr.particiones_mbr[indice].size_particion += size_bytes;
                                    temp_mbr.particiones_mbr[indice + 1].size_particion -= size_bytes;
                                    temp_mbr.particiones_mbr[indice + 1].inicio_particion += size_bytes;
                                    fseek(arch, 0, SEEK_SET);
                                    fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                    std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                                }
                                else
                                {
                                    std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                                }
                            }
                        }
                    }
                    else
                    {
                        //No hay espacio a la deracha xd
                        int sp = temp_mbr.particiones_mbr[indice].inicio_particion + temp_mbr.particiones_mbr[indice].size_particion;
                        int total = temp_mbr.tamanio_mbr + (int)sizeof(mbr); //verificamos el espacio total
                        if ((total - sp) != 0)
                        {
                            if ((total - sp) >= size_bytes)
                            {
                                temp_mbr.particiones_mbr[indice].size_particion += size_bytes;
                                fseek(arch, 0, SEEK_SET);
                                fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                                std::cout << "Se agrego espacio a la partición " << name << " en " << path << endl;
                            }
                            else
                            {
                                std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                            }
                        }
                        else
                        {
                            std::cout << "ERROR, no hay suficiente espacio para agregar a " << name << " en " << path << endl;
                        }
                    }
                }
                else
                {
                    //Ahora si a quitar espacio sin borrar la particion xd
                    if (size_bytes >= temp_mbr.particiones_mbr[indice_extend].size_particion)
                    {
                        std::cout << "ERROR, si se quita ese espacio de " << name << " en " << path << " se carga la particion" << endl;
                    }
                    else
                    {
                        ebr temp_ebr;
                        fseek(arch, temp_mbr.particiones_mbr[indice_extend].inicio_particion, SEEK_SET);
                        fread(&temp_ebr, sizeof(ebr), 1, arch);
                        while ((temp_ebr.sig_particion != -1) && (ftell(arch) < (temp_mbr.particiones_mbr[indice_extend].size_particion + temp_mbr.particiones_mbr[indice_extend].inicio_particion)))
                        {
                            fseek(arch, temp_ebr.sig_particion, SEEK_SET);
                            fread(&temp_ebr, sizeof(ebr), 1, arch);
                        }
                        int lastLogic = temp_ebr.inicio_particion + temp_ebr.tamanio_particion;
                        int temp = (temp_mbr.particiones_mbr[indice_extend].inicio_particion + temp_mbr.particiones_mbr[indice_extend].size_particion) - size_bytes;
                        //Validacion de no borrar ninguna logica
                        if (temp > lastLogic)
                        {
                            temp_mbr.particiones_mbr[indice_extend].size_particion -= size_bytes;
                            fseek(arch, 0, SEEK_SET);
                            fwrite(&temp_mbr, sizeof(mbr), 1, arch);
                            std::cout << "Se redujo el espacio a la partición " << name << " en " << path << endl;
                        }
                        else
                        {
                            std::cout << "ERROR, si borra " << name << " de " << path << " va a dañar una particion logica" << endl;
                        }
                    }
                }
            }
        }
        else
        {
            //PArticion Logica
            if (indice_extend != -1)
            {
                int num_logica = FindLogicPart(path, name);
                if (num_logica != -1)
                {
                    if (tipo == "add")
                    {
                        ebr temp_ebr;
                        fseek(arch, num_logica, SEEK_SET);
                        fread(&temp_ebr, sizeof(ebr), 1, arch);
                    }
                    else
                    {
                        ebr temp_ebr;
                        fseek(arch, num_logica, SEEK_SET);
                        fread(&temp_ebr, sizeof(ebr), 1, arch);
                        if (size_bytes <= temp_ebr.tamanio_particion)
                        {
                            temp_ebr.tamanio_particion -= size_bytes;
                            fseek(arch, num_logica, SEEK_SET);
                            fread(&temp_ebr, sizeof(ebr), 1, arch);
                            std::cout << "Se redujo la particion " << name << " en " << path << " con exito " << endl;
                        }
                        else
                        {
                            std::cout << "ERROR, no puede quitar espacio " << name << " en " << path << " porque se carga la logica " << endl;
                        }
                    }
                }
                else
                {
                    std::cout << "ERROR, no se encuentra particion logica " << name << " en " << path << " para redimensionar" << endl;
                }
            }
            else
            {
                std::cout << "ERROR al quitar espacio de " << name << " en " << path << endl;
            }
        }
    }
    else
    {
        std::cout << "ERROR, desmonte la partición " << name << " del disco " << path << " antes de eliminarla" << endl;
    }
    std::fclose(arch);
}
int disk::FindPrimExtPart(std::string path, std::string name)
{
    FILE *arch = fopen(path.c_str(), "rb+");

    mbr temp_mbr;
    fseek(arch, 0, SEEK_SET);
    fread(&temp_mbr, sizeof(mbr), 1, arch);
    for (int i = 0; i < 4; i++)
    {
        if (temp_mbr.particiones_mbr[i].estado_particion != '1')
        {
            if (temp_mbr.particiones_mbr[i].nombre_particion == name)
            {
                std::fclose(arch);
                return i;
            }
        }
    }
    std::fclose(arch);
    return -1;
}
void disk::mountPart(std::string path, std::string name)
{
    /*int letra = lista_mount->findLetter(path, name);
    int num = lista_mount->findNumber(path, name);*/

    if (existPart(path, name))
    {
        int num_part = FindPrimExtPart(path, name);
        if (num_part != -1)
        {
            FILE *arch = fopen(path.c_str(), "rb+");
            mbr temp_mbr;
            fseek(arch, 0, SEEK_SET);
            fread(&temp_mbr, sizeof(mbr), 1, arch);
            temp_mbr.particiones_mbr[num_part].estado_particion == '2';
            fseek(arch, 0, SEEK_SET);
            fread(&temp_mbr, sizeof(mbr), 1, arch);
            std::fclose(arch);
            int letra = lista_mount->findLetter(path, name);
            if (letra != -1)
            {
                int num = lista_mount->findNumber(path, name);
                char tem_letra = static_cast<char>(letra);
                string id = "70" + to_string(num) + tem_letra;
                nodoLista *nuevo = new nodoLista(path, name, tem_letra, num);
                lista_mount->insertNode(nuevo);
                std::cout << "Particion " << name << " montada con exito en " << path << " con Id: " << id << endl;
                //lista_mount->showList();
            }
            else
            {
                std::cout << "ERROR la particion " << name << " ya esta montada en " << path << std::endl;
            }
        }
        else
        {
            //Particion Logica
            int num_part = FindLogicPart(path, name);
            if (num_part != -1)
            {
                FILE *arch = fopen(path.c_str(), "rb+");
                ebr temp_ebr;
                fseek(arch, num_part, SEEK_SET);
                fread(&temp_ebr, sizeof(ebr), 1, arch);
                temp_ebr.estado_particion = '2';
                fseek(arch, num_part, SEEK_SET);
                fread(&temp_ebr, sizeof(ebr), 1, arch);
                std::fclose(arch);

                int letra = lista_mount->findLetter(path, name);

                if (letra != -1)
                {
                    int num = lista_mount->findNumber(path, name);
                    char tem_letra = static_cast<char>(letra);
                    string id = "70" + to_string(num) + tem_letra;
                    nodoLista *nuevo = new nodoLista(path, name, tem_letra, num);
                    lista_mount->insertNode(nuevo);
                    std::cout << "Particion " << name << " montada con exito en " << path << " con Id: " << id << endl;
                    //lista_mount->showList();
                }
                else
                {
                    std::cout << "ERROR la particion " << name << " ya esta montada en " << path << std::endl;
                }
            }
            else
            {
                std::cout << "ERROR, no se encuentra la particion logica " << name << " en " << path << " para montar" << endl;
            }
        }
    }
    else
    {
        std::cout << "ERROR, no se encontro las particion " << path << " en " << path << endl;
    }
}
void disk::unmount(std::string id)
{
    int temp_unmount = lista_mount->deleteNode(id);

    if (temp_unmount == 1)
    {
        std::cout << "Particion " << id << " desmontada correctamente" << endl;
    }
    else
    {
        std::cout << "ERROR, no se puede desmontar la particion " << id << endl;
    }
}
void disk::rep(std::string name, std::string path, std::string id_param, std::string ruta)
{

    nodoLista *aux = lista_mount->getNodo(id_param);

    if (aux != nullptr)
    {
        string dir = lista_mount->getPathDisk(id_param);
        if (dir != "null")
        {
            vector<string> div_path = dividirPath(path);
            if (verificarPath(div_path[0]) == false)
            {
                if (mkdir(div_path[0].c_str(), S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH == -1))
                {
                    std::cout << "Error al crear carpeta en: " << div_path[0] << endl;
                    return;
                }
            }
            if (name == "mbr")
            {
                ReporteMBR(dir, path, "jpg");
            }
            else if (name == "disk")
            {
                ReporteDisk(dir, path, "jpg");
            }
            else if (name == "sb")
            {
                int indice = FindPrimExtPart(aux->getPath(), aux->getName());
                if (indice != -1)
                {
                    mbr temp_mbr;
                    FILE *arch = fopen(aux->getPath().c_str(), "rb+");
                    fread(&temp_mbr, sizeof(mbr), 1, arch);
                    fseek(arch, temp_mbr.particiones_mbr[indice].inicio_particion, SEEK_SET);
                    fclose(arch);

                    ReporteSuper(dir, path, temp_mbr.particiones_mbr[indice].inicio_particion);
                }
            }
            else if (name == "inode")
            {
                int indice = FindPrimExtPart(aux->getPath(), aux->getName());
                if (indice != -1)
                {
                    mbr temp_mbr;
                    SuperBloque sp;
                    FILE *arch = fopen(aux->getPath().c_str(), "rb+");
                    fread(&temp_mbr, sizeof(mbr), 1, arch);
                    fseek(arch, temp_mbr.particiones_mbr[indice].inicio_particion, SEEK_SET);
                    fread(&sp, sizeof(SuperBloque), 1, arch);
                    fclose(arch);
                    ReporteInodo(aux->getPath(), path, sp.s_bm_inode_start, sp.s_inode_start, sp.s_bm_block_start);
                }
            }
            else if (name == "block")
            {
                int index = FindPrimExtPart(aux->getPath(), aux->getName());
                if (index != -1)
                {
                    mbr temp_mbr;
                    SuperBloque sp;
                    FILE *arch = fopen(aux->getPath().c_str(), "rb+");
                    fread(&temp_mbr, sizeof(mbr), 1, arch);
                    fseek(arch, temp_mbr.particiones_mbr[index].inicio_particion, SEEK_SET);
                    fread(&sp, sizeof(SuperBloque), 1, arch);

                    fclose(arch);
                    ReporteBlock(dir, path, sp.s_bm_block_start, sp.s_block_start, sp.s_inode_start);
                }
            }
            else if (name == "journaling")
            {
                int index = FindPrimExtPart(aux->getPath(), aux->getName());
                if (index != -1)
                { //Primaria|Extendida
                    mbr temp_mbr;
                    SuperBloque super;
                    FILE *fp = fopen(aux->getPath().c_str(), "rb+");
                    fread(&temp_mbr, sizeof(mbr), 1, fp);
                    fseek(fp, temp_mbr.particiones_mbr[index].inicio_particion, SEEK_SET);
                    fread(&super, sizeof(SuperBloque), 1, fp);
                    fclose(fp);
                    ReporteJournaling(dir, path, temp_mbr.particiones_mbr[index].size_particion);
                }
            }
            else if (name == "tree")
            {
                int index = FindPrimExtPart(aux->getPath(), aux->getName());
                if (index != -1)
                {
                    mbr masterboot;
                    FILE *fp = fopen(aux->getPath().c_str(), "rb+");
                    fread(&masterboot, sizeof(mbr), 1, fp);
                    fseek(fp, masterboot.particiones_mbr[index].inicio_particion, SEEK_SET);
                    fclose(fp);
                    ReporteArbol(aux->getPath(), path, masterboot.particiones_mbr[index].inicio_particion);
                }
            }
            else if (name == "bm_inode")
            {
                int index = FindPrimExtPart(aux->getPath(), aux->getName());
                if (index != -1)
                {
                    mbr masterboot;
                    SuperBloque super;
                    FILE *fp = fopen(aux->getPath().c_str(), "rb+");
                    fread(&masterboot, sizeof(mbr), 1, fp);
                    fseek(fp, masterboot.particiones_mbr[index].inicio_particion, SEEK_SET);
                    fread(&super, sizeof(SuperBloque), 1, fp);
                    fclose(fp);

                    ReporteBitmap(aux->getPath(),path, super.s_bm_inode_start, super.s_inodes_count); 

                }
            }
        }
    }
}
void disk::mkfs(std::string id, std::string type, int fs)
{
    nodoLista *aux = lista_mount->getNodo(id);
    if (aux != nullptr)
    {
        int indice = FindPrimExtPart(aux->getPath(), aux->getName());
        if (indice != 1)
        {
            mbr temp_mbr;
            FILE *arch = fopen(aux->getPath().c_str(), "rb+");
            fread(&temp_mbr, sizeof(mbr), 1, arch);
            if (fs == 2)
            {
                FormateoExt2(temp_mbr.particiones_mbr[indice].size_particion, temp_mbr.particiones_mbr[indice].inicio_particion, aux->getPath());
            }
            else
            {
                FormateoExt3(temp_mbr.particiones_mbr[indice].size_particion, temp_mbr.particiones_mbr[indice].inicio_particion, aux->getPath());
            }
        }
        else
        {
        }
    }
    else
    {
        std::cout << "ERROR, no se encuentra ninguna particion " << id << " montada" << endl;
    }
}
void disk::ComandoMKDIR(std::string path_carpeta, bool flag_p)
{
    if (flag_login != false)
    {
        int res = createCarpet(path_carpeta, flag_p);
        if (res == 0)
        {
            std::cout << "ERROR, Carpeta ya existente" << std::endl;
        }
        else if (res == 1)
        {
            if (sesion_activa.tipo_sistema == 3)
            {
                char aux[500];
                char operacion[8];
                char content[5];
                strcpy(aux, path_carpeta.c_str());
                strcpy(operacion, "mkdir");
                strcpy(content, "null");
                SaveOpJournal(operacion, 1, 664, aux, content);
            }
            cout << "Carpeta " << path_carpeta << " creada con exito" << endl;
        }
        else if (res == 2)
        {
            std::cout << "Error, no hay permisos de escrituras para poder crear " << path_carpeta << std::endl;
        }
    }
    else
    {
        std::cout << "ERROR, no hay un usuario logeado para poder crear carpeta " << path_carpeta << endl;
    }
}

void disk::SaveOpJournal(char *op, int type, int rights, char *name, char *content)
{
    SuperBloque sp;
    Journal registro;
    strcpy(registro.journal_op_type, op);
    registro.journal_type = type;
    strcpy(registro.journal_name, name);
    strcpy(registro.journal_content, content);
    registro.journal_date = time(nullptr);
    registro.journal_owner = sesion_activa.id_user;
    registro.journal_permissions = rights;
    FILE *fp = fopen(sesion_activa.direccion.c_str(), "rb+");
    //Buscar el ultimo journal
    Journal registroAux;
    bool ultimo = false;
    fseek(fp, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, fp);
    int inicio_journal = sesion_activa.inicioSuper + static_cast<int>(sizeof(SuperBloque));
    int final_journal = sp.s_bm_inode_start;
    fseek(fp, inicio_journal, SEEK_SET);
    while ((ftell(fp) < final_journal) && !ultimo)
    {
        fread(&registroAux, sizeof(Journal), 1, fp);
        if (registroAux.journal_type != 0 && registroAux.journal_type != 1)
            ultimo = true;
    }
    fseek(fp, ftell(fp) - static_cast<int>(sizeof(Journal)), SEEK_SET);
    fwrite(&registro, sizeof(Journal), 1, fp);
    fclose(fp);
}
void disk::FormateoExt2(int tamanio_particion, int inicio_particion, std::string path)
{
    double cant_inodos = (tamanio_particion - static_cast<int>(sizeof(SuperBloque))) / (4 + static_cast<int>(sizeof(TablaInodo)) + 3 * static_cast<int>(sizeof(BloqueArchivo)));
    int num_estructuras = static_cast<int>(floor(cant_inodos)); //Numero de inodos
    int num_bloques = 3 * num_estructuras;                      // numero de estructuras

    time_t now = time(0);
    tm *dt = localtime(&now);
    std::string mes = std::to_string(dt->tm_mon);
    std::string anio = std::to_string(1900 + dt->tm_year);
    std::string dia = std::to_string(dt->tm_mon);
    std::string hora = std::to_string(dt->tm_hour);
    std::string sec = std::to_string(dt->tm_min);
    std::string fecha = dia + "-" + mes + "-" + anio + " " + hora + ":" + sec;

    SuperBloque sp;
    sp.s_filesystem_type = 2;
    sp.s_inodes_count = num_estructuras;
    sp.s_blocks_count = num_bloques;
    sp.s_free_blocks_count = num_bloques - 2;
    sp.s_free_inodes_count = num_estructuras - 2;
    strcpy(sp.s_mtime, fecha.c_str());
    strcpy(sp.s_umtime, fecha.c_str());
    sp.s_mnt_count = 0;
    sp.s_magic = 0xEF53;
    sp.s_inode_size = sizeof(TablaInodo);
    sp.s_block_size = sizeof(BloqueArchivo);
    sp.s_first_ino = 2;
    sp.s_first_blo = 2;
    sp.s_bm_inode_start = inicio_particion + static_cast<int>(sizeof(SuperBloque));
    sp.s_bm_block_start = inicio_particion + static_cast<int>(sizeof(SuperBloque)) + num_estructuras;
    sp.s_inode_start = inicio_particion + static_cast<int>(sizeof(SuperBloque)) + num_estructuras + num_bloques;
    sp.s_block_start = inicio_particion + static_cast<int>(sizeof(SuperBloque)) + num_estructuras + num_bloques + (static_cast<int>(sizeof(TablaInodo)) * num_estructuras);

    TablaInodo tabla_inodo;
    BloqueCarpeta bloque_carpeta;

    char buffer = '0';
    char buffer_2 = '1';
    char buffer_3 = '2';

    FILE *arch = fopen(path.c_str(), "rb+");

    //superBloque -> Bitmap de Inodos -> Inodos -> Bloques
    fseek(arch, inicio_particion, SEEK_SET);
    fwrite(&sp, sizeof(SuperBloque), 1, arch);

    for (int i = 0; i < num_estructuras; i++)
    {
        fseek(arch, sp.s_bm_inode_start + i, SEEK_SET);
        fwrite(&buffer, sizeof(char), 1, arch);
    }

    fseek(arch, sp.s_bm_inode_start, SEEK_SET);
    fwrite(&buffer_2, sizeof(char), 1, arch);
    fwrite(&buffer_2, sizeof(char), 1, arch);

    for (int i = 0; i < num_bloques; i++)
    {
        fseek(arch, sp.s_bm_block_start + i, SEEK_SET);
        fwrite(&buffer, sizeof(char), 1, arch);
    }

    //bit para usuario.txt
    fseek(arch, sp.s_bm_block_start, SEEK_SET);
    fwrite(&buffer_2, sizeof(char), 1, arch);
    fwrite(&buffer_3, sizeof(char), 1, arch);

    tabla_inodo.i_uid = 1;
    tabla_inodo.i_gid = 1;
    tabla_inodo.i_size = 0;
    tabla_inodo.i_atime = time(nullptr);
    tabla_inodo.i_ctime = time(nullptr);
    tabla_inodo.i_mtime = time(nullptr);
    tabla_inodo.i_block[0] = 0;

    for (int i = 1; i < 15; i++)
    {
        tabla_inodo.i_block[i] = -1;
    }

    tabla_inodo.i_type = '0';
    tabla_inodo.i_perm = 664;
    fseek(arch, sp.s_inode_start, SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);

    //Creacion del bloque para la carpeta /
    strcpy(bloque_carpeta.b_content[0].b_name, "."); //Actual (el mismo)
    bloque_carpeta.b_content[0].b_inodo = 0;

    strcpy(bloque_carpeta.b_content[1].b_name, ".."); //Padre
    bloque_carpeta.b_content[1].b_inodo = 0;

    strcpy(bloque_carpeta.b_content[2].b_name, "users.txt");
    bloque_carpeta.b_content[2].b_inodo = 1;

    strcpy(bloque_carpeta.b_content[3].b_name, ".");
    bloque_carpeta.b_content[3].b_inodo = -1;
    fseek(arch, sp.s_block_start, SEEK_SET);
    fwrite(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);

    //inodo para users.txt
    tabla_inodo.i_uid = 1;
    tabla_inodo.i_gid = 1;
    tabla_inodo.i_size = 27;
    tabla_inodo.i_atime = time(nullptr);
    tabla_inodo.i_ctime = time(nullptr);
    tabla_inodo.i_mtime = time(nullptr);
    tabla_inodo.i_block[0] = 1;
    for (int i = 1; i < 15; i++)
    {
        tabla_inodo.i_block[i] = -1;
    }
    tabla_inodo.i_type = '1';
    tabla_inodo.i_perm = 755;
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);

    //Bloque para users.txt
    tabla_inodo.i_type = '1';
    tabla_inodo.i_perm = 755;
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(tabla_inodo)), SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);
    /*-------------Bloque para users.txt------------*/
    BloqueArchivo bloque_archivo;
    memset(bloque_archivo.b_content, 0, sizeof(bloque_archivo.b_content));
    strcpy(bloque_archivo.b_content, "1,G,root\n1,U,root,root,123\n");
    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)), SEEK_SET);
    fwrite(&bloque_archivo, sizeof(BloqueArchivo), 1, arch);

    std::cout << "Disco " << path << " formateado correctamente en EXT2" << endl;

    fclose(arch);
}
void disk::FormateoExt3(int tamanio_particion, int inicio_particon, std::string path)
{
    //n despejada de la ecuacion
    double n = (tamanio_particion - static_cast<int>(sizeof(SuperBloque))) / (4 + static_cast<int>(sizeof(TablaInodo)) + 3 * static_cast<int>(sizeof(BloqueArchivo)));
    int num_estructuras = static_cast<int>(floor(n));
    int num_bloques = 3 * num_estructuras;
    int super_tamanio = static_cast<int>(sizeof(SuperBloque));
    int journal_tamanio = static_cast<int>(sizeof(Journal)) * num_estructuras;

    time_t now = time(0);
    tm *dt = localtime(&now);
    std::string mes = std::to_string(dt->tm_mon);
    std::string anio = std::to_string(1900 + dt->tm_year);
    std::string dia = std::to_string(dt->tm_mon);
    std::string hora = std::to_string(dt->tm_hour);
    std::string sec = std::to_string(dt->tm_min);
    std::string fecha = dia + "-" + mes + "-" + anio + " " + hora + ":" + sec;

    SuperBloque sp;
    sp.s_filesystem_type = 3;
    sp.s_inodes_count = num_estructuras;
    sp.s_blocks_count = num_bloques;
    sp.s_free_blocks_count = num_bloques - 2;
    sp.s_free_inodes_count = num_estructuras - 2;
    strcpy(sp.s_mtime, fecha.c_str());
    strcpy(sp.s_umtime, fecha.c_str());
    sp.s_mnt_count = 0;
    sp.s_magic = 0xEF53;
    sp.s_inode_size = sizeof(TablaInodo);
    sp.s_block_size = sizeof(BloqueArchivo);
    sp.s_first_ino = 2;
    sp.s_first_blo = 2;
    sp.s_bm_inode_start = inicio_particon + static_cast<int>(sizeof(SuperBloque)) + journal_tamanio;
    sp.s_bm_block_start = inicio_particon + static_cast<int>(sizeof(SuperBloque)) + num_estructuras + journal_tamanio;
    sp.s_inode_start = inicio_particon + static_cast<int>(sizeof(SuperBloque)) + num_estructuras + num_bloques + journal_tamanio;
    sp.s_block_start = inicio_particon + static_cast<int>(sizeof(SuperBloque)) + journal_tamanio + num_estructuras + num_bloques + (static_cast<int>(sizeof(TablaInodo)) * num_estructuras);

    TablaInodo tabla_inodo;
    BloqueCarpeta bloque_carpeta;

    char buffer = '0';
    char buffer_2 = '1';
    char buffer_3 = '2';

    FILE *arch = fopen(path.c_str(), "rb+");

    //superBloque -> Bitmap de Inodos -> Inodos -> Bloques
    fseek(arch, inicio_particon, SEEK_SET);
    fwrite(&sp, sizeof(SuperBloque), 1, arch);

    for (int i = 0; i < num_estructuras; i++)
    {
        fseek(arch, sp.s_bm_inode_start + i, SEEK_SET);
        fwrite(&buffer, sizeof(char), 1, arch);
    }

    fseek(arch, sp.s_bm_inode_start, SEEK_SET);
    fwrite(&buffer_2, sizeof(char), 1, arch);
    fwrite(&buffer_2, sizeof(char), 1, arch);

    for (int i = 0; i < num_bloques; i++)
    {
        fseek(arch, sp.s_bm_block_start + i, SEEK_SET);
        fwrite(&buffer, sizeof(char), 1, arch);
    }

    //bit para root y usuario.txt
    fseek(arch, sp.s_bm_block_start, SEEK_SET);
    fwrite(&buffer_2, sizeof(char), 1, arch);
    fwrite(&buffer_3, sizeof(char), 1, arch);

    tabla_inodo.i_uid = 1;
    tabla_inodo.i_gid = 1;
    tabla_inodo.i_size = 0;
    tabla_inodo.i_atime = time(nullptr);
    tabla_inodo.i_ctime = time(nullptr);
    tabla_inodo.i_mtime = time(nullptr);
    tabla_inodo.i_block[0] = 0;

    for (int i = 1; i < 15; i++)
    {
        tabla_inodo.i_block[i] = -1;
    }

    tabla_inodo.i_type = '0';
    tabla_inodo.i_perm = 664;
    fseek(arch, sp.s_inode_start, SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);

    //Creacion del bloque para la carpeta root
    strcpy(bloque_carpeta.b_content[0].b_name, "."); //Actual
    bloque_carpeta.b_content[0].b_inodo = 0;

    strcpy(bloque_carpeta.b_content[1].b_name, ".."); //Padre
    bloque_carpeta.b_content[1].b_inodo = 0;

    strcpy(bloque_carpeta.b_content[2].b_name, "users.txt");
    bloque_carpeta.b_content[2].b_inodo = 1;

    strcpy(bloque_carpeta.b_content[3].b_name, ".");
    bloque_carpeta.b_content[3].b_inodo = -1;
    fseek(arch, sp.s_block_start, SEEK_SET);
    fwrite(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);

    //Creacion de inodo de users
    //inodo para users.txt
    tabla_inodo.i_uid = 1;
    tabla_inodo.i_gid = 1;
    tabla_inodo.i_size = 27;
    tabla_inodo.i_atime = time(nullptr);
    tabla_inodo.i_ctime = time(nullptr);
    tabla_inodo.i_mtime = time(nullptr);
    tabla_inodo.i_block[0] = 1;
    for (int i = 1; i < 15; i++)
    {
        tabla_inodo.i_block[i] = -1;
    }
    tabla_inodo.i_type = '1';
    tabla_inodo.i_perm = 755;
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);

    //Bloque para users.txt
    tabla_inodo.i_type = '1';
    tabla_inodo.i_perm = 755;
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(tabla_inodo)), SEEK_SET);
    fwrite(&tabla_inodo, sizeof(TablaInodo), 1, arch);
    /*-------------Bloque para users.txt------------*/
    BloqueArchivo bloque_archivo;
    memset(bloque_archivo.b_content, 0, sizeof(bloque_archivo.b_content));
    strcpy(bloque_archivo.b_content, "1,G,root\n1,U,root,root,123\n");
    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)), SEEK_SET);
    fwrite(&bloque_archivo, sizeof(BloqueArchivo), 1, arch);

    std::cout << "Disco " << path << " formateado correctamente en EXT3" << endl;

    fclose(arch);
}
void disk::ReporteMBR(std::string path_disco, std::string path_destino, std::string extension)
{
    FILE *arch = fopen(path_disco.c_str(), "r");
    FILE *arch_dot = fopen("grafo.dot", "w");

    mbr temp_mbr;
    fseek(arch, 0, SEEK_SET);
    fread(&temp_mbr, sizeof(mbr), 1, arch);

    int extendida = -1;

    time_t now = time(0);
    tm *dt = localtime(&now);
    std::string mes = std::to_string(dt->tm_mon);
    std::string anio = std::to_string(1900 + dt->tm_year);
    std::string dia = std::to_string(dt->tm_mon);
    std::string hora = std::to_string(dt->tm_hour);
    std::string sec = std::to_string(dt->tm_min);
    std::string fecha = dia + "-" + mes + "-" + anio + " " + hora + ":" + sec;

    string grafo_string = "";
    grafo_string += "digraph G{ \n subgraph cluster{\n label=\"MBR\"\ntbl[shape=box,label=<\n";
    grafo_string += "<table border=\'0\' cellborder=\'1\' cellspacing=\'0\' width=\'300\'  height=\'200\' >\n";
    grafo_string += "<tr>  <td width=\'150\'> <b>Nombre</b> </td> <td width=\'150\'> <b>Valor</b> </td>  </tr>\n";

    grafo_string += "<tr>  <td><b>mbr_tamaño</b></td><td>" + to_string(temp_mbr.tamanio_mbr) + "</td>  </tr>\n";
    grafo_string += "<tr>  <td><b>mbr_fecha_creacion</b></td> <td>" + fecha + "</td>  </tr>\n";
    grafo_string += "<tr>  <td><b>mbr_disk_signature</b></td> <td>" + to_string(temp_mbr.signature_mbr) + "</td>  </tr>\n";
    grafo_string += "<tr>  <td><b>Disk_fit</b></td> <td>" + string(1, temp_mbr.fit_disco) + "</td>  </tr>\n";

    for (int i = 0; i < 4; i++)
    {
        if (temp_mbr.particiones_mbr[i].inicio_particion != -1 && temp_mbr.particiones_mbr[i].estado_particion != '1')
        {
            if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
            {
                extendida = i;
            }

            grafo_string += "<tr>  <td><b>part_status_" + to_string(i + 1) + "</b></td> <td>" + string(1, temp_mbr.particiones_mbr[i].estado_particion) + "</td>  </tr>\n";
            grafo_string += "<tr>  <td><b>part_type_" + to_string(i + 1) + "</b></td> <td>" + string(1, temp_mbr.particiones_mbr[i].tipo_particion) + "</td>  </tr>\n";
            grafo_string += "<tr>  <td><b>part_fit_" + to_string(i + 1) + "</b></td> <td>" + string(1, temp_mbr.particiones_mbr[i].fit_particion) + "</td>  </tr>\n";
            grafo_string += "<tr>  <td><b>part_start_" + to_string(i + 1) + "</b></td> <td>" + to_string(temp_mbr.particiones_mbr[i].inicio_particion) + "</td>  </tr>\n";
            grafo_string += "<tr>  <td><b>part_size_" + to_string(i + 1) + "</b></td> <td>" + to_string(temp_mbr.particiones_mbr[i].size_particion) + "</td>  </tr>\n";
            grafo_string += "<tr>  <td><b>part_name_" + to_string(i + 1) + "</b></td> <td>" + temp_mbr.particiones_mbr[i].nombre_particion + "</td>  </tr>\n";
        }
    }

    grafo_string += "</table>\n>];\n}\n";

    if (extendida != -1)
    {
        int indx_ebr = 1;
        ebr temp_ebr;
        fseek(arch, temp_mbr.particiones_mbr[extendida].inicio_particion, SEEK_SET);
        while (fread(&temp_ebr, sizeof(ebr), 1, arch) != 0 && (ftell(arch) < (temp_mbr.particiones_mbr[extendida].inicio_particion + temp_mbr.particiones_mbr[extendida].size_particion)))
        {

            if (temp_ebr.estado_particion != '1')
            {

                grafo_string += "subgraph cluster_" + to_string(indx_ebr) + "{\n label=\"EBR_" + to_string(indx_ebr) + "\"\n";
                grafo_string += "\ntbl_" + to_string(indx_ebr) + "[shape=box, label=<\n ";
                grafo_string += "<table border=\'0\' cellborder=\'1\' cellspacing=\'0\'  width=\'300\' height=\'160\' >\n ";
                grafo_string += "<tr>  <td width=\'150\'><b>Nombre</b></td> <td width=\'150\'><b>Valor</b></td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_status_1</b></td> <td>" + string(1, temp_ebr.estado_particion) + "</td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_fit_1</b></td> <td>" + string(1, temp_ebr.fit_particion) + "</td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_start_1</b></td> <td>" + to_string(temp_ebr.inicio_particion) + "</td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_size_1</b></td> <td>" + to_string(temp_ebr.tamanio_particion) + "</td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_next_1</b></td> <td>" + to_string(temp_ebr.sig_particion) + "</td>  </tr>\n";
                grafo_string += "<tr>  <td><b>part_name_1</b></td> <td>" + std::string(temp_ebr.nombre_particion) + "</td>  </tr>\n";
                grafo_string += "</table>\n>];\n}\n";

                indx_ebr++;
            }
            if (temp_ebr.sig_particion == -1)
            {
                break;
            }
            else
            {
                fseek(arch, temp_ebr.sig_particion, SEEK_SET);
            }
        }
    }
    grafo_string += "}\n";
    std::fprintf(arch_dot, grafo_string.c_str());
    std::fclose(arch_dot);
    std::fclose(arch);
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
}
void disk::ReporteDisk(std::string path_disco, std::string path_destino, std::string extension)
{
    FILE *arch = fopen(path_disco.c_str(), "r");
    FILE *arch_dot = fopen("grafo.dot", "w");
    mbr temp_mbr;

    string grafo_string = "digraph G{\n\n";
    grafo_string += "tbl[\nshape=box\nlabel=<\n";
    grafo_string += "<table border=\'0\' cellborder=\'2\' width=\'600\' height=\"200\" color=\'LIGHTSTEELBLUE\'>\n";
    grafo_string += "<tr>\n";
    grafo_string += "<td height=\'200\' width=\'100\'> MBR </td>\n";

    fseek(arch, 0, SEEK_SET);
    fread(&temp_mbr, sizeof(mbr), 1, arch);
    int mbr_size = temp_mbr.tamanio_mbr;
    double espacio_usado = 0;
    for (int i = 0; i < 4; i++)
    {
        particion part_aux = temp_mbr.particiones_mbr[i];
        int size_part_i = part_aux.size_particion;
        if (part_aux.inicio_particion != -1)
        {
            double porc_real = (size_part_i * 100) / mbr_size;
            double porc_aux = (porc_real * 500) / 100;
            espacio_usado += porc_real;

            if (part_aux.estado_particion != '1')
            {
                if (part_aux.tipo_particion == 'p')
                {
                    grafo_string += "<td height=\'200\' width=\'" + to_string(porc_aux) + "\'>PRIMARIA <br/> Ocupado: " + to_string(porc_real) + "</td>\n";
                    if (i != 3)
                    {
                        int aux1 = part_aux.inicio_particion + part_aux.size_particion;
                        int aux2 = temp_mbr.particiones_mbr[i + 1].inicio_particion;

                        if (aux2 != -1)
                        {
                            if ((aux2 - aux1) != 0)
                            {
                                double a = ((aux2 - aux1) * 100) / mbr_size;
                                double b = (a * 500) / 100;
                                grafo_string += "     <td height=\'200\' width=\'" + to_string(b) + "\'>LIBRE<br/> Ocupado: " + to_string(a) + "</td>\n";
                            }
                        }
                    }
                    else
                    {
                        int p1 = part_aux.inicio_particion + part_aux.size_particion;
                        int mbr_s_aux = mbr_size + (int)sizeof(mbr);
                        if ((mbr_s_aux - p1) != 0)
                        {
                            //Espacio libre xdxd
                            double libre = (mbr_s_aux - p1) + sizeof(mbr);
                            double por_real = (libre * 100) / mbr_size;
                            double por_aux = (por_real * 500) / 100;
                            grafo_string += "     <td height=\'200\' width=\'" + to_string(por_aux) + "\'>LIBRE<br/> Ocupado: " + to_string(por_real) + "</td>\n";
                        }
                    }
                }
                else
                {
                    //Extended boot
                    ebr temp_ebr;
                    grafo_string += " <td  height=\'200\' width=\'" + to_string(porc_real) + "\'>\n<table border=\'0\'  height=\'200\' WIDTH=\'" + to_string(porc_real) + "\' cellborder=\'1\'>\n";
                    grafo_string += "<tr>  <td height=\'60\' colspan=\'15\'>EXTENDIDA</td>  </tr>\n     <tr>\n";
                    fseek(arch, part_aux.inicio_particion, SEEK_SET);
                    fread(&temp_ebr, sizeof(ebr), 1, arch);

                    if (temp_ebr.tamanio_particion != 0)
                    { //Si hay mas de alguna logica
                        fseek(arch, part_aux.inicio_particion, SEEK_SET);
                        while (fread(&temp_ebr, sizeof(ebr), 1, arch) != 0 && (ftell(arch) < (part_aux.inicio_particion + part_aux.size_particion)))
                        {
                            size_part_i = temp_ebr.tamanio_particion;
                            porc_real = (size_part_i * 100) / mbr_size;
                            if (porc_real != 0)
                            {
                                if (temp_ebr.estado_particion != '1')
                                {
                                    grafo_string += "<td height=\'140\'>EBR</td>\n";
                                    grafo_string += "     <td height=\'140\'>LOGICA<br/>Ocupado: " + to_string(porc_real) + "</td>\n";
                                }
                                else
                                { //Espacio no asignado
                                    grafo_string += "      <td height=\'150\'>LIBRE 1 <br/> Ocupado: " + to_string(porc_real) + "</td>\n";
                                }
                                if (temp_ebr.sig_particion == -1)
                                {
                                    size_part_i = (part_aux.inicio_particion + part_aux.size_particion) - (temp_ebr.inicio_particion + temp_ebr.tamanio_particion);
                                    porc_real = (size_part_i * 100) / mbr_size;
                                    if (porc_real != 0)
                                    {
                                        grafo_string += "     <td height=\'150\'>LIBRE 2<br/> Ocupado: " + to_string(porc_real) + " </td>\n";
                                    }
                                    break;
                                }
                                else
                                    fseek(arch, temp_ebr.sig_particion, SEEK_SET);
                            }
                        }
                    }
                    else
                    {
                        grafo_string += "     <td height=\'140\'> Ocupado " + to_string(porc_real) + "</td>";
                    }
                    grafo_string += "     </tr>\n     </table>\n     </td>\n";
                    if (i != 3)
                    {
                        int p1 = part_aux.inicio_particion + part_aux.size_particion;
                        int p2 = temp_mbr.particiones_mbr[i + 1].inicio_particion;
                        if (p2 != -1)
                        {
                            if ((p2 - p1) != 0)
                            { //Hay fragmentacion
                                int fragmentacion = p2 - p1;
                                double porcentaje_real = (fragmentacion * 100) / mbr_size;
                                double porcentaje_aux = (porcentaje_real * 500) / 100;
                                grafo_string += "     <td height=\'200\' width=\'" + to_string(porcentaje_aux) + "\'>LIBRE<br/> Ocupado: " + to_string(porcentaje_real) + "</td>\n";
                            }
                        }
                    }
                    else
                    {
                        int p1 = part_aux.inicio_particion + part_aux.size_particion;
                        int mbr_size_1 = mbr_size + (int)sizeof(mbr);
                        if ((mbr_size_1 - p1) != 0)
                        { //Libre
                            double libre = (mbr_size_1 - p1) + sizeof(mbr);
                            double porcentaje_real = (libre * 100) / mbr_size;
                            double porcentaje_aux = (porcentaje_real * 500) / 100;
                            grafo_string += "     <td height=\'200\' width=\'" + to_string(porcentaje_aux) + "\'>LIBRE<br/> Ocupado:" + to_string(porcentaje_real) + "</td>\n";
                        }
                    }
                }
            }
            else
            {
                //Graficamos espacio libre
                grafo_string += "     <td height=\'200\' width=\'" + to_string(porc_aux) + "\'>LIBRE <br/> Ocupado: " + to_string(porc_real) + "</td>\n";
            }
        }
    }
    grafo_string += "     </tr> \n     </table>        \n>];\n\n}";
    std::fprintf(arch_dot, grafo_string.c_str());
    std::fclose(arch_dot);
    std::fclose(arch);
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
}
void disk::ReporteSuper(std::string path_disco, std::string path_destino, int inicio_super)
{
    FILE *arch = fopen(path_disco.c_str(), "r");
    FILE *arch_dot = fopen("grafo.dot", "w");
    SuperBloque sp;
    fseek(arch, inicio_super, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);

    fprintf(arch_dot, "digraph G{\n");
    fprintf(arch_dot, "    nodo [shape=none, fontname=\"Century Gothic\" label=<");
    fprintf(arch_dot, "   <table border=\'0\' cellborder='1\' cellspacing=\'0\' bgcolor=\"cornflowerblue\">");
    fprintf(arch_dot, "    <tr> <td COLSPAN=\'2\'> <b>SUPERBLOQUE</b> </td></tr>\n");
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_file_system_type </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_filesystem_type);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_inodes_count </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_inodes_count);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_blocks_count </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_blocks_count);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_free_block_count </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_free_blocks_count);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_free_inodes_count </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_free_inodes_count);

    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_mtime </td> <td bgcolor=\"white\"> %s </td></tr>\n", "");
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_umtime </td> <td bgcolor=\"white\"> %s </td> </tr>\n", "");
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_mnt_count </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_mnt_count);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_magic </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_magic);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_inode_size </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_inode_size);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_block_size </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_block_size);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_first_ino </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_first_ino);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_first_blo </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_first_blo);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_bm_inode_start </td> <td bgcolor=\"white\"> %d </td></tr>\n", sp.s_bm_inode_start);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_bm_block_start </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_bm_block_start);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_inode_start </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_inode_start);
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> s_block_start </td> <td bgcolor=\"white\"> %d </td> </tr>\n", sp.s_block_start);
    fprintf(arch_dot, "   </table>>]\n");
    fprintf(arch_dot, "\n}");

    fclose(arch);
    fclose(arch_dot);
    string extension = "jpg";
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());

    cout << "Reporte de Super bloque de disco : " << path_disco << " Generado exitosamente" << endl;
}
void disk::ReporteInodo(std::string path_disco, std::string path_destino, int bm_inode_start, int inode_start, int bm_block_start)
{
    FILE *arch = fopen(path_disco.c_str(), "r");

    TablaInodo inodo_aux;
    int aux = bm_inode_start;
    int i = 0;
    char buffer;

    FILE *arch_dot = fopen("grafo.dot", "w");
    fprintf(arch_dot, "digraph G{\n\n");

    while (aux < bm_block_start)
    {
        fseek(arch, bm_inode_start + i, SEEK_SET);
        buffer = static_cast<char>(fgetc(arch));
        aux++;
        if (buffer == '1')
        {
            fseek(arch, inode_start + static_cast<int>(sizeof(TablaInodo)) * i, SEEK_SET);
            fread(&inodo_aux, sizeof(TablaInodo), 1, arch);
            fprintf(arch_dot, "    nodo_%d [ shape=none fontname=\"Century Gothic\" label=<\n", i);
            fprintf(arch_dot, "   <table border=\'0\' cellborder=\'1\' cellspacing=\'0\' bgcolor=\"royalblue\">");
            fprintf(arch_dot, "    <tr> <td colspan=\'2\'> <b>Inodo %d</b> </td></tr>\n", i);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_uid </td> <td bgcolor=\"white\"> %d </td>  </tr>\n", inodo_aux.i_uid);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_gid </td> <td bgcolor=\"white\"> %d </td>  </tr>\n", inodo_aux.i_gid);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_size </td> <td bgcolor=\"white\"> %d </td> </tr>\n", inodo_aux.i_size);
            struct tm *tm;
            char fecha[100];
            tm = localtime(&inodo_aux.i_atime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_atime </td> <td bgcolor=\"white\"> %s </td>  </tr>\n", fecha);
            tm = localtime(&inodo_aux.i_ctime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_ctime </td> <td bgcolor=\"white\"> %s </td>  </tr>\n", fecha);
            tm = localtime(&inodo_aux.i_mtime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_mtime </td> <td bgcolor=\"white\"> %s </td></tr>\n", fecha);
            for (int b = 0; b < 15; b++)
                fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_block_%d </td> <td bgcolor=\"white\"> %d </td> </tr>\n", b, inodo_aux.i_block[b]);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_type </td> <td bgcolor=\"white\"> %c </td> </tr>\n", inodo_aux.i_type);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"> i_perm </td> <td bgcolor=\"white\"> %d </td> </tr>\n", inodo_aux.i_perm);
            fprintf(arch_dot, "   </table>>]\n");
        }
        i++;
    }
    fprintf(arch_dot, "\n}");
    fclose(arch_dot);

    fclose(arch);

    string extension = "jpg";
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
    std::cout << "Reporte de Inodos del disco " << path_disco << " generado exitosamente" << endl;
}
void disk::ReporteBlock(std::string path_disco, std::string path_destino, int bm_block_start, int block_start, int inode_start)
{
    FILE *arch = fopen(path_disco.c_str(), "r");

    BloqueCarpeta bloque_carpeta;
    BloqueArchivo bloque_archivo;
    BloqueApuntador bloque_puntador;

    int bbs = bm_block_start;
    int i = 0;
    char buffer;

    FILE *arch_dot = fopen("grafo.dot", "w");
    fprintf(arch_dot, "digraph G{\n\n");

    while (bbs < inode_start)
    {
        fseek(arch, bm_block_start + i, SEEK_SET);
        buffer = static_cast<char>(fgetc(arch));
        bbs++;
        if (buffer == '1')
        {
            fseek(arch, block_start + static_cast<int>(sizeof(BloqueCarpeta)) * i, SEEK_SET);
            fread(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);
            fprintf(arch_dot, "    nodo_%d [ shape=none, fontname=\"Century Gothic\" label=< \n", i);
            fprintf(arch_dot, "   <table border=\'0\' cellborder='1' cellspacing='0' bgcolor=\"seagreen\">");
            fprintf(arch_dot, "    <tr> <td colspan=\'2\'> <b>Bloque Carpeta %d</b> </td></tr>\n", i);
            fprintf(arch_dot, "    <tr> <td bgcolor=\"mediumseagreen\"> b_name </td> <td bgcolor=\"mediumseagreen\"> b_inode </td></tr>\n");
            for (int c = 0; c < 4; c++)
                fprintf(arch_dot, "    <tr> <td bgcolor=\"white\"> %s </td> <td bgcolor=\"white\"> %d </td></tr>\n", bloque_carpeta.b_content[c].b_name, bloque_carpeta.b_content[c].b_inodo);
            fprintf(arch_dot, "   </table>>]\n\n");
        }
        else if (buffer == '2')
        {
            fseek(arch, block_start + static_cast<int>(sizeof(BloqueArchivo)) * i, SEEK_SET);
            fread(&bloque_archivo, sizeof(BloqueArchivo), 1, arch);
            fprintf(arch_dot, "    nodo_%d [ shape=none, label=< \n", i);
            fprintf(arch_dot, "   <table border=\'0\' cellborder='1' cellspacing='0' bgcolor=\"sandybrown\">");
            fprintf(arch_dot, "    <tr> <td colspan=\'2\'> <b>Bloque Archivo %d </b></td></tr>\n", i);
            fprintf(arch_dot, "    <tr> <td colspan=\'2\' bgcolor=\"white\"> %s </td></tr>\n", bloque_archivo.b_content);
            fprintf(arch_dot, "   </table>>]\n\n");
        }
        else if (buffer == '3')
        {
            fseek(arch, block_start + static_cast<int>(sizeof(BloqueApuntador)) * i, SEEK_SET);
            fread(&bloque_puntador, sizeof(BloqueApuntador), 1, arch);
            fseek(arch, block_start + static_cast<int>(sizeof(BloqueApuntador)) * i, SEEK_SET);
            fread(&bloque_puntador, sizeof(BloqueApuntador), 1, arch);
            fprintf(arch_dot, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", i);
            fprintf(arch_dot, "   <table border=\'0\' bgcolor=\"khaki\">\n");
            fprintf(arch_dot, "    <tr> <td> <b>Pointer block %d</b></td></tr>\n", i);
            for (int j = 0; j < 16; j++)
                fprintf(arch_dot, "    <tr> <td bgcolor=\"white\">%d</td> </tr>\n", bloque_puntador.b_pointer[j]);
            fprintf(arch_dot, "   </table>>]\n\n");
        }
        i++;
    }
    fprintf(arch_dot, "\n}");

    fclose(arch);
    fclose(arch_dot);

    string extension = "jpg";
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
    std::cout << "Reporte de Bloques del disco " << path_disco << " generado exitosamente" << endl;
}
void disk::ReporteJournaling(std::string path_disco, std::string path_destino, int inicio_sp)
{
    FILE *arch = fopen(path_disco.c_str(), "r");

    SuperBloque super;
    Journal journal;
    fseek(arch, inicio_sp, SEEK_SET);
    fread(&super, sizeof(SuperBloque), 1, arch);
    FILE *arch_dot = fopen("grafo.dot", "w");
    fprintf(arch_dot, "digraph G{\n");
    fprintf(arch_dot, "    nodo [shape=none, fontname=\"Century Gothic\" label=<\n");
    fprintf(arch_dot, "   <table border=\'0\' cellborder='1\' cellspacing=\'0\'>\n");
    fprintf(arch_dot, "    <tr> <td COLSPAN=\'7\' bgcolor=\"cornflowerblue\"> <b>JOURNALING</b> </td></tr>\n");
    fprintf(arch_dot, "    <tr> <td bgcolor=\"lightsteelblue\"><b>Operacion</b></td> <td bgcolor=\"lightsteelblue\"><b>Tipo</b></td><td bgcolor=\"lightsteelblue\"><b>Nombre</b></td><td bgcolor=\"lightsteelblue\"><b>Contenido</b></td>\n");
    fprintf(arch_dot, "    <td bgcolor=\"lightsteelblue\"><b>Propietario</b></td><td bgcolor=\"lightsteelblue\"><b>Permisos</b></td><td bgcolor=\"lightsteelblue\"><b>Fecha</b></td></tr>\n");

    fseek(arch, inicio_sp + static_cast<int>(sizeof(SuperBloque)), SEEK_SET);

    struct tm *tm;
    char fecha[100];
    tm = localtime(&journal.journal_date);
    strftime(fecha, 100, "%d/%m/%y %H:%S", tm);

    /*while (ftell(arch) < super.s_bm_inode_start)
    {
        fread(&journal, sizeof(Journal), 1, arch);
        if (journal.journal_type == 0 || journal.journal_type == 1)
        {

            fprintf(arch_dot, "<tr><td>%s</td><td>%d</td><td>%s</td><td>%s</td><td>%d</td><td>%d</td><td>%s</td></tr>\n", journal.journal_op_type, journal.journal_type, journal.journal_name, journal.journal_content, journal.journal_owner, journal.journal_permissions, fecha);
        }
    }
*/
    fprintf(arch_dot, "   </table>>]\n}");

    fclose(arch);
    fclose(arch_dot);

    string extension = "jpg";
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
    std::cout << "Reporte de Journaling del disco " << path_disco << " generado exitosamente" << endl;
}
void disk::ReporteBitmap (std::string path_disco, std::string path_destino, int start_bm, int n){
    FILE *fp = fopen(path_destino.c_str(),"rb+");

    char byte;
    FILE *report = fopen(path_destino.c_str(),"w+");
    fseek(report,0,SEEK_SET);
    int cont = 0;

    for (int i = 0; i < n; i++) {
        fseek(fp,start_bm + i,SEEK_SET);
        byte = static_cast<char>(fgetc(fp));
        if(byte == '0')
            fprintf(report,"0 ");
        else
            fprintf(report,"1 ");
        if(cont == 19){
            cont = 0;
            fprintf(report, "\n");
        }else
            cont++;
    }
    fclose(report);

    fclose(fp);
    cout << "Reporte de Bitmap generado con exito " << endl;
}
void disk::ReporteArbol(std::string path_disco, std::string path_destino, int inicio_super)
{
    FILE *fp = fopen(path_disco.c_str(), "r");
    SuperBloque super;
    TablaInodo inodo;
    BloqueCarpeta carpeta;
    BloqueArchivo archivo;
    BloqueApuntador apuntador;

    fseek(fp, inicio_super, SEEK_SET);
    fread(&super, sizeof(SuperBloque), 1, fp);

    int aux = super.s_bm_inode_start;
    int i = 0;

    char buffer;

    FILE *graph = fopen("grafo.dot", "w");
    fprintf(graph, "digraph G{\n\n");
    fprintf(graph, "    rankdir=\"LR\" \n");

    while (aux < super.s_bm_block_start)
    {
        fseek(fp, super.s_bm_inode_start + i, SEEK_SET);
        buffer = static_cast<char>(fgetc(fp));
        aux++;
        int port = 0;
        if (buffer == '1')
        {
            fseek(fp, super.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * i, SEEK_SET);
            fread(&inodo, sizeof(TablaInodo), 1, fp);
            fprintf(graph, "    inodo_%d [ shape=plaintext fontname=\"Century Gothic\" label=<\n", i);
            fprintf(graph, "   <table bgcolor=\"royalblue\" border=\'0\' >");
            fprintf(graph, "    <tr> <td colspan=\'2\'><b>Inode %d</b></td></tr>\n", i);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_uid </td> <td bgcolor=\"white\"> %d </td>  </tr>\n", inodo.i_uid);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_gid </td> <td bgcolor=\"white\"> %d </td>  </tr>\n", inodo.i_gid);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_size </td><td bgcolor=\"white\"> %d </td> </tr>\n", inodo.i_size);
            struct tm *tm;
            char fecha[100];
            tm = localtime(&inodo.i_atime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_atime </td> <td bgcolor=\"white\"> %s </td> </tr>\n", fecha);
            tm = localtime(&inodo.i_ctime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_ctime </td> <td bgcolor=\"white\"> %s </td> </tr>\n", fecha);
            tm = localtime(&inodo.i_mtime);
            strftime(fecha, 100, "%d/%m/%y %H:%S", tm);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_mtime </td> <td bgcolor=\"white\"> %s </td> </tr>\n", fecha);
            for (int b = 0; b < 15; b++)
            {
                fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_block_%d </td> <td bgcolor=\"white\" port=\"f%d\"> %d </td></tr>\n", port, b, inodo.i_block[b]);
                port++;
            }
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_type </td> <td bgcolor=\"white\"> %c </td>  </tr>\n", inodo.i_type);
            fprintf(graph, "    <tr> <td bgcolor=\"lightsteelblue\"> i_perm </td> <td bgcolor=\"white\"> %d </td>  </tr>\n", inodo.i_perm);
            fprintf(graph, "   </table>>]\n\n");
            //Creamos los bloques relacionados al inodo
            for (int j = 0; j < 15; j++)
            {
                port = 0;
                if (inodo.i_block[j] != -1)
                {
                    fseek(fp, super.s_bm_block_start + inodo.i_block[j], SEEK_SET);
                    buffer = static_cast<char>(fgetc(fp));
                    if (buffer == '1')
                    { //Bloque carpeta
                        fseek(fp, super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * inodo.i_block[j], SEEK_SET);
                        fread(&carpeta, sizeof(BloqueCarpeta), 1, fp);
                        fprintf(graph, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", inodo.i_block[j]);
                        fprintf(graph, "   <table bgcolor=\"seagreen\" border=\'0\'>\n");
                        fprintf(graph, "    <tr> <td colspan=\'2\'><b>Folder block %d</b></td></tr>\n", inodo.i_block[j]);
                        fprintf(graph, "    <tr> <td bgcolor=\"mediumseagreen\"> b_name </td> <td bgcolor=\"mediumseagreen\"> b_inode </td></tr>\n");
                        for (int c = 0; c < 4; c++)
                        {
                            fprintf(graph, "    <tr> <td bgcolor=\"white\" > %s </td> <td bgcolor=\"white\"  port=\"f%d\"> %d </td></tr>\n", carpeta.b_content[c].b_name, port, carpeta.b_content[c].b_inodo);
                            port++;
                        }
                        fprintf(graph, "   </table>>]\n\n");
                        //Relacion de bloques a inodos
                        for (int c = 0; c < 4; c++)
                        {
                            if (carpeta.b_content[c].b_inodo != -1)
                            {
                                if (strcmp(carpeta.b_content[c].b_name, ".") != 0 && strcmp(carpeta.b_content[c].b_name, "..") != 0)
                                    fprintf(graph, "    bloque_%d:f%d -> inodo_%d;\n", inodo.i_block[j], c, carpeta.b_content[c].b_inodo);
                            }
                        }
                    }
                    else if (buffer == '2')
                    { //Bloque archivo
                        fseek(fp, super.s_block_start + static_cast<int>(sizeof(BloqueArchivo)) * inodo.i_block[j], SEEK_SET);
                        fread(&archivo, sizeof(BloqueArchivo), 1, fp);
                        fprintf(graph, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", inodo.i_block[j]);
                        fprintf(graph, "   <table border=\'0\' bgcolor=\"sandybrown\">\n");
                        fprintf(graph, "    <tr> <td> <b>File block %d</b></td></tr>\n", inodo.i_block[j]);
                        fprintf(graph, "    <tr> <td bgcolor=\"white\"> %s </td></tr>\n", archivo.b_content);
                        fprintf(graph, "   </table>>]\n\n");
                    }
                    else if (buffer == '3')
                    { //Bloque apuntador
                        fseek(fp, super.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * inodo.i_block[j], SEEK_SET);
                        fread(&apuntador, sizeof(BloqueApuntador), 1, fp);
                        fprintf(graph, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", inodo.i_block[j]);
                        fprintf(graph, "   <table border=\'0\' bgcolor=\"khaki\">\n");
                        fprintf(graph, "    <tr> <td> <b>Pointer block %d</b></td></tr>\n", inodo.i_block[j]);
                        for (int a = 0; a < 16; a++)
                        {
                            fprintf(graph, "    <tr> <td bgcolor=\"white\" port=\"f%d\">%d</td> </tr>\n", port, apuntador.b_pointer[a]);
                            port++;
                        }
                        fprintf(graph, "   </table>>]\n\n");
                        //Bloques carpeta/archivos del bloque de apuntadores
                        for (int x = 0; x < 16; x++)
                        {
                            port = 0;
                            if (apuntador.b_pointer[x] != -1)
                            {
                                fseek(fp, super.s_bm_block_start + apuntador.b_pointer[x], SEEK_SET);
                                buffer = static_cast<char>(fgetc(fp));
                                if (buffer == '1')
                                {
                                    fseek(fp, super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * apuntador.b_pointer[x], SEEK_SET);
                                    fread(&carpeta, sizeof(BloqueCarpeta), 1, fp);
                                    fprintf(graph, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", apuntador.b_pointer[x]);
                                    fprintf(graph, "   <table border=\'0\' bgcolor=\"seagreen\" >\n");
                                    fprintf(graph, "    <tr> <td colspan=\'2\'> <b>Folder block %d</b> </td></tr>\n", apuntador.b_pointer[x]);
                                    fprintf(graph, "    <tr> <td bgcolor=\"mediumseagreen\"> b_name </td> <td bgcolor=\"mediumseagreen\"> b_inode </td></tr>\n");
                                    for (int c = 0; c < 4; c++)
                                    {
                                        fprintf(graph, "    <tr> <td bgcolor=\"white\"> %s </td> <td bgcolor=\"white\" port=\"f%d\"> %d </td></tr>\n", carpeta.b_content[c].b_name, port, carpeta.b_content[c].b_inodo);
                                        port++;
                                    }
                                    fprintf(graph, "   </table>>]\n\n");
                                    //Relacion de bloques a inodos
                                    for (int c = 0; c < 4; c++)
                                    {
                                        if (carpeta.b_content[c].b_inodo != -1)
                                        {
                                            if (strcmp(carpeta.b_content[c].b_name, ".") != 0 && strcmp(carpeta.b_content[c].b_name, "..") != 0)
                                                fprintf(graph, "    bloque_%d:f%d -> inodo_%d;\n", apuntador.b_pointer[x], c, carpeta.b_content[c].b_inodo);
                                        }
                                    }
                                }
                                else if (buffer == '2')
                                {
                                    fseek(fp, super.s_block_start + static_cast<int>(sizeof(BloqueArchivo)) * apuntador.b_pointer[x], SEEK_SET);
                                    fread(&archivo, sizeof(BloqueArchivo), 1, fp);
                                    fprintf(graph, "    bloque_%d [shape=plaintext fontname=\"Century Gothic\" label=< \n", apuntador.b_pointer[x]);
                                    fprintf(graph, "   <table border=\'0\' bgcolor=\"sandybrown\">\n");
                                    fprintf(graph, "    <tr> <td> <b>File block %d</b></td></tr>\n", apuntador.b_pointer[x]);
                                    fprintf(graph, "    <tr> <td bgcolor=\"white\"> %s </td></tr>\n", archivo.b_content);
                                    fprintf(graph, "   </table>>]\n\n");
                                }
                                else if (buffer == '3')
                                {
                                }
                            }
                        }

                        //Relacion de bloques apuntador a bloques archivos/carpetas
                        for (int b = 0; b < 16; b++)
                        {
                            if (apuntador.b_pointer[b] != -1)
                                fprintf(graph, "    bloque_%d:f%d -> bloque_%d;\n", inodo.i_block[j], b, apuntador.b_pointer[b]);
                        }
                    }
                    //Relacion de inodos a bloques
                    fprintf(graph, "    inodo_%d:f%d -> bloque_%d; \n", i, j, inodo.i_block[j]);
                }
            }
        }
        i++;
    }

    fprintf(graph, "\n\n}");

    fclose(fp);
    fclose(graph);

    string extension = "jpg";
    string comando = "dot -T" + extension + " grafo.dot -o " + path_destino;
    std::system(comando.c_str());
    std::cout << "Reporte Tree del disco " << path_disco << " generado exitosamente" << endl;
}
void disk::login(std::string id, std::string usuario, std::string password)
{
    nodoLista *aux = lista_mount->getNodo(id);
    if (aux != nullptr)
    {
        int index = FindPrimExtPart(aux->getPath(), aux->getName());
        if (index != -1)
        {
            mbr temp_mbr;
            SuperBloque sp;
            TablaInodo inodo;

            FILE *arch = fopen(aux->getPath().c_str(), "rb+");
            fread(&temp_mbr, sizeof(mbr), 1, arch);
            fseek(arch, temp_mbr.particiones_mbr[index].inicio_particion, SEEK_SET);
            fread(&sp, sizeof(SuperBloque), 1, arch);
            fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
            fread(&inodo, sizeof(TablaInodo), 1, arch);
            fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
            inodo.i_atime = time(nullptr);
            fwrite(&inodo, sizeof(TablaInodo), 1, arch);
            fclose(arch);

            sesion_activa.inicioSuper = temp_mbr.particiones_mbr[index].inicio_particion;
            sesion_activa.fit = temp_mbr.particiones_mbr[index].fit_particion;
            sesion_activa.inicioJournal = temp_mbr.particiones_mbr[index].inicio_particion + static_cast<int>(sizeof(SuperBloque));
            sesion_activa.tipo_sistema = sp.s_filesystem_type;

            if (ComprobarDatos(usuario, password, aux->getPath()) == 1)
            {
                flag_login = true;
                std::cout << "Sesion iniciada con usuario " << usuario << endl;
            }
            else if (ComprobarDatos(usuario, password, aux->getPath()) == 2)
            {
                std::cout << "ERROR, la constraseña no coincide con el usuario " << usuario << endl;
            }
            else if (ComprobarDatos(usuario, password, aux->getPath()) == 0)
            {
                std::cout << "ERROR, no se encuentra el usuario:  " << usuario << endl;
            }
        }
    }
    else
    {
        std::cout << "ERROR, no se encuentra una particion montada con el id" << id << endl;
    }
}
void disk::Logout()
{
    if (flag_login == false)
    {
        flag_login = false;
        sesion_activa.id_user = -1;
        sesion_activa.direccion = "";
        sesion_activa.inicioSuper = -1;
        std::cout << "Sesion finalizada" << endl;
    }
    else
    {
        std::cout << "ERROR, no hay sesion iniciada" << endl;
    }
}
int disk::ComprobarDatos(std::string usuario, std::string password, std::string path_disco)
{
    FILE *arch = fopen(path_disco.c_str(), "rb+");

    char cadena[400] = "\0";
    SuperBloque sp;
    TablaInodo inodo;

    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);
    //Se lee el archivo de usuarios (el inodo xd)
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
    fread(&inodo, sizeof(TablaInodo), 1, arch);

    for (int i = 0; i < 15; i++)
    {
        if (inodo.i_block[i] != -1)
        {
            BloqueArchivo bloque_archivo;
            fseek(arch, sp.s_block_start, SEEK_SET);
            for (int j = 0; j <= inodo.i_block[i]; j++)
            {
                fread(&bloque_archivo, sizeof(BloqueArchivo), 1, arch);
            }
            strcat(cadena, bloque_archivo.b_content);
        }
    }

    fclose(arch);

    char *final_string;
    char *tok = strtok_r(cadena, "\n", &final_string);
    while (tok != nullptr)
    {
        char id[2];
        char tipo[2];
        std::string group;
        char user_[12];
        char password_[12];
        char *end_token;
        char *token2 = strtok_r(tok, ",", &end_token);
        strcpy(id, token2);
        if (strcmp(id, "0") != 0)
        {
            //Si no es un usuario o grupo
            token2 = strtok_r(nullptr, ",", &end_token);
            strcpy(tipo, token2);
            if (strcmp(tipo, "U") == 0)
            {
                token2 = strtok_r(nullptr, ",", &end_token);
                group = token2;
                token2 = strtok_r(nullptr, ",", &end_token);
                strcpy(user_, token2);
                token2 = strtok_r(nullptr, ",", &end_token);
                strcpy(password_, token2);
                if (strcmp(user_, usuario.c_str()) == 0)
                {
                    if (strcmp(password_, password.c_str()) == 0)
                    {
                        sesion_activa.direccion = path_disco;
                        sesion_activa.id_user = atoi(id);
                        sesion_activa.id_grp = FindGroup(group);
                        return 1;
                    }
                    else
                        return 2;
                }
            }
        }
        tok = strtok_r(nullptr, "\n", &final_string);
    }
    return 0;
}
vector<string> dividirPath(string entrada)
{
    vector<string> res;
    string path;
    std::string delimiter = "/";
    size_t pos = 0;
    std::string token;

    while ((pos = entrada.find(delimiter)) != std::string::npos)
    {
        token = entrada.substr(0, pos);
        path += token + "/";
        entrada.erase(0, pos + delimiter.length());
    }

    res.push_back(path);
    res.push_back(entrada);

    return res;
}
int disk::FindGroup(std::string name)
{
    FILE *arch = fopen(sesion_activa.direccion.c_str(), "rb+");

    char cadena[400] = "\0";
    SuperBloque sp;
    TablaInodo inodo;

    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);

    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)), SEEK_SET);
    fread(&inodo, sizeof(TablaInodo), 1, arch);

    for (int i = 0; i < 15; i++)
    {
        if (inodo.i_block[i] != -1)
        {
            BloqueArchivo bloque_archivo;
            fseek(arch, sp.s_block_start, SEEK_SET);
            for (int j = 0; j <= inodo.i_block[i]; j++)
            {
                fread(&bloque_archivo, sizeof(BloqueArchivo), 1, arch);
            }
            strcat(cadena, bloque_archivo.b_content);
        }
    }

    fclose(arch);

    char *final_string;
    char *token = strtok_r(cadena, "\n", &final_string);
    while (token != nullptr)
    {
        char id[2];
        char tipo[2];
        char group[12];
        char *fin_token;
        char *token2 = strtok_r(token, ",", &fin_token);
        strcpy(id, token2);
        if (strcmp(id, "0") != 0)
        {
            //Si no es un grupo o usuario eliminado
            token2 = strtok_r(nullptr, ",", &fin_token);
            strcpy(tipo, token2);
            if (strcmp(tipo, "G") == 0)
            {
                strcpy(group, fin_token);
                if (strcmp(group, name.c_str()) == 0)
                    return atoi(id);
            }
        }
        token = strtok_r(nullptr, "\n", &fin_token);
    }

    return -1;
}
int disk::FindLogicPart(std::string path, std::string name)
{
    FILE *arch = fopen(path.c_str(), "rb+");
    int temp_ext = -1;
    mbr temp_mbr;
    fseek(arch, 0, SEEK_SET);
    fwrite(&temp_mbr, sizeof(mbr), 1, arch);

    for (int i = 0; i < 4; i++)
    {
        if (temp_mbr.particiones_mbr[i].tipo_particion == 'e')
        {
            temp_ext = i;
            break;
        }
    }
    if (temp_ext != -1)
    {
        ebr temp_ebr;
        fseek(arch, temp_mbr.particiones_mbr[temp_ext].inicio_particion, SEEK_SET);

        while (fread(&temp_ebr, sizeof(ebr), 1, arch) != 0 && (ftell(arch) < (temp_mbr.particiones_mbr[temp_ext].inicio_particion + temp_mbr.particiones_mbr[temp_ext].size_particion)))
        {
            if (temp_ebr.nombre_particion == name)
            {
                int ret = (ftell(arch) - sizeof(ebr));
                std::fclose(arch);
                return ret;
            }
        }
    }
    std::fclose(arch);
    return -1;
}
int disk::inicioInodoBloque(FILE *arch, int num_inodo, char type)
{
    SuperBloque sp;
    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);
    if (type == '1')
    {
        return (sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * num_inodo);
    }
    else if (type == '2')
        return (sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * num_inodo);
    return 0;
}
TablaInodo disk::InodoCreate(int size_i, char type_i, int perms_i)
{
    TablaInodo inodo;
    inodo.i_uid = sesion_activa.id_user;
    inodo.i_gid = sesion_activa.id_grp;
    inodo.i_size = size_i;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    for (int i = 0; i < 15; i++)
        inodo.i_block[i] = -1;
    inodo.i_type = type_i;
    inodo.i_perm = perms_i;
    return inodo;
}
int disk::MakeNewDIr(FILE *arch, char fit, bool flag_p, char *path, int num_inodo)
{
    SuperBloque sp;
    TablaInodo inodo, inodo_nuevo;
    BloqueCarpeta carpeta, carpeta_nuevo, carpeta_aux;
    BloqueApuntador apuntador;

    std::vector<string> lista;
    char copia_path[500];
    char dir[500];
    char nombre_carpeta[80];

    strcpy(copia_path, path);
    strcpy(dir, dirname(copia_path));
    strcpy(copia_path, path);

    strcpy(nombre_carpeta, basename(copia_path));
    char *token = strtok(path, "/");
    int cont = 0;
    int numInodo = num_inodo;
    int response = 0;

    while (token != nullptr)
    {
        cont++;
        lista.push_back(token);
        token = strtok(nullptr, "/");
    }

    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);

    if (cont == 1)
    {
        //No es un directorio anidado
        int contenido = 0;
        int bloque = 0;
        int puntero = 0;
        int i_libre = LookFreePointerInode(arch, num_inodo, inodo, carpeta, apuntador, contenido, bloque, puntero);

        if (i_libre == 1)
        {
            if (bloque == 12)
            {
                //Apuntador indirecto simple
                bool permissions = WriteRights(inodo.i_perm, (inodo.i_uid == sesion_activa.id_user), (inodo.i_gid == sesion_activa.id_grp));
                if (permissions || (sesion_activa.id_user == 1 && sesion_activa.id_grp == 1))
                {
                    char buffer = '1';
                    int bitInodo = FindFreeBit(arch, 'I', fit);
                    //Agregamos la carpeta al espacio libre en el bloque
                    carpeta.b_content[contenido].b_inodo = bitInodo;
                    strcpy(carpeta.b_content[contenido].b_name, nombre_carpeta);
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * apuntador.b_pointer[puntero], SEEK_SET);
                    fwrite(&carpeta, sizeof(BloqueCarpeta), 1, arch);
                    //Creamos el nuevo inodo carpeta
                    inodo_nuevo = InodoCreate(0, '0', 664);
                    int bitBloque = FindFreeBit(arch, 'B', fit);
                    inodo_nuevo.i_block[0] = bitBloque;
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * bitInodo, SEEK_SET);
                    fwrite(&inodo_nuevo, sizeof(TablaInodo), 1, arch);
                    //Guardamos el bit del inodo en el bitmap
                    fseek(arch, sp.s_bm_inode_start + bitInodo, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Creamos el nuevo bloque carpeta

                    BloqueCarpeta carpeta_nueva;

                    for (int i = 0; i < 4; i++)
                    {
                        strcpy(carpeta_nueva.b_content[i].b_name, "");
                        carpeta_nueva.b_content[i].b_inodo = -1;
                    }

                    carpeta_nueva.b_content[0].b_inodo = bitInodo;
                    carpeta_nueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpeta_nueva.b_content[0].b_name, ".");
                    strcpy(carpeta_nueva.b_content[1].b_name, "..");
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                    fwrite(&carpeta_nueva, sizeof(BloqueCarpeta), 1, arch);
                    //Guardamos el bit del bloque en el bitmap
                    fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Sobreescribimos el super bloque
                    sp.s_free_inodes_count--;
                    sp.s_free_blocks_count--;
                    sp.s_first_ino++;
                    sp.s_first_blo++;
                    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
                    fwrite(&sp, sizeof(SuperBloque), 1, arch);
                    return 1;
                }
                else
                    return 2; // No hay permisos
            }
            else if (bloque == 13 || bloque == 14)
            {
                //Apuntador indirecto doble 13 ó indirecto triple 14
            }
            else
            {
                //Apuntadores directos

                bool permissions = WriteRights(inodo.i_perm, (inodo.i_uid == sesion_activa.id_user), (inodo.i_gid == sesion_activa.id_grp));
                if (permissions || (sesion_activa.id_user == 1 && sesion_activa.id_grp == 1))
                {
                    char buffer = '1';
                    int bitInodo = FindFreeBit(arch, 'I', fit);
                    //Agregamos la carpeta al espacio libre en el bloque
                    carpeta.b_content[contenido].b_inodo = bitInodo;
                    strcpy(carpeta.b_content[contenido].b_name, nombre_carpeta);
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * apuntador.b_pointer[puntero], SEEK_SET);
                    fwrite(&carpeta, sizeof(BloqueCarpeta), 1, arch);
                    //Creamos el nuevo inodo carpeta
                    inodo_nuevo = InodoCreate(0, '0', 664);
                    int bitBloque = FindFreeBit(arch, 'B', fit);
                    inodo_nuevo.i_block[0] = bitBloque;
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * bitInodo, SEEK_SET);
                    fwrite(&inodo_nuevo, sizeof(TablaInodo), 1, arch);
                    //Guardamos el bit del inodo en el bitmap
                    fseek(arch, sp.s_bm_inode_start + bitInodo, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Creamos el nuevo bloque carpeta

                    BloqueCarpeta carpeta_nueva;

                    for (int i = 0; i < 4; i++)
                    {
                        strcpy(carpeta_nueva.b_content[i].b_name, "");
                        carpeta_nueva.b_content[i].b_inodo = -1;
                    }

                    carpeta_nueva.b_content[0].b_inodo = bitInodo;
                    carpeta_nueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpeta_nueva.b_content[0].b_name, ".");
                    strcpy(carpeta_nueva.b_content[1].b_name, "..");
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                    fwrite(&carpeta_nueva, sizeof(BloqueCarpeta), 1, arch);
                    //Guardamos el bit del bloque en el bitmap
                    fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);

                    fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Sobreescribimos el super bloque
                    sp.s_free_inodes_count--;
                    sp.s_free_blocks_count--;
                    sp.s_first_ino++;
                    sp.s_first_blo++;
                    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
                    fwrite(&sp, sizeof(SuperBloque), 1, arch);
                    return 1;
                }
                else
                {
                    return 2;
                }
            }
        }
        else if (i_libre == 0)
        {
            //Bloques llenos
            bool flag_ft = false; //Primera vez
            puntero = -1;
            fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * numInodo, SEEK_SET);
            fread(&inodo, sizeof(TablaInodo), 1, arch);
            for (int i = 0; i < 15; i++)
            {
                if (i == 12)
                { //Apuntador indirecto simple
                    if (inodo.i_block[i] == -1)
                    {
                        //Primera vez
                        bloque = 12;
                        flag_ft = true;
                        break;
                    }
                    else
                    {
                        fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * inodo.i_block[12], SEEK_SET);
                        fread(&apuntador, sizeof(BloqueApuntador), 1, arch);
                        for (int j = 0; j < 16; j++)
                        {
                            if (apuntador.b_pointer[j] == -1)
                            {
                                bloque = 12;
                                puntero = j;
                                break;
                            }
                        }
                    }
                    if (flag_ft || puntero != -1)
                        break;
                }
                else if (i == 13)
                {
                }
                else if (i == 14)
                {
                }
                else
                {
                    if (inodo.i_block[i] == -1)
                    {
                        bloque = i;
                        break;
                    }
                }
            }
            if (bloque == 12 && flag_ft)
            {
                bool permissions = WriteRights(inodo.i_perm, (inodo.i_uid == sesion_activa.id_user), (inodo.i_gid == sesion_activa.id_grp));
                if (permissions || (sesion_activa.id_user == 1 && sesion_activa.id_grp == 1))
                {
                    char buffer = '1';
                    char buffer_3 = '3';
                    int bit_block = FindFreeBit(arch, 'B', fit);
                    inodo.i_block[bloque] = bit_block;
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * numInodo, SEEK_SET);
                    fwrite(&inodo, sizeof(TablaInodo), 1, arch);
                    // bit al bloque apuntador en el bitmap
                    fseek(arch, sp.s_bm_block_start + bit_block, SEEK_SET);
                    fwrite(&buffer_3, sizeof(char), 1, arch);

                    //solo se debe crear un bloque de apuntadores
                    int bitBloqueCarpeta = FindFreeBit(arch, 'B', fit); //Carpeta
                    apuntador.b_pointer[0] = bitBloqueCarpeta;
                    for (int i = 1; i < 16; i++)
                        apuntador.b_pointer[i] = -1;
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * bit_block, SEEK_SET);
                    fwrite(&apuntador, sizeof(BloqueApuntador), 1, arch);

                    //Creamos la carpeta del apuntador
                    int bit_inodo = FindFreeBit(arch, 'I', fit);
                    carpeta_aux = createDirBlock();
                    carpeta_aux.b_content[0].b_inodo = bit_inodo;
                    strcpy(carpeta_aux.b_content[0].b_name, nombre_carpeta);
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloqueCarpeta, SEEK_SET);
                    fwrite(&carpeta_aux, sizeof(BloqueCarpeta), 1, arch);

                    //Escribimos el bit del bloque carpeta en el bitmap
                    fseek(arch, sp.s_bm_block_start + bitBloqueCarpeta, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Creamos el nuevo inodo carpeta
                    inodo_nuevo = InodoCreate(0, '0', 664);
                    bitBloqueCarpeta = FindFreeBit(arch, 'B', fit); //Carpeta
                    inodo_nuevo.i_block[0] = bitBloqueCarpeta;
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * bit_inodo, SEEK_SET);
                    fwrite(&inodo_nuevo, sizeof(TablaInodo), 1, arch);
                    //Escribimos el bit del inodo en el bitmap
                    fseek(arch, sp.s_bm_inode_start + bit_inodo, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);

                    //Creamos el nuevo bloque carpeta
                    carpeta_nuevo = createDirBlock();
                    carpeta_nuevo.b_content[0].b_inodo = bit_inodo;
                    carpeta_nuevo.b_content[1].b_inodo = numInodo;
                    strcpy(carpeta_nuevo.b_content[0].b_name, ".");
                    strcpy(carpeta_nuevo.b_content[1].b_name, "..");
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloqueCarpeta, SEEK_SET);
                    fwrite(&carpeta_nuevo, sizeof(BloqueCarpeta), 1, arch);
                    //Guardamos el bit en el bitmap de bloques
                    fseek(arch, sp.s_bm_block_start + bitBloqueCarpeta, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Sobreescribimos el super bloque
                    sp.s_free_inodes_count--;
                    sp.s_free_blocks_count -= -3;
                    sp.s_first_ino++;
                    sp.s_first_blo += 3;
                    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
                    fwrite(&sp, sizeof(SuperBloque), 1, arch);

                    return 1;
                }
                else
                {
                    return 2;
                }
            }
            else if (bloque == 12 && !flag_ft)
            {
                char buffer = '1';
                //Escribir el numero de bloque en el bloque de apuntadores
                int bitBloque = FindFreeBit(arch, 'B', fit);
                apuntador.b_pointer[puntero] = bitBloque;
                fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * inodo.i_block[12], SEEK_SET);
                fwrite(&apuntador, sizeof(BloqueApuntador), 1, arch);
                //Creamos el bloque auxiliar
                int bitInodo = FindFreeBit(arch, 'I', fit);
                carpeta_aux = createDirBlock();
                carpeta_aux.b_content[0].b_inodo = bitInodo;
                strcpy(carpeta_aux.b_content[0].b_name, nombre_carpeta);
                fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                fwrite(&carpeta_aux, sizeof(BloqueCarpeta), 1, arch);
                //Escribimos el bit del bloque carpeta en el bitmap
                fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                fwrite(&buffer, sizeof(char), 1, arch);
                //Creamos el nuevo inodo carpeta
                inodo_nuevo = InodoCreate(0, '0', 664);
                inodo_nuevo.i_uid = sesion_activa.id_user;
                inodo_nuevo.i_gid = sesion_activa.id_grp;
                bitBloque = FindFreeBit(arch, 'B', fit); //Carpeta
                inodo_nuevo.i_block[0] = bitBloque;
                fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * bitInodo, SEEK_SET);
                fwrite(&inodo_nuevo, sizeof(TablaInodo), 1, arch);
                //Escribimos el bit del inodo en el bitmap
                fseek(arch, sp.s_bm_inode_start + bitInodo, SEEK_SET);
                fwrite(&buffer, sizeof(char), 1, arch);
                //Creamos el nuevo bloque carpeta
                carpeta_nuevo = createDirBlock();
                carpeta_nuevo.b_content[0].b_inodo = bitInodo;
                carpeta_nuevo.b_content[1].b_inodo = numInodo;
                strcpy(carpeta_nuevo.b_content[0].b_name, ".");
                strcpy(carpeta_nuevo.b_content[1].b_name, "..");
                fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                fwrite(&carpeta_nuevo, sizeof(BloqueCarpeta), 1, arch);
                //Guardamos el bit en el bitmap de bloques
                fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                fwrite(&buffer, sizeof(char), 1, arch);
                //Sobreescribimos el sp.bloque
                sp.s_free_inodes_count = sp.s_free_inodes_count - 1;
                sp.s_free_blocks_count = sp.s_free_blocks_count - 2;
                sp.s_first_ino = sp.s_first_ino + 1;
                sp.s_first_blo = sp.s_first_blo + 2;
                fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
                fwrite(&sp, sizeof(SuperBloque), 1, arch);
                return 1;
            }
            else if (bloque == 13 || bloque == 14)
            {
                //Apuntador indirecto doble y triple
            }
            else
            {
                //apuntadores directos

                bool permissions = WriteRights(inodo.i_perm, (inodo.i_uid == sesion_activa.id_user), (inodo.i_gid == sesion_activa.id_grp));
                if (permissions || (sesion_activa.id_user == 1 && sesion_activa.id_grp == 1))
                {
                    char buffer = '1';
                    int bitBloque = FindFreeBit(arch, 'B', fit);
                    inodo.i_block[bloque] = bitBloque;
                    //Sobreescribimos el inodo
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * numInodo, SEEK_SET);
                    fwrite(&inodo, sizeof(TablaInodo), 1, arch);
                    //Bloque carpeta auxiliar
                    int bitInodo = FindFreeBit(arch, 'I', fit);
                    carpeta_aux = createDirBlock();
                    carpeta_aux.b_content[0].b_inodo = bitInodo;
                    strcpy(carpeta_aux.b_content[0].b_name, nombre_carpeta);
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                    fwrite(&carpeta_aux, sizeof(BloqueCarpeta), 1, arch);
                    //Escribimos el bit en el bitmap de blqoues
                    fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Creamos el nuevo inodo
                    inodo_nuevo = InodoCreate(0, '0', 664);
                    bitBloque = FindFreeBit(arch, 'B', fit);
                    inodo_nuevo.i_block[0] = bitBloque;
                    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * bitInodo, SEEK_SET);
                    fwrite(&inodo_nuevo, sizeof(TablaInodo), 1, arch);
                    //Escribimos el bit en el bitmap de inodos
                    fseek(arch, sp.s_bm_inode_start + bitInodo, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Creamos el nuevo bloque carpeta
                    carpeta_nuevo = createDirBlock();
                    carpeta_nuevo.b_content[0].b_inodo = bitInodo;
                    carpeta_nuevo.b_content[1].b_inodo = numInodo;
                    strcpy(carpeta_nuevo.b_content[0].b_name, ".");
                    strcpy(carpeta_nuevo.b_content[1].b_name, "..");
                    fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * bitBloque, SEEK_SET);
                    fwrite(&carpeta_nuevo, sizeof(BloqueCarpeta), 1, arch);
                    //Guardamos el bit en el bitmap de bloques
                    fseek(arch, sp.s_bm_block_start + bitBloque, SEEK_SET);
                    fwrite(&buffer, sizeof(char), 1, arch);
                    //Sobreescribimos el super bloque
                    sp.s_free_inodes_count--;
                    sp.s_free_blocks_count -= 2;
                    sp.s_first_ino++;
                    sp.s_first_blo += 2;
                    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
                    fwrite(&sp, sizeof(SuperBloque), 1, arch);
                    return 1;
                }
                else
                    return 2;
            }
        }
    }
    else
    {
        //Directorio anidado
        int flag_existe = existFileorDir(arch, dir);
        if (flag_existe == -1)
        {
            if (flag_p)
            {
                int index = 0;
                string aux = "";
                //Crear posibles carpetas inexistentes
                for (int i = 0; i < cont; i++)
                {
                    aux += "/" + lista.at(i);
                    char dir[500];
                    char auxDir[500];
                    strcpy(dir, aux.c_str());
                    strcpy(auxDir, aux.c_str());
                    int carpeta = existFileorDir(arch, dir);
                    if (carpeta == -1)
                    {
                        response = MakeNewDIr(arch, fit, false, auxDir, index);
                        if (response == 2)
                            break;
                        strcpy(auxDir, aux.c_str());
                        index = existFileorDir(arch, auxDir);
                    }
                    else
                        index = carpeta;
                }
            }
            else
                return 3;
        }
        else
        {
            char dir[100] = "/";
            strcat(dir, nombre_carpeta);
            return MakeNewDIr(arch, fit, false, dir, flag_existe);
        }
    }
}
int disk::LookFreePointerInode(FILE *arch, int numero_inodo, TablaInodo &tablaInodo, BloqueCarpeta &bloque_carpeta, BloqueApuntador &bloque_apuntador, int &contenido, int &bloque, int &puntero)
{
    int apuntado_libre = 0;
    SuperBloque sp;
    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);
    fseek(arch, sp.s_inode_start + static_cast<int>(sizeof(TablaInodo)) * numero_inodo, SEEK_SET);
    fread(&tablaInodo, sizeof(TablaInodo), 1, arch);

    for (int i = 0; i < 15; i++)
    {
        if (tablaInodo.i_block[i] != -1)
        {

            //Apuntador indirecto simple
            if (i == 12)
            {
                fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * tablaInodo.i_block[i], SEEK_SET);
                fread(&bloque_apuntador, sizeof(BloqueApuntador), 1, arch);
                for (int j = 0; j < 16; j++)
                {
                    if (bloque_apuntador.b_pointer[j] != -1)
                    {
                        fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueApuntador)) * bloque_apuntador.b_pointer[j], SEEK_SET);
                        fread(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);
                        for (int k = 0; k < 4; k++)
                        {
                            if (bloque_carpeta.b_content[k].b_inodo == -1)
                            {
                                apuntado_libre = 1;
                                bloque = i;
                                puntero = j;
                                contenido = k;
                                break;
                            }
                        }
                    }
                    if (apuntado_libre == 1)
                        break;
                }
            }
            else if (i == 12 || i == 13)
            {
                //Indirecto doble y triple
            }
            else
            {
                //Apuntadore directos
                fseek(arch, sp.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)) * tablaInodo.i_block[i], SEEK_SET);
                fread(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);
                for (int j = 0; j < 4; j++)
                {
                    if (bloque_carpeta.b_content[j].b_inodo == -1)
                    {
                        apuntado_libre = 1;
                        bloque = i;
                        contenido = j;
                        break;
                    }
                }
            }
        }
        if (apuntado_libre == 1)
        {
            break;
        }
    }
    return apuntado_libre;
}
int disk::existFileorDir(FILE *arch, char *path)
{
    SuperBloque sp;
    TablaInodo tabla_inodo;
    BloqueCarpeta bloque_carpeta;
    BloqueApuntador bloque_apuntador;

    std::vector<string> lista;
    char *t = strtok(path, "/");
    int contador = 0;
    int num_inodo = 0;

    while (t != nullptr)
    {
        lista.push_back(t);
        contador++;
        t = strtok(nullptr, "/");
    }

    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);
    num_inodo = sp.s_inode_start;

    for (int cont2 = 0; cont2 < contador; cont2++)
    {
        fseek(arch, num_inodo, SEEK_SET);
        fread(&tabla_inodo, sizeof(TablaInodo), 1, arch);
        int sig = 0;
        for (int i = 0; i < 15; i++)
        {
            if (tabla_inodo.i_block[i] != -1)
            {
                //Apuntadores dir
                int byteBloque = inicioInodoBloque(arch, tabla_inodo.i_block[i], '2');
                fseek(arch, byteBloque, SEEK_SET);
                if (i < 12)
                {
                    fread(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);
                    for (int j = 0; j < 4; j++)
                    {
                        if ((cont2 == contador - 1) && (strcasecmp(bloque_carpeta.b_content[j].b_name, lista.at(cont2).c_str()) == 0))
                        {
                            //Se encontro la carpeta
                            return bloque_carpeta.b_content[j].b_inodo;
                        }
                        else if ((cont2 != contador - 1) && (strcasecmp(bloque_carpeta.b_content[j].b_name, lista.at(cont2).c_str()) == 0))
                        {
                            num_inodo = inicioInodoBloque(arch, bloque_carpeta.b_content[j].b_inodo, '1');
                            sig = 1;
                            break;
                        }
                    }
                }
                else if (i == 12)
                { //Apuntador indirecto
                    fread(&bloque_apuntador, sizeof(BloqueApuntador), 1, arch);
                    for (int j = 0; j < 16; j++)
                    {
                        if (bloque_apuntador.b_pointer[j] != -1)
                        {
                            byteBloque = inicioInodoBloque(arch, bloque_apuntador.b_pointer[j], '2');
                            fseek(arch, byteBloque, SEEK_SET);
                            fread(&bloque_carpeta, sizeof(BloqueCarpeta), 1, arch);
                            for (int k = 0; k < 4; k++)
                            {
                                if ((k == contador - 1) && (strcasecmp(bloque_carpeta.b_content[k].b_name, lista.at(k).c_str()) == 0))
                                {
                                    //Retorna el numero de nodo o bloque
                                    return bloque_carpeta.b_content[k].b_inodo;
                                }
                                else if ((k != contador - 1) && (strcasecmp(bloque_carpeta.b_content[k].b_name, lista.at(k).c_str()) == 0))
                                {
                                    num_inodo = inicioInodoBloque(arch, bloque_carpeta.b_content[k].b_inodo, '1');
                                    sig = 1;
                                    break;
                                }
                            }
                            if (sig == 1)
                                break;
                        }
                    }
                }
                else if (i == 13)
                {
                }
                else if (i == 14)
                {
                }
                if (sig == 1)
                    break;
            }
        }
    }
    return -1;
}
bool disk::WriteRights(int pass, bool flag_user, bool flag_group)
{
    string temp = to_string(pass);
    char owner = temp[0];
    char group = temp[1];
    char other = temp[2];
    if ((owner == '2' || owner == '3' || owner == '6' || owner == '7') && flag_user)
        return true;
    else if ((group == '2' || group == '3' || group == '6' || group == '7') && flag_group)
        return true;
    else if (other == '2' || other == '3' || other == '6' || other == '7')
        return true;

    return false;
}
BloqueCarpeta disk::createDirBlock()
{
    BloqueCarpeta aux;

    for (int i = 0; i < 4; i++)
    {
        strcpy(aux.b_content[i].b_name, "");
        aux.b_content[i].b_inodo = -1;
    }

    return aux;
}
int disk::FindFreeBit(FILE *arch, char type, char fit)
{
    SuperBloque sp;
    int inicio_bitmap = 0;
    char temp_bit = '0';
    int bit_libre = -1;
    int tamanio_bitmap = 0;

    fseek(arch, sesion_activa.inicioSuper, SEEK_SET);
    fread(&sp, sizeof(SuperBloque), 1, arch);

    if (type == 'I')
    {
        tamanio_bitmap = sp.s_inodes_count;
        inicio_bitmap = sp.s_bm_inode_start;
    }
    else if (type == 'B')
    {
        tamanio_bitmap = sp.s_blocks_count;
        inicio_bitmap = sp.s_bm_block_start;
    }

    if (fit == 'f')
    { //Primer ajuste
        for (int i = 0; i < tamanio_bitmap; i++)
        {
            fseek(arch, inicio_bitmap + i, SEEK_SET);
            temp_bit = static_cast<char>(fgetc(arch));
            if (temp_bit == '0')
            {
                bit_libre = i;
                return bit_libre;
            }
        }

        if (bit_libre == -1)
            return -1;
    }
    else if (fit == 'b')
    {
        int libres = 0;
        int auxLibres = -1;

        for (int i = 0; i < tamanio_bitmap; i++)
        { //Primer recorrido
            fseek(arch, inicio_bitmap + i, SEEK_SET);
            temp_bit = static_cast<char>(fgetc(arch));
            if (temp_bit == '0')
            {
                libres++;
                if (i + 1 == tamanio_bitmap)
                {
                    if (auxLibres == -1 || auxLibres == 0)
                        auxLibres = libres;
                    else
                    {
                        if (auxLibres > libres)
                            auxLibres = libres;
                    }
                    libres = 0;
                }
            }
            else if (temp_bit == '1')
            {
                if (auxLibres == -1 || auxLibres == 0)
                    auxLibres = libres;
                else
                {
                    if (auxLibres > libres)
                        auxLibres = libres;
                }
                libres = 0;
            }
        }

        for (int i = 0; i < tamanio_bitmap; i++)
        {
            fseek(arch, inicio_bitmap + i, SEEK_SET);
            temp_bit = static_cast<char>(fgetc(arch));
            if (temp_bit == '0')
            {
                libres++;
                if (i + 1 == tamanio_bitmap)
                    return ((i + 1) - libres);
            }
            else if (temp_bit == '1')
            {
                if (auxLibres == libres)
                    return ((i + 1) - libres);
                libres = 0;
            }
        }

        return -1;
    }
}
std::string currentDateTime()
{
    time_t now = time(0);
    struct tm tstruct;
    char buf[80];
    tstruct = *localtime(&now);
    strftime(buf, sizeof(buf), "%Y-%m-%d.%X", &tstruct);
    return buf;
}