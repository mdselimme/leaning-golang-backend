package main

import (
	"fmt"
)

func main() {
	fmt.Println("Noob Function")

	add := func(x int, y int) int {
		return x + y
	}

	var result = add(10, 20)
	fmt.Println("Result:", result)

}
