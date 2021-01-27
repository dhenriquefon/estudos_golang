package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	var resultado int8 = somar(10, 5)
	fmt.Println(resultado)

	var f = func(txt string) string {
		fmt.Println("funcao f")
		return "retorno f"
	}

	var result_f string = f("lalala")
	fmt.Println(result_f)

	resultado_soma, resultado_subtracao := calculosMatematicos(30, 10)

	fmt.Println(resultado_soma, resultado_subtracao)

	resultado_soma_2, _ := calculosMatematicos(30, 10)

	fmt.Println(resultado_soma_2)
}
