package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 1, 3, 4}
	fmt.Println(singleNumber(nums))
}

//试题地址：https://leetcode-cn.com/problems/single-number/
func singleNumber(nums []int) int {
	singlenum := nums[0]
	for i := 1; i < len(nums); i++ {
		singlenum ^= nums[i]
	}
	return singlenum
}
