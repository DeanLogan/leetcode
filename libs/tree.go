package libs

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// example: BuildTree([]int{1, 2, 3, -1, -1, 4, 5}):
//	    1
//	   / \
//	  2   3
//	     / \
//	    4   5
func BuildTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != -1 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != -1 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	var values []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			values = append(values, "null")
		} else {
			values = append(values, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for len(values) > 0 && values[len(values)-1] == "null" {
		values = values[:len(values)-1]
	}

	fmt.Println("[" + strings.Join(values, ",") + "]")
}
