package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	PrintList(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}))
	PrintList(swapPairs(nil))
	PrintList(swapPairs(&ListNode{Val: 1}))
	PrintList(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}))
}

func swapPairs(head *ListNode) *ListNode {
    tempHead := &ListNode{Next: head}
    prev := tempHead

    for head != nil && head.Next != nil {
        first := head
        second := head.Next

        first.Next = second.Next
        second.Next = first
        prev.Next = second

        prev = first
        head = first.Next
    }

    return tempHead.Next
}
