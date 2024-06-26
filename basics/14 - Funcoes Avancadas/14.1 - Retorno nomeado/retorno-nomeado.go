package main

import "fmt"

func somaESub(num1, num2 int) (soma int, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

func main() {
	soma, subtracao := somaESub(10, 10)

	fmt.Println(soma, subtracao)
}
