//linklist version 1.0

package main

import (
	"fmt"
	"strconv"
)

type node struct {
	data []string
	next *node
}

func main() {
	node := create_firstnode()       //create the first node
	add_node(&node, 4, "hao")        //add nodes to the first node
	insert_node(&node, 2, "second")  //insert a node to the linklist
	delete_ndoe(&node, 2)            //delete a node from the link list
	channge_node(&node, 2, "second") //change a node from the linklist
	print_node(&node)                //print out the link list
	search_node(&node, "seco")       //check wether the data is in the list
}
func create_firstnode() node {
	var firstnode node
	firstnode.data = append(firstnode.data, "first")
	firstnode.next = nil
	return firstnode
}
func add_node(currentnode *node, len int, data string) {
	for i := 0; i < len; i++ {
		var node1 node
		node1.data = append(node1.data, data)
		node1.next = nil
		currentnode.next = &node1
		currentnode = &node1
	}
}
func insert_node(currentnode *node, index int, data string) {
	pod := 0
	index--
	for currentnode.next != nil {
		for pod < index {
			pod++
			currentnode = currentnode.next
		}
		var newnode node
		newnode.data = append(newnode.data, data)
		newnode.next = currentnode.next
		currentnode.next = &newnode
		break
	}
}
func print_node(currentnode *node) {
	for currentnode.next != nil {
		fmt.Println(currentnode.data)
		currentnode = currentnode.next
	}
}
func channge_node(currentnode *node, index int, data string) {
	pod := 0
	for currentnode.next != nil {
		for pod < index {
			pod++
			currentnode = currentnode.next
		}
		fmt.Println("The", index, "node befor change", currentnode.data)
		var dataslip []string
		//currentnode.data = append(currentnode.data, data)// this can not change it but append a data after the old data
		//var dataslip = make([]string, len(data)) //this will return a slip but with len(data) size so its not nil
		dataslip = append(dataslip, data)
		//fmt.Println("dataslip", dataslip)
		copy(currentnode.data, dataslip)
		fmt.Println("The", index, "node after done", currentnode.data)
		break
	}
}
func delete_ndoe(currentnode *node, index int) {
	currentnode = loop_node(currentnode, index)
	currentnode.data = currentnode.next.data
	currentnode.next = currentnode.next.next
}
func loop_node(currentnode *node, index int) *node {
	pod := 0
	for currentnode.next != nil {
		for pod < index {
			pod++
			currentnode = currentnode.next
		}
		break
	}
	return currentnode
}
func search_node(currentnode *node, data string) {
	index := 0
	var olddata string
	for currentnode.next != nil {
		olddata = currentnode.data[0]
		if olddata == data {
			string := strconv.Itoa(index)
			fmt.Println("find! the index is " + string)
			break
		} else {
			currentnode = currentnode.next
			index++
		}
	}
	if currentnode.next == nil {
		fmt.Println("the data is not in the list")
	}
}
