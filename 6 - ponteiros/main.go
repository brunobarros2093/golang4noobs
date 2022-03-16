package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	// passagem por valor
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)
	// ponteiro é uma referencia de memória - aponta para o end. de memória da variável

	var variavel3 int = 100
	// FOI PASSADO UM TIPO INTEIRO
	var ponteiro *int = &variavel3
	// O VALOR DE PONTEIRO AQUI É UM ENDEREÇO DE MEMÓRIA
	fmt.Println(variavel3, ponteiro)
	// DESREFERENCIAÇÃO - AQUI VEREMOS O VALOR CONTIDO NO ENDEREÇO DE MEMÓRIA
	fmt.Println(*ponteiro)

}
