package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	res := ""
	shiftKey %= 26
	for _, v := range plain {
		var tmp int = int(v)
		tmp += shiftKey
		//lower
		if v >= 97 && v <= 122 {
			if tmp > 122 {
				tmp -= 26
			}
			res += string(int32(tmp))
		} else if v >= 65 && v <= 90 {
			if tmp > 90 {
				tmp -= 26
			}
			res += string(int32(tmp))
		} else {
			res += string(v)
		}
	}
	return res
}
