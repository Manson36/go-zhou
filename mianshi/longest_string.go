package main

import "fmt"

func main() {
	//最长不重复子串
	s := "jsonzifuchuan"
	sum := 0
	ret := ""
	var mapC = make(map[int32]int, 20)
	for i, v := range s {
		if value, ok := mapC[v]; ok {
			if i-value > sum {
				sum = i - value
				ret = s[value:i]
			}
		}
		mapC[v] = i
	}

	fmt.Println(ret)
}
