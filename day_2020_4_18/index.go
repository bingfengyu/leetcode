package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	maxarea := 0
	i := 0
	j := len(height) - 1
	for j > i {
		if maxarea < min(height[i], height[j])*(j-i) {
			maxarea = min(height[i], height[j]) * (j - i)
		}
		fmt.Println(maxarea)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return maxarea
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
