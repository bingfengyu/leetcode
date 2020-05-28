package main

import "fmt"

func main() {
	nums := []int{4, 5, 0, -2, -3, 1}
	fmt.Println(subarraysDivByK(nums, 5))
}

//试题地址：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		ans += record[modulus] //同余定理 A[0:i]%k=4 A[0:j]%k=4 那么A[i:j]%k==0
		record[modulus]++
	}
	return ans
}

func subarraysDivByK1(nums []int, K int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			if tmp%K == 0 {
				result++
				// fmt.Println(tmp)
			}
		}
	}
	return result
}

func subarraysDivByK2(nums []int, K int) int {
	result := 0
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == 0 {
				if j != 0 {
					tmp[j] += tmp[j-1] + nums[j]
				} else {
					tmp[j] = nums[j]
				}
			} else {
				tmp[j] -= tmp[i-1]
			}
			if tmp[j]%K == 0 {
				result++
			}
		}
		fmt.Println(tmp)
	}
	return result
}
