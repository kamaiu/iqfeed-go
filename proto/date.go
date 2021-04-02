package proto

import (
	"strconv"
	"time"
)

var _YEAR = time.Now().Year()

// HH:mm:ss
func ParseDurationSecs(v string) time.Duration {
	if len(v) < 8 {
		return time.Duration(0)
	}
	var (
		hour, minute, second int64
		err                  error
	)

	hour, err = strconv.ParseInt(v[0:2], 10, 8)
	if err != nil {
		return time.Duration(0)
	}
	minute, err = strconv.ParseInt(v[3:5], 10, 8)
	if err != nil {
		return time.Duration(0)
	}
	second, err = strconv.ParseInt(v[6:8], 10, 8)
	if err != nil {
		return time.Duration(0)
	}

	return (time.Hour * time.Duration(hour)) +
		(time.Minute * time.Duration(minute)) +
		(time.Second * time.Duration(second))
}

// HH:mm:ss.ffffff
func ParseDuration(v string) time.Duration {
	if len(v) < 8 {
		return time.Duration(0)
	}
	var (
		hour, minute, second, micro int64
		err                         error
	)

	hour, err = strconv.ParseInt(v[0:2], 10, 8)
	if err != nil {
		return time.Duration(0)
	}
	minute, err = strconv.ParseInt(v[3:5], 10, 8)
	if err != nil {
		return time.Duration(0)
	}
	second, err = strconv.ParseInt(v[6:8], 10, 8)
	if err != nil {
		return time.Duration(0)
	}

	if len(v) > 9 && v[8] == '.' {
		micro, err = strconv.ParseInt(v[9:], 10, 64)
		if err != nil {
			return time.Duration(0)
		}
	} else {
		micro = 0
	}

	return (time.Hour * time.Duration(hour)) +
		(time.Minute * time.Duration(minute)) +
		(time.Second * time.Duration(second)) +
		(time.Microsecond * time.Duration(micro))
}

// yyyyMMdd HHmmss
func ParseDateTime(v string) time.Time {
	if len(v) < 15 {
		return time.Time{}
	}

	var (
		year, month, day, hour, minute, second int64
		err                                    error
	)

	year, err = strconv.ParseInt(v[0:4], 10, 16)
	if err != nil {
		return time.Time{}
	}

	month, err = strconv.ParseInt(v[4:6], 10, 8)
	if err != nil {
		return time.Time{}
	}
	day, err = strconv.ParseInt(v[6:8], 10, 8)
	if err != nil {
		return time.Time{}
	}
	hour, err = strconv.ParseInt(v[9:11], 10, 8)
	if err != nil {
		return time.Time{}
	}
	minute, err = strconv.ParseInt(v[11:13], 10, 8)
	if err != nil {
		return time.Time{}
	}
	second, err = strconv.ParseInt(v[13:15], 10, 8)
	if err != nil {
		return time.Time{}
	}

	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
}

