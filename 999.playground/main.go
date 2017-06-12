package main

import (
	"fmt"
)

func main() {
	for i := 0; ; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}
