package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Struct in go programming.")

	var user1 User

	user1 = User{
		Name: "Selim",
		Age:  25,
	}

	fmt.Println("User1 Name: ", user1)
	fmt.Println("User1 Name: ", user1.Name)
	fmt.Println("User1 Age: ", user1.Age)

	user2 := User{
		Name: "John",
		Age:  30,
	}
	fmt.Println("User2 Name: ", user2.Name)
	fmt.Println("User2 Age: ", user2.Age)
}
