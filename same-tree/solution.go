package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	p1 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q1 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	p2 := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	q2 := TreeNode{1, nil, &TreeNode{2, nil, nil}}
	p3 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}
	q3 := TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}

	fmt.Println(isSameTree(&p1, &q1))
	fmt.Println(isSameTree(&p2, &q2))
	fmt.Println(isSameTree(&p3, &q3))
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