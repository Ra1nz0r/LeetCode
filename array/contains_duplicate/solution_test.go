package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "example 2",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "example 3",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsDuplicate(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
