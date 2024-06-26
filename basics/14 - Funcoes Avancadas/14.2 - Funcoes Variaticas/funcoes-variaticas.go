package main

import "fmt"

func printTela(nome string, numeros ...int) {
	for _, num := range numeros {
		fmt.Println(nome, num)
	}
}

func main() {
	printTela("Hello World", 2, 3, 4, 5, 5, 1)
}
