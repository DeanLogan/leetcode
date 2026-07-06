package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(printList(reverseList(buildList([]int{0, 1, 2, 3}))))
	fmt.Println(printList(reverseList(buildList([]int{2, 1}))))
	fmt.Println(printList(reverseList(buildList(nil))))
}

func reverseList(head *ListNode) *ListNode {
    prev := (*ListNode)(nil)
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return head
}

func printList(head *ListNode) string {
	if head == nil {
		return "[]"
	}

	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, fmt.Sprint(node.Val))
	}

	return "[" + strings.Join(values, ",") + "]"
}
