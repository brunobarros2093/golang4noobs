package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade string `json: "idade"`
}

func main() {
	fmt.Println("Passar um JSON para STRUCT ou MAP")
	cachorroEmJSON := `{"Nome":"Lili","Raca":"vira-lata","Idade":"1"}`
	// cria-se o cachorro
	var c cachorro
	// passase o json e o endereço de memoria de quem vai receber o conteudo
	// o Unmarshal tbm precisa receber um slice de bytes[] como parametro em JSON
	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println("O JSON em struct: ", c)
	fmt.Println("------------MAP-----------------")
	cachorroMap := map[string]string{}
	if err := json.Unmarshal([]byte(cachorroEmJSON), &cachorroMap); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cachorro em MAP: ", cachorroMap)
}
