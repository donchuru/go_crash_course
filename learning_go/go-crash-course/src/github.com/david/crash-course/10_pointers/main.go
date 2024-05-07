package main

import (
	"fmt"
)

func main() {
	a := [1]string {"yello world"}
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	fmt.Println((*b)[0])

	// change val with pointer
	// Pointers are good for object passing. Just pass the pointer rather than the object itself
	(*b)[0] = "Hello wld"
	fmt.Println((*b)[0])
	fmt.Println(a[0])


}