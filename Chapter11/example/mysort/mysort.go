package mysort

// Sorter 排序接口
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort 排序
func Sort(data Sorter) {
	for pass := 1; pass <= data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

// IsSorted 是否排序
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//IntArray int切片
type IntArray []int

// Len 个数
func (p IntArray) Len() int { return len(p) }

// Less 比较大小
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }

// Swap 交换
func (p IntArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// StringArray string切片
type StringArray []string

// Len 个数
func (p StringArray) Len() int { return len(p) }

// Less 比较大小
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }

// Swap 交换
func (p StringArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SortInts 通用包装
func SortInts(a []int) { Sort(IntArray(a)) }

// StringInts 通用包装
func StringInts(a []string) { Sort(StringArray(a)) }

// IntAreSorted 通用包装
func IntAreSorted(a []int) bool { return IsSorted(IntArray(a)) }

// StringAreSorted 通用包装
func StringAreSorted(a []string) bool { return IsSorted(StringArray(a)) }

// Day 自定义
type Day struct {
	Num       int
	ShortName string
	LongName  string
}

// DayArray 自定义
type DayArray struct {
	Data []*Day //指针数组
}

// Len 个数
func (p *DayArray) Len() int { return len(p.Data) }

// Less 比较大小
func (p *DayArray) Less(i, j int) bool { return p.Data[i].Num < p.Data[j].Num }

// Swap 交换
func (p *DayArray) Swap(i, j int) { p.Data[i], p.Data[j] = p.Data[j], p.Data[i] }
