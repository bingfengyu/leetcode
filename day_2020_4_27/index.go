package main

import "fmt"

func main() {
	nums := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	for _, v := range nums {
		fmt.Println(search(nums, v))
	}
}

//试题地址:https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i := 0
	j := len(nums) - 1
	if len(nums) > 1 && nums[i] > nums[j] { //非有序数组
		for i+1 != j { //查找数组最小值
			mid := (i + j) / 2
			if nums[i] > nums[mid] {
				j = mid
			}
			if nums[mid] > nums[j] {
				i = mid
			}
			// fmt.Println(i, mid, j)
		}
	}
	// fmt.Println(nums[0:j], nums[j:])
	//两个有序数组查找target
	left := midSearch(nums[0:j], target)
	right := midSearch(nums[j:], target)
	if left != -1 {
		return left
	}
	if right != -1 {
		return right + j
	}
	return -1
}

func midSearch(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
		// fmt.Println(i, j)
	}
	return -1
}
