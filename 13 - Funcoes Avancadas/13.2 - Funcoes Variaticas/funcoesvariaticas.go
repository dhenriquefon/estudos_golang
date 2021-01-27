package main

import "fmt"

// funcao que recebe n parametros
func soma(numeros ...int) int {
	total := 0

	// numeros nesse caso se torna um slice, entao posso fazer um for para percorrer
	fmt.Println(numeros)

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(10, 10, 15, 10))
	escrever("douglas", 10, 10, 15, 10)
}
