package best_time_to_buy_and_sell_stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "example 2",
			nums:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "example 3",
			nums:     []int{1, 5},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxProfit(test.nums)
			assert.Equal(t, test.expected, result)
		})
	}
}
