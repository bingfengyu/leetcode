package main

import "fmt"

func main() {
	matrix := [][]int{{0, 0, 1, 1}, {1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 0, 1}}
	fmt.Println(updateMatrix(matrix))
}

//试题地址：https://leetcode-cn.com/problems/01-matrix/submissions/
func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	//初始化为最长路径
	matrixLen := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < m; j++ {
			temp = append(temp, m+n)
		}
		matrixLen = append(matrixLen, temp)
	}

	//正向
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				matrixLen[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if i != 0 && j == 0 {
				matrixLen[i][j] = min(matrixLen[i][j], matrixLen[i-1][j]+1)
				continue
			}
			if j != 0 && i == 0 {
				matrixLen[i][j] = min(matrixLen[i][j], matrixLen[i][j-1]+1)
				continue
			}
			matrixLen[i][j] = min(matrixLen[i][j], min(matrixLen[i-1][j]+1, matrixLen[i][j-1]+1))
		}
	}

	//反向
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				matrixLen[i][j] = 0
				continue
			}
			if i == n-1 && j == m-1 {
				continue
			}
			if i != n-1 && j == m-1 {
				matrixLen[i][j] = min(matrixLen[i][j], matrixLen[i+1][j]+1)
				continue
			}
			if j != m-1 && i == n-1 {
				matrixLen[i][j] = min(matrixLen[i][j], matrixLen[i][j+1]+1)
				continue
			}
			matrixLen[i][j] = min(matrixLen[i][j], min(matrixLen[i+1][j]+1, matrixLen[i][j+1]+1))
		}
	}

	return matrixLen
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

//试题地址：https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for k, v := range nums {
		numMap[v] = k
	}

	for k, v := range nums {
		if _, ok := numMap[target-v]; ok && numMap[target-v] != k {
			return []int{k, numMap[target-v]}
		}
	}
	return []int{}
}
