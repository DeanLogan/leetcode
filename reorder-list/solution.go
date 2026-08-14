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
	list1 := buildList([]int{1,2,3,4})
    list2 := buildList([]int{1,2,3,4,5})
	reorderList(list1)
	reorderList(list2)
	printList(list1)
	printList(list2)
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
	if head == nil {
		fmt.Println("[]")
	}

	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, fmt.Sprint(node.Val))
	}

	fmt.Println("[" + strings.Join(values, ",") + "]")
}
