/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
*/
package remove_nth_node_from_end_of_list

import (
	"leetcode/utils"
)

/*
RemoveNthFromEnd returns the head of the linked list head after removing the
nth node from the end.
RemoveNthFromEnd assumes that:
the number of nodes in the list is sz,
1 <= sz <= 30,
0 <= ListNode.val <= 100,
and 1 <= n <= sz.
*/
func RemoveNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {

	pointer_list := make([]*utils.ListNode, 0)
	var cursor *utils.ListNode = nil
	var return_head *utils.ListNode = nil

	for head != nil {
		new_node := utils.ListNode{head.Val, nil}
		if return_head == nil {
			return_head = &new_node
			cursor = return_head
		} else {
			cursor.Next = &new_node
			cursor = cursor.Next
		}
		pointer_list = append(pointer_list, cursor)
		head = head.Next
	}

	index_to_remove := len(pointer_list) - n
	var right_node *utils.ListNode = nil
	right_index := index_to_remove + 1
	if right_index < len(pointer_list) {
		right_node = pointer_list[right_index]
	}
	left_index := index_to_remove - 1
	if index_to_remove > 0 {
		pointer_list[left_index].Next = right_node
	} else {
		return_head = right_node
	}

	return return_head
}
