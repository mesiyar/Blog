package util

import "time"

func StrToUnix(str string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return t.Unix()
}

func UnixToStr(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006-01-02 15:04:05")
}

func TodayZeroTime() int64 {
	t := time.Now()
	return t.Unix() - int64(t.Hour()*3600+t.Minute()*60+t.Second())
}

func TodayEndTime() int64 {
	t := time.Now()
	return t.Unix() + int64((24-t.Hour())*3600+(60-t.Minute())*60+(60-t.Second()))
}
