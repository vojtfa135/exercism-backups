package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	var (
		fancyNumConv int
		err          error
	)

	if fancyNum, ok := fnb.(FancyNumber); ok {
		if fancyNumConv, err = strconv.Atoi(fancyNum.n); err != nil {
			return 0
		}
	}

	return fancyNumConv
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf(
		"This is a fancy box containing the number %d.0",
		ExtractFancyNumber(fnb),
	)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch toDescribe := i.(type) {
	case float64:
		return DescribeNumber(toDescribe)
	case int:
		return DescribeNumber(float64(toDescribe))
	case NumberBox:
		return DescribeNumberBox(toDescribe)
	case FancyNumberBox:
		return DescribeFancyNumberBox(toDescribe)
	default:
		return "Return to sender"
	}
}
