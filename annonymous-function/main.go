package main

import "fmt"

func main() {
	func(a int, b int) {
		fmt.Println("Anonymous function called with values:", a, b)
	}(5, 10)

	sum := func(a int, b int) int {
		return a + b
	}

	result := sum(3, 7)
	fmt.Println("Sum of 3 and 7 is:", result)
}

func init() {
	fmt.Println("init() function called")
}
