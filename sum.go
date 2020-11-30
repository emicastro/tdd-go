package main

// Sum given an array of numbers, it add all and return the result
func Sum(numbers []int) (total int) {
	for _, v := range numbers {
		total += v
	}
	return
}

// SumAll take a varying number of slices, returning a new slice containing the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// SumAllTails take a varying number of slices and calculates the totals of the "tails" of each slice.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return
}
