package main

// Pointers allow you to store a values address in memory
// allowing programs to more memory efficient. To access their
// value, you have to dereference it.

// To reference use an &
// To deference use *

import "fmt"

var name = "Waylon"

func changeName(name *string) string {
	*name = "Duff"
	return *name
}

func main() {
	// This creates a reference
	nameReference := &name
	fmt.Println("The value reference value is", nameReference)
	fmt.Println("The value of name is", name)
	changeName(nameReference)
	fmt.Println("Now the name is", name)
}
