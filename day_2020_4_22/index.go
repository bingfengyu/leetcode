package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//试题地址：https://leetcode-cn.com/problems/binary-tree-right-side-view/solution/
func rightSideView(root *TreeNode) []int {
	treeMap := [][]TreeNode{}
	temp := []TreeNode{}
	temp = append(temp, (*root))
	treeMap = append(treeMap, temp)
	i := 0
	for {
		temp = []TreeNode{}
		for _, v := range treeMap[i] {
			if v.Left != nil {
				temp = append(temp, (*v.Left))
			}
			if v.Right != nil {
				temp = append(temp, (*v.Right))
			}
		}
		if len(temp) != 0 {
			treeMap = append(treeMap, temp)
			i++
			continue
		}
		break
	}
	result := []int{}
	for _, v := range treeMap {
		result = append(result, v[len(v)-1].Val)
	}
	return result
}
