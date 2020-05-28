package main

import (
	"fmt"
	"strconv"
)

func main() {
	testing := []string{"2[20[bc]31[xy]]xd4[rt]", "3[a]20[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}
	for _, v := range testing {
		fmt.Println(decodeString(v))
	}
}

//试题地址：https://leetcode-cn.com/problems/decode-string/
func decodeString(s string) string {
	flag := true
	for flag {
		i := 0
		flag = false
		tmp := ""
		k := 0
		for ; k < len(s); k++ {
			if string(s[k]) == "[" {
				i = k
			}
			if string(s[k]) == "]" {
				flag = true
				j := getNumBegin(s, i)
				num, _ := strconv.Atoi(string(s[j:i]))
				fmt.Println(num)
				if j == 0 {
					tmp = getNumString(s[i+1:k], num) + s[k+1:]
				} else {
					tmp = s[0:j] + getNumString(s[i+1:k], num) + s[k+1:]
				}
				break
			}
		}
		if flag {
			s = tmp
		}
	}
	return s
}

func getNumBegin(s string, i int) int {
	k := i - 1
	for ; k >= 0; k-- {
		if string(s[k]) >= "0" && string(s[k]) <= "9" {
			continue
		}
		break
	}
	return k + 1
}

func getNumString(str string, n int) string {
	tmp := ""
	for i := 0; i < n; i++ {
		tmp += str
	}
	return tmp
}
