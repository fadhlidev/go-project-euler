package problems_001_to_100

func MultiplesOf3And5(number int) int {
	return sumOfMultiples(number, 3) + sumOfMultiples(number, 5) - sumOfMultiples(number, 15)
}

func sumOfMultiples(number, divisor int) int {
	n := (number - 1) / divisor
	return divisor * n * (n + 1) / 2
}
