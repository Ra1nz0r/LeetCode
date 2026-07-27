package top_k_frequent_elements

func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for i := range nums {
		freq[nums[i]]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	res := make([]int, 0, k)

	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}
