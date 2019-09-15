package main

import "fmt"

// Go is a typed language, but usually the Go compiller is smart enough
// to infer types

// This is a constant variable
const bossesName = "Boss"

// You can declare multiple go variables inline
var (
	number   = 10
	floatVar = 10.0
	name     = "Waylon"
)

func main() {
	// You can change vars
	number = 12
	// This is the short hand way of declaring variables in Go
	// This is usually done at the function level
	shortHandString := "Duff"
	fmt.Println(number, floatVar, name, shortHandString, bossesName)
}
