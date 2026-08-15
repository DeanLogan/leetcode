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
	list1 := buildList([]int{1,2,3,4,5})
	ans1 := removeNthFromEnd(list1, 4)
	printList(ans1)
    
	list2 := buildList([]int{1})
	ans2 := removeNthFromEnd(list2, 1)
	printList(ans2)
    
	list3 := buildList([]int{1,2})
	ans3 := removeNthFromEnd(list3, 1)
	printList(ans3)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    currentHead := head
	nthHead := head

	for i:=0; i<n; i++ {
		if nthHead != nil {
			nthHead = nthHead.Next
		}
	}

	if nthHead == nil {
		return currentHead.Next
	}

	for currentHead.Next != nil {
		if nthHead.Next == nil {
			currentHead.Next = currentHead.Next.Next
			break
		}
		currentHead = currentHead.Next
		nthHead = nthHead.Next
	}
	return head
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
