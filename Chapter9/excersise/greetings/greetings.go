package greetings

import (
	"time"
)

// GoodDay 早安
func GoodDay(name string) string {
	return "Good Day " + name
}

// GoodNight 晚安
func GoodNight(name string) string {
	return "Good Night " + name
}

// IsAm 是否上午
func IsAm() bool {
	localTime := time.Now()
	return localTime.Hour() <= 12
}

// IsAfternoon 是否下午
func IsAfternoon() bool {
	localTime := time.Now()
	return localTime.Hour() <= 18
}

// IsEvening 是否晚上
func IsEvening() bool {
	localTime := time.Now()
	return localTime.Hour() <= 22
}
