package arrays_and_slices

func SumArray(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
