package clone_graph

import (
	"leetcode/utils"
)

func CloneGraph(node *utils.GraphNode) *utils.GraphNode {

	if node == nil {
		return nil
	}

	node_to_connected := make( map[int][]int )
	buildNodeToConnected( node, node_to_connected )

	val_to_pointer := make( map[int]*utils.GraphNode )

	for node_val, _ := range node_to_connected {
		new_node := &utils.GraphNode{ node_val, make( []*utils.GraphNode, 0 ) }
		val_to_pointer[ node_val ] = new_node
	}

	for node_val, connected_vals := range node_to_connected {
		node_to_update := val_to_pointer[ node_val ]
		for _, val := range connected_vals {
			node_to_update.Neighbors = append( node_to_update.Neighbors, val_to_pointer[ val ] )
		}
	}

	return val_to_pointer[ node.Val ]
}


func buildNodeToConnected( node *utils.GraphNode, node_to_connected map[int][]int ) {
	_, exists := node_to_connected[ node.Val ]
	if !exists {
		connected_vals := make( []int, 0 )
		for _, connected_node := range node.Neighbors {
			if connected_node != nil {
				connected_vals = append( connected_vals, connected_node.Val )
			}
		}

		node_to_connected[ node.Val ] = connected_vals

		for _, connected_node := range node.Neighbors {
			if connected_node != nil {
				buildNodeToConnected( connected_node, node_to_connected )
			}
		}
	}
}
