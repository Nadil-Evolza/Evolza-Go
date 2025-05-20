package main

import (
	"fmt"
	"unicode"
)

func CheckPassword() {
	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	if isStrong(password) {
		fmt.Println("Password is strong!")
	} else {
		fmt.Println("Password is weak. Please include at least 8 characters with uppercase, lowercase, and numbers.")
	}
}

func isStrong(password string) bool {
	var hasUpper, hasLower, hasNumber bool

	if len(password) < 8 {
		return false
	}

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}
