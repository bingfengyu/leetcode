package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//试题地址：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/submissions/
func levelOrder(root *TreeNode) [][]int {
	levelTree := [][]int{}
	levelNode := [][]*TreeNode{}
	if root == nil {
		return levelTree
	}
	tmp := []*TreeNode{}
	tmpNode := []int{}
	tmp = append(tmp, root)
	tmpNode = append(tmpNode, root.Val)
	levelNode = append(levelNode, tmp)
	levelTree = append(levelTree, tmpNode)
	for {
		tmp = []*TreeNode{}
		tmpNode = []int{}
		for _, v := range levelNode[len(levelNode)-1] {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
				tmpNode = append(tmpNode, v.Left.Val)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
				tmpNode = append(tmpNode, v.Right.Val)
			}
		}
		if len(tmp) == 0 {
			break
		}
		levelNode = append(levelNode, tmp)
		levelTree = append(levelTree, tmpNode)
	}
	return levelTree
}
