package luhn

import (
	"errors"
	"strconv"
	"strings"
)

func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}
	checkDigit, err := checkSum(s)

	if err != nil {
		return false
	}

	return checkDigit%10 == 0
}

func checkSum(s string) (int, error) {
	chars := strings.Split(s, "")
	result := make([]string, len(s))

	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	r := strings.Join(result, "")
	sum := 0
	for i, v := range r {
		digit, err := strconv.Atoi(string(v))

		if err != nil {

			return 0, errors.New("")
		}

		if (i % 2) == 0 {
			sum += digit
		} else {
			product := digit * 2
			if product > 9 {
				product -= 9
			}
			sum += product
		}
	}
	return sum, nil
}
