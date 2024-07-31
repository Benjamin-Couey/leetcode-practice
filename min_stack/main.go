package min_stack

/* Assumes that:
-2^31 <= val <= 2^31 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 10^4 calls will be made to push, pop, top, and getMin. */
type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	min_val := val
	last_min_index := len(this.min) - 1
	if last_min_index >= 0 && this.min[last_min_index] < min_val {
		min_val = this.min[last_min_index]
	}
	this.min = append(this.min, min_val)
}

func (this *MinStack) Pop() {
	last_index := len(this.stack) - 1
	this.stack = this.stack[:last_index]
	this.min = this.min[:last_index]
}

func (this *MinStack) Top() int {
	last_index := len(this.stack) - 1
	return this.stack[last_index]
}

func (this *MinStack) GetMin() int {
	last_index := len(this.min) - 1
	return this.min[last_index]
}
