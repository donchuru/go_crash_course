// Closures are Anonymous functions

package main

import (
	"fmt"
)

// why the hell would anyone want to do this
func adder () func (int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {

	sum := adder ()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}