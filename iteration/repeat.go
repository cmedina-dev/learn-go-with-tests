package iteration

// Repeat accepts a string and returns a new string with the original appended five times.
func Repeat(input string, count int) string {
	if count <= 0 {
		return input
	}

	var repeated string
	for i := 0; i < count; i++ {
		repeated += input
	}
	return repeated
}
