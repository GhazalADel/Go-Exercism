package main

import "fmt"

func RotationalCipher(plain string, shiftKey int) string {
	res := ""
	shiftKey %= 26
	for _, v := range plain {
		var tmp int = int(v)
		tmp += shiftKey
		//lower
		if v >= 97 {
			if tmp > 122 {
				tmp -= 26
			}
		} else {
			if tmp > 90 {
				tmp -= 26
			}
		}
		res += string(tmp)
	}
	return res
}

func main() {
	fmt.Println(RotationalCipher("a", 0))
}
