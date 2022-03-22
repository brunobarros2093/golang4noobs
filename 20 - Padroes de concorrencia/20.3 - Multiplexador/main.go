package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola mundo"), escrever("Tchau mundo"))

	for range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		fmt.Println(<-canal)
	}
}

// recebe dois ou mais canais e retorna um
func multiplexar(canalEntrada1 <-chan string, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canalEntrada1:
				canalDeSaida <- msg
			case msg := <-canalEntrada2:
				canalDeSaida <- msg
			}
		}
	}()
	return canalDeSaida
}
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
