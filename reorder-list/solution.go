package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	list1 := BuildList([]int{1,2,3,4})
    list2 := BuildList([]int{1,2,3,4,5})
	reorderList(list1)
	reorderList(list2)
	PrintList(list1)
	PrintList(list2)
}

func reorderList(head *ListNode)  {
	listLen := 0
	originalHead := head
	for head != nil {
		head = head.Next
		listLen++
	}
	if listLen <= 1 {
		return
	}
	head = originalHead
	secondHalf := getMidpointOfLinkedList(head, listLen)
	secondHalf = reverseLinkedList(secondHalf)
	mergeLinkedLists(head, secondHalf)
}

func getMidpointOfLinkedList(head *ListNode, listLen int) *ListNode {
	halfwayPoint := listLen / 2
	prevNode := head
	for i := 0; i < halfwayPoint; i++ {
		prevNode = head
		head = head.Next
	}
	prevNode.Next = nil
	return head
}

func reverseLinkedList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func mergeLinkedLists(head1 *ListNode, head2 *ListNode) {
	for head1.Next != nil {
		current1 := head1.Next
		current2 := head2.Next
		head1.Next = head2
		head2.Next = current1
		head1 = current1
		head2 = current2
	}
	head1.Next = head2
}
