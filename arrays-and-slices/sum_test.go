package arrays_and_slices

import (
	"fmt"
	"slices"
	"testing"
)

func BenchmarkSumArray(b *testing.B) {
	data := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		SumArray(data)
	}
}

func ExampleSumArray() {
	data := [5]int{1, 2, 3, 4, 5}
	fmt.Println(SumArray(data))
	// Output: 15
}

func TestSumArray(t *testing.T) {
	t.Run("sums all integers within a fixed array", func(t *testing.T) {
		data := [5]int{1, 2, 3, 4, 5}
		got := SumArray(data)
		want := 15
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, data)
		}
	})
}

func BenchmarkSumSlice(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		SumSlice(data)
	}
}

func ExampleSumSlice() {
	data := []int{1, 2, 3}
	fmt.Println(SumSlice(data))
	// Output: 6
}

func TestSumSlice(t *testing.T) {
	t.Run("sums three integers within a slice", func(t *testing.T) {
		data := []int{1, 2, 3}
		got := SumSlice(data)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, data)
		}
	})

	t.Run("sums five integers within a slice", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		got := SumSlice(data)
		want := 15
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, data)
		}
	})
}

func BenchmarkSumSlices(b *testing.B) {
	dataA := []int{1, 2, 3}
	dataB := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		SumSlices(dataA, dataB)
	}
}

func ExampleSumSlices() {
	dataA := []int{1, 2, 3}
	dataB := []int{4, 5, 6}
	fmt.Println(SumSlices(dataA, dataB))
	// Output: [6 15]
}

func TestSumSlices(t *testing.T) {
	t.Run("sums the values of two slices", func(t *testing.T) {
		dataA := []int{1, 2, 3}
		dataB := []int{1, 2, 3, 4, 5}
		got := SumSlices(dataA, dataB)
		want := []int{6, 15}
		checkSliceSums(t, got, want, dataA, dataB)
	})
}

func BenchmarkSumSlicesTails(b *testing.B) {
	dataA := []int{1, 2, 3}
	dataB := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		SumSlicesTails(dataA, dataB)
	}
}

func ExampleSumSlicesTails() {
	dataA := []int{1, 2, 3}
	dataB := []int{1, 2, 3, 4, 5}
	fmt.Println(SumSlicesTails(dataA, dataB))
	// Output: [5 14]
}

func TestSumSlicesTails(t *testing.T) {
	t.Run("sums the values of all slices excluding the head", func(t *testing.T) {
		dataA := []int{1, 2, 3}
		dataB := []int{1, 2, 3, 4, 5}
		got := SumSlicesTails(dataA, dataB)
		want := []int{5, 14}
		checkSliceSums(t, got, want, dataA, dataB)
	})

	t.Run("returns 0 when an empty slice is provided", func(t *testing.T) {
		var dataA []int
		dataB := []int{1, 2, 3}
		got := SumSlicesTails(dataA, dataB)
		want := []int{0, 5}
		checkSliceSums(t, got, want, dataA, dataB)
	})
}

func checkSliceSums(t *testing.T, got, want, dataA, dataB []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %d, want %d, given {%v, %v}", got, want, dataA, dataB)
	}
}
