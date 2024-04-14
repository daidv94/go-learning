package arrays

func Sum(a []int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}

	return sums
}
