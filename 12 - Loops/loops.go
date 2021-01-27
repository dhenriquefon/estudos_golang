package main

import (
	"fmt"
)

func main() {
	//i := 0

	// interacao como a do while
	// for i < 3 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// interacao convencional do for
	// for j := 0; j < 3; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// interacao de arrasys
	nomes := [3]string{"Joao", "Mariane", "Gabriel"}

	// obrigatoriamente o primeiro parametro eh o indice, se nao quiser usar o indice usar _
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// percorrendo uma string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	// percorrendo map
	usuario := map[string]string{
		"user":  "dhenrique",
		"paswd": "123",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// percorrendo em struct nao da pra usar o range
	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Douglas", "Fonseca"}

	// for chave, valor := range usuario2 {

	// }
}
