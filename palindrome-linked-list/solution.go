package main

import (
	"fmt"

	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	fmt.Println(isPalindrome(BuildList([]int{1, 2, 2, 1})))
	fmt.Println(isPalindrome(BuildList([]int{1, 2})))
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
