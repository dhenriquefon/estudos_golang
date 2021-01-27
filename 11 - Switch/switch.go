package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terça feira"
	case 4:
		return "Quarta feira"
	case 5:
		return "Quinta feira"
	case 6:
		return "Sexta feira"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}

}

func diaDaSemana2(numero int) string {
	var dia string

	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough // joga o codigo obrigatoriamente para proxima condicao
	case numero == 2:
		dia = "Segunda Feira"
	case numero == 3:
		dia = "Terça feira"
	case numero == 4:
		dia = "Quarta feira"
	case numero == 5:
		dia = "Quinta feira"
	case numero == 6:
		dia = "Sexta feira"
	case numero == 7:
		dia = "Sabado"
	default:
		dia = "Numero invalido"
	}

	return dia
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)
}
