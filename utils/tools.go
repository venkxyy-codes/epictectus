package utils

import (
	"regexp"
)

// ValidatePhoneNumber validates if the given phone number is a valid Indian number (10 digits).
func ValidatePhoneNumber(phone string) bool {
	// Regular expression to match exactly 10 digits.
	var phoneRegex = regexp.MustCompile(`^[6-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

func ValidatePassword(password string) string {
	if len(password) < 8 {
		return "err-password-must-be-atleast-8-characters-long"
	}

	hasDigit := regexp.MustCompile(`[0-9]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\.\?\/\\|~]`)

	if !hasDigit.MatchString(password) {
		return "err-password-must-contain-atleast-1-digit"
	}
	if !hasUpper.MatchString(password) {
		return "err-password-must-contain-atleast-1-uppercase-letter"
	}
	if !hasSpecial.MatchString(password) {
		return "err-password-must-contain-atleast-1-special-character"
	}
	return ""
}
