package bottlesong

import (
	"fmt"
	"unicode"
)

func extenseNum(num int) string {
	switch num {
	case 10:
		return "ten"
	case 9:
		return "nine"
	case 8:
		return "eight"
	case 7:
		return "seven"
	case 6:
		return "six"
	case 5:
		return "five"
	case 4:
		return "four"
	case 3:
		return "three"
	case 2:
		return "two"
	case 1:
		return "one"
	case 0:
		return "no"
	default:
		panic("Number not implemented")
	}
}
func Capitalize(s string) string {
	rs := []rune(s)
	if len(rs) > 0 {
		rs[0] = unicode.ToUpper(rs[0])
	}
	return string(rs)
}
func bottle(num int) string {
	if num == 1 {
		return "bottle"
	}
	return "bottles"
}
func Recite(startBottles, takeDown int) []string {
	poem := []string{}
	bottleCount := startBottles
	takeLeft := takeDown
	for bottleCount > 0 {
		for i := 0; i < 2; i++ {
			number := Capitalize(extenseNum(bottleCount))
			poem = append(poem, fmt.Sprintf("%s green %s hanging on the wall,", number, bottle(bottleCount)))
		}
		poem = append(poem, "And if one green bottle should accidentally fall,")
		bottleCount--
		takeLeft--
		poem = append(poem, fmt.Sprintf("There'll be %s green %s hanging on the wall.", extenseNum(bottleCount), bottle(bottleCount)))
		if takeLeft == 0 {
			break
		}
		poem = append(poem, "")
	}
	return poem
}
