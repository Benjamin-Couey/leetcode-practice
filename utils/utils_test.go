package utils

import (
	"fmt"
	"testing"
)

func TestSliceEqual(t *testing.T) {
	testcases := []struct {
		a, b []int
		solution bool
	}{
    { []int{ 1, 2, 3 }, []int{ 1, 2, 3 }, true },
		{ []int{ 1, 2, 4 }, []int{ 1, 2, 3 }, false },
		{ []int{ 1, 2 }, []int{ 1, 2, 3 }, false },
		{ []int{ 1, 2, 3, 0, 0, 0 }, []int{ 1, 2, 2, 3, 5, 6 }, false },
		{ []int{}, []int{}, true },
  }

	for _, testcase := range testcases {
		result := SliceEqual( testcase.a, testcase.b )

		if result != testcase.solution {
			t.Errorf(
				"SliceEqual: %t, want %t for slices %v and %v",
				result,
				testcase.solution,
				testcase.a,
				testcase.b,
			)
		}
	}
}

func TestSortFirstKInts(t *testing.T) {
	testcases := []struct {
		x, sorted_x []int
		k int
	}{
		{ []int{ 3, 2, 2, 3 }, []int{ 2, 3, 2, 3 }, 2 },
		{ []int{ 3, 2, 2, 3 }, []int{ 3, 2, 2, 3 }, 10 },
		{ []int{ 3, 2, 2, 3 }, []int{ 3, 2, 2, 3 }, 0 },
		{ []int{ 0, 1, 2, 2, 3, 0, 4, 2 }, []int{ 0, 1, 2, 2, 3, 0, 4, 2 }, 5 },
	}

	for _, testcase := range testcases {
		SortFirstKInts( testcase.x, testcase.k )
			if !SliceEqual( testcase.x, testcase.sorted_x ) {
			t.Errorf("SortFirstKInts: %v, want %v", testcase.x, testcase.sorted_x)
		}
	}
}

func TestIsSameLinkedList(t *testing.T) {

	head_a := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{ 3, nil },
		},
	}

	head_b := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{ 3, nil },
		},
	}

	head_c := &ListNode{
		1,
		&ListNode{ 2, nil },
	}

	head_d := &ListNode{
		3,
		&ListNode{
			2,
			&ListNode{ 1, nil },
		},
	}

	head_e := &ListNode{ 1, nil }
	head_e_a := &ListNode{ 2, nil }
	head_e_b := &ListNode{ 3, nil }

	head_e.Next = head_e_a
	head_e_a.Next = head_e_b
	head_e_b.Next = head_e

	head_f := &ListNode{ 1, nil }
	head_f_a := &ListNode{ 2, nil }
	head_f_b := &ListNode{ 3, nil }

	head_f.Next = head_f_a
	head_f_a.Next = head_f_b
	head_f_b.Next = head_f

	head_g := &ListNode{ 1, nil }
	head_g_a := &ListNode{ 2, nil }
	head_g_b := &ListNode{ 3, nil }

	head_g.Next = head_g_a
	head_g_a.Next = head_g_b
	head_g_b.Next = head_g_a

	testcases := []struct {
		p, q *ListNode
		solution bool
		summary string
	}{
		{ head_a, head_b, true, "head_a and head_b" },
		{ head_a, head_c, false, "head_a and head_c" },
		{ head_b, head_d, false, "head_b and head_d" },
		{ head_c, head_d, false, "head_c and head_d" },
		{ head_a, head_e, false, "head_a and head_e" },
		{ head_e, head_f, true, "head_e and head_f" },
		{ head_e, head_g, false, "head_e and head_g" },
		{ nil, nil, true, "nil and nil" },
		{ head_a, nil, false, "head_a and nil" },
	}

	for _, testcase := range testcases {
		same := IsSameLinkedList( testcase.p, testcase.q )

		if same != testcase.solution {
			t.Errorf("IsSameLinkedList: for %v returned %v, want %v", testcase.summary, same, testcase.solution)
		}
	}
}

func TestSliceToLinkedList(t *testing.T) {

	head_a := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{ 3, nil },
		},
	}

	head_b := &ListNode{
		3,
		&ListNode{
			2,
			&ListNode{ 1, nil },
		},
	}

	testcases := []struct {
		list []int
		solution *ListNode
	}{
		{ []int{ 1, 2, 3 }, head_a },
		{ []int{ 3, 2, 1 }, head_b },
		{ []int{}, nil },
		{ nil, nil },
	}

	for _, testcase := range testcases {
		head := SliceToLinkedList( testcase.list )

		if !IsSameLinkedList( head, testcase.solution ) {
			t.Errorf("SliceToLinkedList: returned incorrect linked list for list %v", testcase.list)
		}
	}
}

func TestSliceToLinkedListWithCycle(t *testing.T) {

	head_a := &ListNode{ 1, nil }
	head_a_b := &ListNode{ 2, nil }
	head_a_c := &ListNode{ 3, nil }
	head_a_d := &ListNode{ 4, nil }

	head_a.Next = head_a_b
	head_a_b.Next = head_a_c
	head_a_c.Next = head_a_d
	head_a_d.Next = head_a_b

	head_b := &ListNode{ 1, nil }
	head_b.Next = head_b

	head_c := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{ 3, nil },
		},
	}

	testcases := []struct {
		list []int
		cycle_index int
		solution *ListNode
	}{
		{ []int{ 1, 2, 3, 4 }, 1, head_a },
		{ []int{ 1 }, 0, head_b },
		{ []int{ 1, 2, 3 }, -1, head_c },
		{ []int{}, 0, nil },
		{ nil, 0, nil },
	}

	for _, testcase := range testcases {
		head := SliceToLinkedListWithCycle( testcase.list, testcase.cycle_index )

		if !IsSameLinkedList( head, testcase.solution ) {
			t.Errorf(
				"SliceToLinkedListWithCycle: returned incorrect linked list for list %v and cycle index %v",
				testcase.list,
				testcase.cycle_index,
			)
		}
	}
}

