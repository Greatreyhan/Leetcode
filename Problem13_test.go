package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carry = sum / 2
	}

	if carry > 0 {
		result = fmt.Sprintf("%d%s", carry, result)
	}

	return result
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{a: "11", b: "1", expected: "100"},
		{a: "1010", b: "1011", expected: "10101"},
	}

	for _, tt := range tests {
		result := addBinary(tt.a, tt.b)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input a=%q and b=%q", tt.expected, result, tt.a, tt.b)
	}
}
