package maps

import (
	"errors"
)

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("the word you tried to add already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update because it does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrWordNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
