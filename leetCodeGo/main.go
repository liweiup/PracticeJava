package main

import "leetcode/linked"

func main()  {
	list1 := linked.ListNode{Val: 1,Next:&linked.ListNode{Val: 4,Next: &linked.ListNode{Val: 5,Next: nil}}}
	list2 := linked.ListNode{Val: 1,Next:&linked.ListNode{Val: 3,Next: &linked.ListNode{Val: 4,Next: nil}}}
	list3 := linked.ListNode{Val: 2,Next:&linked.ListNode{Val: 6}}
	listNode := []*linked.ListNode{
		&list1,&list2,&list3,
	}
	linked.MergeKLists(listNode);
}

