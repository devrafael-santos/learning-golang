package main

import "fmt"

func main() {
	// ARITMETICOS
	// +
	// -
	// /
	// %
	// *

	// ATRIBUIÇÁO
	// =
	// :=

	// RELACIONAIS
	// <
	// >
	// <=
	// >=
	// ==
	// !=

	// OPERADORES LÓGICOS
	// E : &&
	// OU : ||
	// NÃO : !

	// OPERADORES UNIARIOS
	// Não existe --numero, apenas numero--
	// --
	// ++
	// += 15
	// -= 15
	// /= 15
	// *= 15

	// Operador Ternário não existe

	numero := 4

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)

}
