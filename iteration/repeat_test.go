package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("repeats entered character five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(got, want, t)
	})

	t.Run("repeats entered string three times", func(t *testing.T) {
		got := Repeat("abc", 3)
		want := "abcabcabc"
		assertCorrectMessage(got, want, t)
	})

	t.Run("returns the original string if count is equal to zero", func(t *testing.T) {
		got := Repeat("abc", 0)
		want := "abc"
		assertCorrectMessage(got, want, t)
	})

	t.Run("returns the original string if count is less than zero", func(t *testing.T) {
		got := Repeat("abc", -5)
		want := "abc"
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
