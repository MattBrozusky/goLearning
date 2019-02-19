package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	person1 := person{
		firstName: "alex",
		lastName:  "bob",
	}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.lastName)
	person1.firstName = "tony"
	fmt.Println(person1.firstName)
}
