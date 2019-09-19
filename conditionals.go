package main

import "fmt"

func functionShouldStdOutSpecial() {
	// conditional if with simple statement
	if specialRelativity, generalRelativity := 1905, 1915; specialRelativity < generalRelativity {
		fmt.Println("Special")
	} else {
		fmt.Println("General")
	}
}

func functionShouldStdOutGeneral() {
	switch generalRelativity := 1915; generalRelativity {
	case 1905:
		fmt.Println("General Relativity was not published in 1905.")
	case 1915:
		fmt.Println("That's correct. General Relativity was published in 1915.")
	default:
		fmt.Println("Well, it was for sure published at some point in history!")
	}
}

func main() {
	functionShouldStdOutSpecial()
	functionShouldStdOutGeneral()
}
