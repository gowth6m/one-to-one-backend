package one_to_one

import (
	"time"
)

func GetCurrentWeekAndYear() (int, int) {
	now := time.Now()
	_, week := now.ISOWeek()
	return week, now.Year()
}

// Helper function to filter out empty strings
func FilterEmptyLabels[T any](items []T, getLabel func(T) string) []T {
	var filteredItems []T
	for _, item := range items {
		if label := getLabel(item); label != "" {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
