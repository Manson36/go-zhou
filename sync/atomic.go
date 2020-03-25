package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	//lock.Lock()
	//x++
	//lock.Unlock()
	atomic.AddInt64(&x, 1) //这里相当于前面的三部操作
	wg.Done()
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
