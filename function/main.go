package main

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getUserAge() int {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)
	return age
}

func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func addTwoNumbers() {
	num1, num2 := getTwoNumbers()
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
}

func userInformation() {
	fmt.Println("Welcome to the application.")

	name := getUserName()
	fmt.Printf("Hello, %s! Enjoy using the application.\n", name)

	age := getUserAge()
	fmt.Printf("You are %d years old.\n", age)
}

func main() {
	fmt.Println("Functions in Go")
	result := Sum(5, 10)
	fmt.Println("The sum is:", result)
	userInformation()
	addTwoNumbers()
}
