package atbash

import "strings"

func Atbash(s string) string {
	s = strings.ToLower(s)
	res := ""
	cLet := 0
	for _, v := range s {
		if v == 32 {
			continue
		}
		if v >= 97 && v <= 122 {
			tmp := v - 97
			res += string(122 - tmp)
			cLet++
		} else if v >= 48 && v <= 57 {
			res += string(v)
			cLet++
		}

		if cLet == 5 {
			res += " "
			cLet = 0
		}
	}
	if strings.HasSuffix(res, " ") {
		res = res[:len(res)-1]
	}
	return res
}
