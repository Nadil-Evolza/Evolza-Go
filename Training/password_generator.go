package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	lowerLetters   = "abcdefghijklmnopqrstuvwxyz"
	upperLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func GeneratePassword(length int, useUpper, useLower, useDigits, useSpecial bool) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	charset := ""
	if useLower {
		charset += lowerLetters
	}
	if useUpper {
		charset += upperLetters
	}
	if useDigits {
		charset += digits
	}
	if useSpecial {
		charset += specialChars
	}

	if len(charset) == 0 {
		return "", fmt.Errorf("at least one character type must be selected")
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}
