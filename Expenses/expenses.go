package expenses

import "errors"

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
	res := make([]Record, 0)
	for _, v := range in {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	f := func(record Record) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		}
		return false
	}
	return f
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	f := func(record Record) bool {
		if record.Category == c {
			return true
		}
		return false
	}
	return f
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var tot float64 = 0
	for _, v := range in {
		if ByDaysPeriod(p)(v) {
			tot += v.Amount
		}
	}
	return tot
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	found := false
	var tot float64 = 0
	for _, v := range in {
		if v.Category == c {
			found = true
		}
		if ByDaysPeriod(p)(v) {
			if v.Category == c {
				tot += v.Amount
			}
		}
	}
	if tot == 0 && !found {
		return 0, errors.New("")
	}
	return tot, nil
}
