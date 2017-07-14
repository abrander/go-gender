// Package gender provides a few functions to validate a persons gender from a
// string.
package gender

// Valid returns true if gender is a valid gender.
func Valid(gender string) bool {
	return (Validate(gender) == nil)
}

// Validate returns an error if the gender is not valid. This error may be
// passed to the end-user.
func Validate(gender string) error {
	return nil
}
