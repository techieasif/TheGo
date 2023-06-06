package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

//Form is a custom struct type, embed url.Values object
type Form struct {
	url.Values
	Errors errors
}

//Valid returns true if there are no errors, otherwise false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initializes the form struct
func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: map[string][]string{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// Has checks if form field is in post and is not empty
func (f *Form) Has(field string) bool {
	fieldValue := f.Get(field)
	if fieldValue == "" {
		return false
	}
	return true
}

//MinLength checks for minimum length required for a particular field.
func (f *Form) MinLength(field string, length int) bool {
	val := f.Get(field)

	if len(val) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be atleast of %d characters long", length))
		return false
	}

	return true
}

//IsEmail checks provided email is valid or not
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Values.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
