package main

import (
	"fmt"
	"strconv"
)

func main() {
	digit := 0
	digits := 999

	for i := 1; i <= digits; i++ {
		for j := 1; j <= digits; j++ {
			r := i * j
			s := strconv.Itoa(r)
			reversed := reverse(s)

			if s == reversed && r > digit {
				digit = r
			}
		}
	}

	fmt.Printf("result is: %v \n", digit)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
