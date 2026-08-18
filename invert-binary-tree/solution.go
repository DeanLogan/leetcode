package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	tree1 := BuildTree([]int{4,2,7,1,3,6,9})
	ans1 := invertTree(tree1)
	PrintTree(ans1)
	PrintTreeFancy(ans1)

	tree2 := BuildTree([]int{2,1,3})
	ans2 := invertTree(tree2)
	PrintTree(ans2)
	PrintTreeFancy(ans2)
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
        root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    }
    return root
}