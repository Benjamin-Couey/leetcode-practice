package majority_element

import (
	"errors"
)

/* Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array. */
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
