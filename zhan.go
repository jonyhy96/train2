//栈的顺序存储结构实现 vsersion 1.0
//func init_node(size int) *node  定义栈
//func empty(zhan *node) bool	判断栈是否为空
//func push(zhan *node, data string)	入栈
//func pop(zhan *node) string	出栈
//func peek(zhan *node) string	返回栈顶但不删除栈顶
//func search(zhan *node, data string) bool	查找一个元素是否在栈内
package main

import (
	"fmt"
)

type node struct {
	data    []string
	top     int
	maxsize int
}

func main() {
	zhan := init_node(3)
	push(zhan, "first")
	push(zhan, "second")
	push(zhan, "third")
	push(zhan, "fourth")
	push(zhan, "first")
	fmt.Println(zhan)
	//println(pop(zhan))
	//fmt.Println(zhan)
	println(search(zhan, "fourth"))
}
func init_node(size int) *node {
	var nodezhan node
	nodezhan.data = nil
	nodezhan.top = -1
	nodezhan.maxsize = size
	return &nodezhan
}
func empty(zhan *node) bool {
	if zhan.top == -1 {
		return true
	} else {
		return false
	}
}
func push(zhan *node, data string) {
	if zhan.top == zhan.maxsize-1 {
		fmt.Println("the zhan is full!!")
	} else {
		zhan.data = append(zhan.data, data)
		zhan.top++
	}
}
func pop(zhan *node) string {
	var e string
	top := zhan.top
	if top == -1 {
		return "This zhan is null"
	} else {
		e = zhan.data[top]
		zhan.data[top] = ""
		top--
		return e
	}
}
func peek(zhan *node) string {
	var e string
	top := zhan.top
	if top == -1 {
		return "This zhan is null"
	} else {
		e = zhan.data[top]
		return e
	}
}
func search(zhan *node, data string) bool {
	top := zhan.top
	b := false
	i := 0
	for i < top+1 {
		if data == zhan.data[i] {
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
