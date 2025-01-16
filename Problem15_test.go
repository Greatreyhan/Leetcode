package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func climbStairs(n int) int {
	memo := make(map[int]int) // Memoization to store results

	var helper func(int) int
	helper = func(steps int) int {
		if steps <= 2 {
			return steps // Base cases: 1 way for 1 step, 2 ways for 2 steps
		}
		if val, found := memo[steps]; found {
			return val // Return cached result if already computed
		}
		memo[steps] = helper(steps-1) + helper(steps-2) // Store and return result
		return memo[steps]
	}

	return helper(n)
}
func TestClimbStairs(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{x: 2, expected: 2},
		{x: 3, expected: 3},
		{x: 4, expected: 5},
	}

	for _, tt := range tests {
		result := climbStairs(tt.x)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input x=%q", tt.expected, result, tt.x)
	}
}
