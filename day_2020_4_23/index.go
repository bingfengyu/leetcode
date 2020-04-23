package main

import "fmt"

func main() {
	fmt.Println(waysToChange(30))
}

//试题地址：https://leetcode-cn.com/problems/coin-lcci/
func waysToChange(n int) int {
	if n == 0 {
		return 1
	}
	dpArray := make([][]int, 4)
	for i := 0; i < 4; i++ {
		dpArray[i] = make([]int, n+1)
	}
	array := []int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 {
				dpArray[i][j] = 1
				continue
			}
			if j < array[i] {
				dpArray[i][j] = dpArray[i-1][j]
			} else if j == array[i] {
				dpArray[i][j] = dpArray[i-1][j] + 1
			} else if j%5 == 0 {
				dpArray[i][j] = dpArray[i-1][j] + dpArray[i][j-array[i]]
			} else {
				dpArray[i][j] = dpArray[i][j-1]
			}
		}
	}
	fmt.Println(dpArray)
	return dpArray[3][n] % 1000000007
}
