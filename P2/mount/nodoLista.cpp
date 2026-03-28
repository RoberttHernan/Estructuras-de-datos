#include "nodoLista.h"

nodoLista::nodoLista (std::string path, std::string name, char letra, int num){
    this->path = path; 
    this->name = name; 
    this->letra = letra; 
    this->numero = num; 
    this->siguiente = nullptr; 
}


std::string nodoLista::getPath()
{
    return path;
}
std::string nodoLista::getName()
{
    return name;
}
char nodoLista::getLetra()
{
    return letra;
}
int nodoLista ::getNumero()
{
    return numero;
}

nodoLista *nodoLista::getSiguiente()
{
    return siguiente;
}

void nodoLista::setPath(std::string path)
{
    this->path = path;
}
void nodoLista::setName(std::string name)
{
    this->name = name;
}
void nodoLista::setLetra(char letra)
{
    this->letra = letra;
}
void nodoLista::setNumero(int numero)
{
    this->numero = numero;
}
void nodoLista::setSiguiente(nodoLista *siguiente)
{
    this->siguiente = siguiente;
}