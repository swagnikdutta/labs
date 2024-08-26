package maps

import (
	"errors"
)

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word since because does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary is a wrapper around a map. Using this type we've created the search method that will work on the dictionary instance.
// This technique improves the dictionary's usage, just that I don't know why exactly.
// Need some stronger reasons
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, exists := d[key]
	if !exists {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case err == nil:
		return ErrWordExists
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
