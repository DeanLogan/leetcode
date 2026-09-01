package main

import (
	"fmt"
	"math"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{2,1,3})
	fmt.Println(isValidBST(root1))

	root2 := BuildTree([]int{5,1,4,-1,-1,3,6})
	fmt.Println(isValidBST(root2))
}

func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}

	val := root.Val

	if val >= maxVal || val <= minVal {
		return false
	}

	return dfs(root.Left, minVal, val) && dfs(root.Right, val, maxVal)
}