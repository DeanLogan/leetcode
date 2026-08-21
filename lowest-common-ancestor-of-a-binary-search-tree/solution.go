package main

import (
	. "github.com/DeanLogan/leetcode/libs"
)

func main() {
	root1 := BuildTree([]int{6,2,8,0,4,7,9,-1,-1,3,5})
    p1 := root1.Left
    q1 := root1.Right
	ans1 := lowestCommonAncestor(root1, p1, q1)
	PrintTree(ans1)

	root2 := BuildTree([]int{6,2,8,0,4,7,9,-1,-1,3,5})
    p2 := root2.Left
    q2 := root2.Left.Right
	ans2 := lowestCommonAncestor(root2, p2, q2)
	PrintTree(ans2)

	root3 := BuildTree([]int{2,1})
    p3 := root3
    q3 := root3.Right
	ans3 := lowestCommonAncestor(root3, p3, q3)
	PrintTree(ans3)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for (root != p && root != q) {
        if (root.Val > p.Val && root.Val > q.Val) {
            root = root.Left
        } else if (root.Val < p.Val && root.Val < q.Val) {
            root = root.Right
        } else{
            return root
        }
    }
    return root
}