package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, -3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}

//试题地址：https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	count := 0
	arraysum := []int{}
	arraysum = append(arraysum, nums[0])
	if nums[0] == k {
		count++
	}
	for i := 1; i < len(nums); i++ {
		if arraysum[i-1]+nums[i] == k {
			count++
		}
		arraysum = append(arraysum, arraysum[i-1]+nums[i])
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			arraysum[j] -= nums[i-1]
			if arraysum[j] == k {
				count++
			}
		}
	}
	return count
}
