package group_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "example 2",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "example 3",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GroupAnagrams(test.strs)
			assert.ElementsMatch(t, test.expected, result)
		})
	}
}
