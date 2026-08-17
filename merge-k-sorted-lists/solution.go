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
	list1 := buildList([]int{1,4,5})
	list2 := buildList([]int{1,3,4})
	list3 := buildList([]int{2,6})
	ans1 := mergeKLists([]*ListNode{list1, list2, list3})
	printList(ans1)
	
	ans2 := mergeKLists([]*ListNode{buildList([]int{})})
	printList(ans2)
	
	ans3 := mergeKLists([]*ListNode{})
	printList(ans3)
}

func mergeKLists(lists []*ListNode) *ListNode {
    dummyHead := &ListNode{Val: -2147483648}

	for _, head := range lists {
		for head != nil {
			addValueToList(head.Val, dummyHead)
			head = head.Next
		}
	}

	return dummyHead.Next
}

func addValueToList(value int, head *ListNode) {
	prevNode := (*ListNode)(nil)
	for head != nil {
		if head.Val >= value {
			node := &ListNode{Val: value, Next: head}
			prevNode.Next = node
			break
		}
		prevNode = head
		head = head.Next
	}

	if head == nil {
		prevNode.Next = &ListNode{Val: value}
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
	if head == nil {
		fmt.Println("[]")
	}

	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, fmt.Sprint(node.Val))
	}

	fmt.Println("[" + strings.Join(values, ",") + "]")
}
