package main

import "fmt"

func func1() {
	fmt.Println("Func 1")
}

func func2() {
	fmt.Println("Func 2")
}

func main() {
	defer func1()
	func2()
}
