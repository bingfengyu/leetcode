package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(movingCount(7, 2, 3))   //7
	fmt.Println(movingCount(11, 8, 16)) //88
	fmt.Println(movingCount(16, 8, 4))  //15
	fmt.Println(movingCount(38, 15, 9)) //135
}

//试题地址:https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/submissions/
//1、超过10位的问题
//2、能否到达的问题
func movingCount(m int, n int, k int) int {
	flag := make([][]int, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}
	sum := 1
	flag[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i-1 >= 0 && flag[i-1][j] == 1) ||
				(j-1 >= 0 && flag[i][j-1] == 1) {
				if get_sum(i)+get_sum(j) <= k {
					flag[i][j] = 1
					sum += 1
				}
			}
		}
	}
	return sum
}

// func movingCount(m int, n int, k int) int {
// 	sum := 0
// 	flag := ","
// 	dfs(0, 0, n, m, k, &sum, &flag)
// 	fmt.Println(flag)
// 	return sum
// }

//广度优先遍历
func dfs(i, j, n, m, k int, sum *int, flag *string) {
	if i >= n || j >= m {
		return
	}
	if get_sum(i)+get_sum(j) <= k {
		*sum += 1
		*flag += strconv.Itoa(i) + "*" + strconv.Itoa(j) + ","
		if !strings.Contains(*flag, ","+strconv.Itoa(i)+"*"+strconv.Itoa(j+1)+",") {
			dfs(i, j+1, n, m, k, sum, flag)
		}
		if !strings.Contains(*flag, ","+strconv.Itoa(i+1)+"*"+strconv.Itoa(j)+",") {
			dfs(i+1, j, n, m, k, sum, flag)
		}
	}
	return
}

//获取位数和
func get_sum(n int) int {
	temp := 0
	for {
		temp += n % 10
		n = n / 10
		if n == 0 {
			break
		}
	}
	return temp
}
