package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{}
	pre := res
	for head != nil && head.Next != nil {
		nx := head.Next.Next
		pre.Next = head.Next
		head.Next = nil
		pre.Next.Next = head
		pre = pre.Next.Next
		head = nx

	}
	if head != nil {
		pre.Next = head
	}
	//for q.Next != nil {
	//	q = q.Next
	//	fmt.Println(q.Val)
	//}

	return res.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{}
	pre := res

	for head != nil && head.Next != nil {
		nextHead := head.Next.Next
		pre.Next = head.Next
		head.Next = nil
		pre.Next.Next = head
		pre = pre.Next.Next
		head = nextHead
	}
	if head != nil {
		pre.Next = head
	}
	return res.Next
}

func main() {
	head := &ListNode{
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
	swapPairs2(head)
}
