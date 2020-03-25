package main

import (
	"fmt"
	"strings"
)

func Split(str string, sep string) []string {
	var ret = make([]string, 0)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}

func main() {
	ret := Split("babedberdbd", "b")
	ret2 := Split("e", "b")
	fmt.Println(ret)
	fmt.Println(ret2)
}
