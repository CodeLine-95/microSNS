package LocalTime

import (
	"database/sql/driver"
	"fmt"
	"github.com/jobhandsome/microSNS/common/define"
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type LocalTime struct {
	time.Time
}

// MarshalJSON 获取当前时间 返回时间戳
func (t LocalTime) MarshalJSON() []byte {
	second := t.Unix()
	return []byte(strconv.FormatInt(second, 10))
}

// DateTimeToStamp 任意日期格式转换成 时间戳
func (t *LocalTime) DateTimeToStamp(data string) int64 {
	now, err := time.ParseInLocation(timeFormat, data, time.Local)
	if err != nil {
		return time.Now().Unix()
	}
	return now.Unix()
}
func (t LocalTime) Value() driver.Value {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil
	}
	return t.Time
}

// 当前时间转换成日期格式
func (t LocalTime) String() string {
	return time.Now().Format(timeFormat)
}

// FormatString 当前时间转换成日期格式
func (t LocalTime) FormatString(Format string) string {
	return time.Now().Format(Format)
}

// 1. 设置时区
// 2. time.LoadLocation() 返回time.Time 类型
func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc)
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// SpecifiedDate 获取指定 年、月、日
func (t LocalTime) SpecifiedDate(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format(define.DateDayFormat)
}

// SpecifiedTimeForDayAfter 获取指定日期的后一天
func (t LocalTime) SpecifiedTimeForDayAfter(LastTime string) string {
	now, _ := time.ParseInLocation(timeFormat, LastTime, time.Local)
	dh, _ := time.ParseDuration(fmt.Sprintf("+%dh", 24))
	return now.Add(dh).Format(timeFormat)
}

// SpecifiedTimeForDayBefore 获取指定日期的前一天
func (t LocalTime) SpecifiedTimeForDayBefore(LastTime string) string {
	now, _ := time.ParseInLocation(timeFormat, LastTime, time.Local)
	dh, _ := time.ParseDuration(fmt.Sprintf("-%dh", 24))
	return now.Add(dh).Format(timeFormat)
}

// FormatDateString 将指定格式的日期，转成 YYYY-MM-DD HH:ii:ss
func (t LocalTime) FormatDateString(MyDateString string) string {
	// 1. 设置时区
	// 2. time.LoadLocation() 返回time.Time 类型
	loc, _ := time.LoadLocation(timezone)
	// timeFormat = 2006-01-02 15:04:05  go的默认日期
	// 将日期字符串解析为Go的time对象第一个参数指定格式，第二个是日期字符串，转换成时间戳(秒)
	myDate, err := time.ParseInLocation(timeFormat, MyDateString, loc)
	if err != nil {
		return ""
	}
	// 获取时间戳
	// 1.把时间戳转使用time.Unix()转化为 time.Time 类型时间
	// 2.使用Format() 方法进行格式化
	DateString := time.Unix(myDate.Unix(), 0).Format(timeFormat)
	return DateString
}
