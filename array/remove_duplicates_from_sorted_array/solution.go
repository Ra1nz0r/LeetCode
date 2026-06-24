package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	writePos := 0

	for _, v := range nums {
		if writePos == 0 || nums[writePos-1] != v {
			nums[writePos] = v
			writePos++
		}
	}

	return writePos
}
