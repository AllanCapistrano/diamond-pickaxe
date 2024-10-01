package handler

import "time"

// Return the current date and time using the following format:
// d/m/Y H:m:s
func CurrentTimestampFormatted() string {
	now := time.Now()

	formattedTimestamp := now.Format("02/01/2006 15:04:05")

	return formattedTimestamp
}
