package main

// TZ 时区
type TZ int

const (
	// HOUR HOUR
	HOUR TZ = 60 * 60
	// UTC UTC
	UTC TZ = 0 * HOUR
	// EST EST
	EST TZ = -5 * HOUR
	// CST CST
	CST TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich Time",
	EST: "Eastern Standard Time",
	CST: "Central Standard Time"}

func (t TZ) String() string {
	if zone, ok := timeZones[t]; ok {
		return zone
	}
	return ""
}
