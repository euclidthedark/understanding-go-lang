package maps

import "fmt"

func DeclareAMap () map {
	myMap = make(map[int]bool)

	return myMap
}

func AccessMapProperties () {
	theMap = DeclareAMap()
	theMap[1] = true

	fmt.Println("Log a map attribute: ", theMap)
}