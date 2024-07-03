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
    var emptyRoot *TreeNode = nil

    singleNodeRoot := &TreeNode{Val: 1}

    multipleNodesRoot := &TreeNode{Val: 1}
    multipleNodesRoot.Left = &TreeNode{Val: 2}
    multipleNodesRoot.Right = &TreeNode{Val: 3}
    multipleNodesRoot.Left.Left = &TreeNode{Val: 4}
    multipleNodesRoot.Left.Right = &TreeNode{Val: 5}
    multipleNodesRoot.Right.Left = &TreeNode{Val: 6}

    fmt.Println(countNodes(emptyRoot))
    fmt.Println(countNodes(singleNodeRoot))
    fmt.Println(countNodes(multipleNodesRoot))
}

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := getDepth(root.Left, true)
    rightDepth := getDepth(root.Right, false)
    if leftDepth == rightDepth {
        return (1 << (leftDepth + 1)) - 1 // 2^n - 1
    } 
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func getDepth(node *TreeNode, isLeft bool) int {
    depth := 0
    for node != nil {
        depth++
        if isLeft {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return depth
}