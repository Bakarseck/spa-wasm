package utils

import (
	"regexp"
	"strconv"
)

func matchesRegex(s string, pattern string) bool {
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(s)
}

type Validator interface {
	Validate(string) bool
}

type UsernameValidator struct{}

func (UsernameValidator) Validate(s string) bool {
	return matchesRegex(s, `^[\w\d-]+$`)
}

type AgeValidator struct{}

func (AgeValidator) Validate(s string) bool {
	ageInt, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return ageInt >= 0
}

type GenderValidator struct{}

func (GenderValidator) Validate(s string) bool {
	return matchesRegex(s, `^[MF]$`)
}

type EmailValidator struct{}

func (EmailValidator) Validate(s string) bool {
	return matchesRegex(s, `^[\w\d._%+]+@[\w\d.]+\.[a-z]{2,}$`)
}

type PasswordValidator struct{}

func (PasswordValidator) Validate(s string) bool {
	has := func(pattern string) bool { return matchesRegex(s, pattern) }
	return len(s) >= 8 &&
		has(`[a-z]`) &&
		has(`[A-Z]`) &&
		has(`[0-9]`)
}

// FirstNameValidator is a type that can validate first names.
type FirstNameValidator struct{}

// Validate checks if the given string is a valid first name.
// A valid first name is a string of letters that is between 4 and 10 characters long.
func (FirstNameValidator) Validate(s string) bool {
	return matchesRegex(s, `^[a-zA-Z]{4,10}$`)
}

// LastNameValidator is a type that can validate last names.
type LastNameValidator struct{}

// Validate checks if the given string is a valid last name.
// A valid last name is a string of letters that is between 2 and 10 characters long.
func (LastNameValidator) Validate(s string) bool {
	return matchesRegex(s, `^[a-zA-Z]{2,10}$`)
}