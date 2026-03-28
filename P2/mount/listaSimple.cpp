#include "listaSimple.h"

listaSimple::listaSimple()
{
    head = nullptr;
}
int listaSimple::findNumber(std::string path, std::string nombre)
{
    int num = 1;
    nodoLista *aux = this->head;
    while (aux != NULL)
    {
        if ((path == aux->getPath()) && nombre == aux->getName())
        {
            return -1;
        }
        else
        {
            if (aux->getPath() == path)
            {
                return aux->getNumero();
            }
            else if (num <= aux->getNumero())
            {
                num++;
            }
        }
        aux = aux->getSiguiente();
    }
    return num;
}
nodoLista * listaSimple::getNodo (std:: string id){
    nodoLista * aux = head; 
    while (aux != nullptr)
    {
        std::string temp_id = "70" + std::to_string(aux->getNumero()) + aux->getLetra();
        if (id == temp_id){
            return aux; 

        }
        aux = aux->getSiguiente(); 
    }
    return nullptr; 
    
}
int listaSimple::findLetter(std::string path, std::string nombre)
{
    nodoLista *aux = head;
    int letra = 'a';
    while (aux != NULL)
    {
        if (aux->getPath() == path && aux->getLetra() == letra)
        {
            letra++;
        }
        aux = aux->getSiguiente();
    }
    return letra;
}
void listaSimple::showList()
{
    std::cout << "---------------------------------" << std::endl;
    std::cout << "|       Particiones montadas    |" << std::endl;
    std::cout << "---------------------------------" << std::endl;
    std::cout << "|      Nombre    |    ID        |" << std::endl;
    std::cout << "---------------------------------" << std::endl;

    nodoLista *aux = head;
    while (aux != nullptr)
    {
        std::cout << "|     " << aux->getName() << "          "
                  << "70" << aux->getNumero() << aux->getLetra() << std::endl;
        std::cout << "---------------------------------" << std::endl;
        aux = aux->getSiguiente();
    }
}
void listaSimple::insertNode(nodoLista *nodo_nuevo)
{
    nodoLista *aux = head;
    if (head == nullptr)
    {
        head = nodo_nuevo;
    }
    else
    {
        while (aux->getSiguiente() != nullptr)
        {
            aux = aux->getSiguiente();
        }
        aux->setSiguiente(nodo_nuevo);
    }
}
int listaSimple::deleteNode(std::string id)
{
    nodoLista *aux = head;
    std::string temp_id = "70" + std::to_string(aux->getNumero()) + std::to_string(aux->getLetra());

    if (id == temp_id)
    {
        head = aux->getSiguiente();
        free(aux);
        return 1;
    }
    else
    {
        nodoLista *aux2 = nullptr;
        while (aux != nullptr)
        {
            temp_id = "70" + std::to_string(aux->getNumero()) + aux->getLetra();
            if (id == temp_id)
            {
                aux2->setSiguiente(aux->getSiguiente());
                return 1;
            }
            aux2 = aux;
            aux = aux->getSiguiente();
        }
    }
    return 0;
}
bool listaSimple::findNode(std::string path, std::string name)
{
    nodoLista *aux = this->head;
    while (aux != nullptr)
    {
        if (aux->getPath() == path && aux->getName() == name)
        {
            return true;
        }
        aux = aux->getSiguiente();
    }
    return false;
}
std::string listaSimple::getPathDisk(std::string id_param)
{
    nodoLista *aux = head;
    while (aux != nullptr)
    {
        std::string temp_id = "70" + std::to_string(aux->getNumero()) + aux->getLetra();
        //std::cout <<"temp: " <<temp_id << " id_param: " <<id_param<<std::endl;
        if (id_param == temp_id)
        {
            return aux->getPath();
        }
        aux = aux->getSiguiente();
    }
    return "null";
}