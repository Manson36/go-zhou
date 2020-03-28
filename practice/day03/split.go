package main

import (
	"fmt"
	"strings"
)

func Split(str string, sep string) []string {
	reply := make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		reply = append(reply, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	reply = append(reply, str)
	return reply
}

func main() {
	reply := Split("a:b:c:d", ":")
	fmt.Println(reply)
}
