package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	tree1 := BuildTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(maxDepth(tree1))
	
	tree2 := BuildTree([]int{1,-1,2})
	fmt.Println(maxDepth(tree2))
}

func maxDepth(root *TreeNode) int {
    for root == nil {
		return 0
	}
    left := maxDepth(root.Left)
	right := maxDepth(root.Right)

    return 1 + max(left, right)
}