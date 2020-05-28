package main

import "fmt"

func main() {
	minDistance("zoologicoarchaeologist", "zoogeologist")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//试题地址：https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[len(lists)-1]
	}

	i := 0
	for {
		lists = append(lists, merge2Lists(lists[i], lists[i+1]))
		i += 2
		if i == len(lists)-1 {
			break
		}
	}

	return lists[len(lists)-1]
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	list := &ListNode{}
	temp := list
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			temp.Next = list2
			list2 = list2.Next
		} else {
			temp.Next = list1
			list1 = list1.Next
		}
		temp = temp.Next
	}
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}
	return list.Next
}

//试题地址：https://leetcode-cn.com/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	result := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = make([]int, m+1)
		if i == 0 {
			for j := 0; j <= m; j++ {
				result[i][j] = j
			}
		} else {
			result[i][0] = i
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				result[i][j] = result[i-1][j-1]
				continue
			}
			result[i][j] = min(min(result[i-1][j-1], result[i][j-1]), result[i-1][j]) + 1
		}
	}
	for _, v := range result {
		fmt.Println(v)
	}
	return result[n][m]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
