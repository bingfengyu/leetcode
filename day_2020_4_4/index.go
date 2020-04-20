package main

import (
	"fmt"
	"strings"
)

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// fmt.Println(trap(height))

	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()()"))
	fmt.Println(longestValidParentheses("()(())("))
	fmt.Println(longestValidParentheses(")("))
	fmt.Println(longestValidParentheses("(()))())("))
	fmt.Println(longestValidParentheses("))))))()((()(((()((()((())()((()))()()(()()))))()())()))())())))(((()()()()(()()))()((()))))()(()()(((()())(((())(((())((()))))()(()(())(()(()(((((())(())))(()))())(((((())))))))()(())((((((()())()))))(()))(((()()(()(())()((())))()(())))())()())((((()))))()()((((((())()((((()(())((((()()())(())()())))()(()((())))))))()())((((()))())())))()))((())))()((()(())((()((())))((())((()())()))))())))())(((((((()))(((())(((()((()))(((()(())))((((()()())))))))())()))(()))())))()(()))((())()))((())))((())(()())(((()((()()))))()()(())))))()))))()((()((()(((()())(((())))(()(())())))((((((()()()))))((()))))(()(())()(()((())((()))()()()()(()((()())((((())()((()()))()))))()(((()())())))(()(((()()))()())))()()))))((()((((()(()))()(()()(((()()()()())((())()(((())(())))))()))))()())()()())(())))((())(()))(((())()))))))())()())((()((()()(())))())((()(((()(((((()(())(())()()))()))(()((((((((()(()()))()(()(()))(())(())())))))(())))))))()))))()())()()(((())()))((()))))))(())((()))(()((()()())())((()))))((()()())(()(((()))))(())(((()(()()()()((())())(())()))()())))())))()(((()))(()))()))()((((()()()(()(()))(((())()()(((())()((()(())((())())(((()()))((()())))())(()(())(()()))(((()(())(())))(()())(()()()()(()((()()()))(())(()(((()())))()((((())()())((()(()(())))(()()(()(()))))))(()())((()))()))()(()))((()()))))()()()))(())(())(())))()(()))(()()(((()((((())())))(()(())(()()))))))((()((()()))())))((((((())()))())(())())(()()(()(()()))()))((()((())(()((()(())))()((()))()(((()()()())()((())))(()()()()(()((()()((()((()()()(()(((())))))(())(()()(()())((((()()()()((()(())()(((()))(()(())(((()())())((()()))("))
}

//试题地址：https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	i := 0
	j := len(height) - 1
	sum := 0

	//查找左右界限
	var temp int
	for {
		for m := i; m < len(height)-1; m++ {
			if height[m] > height[m+1] {
				temp = m
				break
			}
			if m == len(height)-1 {
				temp = m
			}
		}
		i = temp

		for m := j; m > 1; m-- {
			if height[m] > height[m-1] {
				temp = m
				break
			}
			if m == 1 {
				temp = m
			}
		}
		j = temp

		//上移x轴
		if height[i] > height[j] {
			sum += list(&height, i, j, height[j])
			j -= 2
		} else if height[i] < height[j] {
			sum += list(&height, i, j, height[i])
			i += 2
		} else if height[i] == height[j] {
			sum += list(&height, i, j, height[i])
			i += 2
			j -= 2
		}
		if i >= j || i == j-1 {
			break
		}
	}
	return sum
}

func value(n, m int) int {
	temp := 0
	if n-m >= 0 {
		temp = n - m
	}
	return temp
}

func list(height *[]int, i, j, k int) int {
	temp := 0
	for m := i; m <= j; m++ {
		temp += value(k, (*height)[m])
		(*height)[m] = value((*height)[m], k)
	}
	return temp
}

func list_print(height []int) {
	for i := 0; i < len(height); i++ {
		fmt.Print(height[i], " ")
	}
	fmt.Println()
}

//试题地址：https://leetcode-cn.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var temp string
	for {
		temp = re_s(s)
		if s == temp {
			break
		}
		s = temp
	}
	max := 0
	for _, s1 := range strings.Split(s, "(") {
		for _, s2 := range strings.Split(s1, ")") {
			if len(s2) > max {
				max = len(s2)
			}
		}
	}
	return max
}

//处理字符串
func re_s(s string) string {
	i := 0
	for {
		j := i + 1
		for {
			if j >= len(s) || s[j] != 'o' {
				break
			}
			j++
		}
		if i >= len(s) || j >= len(s) {
			break
		}
		if s[i] == '(' && s[j] == ')' {
			s = s[0:i] + "o" + s[i+1:]
			s = s[0:j] + "o" + s[j+1:]
			i = j
			continue
		}
		i++
	}
	return s
}
