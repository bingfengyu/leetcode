package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//试题地址：https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	list := []int{}
	flag := true
	recursive(root, &list, &flag)
	// fmt.Println(list)
	return flag
}

func recursive(root *TreeNode, list *[]int, flag *bool) {
	if root == nil || !(*flag) {
		return
	}

	recursive(root.Left, list, flag)
	if len(*list) == 0 {
		*list = append(*list, root.Val)
	} else if root.Val > (*list)[len(*list)-1] {
		*list = append(*list, root.Val)
	} else {
		*flag = false
	}
	recursive(root.Right, list, flag)
}
