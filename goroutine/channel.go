package main

import (
	"fmt"
	"sync"
)

var b chan int
var wg sync.WaitGroup

func main() {
	fmt.Println(b)
	b = make(chan int)
	wg.Add(1)
	go func() {
		x := <-b
		wg.Done()
		fmt.Println("b中取出", x)
	}()
	b <- 10
	fmt.Println("10发送到b中了")
	fmt.Println(b)
	wg.Wait()
}
