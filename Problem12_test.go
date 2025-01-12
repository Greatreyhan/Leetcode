package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plusOne(digits []int) []int {
	rem := 1
	for i, _ := range digits {
		res := digits[len(digits)-1-i] + rem
		if res > 9 {
			rem = 1
			digits[len(digits)-1-i] = 0
			continue
		} else {
			digits[len(digits)-1-i] = res
			rem = 0
			break
		}
	}
	if rem != 0 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits   []int
		expected []int
	}{
		{digits: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{digits: []int{9}, expected: []int{1, 0}},
	}

	for _, tt := range tests {
		result := plusOne(tt.digits)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input haystack=%q and needle=%q", tt.expected, result, tt.digits, tt.expected)
	}
}
