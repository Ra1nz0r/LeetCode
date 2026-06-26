package single_number

func SingleNumber(nums []int) int {
	result := 0

	for _, v := range nums {
		result ^= v
	}

	return result
}

// Обычное решение через map, затраты выше.
func SingleNumberMap(nums []int) int {
	m := make(map[int]int, len(nums))

	for _, v := range nums {
		m[v]++

	}

	for k := range m {
		if m[k] == 1 {
			return k
		}
	}

	return 0
}
