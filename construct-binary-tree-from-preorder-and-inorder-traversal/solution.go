package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	PrintTree(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
	PrintTree(buildTree([]int{-1}, []int{-1}))
	PrintTree(buildTree([]int{1,2}, []int{2,1}))
	PrintTree(buildTree([]int{1,2}, []int{1,2}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	index := make(map[int]int)

	for i, v := range inorder {
		index[v] = i
	}

	preIdx := 0

	return build(preorder, index, &preIdx, 0, len(inorder)-1)
}

func build(preorder []int, index map[int]int, preIdx *int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := preorder[*preIdx]
	*preIdx++
	root := &TreeNode{Val: rootVal}
	mid := index[rootVal]

	root.Left = build(preorder, index, preIdx, left, mid-1)
	root.Right = build(preorder, index, preIdx, mid+1, right)

	return root
}