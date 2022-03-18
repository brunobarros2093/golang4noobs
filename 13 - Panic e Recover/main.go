package main

import "fmt"

func recuperaExecutacao() {
	fmt.Println("TENTATIVA DE RECUPERAR EXECUTACAO")
	// se r (retorno de recover() for diferente de null/nil, então é sucesso)
	if r := recover(); r != nil {
		fmt.Println("EXC RECUPERADA COM SUCESSO")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperaExecutacao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// interrompe o fluxo normal e para TUDO
	panic("!!!!A media é exatamente 6!!!!")
}

func main() {
	// panic and recover
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
