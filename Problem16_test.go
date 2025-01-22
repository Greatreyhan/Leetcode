package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNodeP16 struct {
	Val  int
	Next *ListNodeP16
}

func deleteDuplicates(head *ListNodeP16) *ListNodeP16 {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input    *ListNodeP16
		expected []int
	}{
		{
			input:    &ListNodeP16{1, &ListNodeP16{1, &ListNodeP16{2, nil}}},
			expected: []int{1, 2},
		},
		{
			input:    &ListNodeP16{1, &ListNodeP16{1, &ListNodeP16{2, &ListNodeP16{3, &ListNodeP16{3, nil}}}}},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		result := deleteDuplicates(tt.input)
		var resultSlice []int
		for result != nil {
			resultSlice = append(resultSlice, result.Val)
			result = result.Next
		}
		assert.Equal(t, tt.expected, resultSlice, "Expected %v but got %v", tt.expected, resultSlice)
	}
}
