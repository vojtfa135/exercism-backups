package booking

import (
    "time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	time, _ := time.Parse(layout, date)
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	timeDate, _ := time.Parse(layout, date)
	return timeDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	timeDate, _ := time.Parse(layout, date)

	if appointHour := timeDate.Hour(); appointHour >= 12 && appointHour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	timeDate, _ := time.Parse(layout, date)
	appointWeekDay := timeDate.Weekday().String()
	appointMonth := timeDate.Month().String()
	appointDay := timeDate.Day()
	appointYear := timeDate.Year()
	appointHour := timeDate.Hour()
	appointMin := timeDate.Minute()

	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %d:%d.",
		appointWeekDay,
		appointMonth,
		appointDay,
		appointYear,
		appointHour,
		appointMin,
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05 +0000 UTC"
	thisYearAnns := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", time.Now().Year())
	thisYearAnsDate, _ := time.Parse(layout, thisYearAnns)
	return thisYearAnsDate
}
