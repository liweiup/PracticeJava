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
	for {
		nlr := nl.Next
		for {
			if nl.Val > nlr.Val {
				pp.Next = nlr
			} else {
				pp.Next = &nl
			}
			if nlr == nil || nlr.Next == nil {
				break
			}
			nlr = nlr.Next
		}
		if nl.Next == nil {
			break
		}
		nl = *nl.Next
	}
	fmt.Println(pp)
	return nil
}
