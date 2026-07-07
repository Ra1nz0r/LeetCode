package longest_consecutive_sequence

func LongestConsecutive(nums []int) int {
	seen := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		seen[v] = struct{}{}
	}

	longestLen := 0

	for v := range seen {
		if _, ok := seen[v-1]; ok {
			continue
		}

		length := 1
		current := v

		for {
			if _, ok := seen[current+1]; !ok {
				break
			}

			length++
			current++
		}

		if length > longestLen {
			longestLen = length
		}
	}

	return longestLen
}