// MM/dd/yyyy
func ParseDate(v string) time.Time {
	if len(v) < 10 {
		return time.Time{}
	}
	var (
		month, day, year int64
		err              error
	)

	month, err = strconv.ParseInt(v[0:2], 10, 8)
	if err != nil {
		return time.Time{}
	}
	day, err = strconv.ParseInt(v[3:5], 10, 16)
	if err != nil {
		return time.Time{}
	}
	year, err = strconv.ParseInt(v[6:], 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
}

// MM/dd/yyyy
func AppendDate(o []byte, t time.Time) []byte {
	switch t.Month() {
	case time.January:
		o = append(o, "01"...)
	case time.February:
		o = append(o, "02"...)
	case time.March:
		o = append(o, "03"...)
	case time.April:
		o = append(o, "04"...)
	case time.May:
		o = append(o, "05"...)
	case time.June:
		o = append(o, "06"...)
	case time.July:
		o = append(o, "07"...)
	case time.August:
		o = append(o, "08"...)
	case time.September:
		o = append(o, "09"...)
	case time.October:
		o = append(o, "10"...)
	case time.November:
		o = append(o, "11"...)
	case time.December:
		o = append(o, "12"...)
	}
	o = append(o, '/')
	switch t.Day() {
	case 0:
		o = append(o, "01"...)
	case 1:
		o = append(o, "01"...)
	case 2:
		o = append(o, "02"...)
	case 3:
		o = append(o, "03"...)
	case 4:
		o = append(o, "04"...)
	case 5:
		o = append(o, "05"...)
	case 6:
		o = append(o, "06"...)
	case 7:
		o = append(o, "07"...)
	case 8:
		o = append(o, "08"...)
	case 9:
		o = append(o, "09"...)
	case 10:
		o = append(o, "10"...)
	case 11:
		o = append(o, "11"...)
	case 12:
		o = append(o, "12"...)
	case 13:
		o = append(o, "13"...)
	case 14:
		o = append(o, "14"...)
	case 15:
		o = append(o, "15"...)
	case 16:
		o = append(o, "16"...)
	case 17:
		o = append(o, "17"...)
	case 18:
		o = append(o, "18"...)
	case 19:
		o = append(o, "19"...)
	case 20:
		o = append(o, "20"...)
	case 21:
		o = append(o, "21"...)
	case 22:
		o = append(o, "22"...)
	case 23:
		o = append(o, "23"...)
	case 24:
		o = append(o, "24"...)
	case 25:
		o = append(o, "25"...)
	case 26:
		o = append(o, "26"...)
	case 27:
		o = append(o, "27"...)
	case 28:
		o = append(o, "28"...)
	case 29:
		o = append(o, "29"...)
	case 30:
		o = append(o, "30"...)
	case 31:
		o = append(o, "31"...)
	}
	o = append(o, '/')
	o = append(o, strconv.Itoa(t.Year())...)
	return o
}

// MMM dd H:mmtt
func ParseMonthDayTime(v string) time.Time {
	if len(v) < 13 {
		return time.Time{}
	}
	var month time.Month
	switch v[0:3] {
	case "Jan", "jan", "JAN":
		month = time.January
	case "Feb", "feb", "FEB":
		month = time.February
	case "Mar", "mar", "MAR":
		month = time.March
	case "Apr", "apr", "APR":
		month = time.April
	case "Jun", "jun", "JUN":
		month = time.June
	case "Jul", "jul", "JUL":
		month = time.July
	case "Aug", "aug", "AUG":
		month = time.August
	case "Sep", "sep", "SEP":
		month = time.September
	case "Oct", "oct", "OCT":
		month = time.October
	case "Nov", "nov", "NOV":
		month = time.November
	case "Dec", "dec", "DEC":
		month = time.December
	}

	day := 0
	switch v[4] {
	case '0':
		switch v[5] {
		case '0':
		case '1':
			day = 1
		case '2':
			day = 2
		case '3':
			day = 3
		case '4':
			day = 4
		case '5':
			day = 5
		case '6':
			day = 6
		case '7':
			day = 7
		case '8':
			day = 8
		case '9':
			day = 9
		}
	case '1':
		switch v[5] {
		case '0':
			day = 10
		case '1':
			day = 11
		case '2':
			day = 12
		case '3':
			day = 13
		case '4':
			day = 14
		case '5':
			day = 15
		case '6':
			day = 16
		case '7':
			day = 17
		case '8':
			day = 18
		case '9':
			day = 19
		}
	case '2':
		switch v[5] {
		case '0':
			day = 20
		case '1':
			day = 21
		case '2':
			day = 22
		case '3':
			day = 23
		case '4':
			day = 24
		case '5':
			day = 25
		case '6':
			day = 26
		case '7':
			day = 27
		case '8':
			day = 28
		case '9':
			day = 29
		}
	case '3':
		switch v[5] {
		case '0':
			day = 30
		default:
			day = 31
		}
	}

	hour := 0
	switch v[7] {
	case '0':
	case '1':
		hour = 1
	case '2':
		hour = 2
	case '3':
		hour = 3
	case '4':
		hour = 4
	case '5':
		hour = 5
	case '6':
		hour = 6
	case '7':
		hour = 7
	case '8':
		hour = 8
	case '9':
		hour = 9
	}
	if v[8] != ':' {
		switch hour {
		case 0:
			switch v[8] {
			case '0':
			case '1':
				hour = 1
			case '2':
				hour = 2
			case '3':
				hour = 3
			case '4':
				hour = 4
			case '5':
				hour = 5
			case '6':
				hour = 6
			case '7':
				hour = 7
			case '8':
				hour = 8
			case '9':
				hour = 9
			}
		case 1:
			switch v[8] {
			case '0':
				hour = 10
			case '1':
				hour = 11
			case '2':
				hour = 12
			}
		}
	}

	minute := 0
	switch v[len(v)-4] {
	case '0':
		switch v[len(v)-3] {
		case '0':
		case '1':
			minute = 1
		case '2':
			minute = 2
		case '3':
			minute = 3
		case '4':
			minute = 4
		case '5':
			minute = 5
		case '6':
			minute = 6
		case '7':
			minute = 7
		case '8':
			minute = 8
		case '9':
			minute = 9
		}
	case '1':
		switch v[len(v)-3] {
		case '0':
			minute = 10
		case '1':
			minute = 11
		case '2':
			minute = 12
		case '3':
			minute = 13
		case '4':
			minute = 14
		case '5':
			minute = 15
		case '6':
			minute = 16
		case '7':
			minute = 17
		case '8':
			minute = 18
		case '9':
			minute = 19
		}
	case '2':
		switch v[len(v)-3] {
		case '0':
			minute = 20
		case '1':
			minute = 21
		case '2':
			minute = 22
		case '3':
			minute = 23
		case '4':
			minute = 24
		case '5':
			minute = 25
		case '6':
			minute = 26
		case '7':
			minute = 27
		case '8':
			minute = 28
		case '9':
			minute = 29
		}
	case '3':
		switch v[len(v)-3] {
		case '0':
			minute = 30
		case '1':
			minute = 31
		case '2':
			minute = 32
		case '3':
			minute = 33
		case '4':
			minute = 34
		case '5':
			minute = 35
		case '6':
			minute = 36
		case '7':
			minute = 37
		case '8':
			minute = 38
		case '9':
			minute = 39
		}
	case '4':
		switch v[len(v)-3] {
		case '0':
			minute = 40
		case '1':
			minute = 41
		case '2':
			minute = 42
		case '3':
			minute = 43
		case '4':
			minute = 44
		case '5':
			minute = 45
		case '6':
			minute = 46
		case '7':
			minute = 47
		case '8':
			minute = 48
		case '9':
			minute = 49
		}
	case '5':
		switch v[len(v)-3] {
		case '0':
			minute = 50
		case '1':
			minute = 51
		case '2':
			minute = 52
		case '3':
			minute = 53
		case '4':
			minute = 54
		case '5':
			minute = 55
		case '6':
			minute = 56
		case '7':
			minute = 57
		case '8':
			minute = 58
		case '9':
			minute = 59
		}
	}

	switch v[len(v)-2:] {
	case "AM", "am":
		if hour == 12 {
			hour = 0
		}
	case "PM", "pm":
		if hour < 12 {
			hour += 12
		}
	}

	return time.Date(_YEAR, month, day, hour, minute, 0, 0, time.Local)
}
