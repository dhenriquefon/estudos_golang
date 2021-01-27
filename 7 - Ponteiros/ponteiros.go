package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// PONTEIRO EH UMA REFERENCIA DE MEMORIA
	var variavel3 int = 100
	var ponteirovar3 *int = &variavel3

	fmt.Println(variavel3, ponteirovar3)
	fmt.Println(variavel3, *ponteirovar3)

	variavel3 = 150
	fmt.Println(variavel3, ponteirovar3)
	fmt.Println(variavel3, *ponteirovar3)
}
