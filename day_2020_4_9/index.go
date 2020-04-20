package main

import "fmt"

func main() {
	// generateParenthesis(7)
	for _, v := range generateParenthesis(7) {
		fmt.Println(v)
	}
}

//试题链接：https://leetcode-cn.com/problems/generate-parentheses/submissions/
func generateParenthesis(n int) []string {
	parenthesis := []string{}
	parenthesisMap := map[string]bool{}

	startString := "()"
	orgnizeParenthesis(startString, &parenthesisMap, n-1, &parenthesis)

	return parenthesis
}

func orgnizeParenthesis(str string, parenthesisMap *map[string]bool, n int, parenthesis *[]string) {
	if (*parenthesisMap)[str] { //剪枝
		return
	} else {
		(*parenthesisMap)[str] = true
		if n == 0 { //最后结果
			(*parenthesis) = append((*parenthesis), str)
			return
		}
	}

	for i := 0; i <= len(str); i++ {
		tempStr1 := str[0:i] + "(" + str[i:]
		for j := i + 1; j <= len(tempStr1); j++ {
			tempStr2 := tempStr1[0:j] + ")" + tempStr1[j:]
			orgnizeParenthesis(tempStr2, parenthesisMap, n-1, parenthesis)
		}
	}
}
