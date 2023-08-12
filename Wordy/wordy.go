package wordy

import (
	"regexp"
	"strconv"
)

func ExtractNumbers(question string, middle int) (int, int, error) {
	return 0, 0, nil
}

func Answer(question string) (int, bool) {
	if re := regexp.MustCompile("what is \\d+\\?"); re.MatchString(question) {
		i := 8
		numStr := ""
		for i < len(question)-1 && numStr[i] >= 48 && numStr[i] <= 57 {
			numStr += string(question[i])
			i++
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, true
		}
		return num, false
	} else if re := regexp.MustCompile("what is \\d+ plus \\d+\\?"); re.MatchString(question) {
		firstNumStr, secondNumStr := "", ""
		i := 8
		for i < len(question)-1 && firstNumStr[i] >= 48 && firstNumStr[i] <= 57 {
			firstNumStr += string(question[i])
			i++
		}
		i += 5
		for i < len(question)-1 && secondNumStr[i] >= 48 && secondNumStr[i] <= 57 {
			secondNumStr += string(question[i])
			i++
		}
		num1, err := strconv.Atoi(firstNumStr)
		if err != nil {
			return 0, true
		}
		num2, err := strconv.Atoi(secondNumStr)
		if err != nil {
			return 0, true
		}
		return num1 + num2, false
	} else if re := regexp.MustCompile("what is \\d+ plus \\d+\\?"); re.MatchString(question) {

	}
	return 0, false
}
