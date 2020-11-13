package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var ret = new(ListNode)
	now := ret

	for {
		if l1.Val > l2.Val {
			now.Val = l2.Val
			l2 = l2.Next
		} else {
			now.Val = l1.Val
			l1 = l1.Next
		}

		newNode := &ListNode{}
		now.Next = newNode
		now = now.Next

		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			now.Val = l2.Val
			now.Next = l2.Next
			break
		}

		if l2 == nil {
			now.Val = l1.Val
			now.Next = l1.Next
			break
		}
	}
	return ret
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	mergeTwoLists(l1, l2)
}
