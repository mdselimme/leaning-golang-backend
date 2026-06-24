package main

import "fmt"

var x int = 5 // Global variable

func main() {
	x := 10
	fmt.Println("Outer x:", x)
	{
		x := 20
		fmt.Println("Inner x:", x)
	}
	fmt.Println("Outer x after inner block:", x)
}
