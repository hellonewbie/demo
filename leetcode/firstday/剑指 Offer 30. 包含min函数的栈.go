package main

//剑指offer 30
//卡了的原因，三个点：1.没有考虑好范围2.return的错误使用3.没有考虑栈空的情况
type MinStack struct {
	stack []int
	top   *int
	tail  *int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.top = &this.stack[len(this.stack)-1]
	if len(this.stack) == 1 {
		this.tail = &this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) > 1 {
		this.top = &this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return *this.top
}

func (this *MinStack) Min() int {
	var min int
	min = *this.top
	for _, v := range this.stack {
		if min > v {
			min = v
		} else {
		}
	}
	return min
}
