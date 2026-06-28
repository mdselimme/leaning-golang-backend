package main

import "fmt"

func main() {
	fmt.Println("Array in go.")

	// var num [5]int

	// num[0] = 10
	// num[1] = 20
	// num[2] = 30
	// num[3] = 40
	// num[4] = 50

	num := [5]int{10, 20, 30, 40, 50}

	name := [5]string{"John", "Doe", "Alice", "Bob", "Charlie"}

	fmt.Println("Array of integers:", num)
	fmt.Println("Array of strings:", name[2])

	name[2] = "Eve"
	fmt.Println("Array of strings after change:", name[2])

	fmt.Println("Array of integers:", num[4])

}
