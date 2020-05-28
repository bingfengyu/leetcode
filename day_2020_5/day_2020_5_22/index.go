package main

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//试题：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var tmp int
	for k, v := range inorder {
		if v == preorder[0] {
			tmp = k
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:tmp+1], inorder[0:tmp]),
		Right: buildTree(preorder[tmp+1:], inorder[tmp+1:]),
	}
}
