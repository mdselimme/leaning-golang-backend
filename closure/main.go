package main

import (
	"fmt"
)

const a = 10

var p = 100

func outer() func() {
	money := 1000
	age := 20

	fmt.Println("age = ", age)

	show := func() {
		money = money + a + p
		fmt.Println("money = ", money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("init() function called")
}
