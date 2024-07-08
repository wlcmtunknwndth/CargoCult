package twoPointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	given    []int
	expected []int
}{
	{
		given:    []int{0, 42, 1, 31, 123, 4, 5, 3, 1},
		expected: []int{0, 1, 1, 3, 4, 5, 31, 42, 123},
	},
	{
		given:    []int{3, 1, 2, 1, 15, 14, 13},
		expected: []int{1, 1, 2, 3, 13, 14, 15},
	},
	{
		given:    []int{0, 1, 2, 3, 4},
		expected: []int{0, 1, 2, 3, 4},
	},
}

func TestTwoPointers(t *testing.T) {
	for _, tc := range testCases {
		var tmp = make([]int, len(tc.given))
		copy(tmp, tc.given)
		QuickSortStart(tmp, func(i1 int, i2 int) int {
			if i1 > i2 {
				return -1
			} else if i1 == i2 {
				return 0
			} else {
				return 1
			}
		})

		assert.Equal(t, tmp, tc.expected)
	}
}
