package forms

import "reflect"

type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// GetFirst returns the first error message with the key
func (e errors) GetFirst() (string, string) {
	keys := reflect.ValueOf(e).MapKeys()

	for _, key := range keys {
		res := e.Get(key.String())
		return res, key.String()
	}

	return "", ""
}

// Get returns the first error message of current field
func (e errors) Get(field string) string {
	yes := e[field]
	if len(yes) == 0 {
		return ""
	}

	return yes[0]
}
