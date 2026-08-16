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
	ans1 := addTwoNumbers(buildList([]int{2,4,3}), buildList([]int{5,6,4}))
	printList(ans1)
	ans2 := addTwoNumbers(buildList([]int{0}), buildList([]int{0}))
	printList(ans2)
	ans3 := addTwoNumbers(buildList([]int{9,9,9,9,9,9,9}), buildList([]int{9,9,9,9}))
	printList(ans3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	dummyNode := &ListNode{}
	prevNode := dummyNode
    for l1 != nil || l2 != nil || carry {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		computedVal := val1 + val2
		if carry {
			computedVal = computedVal + 1
			carry = false
		}
		if computedVal >= 10 {
			computedVal = computedVal - 10
			carry = true
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		prevNode.Next = &ListNode{Val: computedVal}
		prevNode = prevNode.Next
	}
	return dummyNode.Next
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
