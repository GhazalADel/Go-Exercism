// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "'", "")
	tmp := ""
	for _, v := range s {
		if v == 32 || (v >= 65 && v <= 90) {
			tmp += string(v)
		} else {
			tmp += " "
		}
	}
	tmpSplitted := strings.Split(tmp, " ")
	res := ""
	for _, v := range tmpSplitted {
		if strings.TrimSpace(v) != "" {
			res += string(v[0])
		}
	}
	return res
}
