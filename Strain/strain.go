package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	res := make([]int, 0)
	for _, v := range i {
		if filter(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	res := make([]int, 0)
	for _, v := range i {
		if !filter(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	res := make([][]int, 0)
	for _, v := range l {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	res := make([]string, 0)
	for _, v := range s {
		if filter(string(v)) {
			res = append(res, string(v))
		}
	}
	return res
}
