package main

import "fmt"

func main() {
	var ss = make([]map[int]string, 10, 10)
	ss[0] = make(map[int]string, 10)
	ss[0][0] = "liqin"

	for _, v := range ss {
		fmt.Println(v)
	}
}
