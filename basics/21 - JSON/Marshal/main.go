package main

import (
	"bytes"
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
	c := cachorro{"Bolt", "Shiba inu", 4}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))
}
