package two_sum

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for k, v := range nums {
		needed := target - v

		if val, ok := seen[needed]; ok {
			return []int{val, k}
		}

		seen[nums[k]] = k
	}

	return nil
}
