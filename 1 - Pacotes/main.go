package main

import (	
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"	
)

func main() {
	fmt.Println("Escrevendo do main");
	auxiliar.Escrever();

	//validar email
	erro := checkmail.ValidateFormat("devbook@gmail.com");
	fmt.Println(erro);
}