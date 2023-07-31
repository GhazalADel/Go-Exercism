package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	res := ""
	for _, char := range log {
		if char == 10071 {
			res = "recommendation"
		} else if char == 128269 {
			res = "search"
		} else if char == 9728 {
			res = "weather"
		}
		if res != "" {
			break
		}
	}
	if res == "" {
		res = "default"
	}
	return res
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	tmp := ""
	for _, char := range log {
		if char == oldRune {
			tmp += string(newRune)
		} else {
			tmp += string(char)
		}
	}
	return tmp
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	}
	return false
}
