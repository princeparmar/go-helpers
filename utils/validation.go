package utils

import "regexp"

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	// Use regex to validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidateMobile(mobile string) bool {
	if mobile == "" {
		return false
	}

	// Use regex to validate mobile format
	mobileRegex := regexp.MustCompile(`^[0-9]{10}$`)
	return mobileRegex.MatchString(mobile)
}
