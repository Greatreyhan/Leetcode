package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums      []int
		val       int
		expected  []int
		expectedK int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expected: []int{2, 2}, expectedK: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expected: []int{0, 1, 3, 0, 4}, expectedK: 5},
	}

	for _, tt := range tests {
		original := make([]int, len(tt.nums))
		copy(original, tt.nums)

		k := removeElement(tt.nums, tt.val)

		// Verify length
		assert.Equal(t, tt.expectedK, k, "Expected k=%v but got %v", tt.expectedK, k)

		// Verify first k elements
		assert.ElementsMatch(t, tt.expected, tt.nums[:k], "Expected nums=%v but got %v for input %v", tt.expected, tt.nums[:k], original)
	}
}
