package util

import "strings"

// IsBlank Checks if a given string is empty or whitespaces only
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
