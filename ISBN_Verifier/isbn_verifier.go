package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	starter := 10
	sum := 0
	if len(isbn) < 10 {
		return false
	}
	for i, v := range isbn {
		if i == len(isbn)-1 {
			if string(v) == "X" {
				sum += 10
			} else {
				digit, err := strconv.Atoi(string(v))
				if err != nil {
					return false
				}
				sum += digit
			}
		} else {
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return false
			}
			sum += (starter * digit)
		}
		starter--
	}
	return sum%11 == 0
}
