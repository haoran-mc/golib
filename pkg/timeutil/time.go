package timeutil

import "time"

func NowInMilliseconds() int64 {
	return time.Now().UnixMilli()
}

func NowInMicroseconds() int64 {
	return time.Now().UnixMicro()
}

func NowInSeconds() int64 {
	return time.Now().Unix()
}

const (
	timeFormatStr     = "15:04:05"
	datetimeFormatStr = "2006-01-02 15:04:05"
)

func ToTime(s string) (time.Time, error) {
	return time.Parse(timeFormatStr, s)
}

func ToTimeString(t time.Time) string {
	return t.Format(timeFormatStr)
}

func BetweenTime(interval [2]time.Time, t time.Time) bool {
	return t.After(interval[0]) && t.Before(interval[1])
}

func ToDatetime(s string) (time.Time, error) {
	return time.Parse(datetimeFormatStr, s)
}

func ToDatetimeString(t time.Time) string {
	return t.Format(datetimeFormatStr)
}

func BetweenDatetime(interval [2]time.Time, t time.Time) bool {
	return t.After(interval[0]) && t.Before(interval[1])
}
