package main

import (
	"fmt"

	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	fmt.Println(hasCycle(buildCycleList([]int{3,2,0,-4}, 1)))
	fmt.Println(hasCycle(buildCycleList([]int{1,2}, 0)))
	fmt.Println(hasCycle(buildCycleList([]int{1}, -1)))
}

func hasCycle(head *ListNode) bool {
	slowPointer := head
    fastPointer := head

	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

func hasCycleHashMapApproach(head *ListNode) bool {
    vistedNodes := make(map[*ListNode]bool)

	for head != nil {
		vistedNodes[head] = true
		if head.Next != nil && vistedNodes[head.Next] {
			return true
		}
		head = head.Next
	}
	return false
}

func buildCycleList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}
