package main

import "fmt"

// como não há tipo definido, essa interface aceita QUALQUER TIPO, logo, é genérica
func generica(interf interface{}) {
	fmt.Println(interf)
}

// interfaces apenas possuem assinatura de método
type forma interface {
}

func main() {
	generica("Sou uma String")
	generica(123)
	generica(true)
}
