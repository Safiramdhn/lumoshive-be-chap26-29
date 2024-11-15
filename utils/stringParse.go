package utils

import "strings"

func ParseActivities(data []byte) []string {
	// Convert byte slice to string and remove curly braces
	str := strings.Trim(string(data), "{}")

	// Split the string by commas to get individual items
	activities := strings.Split(str, ",")

	// Remove any extra spaces and surrounding quotes from each item
	for i, activity := range activities {
		activities[i] = strings.TrimSpace(activity)
		activities[i] = strings.Trim(activities[i], "\"") // Remove surrounding quotes
	}
	return activities
}
