package main

import (
	"fmt"
)

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
	Cap   int
	Cache map[int]*ListNode
	Head  *ListNode
	Tail  *ListNode
}

func main() {
	testCommands := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	testValues := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	var obj LRUCache
	var output []interface{}

	for i, cmd := range testCommands {
		switch cmd {
		case "LRUCache":
			obj = Constructor(testValues[i][0])
			output = append(output, nil)
		case "put":
			obj.Put(testValues[i][0], testValues[i][1])
			output = append(output, nil)
		case "get":
			val := obj.Get(testValues[i][0])
			output = append(output, val)
		}
	}

	fmt.Println(output)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*ListNode),
		Head:  nil,
		Tail:  nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.Cache[key]; exists {
		if node == this.Head {
			return node.Val
		}
		if node == this.Tail {
			this.Tail = node.Prev
			this.Tail.Next = nil
		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}

		moveToHead(node, this)

		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.Cache[key]; exists {
		node.Val = value
		this.Get(key)
		return
	}

    if len(this.Cache) >= this.Cap {
        delete(this.Cache, this.Tail.Key)
        if this.Head == this.Tail {
            this.Head = nil
            this.Tail = nil
        } else {
            this.Tail = this.Tail.Prev
            this.Tail.Next = nil
        }
    }
	node := &ListNode{Key: key, Val: value}
	moveToHead(node, this)
	this.Cache[key] = node
}

func moveToHead(node *ListNode, this *LRUCache) {
    if this.Head == nil {
        this.Head = node
        this.Tail = node
    } else {
        node.Next = this.Head
        this.Head.Prev = node
        this.Head = node
    }
}
