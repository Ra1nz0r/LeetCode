package maximum_subarray

func MaxSubArray(nums []int) int {
	current := nums[0]
	maximum := nums[0]

	for i := 1; i < len(nums); i++ {
		if current < 0 {
			current = nums[i]
		} else {
			current += nums[i]
		}

		if current > maximum {
			maximum = current
		}
	}

	fmt.Println("hello")

	return maximum
}
