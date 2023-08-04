package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n == 1 {
		return 0, nil
	}
	if n < 1 {
		return 0, errors.New("")
	}
	count := 0
	for n != 1 {
		count++
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n++
		}
	}
	return count, nil
}
