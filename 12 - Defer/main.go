package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao1")
}
func funcao2() {
	fmt.Println("Executando funcao2")
}

func main() {
	// defer = adiar a execução de um código até o ultimo momento possivel
	defer funcao1()
	funcao2()
}
