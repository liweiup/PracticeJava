package linked

import "fmt"

type ListNode struct {
   Val int
   Next *ListNode
}

func main()  {
	list1 := ListNode{Val: 1,Next:&ListNode{Val: 2,Next: &ListNode{Val: 4,Next: nil}}}
	list2 := ListNode{Val: 1,Next:&ListNode{Val: 3,Next: &ListNode{Val: 4,Next: nil}}}
	mergeTwoLists(&list1,&list2);
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) {
	nl := ListNode{Val: -1}
	p := &nl
	for {
		fmt.Println(list1)
		fmt.Println(list2)
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
	fmt.Println("===========")
	//for {
	//	fmt.Println(nl)
	//	if nl.Next == nil {
	//		break
	//	}
	//	nl = *nl.Next
	//}
	nx := nl.Next
	fmt.Println(nx)
}
