package main

import (
	"fmt"
	"time"
)

func ParseStringToTime(dateString, format string) (time.Time, error) {
	res, ok := time.Parse(format, dateString)
	if ok != nil {
		return time.Time{}, fmt.Errorf("error when translated time")
	}
	return res, nil
}
