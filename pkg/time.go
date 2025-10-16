package pkg

import "time"

// Function to abstract timestamp logic
// It returns the time at current time in the RFC3339 format
func Timestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
