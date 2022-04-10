package main

import "fmt"

func main() {
	// list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// list2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	// nx := mergeTwoLists2(&list1, &list2)
	list3 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// nx := reverseList(&list3)
	nx := reverseList2(&list3)
	// fmt.Println(nx)

	for {
		fmt.Println(nx)
		if nx == nil || nx.Next == nil {
			break
		}
		nx = nx.Next
	}

}
