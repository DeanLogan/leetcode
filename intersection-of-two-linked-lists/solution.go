package main

import (
	"fmt"

	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	// Example 1: intersection at node with value 8
	shared1 := BuildList([]int{8, 4, 5})
	headA1 := BuildList([]int{4, 1})
	headB1 := BuildList([]int{5, 6, 1})
	getTail(headA1).Next = shared1
	getTail(headB1).Next = shared1
	fmt.Println(getIntersectionNode(headA1, headB1))

	// Example 2: intersection at node with value 2
	shared2 := BuildList([]int{2, 4})
	headA2 := BuildList([]int{1, 9, 1})
	headB2 := BuildList([]int{3})
	getTail(headA2).Next = shared2
	getTail(headB2).Next = shared2
	fmt.Println(getIntersectionNode(headA2, headB2))

	// Example 3: no intersection
	headA3 := BuildList([]int{2, 6, 4})
	headB3 := BuildList([]int{1, 5})
	fmt.Println(getIntersectionNode(headA3, headB3))
}

func getTail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
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
