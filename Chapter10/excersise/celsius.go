package main

import "fmt"

// Celsius 摄氏度
type Celsius float64

// String Celsius方法
func (f *Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(*f))
}
