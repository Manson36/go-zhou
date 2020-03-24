package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中有几个汉字
	var s = "hello, 李沁女神"
	count := 0

	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}

	println(count)

	//2.判断一段英文中单词出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")

	wordMap := make(map[string]int, 10)
	for _, v := range s3 {
		//wordMap[v] +=1
		if _, ok := wordMap[v]; ok {
			wordMap[v] += 1
		} else {
			wordMap[v] = 1
		}
	}

	for i, v := range wordMap {
		fmt.Println(i, v)
	}

	//回文判断
	ss := "黄山落叶松叶落山黄"
	//发现问题：首先要将字符串放入rune切片中
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}

	var flag bool
	for i := 0; i < len(r)/2; i++ {
		if r[i] == r[len(r)-i-1] {
			flag = true
		}
	}

	if flag {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}
