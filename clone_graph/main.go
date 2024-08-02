/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/clone-graph/description/
*/
package clone_graph

import (
	"leetcode/utils"
)

/*
CloneGraph returns a deep copy of the graph reachable from node.
CloneGraph assumes that:
The number of nodes in the graph is in the range [0, 100],
and 1 <= GraphNode.val <= 100.
*/
func CloneGraph(node *utils.GraphNode) *utils.GraphNode {

	if node == nil {
		return nil
	}

	// Build map of node vals to connected vals.
	node_to_connected := make(map[int][]int)
	buildNodeToConnected(node, node_to_connected)

	// Build map of node vals to pointers.
	val_to_pointer := make(map[int]*utils.GraphNode)

	for node_val, _ := range node_to_connected {
		// Make new node for each val.
		new_node := &utils.GraphNode{node_val, make([]*utils.GraphNode, 0)}
		val_to_pointer[node_val] = new_node
	}

	// Connect new nodes according to node_to_connected.
	for node_val, connected_vals := range node_to_connected {
		node_to_update := val_to_pointer[node_val]
		for _, val := range connected_vals {
			node_to_update.Neighbors = append(node_to_update.Neighbors, val_to_pointer[val])
		}
	}

	return val_to_pointer[node.Val]
}

func buildNodeToConnected(node *utils.GraphNode, node_to_connected map[int][]int) {
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
