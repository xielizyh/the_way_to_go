package main

import (
	"fmt"
	"sort"
)

func main() {
	week := make(map[int]string)
	week[0] = "Monday"
	week[1] = "Tuesday"
	week[2] = "Wednesday"
	week[3] = "Thursday"
	week[4] = "Friday"
	week[5] = "Saturday"
	week[6] = "Sunday"
	for key, value := range week {
		fmt.Printf("The day %d is %s\n", key, value)
	}
	isTuesdayExist := false
	isHollydayExist := false
	for _, value := range week {
		if value == "Tuesday" {
			isTuesdayExist = true
			break
		} else if value == "Hollyday" {
			isHollydayExist = true
			break
		}
	}
	if isTuesdayExist {
		fmt.Printf("Tuesday is present\n")
	} else {
		fmt.Printf("Tuesday is not present\n")
	}
	if isHollydayExist {
		fmt.Printf("Hollday is present\n")
	} else {
		fmt.Printf("Hollday is not present\n")
	}
	// 8.2
	fmt.Println("------Excersise 8.2------")
	drink := make(map[string]string)
	drink["coffee"] = "咖啡"
	drink["cola"] = "可乐"
	drink["milk"] = "牛奶"
	drink["green tea"] = "绿茶"
	for key, value := range drink {
		fmt.Printf("%s=%s\n", key, value)
	}
	i := 0
	keys := make([]string, len(drink))
	for k := range drink {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v\n", k, drink[k])
	}
}
