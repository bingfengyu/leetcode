package main

import "fmt"

func main() {
	// nums := []int{1, -6, 7, 2, 3, -16, 2, 3, 7}
	// fmt.Println(maxSubArray(nums))

	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}
	// fmt.Println(findMedianSortedArrays(nums1, nums2))

	fmt.Println(longestPalindrome("1212"))
}

//试题地址：https://leetcode-cn.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for j > i {
		if s[j] == s[i] {
			fmt.Println(string(s[i]), string(s[j]))
			i++
			j--
			continue
		}
		return valid(s[i:j]) || valid(s[i+1:j+1])
	}
	return true
}

func valid(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

//试题地址：https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	maxnum, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxnum = max(nums[i], maxnum+nums[i])
		result = max(maxnum, result)
		// fmt.Println(maxnum, result)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 试题地址：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 1 {
		if len(nums1) == 0 {
			return float64(nums2[0])
		}
		return float64(nums1[0])
	}
	if len(nums1) == 1 && len(nums2) == 1 {
		return (float64(nums1[0]) + float64(nums2[0])) / 2
	}

	tmp := (len(nums1)+len(nums2))/2 - 1
	tmp1 := tmp + 1
	result1, result2 := 0, 0
	flag1, flag2 := true, true
	i, j := 0, 0
	for {
		if !flag1 && !flag2 {
			break
		}
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				if tmp == i+j && flag1 {
					result1 = nums2[j]
					flag1 = false
				}
				if tmp1 == i+j && flag2 {
					result2 = nums2[j]
					flag2 = false
				}
				j++
				continue
			} else {
				if tmp == i+j && flag1 {
					result1 = nums1[i]
					flag1 = false
				}
				if tmp1 == i+j && flag2 {
					result2 = nums1[i]
					flag2 = false
				}
				i++
				continue
			}
		}
		if i < len(nums1) {
			if tmp == i+j && flag1 {
				result1 = nums1[i]
				flag1 = false
			}
			if tmp1 == i+j && flag2 {
				result2 = nums1[i]
				flag2 = false
			}
			i++
			continue
		}
		if j < len(nums2) {
			if tmp == i+j && flag1 {
				result1 = nums2[j]
				flag1 = false
			}
			if tmp1 == i+j && flag2 {
				result2 = nums2[j]
				flag2 = false
			}
			j++
			continue
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(result1) + float64(result2)) / 2
	}
	return float64(result2)
}

//试题地址：https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if trueString(s[i:j]) {
				if len(str) < len(s[i:j]) {
					str = s[i:j]
				}
			}
		}
		if len(str) > len(s)-i {
			break
		}
	}
	return str
}

func trueString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
