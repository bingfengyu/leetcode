package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("  hello     world!  "))
	fmt.Println(reverseWords("  sa dsf rger  er    "))
	fmt.Println(reverseWords("      "))
}

//试题地址：https://leetcode-cn.com/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	strList := strings.Fields(s)
	if len(strList) == 0 {
		return ""
	}
	reverseS := ""
	for i := len(strList) - 1; i > 0; i-- {
		reverseS += strList[i] + " "
	}
	reverseS += strList[0]

	return reverseS
}
