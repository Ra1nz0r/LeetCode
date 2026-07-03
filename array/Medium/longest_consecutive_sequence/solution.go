package longest_consecutive_sequence

import (
	"fmt"
	"sort"
)

func LongestConsecutive(nums []int) int {
	sort.Ints(nums)

	fmt.Println("Sort nums", nums)

	seen := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := seen[v-1]; !ok {
			seen[v] = 1
			continue
		} else {
			seen[v-1] = seen[v+1]
			seen[v]++
		}
	}

	fmt.Println(seen)

	return 0
}
