package binary_search

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Search1(nums []int, target int) int {
	candidate := 0

	for len(nums) > 1 && candidate != target {
		candidate = nums[(len(nums) / 2)]

		if candidate != target && candidate < target {
			nums = nums[(len(nums) / 2):]
		} else {
			nums = nums[:(len(nums) / 2)]
		}

		if target == candidate {
			return candidate
		}
	}

	return -1
}
