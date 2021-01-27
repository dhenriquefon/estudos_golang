package main

import "fmt"

func fibonacci(posicao uint) uint {
	// condicao de parada
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	fmt.Println(fibonacci(posicao))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
