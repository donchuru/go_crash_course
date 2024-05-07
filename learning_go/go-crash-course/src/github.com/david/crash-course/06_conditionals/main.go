package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10

	// common convention is to not use parentheses
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	}


	color := "yellow";

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
