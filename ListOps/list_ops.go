package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, 0)
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

func (s IntList) Reverse() IntList {
	res := make(IntList, 0)
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	for _, v := range lst {
		s = append(s, v)
	}
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	res := make(IntList, 0)
	res = append(res, s...)
	for _, l := range lists {
		for _, v := range l {
			res = append(res, v)
		}
	}
	return res
}
