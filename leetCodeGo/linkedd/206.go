package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	end := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return end
}

func reverseList2(head *ListNode) *ListNode {
	cur := head
	for {
		if head == nil {
			break
		}
		pre := cur.Next
		cur.Next = nil
		pre.Next = cur
		head = head.Next
	}
	return cur
}
