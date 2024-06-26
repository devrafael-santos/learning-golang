package main

import "fmt"

func recuperarAplicacao() {
	if i := recover(); i != nil {
		fmt.Println("Aplicação recuperada com sucesso.")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarAplicacao()

	if (n1+n2)/2 < 6 {
		fmt.Println("Reprovado")
	} else if (n1+n2)/2 > 6 {
		fmt.Println("Aprovado")
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(5, 4))
}
