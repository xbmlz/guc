package utime

import (
	"fmt"
	"time"
)

var formatMap = map[string]string{
	"yyyy-MM-dd HH:mm:ss":          "2006-01-02 15:04:05",
	"yyyy/MM/dd HH:mm:ss":          "2006/01/02 15:04:05",
	"yyyy.MM.dd HH:mm:ss":          "2006.01.02 15:04:05",
	"yyyy年MM月dd日 HH时mm分ss秒":        "2006年01月02日 15时04分05秒",
	"yyyy-MM-dd":                   "2006-01-02",
	"yyyy/MM/dd":                   "2006/01/02",
	"yyyy.MM.dd":                   "2006.01.02",
	"HH:mm:ss":                     "15:04:05",
	"HH时mm分ss秒":                    "15时04分05秒",
	"yyyy-MM-dd HH:mm":             "2006-01-02 15:04",
	"yyyy-MM-dd HH:mm:ss.SSS":      "2006-01-02 15:04:05.000",
	"yyyyMMddHHmmss":               "20060102150405",
	"yyyyMMddHHmmssSSS":            "20060102150405000",
	"yyyyMMdd":                     "20060102",
	"EEE, dd MMM yyyy HH:mm:ss z":  "Mon, 02 Jan 2006 15:04:05 MST",
	"EEE MMM dd HH:mm:ss zzz yyyy": "Mon Jan 02 15:04:05 -0700 2006",
	"yyyy-MM-dd'T'HH:mm:ss'Z'":     "2006-01-02T15:04:05Z",
	"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'": "2006-01-02T15:04:05.000Z",
	"yyyy-MM-dd'T'HH:mm:ssZ":       "2006-01-02T15:04:05-0700",
	"yyyy-MM-dd'T'HH:mm:ss.SSSZ":   "2006-01-02T15:04:05.000-0700",
}

// Format formats a time value using the given layout.
func Format(t time.Time, layout string) string {
	if layout, ok := formatMap[layout]; ok {
		return t.Format(layout)
	}
	return t.Format(layout)
}

// Parse parses a formatted string and returns the time value it represents.
func Parse(value, layout string) (time.Time, error) {
	if layout, ok := formatMap[layout]; ok {
		return time.Parse(layout, value)
	}
	return time.Parse(layout, value)
}

// Now returns the current local time.
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

type OffsetType int

const (
	OffsetYear OffsetType = iota
	OffsetMonth
	OffsetDay
	OffsetHour
	OffsetMinute
	OffsetSecond
)

// Offset returns the time t+d.
func Offset(t time.Time, offset OffsetType, num int) time.Time {
	switch offset {
	case OffsetYear:
		return t.AddDate(num, 0, 0)
	case OffsetMonth:
		return t.AddDate(0, num, 0)
	case OffsetDay:
		return t.AddDate(0, 0, num)
	case OffsetHour:
		return t.Add(time.Hour * time.Duration(num))
	case OffsetMinute:
		return t.Add(time.Minute * time.Duration(num))
	case OffsetSecond:
		return t.Add(time.Second * time.Duration(num))
	default:
		return t
	}
}

type BetweenType int

const (
	BetweenYear BetweenType = iota
	BetweenMonth
	BetweenDay
	BetweenHour
	BetweenMinute
	BetweenSecond
	BetweenMillisecond
)

// Between returns the time t+d.
func Between(t1, t2 time.Time, between BetweenType) int64 {
	duration := t2.Sub(t1)
	switch between {
	case BetweenYear:
		return int64(duration.Hours() / 24 / 365)
	case BetweenMonth:
		return int64(duration.Hours() / 24 / 30)
	case BetweenDay:
		return int64(duration.Hours() / 24)
	case BetweenHour:
		return int64(duration.Hours())
	case BetweenMinute:
		return int64(duration.Minutes())
	case BetweenSecond:
		return int64(duration.Seconds())
	case BetweenMillisecond:
		return int64(duration.Milliseconds())
	default:
		return 0
	}
}

type FmtBetweenType int

const (
	FmtBetweenDay FmtBetweenType = iota
	FmtBetweenHour
	FmtBetweenMinute
	FmtBetweenSecond
	FmtBetweenMillisecond
)

var FmtBetweenNames = map[FmtBetweenType]string{
	FmtBetweenDay:         "天",
	FmtBetweenHour:        "小时",
	FmtBetweenMinute:      "分钟",
	FmtBetweenSecond:      "秒",
	FmtBetweenMillisecond: "毫秒",
}

// FmtBetween returns the time t+d.
func FmtBetween(t1, t2 time.Time, level FmtBetweenType) (ret string) {
	betweenMs := Between(t1, t2, BetweenMillisecond)
	if betweenMs > 0 {
		day := betweenMs / 1000 / 60 / 60 / 24
		hour := betweenMs / 1000 / 60 / 60 % 24
		minute := betweenMs / 1000 / 60 % 60

		betweenOfSecond := ((day*24+hour)*60 + minute) * 60
		second := betweenMs/1000 - betweenOfSecond

		millisecond := betweenMs - (betweenOfSecond+second)*1000

		levelCount := 0
		if day != 0 && level >= FmtBetweenDay {
			ret += fmt.Sprintf("%d%s", day, FmtBetweenNames[FmtBetweenDay])
			levelCount++
		}
		if hour != 0 && level >= FmtBetweenHour {
			ret += fmt.Sprintf("%d%s", hour, FmtBetweenNames[FmtBetweenHour])
			levelCount++
		}
		if minute != 0 && level >= FmtBetweenMinute {
			ret += fmt.Sprintf("%d%s", minute, FmtBetweenNames[FmtBetweenMinute])
			levelCount++
		}
		if second != 0 && level >= FmtBetweenSecond {
			ret += fmt.Sprintf("%d%s", second, FmtBetweenNames[FmtBetweenSecond])
			levelCount++
		}
		if millisecond != 0 && level >= FmtBetweenMillisecond {
			ret += fmt.Sprintf("%d%s", millisecond, FmtBetweenNames[FmtBetweenMillisecond])
		}
	}
	if ret == "" {
		ret = "0" + FmtBetweenNames[level]
	}
	return ret
}
