package subarray_sum_equals_k

import "fmt"

func SubarraySum(nums []int, k int) int {
	seen := make(map[int]int)

	for i := range nums {
		seen[i] = nums[i]
	}

	fmt.Println(seen)

	return 0
}
