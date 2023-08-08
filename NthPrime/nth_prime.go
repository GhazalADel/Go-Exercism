package prime

import (
	"errors"
	"math"
)

func isPrime(number int) bool {
	if number == 1 {
		return false
	}
	if number == 2 {
		return true
	}
	for i := 2; i < int(math.Sqrt(float64(number)))+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n should be greater than one")
	}
	if n == 1 {
		return 2, nil
	}
	count := 1
	start := 3
	for count != n {
		if isPrime(start) {
			count++
		}
		start += 2
	}
	return start - 2, nil
}
