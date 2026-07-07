package longest_consecutive_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "example 3",
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "example 4",
			nums:     []int{1, 100, 6, 101, 102, 99},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LongestConsecutive(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}

}
