package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	res := make([]string, 0)
	subject = strings.ToLower(subject)
	chars := []rune(subject)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	sortedSbj := string(chars)
	for _, v := range candidates {
		tmp := string(v)
		v = strings.ToLower(v)
		if v == subject {
			continue
		}
		chars = []rune(v)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		sortedV := string(chars)
		if sortedV == sortedSbj {
			res = append(res, tmp)
		}
	}
	return res
}
