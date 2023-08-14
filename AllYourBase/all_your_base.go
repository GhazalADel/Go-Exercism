package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	allZero := true
	for _, v := range inputDigits {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero || len(inputDigits) == 0 {
		return []int{0}, nil
	}
	baseTen := 0
	for i, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		baseTen += v * int(math.Pow(float64(inputBase), float64(len(inputDigits)-i-1)))
	}
	res := make([]int, 0)
	for baseTen != 0 {
		res = append([]int{baseTen % outputBase}, res...)
		baseTen = int(math.Floor(float64(baseTen / outputBase)))
	}

	return res, nil
}
