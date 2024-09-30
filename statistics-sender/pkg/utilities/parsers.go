package utils

import (
	"fmt"
	"time"
)

func ParseTimeStr(timeStr string) (time.Time, error) {
	layout := "15:04:05"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %w", err)
	}

	return t, nil
}
