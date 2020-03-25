package main

import "fmt"

func main() {
	var chanInt = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case chanInt <- i:
		case d := <-chanInt:
			fmt.Println(d)
		}
	}
}
