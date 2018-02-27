package timeformatutil

import (
	"strconv"
	"strings"
	"time"
)

// TimeEncoder 时间格式化器
// t 被格式化的时间，格式化为YYYYmmddhhiiss格式或YYYYmmddhhiissnn格式(到毫秒)
func TimeFormator(t time.Time) (string, string) {
	second := t.Unix()                        //秒，int64
	necond := t.UnixNano()                    //纳秒，int64 1s=10^9 ns
	mtime := int((necond - second*1e9) / 1e6) //毫秒
	h := strconv.Itoa(t.Hour())
	i := strconv.Itoa(t.Minute())

	s := strconv.Itoa(t.Second())
	n := strconv.Itoa(mtime)
	if len(h) == 1 {
		h = "0" + h
	}
	if len(i) == 1 {
		i = "0" + i
	}
	if len(s) == 1 {
		s = "0" + s
	}
	timeString := time.Now().Format("2006-01-02 15:04:05")
	ym := strings.Split(strings.Split(timeString, " ")[0], "-")
	rtn1 := ym[0] + ym[1] + ym[2] + h + i + s + n
	rtn2 := ym[0] + ym[1] + ym[2] + h + i + s
	return rtn1, rtn2
}

// TimeDecoder 反时间格式化
// 返回精确到秒或毫秒的时间字符串
func TimeDecoder(timeString string) string {
	var validTimeString, s string
	length := len(timeString)
	y := timeString[0:4]
	m := timeString[4:6]
	d := timeString[6:8]
	h := timeString[8:10]
	i := timeString[10:12]
	if length == 14 {
		s = timeString[12:14]
		validTimeString = y + "-" + m + "-" + d + " " + h + ":" + i + ":" + s
	} else {
		s = timeString[12:14]
		ms := timeString[14:]
		validTimeString = y + "-" + m + "-" + d + " " + h + ":" + i + ":" + s + "." + ms
	}
	return validTimeString
}

// WeekDateTimeInteger 截取年月日时分秒对应的整数
// t 指定的时间
// flag 截取的信息
// y、Y表示年，m、M表示月，d、D表示日，h、H表示时,i、I表示分，s、S表示秒，w、W表示周
func WeekDateTimeInteger(t time.Time, flag byte) int {
	str := t.Format("2006-01-02 15:04:05")
	str = strings.Replace(str, " ", "-", -1)
	str = strings.Replace(str, ":", "-", -1)
	dt := strings.Split(str, "-")

	var n int
	switch flag {
	case 'y', 'Y':
		n, _ = strconv.Atoi(dt[0])
	case 'm', 'M':
		n, _ = strconv.Atoi(dt[1])
	case 'd', 'D':
		n, _ = strconv.Atoi(dt[2])
	case 'h', 'H':
		n, _ = strconv.Atoi(dt[3])
	case 'i', 'I':
		n, _ = strconv.Atoi(dt[4])
	case 's', 'S':
		n, _ = strconv.Atoi(dt[5])
	case 'w', 'W':
		week := t.Weekday().String()
		switch week {
		case "Sunday":
			n = 0
		case "Monday":
			n = 1
		case "Tuesday":
			n = 2
		case "Wednesday":
			n = 3
		case "Thursday":
			n = 4
		case "Friday":
			n = 5
		case "Saturday":
			n = 6
		}
	}
	return n
}
