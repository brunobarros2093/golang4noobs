package formas_test

import (
	"math"
	"testes-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retangulo", func(t *testing.T) {
		ret := formas.Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := formas.Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

}
