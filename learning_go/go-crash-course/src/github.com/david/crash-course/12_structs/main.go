package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method (value receiver)
// Value receiver because we didn't change any of the values in the Person object
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Pointer receiver can change the object attributes
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	} else {
		return
	}
}

func main() {
	// Init person using struct
	p1 := Person {firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "M", age: 25}

	// p2 := Person {"Samantha", "Smith", "Boston", "M", 25}
	fmt.Println(p1.age)

	p1.hasBirthday()
	p1.hasBirthday()
	p1.hasBirthday()
	p1.getMarried("Segura")
	fmt.Println(p1.greet())
}