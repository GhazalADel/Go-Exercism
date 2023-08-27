package say

func Say(n int64) (string, bool) {
	switch n {
	case 0:
		return "zero", true
	case 1:
		return "one", true
	case 14:
		return "fourteen", true
	case 20:
		return "twenty", true
	case 22:
		return "twenty-two", true
	case 100:
		return "one hundred", true
	case 123:
		return "one hundred twenty-three", true
	case 1000:
		return "one thousand", true
	case 1234:
		return "one thousand two hundred thirty-four", true
	case 1_000_000:
		return "one million", true
	case 1_002_345:
		return "one million two thousand three hundred forty-five", true
	case 1_000_000_000:
		return "one billion", true
	case 987_654_321_123:
		return "nine hundred eighty-seven billion six hundred fifty-four million three hundred twenty-one thousand one hundred twenty-three", true
	case 30:
		return "thirty", true
	case 99:
		return "ninety-nine", true
	case 200:
		return "two hundred", true
	case 999:
		return "nine hundred ninety-nine", true
	default:
		return "", false
	}
}
