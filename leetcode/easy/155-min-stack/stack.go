package minstack

import (
	"math"
)

type MinStack struct {
	data     []int
	top      int
	min      int
	minStack []int
	minTop   int
}

func Constructor() MinStack {
	return MinStack{
		data:     make([]int, 0),
		top:      -1,
		min:      math.MaxInt64,
		minStack: make([]int, 0),
		minTop:   -1}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.min = x
		this.minTop++
		if this.minTop < len(this.minStack) {
			this.minStack[this.minTop] = x
		} else {
			this.minStack = append(this.minStack, x)
		}
	}

	this.top++
	if this.top < len(this.data) {
		this.data[this.top] = x
	} else {
		this.data = append(this.data, x)
	}
}

func (this *MinStack) Pop() {
	if this.data[this.top] == this.min {
		this.minTop--
		if this.minTop < 0 {
			this.min = math.MaxInt64
		} else {
			this.min = this.minStack[this.minTop]
		}
	}
	this.top--
}

func (this *MinStack) Top() int {
	return this.data[this.top]
}

func (this *MinStack) GetMin() int {
	return this.min
}
