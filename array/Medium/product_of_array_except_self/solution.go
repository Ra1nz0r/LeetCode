package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	suffix := 1
	for i := len(result) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductExceptSelf1(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	suffix := make([]int, len(nums))
	suffix[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	result := make([]int, len(nums))

	for k := range result {
		result[k] = prefix[k] * suffix[k]
	}

	return result
}
