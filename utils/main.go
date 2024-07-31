package utils

import (
	"errors"
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/* Assumes that:
Node.val is unique for each node. Node.val is also the unique index of the node.
There are no repeated edges and no self-loops in the graph.
The Graph is connected and all nodes can be visited starting from the given node. */
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func SliceEqual[V comparable](a, b []V) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func MatrixEqual[V comparable](a, b [][]V) bool {
	if len(a) != len(b) {
		return false
	}
	for row_i, _ := range a {
		if len(a[row_i]) != len(b[row_i]) {
			return false
		}
		for col_i, _ := range a[row_i] {
			if a[row_i][col_i] != b[row_i][col_i] {
				return false
			}
		}
	}
	return true
}

func SliceEqualAnyOrder[V comparable](a, b []V) bool {
	if len(a) != len(b) {
		return false
	}

	mutable_b := make([]V, len(b))
	copy(mutable_b, b)

	for a_i, _ := range a {
		found_match := false
		b_i := 0
		for !found_match && b_i < len(mutable_b) {
			if a[a_i] == mutable_b[b_i] {
				mutable_b = append(mutable_b[:b_i], mutable_b[b_i+1:]...)
				found_match = true
			}
			b_i++
		}
		if !found_match {
			return false
		}
	}
	return true
}

func MatrixEqualAnyOrder[V comparable](a, b [][]V) bool {
	if len(a) != len(b) {
		return false
	}

	mutable_b := make([][]V, len(b))
	for i := range b {
		mutable_b[i] = make([]V, len(b[i]))
		copy(mutable_b[i], b[i])
	}

	for a_i, _ := range a {
		// Find row in b that matches row at a_i
		found_match := false
		b_i := 0
		for !found_match && b_i < len(mutable_b) {
			// Skip past rows of different length
			if len(a[a_i]) == len(mutable_b[b_i]) {
				// Compare rows
				if SliceEqualAnyOrder(a[a_i], mutable_b[b_i]) {
					found_match = true
					mutable_b = append(mutable_b[:b_i], mutable_b[b_i+1:]...)
				}
			}
			b_i++
		}

		if !found_match {
			return false
		}
	}
	return true
}

func SortFirstKInts(x []int, k int) {
	if k > len(x) || k < 0 {
		return
	}
	first_k := x[0:k]
	rest := x[k:]
	sort.Ints(first_k)
	x = append(first_k, rest...)
}

func IsSameLinkedList(p *ListNode, q *ListNode) bool {

	encountered_p_nodes := make(map[*ListNode]bool)
	encountered_q_nodes := make(map[*ListNode]bool)

	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		_, p_exists := encountered_p_nodes[p]
		_, q_exists := encountered_q_nodes[q]

		if p_exists || q_exists {
			if p.Val != q.Val {
				return false
			} else {
				return true
			}
		}

		encountered_p_nodes[p] = true
		encountered_q_nodes[q] = true

		p = p.Next
		q = q.Next
	}

	if p == q {
		return true
	} else {
		return false
	}
}

func SliceToLinkedList(x []int) *ListNode {
	if len(x) < 1 {
		return nil
	}

	head := ListNode{x[0], nil}
	cursor := &head
	for _, val := range x[1:] {
		new_node := ListNode{val, nil}
		cursor.Next = &new_node
		cursor = cursor.Next
	}

	return &head
}

func SliceToLinkedListWithCycle(x []int, y int) *ListNode {
	if len(x) < 1 {
		return nil
	}

	var head *ListNode = nil
	var cursor *ListNode = nil
	var cycle_node *ListNode = nil
	for index, val := range x {
		new_node := ListNode{val, nil}
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
		if index == len(x)-1 {
			cursor.Next = cycle_node
		}
	}

	return head
}

