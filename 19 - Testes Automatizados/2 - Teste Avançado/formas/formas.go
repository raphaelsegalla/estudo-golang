package formas

import (
	// "fmt"
	"math"
)

type Forma interface {
	area() float64
	name() string
}

// func EscrevendoArea(f Forma) {
// 	fmt.Printf("A área da forma %s é %0.2f\n", f.name(), f.area())
// }

type Retangulo struct {
	Nome    string
	Altura  float64
	Largura float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

func (r Retangulo) name() string {
	return r.Nome
}

type Circulo struct {
	Nome string
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (c Circulo) name() string {
	return c.Nome
}
