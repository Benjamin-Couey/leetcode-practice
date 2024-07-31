package average_of_levels_in_binary_tree

import (
	"leetcode/utils"
)

func AverageOfLevels(root *utils.TreeNode) []float64 {
	sums, counts := recursiveAverageOfLevels(root, []int{0}, []int{0}, 0)

	averages := []float64{}

	for index := range sums {
		average := float64(sums[index]) / float64(counts[index])
		averages = append(averages, average)
	}

	return averages
}

func recursiveAverageOfLevels(root *utils.TreeNode, sums []int, counts []int, index int) ([]int, []int) {
	if root == nil {
		return sums, counts
	}

	sums[index] += root.Val
	counts[index]++

	if root.Left != nil || root.Right != nil {
		sums = append(sums, 0)
		counts = append(counts, 0)

		sums, counts = recursiveAverageOfLevels(root.Left, sums, counts, index+1)
		sums, counts = recursiveAverageOfLevels(root.Right, sums, counts, index+1)
	}

	return sums, counts
}
