package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序
//1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan中
//3. 主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}
type result struct {
	*job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func distribute(d chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{x}
		d <- newJob
		time.Sleep(time.Millisecond * 500)
	}

}

func executor(d <-chan *job, e chan<- *result) {
	defer wg.Done()
	for {
		job := <-d
		n := job.value
		var sum int64
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job,
			sum,
		}
		e <- newResult
	}
}

func main() {
	wg.Add(1)
	go distribute(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go executor(jobChan, resultChan)
	}

	for v := range resultChan {
		fmt.Printf("int64的数：%d 各位和为%d\n", v.value, v.sum)
	}
	wg.Wait()
}
