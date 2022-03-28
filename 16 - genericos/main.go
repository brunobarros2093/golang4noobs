package main

import "fmt"

// como não há tipo definido, essa interface aceita QUALQUER TIPO, logo, é genérica
func generica(interf interface{}) {
	fmt.Println(interf)
}

// interfaces apenas possuem assinatura de método
type forma interface {
}

// type parameter
// type inferance
// type Set
//---------\/ = faz o mesmo que utilizar interface{}
func min[T any](a T, b T) T {
	return a
}

// o tio (tilde em ingles) - diz para o Go que td objeto que compõe o float64 será permitido, assim como INT
type minTypes interface {
	~float64 | int
}

//-----------\/ type Sets
func max[T minTypes](a T) T {
	return a
}

func main() {
	fmt.Println("--- GENERICOS ANTES DA VERSÃO 1.18 -- UTILIZANDO INTERFACE")
	generica("Sou uma String")
	generica(123)
	generica(true)

	fmt.Println("--- GENERICOS APÓS DA VERSÃO 1.18 -- UTILIZANDO UM TIPO GENERICO T interface como parametro")
	//--------------\/ opcional
	fmt.Println(min[int](1, 2))
	fmt.Println(min("bruno", "birros"))
}
