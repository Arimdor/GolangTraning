package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int

	fmt.Println("Numero 1: ")
	fmt.Scan(&n1)
	fmt.Println("Numero 2: ")
	fmt.Scan(&n2)

	println(n1, "/", n2, "=", n1/n2, "sobra", n1%n2)

}
