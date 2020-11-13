package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tmpNode := head
	oldNode := head

	for tmpNode.Next != nil {
		next := tmpNode.Next
		if oldNode.Val == next.Val {
			oldNode.Next = nil
			continue
		}
		oldNode.Next = next
		oldNode = oldNode.Next
		tmpNode = tmpNode.Next
	}
	for head != nil {
		fmt.Printf("val: %d\n", head.Val)
		head = head.Next
	}
	return head
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	deleteDuplicates(l1)
}
