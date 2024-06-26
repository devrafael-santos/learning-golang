package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float32
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(float64(c.Raio), 2)
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func AreaDaForma(f Forma) {
	fmt.Printf("A área da forma é %0.2f", f.Area())
}
