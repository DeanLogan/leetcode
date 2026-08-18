package main

import (
	"container/heap"

	. "github.com/DeanLogan/leetcode/libs"
)

type PriorityQueue []*ListNode

func (h PriorityQueue) Len() int { 
	return len(h) 
}

func (h PriorityQueue) Less(i, j int) bool { 
	return h[i].Val < h[j].Val 
}

func (h PriorityQueue) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i] 
}

func (h *PriorityQueue) Push(x any) {
    node := x.(*ListNode)
    if node != nil {
        *h = append(*h, node)
    }
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	list1 := BuildList([]int{1,4,5})
	list2 := BuildList([]int{1,3,4})
	list3 := BuildList([]int{2,6})
	ans1 := mergeKLists([]*ListNode{list1, list2, list3})
	PrintList(ans1)
	
	ans2 := mergeKLists([]*ListNode{BuildList([]int{})})
	PrintList(ans2)
	
	ans3 := mergeKLists([]*ListNode{})
	PrintList(ans3)
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, head := range lists  {
		heap.Push(pq, head)
	}

	dummy := &ListNode{}
	prevNode := dummy
	for pq.Len() > 0 {
		minNode := heap.Pop(pq).(*ListNode)
		heap.Push(pq, minNode.Next)
		prevNode.Next = minNode
		prevNode = minNode
	}

	return dummy.Next
}
