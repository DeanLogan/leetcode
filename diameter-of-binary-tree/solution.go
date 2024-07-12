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
	root1 := &TreeNode{1, nil, nil}
	root1.Left = &TreeNode{2, nil, nil}
	root1.Right = &TreeNode{3, nil, nil}
	root1.Left.Left = &TreeNode{4, nil, nil}
	root1.Left.Right = &TreeNode{5, nil, nil}
	fmt.Println(diameterOfBinaryTree(root1))
	root2 := &TreeNode{1, nil, nil}
	root2.Left = &TreeNode{2, nil, nil}
	fmt.Println(diameterOfBinaryTree(root2))
}

func diameterOfBinaryTree(root *TreeNode) int {
    maxPath := 0
    var diameter func(*TreeNode) int
    diameter = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        leftPath := diameter(root.Left)
        rightPath := diameter(root.Right)
        maxPath = max(maxPath, leftPath+rightPath )
        return max(leftPath,rightPath) + 1
    }
    diameter(root)
    return maxPath  
}