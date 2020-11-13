package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Add(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var ret *ListNode
	ret = new(ListNode)
	now := ret
	var i, s int

	for {

		if l1 == nil && l2 == nil {
			if i > 0 {
				now.Val = i
			}
			break
		}

		if l1 == nil {
			s = l2.Val + i
			if s >= 10 {
				now.Val = s % 10
				i = s / 10
			} else {
				now.Val = s
				i = 0
			}
			l2 = l2.Next
		} else if l2 == nil {
			s = l1.Val + i
			if s >= 10 {
				now.Val = s % 10
				i = s / 10
			} else {
				now.Val = s
				i = 0
			}
			l1 = l1.Next
		} else {

			if i > 0 {
				s = l1.Val + l2.Val + 1
			} else {
				s = l1.Val + l2.Val
			}
			if s >= 10 {
				now.Val = s % 10
				i = s / 10
			} else {
				now.Val = s
				i = 0
			}
			l1 = l1.Next
			l2 = l2.Next
		}

		if l1 != nil || l2 != nil || i > 0 {
			now.Next = &ListNode{}
			now = now.Next
		}
	}

	return ret
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	ret := Add(l1, l2)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
