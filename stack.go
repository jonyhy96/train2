//栈的顺序存储结构实现 vsersion 1.0
//func init_node(size int) *stack  定义栈
//func empty(stack *stack) bool	判断栈是否为空
//func push(stack *stack, data string)	入栈
//func pop(stack *stack) string	出栈
//func peek(stack *stack) string	返回栈顶但不删除栈顶
//func search(stack *stack, data string) bool	查找一个元素是否在栈内
package main

import (
	"fmt"
)

type stack struct {
	data    []string
	top     int
	maxsize int
}

func main() {
	stack := init_node(3)
	push(stack, "first")
	push(stack, "second")
	push(stack, "third")
	push(stack, "fourth")
	push(stack, "first")
	fmt.Println(stack)
	//println(pop(stack))
	//fmt.Println(stack)
	println(search(stack, "fourth"))
}
func init_node(size int) *stack {
	var nodestack stack
	nodestack.data = nil
	nodestack.top = -1
	nodestack.maxsize = size
	return &nodestack
}
func empty(stack *stack) bool {
	if stack.top == -1 {
		return true
	} else {
		return false
	}
}
func push(stack *stack, data string) {
	if stack.top == stack.maxsize-1 {
		fmt.Println("the stack is full!!")
	} else {
		stack.data = append(stack.data, data)
		stack.top++
	}
}
func pop(stack *stack) string {
	var e string
	top := stack.top
	if top == -1 {
		return "This stack is null"
	} else {
		e = stack.data[top]
		stack.data[top] = ""
		top--
		return e
	}
}
func peek(stack *stack) string {
	var e string
	top := stack.top
	if top == -1 {
		return "This stack is null"
	} else {
		e = stack.data[top]
		return e
	}
}
func search(stack *stack, data string) bool {
	top := stack.top
	b := false
	i := 0
	for i < top+1 {
		if data == stack.data[i] {
			b = true
			break
		} else {
			i++
		}
	}
	if i == top+1 {
		b = false
	}

	return b
}
