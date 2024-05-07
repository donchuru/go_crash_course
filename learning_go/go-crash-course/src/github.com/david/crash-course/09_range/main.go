package main

import (
	"fmt"
)

func main() {
	ids := []int {44, 8888}

	// Loop through ids slice
	for i, id := range ids {
		fmt.Printf("%d - ID:%d\n", i, id)
	}

	// no need for index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
}