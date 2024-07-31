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
	root1.Left = &TreeNode{9, nil, nil}
	root1.Right = &TreeNode{20, nil, nil}
	root1.Right.Left = &TreeNode{15, nil, nil}
	root1.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(sumOfLeftLeaves(root1))
	root2 := &TreeNode{1, nil, nil}
	fmt.Println(sumOfLeftLeaves(root2))
}

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    sum := 0
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        if node.Left != nil {
            if node.Left.Left == nil && node.Left.Right == nil {
                sum += node.Left.Val
            } else {
                queue = append(queue, node.Left)
            }
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }

    return sum
}