package helpers

import (
	"go-boilerplate-api/internal/api/config"
	"time"
)

func GetTodayDate() string {
	loc, err := time.LoadLocation(config.TIMEZONE)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	today := time.Now().In(loc)
	return today.Format("2006-01-02")
}

func IsDateToday(date string) bool {
	loc, err := time.LoadLocation(config.TIMEZONE)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	todayRaw := time.Now().In(loc)
	today := todayRaw.Format("2006-01-02")
	return date == today
}
