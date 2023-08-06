package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	mapping := make(map[string]bool)
	for _, v := range input {
		if len(mapping) == 26 {
			break
		}
		if !(v >= 97 && v <= 122) {
			continue
		}
		if _, ok := mapping[string(v)]; !ok {
			mapping[string(v)] = true
		}
	}
	return len(mapping) == 26
}
