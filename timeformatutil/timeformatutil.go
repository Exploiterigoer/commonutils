package timeformatutil

import (
	"strconv"
	"strings"
	"time"
)

// TimeEncoder time formatter
// parameter "t", the time to be formated，
// first return value with the form of YYYYmmddhhiiss
// second return value with the form of YYYYmmddhhiissnn
func TimeEncoder(t time.Time) (string, string) {
	second := t.Unix()                        // second,int64
	necond := t.UnixNano()                    // nanosecond,int64 1s=10^9 ns
	mtime := int((necond - second*1e9) / 1e6) // millisecond
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
	if len(n) == 1 { // millisecond string in the form of 3 number
		n = "00" + n
	} else if len(n) == 2 {
		n = "0" + n
	}

	timeString := time.Now().Format("2006-01-02 15:04:05")
	ym := strings.Split(strings.Split(timeString, " ")[0], "-")
	rtn1 := ym[0] + ym[1] + ym[2] + h + i + s
	rtn2 := ym[0] + ym[1] + ym[2] + h + i + s + n
	return rtn1, rtn2
}

// TimeDecoder formats a string to a normal time string
// returns the time string with the form of second or millisecond
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

// WeekDateTimeInteger  Gets the integer of the year,month,day,hour,minutes,second and the week
// parameter "t",the time to be cut
// parameter "flag",type for needding,
// y or Y for year,m or M for month,d or D for day,h or H for hour,
// i or I for minutes,s or S for second,w or W for week
// NOTE:if the flag is w or W,the return value from 0 to 7,which corresponds of Sunday to Saturday
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
