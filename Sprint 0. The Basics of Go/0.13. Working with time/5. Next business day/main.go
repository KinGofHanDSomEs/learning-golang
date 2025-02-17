package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	i := 0
	for {
		i++
		next_day := start.AddDate(0, 0, i)
		if next_day.Format("Monday") != "Saturday" && next_day.Format("Monday") != "Sunday" {
			return next_day
		}
	}
}
