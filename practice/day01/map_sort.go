package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//
	//var scoreMap = make(map[string]int)
	//
	//for i:= 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%2d", i)
	//	value := rand.Intn(100)
	//	scoreMap[key] = value
	//}
	//
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys,key)
	//}
	//
	//sort.Strings(keys)
	//for _, k := range keys {
	//	fmt.Println(k, scoreMap[k])
	//}
	kk()
}

func kk() {
	rand.Seed(time.Now().UnixNano())
	stuMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100)
		stuMap[key] = value
	}

	var ss = make([]string, 0, 200)
	for k := range stuMap {
		ss = append(ss, k)
	}

	sort.Strings(ss)
	for _, v := range ss {
		fmt.Println(v, stuMap[v])
	}
}
