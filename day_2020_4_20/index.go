package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	islandnum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				islandDfs(i, j, n, m, grid)
				islandnum++
			}
		}
	}
	return islandnum
}

func islandDfs(i, j, n, m int, grid [][]byte) {
	if safelocation(i, j, n, m, grid) == '0' {
		return
	}
	grid[i][j] = '0'
	islandDfs(i-1, j, n, m, grid)
	islandDfs(i+1, j, n, m, grid)
	islandDfs(i, j-1, n, m, grid)
	islandDfs(i, j+1, n, m, grid)
}

func safelocation(i, j, n, m int, grid [][]byte) byte {
	if i < 0 {
		return '0'
	}
	if i >= n {
		return '0'
	}
	if j < 0 {
		return '0'
	}
	if j >= m {
		return '0'
	}
	return grid[i][j]
}
