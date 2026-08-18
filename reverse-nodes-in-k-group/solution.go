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
	printList(reverseKGroup(buildList([]int{1,2,3,4,5}), 2))
	// printList(reverseKGroup(buildList([]int{1,2,3,4,5}), 3))
	// printList(reverseKGroup(buildList([]int{1,2}), 2))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
    prevGroupEnd := dummy

	for {
		kthNode := getKthNode(prevGroupEnd, k)
		if kthNode == nil {
			break
		}
		groupStart := prevGroupEnd.Next
        nextGroupStart := kthNode.Next

		reverseList(groupStart, nextGroupStart)

		prevGroupEnd.Next = kthNode
        prevGroupEnd = groupStart
	}	


	return dummy.Next
}

func getKthNode(head *ListNode, k int) *ListNode {
	for head != nil && k > 0 {
		head = head.Next
		k--
	}
	return head
}

func reverseList(head *ListNode, endOfList *ListNode) {
    prev := endOfList
    for head != endOfList {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
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

func printList(head *ListNode) {
	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, fmt.Sprint(node.Val))
	}

	fmt.Println("[" + strings.Join(values, ",") + "]")
}
