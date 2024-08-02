/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/linked-list-cycle/description/
*/
package linked_list_cycle

import (
	"leetcode/utils"
)

/*
HasCycle reports whether the linked list starting at head contains a cycle.
HasCycle assumes that:
the number of the nodes in the list is in the range [0, 10^4],
and -10^5 <= Node.val <= 10^5.
*/
func HasCycle(head *utils.ListNode) bool {

	encountered_nodes := make(map[*utils.ListNode]bool)
	cursor := head

	for cursor != nil {
		_, exists := encountered_nodes[cursor]
		if exists {
			return true
		}
		encountered_nodes[cursor] = true
		cursor = cursor.Next
	}

	return false
}
