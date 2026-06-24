package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		uniqueNumsCnt int
		expected      []int
	}{
		{
			name:          "example 1",
			nums:          []int{1, 1, 2},
			uniqueNumsCnt: 2,
			expected:      []int{1, 2},
		},
		{
			name:          "example 2",
			nums:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			uniqueNumsCnt: 5,
			expected:      []int{0, 1, 2, 3, 4},
		},
		{
			name:          "no duplicates",
			nums:          []int{1, 2, 3},
			uniqueNumsCnt: 3,
			expected:      []int{1, 2, 3},
		},
		{
			name:          "all duplicates",
			nums:          []int{2, 2, 2},
			uniqueNumsCnt: 1,
			expected:      []int{2},
		},
		{
			name:          "with negative numbers",
			nums:          []int{-3, -3, -1, 0, 0, 2},
			uniqueNumsCnt: 4,
			expected:      []int{-3, -1, 0, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveDuplicates(test.nums)

			require.Equal(t, test.uniqueNumsCnt, result)
			assert.Equal(t, test.expected, test.nums[:result])
		})
	}
}
