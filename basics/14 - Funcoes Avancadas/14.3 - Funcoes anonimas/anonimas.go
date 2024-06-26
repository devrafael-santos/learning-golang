package main

import "fmt"

func main() {

	resultado := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(resultado)
}
