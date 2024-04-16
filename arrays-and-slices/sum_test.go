package arrays_and_slices

import (
	"fmt"
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
