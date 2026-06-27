package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUser(usr User) {
	fmt.Println("Name is = ", usr.Name)
	fmt.Println("Age is = ", usr.Age)
}

//receiver function in go
func (usr User) printDetails() {
	fmt.Println("Name is = ", usr.Name)
	fmt.Println("Age is = ", usr.Age)
}

//receiver function in go
func (usr User) call(a string) {
	fmt.Println("Name is = ", usr.Name)
	fmt.Println("Age is = ", usr.Age)
	fmt.Println("Cousin name is = ", a)
}

func main() {
	fmt.Println("Receiver Function in go programming.")

	user1 := User{
		Name: "Selim",
		Age:  25,
	}

	// printUser(user1)
	user1.printDetails()
	user2 := User{
		Name: "Selim",
		Age:  25,
	}

	// printUser(user2)
	user2.call("2")

}
