package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	sp := strings.Split(s, " ")
	for i, _ := range sp {
		if len(sp[len(sp)-1-i]) > 0 {
			return len(sp[len(sp)-1-i])
		}
	}
	return 0
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{s: "Hello World", expected: 5},
		{s: "   fly me   to   the moon  ", expected: 4},
		{s: "luffy is still joyboy", expected: 6},
	}

	for _, tt := range tests {
		result := lengthOfLastWord(tt.s)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input haystack=%q and needle=%q", tt.expected, result, tt.s, tt.expected)
	}
}
