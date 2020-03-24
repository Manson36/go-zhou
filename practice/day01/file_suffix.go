package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffix(".jpg")
	txtFunc := makeSuffix(".txt")

	fmt.Println(jpgFunc("dalin"))
	fmt.Println(txtFunc("xioa.txt"))
}
