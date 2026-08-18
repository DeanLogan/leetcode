package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	PrintList(reverseList(BuildList([]int{0, 1, 2, 3})))
	PrintList(reverseList(BuildList([]int{2, 1})))
	PrintList(reverseList(BuildList(nil)))
}

func reverseList(head *ListNode) *ListNode {
    prev := (*ListNode)(nil)
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}
