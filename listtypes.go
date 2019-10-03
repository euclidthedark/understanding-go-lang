package main

import "fmt"

// This is an array literal
var arrayForFunction = [5]int{}

func printAnArrayLiteral() {
	fmt.Println(arrayForFunction)
}

func declareASlice() {
	// slices are pointers to an array, so you can manipulate the list
	// below is the short hand syntax to create slices.
	// you can also use the make function s := make([]int, len=3, cap=3)
	s := []int{1, 2, 3}
	fmt.Println(s)
}

func appendElementsToASlice() {
	// append ones to the slice
	sliceOfOnes := []int{}

	for i := 0; i < 3; i++ {
		sliceOfOnes = append(sliceOfOnes, i)
	}
	fmt.Println(sliceOfOnes)
}

func sliceTheSlice() {
	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Manipulating slices is very similar to list types
	// in Python
	fmt.Println(aSlice[:5])
	fmt.Println(aSlice[5:])
}

func main() {
	printAnArrayLiteral()
	declareASlice()
	appendElementsToASlice()
	sliceTheSlice()
}
