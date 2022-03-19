package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// criado o método area() que retorna um float para o retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// interfaces apenas possuem assinatura de método
type forma interface {
	area() float64
}

// o método recebe um tipo da interface
func escreverArea(f forma) {
	fmt.Printf("A area é: %f \n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
}
