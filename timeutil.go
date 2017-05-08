package main

import (
	"log"
	"math"
	"strconv"
	"time"
)

func formatScheduleTime(zone string, scheduledTime string) string {
	loc, _ := time.LoadLocation(zone)
	t, err := time.ParseInLocation(time.RFC822Z, scheduledTime, loc)
	if checkErr("timeutil.go: ", err) {
		return ""
	}

	log.Println("default format:", t)
	log.Println("Unix format:", t.Format(time.UnixDate))
	utcT := t.UTC()
	log.Println("Same, in UTC:", utcT.Format(time.UnixDate))

	h, m := utcT.Hour(), utcT.Minute()
	hStr := strconv.Itoa(h)
	mStr := strconv.Itoa(m)
	if h < 10 {
		hStr = "0" + hStr
	}
	if m < 10 {
		mStr = "0" + mStr
	}
	return hStr + ":" + mStr
}

func timeAgo(then int64) string {
	now := time.Now().Unix()
	seconds := int(math.Abs(float64(now - then)))

	var suffix string
	if then < now {
		suffix = "ago"
	} else {
		suffix = "from now"
	}

	var value int
	var unit string

	if seconds < 60 {
		value = seconds
		unit = "second"
	} else if seconds < 60*60 {
		value = seconds / 60
		unit = "minute"
	} else if seconds < 60*60*24 {
		value = (seconds / (60 * 60))
		unit = "hour"
	} else if seconds < 60*60*24*7 {
		value = (seconds / (60 * 60 * 24))
		unit = "day"
	} else if seconds < 60*60*24*30 {
		value = (seconds / (60 * 60 * 24 * 7))
		unit = "week"
	} else if seconds < 60*60*24*365 {
		value = (seconds / (60 * 60 * 24 * 30))
		unit = "month"
	} else {
		value = (seconds / (60 * 60 * 24 * 365))
		unit = "year"
	}

	if value != 1 {
		unit += "s"
	}

	return strconv.Itoa(value) + " " + unit + " " + suffix
}
