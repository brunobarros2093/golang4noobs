package formas

import (
	"math"
)

// Forma interfaces apenas possuem assinatura de método
type Forma interface {
	area() float64
}

// Retangulo exportavel
type Retangulo struct {
	Altura  float64
	Largura float64
}

//Circulo exportavel
type Circulo struct {
	// Raio float
	Raio float64
}

//Area criado o método area() que retorna um float para o retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

//Area criado o método area() que retorna um float para o retangulo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
