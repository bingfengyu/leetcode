package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

//https://leetcode-cn.com/problems/permutations/submissions/
func permute(nums []int) [][]int {
	result := [][][]int{}
	temp2 := [][]int{}
	temp1 := []int{}
	temp1 = append(temp1, nums[0])
	temp2 = append(temp2, temp1)
	if len(nums) <= 0 {
		return temp2
	}
	result = append(result, temp2)
	for i := 1; i < len(nums); i++ {
		result = append(result, dp(result[i-1], nums[i]))
	}
	return result[len(nums)-1]
}

func dp(temp2 [][]int, n int) [][]int {
	temp := [][]int{}
	for i := 0; i < len(temp2); i++ {
		for j := 0; j <= len(temp2[i]); j++ {
			temp = append(temp, dpsort(temp2[i], n, j))
		}
	}
	return temp
}

func dpsort(temp []int, n, j int) []int {
	if j == len(temp) {
		temp = append(temp, n)
		return temp
	}

	tmp := []int{}
	for i := 0; i < len(temp); i++ {
		if j == i {
			tmp = append(tmp, n)
		}
		tmp = append(tmp, temp[i])
	}
	return tmp
}
