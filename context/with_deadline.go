package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	//尽管ctx会过期而结束，但是在任何情况下调用它的cancel函数都是很好的习惯。
	//如果不这样做，可能会使context和其父类存活时间超过设置的时间
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(1 * time.Second): //间隔1秒钟
		fmt.Println("李沁 女神")
	}
}
