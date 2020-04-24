package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 2, 1}
	fmt.Println(reversePairs(nums))
}

//试题地址：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

//归并排序
func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	sumStep := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	//开始合并
	temp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] { //此时nums[i]>=nums[j],那么mid+1~j都是小于num[i]
			sumStep += j - mid - 1
			temp = append(temp, nums[i])
			i++
			continue
		}
		temp = append(temp, nums[j])
		j++
	}
	for j <= end {
		temp = append(temp, nums[j])
		j++
	}
	for i <= mid { //此时i仍小于mid，表明剩下得数字大于后面得最大值
		temp = append(temp, nums[i])
		sumStep += end - mid
		i++
	}
	//将排序好的数组放入原始数组种
	for i := start; i <= end; i++ {
		nums[i] = temp[i-start]
	}
	return sumStep
}
