package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Rotate(test.nums, test.k)
			assert.Equal(t, test.expected, test.nums)
		})
	}
}
