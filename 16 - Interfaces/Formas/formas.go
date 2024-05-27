package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
	name() string
}

func escrevendoArea(f forma) {
	fmt.Printf("A área da forma %s é %0.2f\n", f.name(), f.area())
}

type retangulo struct {
	nome    string
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (r retangulo) name() string {
	return r.nome
}

type circulo struct {
	nome string
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) name() string {
	return c.nome
}

func main() {
	fmt.Println("interfaces")
	r := retangulo{"Retangulo", 10, 15}
	escrevendoArea(r)
	c := circulo{"Circulo", 10}
	escrevendoArea(c)
}
