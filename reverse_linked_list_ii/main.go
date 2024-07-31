package reverse_linked_list_ii

import (
	"leetcode/utils"
)

/* Assumes that:
The number of nodes in the list is n.
1 <= n <= 500.
-500 <= Node.val <= 500.
1 <= left <= right <= n. */
func ReverseBetween(head *utils.ListNode, left int, right int) *utils.ListNode {
	var cursor *utils.ListNode = nil
	var return_head *utils.ListNode = nil
	var last_forward *utils.ListNode = nil
	var first_reverse *utils.ListNode = nil
	var last_reverse *utils.ListNode = nil

	index := 1

	for head != nil {
		new_node := utils.ListNode{head.Val, nil}
		if index < left || index > right {
			if return_head == nil && last_reverse == nil {
				return_head = &new_node
				cursor = return_head
			} else {
				if first_reverse != nil && first_reverse.Next == nil {
					first_reverse.Next = &new_node
					cursor = first_reverse.Next
				} else {
					cursor.Next = &new_node
					cursor = cursor.Next
				}
			}
			last_forward = cursor
		} else {
			if first_reverse == nil {
				cursor = &new_node
				first_reverse = cursor
			} else {
				new_node.Next = cursor
				cursor = &new_node
			}
			if last_forward != nil {
				last_forward.Next = cursor
			}
			last_reverse = cursor
		}
		head = head.Next
		index++
	}

	if return_head == nil && last_reverse != nil {
		return_head = last_reverse
	}

	return return_head
}
