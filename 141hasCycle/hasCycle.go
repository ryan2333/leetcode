package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	resMap := make(map[*ListNode]int)

	for head != nil {
		if _, ok := resMap[head]; ok {
			return true
		}
		resMap[head] = head.Val
		head = head.Next
	}

	return false
}

func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var fast = head
	var slow = head

	for slow != nil && fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	n1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	//n3 := &ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//n4 := &ListNode{
	//	Val:  -4,
	//	Next: nil,
	//}

	n1.Next = n2
	//n2.Next = n1
	//n3.Next = n4
	//n4.Next = n2

	fmt.Println(hasCycle1(n1))

}
