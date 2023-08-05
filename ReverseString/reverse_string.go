package reverse

import "strings"

func Reverse(s string) string {
	chars := strings.Split(s, "")
	res := make([]string, 0)

	for i := len(chars) - 1; i >= 0; i-- {
		res = append(res, chars[i])
	}
	return strings.Join(res, "")
}
