package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAres(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "example 2",
			nums:     []int{1, 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxArea(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
