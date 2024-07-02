package dateutils

import (
	"strings"
	"time"
)

func GetDate(dateValue string) time.Time {
	const dateFormat = "2006-01-02"

	d := strings.Split(dateValue, "/")

	if len(d[0]) < 2 {
		d[0] = "0" + d[0]
	}

	if len(d[1]) < 2 {
		d[1] = "0" + d[1]
	}

	year := d[2]
	d[2] = d[0]
	d[0] = year

	newDateText := strings.Join(d, "-")
	newDate, _ := time.Parse(dateFormat, newDateText)
	return newDate
}
