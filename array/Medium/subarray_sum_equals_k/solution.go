package subarray_sum_equals_k

func SubarraySum(nums []int, k int) int {
	currentSum := 0
	answer := 0

	seenSums := make(map[int]int)
	seenSums[0] = 1

	for _, num := range nums {
		currentSum += num
		previousSum := currentSum - k

		answer += seenSums[previousSum]

		seenSums[currentSum]++
	}

	return answer
}
