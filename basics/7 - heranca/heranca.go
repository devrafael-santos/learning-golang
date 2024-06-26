package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Rafael", "Santos", 19, 182}

	e1 := estudante{p1, "ADS", "UCSAL"}

	fmt.Println(e1)
}
