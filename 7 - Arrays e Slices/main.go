package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// DECLARAÇÃO DE TAMANHO E TIPO - TUDO CONTIDO NELE SERÁ INT
	var arrayForma1 [5]string
	arrayForma1[4] = "Posição 1"
	fmt.Println(arrayForma1)
	// OUTRA FORMA DE DECLARAÇÃO
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	// SLICES
	// Slices ainda possuem a limitação de tipo, mas são flexiveis em tamanho
	// Slice aponta para um array
	slice := []int{}
	slice = append(slice, 34)
	fmt.Println(slice)
	// Diferença entre slices e arrays
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println("*------------ARRAYS INTERNOS-------------*")
	// ARRAYS INTERNOS
	slice3 := make([]float32, 10, 15)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
