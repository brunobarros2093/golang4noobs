package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

// todos os usuarios terão o método salvar
// receiver
func (u usuario) salvar() {
	fmt.Println(u.idade)
	fmt.Println("Metodo salvar do usuario")
}

func main() {
	usuario1 := usuario{"usuario1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

}
