package main

import "fmt"

type pessoa struct {
	nome      string
	idade     int
	altura    int
	sobrenome string
}

// estudante é uma pessoa
type estudante struct {
	// pegando todos os campos de pessoa e jogando em estudante
	// sem tipo, pois assim os campos serão copiados diretamente
	pessoa
	curso     string
	faculdade string
}

// IMPORTANTE: NÃO HÁ HERENÇA EM GO - ISSO É O MAIS PROXIMO QUE CHEGA
func main() {
	fmt.Println("Herança que não é herança")

	p1 := pessoa{"Joao", 20, 194, "pedrim"}
	e1 := estudante{p1, "filosofia", "usp"}
	fmt.Println(p1)
	// por ter passado pessoa sem um tipo no struct, os campos foram copiados
	// se passar "pessoa pessoa", acessariamos os campos por e1.pessoa.nome, pois seria criado uma propriedade de nome pessoa
	fmt.Println(e1, e1.nome)
}
