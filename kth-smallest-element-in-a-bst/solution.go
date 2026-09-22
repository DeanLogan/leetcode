package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	fmt.Println(kthSmallest(BuildTree([]int{3,1,4,-1,2}), 1))
	fmt.Println(kthSmallest(BuildTree([]int{5,3,6,2,4,-1,-1,1}), 3))
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	currentNode := root

	for currentNode != nil || len(stack) > 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return currentNode.Val
		}

		currentNode = currentNode.Right
	}

	return -1
}