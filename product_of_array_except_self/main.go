/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/product-of-array-except-self/description/
*/
package product_of_array_except_self

/*
ProductExceptSelf returns a slice such that return_slice[i] is equal to the
product of all elements of nums except nums[i]. For the sake of the problem,
ProductExceptSelf is implemented so it runs in  O(n) time and without using the
division operation.
ProductExceptSelf assumes that:
2 <= nums.length <= 10^5,
-30 <= nums[i] <= 30,
and the product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/
func ProductExceptSelf(nums []int) []int {
	prefix_products := make([]int, len(nums))
	product := 1
	for index, num := range nums {
		product *= num
		prefix_products[index] = product
	}

	suffix_products := make([]int, len(nums))
	product = 1
	for index := len(nums) - 1; index >= 0; index-- {
		product *= nums[index]
		suffix_products[index] = product
	}

	return_slice := make([]int, len(nums))

	for index, _ := range return_slice {
		if index == 0 {
			return_slice[index] = suffix_products[index+1]
		} else if index == len(return_slice)-1 {
			return_slice[index] = prefix_products[index-1]
		} else {
			return_slice[index] = prefix_products[index-1] * suffix_products[index+1]
		}
	}

	return return_slice
}


/*
AltProductExceptSelf is an alternative implementation of ProductExceptSelf
which uses O(1) extra space instead of O(n), as defiend by the follow up to
the problem which states that "the output array does not count as extra space
for space complexity analysis".

Arguably, this implementation still uses O(n) extra space since if you are
willing to modify nums, there's no reason you can't put the result there,
making the return slice an extra space allocation.
*/
func AltProductExceptSelf(nums []int) []int {

	suffix_products := make([]int, len(nums))
	product := 1
	for index := len(nums) - 1; index >= 0; index-- {
		product *= nums[index]
		suffix_products[index] = product
	}

	product = 1
	for index, num := range nums {
		product *= num
		nums[index] = product
	}

	for index, _ := range suffix_products {
		if index == 0 {
			suffix_products[index] = suffix_products[index+1]
		} else if index == len(suffix_products)-1 {
			suffix_products[index] = nums[index-1]
		} else {
			suffix_products[index] = nums[index-1] * suffix_products[index+1]
		}
	}

	return suffix_products
}
