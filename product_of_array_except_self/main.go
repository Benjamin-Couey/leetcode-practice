package product_of_array_except_self

/* Assumes that:
Needs to be implemented so it runs in  O(n) time and without using the division
operation.

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer. */
func ProductExceptSelf(nums []int) []int {
	prefix_products := make([]int, len(nums))
	product := 1
	for index, num := range nums {
		product *= num
		prefix_products[ index ] = product
	}

	suffix_products := make([]int, len(nums))
	product = 1
	for index:=len(nums)-1; index >= 0; index-- {
		product *= nums[index]
		suffix_products[ index ] = product
	}

	return_slice := make([]int, len(nums))

	for index, _ := range return_slice {
		if index == 0 {
			return_slice[ index ] = suffix_products[ index + 1 ]
		} else if index == len(return_slice) - 1 {
			return_slice[ index ] = prefix_products[ index - 1 ]
		} else {
			return_slice[ index ] = prefix_products[ index - 1 ] * suffix_products[ index + 1 ]
		}
	}

	return return_slice
}

/* Alternative that uses O(1) extra space by using the necessary nums and return
slice.
Arguably still O(n) space complexity since if you are willing to modify nums,
there's no reason you can't put the result there, making the return slice an
extra space allocation. */
func AltProductExceptSelf(nums []int) []int {

	suffix_products := make([]int, len(nums))
	product := 1
	for index:=len(nums)-1; index >= 0; index-- {
		product *= nums[index]
		suffix_products[ index ] = product
	}

	product = 1
	for index, num := range nums {
		product *= num
		nums[ index ] = product
	}

	for index, _ := range suffix_products {
		if index == 0 {
			suffix_products[ index ] = suffix_products[ index + 1 ]
		} else if index == len(suffix_products) - 1 {
			suffix_products[ index ] = nums[ index - 1 ]
		} else {
			suffix_products[ index ] = nums[ index - 1 ] * suffix_products[ index + 1 ]
		}
	}

	return suffix_products
}
