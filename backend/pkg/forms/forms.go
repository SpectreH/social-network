package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// Form creates a custom from struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// NewForm initializes a form struct
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("%s cannot be blank", splitString(field)))
		}
	}
}

// MinLenght checks for string minimum lenght
func (f *Form) MinLenght(field string, lenght int) bool {
	x := f.Get(field)
	if len(x) < lenght {
		f.Errors.Add(field, fmt.Sprintf("%s must be at least %d characters long", splitString(field), lenght))
		return false
	}

	return true
}

// IsEmail checks for valid email
func (f *Form) IsEmail(field string) bool {
	res, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, f.Get(field))
	if !res {
		f.Errors.Add(field, "Please provide a valid email address")
		return false
	}
	return true
}

// splitString splits a string at uppercase letters
func splitString(str string) string {
	str = strings.Title(str)
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	submatchall := re.FindAllString(str, -1)
	return strings.Join(submatchall, " ")
}
