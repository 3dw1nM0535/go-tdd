package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrorKeyNotFound = errors.New("could not find the word you looking for")

func (d Dictionary) Search(k string) (string, error) {
	definition, ok := d[k]
	if !ok {
		return "", ErrorKeyNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
