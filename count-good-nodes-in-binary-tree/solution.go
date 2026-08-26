package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{3,1,4,3,-1,1,5})
	fmt.Println(goodNodes(root1))

	root2 := BuildTree([]int{3,3,-1,4,2})
	fmt.Println(goodNodes(root2))

	root3 := BuildTree([]int{1})
	fmt.Println(goodNodes(root3))
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, maxVal int) int {
	goodNodeCount := 0
	if root == nil {
		return goodNodeCount
	}

	if root.Val >= maxVal {
		goodNodeCount++
		maxVal = root.Val
	}

	goodNodeCount += dfs(root.Left, maxVal)
	goodNodeCount += dfs(root.Right, maxVal)
	return  goodNodeCount
}