package validator

import (
	"fmt"
	"time"
)

// Function to validate a string value based on the minimum and maximum length
func ValidateString(value string, minLen int, maxLen int) error {
	n := len(value)
	if n < minLen || n > maxLen {
		return fmt.Errorf("must contain %d-%d characters", minLen, maxLen)
	}
	return nil
}

// Function to validate UUID
func ValidateUuid(uuid_str string) error {
	// Assuming the Id should be a positive integer
	if len(uuid_str) == 0 {
		return fmt.Errorf("'uuid' must not be empty")
	}
	return nil
}

func ValidateUsername(username string) error {
	// Should be at least 3 and maximum 24 characters long
	return ValidateString(username, 3, 24)
}

func ValidateUserAgent(userAgent string) error {
	if len(userAgent) == 0 {
		return fmt.Errorf("'user_agent' must not be empty")
	}
	return nil
}

func ValidateClientIp(clientIp string) error {
	if len(clientIp) == 0 {
		return fmt.Errorf("'client_ip' must not be empty")
	}
	return nil
}

func ValidateExpiration(expires_at time.Time) error {
	if expires_at.IsZero() {
		return fmt.Errorf("'expires_at' is not a valid time")
	}
	
	if expires_at.Before(time.Now()) {
		return fmt.Errorf("'refresh_token' already expired")
	}
	return nil
}
