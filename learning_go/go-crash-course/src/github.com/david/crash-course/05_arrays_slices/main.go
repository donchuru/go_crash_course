package main

import (
	"fmt"
)


func main() {
	// Arrays have fixed size, 	// Slices dont

	// var fruitArr [2] string
	// fruitArr [0] = "Apple"
	// fruitArr [1] = "Orange"

	// initialize
	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice[1:])

}
