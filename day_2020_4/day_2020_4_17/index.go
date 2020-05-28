package main

import "fmt"

func main() {
	nums := []int{2, 0, 0}
	fmt.Println(canJump2(nums))

}

//试题地址：https://leetcode-cn.com/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	reach := make([]bool, len(nums))

	reach[0] = true

	for i := 0; i < len(nums); i++ {
		if !reach[i] {
			return false
		}
		if i+nums[i] >= len(nums)-1 {
			return true
		}
		for j := i + 1; j <= i+nums[i]; j++ {
			reach[j] = true
		}
		fmt.Println(reach)
	}
	return false
}

func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	i := len(nums) - 1
	for {
		j := i - 1
		for {
			if j < 0 || nums[j] >= i-j {
				break
			}
			j--
		}

		if j < 0 {
			return false
		}
		if j == 0 {
			return true
		}
		i = j
	}
	return false
}
