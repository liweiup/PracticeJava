package main

import "fmt"

func reverseList3(nodeList *ListNode) *ListNode {
	if nodeList.Next == nil {
		return nodeList
	}
	endNode := reverseList3(nodeList.Next)
	nodeList.Next.Next = nodeList
	nodeList.Next = nil
	return endNode
}

func reverseList4(head *ListNode) *ListNode {
	temp := ListNode{}
	for {
		if head == nil {
			break
		}
		next := head.Next
		head.Next = temp.Next
		temp.Next = head
		head = next
	}
	return temp.Next
}

func main() {
	// list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// list2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	// nx := mergeTwoLists2(&list1, &list2)
	list3 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// nx := reverseList(&list3)
	nx := reverseList4(&list3)
	// fmt.Println(nx)

	for {
		fmt.Println(nx)
		if nx == nil || nx.Next == nil {
			break
		}
		nx = nx.Next
	}

}
