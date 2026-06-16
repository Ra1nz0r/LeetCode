package move_zeroes

import "fmt"

func MoveZeroes(nums []int) {
	for k, v := range nums {
		if v == 0 {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums, 0)

			continue
		}
	}

	fmt.Println(nums)
}
