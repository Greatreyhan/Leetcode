package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums      []int
		expected  []int
		expectedK int
	}{
		{nums: []int{1, 1, 2}, expected: []int{1, 2}, expectedK: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: []int{0, 1, 2, 3, 4}, expectedK: 5},
	}

	for _, tt := range tests {
		original := make([]int, len(tt.nums))
		copy(original, tt.nums)

		k := removeDuplicates(tt.nums)

		assert.Equal(t, tt.expectedK, k, "Expected k=%v but got %v", tt.expectedK, k)
		assert.Equal(t, tt.expected, tt.nums[:k], "Expected nums=%v but got %v for input %v", tt.expected, tt.nums[:k], original)
	}
}
