/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/merge-two-sorted-lists/description/
*/
package merge_two_sorted_lists

import (
	"leetcode/utils"
)

/*
MergeTwoLists merges the linked lists starting at list1 and list2 and returns
the resulting sorted linked list.
MergeTwoLists assumes that:
two nodes with the same value from different original lists are interchangeable
in the merged list,
the number of nodes in both lists is in the range [0, 50],
-100 <= Node.val <= 10,
and both list1 and list2 are sorted in non-decreasing order.
*/
func MergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {

	var head *utils.ListNode = nil
	var cursor *utils.ListNode = nil
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {

			if head == nil && cursor == nil {
				head = list2
				cursor = head
			} else {
				cursor.Next = list2
				cursor = cursor.Next
			}
			list2 = list2.Next

		} else {

			if head == nil && cursor == nil {
				head = list1
				cursor = head
			} else {
				cursor.Next = list1
				cursor = cursor.Next
			}
			list1 = list1.Next
		}
	}

	var remainder *utils.ListNode = nil
	if list1 != nil {
		remainder = list1
	} else {
		remainder = list2
	}

	if head == nil && cursor == nil {
		head = remainder
	} else {
		cursor.Next = remainder
	}

	return head
}
