package min_stack

import (
	"reflect"
	"testing"
)

/* Constants used by MinStackOperation to specify which method should be
called on MinStack. */
const PUSH int = 0
const POP int = 1
const TOP int = 2
const GETMIN int = 3

/* MinStackOperation describes an operation that should be performed with a
MinStack, including the val passed to the method and an appropriate expected
solution. */
type MinStackOperation struct {
	val       int
	operation int
	solution  int
}

func TestConstructor(t *testing.T) {

	stack := MinStack{
		make([]int, 0),
		make([]int, 0),
	}

	// TODO: Find a more elegant way to check if result is of the type MinStack
	result := Constructor()
	if reflect.TypeOf(result) != reflect.TypeOf(stack) {
		t.Errorf("Constructor: did not return a MinStack")
	}
}

func TestMinStack(t *testing.T) {

	/* Define series of MinStackOperation to call to test MinStack.
	If operation is PUSH or POP then solution will be ignored as these functions do
	not return anything.
	If operation is POP, TOP, or GETMIN then val will be ignored as these functions
	do not accept parameters. */
	testcases := []struct {
		name       string
		operations []MinStackOperation
	}{
		{"First case", []MinStackOperation{
			{-2, PUSH, 0},
			{0, TOP, -2},
			{0, POP, 0},
		}},
		{"Second case", []MinStackOperation{
			{-2, PUSH, 0},
			{0, PUSH, 0},
			{-3, PUSH, 0},
			{0, GETMIN, -3},
			{0, POP, 0},
			{0, TOP, 0},
			{0, GETMIN, -2},
		}},
	}

	for _, testcase := range testcases {

		stack := Constructor()

		for index, operation := range testcase.operations {
			switch operation.operation {
			case PUSH:
				stack.Push(operation.val)
			case POP:
				stack.Pop()
			case TOP:
				result := stack.Top()
				if result != operation.solution {
					t.Errorf(
						"MinStack: Testcase %v Top operation %v returned %v, want %v",
						testcase.name,
						(index + 1),
						result,
						operation.solution,
					)
				}
			case GETMIN:
				result := stack.GetMin()
				if result != operation.solution {
					t.Errorf(
						"MinStack: Testcase %v GetMin operation %v returned %v, want %v",
						testcase.name,
						(index + 1),
						result,
						operation.solution,
					)
				}
			}
		}
	}
}
