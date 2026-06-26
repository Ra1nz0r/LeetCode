package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "no zeroes",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all zeroes",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "zeroes at the end",
			input:    []int{1, 2, 0, 0},
			expected: []int{1, 2, 0, 0},
		},
		{
			name:     "consecutive zeroes",
			input:    []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			MoveZeroes(test.input)
			assert.Equal(t, test.expected, test.input)
		})
	}
}
