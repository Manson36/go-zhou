package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("女神，李沁")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //context的Done方法
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //参数是根目录的context
	wg.Add(1)
	go f2(ctx) //传入context参数

	time.Sleep(time.Second * 5)
	//如何通知子goroutine退出
	cancel()
	wg.Wait()
}
