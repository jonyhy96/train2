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
var instance *singleton

func getinstance(a int) *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{a}
		}
	})
	return instance
}
func main() {
	instance := getinstance(3)
	instance1 := getinstance(5555)
	instance.f()
	instance1.f()
}