func TestIsSameTree(t *testing.T) {

	root_a := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{ 3, nil, nil },
	}

	root_b := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{ 3, nil, nil },
	}

	root_c := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		nil,
	}

	root_d := &TreeNode{
		1,
		nil,
		&TreeNode{ 2, nil, nil },
	}

	root_e := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{ 1, nil, nil },
	}

	root_f := &TreeNode{
		1,
		&TreeNode{ 1, nil, nil },
		&TreeNode{ 2, nil, nil },
	}

	testcases := []struct {
		p, q *TreeNode
		solution bool
		summary string
	}{
		{ root_a, root_b, true, "root_a and root_b" },
		{ root_c, root_d, false, "root_c and root_d" },
		{ root_e, root_f, false, "root_e and root_f" },
		{ nil, nil, true, "nil and nil" },
		{ root_a, nil, false, "root_a and nil" },
	}

	for _, testcase := range testcases {
		same := IsSameTree( testcase.p, testcase.q )

		if same != testcase.solution {
			t.Errorf("IsSameTree: for %v returned %v, want %v", testcase.summary, same, testcase.solution)
		}
	}
}

func TestLevelOrderToTree(t *testing.T) {

	root_a := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{ 3, nil, nil },
	}

	root_b := &TreeNode{
		1,
		nil,
		&TreeNode{ 2, nil, nil },
	}

	root_c := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{
			3,
			&TreeNode{ 4, nil, nil },
			&TreeNode{ 5, nil, nil },
		},
	}

	root_d := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{ 4, nil, nil },
			&TreeNode{ 5, nil, nil },
		},
		&TreeNode{
			3,
			&TreeNode{ 6, nil, nil },
			&TreeNode{ 7, nil, nil },
		},
	}

	testcases := []struct {
		list []int
		solution *TreeNode
	}{
		{ []int{}, nil },
		{ []int{ 1, 2, 3 }, root_a },
		{ []int{ 1, 0, 2 }, root_b },
		{ []int{ 1, 2, 3, 0, 0, 4, 5 }, root_c },
		{ []int{ 1, 2, 3, 4, 5, 6, 7 }, root_d },
	}

	for _, testcase := range testcases {
		root := LevelOrderToTree( testcase.list )
		if !IsSameTree( root, testcase.solution ) {
			t.Errorf( "LevelOrderToTree: Returned incorrect tree for level order %v\n", testcase.list )
		}
	}
}

func TestNillableLevelOrderToTree(t *testing.T) {

	root_a := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{ 3, nil, nil },
	}

	root_b := &TreeNode{
		1,
		nil,
		&TreeNode{ 2, nil, nil },
	}

	root_c := &TreeNode{
		1,
		&TreeNode{ 2, nil, nil },
		&TreeNode{
			3,
			&TreeNode{ 4, nil, nil },
			&TreeNode{ 5, nil, nil },
		},
	}

	root_d := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{ 4, nil, nil },
			&TreeNode{ 5, nil, nil },
		},
		&TreeNode{
			3,
			&TreeNode{ 6, nil, nil },
			&TreeNode{ 7, nil, nil },
		},
	}

	testcases := []struct {
		list []NilInt
		solution *TreeNode
	}{
		{ []NilInt{}, nil },
		{ []NilInt{ NewInt(1), NewInt(2), NewInt(3) }, root_a },
		{ []NilInt{ NewInt(1), NewNil(), NewInt(2) }, root_b },
		{ []NilInt{ NewInt(1), NewInt(2), NewInt(3), NewNil(), NewNil(), NewInt(4), NewInt(5) }, root_c },
		{ []NilInt{ NewInt(1), NewInt(2), NewInt(3), NewInt(4), NewInt(5), NewInt(6), NewInt(7) }, root_d },
	}

	for _, testcase := range testcases {
		root := NillableLevelOrderToTree( testcase.list )
		if !IsSameTree( root, testcase.solution ) {
			t.Errorf( "NillableLevelOrderToTree: Returned incorrect tree for level order %v\n", testcase.list )
		}
	}
}

func TestSliceInTree(t *testing.T) {
	testcases := []struct {
		nums, tree_list []int
		solution bool
	}{
		{ []int{ 1, 2, 3 }, []int{ 1, 2, 3 }, true },
		{ []int{ 3, 2, 1 }, []int{ 1, 2, 3 }, true },
		{ []int{ 1, 2, 3, 4 }, []int{ 1, 2, 3 }, false },
		{ []int{ 1, 2, 3 }, []int{ 1, 2, 3, 4, 5, 0, 0 }, false },
		{ []int{}, []int{}, true },
		{ []int{ 1 }, []int{}, false },
		{ []int{}, []int{ 1 }, false },
	}

	for _, testcase := range testcases {
		root := LevelOrderToTree( testcase.tree_list )
		fmt.Printf("%v\n", testcase.nums)
		result := SliceInTree( testcase.nums, root )
		fmt.Printf("%v\n", testcase.nums)
		if result != testcase.solution {
			t.Errorf(
				"SliceInTree: for slice %v and tree %v returned %v, want %v",
				testcase.nums,
				testcase.tree_list,
				result,
				testcase.solution)
		}
	}
}
