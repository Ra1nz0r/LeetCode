package move_zeroes

func MoveZeroes(nums []int) {
	writePos := 0

	for _, v := range nums {
		if v != 0 {
			nums[writePos] = v
			writePos++
		}
	}

	for i := writePos; i < len(nums); i++ {
		nums[i] = 0
	}
}
