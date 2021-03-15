package exercises

func fibOne(number int) int {
	var x, y = 0, 1

	for i := 2; i <= number; i++ {
		r := x + y
		x, y = y, r
	}

	return y
}
