package main

import (
	"fmt"
)

func main() {
    obj := Constructor()

    testCommands := []string{"push", "push", "push", "getMin", "pop", "top", "getMin"}
    testValues := [][]int{{-2}, {0}, {-3}, {}, {}, {}, {}}

    var output []interface{}

    for i, cmd := range testCommands {
        switch cmd {
        case "push":
            obj.Push(testValues[i][0])
            output = append(output, nil)
        case "pop":
            obj.Pop()
            output = append(output, nil)
        case "top":
            top := obj.Top()
            output = append(output, top)
        case "getMin":
            min := obj.GetMin()
            output = append(output, min)
        }
    }

    fmt.Println(output)
}

type MinStack struct {
    top *StackNode
    min int
}

type StackNode struct {
	data int
	next *StackNode
    lastmin int
}

var mystack MinStack = MinStack{top: nil}
var newtop *StackNode

func Constructor() MinStack {
    return mystack
}

func (this *MinStack) Push(val int)  {
    if this.top == nil {
        newtop = &StackNode{data: val, next: this.top}
        this.min = val
    } else {
        newtop = &StackNode{data: val, next: this.top, lastmin: this.min}
    }
    this.top = newtop
    if this.top.data < this.min {
        this.min = this.top.data
    }
}


func (this *MinStack) Pop()  {
    if this.top.next == nil {
        this.top = nil
        return
    } 
    this.min = this.top.lastmin
    *this.top = *this.top.next
}


func (this *MinStack) Top() int {
    return this.top.data
}


func (this *MinStack) GetMin() int {
    return this.min; 
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */