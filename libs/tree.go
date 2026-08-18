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

// PrintTreeFancy prints trees like this:
//	        4
//	       / \
//	      2   7
//	     / \ / \
//	    1  3 6  9
func PrintTreeFancy(root *TreeNode) {
	if root == nil {
		fmt.Println("<empty>")
		return
	}

	lines := buildFancyLines(root)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func buildFancyLines(node *TreeNode) []string {
	if node == nil {
		return []string{}
	}

	valStr := strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		return []string{valStr}
	}

	if node.Left == nil {
		rightLines := buildFancyLines(node.Right)
		rightWidth := len(rightLines[0])
		firstLine := valStr + strings.Repeat(" ", rightWidth)
		secondLine := strings.Repeat(" ", len(valStr)) + "\\" + strings.Repeat(" ", rightWidth-1)
		result := []string{firstLine, secondLine}
		padding := strings.Repeat(" ", len(valStr))
		for _, line := range rightLines {
			result = append(result, padding+line)
		}
		return result
	}

	if node.Right == nil {
		leftLines := buildFancyLines(node.Left)
		leftWidth := len(leftLines[0])
		firstLine := strings.Repeat(" ", leftWidth) + valStr
		secondLine := strings.Repeat(" ", leftWidth-1) + "/" + strings.Repeat(" ", len(valStr))
		result := []string{firstLine, secondLine}
		padding := strings.Repeat(" ", 0)
		for _, line := range leftLines {
			result = append(result, padding+line)
		}
		return result
	}

	leftLines := buildFancyLines(node.Left)
	rightLines := buildFancyLines(node.Right)

	leftWidth := len(leftLines[0])
	rightWidth := len(rightLines[0])

	// Build first line: spaces + root value centered above
	firstLine := strings.Repeat(" ", leftWidth) + valStr + strings.Repeat(" ", rightWidth)

	// Build connector line with / and \
	secondLine := strings.Repeat(" ", leftWidth-1) + "/" + strings.Repeat(" ", len(valStr)) + "\\" + strings.Repeat(" ", rightWidth-1)

	result := []string{firstLine, secondLine}

	// Merge left and right subtree lines side by side
	leftLen := len(leftLines)
	rightLen := len(rightLines)
	maxLen := leftLen
	if rightLen > maxLen {
		maxLen = rightLen
	}

	gap := strings.Repeat(" ", len(valStr))
	for i := 0; i < maxLen; i++ {
		left := strings.Repeat(" ", leftWidth)
		if i < leftLen {
			left = leftLines[i]
		}
		right := strings.Repeat(" ", rightWidth)
		if i < rightLen {
			right = rightLines[i]
		}
		result = append(result, left+gap+right)
	}

	return result
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
