package main

// Day 天
type Day int

const (
	// MO 星期一
	MO Day = iota
	// TU 星期二
	TU
	// WE 星期三
	WE
	// TH 星期四
	TH
	// FR 星期五
	FR
	// SA 星期六
	SA
	// SU 星期天
	SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	return dayName[d]
}
