package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 0:
		return "domingo"
	case 1:
		return "segunda"
	case 2:
		return "terça"
	default:
		return "Não é um dia da semana"
	}
}

func main() {
	fmt.Println("Estrutudas de controle condicionals")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}
	// cria a variavel outroNumero, atribui o valor de numero e compara se é maior que ZERO
	// variavel local ao escopo - não existe fora do if-init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	}

	fmt.Println("*--------SWITCH-------*")

	fmt.Println(diaDaSemana(2))

}
