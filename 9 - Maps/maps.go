package main

import "fmt"

func main() {
	fmt.Println("maps")

	// dentro do colchetes, tipo da chave
	// fora dos colchetes, tipo dos valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Gomes",
	}

	fmt.Println("Usario-nome: ", usuario["nome"])
	fmt.Println("Usario-sobrenome: ", usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"ultimo":   "Sanefuji",
		},
		"curso": {
			"nome":   "Biologia",
			"campus": "UFSCAR",
		},
	}

	fmt.Println(usuario2)

	// adicionando uma chave
	usuario2["signo"] = map[string]string{
		"signo": "Aries",
	}

	fmt.Print(usuario2["signo"])

	// deletando uma chave
	delete(usuario2, "nome")
}
