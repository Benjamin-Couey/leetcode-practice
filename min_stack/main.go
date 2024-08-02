/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/min-stack/description/
*/
package min_stack

/*
MinStack is an integer stack which implements push, pop, top, and retreiving
the minimum element in constant time.
The zero value for MinStack is an empty stack ready to use.
MinStack assumes that:
it is given values such that -2^31 <= val <= 2^31 - 1,
the methods pop, top and getMin operations will always be called on non-empty stacks,
and at most 3 * 10^4 calls will be made to Push, Pop, Top, and GetMin.
*/
type MinStack struct {
	stack []int
	min   []int
}

/* Constructor returns zero value for MinStack. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

/* Push val on top of MinStack this. */
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	min_val := val
	last_min_index := len(this.min) - 1
	if last_min_index >= 0 && this.min[last_min_index] < min_val {
		min_val = this.min[last_min_index]
	}
	this.min = append(this.min, min_val)
}

/* Pop remove top value of MinStack this. Does not return the removed value. */
func (this *MinStack) Pop() {
	last_index := len(this.stack) - 1
	this.stack = this.stack[:last_index]
	this.min = this.min[:last_index]
}

/* Top returns the top value of MinStack this. */
func (this *MinStack) Top() int {
	last_index := len(this.stack) - 1
	return this.stack[last_index]
}

/* GetMin returns the minimum value in MinStack this. */
func (this *MinStack) GetMin() int {
	last_index := len(this.min) - 1
	return this.min[last_index]
}
