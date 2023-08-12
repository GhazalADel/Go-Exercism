package cryptosquare

import "strings"

func NormalizeString(pt string) string {
	res := ""
	pt = strings.ToLower(pt)
	i := 0

	for i < len(pt) {
		if pt[i] >= 97 && pt[i] <= 122 {
			res += string(pt[i])
			i++
		} else if pt[i] == 32 {
			index := i
			for pt[index] == 32 {
				index++
			}
			index--
			if index-i == 0 {
				i++
			} else {
				res += " "
				i += (index - i + 1)
			}
		} else {
			i++
		}
	}
	return res
}
func Encode(pt string) string {
	pt = NormalizeString(pt)
	c := 1
	r := 1
	for r*c < len(pt) {
		c++
		r = c
		if r*c >= len(pt) {
			break
		}
		r--
	}
	res := ""
	for i := 0; i < r; i++ {
		count := 0
		for j := i * c; i < j*(c+1) && j < len(pt); j++ {
			count++
			res += string(pt[j])
		}
		if count != c {
			for k := 0; k < c-count; k++ {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}
