package main

import "fmt"

func main() {
	fmt.Println("Estruturas e Controle")

	numero := 10

	if numero > 10 {
		fmt.Println("valor maior que 10")
	} else {
		fmt.Println("valor menor que 10")
	}

	// inicializando valor no if
	// ';' separa a condicao do if
	if outronumero := numero; outronumero > 10 {
		fmt.Println("maior que 10")
	} else if outronumero > 1 {
		fmt.Println("maior que 1")
	} else {
		fmt.Println("valor menor que 1")
	}

	// a variavel fica restrita ao escopo do if
	//fmr.Println(outronumero) -> esse comando nao funciona

}
