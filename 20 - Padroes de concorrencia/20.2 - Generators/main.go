package main

import (
	"fmt"
	"time"
)

func main() {
	// A Função já retorna um canal e já escreve de forma assincrona, por isso não há necessidade de
	// utilizar o 'go' aqui na frente - afinal, ele já é concorrente e já retorna um canal
	canal := escrever("Ola canal generator")
	for i := 0; i <= 10; i++ {
		fmt.Println(<-canal)
	}
}

// função generator - gera um canal assincrono
// retorna um canal de escrita (só recebe)
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
