package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{1,2,3,4,5,-1,-1})
	fmt.Println(diameterOfBinaryTree(root1))
	root2 := BuildTree([]int{1,2})
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