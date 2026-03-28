#pragma once

#include <iostream>
class nodoLista
{
private:
    std::string path;
    std::string name;
    char letra;
    int numero;
    nodoLista *siguiente;

public:
    nodoLista(std::string, std::string, char letra, int num);
    std::string getPath();
    std::string getName();
    char getLetra();
    int getNumero();
    nodoLista *getSiguiente();

    void setPath(std::string);
    void setName(std::string);
    void setLetra(char);
    void setNumero(int);
    void setSiguiente(nodoLista *);
};

