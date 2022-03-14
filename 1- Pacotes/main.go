package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("bruno@email.com")
	fmt.Println(erro)
	// auxiliar.Escrever()

}
