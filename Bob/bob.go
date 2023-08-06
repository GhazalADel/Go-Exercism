// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.ReplaceAll(remark, "\n", "")
	remark = strings.ReplaceAll(remark, " ", "")

	hasLetter := false
	for _, v := range remark {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			hasLetter = true
			break
		}
	}
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	} else if strings.ToUpper(remark) == remark && string(remark[len(remark)-1]) == "?" && hasLetter {
		return "Calm down, I know what I'm doing!"
	} else if string(strings.TrimSpace(remark)[len(remark)-1]) == "?" {
		return "Sure."
	} else if strings.ToUpper(remark) == remark && hasLetter {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
