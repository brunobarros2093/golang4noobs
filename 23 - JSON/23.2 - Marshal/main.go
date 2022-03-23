package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json: "nome"`
	Raca string `json: "raca"`
	//Idade uint   `json: "idade"`
}

func main() {
	fmt.Println("PASSANDO STRUCT OU MAP PARA JSON ")
	cachorro := cachorro{"Lili", "vira-lata"}
	fmt.Println(cachorro)
	// Marshal retorna um slice de bytes[]
	cachorroEmJSON, err := json.Marshal(cachorro)
	if err != nil {
		log.Fatal(err)
	}
	// NOS RETORNA UM SLICE DE []UINT8
	fmt.Println(cachorroEmJSON)
	// retorna um json
	// struct PARA json
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	cachorroMap := map[string]string{
		"nome": "Mel",
		"raca": "vira-lata",
	}
	cachorro2emJSON, err := json.Marshal(cachorroMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewBuffer(cachorro2emJSON))

}
