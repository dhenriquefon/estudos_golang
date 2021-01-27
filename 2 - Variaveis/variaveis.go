package main

import (
	"fmt"
)

func main () {
	var variavel string  = "var 1"
	variavel2 := "var2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste1"
		variavel4 string = "teste2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel5, variavel6)

	const constant1 string = "constante"

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}