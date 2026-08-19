package main

import (
	"fmt"
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	p1 := BuildTree([]int{1,2,3})
	q1 := BuildTree([]int{1,2,3})
	fmt.Println(isSameTree(p1, q1))

	p2 := BuildTree([]int{1,2})
	q2 := BuildTree([]int{1,-1, 2})
	fmt.Println(isSameTree(p2, q2))

	p3 := BuildTree([]int{1,2,1})
	q3 := BuildTree([]int{1,1,2})
	fmt.Println(isSameTree(p3, q3))
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