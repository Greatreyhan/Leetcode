package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x
	var result int

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
func TestMySqrt(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{x: 4, expected: 2},
		{x: 8, expected: 2},
	}

	for _, tt := range tests {
		result := mySqrt(tt.x)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input x=%q", tt.expected, result, tt.x)
	}
}
