package sortpersons

import (
	"sort"
)

// Person 人
type Person struct {
	FirstName string
	LastName  string
}

// Persons 人员
type Persons []Person

// Len 数量
func (p Persons) Len() int { return len(p) }

// Less 比较
func (p Persons) Less(i, j int) bool {
	if p[i].FirstName < p[j].FirstName {
		return true
	} else if p[i].FirstName > p[j].FirstName {
		return false
	}
	return p[i].LastName < p[j].LastName
}

// Swap 交换
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Sort 排序
// 调用sort包中的方法Sort，进而调用本包中的Len()、Less()、Swap()完成实现
func (p Persons) Sort() {
	sort.Sort(p)
}

// IsSorted 是否已排序
func (p Persons) IsSorted() bool {
	return sort.IsSorted(p)
}
