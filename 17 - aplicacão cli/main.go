package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	applicacao := app.Gerar()
	// if init
	if err := applicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
