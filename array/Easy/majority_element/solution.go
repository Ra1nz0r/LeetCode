package majority_element

func MajorityElementBoyerMooreType(nums []int) int {
	candidate := 0
	count := 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if candidate == v {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func MajorityElementMapVersion(nums []int) int {
	seen := make(map[int]int, len(nums))
	threshold := len(nums)/2 + 1

	for _, v := range nums {
		seen[v]++

		if seen[v] >= threshold {
			return v
		}
	}

	return 0
}
