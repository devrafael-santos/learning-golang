package main

import "fmt"

// Pode ter uma funçao init por arquivo (não por pacote)
func init() {
	fmt.Println("Executando a funcao antes do metodo main")
}

func main() {
	fmt.Println("Executando a funcao min")

}
