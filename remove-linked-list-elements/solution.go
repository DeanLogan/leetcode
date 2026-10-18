package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	PrintList(removeElements(BuildList([]int{1,2,6,3,4,5,6}), 6))
	PrintList(removeElements(BuildList([]int{}), 1))
	PrintList(removeElements(BuildList([]int{7,7,7,7}), 7))
}

func removeElements(head *ListNode, val int) *ListNode {
    temp := &ListNode{Next: head}
    prev := temp

    for curr := head; curr != nil; curr = curr.Next {
        if curr.Val == val {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
    }
    return temp.Next
}