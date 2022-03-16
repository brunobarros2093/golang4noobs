package main

import "fmt"

func main() {
	fmt.Println("-----MAPS-----")
	//--------tipo-chave--valor
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "pedrim",
	}
	fmt.Println(usuario)
	fmt.Printf("--- usuario[nome]: %s ----\n", usuario["nome"])

	fmt.Println("------------------")
	usuario2 := map[string]map[string]string{
		"312321": {"primeiro": "ultimo"},
		"222312": {"primeiro": "ultimo"},
	}

	fmt.Println(usuario2)

}
