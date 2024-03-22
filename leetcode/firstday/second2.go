package main

import "math"

//可以利用一个辅助栈极大的剪短了程序运行的时间
type MinStack2 struct {
	stack  []int
	Fstack []int
}

/** initialize your data structure here. */
func Constructor2() MinStack2 {
	return MinStack2{
		stack:  []int{}, //初始化要加{}
		Fstack: []int{math.MaxInt64},
	}
}

func (this *MinStack2) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.Fstack[len(this.stack)-1]
	this.Fstack = append(this.Fstack, min(top, x))
}

func (this *MinStack2) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.Fstack = this.Fstack[:len(this.Fstack)-1]
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) Min() int {
	return this.Fstack[len(this.Fstack)-1]
}

func min(top int, x int) int {
	if top > x {
		return x
	}
	return top
}
