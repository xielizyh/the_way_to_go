package main

import "fmt"

func main() {
	// 8.1
	fmt.Println("------Example 8.1------")
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	// mapAssigned 也是 mapList 的引用
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\": %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\": %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\": %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\": %d\n", mapLit["ten"])

	// 8.2
	fmt.Println("------Example 8.2------")
	// 使用func() int作为map的值
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	// 8.4
	fmt.Println("------Example 8.4------")
	map1 := map[string]int{"New Delhi": 55, "Beijing": 20, "Washington": 25}
	value, isPresent := map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain \"Beijing\"")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("The value of \"Paris\" in map1 is: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Washington")
	}
	// 8.5
	fmt.Println("------Example 8.5------")
	map2 := make(map[int]float32)
	map2[1] = 1.0
	map2[2] = 2.0
	map2[3] = 3.0
	map2[4] = 4.0
	for key, value := range map2 {
		fmt.Printf("key is： %d - value is: %f\n", key, value)
	}
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
	// 8.4
	fmt.Println("------Example 8.4------")
	// 切片里每一个元素都是map类型
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
