package main

import "fmt"

func main() {

	numero := 0

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior que 0")
	} else {
		fmt.Println("Numero menor ou igual a 0")
	}

}
