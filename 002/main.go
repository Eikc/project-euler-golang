package main

import "fmt"

func main() {
	x, y := 1, 2
	sum := 0

	for y <= 4000000 {
		if y%2 == 0 {
			sum += y
		}

		x, y = y, x+y
	}

	fmt.Printf("sum is: %v \n", sum)
}
