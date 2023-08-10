package utils

import (
	"log"
	"time"
)

func ParseTime(s string) (dt time.Time) {
	// layout must be at 2006/01/02 15:04:05 (MST), with any format
	layout := "2006-01-02T15:04"
	dt, err := time.Parse(layout, s)

	if err != nil {
		log.Println(err)
	}

	return dt
}
