package transpose

func Transpose(input []string) []string {
	res := make([]string, 0)
	if len(input) == 0 {
		return res
	}
	max := len(input[0])
	for _, v := range input {
		if len(string(v)) > max {
			max = len(v)
		}
	}

	for i := 0; i < max; i++ {
		tmp := ""
		for j := 0; j < len(input); j++ {
			if len(input[j])-1 >= i {
				tmp += string(input[j][i])
			} else {
				tmp += " "
			}
		}
		res = append(res, tmp)
	}
	return res
}
