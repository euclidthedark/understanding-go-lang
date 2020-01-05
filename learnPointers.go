package pointers

// Pointers allow you to store a values address in memory
// allowing programs to more memory efficient. To access their
// value, you have to dereference it.

// Also, remeber golang itself will create copy's of data types
// as you pass them around in functions

// To reference use an &
// To dereference use *

import "fmt"

var name = "Waylon"

func ChangeName(name *string) string {
	// deference the name and change the value
	*name = "Duff"
	return *name
}

func SeeHowPointersWork() {
	// This creates a reference
	nameReference := &name
	fmt.Println("The value reference value is", nameReference)
	fmt.Println("The value of name is", name)
	changeName(nameReference)
	fmt.Println("Now the name is", name)
}
