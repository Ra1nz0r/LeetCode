package rotate_array

import "fmt"

func Rotate(nums []int, k int) {
	k %= len(nums)

	reverse(nums, 0, len(nums)-1)

	reverse(nums, 0, k-1)

	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
