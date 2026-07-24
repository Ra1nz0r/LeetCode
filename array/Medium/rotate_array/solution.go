package rotate_array

import "fmt"

func Rotate(nums []int, k int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	left = 0
	right = k - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	left = k
	right = len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	fmt.Println(nums)
}
