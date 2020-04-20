package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// intervals := [][]int{
	// 	{1, 3},
	// 	{2, 6},
	// 	{8, 10},
	// 	{2, 9},
	// 	{15, 18},
	// }
	// intervals := [][]int{
	// 	{1, 6},
	// 	{4, 5},
	// }
	// fmt.Println(merge(intervals))

	// s := "abcdefaa"
	// fmt.Println(s[2:8])
	// fmt.Println(startLocation(s[2:8], "a"), string(s[2+startLocation(s[2:8], "a")]))

	fmt.Println(lengthOfLongestSubstring("abcddfrgfdytuvilliuhu"))
}

//试题地址：https://leetcode-cn.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	result := [][][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals)
	i := 0
	for {
		k := 0
		result = append(result, [][]int{})
		for j := 0; j < len(result[i])-1; j++ {
			result[i+1] = append(result[i+1], []int{})
			fmt.Println(result)
			if result[i][j][1] >= result[i][j+1][0] {
				if result[i][j][1] >= result[i][j+1][1] {
					result[i+1][k] = append(result[i+1][k], result[i][j][0])
					result[i+1][k] = append(result[i+1][k], result[i][j][1])
				} else {
					result[i+1][k] = append(result[i+1][k], result[i][j][0])
					result[i+1][k] = append(result[i+1][k], result[i][j+1][1])
				}
				j++
			} else {
				result[i+1][k] = append(result[i+1][k], result[i][j][0])
				result[i+1][k] = append(result[i+1][k], result[i][j][1])
			}
			k++
			if j == len(result[i])-2 {
				result[i+1] = append(result[i+1], []int{})
				result[i+1][k] = append(result[i+1][k], result[i][j+1][0])
				result[i+1][k] = append(result[i+1][k], result[i][j+1][1])
			}
		}
		fmt.Println(result)
		if len(result[i]) == len(result[i+1]) || len(result[i]) == 1 {
			break
		}
		i++
	}
	return result[i]
}

//试题地址：https://leetcode-cn.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := l
	flag := 0
	for {
		if l1 == nil && l2 != nil {
			for {
				temp.Next = &ListNode{
					Val:  (l2.Val + flag) % 10,
					Next: nil,
				}
				flag = (l2.Val + flag) / 10
				l2 = l2.Next
				temp = temp.Next
				if l2 == nil {
					if flag == 0 {
						temp.Next = nil
					} else {
						temp.Next = &ListNode{
							Val:  1,
							Next: nil,
						}
					}
					break
				}
			}
			break
		}
		if l2 == nil && l1 != nil {
			for {
				temp.Next = &ListNode{
					Val:  (l1.Val + flag) % 10,
					Next: nil,
				}
				flag = (l1.Val + flag) / 10
				l1 = l1.Next
				temp = temp.Next
				if l1 == nil {
					if flag == 0 {
						temp.Next = nil
					} else {
						temp.Next = &ListNode{
							Val:  1,
							Next: nil,
						}
					}
					break
				}
			}
			break
		}

		temp.Next = &ListNode{
			Val:  (l1.Val + l2.Val + flag) % 10,
			Next: nil,
		}
		flag = (l1.Val + l2.Val + flag) / 10
		l1 = l1.Next
		l2 = l2.Next
		temp = temp.Next
		if l1 == nil && l2 == nil {
			if flag == 0 {
				temp.Next = nil
			} else {
				temp.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			break
		}

	}
	return l.Next
}

//试题地址：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	subMaxLen := 0

	i := 0
	j := 0
	for {
		j++
		if i >= len(s) || j >= len(s) {
			fmt.Println(s[i:j])
			if j-i > subMaxLen {
				subMaxLen = j - i
			}
			break
		}

		if !strings.Contains(s[i:j], string(s[j])) {
			continue
		} else {
			fmt.Println(s[i:j], "i=", i, "j=", j)
			if j-i > subMaxLen {
				subMaxLen = j - i
			}
			i = i + startLocation(s[i:j], string(s[j])) + 1
		}
	}
	return subMaxLen
}

func startLocation(str, s string) int {
	for i := 0; i < len(str); i++ {
		if string(str[i]) == s {
			return i
		}
	}
	return len(str)
}
