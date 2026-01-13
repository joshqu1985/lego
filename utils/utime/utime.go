package utime

import (
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Time time.
type Time struct {
	time.Time
}

// Now 基于当前时间构建Time.
func Now() *Time {
	return &Time{
		Time: time.Now(),
	}
}

// New 基于时间字符串构建Time.
func New(s string, layout ...string) *Time {
	if len(layout) == 0 {
		layout = []string{time.DateTime}
	}
	t, _ := time.Parse(layout[0], s)
	if t.IsZero() {
		return nil
	}

	return &Time{t}
}

// String 转换成字符串格式.
func (t *Time) String(layout ...string) string {
	if t == nil || t.IsZero() {
		return ""
	}
	if len(layout) == 0 {
		layout = []string{time.DateTime}
	}

	return t.Format(layout[0])
}

// InNumDays 是否在n天内.
func (t *Time) InNumDays(n int) bool {
	begYear, begMonth, begDay := t.Date()
	beg := time.Date(begYear, begMonth, begDay, 0, 0, 0, 0, t.Location())

	endYear, endMonth, endDay := time.Now().Date()
	end := time.Date(endYear, endMonth, endDay, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())

	return beg.AddDate(0, 0, n).After(end)
}

// DayBegin 天开始时间.
func (t *Time) DayBegin() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, m, d := t.Date()

	return &Time{time.Date(y, m, d, 0, 0, 0, 0, t.Location())}
}

// DayEnd 天结束时间.
func (t *Time) DayEnd() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, m, d := t.Date()

	return &Time{time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())}
}

// WeekBegin 周开始时间 NOTICE: 一周从周日开始.
func (t *Time) WeekBegin() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, m, d := t.AddDate(0, 0, 0-int(t.DayBegin().Weekday())).Date()

	return &Time{time.Date(y, m, d, 0, 0, 0, 0, t.Location())}
}

// WeekEnd 周结束时间 NOTICE: 一周从周日开始.
func (t *Time) WeekEnd() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, m, d := t.WeekBegin().AddDate(0, 0, 7).Add(-time.Nanosecond).Date()

	return &Time{time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())}
}

// MonthBegin 月开始时间.
func (t *Time) MonthBegin() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, m, _ := t.Date()

	return &Time{time.Date(y, m, 1, 0, 0, 0, 0, t.Location())}
}

// MonthEnd 月结束时间.
func (t *Time) MonthEnd() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}

	return &Time{t.MonthBegin().AddDate(0, 1, 0).Add(-time.Nanosecond)}
}

// YearBegin 年开始时间.
func (t *Time) YearBegin() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}
	y, _, _ := t.Date()

	return &Time{time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())}
}

// YearEnd 年结束时间.
func (t *Time) YearEnd() *Time {
	if t == nil || t.IsZero() {
		return &Time{}
	}

	return &Time{t.YearBegin().AddDate(1, 0, 0).Add(-time.Nanosecond)}
}

// MarshalJSON 序列化.
func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("0"), nil
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t.UnixMilli()))

	return buf, nil
}

// UnmarshalJSON 反序列化.
func (t *Time) UnmarshalJSON(bytes []byte) error {
	if t == nil || len(bytes) == 0 {
		return errors.New("time or args is nil")
	}
	if bytes[0] == bytes[len(bytes)-1] && bytes[0] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}

	intval, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(intval/1e3, intval%1e3*1e6)

	return nil
}

// Scan sql.Scanner.
func (t *Time) Scan(v any) error {
	if t == nil {
		return errors.New("time is nil")
	}

	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to time.Time", v)
	}
	*t = Time{Time: value}

	return nil
}

// Value sql.driver.Valuer.
func (t *Time) Value() (driver.Value, error) {
	if t == nil || t.IsZero() {
		return time.Time{}, nil
	}

	return t.Time, nil
}
