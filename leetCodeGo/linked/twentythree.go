package linked

import "fmt"

func MergeKLists(lists []*ListNode) *ListNode {
	nl := ListNode{Val: -1}
	p := &nl
	pp := &nl
	for _, list := range lists {
		for {
			p.Next = list
			p = p.Next
			if list.Next == nil {
				break;
			}
			list = list.Next
		}
	}
	fmt.Println(pp)
	return nil
}
