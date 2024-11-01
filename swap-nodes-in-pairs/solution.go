package main

import (
	"fmt"
)

func main() {
	printList(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
	printList(swapPairs(nil))
	printList(swapPairs(&ListNode{1, nil}))
	printList(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    tempHead := &ListNode{Next: head}
    prev := tempHead

    for head != nil && head.Next != nil {
        first := head
        second := head.Next

        first.Next = second.Next
        second.Next = first
        prev.Next = second

        prev = first
        head = first.Next
    }

    return tempHead.Next
}

// a function to print out a linked list in a nice way to the console
func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}