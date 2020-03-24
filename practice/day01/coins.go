package main

import (
	"fmt"
)

var users = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
var coins = 50
var userCoins = make(map[string]int)

func distribute() int {
	for _, user := range users {
		sum := 0
		for _, v := range user {
			switch v {
			case 'e', 'E':
				sum += 1
			case 'i', 'I':
				sum += 2
			case 'o', 'O':
				sum += 3
			case 'u', 'U':
				sum += 4
			}
		}
		userCoins[user] = sum
		fmt.Printf("%s 得到金币：%d\n", user, sum)
		coins -= sum
	}

	return coins
}

func main() {
	ret := distribute()
	fmt.Println("剩余金币：", ret)
}
