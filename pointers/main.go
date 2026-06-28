package main

import "fmt"

func passByReference(numbers *[3]int) {
	fmt.Println("Array is : ", numbers)
}

func passByValue(numbers [3]int) {
	fmt.Println("Array is : ", numbers)
}

type Person struct {
	name   string
	age    int
	salary float64
}

func main() {
	fmt.Println("Hello world pointer.")

	x := 10

	fmt.Println("Number is : ", x)

	addr := &x
	fmt.Println("Address of x is : ", addr)

	valueOfX := *addr
	fmt.Println("Value of x is : ", valueOfX)

	*addr = 20
	fmt.Println("New value of x is : ", x)

	numbers := [3]int{1, 2, 3}
	passByReference(&numbers)
	passByValue(numbers)

	person1 := Person{name: "John", age: 30, salary: 50000.0}

	fmt.Println("Person1 details: ", person1)

	personAddr := &person1
	fmt.Println("Address of person1 is : ", *personAddr)

}
