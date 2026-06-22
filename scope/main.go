package main

import "fmt"

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func main() {
	result := Sum(10, 20)
	fmt.Println("Sum of 10 and 20 is:", result)
	// fmt.Println("Sum of 5 and 3 is:", Sum(5, 3))
	fmt.Println("Your name is:", getUserName())
}
