package arrays_and_slices

// SumArray adds all integers within a fixed array of size five.
func SumArray(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumSlice adds all integers within a slice of any size.
func SumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumSlices returns a new slice with each index representing the sum of one slice.
func SumSlices(slices ...[]int) (result []int) {
	for _, valA := range slices {
		result = append(result, SumSlice(valA))
	}
	return
}

// SumSlicesTails returns a new slice with each index representing the sum of the slice minus the head.
func SumSlicesTails(slices ...[]int) (result []int) {
	for _, valA := range slices {
		if len(valA) == 0 {
			result = append(result, 0)
		} else {
			tail := valA[1:]
			result = append(result, SumSlice(tail))
		}
	}
	return
}
