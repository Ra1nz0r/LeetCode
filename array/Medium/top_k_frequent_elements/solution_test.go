package top_k_frequent_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{10, 10, 10, 20, 20, 30},
			k:        2,
			expected: []int{10, 20},
		},
		{
			name:     "example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "example 3",
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:        2,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TopKFrequent(test.nums, test.k)
			assert.ElementsMatch(t, test.expected, result)
		})
	}
}
