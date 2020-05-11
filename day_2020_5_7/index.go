package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := strings.Split("014235", "10142")
	fmt.Println(len(s), s)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//试题地址：https://leetcode-cn.com/problems/subtree-of-another-tree/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if t == nil {
		return true
	}

	sString := ""
	recursive(s, &sString)
	tString := ""
	recursive(t, &tString)
	return strings.Contains(sString, tString)
}

func recursive(tree *TreeNode, s *string) {
	if tree == nil {
		return
	}
	if tree.Left != nil {
		recursive(tree.Left, s)
	} else {
		*s += "L"
	}
	*s += strconv.Itoa(tree.Val)
	if tree.Right != nil {
		recursive(tree.Right, s)
	} else {
		*s += "R"
	}
}
