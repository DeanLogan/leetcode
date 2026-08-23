package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(levelOrder(root1))

	root2 := BuildTree([]int{1})
	fmt.Println(levelOrder(root2))
	
	root3 := BuildTree([]int{})
	fmt.Println(levelOrder(root3))
}

func levelOrder(root *TreeNode) [][]int {
	list := make([][]int,0)
	if root == nil {
		return nil
	}
	
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelList := make([]int, 0, levelSize)
		for i:=0; i<levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelList = append(levelList, node.Val)
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		list = append(list, levelList)
	}

	return list
}
