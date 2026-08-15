package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
	Random *Node
}

func main() {
	list1 := buildRandomList([]int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0})
	ans1 := copyRandomList(list1)
	printRandomList(list1)
	printRandomList(ans1)

	list2 := buildRandomList([]int{1, 2}, []int{1, 1})
	ans2 := copyRandomList(list2)
	printRandomList(list2)
	printRandomList(ans2)

	list3 := buildRandomList([]int{3, 3, 3}, []int{-1, 0, -1})
	ans3 := copyRandomList(list3)
	printRandomList(list3)
	printRandomList(ans3)

	list4 := buildRandomList([]int{}, []int{})
	ans4 := copyRandomList(list4)
	printRandomList(list4)
	printRandomList(ans4)
}

func copyRandomList(head *Node) *Node {
	oldToNewMap := make(map[*Node]*Node)

	prevNewNode := &Node{Next: nil}
	copyOfDummy := prevNewNode

	for head != nil {
		newNode := &Node{Val: head.Val}
		prevNewNode.Next = newNode
		oldToNewMap[head] = newNode
		prevNewNode = newNode
		head = head.Next
	}

	for oldNode, newNode := range oldToNewMap {
		newNode.Random = oldToNewMap[oldNode.Random]
	}

    return copyOfDummy.Next
}

func buildRandomList(values []int, randomIdx []int) *Node {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*Node, len(values))
	for i, v := range values {
		nodes[i] = &Node{Val: v}
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	for i, ri := range randomIdx {
		if ri != -1 {
			nodes[i].Random = nodes[ri]
		}
	}

	return nodes[0]
}

func printRandomList(head *Node) {
	if head == nil {
		fmt.Println("[]")
		return
	}

	index := make(map[*Node]int)
	i := 0
	for node := head; node != nil; node = node.Next {
		index[node] = i
		i++
	}

	var pairs []string
	for node := head; node != nil; node = node.Next {
		randomStr := "null"
		if node.Random != nil {
			randomStr = fmt.Sprint(index[node.Random])
		}
		pairs = append(pairs, fmt.Sprintf("[%d,%s]", node.Val, randomStr))
	}

	fmt.Println("[" + strings.Join(pairs, ",") + "]")
}
