package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10

	// Copiando o valor da variavel1 e jogando para a variavel2
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	// Ele não faz uma cópia do valor e coloca em outra variável
	// Você passa o endereço na memória

	// Guarda um valor inteiro
	var variavel3 int = 0

	// Guarda um endereço de memória de um valor inteiro
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	// Desreferênciação
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150

	fmt.Println(variavel3, *ponteiro)

}
