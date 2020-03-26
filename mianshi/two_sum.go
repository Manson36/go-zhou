package main

import "fmt"

//两数求和为目标数
func main() {
	nums := []int{2, 7, 9, 11}
	target := 90
	ret := twoNum(nums, target)
	fmt.Println(ret)
}

func twoNum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
