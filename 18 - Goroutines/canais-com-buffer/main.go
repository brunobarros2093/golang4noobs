package main

import "fmt"

func main() {
	// é a mesma instanciação de um buffer comum, apenas dizemos o numero de canais que serão abertos
	// dessa forma o Go sabe que abriremos dois canais
	// o canal só bloqueará quando o buffer chegar no limite
	canal := make(chan string, 2)
	canal <- "Ola mundo!"
	canal <- "Ola mundo 2!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
