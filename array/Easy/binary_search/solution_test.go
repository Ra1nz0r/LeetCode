package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{-1, 0, 3, 5, 9, 12, 15, 16, 17, 18},
			target:   12,
			expected: 5,
		},
		{
			name:     "example 2",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Search(test.nums, test.target)
			assert.Equal(t, test.expected, result)
		})
	}
}
