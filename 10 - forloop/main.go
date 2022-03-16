package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	// O unico loop em Go é o FOR, que é usado de várias formas
	for i < 10 {
		time.Sleep(time.Second)
		i++
	}
	fmt.Println(i)

	// estilo o if-init (inicia uma variavel)
	fmt.Println("--------for padrao- -----")
	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}
}
