package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func shortest_str(strs []string) int {
	l := 9999
	for _, s := range strs {
		if l > len(s) {
			l = len(s)
		}
	}
	return l
}

func longestCommonPrefix(strs []string) string {
	str := ""
	shrtest := shortest_str(strs)
	fmt.Println(shrtest)
	last_c := ""
	for i := 0; i < shrtest; i++ {
		for _, s := range strs {
			// fmt.Println(last_c)
			if last_c == "" {
				last_c = string(s[i])
			} else if last_c != string(s[i]) {
				last_c = ""
				break
			}
		}
		if last_c != "" {
			str += last_c
			last_c = ""
		} else {
			break
		}
	}
	// fmt.Println(str)
	return str
}

func Test4(t *testing.T) {
	test4 := []struct {
		x        []string
		expected string
	}{
		{x: []string{"flower", "flow", "flight"}, expected: "fl"},
		{x: []string{"dog", "racecar", "car"}, expected: ""},
		{x: []string{"a"}, expected: "a"},
		{x: []string{"", "b"}, expected: ""},
	}

	for _, tt := range test4 {
		result := longestCommonPrefix(tt.x)
		assert.Equal(t, tt.expected, result)
	}
}
