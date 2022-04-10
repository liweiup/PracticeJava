package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	nl := ListNode{Val: -1}
	p := &nl
	for {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
		if list1 == nil || list2 == nil {
			break
		}
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}

	nx := nl.Next
	return nx
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next,list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1,list2.Next)
		return list2
	}
}
