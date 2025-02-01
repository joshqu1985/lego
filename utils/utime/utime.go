package utime

import (
	"fmt"
	"time"

	"database/sql/driver"
)

// Time time
type Time struct {
	time.Time
}

// Now 基于当前时间构建Time
func Now() *Time {
	return &Time{
		Time: time.Now(),
	}
}

// New 基于时间字符串构建Time
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

// String 转换成字符串格式
func (t *Time) String(layout ...string) string {
	if t == nil || t.Time.IsZero() {
		return ""
	}

	if len(layout) == 0 {
		layout = []string{time.DateTime}
	}
	return t.Time.Format(layout[0])
}

// Parse 解析时间字符串
func (t *Time) Parse(s string, layout ...string) (err error) {
	if t == nil {
		return nil
	}

	if len(layout) == 0 {
		layout = []string{time.DateTime}
	}
	t.Time, err = time.Parse(layout[0], s)
	return err
}

// MarshalJSON 序列化
func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil || t.Time.IsZero() {
		return []byte(""), nil
	}
	data := fmt.Sprintf("\"%s\"", t.Time.Format(time.DateTime))
	return []byte(data), nil
}

// UnmarshalJSON 反序列化
func (t *Time) UnmarshalJSON(bytes []byte) (err error) {
	if len(bytes) == 0 {
		return nil
	}
	if t == nil {
		return fmt.Errorf("time is nil")
	}

	if bytes[0] == bytes[len(bytes)-1] && bytes[0] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}

	t.Time, err = time.Parse(time.DateTime, string(bytes))
	return err
}

// Scan sql.Scanner
func (t *Time) Scan(v interface{}) error {
	if t == nil {
		return fmt.Errorf("time is nil")
	}

	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to time.Time", v)
	}

	*t = Time{Time: value}
	return nil
}

// Value sql.driver.Valuer
func (t *Time) Value() (driver.Value, error) {
	if t == nil || t.Time.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}
