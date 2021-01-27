package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end   endereco
}

type endereco struct {
	logradouro string
	cep        string
}

func main() {
	var u usuario
	fmt.Println(u)

	u.nome = "Lenin"
	u.idade = 52

	fmt.Println(u)

	localEndereco := endereco{"rua lidia coelho", "02035030"}
	usuario2 := usuario{"Marx", 33, localEndereco}

	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Engels"}

	fmt.Println(usuario3)
}
