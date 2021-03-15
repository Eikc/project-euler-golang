package main

import "fmt"

func main() {

	value := 1000

	//bruteforcing:
	for a := 3; a < value; a++ {
		for b := a + 1; b < value; b++ {
			c := value - b - a

			// a < b < c
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				fmt.Println("Found: ", a, b, c)
				fmt.Println("Product: ", a*b*c)
			}
		}
	}
}
