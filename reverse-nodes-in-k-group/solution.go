package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	PrintList(reverseKGroup(BuildList([]int{1,2,3,4,5}), 2))
	// PrintList(reverseKGroup(BuildList([]int{1,2,3,4,5}), 3))
	// PrintList(reverseKGroup(BuildList([]int{1,2}), 2))
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
