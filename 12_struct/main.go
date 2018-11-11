package main

import (
	"fmt"
	"strconv"
)

//Define the Person Struct
type person struct {
	firstname string
	lastname  string
	city      string
	age       int
}

func main() {

	person1 := person{firstname: "Omkar", lastname: "khatavkar",
		city: "Pune", age: 23}
	fmt.Println(person1)
	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
}

//greeting nethod
func (p person) greet() string {
	return "Hello Mr " +
		p.firstname + " " + p.lastname + " and your age is " +
		strconv.Itoa(p.age)

}

func (p *person) hasBirthday() {
	p.age++
}
