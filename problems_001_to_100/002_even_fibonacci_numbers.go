package problems_001_to_100

func FiboEvenSum(n int) int {
	sum, prev, current := 0, 1, 2

	for current <= n {
		if isEven(current) {
			sum += current
		}

		prev, current = current, prev+current
	}

	return sum
}

func isEven(n int) bool {
	return n%2 == 0
}
