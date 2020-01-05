package functions

// When declaring functions, the signatures must be type declared
// You do not have to declare return types

// Below is an example of a function that returns two types
func MultiplyTwoIntsByTwoAndThree(i1, i2 int) (int1, int2 int) {
	i1 = i1 * 2
	i2 = i2 * 3

	return i1, i2
}

// When you don't know how many parameters are going to passed to a function
// you should declare a variadic function.
// Notice the elipses in the function signature.

func SumSuppliedNumbers(numbers ...int) int {
	origin := 0

	for _, number := range numbers {
		origin += number
	}
	return origin
}
