package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{3,9,20,-1,-1,15,7})
	PrintTreeFancy(root1)
	fmt.Println(isBalanced(root1))
	
	root2 := BuildTree([]int{1,2,2,3,3,-1,-1,4,4})
	PrintTreeFancy(root2)
	fmt.Println(isBalanced(root2))

	root3 := BuildTree([]int{})
	PrintTreeFancy(root3)
	fmt.Println(isBalanced(root3))

	root4 := BuildTree([]int{1,-1,2,-1,3})
	PrintTreeFancy(root4)
	fmt.Println(isBalanced(root4))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	
	if (leftHeight-rightHeight) <= 1 && (leftHeight-rightHeight) >= -1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

    return false
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}