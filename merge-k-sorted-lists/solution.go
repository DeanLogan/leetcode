package main

import (
	"fmt"
	"strings"
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	list1 := buildList([]int{1,4,5})
	list2 := buildList([]int{1,3,4})
	list3 := buildList([]int{2,6})
	ans1 := mergeKLists([]*ListNode{list1, list2, list3})
	printList(ans1)
	
	ans2 := mergeKLists([]*ListNode{buildList([]int{})})
	printList(ans2)
	
	ans3 := mergeKLists([]*ListNode{})
	printList(ans3)
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

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return head
}

func printList(head *ListNode) {
	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, fmt.Sprint(node.Val))
	}

	fmt.Println("[" + strings.Join(values, ",") + "]")
}
