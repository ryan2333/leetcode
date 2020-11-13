package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode, val int) {
	if node.Val == val {
		node = node.Next
		return
	}
	pre, cur := node, node.Next

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	cur = node

	for cur != nil {
		fmt.Printf("val: %d ", node.Val)
		node = node.Next
	}
	fmt.Println("")
}

func main() {
	l := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(l, 4)
	l = &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(l, 5)
	l = &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(l, 1)
	l = &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(l, 9)
}
