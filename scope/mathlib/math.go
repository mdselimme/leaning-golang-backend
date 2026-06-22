package mathlib

import "fmt"

func Add(a int, b int) int {
	sum := a + b
	fmt.Println("Number is: ", sum)
	return sum
}

var Num = 10