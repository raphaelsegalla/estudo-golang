package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{"Retangulo", 10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida {
			// t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) // para a execução
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{"Circulo", 10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
