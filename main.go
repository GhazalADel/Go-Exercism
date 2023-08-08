package main

import (
	"fmt"
)

func SumMultiples(limit int, divisors ...int) int {
	tot := 0
	mapping := make(map[int]bool)
	for _, v := range divisors {
		start := 1
		for start*v < limit {
			if _, ok := mapping[start*v]; !ok {
				mapping[start*v] = true
				tot += start * v
			}
			start++
		}
	}
	return tot
}

func main() {
	fmt.Println(SumMultiples(7, 3))
}
