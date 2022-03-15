package main

import "fmt"

func main() {
	fmt.Println(somar(4, 4))

	// funções são tipo!
	var f = func() {
		fmt.Println("Vim da F")
	}

	f()
	// retorno da função com multiplos valores
	resultadoSoma, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSub)
}

// funcoes com multiplos retornos (lindo)
func calculosMatematicos(numero1, numero2 int) (int, int) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao
}

//---nome--parametros/tipo---tipo do retorno
func somar(n1 int, n2 int) int {
	return n1 + n2
}
