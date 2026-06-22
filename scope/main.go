package main

import (
	"fmt"

	"example.com/mathlib"
)

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func main() {
	result := mathlib.Add(10, 20)
	fmt.Println("Sum of 10 and 20 is:", result)
	fmt.Println("Value of Num is:", mathlib.Num)
	// fmt.Println("Sum of 5 and 3 is:", mathlib.Add(5, 3))
	fmt.Println("Your name is:", getUserName())
}
