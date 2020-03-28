package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker2(ctx context.Context) {
	defer wg.Done()

	key := TraceCode("TRACE_CODE")
	tracecode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Println("worker, trace code :", tracecode)
		time.Sleep(time.Millisecond * 10) //假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
}

func main() {
	//设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	//在系统的入口设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123456")
	wg.Add(1)
	go worker2(ctx)
	time.Sleep(time.Second * 5)
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
