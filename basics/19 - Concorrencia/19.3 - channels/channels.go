package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Hello World", canal)

	for msg := range canal {
		fmt.Println(msg)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
