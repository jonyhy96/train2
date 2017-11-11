//Hashmap version 1.0

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const hashsize = 10

type node struct {
	key   int
	value string
	next  *node
}
type HashMap struct {
	size int
	list []node
}

func main() {
	hm := create_HashMap()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		hm.insert(r.Intn(10), "jony")
	}
	hm.print_hashmap()
	key := r.Intn(10)
	hm.search(key)
	hm.delete(key)
	hm.search(key)
	hm.print_hashmap()
}

func (hm *HashMap) hash(key int) int {
	return key % hashsize
}
func create_HashMap() HashMap {
	hm := HashMap{}
	hm.size = 0
	hm.list = make([]node, hashsize)
	return hm
}
func (hm *HashMap) insert(key int, data string) {
	newnode := node{}
	newnode.key = key
	newnode.value = data
	index := hm.hash(key)
	p := &hm.list[index]
	for p.next != nil {
		p = p.next
	}
	//fmt.Println("start to insert data")
	p.next = &newnode
	newnode.next = nil
	hm.size++
}
func (hm *HashMap) delete(key int) {
	index := hm.hash(key)
	p := &hm.list[index]
	for p.next != nil {
		if p.next.key == key {
			fmt.Println("key fond,prepare to delete")
			p.next = p.next.next
			hm.size--
			fmt.Println("delete successful!")
			return
		} else {
			p = p.next
		}
	}
}
func (hm *HashMap) search(key int) {
	index := hm.hash(key)
	p := &hm.list[index]
	for p.next != nil {
		if p.next.key == key {
			fmt.Printf("key %d fond, the value is %s  \n", p.next.key, p.next.value)
			return
		} else {
			p = p.next
		}
	}
	fmt.Println("key is not in the hashmap")
}
func (hm *HashMap) print_hashmap() {
	oldsize := hm.size
	var p []*node = make([]*node, hashsize)
	for i := 0; i < hashsize; i++ {
		p[i] = &hm.list[i]
	}
	for oldsize > 0 {
	LOOP:
		for i := 0; i < hashsize; i++ {
			if p[i].next != nil && p[i].next.value != "" {
				fmt.Printf("%2d : %4s  ", p[i].next.key, p[i].next.value)
				p[i] = p[i].next
				oldsize--
				continue
			} else {
				fmt.Printf("   :       ")
				continue
			}
		}
		fmt.Println()
		fmt.Printf("============================================================================================================")
		fmt.Println()
		if oldsize > 0 {
			goto LOOP
		}
	}
}
