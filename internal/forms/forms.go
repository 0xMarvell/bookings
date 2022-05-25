package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct and embeds a url.Values object.
type Form struct {
	url.Values
	Errors errors
}

// Required checks if the required form fields have the right value
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Valid returns true if there are no error in inputted form data,
// otherwise Valid returns false.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initilizes a form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if required form field(s) is not empty.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		// f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// Minlength checks if the inputted field data surpasses
// the minimum length of data expected.
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	fieldData := r.Form.Get(field)

	if len(fieldData) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d characters long", length))
		return false
	}

	return true
}

// IsEmail checks if email address is valid.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "invalid email address")
	}
}
