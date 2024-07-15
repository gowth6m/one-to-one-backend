package one_to_one

import (
	"time"
)

func GetCurrentWeekAndYear() (int, int) {
	now := time.Now()
	_, week := now.ISOWeek()
	return week, now.Year()
}
