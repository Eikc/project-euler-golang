package main

import (
	"fmt"
	"math"
)

func main() {
	count := 1
	candidate := 1

	for count < 10001 {
		candidate = candidate + 2
		if isPrime(candidate) {
			count++
		}

	}

	fmt.Printf("Result is: %v \n", candidate)
}

func isPrime(number int) bool {
	r := int(math.Floor(math.Sqrt(float64(number))))
	for i := 2; i <= r; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
