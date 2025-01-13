package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addBinary(a string, b string) string {
	anum, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		panic(err)
	}
	bnum, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	return strconv.FormatInt(anum+bnum, 2)
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
