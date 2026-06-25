package main

import "fmt"

func processOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Higher Order function.")
	result := processOperation(5, 10, add)
	fmt.Println("The result is:", result)
}
