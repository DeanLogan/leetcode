package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	list1 := BuildList([]int{1,2,3,4,5})
	ans1 := removeNthFromEnd(list1, 4)
	PrintList(ans1)
    
	list2 := BuildList([]int{1})
	ans2 := removeNthFromEnd(list2, 1)
	PrintList(ans2)
    
	list3 := BuildList([]int{1,2})
	ans3 := removeNthFromEnd(list3, 1)
	PrintList(ans3)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    currentHead := head
	nthHead := head

	for i:=0; i<n; i++ {
		if nthHead != nil {
			nthHead = nthHead.Next
		}
	}

	if nthHead == nil {
		return currentHead.Next
	}

	for currentHead.Next != nil {
		if nthHead.Next == nil {
			currentHead.Next = currentHead.Next.Next
			break
		}
		currentHead = currentHead.Next
		nthHead = nthHead.Next
	}
	return head
}
