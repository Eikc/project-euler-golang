package exercises

import (
	"math"
)

func getClosesNumberToZeroIn(numbers []int) int {
	lowestNumberInArray := int(math.Abs(float64(numbers[0])))

	for _, number := range numbers {
		absoluteNumber := math.Abs(float64(number))
		if int(absoluteNumber) < lowestNumberInArray {
			lowestNumberInArray = number
		}
	}

	return lowestNumberInArray
}
