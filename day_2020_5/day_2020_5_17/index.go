package main

import "fmt"

func main() {
	// nums := [][]int{
	// 	{3, 0},
	// 	{1, 6},
	// 	{5, 1},
	// 	{0, 1},
	// 	{5, 6},
	// 	{2, 5},
	// 	{2, 0},
	// }

	nums := [][]int{{1, 0}, {0, 1}}
	fmt.Println(findOrder(2, nums))
}

//试题地址：https://leetcode-cn.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	order := []int{}

	maporder := map[int]bool{}
	for i := 0; i < len(prerequisites); i++ {
		maporder[i] = true
	}

	for {
		temp := 0
		for i := 0; i < len(prerequisites); i++ {
			if maporder[i] {
				temp = prerequisites[i][1]
				break
			}
		}
		tmp := 0
		for {
			flag := false
			for i := 0; i < len(prerequisites); i++ {
				if maporder[i] && prerequisites[i][0] == temp {
					tmp++
					flag = true
					temp = prerequisites[i][1]
				}
			}
			if !flag {
				break
			}
			if tmp >= len(prerequisites) { //出现循环 直接retuen
				return []int{}
			}

		}

		count := 0
		for i := 0; i < len(prerequisites); i++ {
			if prerequisites[i][1] == temp {
				maporder[i] = false
			}
			if maporder[i] {
				count++
			}
		}
		order = append(order, temp)
		if count == 0 {
			break
		}
	}

	for i := 0; i < numCourses; i++ { //未出现的课程随意安排
		falg := false
		for j := 0; j < len(order); j++ {
			if i == order[j] {
				falg = true
				break
			}
		}
		if !falg {
			order = append(order, prerequisites[i][0])
		}
	}

	return order
}
