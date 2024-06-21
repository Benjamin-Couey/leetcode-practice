package min_stack

import (
	"reflect"
	"testing"
)

const PUSH int = 0
const POP int = 1
const TOP int = 2
const GETMIN int = 3

type MinStackOperation struct {
	val int
	operation int
	solution int
}

func TestConstructor(t *testing.T) {

	stack := MinStack {
		make( []int, 0 ),
		make( []int, 0 ),
	}

	// TODO: Find a more elegant way to check if result is of the type RandomizedSet
	result := Constructor()
	if reflect.TypeOf( result ) != reflect.TypeOf( stack ) {
		t.Errorf( "Constructor: did not return a RandomizedSet" )
	}
}

func TestMinStack(t *testing.T) {

	testcases := []struct {
		name string
		operations []MinStackOperation
	} {
		{ "First case", []MinStackOperation{
			{ -2, PUSH, 0, },
			{ 0, POP, 0, },
		}, },
		{ "Second case", []MinStackOperation{
			{ -2, PUSH, 0, },
			{ 0, PUSH, 0, },
			{ -3, PUSH, 0, },
			{ 0, GETMIN, -3, },
			{ 0, POP, 0, },
			{ 0, TOP, 0, },
			{ 0, GETMIN, -2, },
		}, },
	}

	for _, testcase := range testcases {

		stack := Constructor()

		for index, operation := range testcase.operations {
			switch operation.operation {
				case PUSH:
					stack.Push( operation.val )
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
