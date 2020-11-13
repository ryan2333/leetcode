package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func loopHead(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
func removeelements(head *ListNode, val int) *ListNode {
	var pre, cur, temp *ListNode
	temp = &ListNode{0, head}
	pre, cur = temp, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return temp.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 6,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 5, Next: &ListNode{6, nil}}}}}},
	}
	loopHead(removeelements(head, 6))
}
