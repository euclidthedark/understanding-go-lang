package loops

import "fmt"

// Below is an example of an infinite loop
func InfiniteLoop() {
	for {
		fmt.Println("What is love?")
	}
}

// Below is an example of an iterable loop
func IteratingOverArraysOrSlices() {
	theSlice := []string{"a", "b", "c"}

	for _, i := range theSlice {
		fmt.Println(i)
	}
}

// Below is a regular for loop
func AnExampleForLoop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		}
		if i == 8 {
			break
		} // This will break the flow of the code, even though the simple statement in the loop syntax has a different conditional.
	}
}

// I think the declaration in the loop will scope up to the next context.
func AnExampleWhileLoop () {
	for left, right := 1, 10; left < right {
		fmt.Println("Working in a while loop!")

		left += 1
	}
}
