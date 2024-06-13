package utils

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val int
	Next *ListNode
}

func SliceEqual[V comparable]( a, b []V) bool {
	if len(a) != len(b) {
    return false
  }
	for i, _ := range(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func MatrixEqual[V comparable]( a, b [][]V) bool {
	if len(a) != len(b) {
		return false
	}
	for row_i, _ := range(a) {
		if len(a[row_i]) != len(b[row_i]) {
			return false
		}
		for col_i, _ := range(a[row_i]) {
			if a[row_i][col_i] != b[row_i][col_i] {
				return false
			}
		}
	}
	return true
}

func SortFirstKInts( x []int, k int ) {
	if k > len( x ) || k < 0 {
    return
  }
	first_k := x[0:k]
	rest := x[k:]
	sort.Ints( first_k )
	x = append( first_k, rest... )
}

func IsSameLinkedList(p *ListNode, q *ListNode) bool {

	encountered_p_nodes := make( map[*ListNode]bool )
	encountered_q_nodes := make( map[*ListNode]bool )

	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		_, p_exists := encountered_p_nodes[ p ]
		_, q_exists := encountered_q_nodes[ q ]

		if p_exists || q_exists {
			if p.Val != q.Val {
				return false
			} else {
				return true
			}
		}

		encountered_p_nodes[ p ] = true
		encountered_q_nodes[ q ] = true

		p = p.Next
		q = q.Next
	}

	if p == q {
		return true
	} else {
		return false
	}
}

func SliceToLinkedList( x []int ) *ListNode {
	if len(x) < 1 {
		return nil
	}

	head := ListNode{ x[0], nil }
	cursor := &head
	for _, val := range( x[1:] ) {
		new_node := ListNode{ val, nil }
		cursor.Next = &new_node
		cursor = cursor.Next
	}

	return &head
}

func SliceToLinkedListWithCycle( x []int, y int ) *ListNode {
	if len(x) < 1 {
		return nil
	}

	var head *ListNode = nil
	var cursor *ListNode = nil
	var cycle_node *ListNode = nil
	for index, val := range( x ) {
		new_node := ListNode{ val, nil }
		if head == nil && cursor == nil {
			head = &new_node
			cursor = head
		} else {
			cursor.Next = &new_node
			cursor = cursor.Next
		}
		if index == y {
			cycle_node = cursor
		}
		if index == len(x) - 1 {
			cursor.Next = cycle_node
		}
	}

	return head
}

func TraverseAndPrintTree(node *TreeNode) {
	if node == nil {
		return
	}
	left_val := 0
	if node.Left != nil {
		left_val = node.Left.Val
	}
	right_val := 0
	if node.Right != nil {
		right_val = node.Right.Val
	}
	fmt.Printf(
		"Node with value %v, left value %v, right value %v\n",
		node.Val,
		left_val,
		right_val,
	)
	TraverseAndPrintTree(node.Left)
	TraverseAndPrintTree(node.Right)
}


func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		} else {
			return false
		}
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree( p.Left, q.Left ) && IsSameTree( p.Right, q.Right )
}

/* Given a level order list of integer values, builds a tree with those values.
Node which are nil should be indicated with a 0. Trailings 0s may be
excluded. Assumes that nodes in the tree do not contain the value 0. */
func LevelOrderToTree( x []int ) *TreeNode {
	if len( x ) < 1 {
		return nil
	}
	return recursiveLevelOrderToTree( x, 0 )
}

func recursiveLevelOrderToTree( x []int, index int ) *TreeNode {
	if index >= len(x) || index < 0 {
		return nil
	}
	if x[ index ] == 0 {
		return nil
	}
	root := TreeNode{ x[index], nil, nil }

	left_index := index * 2 + 1
	right_index := index * 2 + 2

	root.Left = recursiveLevelOrderToTree( x, left_index )
	root.Right = recursiveLevelOrderToTree( x, right_index )

	return &root
}

type NilInt struct {
  value int
  null  bool
}

// Have NilInt implicitly implement the Value interface. interface{} specifies
// that the function returns an interface. In order to get an int, you will
// need to use a type assertion..
func (n *NilInt) Value() interface{} {
  if n.null {
    return nil
  }
  return n.value
}

func NewInt(x int) NilInt {
  return NilInt{x, false}
}

func NewNil() NilInt {
  return NilInt{0, true}
}

func NillableLevelOrderToTree( x []NilInt ) *TreeNode {
	if len( x ) < 1 {
		return nil
	}
	return recursiveNillableLevelOrderToTree( x, 0 )
}

func recursiveNillableLevelOrderToTree( x []NilInt, index int ) *TreeNode {
	if index >= len(x) || index < 0 {
		return nil
	}
	// Get the Value interface from the NilInt struct and assert its type is int
	value, ok := x[ index ].Value().(int)
	if !ok {
		return nil
	}
	root := TreeNode{ value, nil, nil }

	left_index := index * 2 + 1
	right_index := index * 2 + 2

	root.Left = recursiveNillableLevelOrderToTree( x, left_index )
	root.Right = recursiveNillableLevelOrderToTree( x, right_index )

	return &root
}

/* Given a slice of values and the root of a tree, returns true if every value
in the slice appears as exactly one node in the tree. */
func SliceInTree( nums []int, root *TreeNode ) bool {

	/* recursiveSliceInTree is going to remove values from nums while it runs.
	Since slices are pointers, to avoid whatever slice was passed to SliceInTree
	also being mutated, we'll make a new slice for this function to use. */
	mutable_slice := make([]int, 0)
	mutable_slice = append(mutable_slice, nums...)
	if len(mutable_slice) == 0 || root == nil {
		return len(mutable_slice) == 0 && root == nil
	}

	var tree_in_nums bool
	tree_in_nums, mutable_slice = recursiveSliceInTree( mutable_slice, root )
	return tree_in_nums && len(mutable_slice) == 0
}

func recursiveSliceInTree( nums []int, root *TreeNode ) ( bool, []int ) {
	found_root_val := false
	for index, num := range( nums ) {
		if num == root.Val {
			nums = append( nums[:index], nums[index+1:]... )
			found_root_val = true
		}
	}

	if !found_root_val {
		return false, nums
	}

	left_bool := true
	if root.Left != nil {
		left_bool, nums = recursiveSliceInTree( nums, root.Left )
	}
	right_bool := true
	if root.Right != nil {
		right_bool, nums = recursiveSliceInTree( nums, root.Right )
	}
	return left_bool && right_bool, nums
}
