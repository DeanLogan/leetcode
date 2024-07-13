package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	ListNode1 := &ListNode{Val: 1}
	ListNode2 := &ListNode{Val: 2}
	ListNode3 := &ListNode{Val: 2}
	ListNode4 := &ListNode{Val: 1}
	ListNode1.Next = ListNode2
	ListNode2.Next = ListNode3
	ListNode3.Next = ListNode4
	fmt.Println(isPalindrome(ListNode1))
	ListNode5 := &ListNode{Val: 1}
	ListNode6 := &ListNode{Val: 1}
	ListNode5.Next = ListNode6
	fmt.Println(isPalindrome(ListNode5))
}

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = reverseList(slow.Next)
    slow = slow.Next

    p1, p2 := head, slow
    for p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    return true
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }
    return prev
}