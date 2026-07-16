package three_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "example 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "example 4",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ThreeSum(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
