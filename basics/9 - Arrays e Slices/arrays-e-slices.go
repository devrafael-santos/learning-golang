package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//Array

	// Todos os valores têm que ser do mesmo tipo
	var array1 [5]string

	array1[0] = "Posicao 0"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3"}
	fmt.Println(array2)

	// Tamanho baseado na quantidade de valores passados na inicialização
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice não precisa se preocupar com o tamanho
	// Aponta para um array, não é um array
	// Slice : fatia
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array3))

	slice = append(slice, 4)
	fmt.Println(slice)

	// inclusivo e exclusivo, respectivamente
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNOS | AULA 10

	// MAKE ARMAZENA UM LOCAL NA MEMÓRIA
	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
