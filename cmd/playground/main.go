package main

import (
	"fmt"
	"leetcode/array/remove_duplicates_from_sorted_array"
)

func main() {
	fmt.Println(remove_duplicates_from_sorted_array.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(remove_duplicates_from_sorted_array.RemoveDuplicates([]int{1, 1, 2}))
}
