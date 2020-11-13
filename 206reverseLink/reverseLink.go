package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	cur = head

	for cur != nil {
		tmpNode := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNode
	}
	return pre
}

func main() {
	var head *ListNode

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ret := ReverseList(head)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
