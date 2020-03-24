package main

import "fmt"

func jiecheng(n int) int {
	if n <= 1 {
		return 1
	}
	return jiecheng(n-1) * n
}

func taijie(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}

func main() {
	fmt.Println(jiecheng(7))
	fmt.Println(taijie(4))
}
