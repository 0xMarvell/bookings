package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct and embeds a url.Values object.
type Form struct {
	url.Values
	Errors errors
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
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}
