/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2023/1/11 16:59
 */

package test

import (
	"fmt"
	"testing"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func TestRun(t *testing.T) {
	LastTime := "2022-12-31 00:00:00"
	now, _ := time.ParseInLocation(timeFormat, LastTime, time.Local)
	dh, _ := time.ParseDuration(fmt.Sprintf("+%dh", 24))
	expTime := now.Add(dh).Format(timeFormat)
	fmt.Println(expTime)
	fmt.Println(time.Duration(18000000))

	m, _ := time.ParseDuration("30m")
	fmt.Println(time.Duration(m.Seconds() * float64(time.Second)))
}
