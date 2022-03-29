package main

import "fmt"

type ListNode struct {
   Val int
   Next *ListNode
}

func main()  {
	//list1 := [3]new(ListNode){1,2,3}
	list1 := ListNode{Val: 1,Next:&ListNode{Val: 2,Next: &ListNode{Val: 4,Next: nil}}}
	list2 := ListNode{Val: 1,Next:&ListNode{Val: 3,Next: &ListNode{Val: 4,Next: nil}}}
	mergeTwoLists(&list1,&list2);

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) {
	newList := ListNode{Val: -1}
	for {
		newList.Next = list1
		if list1.Next == nil {
			break
		}
		list1 = list1.Next
	}


	for {
		fmt.Println(newList)
		if newList.Next == nil {
			break
		}
		newList = *newList.Next
	}

}
