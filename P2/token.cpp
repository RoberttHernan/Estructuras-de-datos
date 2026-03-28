#include "token.h"

using namespace N; 


 token::token (string tipo_token, string valor){
     this->tipo_token = tipo_token;
     this->valor = valor; 
}

string token::getTipoToken (){
    return this->tipo_token;
}

string token::getValor (){
    return this->valor;
}

void token::setTipoToken (string tipoToken){
    this->tipo_token = tipoToken; 
}

void token::setValor (string valor){
    this->valor = valor; 
}
