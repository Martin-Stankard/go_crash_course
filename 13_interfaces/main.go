package main

import (
	"fmt"
	"strconv"
)

//define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//pointer receiver
func (p *Person) hasBirthday() {

	p.age++
}

//struct related ...value receiver
func (p Person) greet() string {
	return "hello my name is " + p.firstName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {

		p.lastName = spouseLastName
	}
}

func main() {

	fmt.Println("structs")

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "asdf", gender: "f", age: 25}
	person2 := Person{"bob", "Su", "Town", "m", 11}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.greet())
	person2.hasBirthday()
	person2.getMarried("marriedName")
	fmt.Println(person2)
	person1.getMarried("marriedName")
	fmt.Println(person1)
}
