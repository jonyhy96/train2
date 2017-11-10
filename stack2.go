//stack 1.0
//func init_stack() 初始化一个栈
//func push(data string) 入栈
//func pop() 出栈
//func peek() 返回栈顶元素
//func empty() string 判断栈是否为空
//func loop() 从栈顶到栈底输出栈

package main

import "fmt"

type stack struct {
	data string
	next *stack
}

var size int
var top *stack

func main() {
	init_stack()
	push("a")
	peek()
	pop()
	loop()
}

func init_stack() {
	var basenode stack
	basenode.data = ""
	basenode.next = nil
	size = 0
	top = &basenode
}
func push(data string) {
	var newnode stack
	newnode.data = data
	newnode.next = top
	top = &newnode
	size++
}
func pop() {
	if size == 0 {
		fmt.Println("The stack is null")
	} else {
		data := top.data
		top = top.next
		size--
		fmt.Println(data)
	}
}
func peek() {
	if size == 0 {
		fmt.Println("The stack is null")
	} else {
		fmt.Println(top.data)
	}
}
func empty() string {
	if top.next == nil {
		return "The starck is null"
	} else {
		return "The stark is not null"
	}
}
func loop() {
	if size == 0 {
		fmt.Println("The stack is null")
	} else {
		var currentnode *stack
		currentnode = top
		i := size
		for size > 0 {
			fmt.Println(currentnode.data)
			currentnode = currentnode.next
			size--
		}
		size = i
	}
}