/* Will stop upon encountering a cycle. */
func LinkedListToSlice(head *ListNode) []int {
	return_slice := make([]int, 0)
	encountered_nodes := make(map[*ListNode]bool)
	for head != nil && !encountered_nodes[head] {
		return_slice = append(return_slice, head.Val)
		encountered_nodes[head] = true
		head = head.Next
	}
	return return_slice
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

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

/* Given a level order list of integer values, builds a tree with those values.
Node which are nil should be indicated with a 0. Trailings 0s may be
excluded. Assumes that nodes in the tree do not contain the value 0. */
func LevelOrderToTree(x []int) *TreeNode {
	if len(x) < 1 {
		return nil
	}
	return recursiveLevelOrderToTree(x, 0)
}

func recursiveLevelOrderToTree(x []int, index int) *TreeNode {
	if index >= len(x) || index < 0 {
		return nil
	}
	if x[index] == 0 {
		return nil
	}
	root := TreeNode{x[index], nil, nil}

	left_index := index*2 + 1
	right_index := index*2 + 2

	root.Left = recursiveLevelOrderToTree(x, left_index)
	root.Right = recursiveLevelOrderToTree(x, right_index)

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

func NillableLevelOrderToTree(x []NilInt) *TreeNode {
	if len(x) < 1 {
		return nil
	}
	return recursiveNillableLevelOrderToTree(x, 0)
}

func recursiveNillableLevelOrderToTree(x []NilInt, index int) *TreeNode {
	if index >= len(x) || index < 0 {
		return nil
	}
	// Get the Value interface from the NilInt struct and assert its type is int
	value, ok := x[index].Value().(int)
	if !ok {
		return nil
	}
	root := TreeNode{value, nil, nil}

	left_index := index*2 + 1
	right_index := index*2 + 2

	root.Left = recursiveNillableLevelOrderToTree(x, left_index)
	root.Right = recursiveNillableLevelOrderToTree(x, right_index)

	return &root
}

/* Given a slice of values and the root of a tree, returns true if every value
in the slice appears as exactly one node in the tree. */
func SliceInTree(nums []int, root *TreeNode) bool {

	/* recursiveSliceInTree is going to remove values from nums while it runs.
	Since slices are pointers, to avoid whatever slice was passed to SliceInTree
	also being mutated, we'll make a new slice for this function to use. */
	mutable_slice := make([]int, 0)
	mutable_slice = append(mutable_slice, nums...)
	if len(mutable_slice) == 0 || root == nil {
		return len(mutable_slice) == 0 && root == nil
	}

	var tree_in_nums bool
	tree_in_nums, mutable_slice = recursiveSliceInTree(mutable_slice, root)
	return tree_in_nums && len(mutable_slice) == 0
}

func recursiveSliceInTree(nums []int, root *TreeNode) (bool, []int) {
	found_root_val := false
	for index, num := range nums {
		if num == root.Val {
			nums = append(nums[:index], nums[index+1:]...)
			found_root_val = true
		}
	}

	if !found_root_val {
		return false, nums
	}

	left_bool := true
	if root.Left != nil {
		left_bool, nums = recursiveSliceInTree(nums, root.Left)
	}
	right_bool := true
	if root.Right != nil {
		right_bool, nums = recursiveSliceInTree(nums, root.Right)
	}
	return left_bool && right_bool, nums
}

func IsSameGraph(p *GraphNode, q *GraphNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		} else {
			return false
		}
	}

	p_node_to_connected := make(map[int][]int)

	// Build map of p's nodes to the nodes they connect to
	buildNodeToConnected(p, p_node_to_connected)

	q_node_to_connected := make(map[int][]int)

	// Build map of q's nodes to the nodes they connect to
	buildNodeToConnected(q, q_node_to_connected)

	if len(p_node_to_connected) != len(q_node_to_connected) {
		return false
	}

	// Compare maps
	for key, p_connected := range p_node_to_connected {
		q_connected, exists := q_node_to_connected[key]
		if !exists {
			return false
		}
		if !SliceEqual(p_connected, q_connected) {
			return false
		}
	}

	return true
}

func buildNodeToConnected(node *GraphNode, node_to_connected map[int][]int) {
	_, exists := node_to_connected[node.Val]
	if !exists {
		connected_vals := make([]int, 0)
		for _, connected_node := range node.Neighbors {
			if connected_node != nil {
				connected_vals = append(connected_vals, connected_node.Val)
			}
		}

		node_to_connected[node.Val] = connected_vals

		for _, connected_node := range node.Neighbors {
			if connected_node != nil {
				buildNodeToConnected(connected_node, node_to_connected)
			}
		}
	}
}

func AdjacencyToGraph(list [][]int) (*GraphNode, error) {

	if len(list) < 1 {
		return nil, nil
	}

	nodes := make([]*GraphNode, 0)

	for index := 1; index <= len(list); index++ {
		new_node := &GraphNode{index, make([]*GraphNode, 0)}
		nodes = append(nodes, new_node)
	}

	for index, adjacencies := range list {
		existing_adjacencies := make(map[int]bool)
		for _, adjacency := range adjacencies {
			_, exists := existing_adjacencies[adjacency]
			if adjacency > len(list) || adjacency < 1 {
				return nil, errors.New("Node index out of range")
			} else if index == adjacency-1 {
				return nil, errors.New("Adjacency defined self loop")
			} else if exists {
				return nil, errors.New("Adjacency defined repeated edge")
			} else {
				nodes[index].Neighbors = append(nodes[index].Neighbors, nodes[adjacency-1])
				existing_adjacencies[adjacency] = true
			}
		}
	}

	return nodes[0], nil
}
