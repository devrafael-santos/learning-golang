package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	nomes := [3]string{"Rafael", "Caio", "Rian"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
}
