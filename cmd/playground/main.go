package main

import (
	"leetcode/array/merge_sorted_array"
)

func main() {
	merge_sorted_array.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge_sorted_array.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
