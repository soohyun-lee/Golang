package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errornotfound = errors.New("Not found")
var errwordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errornotfound
}

// Add a word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errornotfound:
		d[word] = def
	case nil:
		return errwordExists
	}
	return nil
}
