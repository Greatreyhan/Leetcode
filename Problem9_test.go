package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{haystack: "sadbutsad", needle: "sad", expected: 0},
		{haystack: "leetcode", needle: "leeto", expected: -1},
		{haystack: "hello", needle: "ll", expected: 2},
		{haystack: "aaaaa", needle: "bba", expected: -1},
		{haystack: "", needle: "", expected: 0},
		{haystack: "a", needle: "a", expected: 0},
	}

	for _, tt := range tests {
		result := strStr(tt.haystack, tt.needle)
		assert.Equal(t, tt.expected, result, "Expected %v but got %v for input haystack=%q and needle=%q", tt.expected, result, tt.haystack, tt.needle)
	}
}
