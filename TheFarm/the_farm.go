package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, num int) (float64, error) {
	if total, err := fc.FodderAmount(num); err != nil {
		return 0, err
	} else if factor, err := fc.FatteningFactor(); err != nil {
		return 0, err
	} else {
		return total / float64(num) * factor, nil
	}
}
func ValidateInputAndDivideFood(fc FodderCalculator, num int) (float64, error) {
	if num <= 0 {
		return 0, errors.New("invalid number of cows")
	} else {
		return DivideFood(fc, num)
	}
}

type InvalidCowsError struct {
	num int
	msg string
}

func (rev *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", rev.num, rev.msg)
}
func NewInvalidCowsError(num int, msg string) *InvalidCowsError {
	return &InvalidCowsError{num, msg}
}
func ValidateNumberOfCows(num int) error {
	if num < 0 {
		return NewInvalidCowsError(num, "there are no negative cows")
	} else if num == 0 {
		return NewInvalidCowsError(num, "no cows don't need food")
	} else {
		return nil
	}
}
