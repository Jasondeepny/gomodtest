package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	soceMap := make(map[string]int)
	soceMap["小明"] = 90
	soceMap["小丽"] = 100
	soceMap["小在"] = 20
	delete(soceMap, "xiao")
	for k, v := range soceMap {
		fmt.Println(k, v)
	}

	mapSort()

}

func mapSort() {
	rand.Seed(time.Now().UnixNano())
	sortMap := make(map[string]int)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		//key := "stu " + strconv.Itoa(i)
		value := rand.Intn(100)
		sortMap[key] = value
	}
	slice := make([]string, 200)
	for k := range sortMap {
		slice = append(slice, k)
	}
	sort.Strings(slice)
	for _, s := range slice {
		fmt.Println(s, sortMap[s])
	}

}
