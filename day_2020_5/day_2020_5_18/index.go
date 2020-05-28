package main

import "fmt"

func main() {
	nums := []int{9, 2, -6, 8, 8, 0, 3, -4, 5, -6}
	fmt.Println(maxProduct(nums))
}

//试题地址:https://leetcode-cn.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	maxnum, minnum, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpmax, tmpmin := maxnum, minnum
		maxnum = max(tmpmax*nums[i], max(nums[i], tmpmin*nums[i]))
		minnum = min(tmpmin*nums[i], min(nums[i], tmpmax*nums[i]))
		result = max(result, maxnum)
		fmt.Println(maxnum, minnum, result)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
