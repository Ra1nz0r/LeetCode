package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "example 2",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "example 3",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MajorityElementBoyerMooreType(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
