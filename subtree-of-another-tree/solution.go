package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{3,4,5,1,2})
	subRoot1 := BuildTree([]int{4,1,2})
	fmt.Println(isSubtree(root1, subRoot1))

	root2 := BuildTree([]int{3,4,5,1,2,-1,-1,-1,-1,0})
	subRoot2 := BuildTree([]int{4,1,2})
	fmt.Println(isSubtree(root2, subRoot2))
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    
    if p == nil || q == nil {
        return false
    }
    
    if p.Val != q.Val {
        return false
    }
    
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}