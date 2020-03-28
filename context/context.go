package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()

	for {
		if notify {
			break
		}
		fmt.Println("女神，李沁")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	wg.Add(1)
	go f()
	//如何通知子goroutine退出
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
