package expenses

import (
	"fmt"
)

type UnknownCategoryError struct {
	category string
}

func (e *UnknownCategoryError) Error() string {
	return fmt.Sprintf("unknown category %s", e.category)
}

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	if predicate == nil {
		return in
	}

	filteredRecord := []Record{}

	for _, val := range in {
		if predicate(val) {
			filteredRecord = append(filteredRecord, val)
		}
	}

	return filteredRecord
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.From <= r.Day && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var periodTotal float64

	for _, val := range in {
		if ByDaysPeriod(p)(val) {
			periodTotal += val.Amount
		}
	}

	return periodTotal
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var (
		periodTotal float64
		numOfFails  int
		err         error
	)

	for _, val := range in {
		periodMatch := ByDaysPeriod(p)(val)
		categoryMatch := ByCategory(c)(val)

		if !categoryMatch {
			numOfFails++
		}

		if periodMatch && categoryMatch {
			periodTotal += val.Amount
		}
	}

	if numOfFails == len(in) {
		err = &UnknownCategoryError{
			category: c,
		}
	}

	return periodTotal, err
}
