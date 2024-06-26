package main

import "fmt"

type customer struct {
	nome  string
	idade uint8
}

func (c customer) buy() {
	fmt.Printf("The customer %s have buy something", c.nome)
}

func main() {
	custumer1 := customer{"Alex", 19}

	custumer1.buy()
}
