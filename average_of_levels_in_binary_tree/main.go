/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
*/
package average_of_levels_in_binary_tree

import (
	"leetcode/utils"
)

/*
AverageOfLevels returns a slice of the average value of each level of tree
starting at root.
AverageOfLevels assumes that:
the number of nodes in the tree is in the range [1, 104],
and -2^31 <= TreeNode.val <= 2^31 - 1.
*/
func AverageOfLevels(root *utils.TreeNode) []float64 {
	sums, counts := recursiveSumOfLevels(root, []int{0}, []int{0}, 0)

	averages := []float64{}

	for index := range sums {
		average := float64(sums[index]) / float64(counts[index])
		averages = append(averages, average)
	}

	return averages
}

func recursiveSumOfLevels(root *utils.TreeNode, sums []int, counts []int, index int) ([]int, []int) {
	if root == nil {
		return sums, counts
	}

	sums[index] += root.Val
	counts[index]++

	if root.Left != nil || root.Right != nil {
		sums = append(sums, 0)
		counts = append(counts, 0)

		sums, counts = recursiveSumOfLevels(root.Left, sums, counts, index+1)
		sums, counts = recursiveSumOfLevels(root.Right, sums, counts, index+1)
	}

	return sums, counts
}
