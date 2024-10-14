package utils

import (
	"time"
)

// check if targetTime is between startTime and endTime
func IsTimeBetween(targetTime, startTime, endTime time.Time) (bool, error) {
	// Check if startTime is zero value (optional startTime check)
	if startTime.IsZero() && !endTime.IsZero() {
		// Check if targetTime is before endTime
		if targetTime.Before(endTime) {
			return true, nil
		}
		return false, nil
	}

	// Check if endTime is zero value (optional endTime check)
	if endTime.IsZero() && !startTime.IsZero() {
		// Check if targetTime is after startTime
		if targetTime.After(startTime) {
			return true, nil
		}
		return false, nil
	}

	// If both startTime and endTime are provided, check if targetTime is between them
	if targetTime.After(startTime) && targetTime.Before(endTime) {
		return true, nil
	}

	return false, nil
}