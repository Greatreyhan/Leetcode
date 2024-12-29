package leetcode

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := strconv.Itoa(x)
	for i := 0; i < len(arr); i++ {
		if string(arr[i]) != string(arr[len(arr)-i-1]) {
			return false
		}
	}
	return true
}

func Test2(t *testing.T) {
	tests := []struct {
		x        int
		expected bool
	}{
		{x: 121, expected: true},
		{x: -121, expected: false},
		{x: 10, expected: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("x=%v", tt.x), func(t *testing.T) {
			result := isPalindrome(tt.x)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("isPalindrome(%v) = %v; expected %v", tt.x, result, tt.expected)
			}
		})
	}
}
