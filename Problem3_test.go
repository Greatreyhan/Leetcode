package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func romanToInt(s string) int {
	set := map[string]int{"": 0, "I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	last_index := ""
	total := 0
	for _, c := range s {

		fmt.Printf("char : %c %v-%v %v \n", c, set[last_index], set[string(c)], total)

		if set[string(c)] > 0 {

			if set[string(c)] > set[last_index] && last_index != "" {
				total -= 2 * set[last_index]
				total += set[string(c)]
			} else {
				total += set[string(c)]
				last_index = string(c)
			}

		}

	}
	return total
}

func Test3(t *testing.T) {
	test3 := []struct {
		x        string
		expected int
	}{
		{x: "III", expected: 3},
		{x: "LVIII", expected: 58},
		{x: "MCMXCIV", expected: 1994},
	}

	for _, tt := range test3 {
		result := romanToInt(tt.x)
		assert.Equal(t, tt.expected, result)
	}
}
