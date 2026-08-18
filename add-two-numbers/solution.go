package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	ans1 := addTwoNumbers(BuildList([]int{2,4,3}), BuildList([]int{5,6,4}))
	PrintList(ans1)
	ans2 := addTwoNumbers(BuildList([]int{0}), BuildList([]int{0}))
	PrintList(ans2)
	ans3 := addTwoNumbers(BuildList([]int{9,9,9,9,9,9,9}), BuildList([]int{9,9,9,9}))
	PrintList(ans3)
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
