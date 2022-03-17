package main

import "fmt"

// função com retorno nomeado
func calcuosMath(num1, num2 int) (soma int, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

// função com numero variavel de parametros - transforma os parametros em um slice
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	fmt.Println("Funçoes com retorno nomeado")

	sum, sub := calcuosMath(2, 5)
	fmt.Println(sum, sub)
	fmt.Println("Funçoes com numero indefinido de parametros")
	fmt.Println(soma(1, 2, 3, 4, 5, 7, 8))
	fmt.Println("Funçoes anonimas")

	func() {
		fmt.Println("Conteudo da Funcao anonima!!!!")
	}()

}
