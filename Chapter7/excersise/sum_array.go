package main

// SumArray 计算float数组的和
func SumArray(arrF []float32) float64 {
	sum := 0.0
	for _, v := range arrF {
		sum += float64(v)
	}
	return sum
}

// SumArray1 计算float数组的和
func SumArray1(arrF ...float32) float64 {
	sum := 0.0
	for _, v := range arrF {
		sum += float64(v)
	}
	return sum
}

// SumAndAverage 计算和和平均值
func SumAndAverage(a int, b float32) (float32, float32) {
	return float32(a) + b, (float32(a) + b) / 2.0
}
