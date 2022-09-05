package main

import (
	"fmt"
	"time"
)

func main() {
	timeLayout := "2006-01-02"

	var dfTimeValue time.Time

	timeString := "2022-5-01"

	t := parseTime(timeString, dfTimeValue, timeLayout)

	fmt.Println(t)
}

func parseTime(ts string, defaultValue time.Time, layout string) time.Time {
	if ts == "" {
		return defaultValue
	}

	t, err := time.Parse(layout, ts)
	if err != nil {
		return defaultValue
	}

	return t
}
