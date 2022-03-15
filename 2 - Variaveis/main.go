package main

import "fmt"

func main() {
	// expondo o tipo
	var variavel1 string = "var 1"
	// inferencia de tipo
	variavel2 := "Variavel 2"
	// explicitando o tipo
	var (
		var3 string
		var4 string = "lalalala"
	)
	// varias instancias com inferencia de tipo
	vari5, var6 := "variavel5", "variavel6"
	//constante
	const contante = "contante"

	fmt.Println(variavel1, variavel2, var3, var4, vari5, var6, contante)
}
