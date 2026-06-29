package product_of_array_except_self

import "fmt"

func ProductExceptSelf(nums []int) []int {
	mltpnTotal := 1

	for _, v := range nums {
		if v == 0 {
			continue
		} else {
			mltpnTotal *= v
		}
	}

	fmt.Println(mltpnTotal)

	for k := range nums {
		if nums[k] == 0 {
			nums[k] = mltpnTotal / 1
		} else {
			nums[k] = mltpnTotal / nums[k]
		}
	}

	return nums
}
