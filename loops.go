package main

import "fmt"

// Below is an example of an infinite loop
func infiniteLoop() {
	for {
		fmt.Println("What is love?")
	}
}

// Below is an example of an iterable loop
func iteratingOverArraysOrSlices() {
	theSlice := []string{"a", "b", "c"}

	for _, i := range theSlice {
		fmt.Println(i)
	}
}

// Below is a regular for loop
func anExampleForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	// infiniteLoop()
	//iteratingOverArraysOrSlices()
	anExampleForLoop()
}
