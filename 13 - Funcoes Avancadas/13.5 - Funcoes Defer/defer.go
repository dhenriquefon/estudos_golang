package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoEstaAprovado(nota1 int, nota2 int) bool {
	// sempre no final essa funcao sera chamada
	// imagina que aqui gostariamos que em qualquer sitauacao de erro queremos fechar a cconexao no banco, defer eh o ideal
	defer fmt.Println("Media calculada, resultado sera retornado")

	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	}

	return false
}

func main() {
	funcao1()
	funcao2()

	// qunando usamos a clausa DEFER estamos adiando a execucao da func1
	// o defer eh executado no momento antes do RETURN da funcao
	// imagina
	defer funcao1()
	funcao2()

	alunoEstaAprovado(10, 10)

	alunoEstaAprovado(4, 4)
}
