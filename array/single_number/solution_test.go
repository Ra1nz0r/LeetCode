package single_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SingleNumber(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
