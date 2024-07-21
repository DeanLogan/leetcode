package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	head1 := ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head2 := ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	fmt.Println(deleteDuplicates(&head1))
	fmt.Println(deleteDuplicates(&head2))
	printList(deleteDuplicates(&head1))
	printList(deleteDuplicates(&head2))
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

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d -> ", head.Val)
        head = head.Next
    }
    fmt.Println("nil")
}