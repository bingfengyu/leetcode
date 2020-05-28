package main

import "fmt"

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	// nums := []int{1, 1, 1, 1, 1}
	// nums := []int{1, 2, 0, 1}
	nums := []int{3, 1, 2, 3}
	fmt.Println(jump(nums))
}

//试题地址：https://leetcode-cn.com/problems/jump-game-ii/
//dp做法 超时
func jumpDp(nums []int) int {
	result := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == 0 {
				if j <= i+nums[i] && j != 0 {
					result[i][j] = result[i][i] + 1
				}
			} else {
				if j < i+1 || j > i+nums[i] {
					result[i][j] = result[i-1][j]
				} else {
					if result[i-1][j] != 0 {
						result[i][j] = min(result[i][i]+1, result[i-1][j])
					} else {
						result[i][j] = result[i][i] + 1
					}
				}
			}
		}
		if i+nums[i] == len(nums)-1 {
			return result[i][len(nums)-1]
		}
	}
	fmt.Println(result)
	return result[len(nums)-1][len(nums)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//贪心做法
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	maxlen := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-2; i++ {
		maxlen = max(maxlen, i+nums[i]) //改step下 最远可到达距离
		if maxlen >= len(nums)-1 {
			return step + 1 //下一步一定可以到达
		}
		if i == end {
			end = maxlen
			step++
		}
		fmt.Println(i, end, maxlen, step)
	}
	return step + 1 //下一步一定可以到达
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
