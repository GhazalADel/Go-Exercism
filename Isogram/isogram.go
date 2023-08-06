package isogram

import "strings"

func IsIsogram(word string) bool {
	mapping := make(map[string]int)
	word = strings.ToLower(word)
	for _, v := range word {
		if v == 32 || v == 45 {
			continue
		}
		if _, ok := mapping[string(v)]; ok {
			return false
		}
		mapping[string(v)] = 1
	}
	return true
}
