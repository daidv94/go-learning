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

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, n := range numbers {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			tail := n[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
