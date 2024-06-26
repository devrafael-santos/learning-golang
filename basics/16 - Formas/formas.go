package main

import "fmt"

type forma interface {
	area() float32
}

type retangulo struct {
	altura  float32
	largura float32
}

func (r retangulo) area() float32 {
	return r.altura * r.largura
}

func areaDaForma(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

func main() {
	r := retangulo{10, 10}

	fmt.Println(r.area())
}
