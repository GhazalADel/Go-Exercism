package piglatin

import (
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}
var consonantSounds = []string{"squ", "thr", "th", "sch", "qu", "ch", "rh", "p", "k", "x", "q", "y", "m", "f", "r"}

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	translated := []string{}
	for _, word := range words {
		translated = append(translated, translate(word))
	}
	return strings.Join(translated, " ")
}

func translate(word string) string {
	isStart := false
	for _, vowel := range vowels {
		if strings.HasPrefix(word, vowel) {
			isStart = true
			break
		}
	}
	if isStart ||
		strings.HasPrefix(word, "xr") ||
		strings.HasPrefix(word, "yt") {
		return word + "ay"
	}
	if startsWithConsonantSound(word) {
		return handleConsonantSound(word)
	}
	return ""
}

func startsWithConsonantSound(sentence string) bool {
	for _, cs := range consonantSounds {
		if strings.HasPrefix(sentence, cs) {
			return true
		}
	}
	return false
}

func handleConsonantSound(sentence string) string {
	for _, cs := range consonantSounds {
		if strings.HasPrefix(sentence, cs) {
			return strings.TrimPrefix(sentence, cs) + cs + "ay"
		}
	}
	return ""
}
