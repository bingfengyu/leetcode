package main

import (
	"fmt"
)

func main() {
	// nums := []int{2, 4, 6}
	nums := []int{1, 2, 3, 1, 2, 4, 1, 2, 3, 2, 1, 3}
	fmt.Println(numberOfSubarrays(nums, 4))
}

//试题地址：https://leetcode-cn.com/problems/count-number-of-nice-subarrays/submissions/
func numberOfSubarrays(nums []int, k int) int {
	num, i, j, sum, temp := 0, 0, -1, 0, 0
	for ; i < len(nums); i++ { //找出第一个符合的数组
		if nums[i]%2 == 1 {
			sum++
		}
		if sum == k {
			j = i
			break
		}
	}
	if j == -1 { //数组中的奇数总个数小于k个
		return num
	}
	// fmt.Println(nums, j)
	i = 0
	flag := true
	for i < len(nums) {
		for {
			if sum > k {
				sum-- //回溯
				break
			}
			if sum == k {
				fmt.Println(i, j)
				if flag {
					temp = j //记住j的上一个基数位置
					flag = false
				}
				num++
			}
			j++
			if j >= len(nums) {
				break
			}
			sum += nums[j] % 2
		}
		if j >= len(nums) && nums[i]%2 == 1 { //保证这是最后一个
			break
		}
		sum -= nums[i] % 2
		j = temp
		flag = true
		i++
	}
	return num
}
