package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := listLen(l1)
	num2 := listLen(l2)
	num := max(num1, num2) + 1 //最高位可能进位
	array1 := listToArray(l1, num, num1)
	array2 := listToArray(l2, num, num2)
	for i := num - 1; i >= 0; i-- {
		if array1[i]+array2[i] >= 10 {
			array1[i] = (array1[i] + array2[i]) % 10
			array1[i-1]++
			continue
		}
		array1[i] = array1[i] + array2[i]
	}

	l := &ListNode{
		Val:  array1[0],
		Next: nil,
	}

	temp := l
	for i := 1; i < num; i++ {
		temp.Next = &ListNode{
			Val:  array1[i],
			Next: nil,
		}
		temp = temp.Next
	}
	if array1[0] == 0 {
		l = l.Next
	}
	return l
}

func listToArray(l *ListNode, n, m int) []int {
	numArray := make([]int, n)
	for i := n - m; i < n; i++ {
		numArray[i] = l.Val
		l = l.Next
	}
	return numArray
}

func listLen(l *ListNode) int {
	num := 0
	temp := l
	for temp.Next != nil {
		num++
		temp = temp.Next
	}
	return num
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
