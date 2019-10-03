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
		if i%2 == 0 {
			fmt.Println("Even")
		}
		if i == 8 {
			break
		} // This will break the flow of the code, even though the simple statement in the loop syntax has a different conditional.
	}
}

func main() {
	// infiniteLoop()
	//iteratingOverArraysOrSlices()
	anExampleForLoop()
}
