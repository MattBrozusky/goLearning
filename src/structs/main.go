package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	person1 := person{
		firstName: "alex",
		lastName:  "bob",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 23456,
		},
	}

	person1.updateName("matt")
	person1.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
	fmt.Printf("%v\n", p.contactInfo.email)
	fmt.Printf("%v\n", p.firstName)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName

}
