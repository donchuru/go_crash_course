package main

import (
	"fmt"
)

func main() {
	// emails := make(map[string][]string)

	// // Assign kv
	// emails["Bob"] = append(emails["Bob"], "bob@gmail.com", "gibberish")
	// emails["Sharon"] = append(emails["Bob"], "bob@gmail.com") // Truly cursed feature here
	// emails["Mike"] = append(emails["Mike"], "mike@gmail.com", "gibberish")

	emails := map[string]string{"Bob": "dont worry", "Sharon": "about a thing"}

	emails["Mike"] = "everything gonna be alright"

	fmt.Println(emails)
	fmt.Println(emails["Sharon"])

	// Delete from map
	// delete(emails, "Bob")
	// fmt.Println(emails);

}