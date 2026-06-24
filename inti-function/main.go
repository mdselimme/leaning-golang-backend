package main

import "fmt"

var num int = 42 // Global variable

func main() {
	fmt.Println("Main function")
	fmt.Println("Global variable num:", num)
}

func init() {
	fmt.Println("Init function called before main")
	num = 100 // Local variable shadows the global variable
}
