package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{Val: 10}
	root1.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 6}
	fmt.Println(checkTree(root1))
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 1}
	fmt.Println(checkTree(root1))
	fmt.Println(checkTree(root2))
}

func checkTree(root *TreeNode) bool {
    return root.Val == root.Left.Val + root.Right.Val
}