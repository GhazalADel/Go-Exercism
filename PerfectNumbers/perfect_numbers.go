package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification string

const ClassificationAbundant Classification = "ClassificationAbundant"
const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationPerfect Classification = "ClassificationPerfect"

var ErrOnlyPositive = errors.New("error only positive numbers can be classified")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassificationDeficient, ErrOnlyPositive
	}
	if n == 1 || n == 2 {
		return ClassificationDeficient, nil
	}
	var tot int64 = 1
	var i int64
	for i = 2; i <= int64(float64(n)/2); i++ {
		if n%i == 0 {
			tot += i
		}
	}
	if tot == n {
		return ClassificationPerfect, nil
	} else if tot > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
