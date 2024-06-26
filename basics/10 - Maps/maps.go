package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Rafael",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"], usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Rafael",
			"ultimo":   "Costa",
		},
		"curso": {
			"nome":   "ADS",
			"Campus": "Pituaçu",
		},
	}

	fmt.Println("-------")

	fmt.Println(usuario2)

	fmt.Println(usuario2)
}
