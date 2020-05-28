package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 8, 5, 1, 5, 2, 2}
	fmt.Println(singleNumbers(nums))

}

//试题地址：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func singleNumbers1(nums []int) []int {
	result := []int{}
	if len(nums) == 2 {
		return nums
	}

	sort.Ints(nums) //调用内部排序 应该不是这题的考点
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			i++
		} else {
			result = append(result, nums[i])
		}
	}
	if len(result) == 1 {
		result = append(result, nums[len(nums)-1])
	}
	return result
}

func singleNumbers(nums []int) []int {
	var a int
	for i := range nums {
		a ^= nums[i] //对于相同数字做^运算结果是0，也就是a最后的结果是剩下的两个数字相^
	}
	mask := a & (-a)
	fmt.Println(a, mask)
	res := make([]int, 2)
	for _, v := range nums {
		if (v & mask) == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
