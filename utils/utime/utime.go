package utime

import (
	"database/sql/driver"
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
// 注（x-stock 2026-08-31 修复）：按 Local 时区解析——原 time.Parse 按 UTC 解析，
// 与 Now()/DB(loc=Local) 差 8 小时，日期边界比较（回测 from/to、建仓日）会被
// 钳到次日（详见 x-stock rewrite-five-modules 5.9 验收发现的时区偏移 bug）。
func New(s string, layout ...string) *Time {
	if len(layout) == 0 {
		layout = []string{time.DateTime}
	}
	t, _ := time.ParseInLocation(layout[0], s, time.Local)
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

// MarshalJSON 序列化为 unix 毫秒 JSON 数字（值接收者，值/指针字段输出一致）。
// 原实现返回裸 8 字节二进制（非合法 JSON），导致 json.Marshal 整体失败，此处修正。
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.UnixMilli(), 10)), nil
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
	if v == nil {
		*t = Time{} // NULL → 零值

		return nil
	}

	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to time.Time", v)
	}
	*t = Time{Time: value}

	return nil
}

// Value sql.driver.Valuer。
// 用值接收者：gorm 模型结构体字段为 utime.Time 值类型时，driver 要求值类型实现 Valuer
// （指针接收者版本在 gorm 写入时会报 unsupported type utime.Time, a struct）。
// 零值返回 nil（NULL）：零值 time.Time 会被驱动写成 '0000-00-00'，MySQL 8 默认 sql_mode 拒绝。
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}
