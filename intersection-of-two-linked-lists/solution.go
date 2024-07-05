package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headA1 := &ListNode{Val: 4}
	headA1.Next = &ListNode{Val: 1}
	headA1.Next.Next = &ListNode{Val: 8}
	headA1.Next.Next.Next = &ListNode{Val: 4}
	headA1.Next.Next.Next.Next = &ListNode{Val: 5}
	headB1 := &ListNode{Val: 5}
	headB1.Next = &ListNode{Val: 6}
	headB1.Next.Next = &ListNode{Val: 1}
	headB1.Next.Next.Next = &ListNode{Val: 8}
	headB1.Next.Next.Next.Next = &ListNode{Val: 4}
	headB1.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(getIntersectionNode(headA1, headB1))

	headA2 := &ListNode{Val: 1}
	headA2.Next = &ListNode{Val: 9}
	headA2.Next.Next = &ListNode{Val: 1}
	headA2.Next.Next.Next = &ListNode{Val: 2}
	headA2.Next.Next.Next.Next = &ListNode{Val: 4}
	headB2 := &ListNode{Val: 3}
	headB2.Next = &ListNode{Val: 2}
	headB2.Next.Next = &ListNode{Val: 4}
	fmt.Println(getIntersectionNode(headA2, headB2))

	headA3 := &ListNode{Val: 2}
	headA3.Next = &ListNode{Val: 6}
	headA3.Next.Next = &ListNode{Val: 4}
	headB3 := &ListNode{Val: 1}
	headB3.Next = &ListNode{Val: 4}
	fmt.Println(getIntersectionNode(headA3, headB3))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := getLength(headA), getLength(headB)

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}

	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}