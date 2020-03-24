package main

import "fmt"

func hello(i int) {
	fmt.Println("hello", i)
}

//func main() {
//	for i:=0; i<100;i++ {
//		go hello(i) //开启一个单独的goroutine执行hello函数
//	}
//
//	fmt.Println("hello main")
//}

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("hello")
}
