package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
	pos := 0
	for i, e := range nums {
		fmt.Println(i, e, target, pos)
		if e <= target {
			pos = i + 1
		}
		if target == e {
			return i
		}
	}
	return pos
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
	}

	for _, tt := range tests {
		result := searchInsert(tt.nums, tt.target)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input haystack=%q and needle=%q", tt.expected, result, tt.nums, tt.target)
	}
}
