package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func indexOf(c string, arr []string) int {
	for i, a := range arr {
		if c == a {
			return i
		}
	}

	return -1
}

func isValid(s string) bool {
	stack := []string{}
	open := []string{"(", "{", "["}
	close := []string{")", "}", "]"}
	for _, c := range s {
		isOpen := indexOf(string(c), open)
		// fmt.Printf("%v - %v - %v \n", string(c), stack, isOpen)

		// Handle Open
		if isOpen >= 0 {
			stack = append(stack, string(c))
		} else {
			// Handle Close
			if len(stack) > 0 {
				// fmt.Printf("Comparing %v - %v\n", indexOf(string(c), close), indexOf(stack[len(stack)-1], open))
				// check last stack
				if indexOf(string(c), close) == indexOf(stack[len(stack)-1], open) {
					// pop from last stack
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func Test5(t *testing.T) {
	test5 := []struct {
		x        string
		expected bool
	}{
		{x: "()", expected: true},
		{x: "()[]{}", expected: true},
		{x: "(]", expected: false},
		{x: "([])", expected: true},
	}

	for _, tt := range test5 {
		fmt.Printf("Test %s is %v \n", tt.x, tt.expected)
		result := isValid(tt.x)
		assert.Equal(t, tt.expected, result)
	}
}
