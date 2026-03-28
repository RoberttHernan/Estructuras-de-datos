#include <iostream>
#include "analizador.h"
#include "ejecucion.h"
using namespace std;
using namespace N;
using namespace ejc;

// Divide una cadena en tokens usando el delimitador indicado
std::vector<std::string> separarCadena(std::string entrada, std::string delimitador)
{
  std::vector<std::string> listaRetorno;
  size_t pos = 0;
  std::string token;
  while ((pos = entrada.find(delimitador)) != std::string::npos)
  {
    token = entrada.substr(0, pos);
    listaRetorno.push_back(token);
    entrada.erase(0, pos + delimitador.length());
  }
  listaRetorno.push_back(entrada);
  return listaRetorno;
}

// Si el comando es "exec -path=<ruta>", retorna la ruta del archivo de script;
// de lo contrario retorna cadena vacia
std::string isValidaExec(std::string comando)
{
  std::vector<std::string> lista1 = separarCadena(comando, " ");
  if (lista1[0] == "exec")
  {
    std::vector<std::string> lista2 = separarCadena(lista1[1], "=");
    return lista2[1];
  }
  return "";
}

void imprimirEncabezado()
{
  std::cout << "|---------------Sistema De Archivos-----------|" << endl;
}

int main(int argc, char const *argv[])
{
  std::string comando = "";
  std::string entrada;

  // Bucle interactivo: acepta un comando por iteracion.
  // Si el comando es "exec -path=<ruta>", lee el script desde el archivo indicado
  // y lo procesa como una sola cadena de entrada.
  while (true)
  {
    imprimirEncabezado();
    std::cout << "Ingrese Comando" << endl;
    getline(cin, comando);

    if (isValidaExec(comando) != "")
    {
      // Leer el script desde archivo
      std::string com = isValidaExec(comando);
      ifstream fe(com);
      string linea;
      while (getline(fe, linea))
      {
        entrada += linea + "\n";
      }
      fe.close();
    }
    else
    {
      entrada = comando;
    }

    analizador scanner;
    ejecucion parser;
    list<token> tokens = scanner.analizarCadena(entrada);

    if (tokens.size() > 0)
    {
      parser.ejecutarComandos(tokens);
    }
    else
    {
      std::cout << "ERROR en analisis Lexico" << endl;
    }

    std::cout << " " << endl;
    std::string opcion;
    std::cout << "Desea ingresar otro Comando? (S/N)" << endl;
    getline(cin, opcion);

    if (opcion != "S" && opcion != "s")
    {
      break;
    }
  }
}
