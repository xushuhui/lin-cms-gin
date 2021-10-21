package utils

import "time"

// 获取当前的时间 - 字符串
func GetCurrentTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetCurrentDate() string {
	return time.Now().Format("2006/01/02")
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

func String2time(dateString string) time.Time {
	date, _ := time.Parse("2006-01-02 15:04:05", dateString)
	return date
}
