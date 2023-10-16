package utils

import "time"

func ParseStringDate(dateString string, local string) (time.Time, error) {
	tz, _ := time.LoadLocation(local)
	date, err := time.ParseInLocation("01/02 15:04", dateString, tz)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func ConvertDateInput(dateString string, local string) (time.Time, error) {
	date, err := ParseStringDate(dateString, local)
	if err != nil {
		return time.Time{}, err
	}
	currentDate := time.Now()
	date = date.AddDate(currentDate.Year(), 0, 0)
	if date.Before(currentDate) {
		date = date.AddDate(1, 0, 0)
	}
	return date, nil
}
