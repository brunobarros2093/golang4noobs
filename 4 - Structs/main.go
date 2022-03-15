package main

import "fmt"

// definindo um tipo usuario que é um struct
type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Estudo de strutcs")
	//jeitos de definir um struct
	var user usuario
	user.idade = 15
	user.nome = "bruno"
	// aninhamento de structs
	endereco := endereco{"avenida getulio vargas", 425}
	//com  inferencia de tipo
	user2 := usuario{"bruno2", 21, endereco}
	fmt.Println(user2, user2.endereco.logradouro)

	// com inferencia de tipo sem dados completos
	user3 := usuario{idade: 21}
	fmt.Println(user3)
}
