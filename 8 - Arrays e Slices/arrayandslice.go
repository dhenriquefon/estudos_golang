package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "teste1"
	array1[1] = "teste2"
	array1[2] = "teste3"
	array1[3] = "teste4"
	array1[4] = "teste5"

	fmt.Println(array1)

	array2 := [5]string{"lalala", "lelele", "lililili", "lololo", "lululu"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	//fmt.Println(reflect.typeOf(array2))
	//fmt.Println(reflect.typeOf(slice))
	slice = append(slice, 17)
	fmt.Println(slice)

	// funciona da mesma forma que ponteiro
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "alterado"
	fmt.Println(slice2)

	// array interno
	// tio, tamanho, tamanho maximo/capacidade
	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 20)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
}
