package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	str  string
	next *Node
}

type Stack struct {
	head *Node
}

func createStack() *Stack {
	return &Stack{}
}

func (stack *Stack) PushBack(arg string) {
	node := &Node{str: arg}
	if stack.head == nil {
		stack.head = node
		return
	}

	iter := stack.head
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = node
}

func (stack *Stack) PopFront() string {
	iter := stack.head
	if stack.head == nil {
		return ""
	}
	if stack.head.next == nil {
		stack.head = nil
		return iter.str
	}
	var prev *Node
	for iter.next != nil {
		prev = iter
		iter = iter.next
	}
	prev.next = nil
	return iter.str
}

func evalRPN(tokens []string) int {
	size := len(tokens)
	stack := createStack()
	for i := 0; i != size; i++ {
		switch tokens[i] {
		case "+":
			num1, err1 := strconv.Atoi(stack.PopFront())
			num2, err2 := strconv.Atoi(stack.PopFront())
			if err1 == nil && err2 == nil {
				stack.PushBack(strconv.Itoa(num1 + num2))
			}
		case "-":
			num1, err1 := strconv.Atoi(stack.PopFront())
			num2, err2 := strconv.Atoi(stack.PopFront())
			if err1 == nil && err2 == nil {
				stack.PushBack(strconv.Itoa(num2 - num1))
			}
		case "*":
			num1, err1 := strconv.Atoi(stack.PopFront())
			num2, err2 := strconv.Atoi(stack.PopFront())
			if err1 == nil && err2 == nil {
				stack.PushBack(strconv.Itoa(num1 * num2))
			}
		case "/":
			num1, err1 := strconv.Atoi(stack.PopFront())
			num2, err2 := strconv.Atoi(stack.PopFront())
			if err1 == nil && err2 == nil {
				stack.PushBack(strconv.Itoa(num2 / num1))
			}
		default:
			stack.PushBack(tokens[i])
		}
	}
	ans, e := strconv.Atoi(stack.PopFront())
	if e == nil {
		return ans
	}
	return 1
}

func main() {
	fmt.Print(evalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}
