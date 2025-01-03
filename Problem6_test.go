package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

// Helper function to compare two linked lists
func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func Test6(t *testing.T) {
	test6 := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{l1: []int{1, 2, 4}, l2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{l1: []int{}, l2: []int{}, expected: []int{}},
		{l1: []int{}, l2: []int{0}, expected: []int{0}},
	}

	for _, tt := range test6 {
		fmt.Printf("Test %v and %v is %v \n", tt.l1, tt.l2, tt.expected)
		list1 := createList(tt.l1)
		list2 := createList(tt.l2)
		expectedList := createList(tt.expected)

		result := mergeTwoLists(list1, list2)
		assert.True(t, compareLists(result, expectedList), "Merged list did not match the expected output")
	}
}
