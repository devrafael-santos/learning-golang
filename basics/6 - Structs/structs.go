package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario

	fmt.Println(u)

	u.nome = "Rafael Santos"
	u.idade = 19

	fmt.Println(u)

	enderecoExemplo := endereco{"New Horizon", 0}

	// Criando com inferencia
	usuario2 := usuario{"Rafael Santos", 20, enderecoExemplo}

	fmt.Println(usuario2)

	// Criando com inferencia com apenas um campo
	usuario3 := usuario{nome: "Rodrigo"}

	fmt.Println(usuario3)

	usuario4 := usuario{idade: 22}

	fmt.Println(usuario4)

}
