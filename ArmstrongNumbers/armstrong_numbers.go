package armstrong

import (
	"fmt"
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	digits := fmt.Sprint(n)
	sum := 0
	for _, digit := range digits {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		p := int(math.Pow(float64(d), float64(len(digits))))
		sum += p
	}
	return n == sum
}
