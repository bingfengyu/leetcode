package main

import "fmt"

func main() {
	nums := []int{4, 1, 3, 2, 3}
	fmt.Println(findDuplicate(nums))
}

//试题地址：https://leetcode-cn.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	n := len(nums)
	l := 1
	r := n - 1
	for r > l {
		mid := (l + r) / 2
		count := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid && nums[i] >= l {
				count++
			}
		}
		// fmt.Println(l, mid, r, count)
		if count > mid-l+1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
