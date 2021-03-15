package main

import "fmt"

func main() {
	current := 20
	smallest := current

	for {
		if moduloOfAll(smallest, current) {
			break
		}
		smallest += current
	}

	fmt.Printf("result is: %v \n", smallest)
}

func moduloOfAll(number int, maxNumber int) bool {
	for i := 2; i <= maxNumber; i++ {
		if number%i != 0 {
			return false
		}
	}

	return true
}
