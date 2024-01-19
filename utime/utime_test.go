package utime

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name   string
		t      time.Time
		layout string
		want   string
	}{
		{"yyyy-MM-dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd HH:mm:ss", "2019-01-01 01:01:01"},
		{"yyyy/MM/dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy/MM/dd HH:mm:ss", "2019/01/01 01:01:01"},
		{"yyyy.MM.dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy.MM.dd HH:mm:ss", "2019.01.01 01:01:01"},
		{"yyyy年MM月dd日 HH时mm分ss秒", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy年MM月dd日 HH时mm分ss秒", "2019年01月01日 01时01分01秒"},
		{"yyyy-MM-dd", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd", "2019-01-01"},
		{"yyyy/MM/dd", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy/MM/dd", "2019/01/01"},
		{"yyyy.MM.dd", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy.MM.dd", "2019.01.01"},
		{"HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "HH:mm:ss", "01:01:01"},
		{"HH时mm分ss秒", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "HH时mm分ss秒", "01时01分01秒"},
		{"yyyy-MM-dd HH:mm", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd HH:mm", "2019-01-01 01:01"},
		{"yyyy-MM-dd HH:mm:ss.SSS", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd HH:mm:ss.SSS", "2019-01-01 01:01:01.000"},
		{"yyyyMMddHHmmss", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyyMMddHHmmss", "20190101010101"},
		{"yyyyMMddHHmmssSSS", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyyMMddHHmmssSSS", "20190101010101000"},
		{"yyyyMMdd", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyyMMdd", "20190101"},
		{"EEE, dd MMM yyyy HH:mm:ss z", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "EEE, dd MMM yyyy HH:mm:ss z", "Tue, 01 Jan 2019 01:01:01 CST"},
		{"EEE MMM dd HH:mm:ss zzz yyyy", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "EEE MMM dd HH:mm:ss zzz yyyy", "Tue Jan 01 01:01:01 +0800 2019"},
		{"yyyy-MM-dd'T'HH:mm:ss'Z'", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd'T'HH:mm:ss'Z'", "2019-01-01T01:01:01Z"},
		{"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'", "2019-01-01T01:01:01.000Z"},
		{"yyyy-MM-dd'T'HH:mm:ssZ", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd'T'HH:mm:ssZ", "2019-01-01T01:01:01+0800"},
		{"yyyy-MM-dd'T'HH:mm:ss.SSSZ", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), "yyyy-MM-dd'T'HH:mm:ss.SSSZ", "2019-01-01T01:01:01.000+0800"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.t, tt.layout); got != tt.want {
				t.Errorf("Format(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		layout string
		want   time.Time
	}{
		// name yyyy-MM-dd HH:mm:ss , got 2019-01-01 01:01:01 +0000 UTC, want 2019-01-01 01:01:01.000000001 +0800 CST
		{"yyyy-MM-dd HH:mm:ss", "2019-01-01 01:01:01", "yyyy-MM-dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy/MM/dd HH:mm:ss", "2019/01/01 01:01:01", "yyyy/MM/dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy.MM.dd HH:mm:ss", "2019.01.01 01:01:01", "yyyy.MM.dd HH:mm:ss", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy年MM月dd日 HH时mm分ss秒", "2019年01月01日 01时01分01秒", "yyyy年MM月dd日 HH时mm分ss秒", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy-MM-dd", "2019-01-01", "yyyy-MM-dd", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"yyyy/MM/dd", "2019/01/01", "yyyy/MM/dd", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"yyyy.MM.dd", "2019.01.01", "yyyy.MM.dd", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"HH:mm:ss", "01:01:01", "HH:mm:ss", time.Date(0, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"HH时mm分ss秒", "01时01分01秒", "HH时mm分ss秒", time.Date(0, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy-MM-dd HH:mm", "2019-01-01 01:01", "yyyy-MM-dd HH:mm", time.Date(2019, 1, 1, 1, 1, 0, 0, time.UTC)},
		{"yyyy-MM-dd HH:mm:ss.SSS", "2019-01-01 01:01:01.000", "yyyy-MM-dd HH:mm:ss.SSS", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyyMMddHHmmss", "20190101010101", "yyyyMMddHHmmss", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyyMMddHHmmssSSS", "20190101010101000", "yyyyMMddHHmmssSSS", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyyMMdd", "20190101", "yyyyMMdd", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"EEE, dd MMM yyyy HH:mm:ss z", "Tue, 01 Jan 2019 01:01:01 CST", "EEE, dd MMM yyyy HH:mm:ss z", time.Date(2019, 1, 1, 1, 1, 1, 0, time.Local)},
		{"EEE MMM dd HH:mm:ss zzz yyyy", "Tue Jan 01 01:01:01 +0800 2019", "EEE MMM dd HH:mm:ss zzz yyyy", time.Date(2019, 1, 1, 1, 1, 1, 0, time.Local)},
		{"yyyy-MM-dd'T'HH:mm:ss'Z'", "2019-01-01T01:01:01Z", "yyyy-MM-dd'T'HH:mm:ss'Z'", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'", "2019-01-01T01:01:01.000Z", "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'", time.Date(2019, 1, 1, 1, 1, 1, 0, time.UTC)},
		{"yyyy-MM-dd'T'HH:mm:ssZ", "2019-01-01T01:01:01+0800", "yyyy-MM-dd'T'HH:mm:ssZ", time.Date(2019, 1, 1, 1, 1, 1, 0, time.Local)},
		{"yyyy-MM-dd'T'HH:mm:ss.SSSZ", "2019-01-01T01:01:01.000+0800", "yyyy-MM-dd'T'HH:mm:ss.SSSZ", time.Date(2019, 1, 1, 1, 1, 1, 0, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Parse(tt.value, tt.layout); got != tt.want || err != nil {
				t.Errorf("Parse(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestOffset(t *testing.T) {
	tests := []struct {
		name   string
		t      time.Time
		offset OffsetType
		num    int
		want   time.Time
	}{
		{"OffsetYear", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetYear, 1, time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local)},
		{"OffsetMonth", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetMonth, 1, time.Date(2019, 2, 1, 1, 1, 1, 1, time.Local)},
		{"OffsetDay", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetDay, 1, time.Date(2019, 1, 2, 1, 1, 1, 1, time.Local)},
		{"OffsetHour", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetHour, 1, time.Date(2019, 1, 1, 2, 1, 1, 1, time.Local)},
		{"OffsetMinute", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetMinute, 1, time.Date(2019, 1, 1, 1, 2, 1, 1, time.Local)},
		{"OffsetSecond", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetSecond, 1, time.Date(2019, 1, 1, 1, 1, 2, 1, time.Local)},
		{"-OffsetYear", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetYear, -1, time.Date(2018, 1, 1, 1, 1, 1, 1, time.Local)},
		{"-OffsetMonth", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetMonth, -1, time.Date(2018, 12, 1, 1, 1, 1, 1, time.Local)},
		{"-OffsetDay", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetDay, -1, time.Date(2018, 12, 31, 1, 1, 1, 1, time.Local)},
		{"-OffsetHour", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetHour, -1, time.Date(2019, 1, 1, 0, 1, 1, 1, time.Local)},
		{"-OffsetMinute", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetMinute, -1, time.Date(2019, 1, 1, 1, 0, 1, 1, time.Local)},
		{"-OffsetSecond", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), OffsetSecond, -1, time.Date(2019, 1, 1, 1, 1, 0, 1, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Offset(tt.t, tt.offset, tt.num); got != tt.want {
				t.Errorf("Offset(): name %v , got %v, want %v", tt.name, got, tt.want)

			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		b    BetweenType
		want int64
	}{
		{"BetweenYear", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2020, 1, 1, 1, 1, 1, 1, time.Local), BetweenYear, 1},
		{"BetweenMonth", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 2, 1, 1, 1, 1, 1, time.Local), BetweenMonth, 1},
		{"BetweenDay", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 2, 1, 1, 1, 1, time.Local), BetweenDay, 1},
		{"BetweenHour", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 2, 1, 1, 1, time.Local), BetweenHour, 1},
		{"BetweenMinute", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 1, 2, 1, 1, time.Local), BetweenMinute, 1},
		{"BetweenSecond", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 1, 1, 2, 1, time.Local), BetweenSecond, 1},
		{"BetweenYear", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2018, 1, 1, 1, 1, 1, 1, time.Local), BetweenYear, -1},
		{"BetweenMonth", time.Date(2019, 3, 1, 1, 1, 1, 1, time.Local), time.Date(2018, 12, 1, 1, 1, 1, 1, time.Local), BetweenMonth, -3},
		{"BetweenDay", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2018, 12, 31, 1, 1, 1, 1, time.Local), BetweenDay, -1},
		{"BetweenHour", time.Date(2018, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 0, 1, 1, 1, time.Local), BetweenHour, 8759},
		{"BetweenMinute", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 1, 0, 1, 1, time.Local), BetweenMinute, -1},
		{"BetweenSecond", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2018, 1, 1, 1, 1, 1, 1, time.Local), BetweenSecond, -31536000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.t1, tt.t2, tt.b); got != tt.want {
				t.Errorf("Between(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFmtBetween(t *testing.T) {
	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		tt   FmtBetweenType
		want string
	}{
		{"1天", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 2, 1, 1, 1, 1, time.Local), FmtBetweenDay, "1天"},
		{"1小时", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 2, 1, 1, 1, time.Local), FmtBetweenHour, "1小时"},
		{"1分钟", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 1, 2, 1, 1, time.Local), FmtBetweenMinute, "1分钟"},
		{"1秒", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 1, 1, 1, 1, 2, 1, time.Local), FmtBetweenSecond, "1秒"},
		{"31天1小时", time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local), time.Date(2019, 2, 1, 2, 1, 1, 1, time.Local), FmtBetweenMinute, "31天1小时"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FmtBetween(tt.t1, tt.t2, tt.tt); got != tt.want {
				t.Errorf("FmtBetween(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
