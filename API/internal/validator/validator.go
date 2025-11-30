package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]))+$")
)

// Define a new Validator type which contains a map of validation errors.
type Validator struct {
	Errors map[string]string
}

// new function initializes a new Validator instance.
// It returns a pointer to the Validator with an empty Errors map.
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Valid method checks if there are any validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError method adds a validation error for a specific field.
func (v *Validator) AddError(field, message string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}
	v.Errors[field] = message
}

// Check.
func (v *Validator) Check(ok bool, message string, field string) {
	if !ok {
		v.AddError(field, message)
	}
}

// In returns true if a specific value is in a list of strings.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}


// validate passwordplaintext checks if the password is not empty and meets length requirements.
func ValidatePasswordPlaintext(v *Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

// validate hashed password checks if the hashed password is not empty.
func ValidatePasswordHash(v *Validator, hash []byte) {
	v.Check(len(hash) > 0, "password", "must be provided")
}