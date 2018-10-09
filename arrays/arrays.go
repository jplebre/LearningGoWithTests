package arrays

// See how arrays need the lenght as part of the type?
// Passing an [4]int would cause the compiler to blow up
func SumWithArrays(numbers [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}

func SumWithArraysUsingRange(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumWithSlices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// How to make an array with a specified length
	// sums := make([]int, len(numbersToSum))

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, SumWithSlices(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, SumWithSlices(tail))
	}

	return sums
}
