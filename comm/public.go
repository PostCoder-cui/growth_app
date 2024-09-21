package comm

import (
	"fmt"
	"runtime"
	"time"
)

const Time_Layout = "2006-01-02 15:04:05"

var Empty_Time time.Time

func Now() *time.Time {
	now := time.Now()
	return &now
}

func MarkLine() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("打个标记 at %s:%d", file, line)
}
func MarkLineErr(err any) string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("发生错误%v at %s:%d", err, file, line)
}
func TimeFormat(time *time.Time) string {
	return time.Format(Time_Layout)
}
func TimeParse(str string) time.Time {
	if res, err := time.Parse(Time_Layout, str); err != nil {
		return time.Time{}
	} else {
		return res
	}
}
