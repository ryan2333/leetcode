package main

type MyQueue struct {
	Input  []int
	Output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Input = append(this.Input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()
	x := this.Output[len(this.Output)-1]
	this.Output = this.Output[:len(this.Output)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Output) == 0 {
		for len(this.Input) > 0 {
			x := this.Input[len(this.Input)-1]
			this.Input = this.Input[:len(this.Input)-1]
			this.Output = append(this.Output, x)
		}
	}
	return this.Output[len(this.Output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Input) == 0 && len(this.Output) == 0
}

func main() {
	q := MyQueue{}
	q.Push(1)
	q.Push(2)
	q.Peek()
	q.Pop()
	q.Empty()
	q.Pop()
	q.Empty()
}
