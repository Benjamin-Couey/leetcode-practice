/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/majority-element/description/
*/
package majority_element

import (
	"errors"
)

/*
MajorityElement returns the majority element in nums. An error is returned
if nums is empty or does not contain a majority_element.
MajorityElement assumes that:
a majority element always exists in nums,
n == nums.length,
1 <= n <= 5 * 10^4,
and -10^9 <= nums[i] <= 10^9.
*/
func MajorityElement(nums []int) (int, error) {
	if len(nums) < 1 {
		return 0, errors.New("list is empty")
	}

	counts := make(map[int]int)

	for _, value := range nums {
		_, exists := counts[value]
		if !exists {
			counts[value] = 1
		} else {
			counts[value] += 1
		}
		if counts[value] > len(nums)/2 {
			return value, nil
		}
	}

	return 0, errors.New("list does not contain a majority element")
}
