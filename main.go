package main

import "fmt"

func CollatzConjecture(n int) (int, error) {
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
func main() {
	r, _ := CollatzConjecture(12)
	fmt.Println(r)
}
