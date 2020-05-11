package main

import "fmt"

func main() {
	// matrix := [][]byte{{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'}}
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}}
	fmt.Println(maximalSquare1(matrix))
}

//试题地址：https://leetcode-cn.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int { //贪心算法
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])

	temp := 0
	if n > m {
		temp = m
	} else {
		temp = n
	}
	// fmt.Println(n, m, temp)

	for i := temp; i >= 0; i-- {
		for j := 0; j < n-i; j++ {
			for k := 0; k < m-i; k++ {
				// fmt.Println(j, k, i)
				// getMatrix(matrix, j, k, i)
				if getMatrix(matrix, j, k, i) {
					return (i + 1) * (i + 1)
				}
			}
		}
	}
	return 0
}

func getMatrix(maxtric [][]byte, i, j, k int) bool {
	for temp1 := i; temp1 <= i+k; temp1++ {
		for temp2 := j; temp2 <= j+k; temp2++ {
			// fmt.Print(maxtric[temp1][temp2])
			if maxtric[temp1][temp2] == 48 {
				return false
			}
		}
		// fmt.Println()
	}
	// fmt.Println("------------")
	return true
}

func maximalSquare1(matrix [][]byte) int { //dp
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = int(matrix[i][j]) - 48
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}
	return maxLen * maxLen
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
