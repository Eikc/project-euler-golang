package main

import "fmt"

func main() {
	number := 100

	var sumOfSquares int
	var squareOfTheSum int

	for i := 1; i <= number; i++ {
		sumOfSquares += i * i
		squareOfTheSum += i
	}

	squareOfTheSum = squareOfTheSum * squareOfTheSum
	diff := squareOfTheSum - sumOfSquares

	fmt.Printf("result is: %v \n", diff)
}
