package linked_list_cycle

import (
	"leetcode/utils"
)

/* Assumes that:
The number of the nodes in the list is in the range [0, 10^4].
-10^5 <= Node.val <= 10^5. */
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
