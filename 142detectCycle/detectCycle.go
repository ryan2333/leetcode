package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	//var index = 0
	resMap := make(map[*ListNode]*ListNode)

	for head != nil {
		if i, ok := resMap[head]; ok {
			resMap = nil
			return i
		}
		resMap[head] = head
		head = head.Next
	}

	return nil
}

func main() {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	fmt.Println(detectCycle(n1).Val)
}
