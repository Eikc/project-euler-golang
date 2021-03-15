package main

import "fmt"

func main() {
	number := 600851475143

	for i := 2; i < number; i++ {
		for number%i == 0 {
			number = number / i
		}
	}

	fmt.Printf("result is: %v \n", number)

}
