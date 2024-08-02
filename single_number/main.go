/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/single-number/description/
*/
package single_number

/* SingleNumber returns the only element of nums which appears only once.
SingleNumber assumes that:
1 <= nums.length <= 3 * 10^4,
-3 * 10^4 <= nums[i] <= 3 * 10^4,
and each element in the array appears twice except for one element which appears
only once. */
func SingleNumber(nums []int) int {
	// Initialize integer mask to 0.
	var mask int

	/*
	Bitwise XOR all the numbers in nums. This will cause all the elements that
	appear twice to cancel each other out, leaving mask as the one element which
	was only XORed once.
	*/
	for _, num := range nums {
		mask = mask ^ num
	}

	return mask
}
