package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(evalRPN([]string{"2","1","+","3","*"}))
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
	fmt.Println(evalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}

func evalRPN(tokens []string) int {
	var stack []int
	var currentOperand, previousOperand int
	for _, token := range tokens {
		switch token {
		case "+":
			currentOperand, previousOperand, stack = getAndPopLastOperand(stack)
			stack = append(stack, (currentOperand + previousOperand))
		case "-":
			currentOperand, previousOperand, stack = getAndPopLastOperand(stack)
			stack = append(stack, (currentOperand - previousOperand))
		case "*":
			currentOperand, previousOperand, stack = getAndPopLastOperand(stack)
			stack = append(stack, (currentOperand * previousOperand))
		case "/":
			currentOperand, previousOperand, stack = getAndPopLastOperand(stack)
			stack = append(stack, (currentOperand / previousOperand))
		default:
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	return stack[0]
}

func getAndPopLastOperand(stack []int) (int, int, []int) {
	currentOperand := stack[len(stack)-2]
	previousOperand := stack[len(stack)-1]
	stack = stack[:len(stack)-2]

	return currentOperand, previousOperand, stack
}
