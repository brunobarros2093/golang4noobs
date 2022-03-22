package main

import (
	"fmt"
	"intro-a-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Paulista")
	fmt.Println(tipoEndereco)
}
