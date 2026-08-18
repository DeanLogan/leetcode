package main

import (
	"fmt"

	. "github.com/DeanLogan/leetcode/libs"
)

func main(){
	head1 := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	head2 := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	fmt.Println(deleteDuplicates(&head1))
	fmt.Println(deleteDuplicates(&head2))
	PrintList(deleteDuplicates(&head1))
	PrintList(deleteDuplicates(&head2))
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
