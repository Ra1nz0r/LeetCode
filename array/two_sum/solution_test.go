package two_sum

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {

		result := TwoSum(test.nums, test.target)

		assert.Equal(t, result, test.expected, fmt.Sprintf(
			"got %v expected %v",
			result,
			test.expected),
		)
	}
}
