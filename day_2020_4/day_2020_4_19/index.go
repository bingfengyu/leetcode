package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getMaxRepetitions("cba", 5, "abc", 2))
}

//试题地址：https://leetcode-cn.com/problems/count-the-repetitions/
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if strings.Contains(s1, s2) {
		if strings.Repeat(s2, len(strings.Split(s1, s2))-1) == s1 {
			return n1 / n2 * (len(strings.Split(s1, s2)) - 1)
		}
	}
	if n1/n2 > 1 || n2 > n1 {
		s1Len := strings.Repeat(s1, n1)
		s2Len := strings.Repeat(s2, n2)
		fmt.Println(s1Len, s2Len)
		return repetitions(s1Len, s2Len)
	}
	return repetitions(s1, s2)
}

func repetitions(s1Len, s2Len string) int {
	j := 0
	k := 0
	for i := 0; i < len(s1Len); i++ {
		if s1Len[i] == s2Len[j] {
			if j == len(s2Len)-1 {
				fmt.Println(string(s1Len[i]), string(s2Len[j]), i, j, k)
				k++
				j = 0
				continue
			}
			fmt.Println(string(s1Len[i]), string(s2Len[j]), i, j, k)
			j++
		}
		fmt.Println(string(s1Len[i]), string(s2Len[j]), i, j, k)
	}
	return k
}
