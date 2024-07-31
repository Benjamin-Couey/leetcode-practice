package add_two_numbers

import (
	"leetcode/utils"
)

/* Assumes that:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
The list represents a number that does not have leading zeros. */
func AddTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var return_list utils.ListNode
	var return_cursor *utils.ListNode
	l1_cursor := l1
	l2_cursor := l2
	carry := 0
	for l1_cursor != nil || l2_cursor != nil || carry > 0 {
		if return_cursor == nil {
			return_cursor = &return_list
		} else {
			new_node := utils.ListNode{0, nil}
			return_cursor.Next = &new_node
			return_cursor = return_cursor.Next
		}

		sum := 0
		if l1_cursor != nil {
			sum += l1_cursor.Val
			l1_cursor = l1_cursor.Next
		}
		if l2_cursor != nil {
			sum += l2_cursor.Val
			l2_cursor = l2_cursor.Next
		}
		if carry > 0 {
			sum += carry
			carry = 0
		}
		return_cursor.Val = sum % 10
		carry = sum / 10
	}

	return &return_list
}
