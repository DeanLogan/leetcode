package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
						Children: nil,
					},
					{
						Val: 6,
						Children: nil,
					},
				},
			},
			{
				Val: 2,
				Children: nil,
			},
			{
				Val: 4,
				Children: nil,
			},
		},
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    max := 0
    for _, child := range root.Children {
        depth := maxDepth(child)
        if max < depth {
            max = depth
        }
    }

    return max + 1
}