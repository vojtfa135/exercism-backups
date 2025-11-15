package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numCows int
	msg     string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.msg)
}

// TODO: define the 'DivideFood' function

func DivideFood(fodderCalc FodderCalculator, numCows int) (float64, error) {
	var (
		fodderAmnt, fatFactor float64
		err                   error
	)

	if fodderAmnt, err = fodderCalc.FodderAmount(numCows); err != nil {
		return 0, err
	}

	if fatFactor, err = fodderCalc.FatteningFactor(); err != nil {
		return 0, err
	}

	return (fodderAmnt / float64(numCows)) * fatFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalc FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(fodderCalc, numCows)
	}

	return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numCows: numCows,
			msg:     "there are no negative cows",
		}
	} else if numCows == 0 {
		return &InvalidCowsError{
			numCows: numCows,
			msg:     "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
