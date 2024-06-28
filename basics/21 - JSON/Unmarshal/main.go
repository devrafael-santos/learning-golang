package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string
	Raca  string
	Idade uint
}

func main() {
	cachorroJson := `{"Nome":"Bolt","Raca":"Shiba inu","Idade":4}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)
}
