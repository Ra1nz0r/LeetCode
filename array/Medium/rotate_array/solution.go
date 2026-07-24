package rotate_array

import "fmt"

func Rotate(nums []int, k int) {
	cp := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		cp[(i+k)%len(nums)] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = cp[i]
	}

	fmt.Println(nums)

}
