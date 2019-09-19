package main

import (
	"fmt"
	"os"
)

func functionShouldStdOutSpecial() {
	// conditional if with simple statement
	if specialRelativity, generalRelativity := 1905, 1915; specialRelativity < generalRelativity {
		fmt.Println("Special")
	} else {
		fmt.Println("General")
	}
}

func functionShouldStdOutGeneral() {
	// Simple simple statements also apply to switches
	// Go switches have implicit break statements.
	// You can also evaluate multiple expressions with a ,
	switch generalRelativity := 1915; generalRelativity {
	case 1905:
		fmt.Println("General Relativity was not published in 1905.")
	case 1915, 0:
		fmt.Println("That's correct. General Relativity was published in 1915.")
	default:
		fmt.Println("Well, it was for sure published at some point in history!")
	}
}

func throwErrorBecauseFileDoesNotExist() {
	// Important note, you can also incorporate errors in simple statements as well.
	_, err := os.Open("./txt.txt")

	if err != nil {
		fmt.Println("The error is: ", err)
	}
}

func main() {
	functionShouldStdOutSpecial()
	functionShouldStdOutGeneral()
	throwErrorBecauseFileDoesNotExist()
}
