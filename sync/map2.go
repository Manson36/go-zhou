package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m2 = sync.Map{} //不需要make初始化，是结构体

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         //存储必须使用Store方法
			value, _ := m2.Load(key) //获取必须使用Load方法
			fmt.Printf("k=%v, v = %v\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
