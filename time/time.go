package time

import (
	"math"
	"strconv"
	"time"
)

var (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, time.Local)
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowDate() string {
	return time.Now().Format(TimeLayoutDate)
}

func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}

func IsValidTime(t time.Time) bool {
	if t.IsZero() {
		return false
	}
	if t.Unix() <= 0 {
		return false
	}
	return true
}

// SinceForHuman 1小时前 -> 这样的展示方式
func SinceForHuman(t time.Time) string {
	duration := time.Since(t)
	hour := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()

	unit := "秒"
	s := 0
	if hour > (365 * 24) {
		s = int(math.Floor(hour / 365))
		unit = "年"
	} else if hour > 30 {
		s = int(math.Floor(hour / 30))
		unit = "月"
	} else if hour > 1 {
		s = int(math.Floor(hour))
		unit = "小时"
	} else if minutes > 1 {
		s = int(math.Floor(minutes))
		unit = "分钟"
	} else if seconds > 0 {
		return "刚刚"
	}

	return strconv.Itoa(s) + unit + "前"
}
