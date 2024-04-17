package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("returns the definition if the word exists", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		definition := "this is just a test"
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("returns an error when the word does not exist", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrWordNotFound)
	})
	t.Run("adds a word to an existing dictionary", func(t *testing.T) {
		word := "test2"
		dictionary := Dictionary{"test": "this is just a test"}
		definition := "this is a second test"
		err := dictionary.Add("test2", definition)
		if err != nil {
			t.Fatal("error not expected: ", err)
		}
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("does not add a word if it already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "this test has been overwritten")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("updates the definition of an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "this is the updated definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)
		if err != nil {

			t.Fatal("error not expected: ", err)
		}
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("does not update words that do not exist", func(t *testing.T) {
		word := "test2"
		definition := "this is just a test"
		newDefinition := "this is the updated definition"
		dictionary := Dictionary{"test": definition}
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
	t.Run("deletes an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrWordNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("did not expect an error: ", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
