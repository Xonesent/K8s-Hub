package utils

import (
	"errors"
	"time"
)

var errScheduleTimeBeforeNow = errors.New("current time less than specified")

func ValidateTimer(parsedTime time.Time) (time.Duration, error) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return -1, err
	}

	now := time.Now().In(loc)
	scheduledTime := time.Date(now.Year(), now.Month(), now.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, now.Location())

	if scheduledTime.Before(now) {
		return -1, errScheduleTimeBeforeNow
	}

	return scheduledTime.Sub(now), nil
}
