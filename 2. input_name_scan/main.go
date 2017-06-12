package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
