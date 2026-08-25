package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{1,2,3,-1,5,-1,4})
	fmt.Println(rightSideView(root1))

	root2 := BuildTree([]int{1,2,3,4,-1,-1,-1,5})
	fmt.Println(rightSideView(root2))

	root3 := BuildTree([]int{1,-1,3})
	fmt.Println(rightSideView(root3))

	root4 := BuildTree([]int{})
	fmt.Println(rightSideView(root4))
}

func rightSideView(root *TreeNode) []int {
	rightSideVals := []int{}
	if root == nil {
		return rightSideVals
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		rightSideVals = append(rightSideVals, queue[levelSize-1].Val)
		for i:=0; i<levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftNode := node.Left; leftNode != nil {
				queue = append(queue, leftNode)
			}
			if rightNode := node.Right; rightNode != nil {
				queue = append(queue, rightNode)
			}
		}
	}
    return rightSideVals
}