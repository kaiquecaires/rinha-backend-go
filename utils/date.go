package utils

import "time"

func IsDateValid(dateString string) bool {
	// Define the expected date format
	dateFormat := "2006-01-02" // "YYYY-MM-DD"

	// Attempt to parse the date string
	_, err := time.Parse(dateFormat, dateString)
	return err == nil
}
