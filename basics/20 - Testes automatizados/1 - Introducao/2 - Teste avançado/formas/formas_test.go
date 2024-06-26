package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área recebida %0.2f diferente da esperada: %0.2f", areaEsperada, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * math.Pow(cir.Raio, 2))
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área recebida %0.2f diferente da esperada: %0.2f", areaEsperada, areaEsperada)
		}
	})
}
