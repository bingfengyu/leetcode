package main

import "fmt"

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		if this.min[len(this.min)-1] > x {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, this.min[len(this.min)-1])
		}
	}
	fmt.Println(this.stack, this.min)
}

func (this *MinStack) Pop() {
	this.min = this.min[1:]
	this.stack = this.stack[1:]
	fmt.Println(this.stack, this.min)
}

func (this *MinStack) Top() int {
	return this.stack[0]
}

func (this *MinStack) GetMin() int {
	return this.min[0]
}

//试题地址：https://leetcode-cn.com/problems/min-stack/
func main() {
	var minStack MinStack
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
