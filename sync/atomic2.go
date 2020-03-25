package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//比较并交换
	var x int32 = 300
	ok := atomic.CompareAndSwapInt32(&x, x, 500)
	fmt.Println(x, ok)
}
