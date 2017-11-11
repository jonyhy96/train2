package main

import "fmt"
import "sync"

type singleton struct {
	a int
}

func (s singleton) f() {
	fmt.Println(s.a)
}

var once sync.Once

func getinstance(a int) *singleton {
	var instance *singleton
	once.Do(func() {
		instance = &singleton{a}
	})
	return instance
}
func main() {
	instance := getinstance(3)
	instance1 := getinstance(4)
	instance.f()
	instance1.f()
}
