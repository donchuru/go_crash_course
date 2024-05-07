package main

import "fmt" // io

// runs automatically
func main() {
	
	// you don't need string here
	// var name string = "Brad"

	// you have to use every variable you use
	var age = 37

	var isCool = true

	// Shorthand
	name := "David"

	size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
